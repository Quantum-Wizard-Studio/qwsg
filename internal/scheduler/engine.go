package scheduler

import (
	"fmt"
	"sort"
	"time"

	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/policy"
)

type TimeZoneResolver interface {
	Resolve(string) (*time.Location, error)
}
type SystemTimeZones struct{}

func (SystemTimeZones) Resolve(name string) (*time.Location, error) { return time.LoadLocation(name) }

type candidate struct {
	schedule   configuration.Schedule
	occurrence Occurrence
	attempt    int
	retry      bool
}

// Evaluate is the pure deterministic Scheduler Engine. The half-open window is
// (last evaluation, observation]. First evaluation anchors interval schedules
// at the observation and intentionally creates no immediate occurrence.
func Evaluate(config configuration.Effective, state State, observation ClockObservation, zones TimeZoneResolver) (Evaluation, error) {
	if err := configuration.ValidateEffective(config); err != nil {
		return Evaluation{}, fmt.Errorf("invalid effective configuration: %w", err)
	}
	if observation.WallTime.IsZero() || observation.SessionID == "" || observation.MonotonicNS < 0 {
		return Evaluation{}, fmt.Errorf("invalid clock observation")
	}
	observation.WallTime = normalizeTime(observation.WallTime)
	if zones == nil {
		return Evaluation{}, fmt.Errorf("time-zone resolver is required")
	}
	if state.SchemaName == "" {
		state = NewState(config.ID)
	} else if err := ValidateState(state); err != nil {
		return Evaluation{}, err
	}
	if state.ConfigurationID != config.ID {
		return Evaluation{}, fmt.Errorf("scheduler state configuration identity differs from effective configuration")
	}
	state = cloneState(state)
	evaluation := Evaluation{SchemaName: EvaluationSchema, SchemaVersion: SchemaVersion, EngineVersion: EngineVersion, TaxonomyVersion: TaxonomyVersion, ConfigurationID: config.ID, EvaluatedAt: observation.WallTime, Records: []Record{}, Requests: []Request{}, Events: []Event{}}

	first := state.LastWallTime.IsZero()
	restart := !first && state.SessionID != observation.SessionID
	discontinuity := false
	if !first && !restart {
		wallDelta := observation.WallTime.Sub(state.LastWallTime)
		monoDelta := time.Duration(observation.MonotonicNS - state.LastMonotonicNS)
		if wallDelta < 0 || monoDelta < 0 || absDuration(wallDelta-monoDelta) > ClockTolerance {
			discontinuity = true
			evaluation.Events = append(evaluation.Events, newEvent(EventClockDiscontinuity, observation.WallTime, "", "", "wall_and_monotonic_clock_diverged", map[string]string{}))
		}
	}
	if restart {
		for i := range state.Schedules {
			for _, active := range state.Schedules[i].Active {
				result := ExecutionResult{SchemaName: ResultSchema, SchemaVersion: SchemaVersion, RequestID: active.RequestID, ScheduleID: state.Schedules[i].ScheduleID, OccurrenceID: active.OccurrenceID, ScheduledAt: active.ScheduledAt, Attempt: active.Attempt, Outcome: ExecutionInterrupted, StartedAt: active.ReservedAt, CompletedAt: observation.WallTime, StageContracts: []string{}, PolicyEvaluationIDs: []string{}, PolicyOutcomes: []string{}, FailureCode: "scheduler_restart"}
				if schedule, ok := scheduleByID(config, state.Schedules[i].ScheduleID); ok {
					if retry, found := retryPolicy(config, schedule.RetryPolicyID); found && active.Attempt < retry.MaxAttempts {
						result.NextRetryAt = observation.WallTime.Add(retryDelay(retry, active.Attempt))
					}
				}
				result.ID = resultID(result)
				state.Results = append(state.Results, result)
			}
			state.Schedules[i].Active = []ActiveRun{}
		}
		evaluation.Events = append(evaluation.Events, newEvent(EventRestartRecovered, observation.WallTime, "", "", "in_progress_requests_marked_interrupted", map[string]string{}))
	}

	candidates := []candidate{}
	for _, schedule := range config.Values.Schedules {
		item := findScheduleState(&state, schedule.ID)
		if item.AnchorAt.IsZero() {
			item.AnchorAt = observation.WallTime
			item.LastEvaluatedAt = observation.WallTime
			evaluation.Events = append(evaluation.Events, newEvent(EventInitialized, observation.WallTime, schedule.ID, "", "schedule_state_initialized", map[string]string{}))
		}
		record := Record{ScheduleID: schedule.ID, OccurrenceIDs: []string{}, RequestIDs: []string{}, ConfigurationID: config.ID, EvaluatedAt: observation.WallTime}
		if !schedule.Enabled {
			record.Decision, record.Reason = DecisionDisabled, "schedule_disabled"
			item.NextRunAt = time.Time{}
			record.ID = recordID(record)
			evaluation.Records = append(evaluation.Records, record)
			continue
		}
		if len(schedule.CheckIDs) > 0 {
			record.Decision, record.Reason = DecisionInapplicable, "command_1.0_cannot_represent_check_scope"
			item.NextRunAt = nextRun(schedule, *item, observation.WallTime, zones)
			record.NextRunAt = item.NextRunAt
			record.ID = recordID(record)
			evaluation.Records = append(evaluation.Records, record)
			continue
		}
		if discontinuity {
			record.Decision, record.Reason = DecisionIndeterminate, "clock_discontinuity"
			item.LastEvaluatedAt = observation.WallTime
			item.NextRunAt = nextRun(schedule, *item, observation.WallTime, zones)
			record.NextRunAt = item.NextRunAt
			record.ID = recordID(record)
			evaluation.Records = append(evaluation.Records, record)
			continue
		}

		retryDue := false
		if retry, ok := retryCandidate(config, state, schedule, observation.WallTime); ok {
			candidates = append(candidates, retry)
			retryDue = true
			record.Decision, record.Reason = DecisionDue, "retry_eligible"
			record.OccurrenceIDs = []string{retry.occurrence.ID}
		}
		if item.Pending != nil && !retryDue {
			candidates = append(candidates, candidate{schedule: schedule, occurrence: *item.Pending, attempt: 1})
			record.Decision, record.Reason = DecisionDue, "pending_occurrence_eligible"
			record.OccurrenceIDs = append(record.OccurrenceIDs, item.Pending.ID)
		}
		occurrences, err := dueOccurrences(schedule, *item, observation.WallTime, zones)
		if err != nil {
			record.Decision, record.Reason = DecisionIndeterminate, "time_zone_unavailable"
			item.LastEvaluatedAt = observation.WallTime
			record.ID = recordID(record)
			evaluation.Records = append(evaluation.Records, record)
			continue
		}
		if len(occurrences) > 0 {
			for _, occ := range occurrences {
				record.OccurrenceIDs = append(record.OccurrenceIDs, occ.ID)
			}
			selected, decision, reason := selectMissed(schedule, occurrences)
			record.Decision, record.Reason = decision, reason
			if selected != nil {
				candidates = append(candidates, candidate{schedule: schedule, occurrence: *selected, attempt: 1})
			}
			item.LastScheduledAt = occurrences[len(occurrences)-1].ScheduledAt
		}
		if len(occurrences) == 0 && record.Decision == "" {
			record.Decision, record.Reason = DecisionNotDue, "no_occurrence_in_window"
		}
		item.LastEvaluatedAt = observation.WallTime
		item.NextRunAt = nextRun(schedule, *item, observation.WallTime, zones)
		record.NextRunAt = item.NextRunAt
		record.ID = recordID(record)
		evaluation.Records = append(evaluation.Records, record)
	}

	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].retry != candidates[j].retry {
			return candidates[i].retry
		}
		if candidates[i].schedule.Priority != candidates[j].schedule.Priority {
			return candidates[i].schedule.Priority > candidates[j].schedule.Priority
		}
		if !candidates[i].occurrence.ScheduledAt.Equal(candidates[j].occurrence.ScheduledAt) {
			return candidates[i].occurrence.ScheduledAt.Before(candidates[j].occurrence.ScheduledAt)
		}
		return candidates[i].schedule.ID < candidates[j].schedule.ID
	})
	active := 0
	for _, item := range state.Schedules {
		active += len(item.Active)
	}
	available := config.Values.Concurrency - active
	if available < 0 {
		available = 0
	}
	reserved := map[string]bool{}
	for _, value := range candidates {
		item := findScheduleState(&state, value.schedule.ID)
		record := findRecord(evaluation.Records, value.schedule.ID)
		if value.schedule.OverlapPolicy == configuration.OverlapForbid && (len(item.Active) > 0 || reserved[value.schedule.ID]) {
			if value.schedule.MisfirePolicy == configuration.MisfireRunOnce {
				copy := value.occurrence
				item.Pending = &copy
				record.Decision, record.Reason = DecisionQueued, "overlap_coalesced_once"
			} else if value.schedule.MisfirePolicy == configuration.MisfireIndeterminate {
				record.Decision, record.Reason = DecisionIndeterminate, "overlap_indeterminate"
			} else {
				record.Decision, record.Reason = DecisionSkipped, "overlap_skipped"
			}
			continue
		}
		if available == 0 {
			copy := value.occurrence
			item.Pending = &copy
			record.Decision, record.Reason = DecisionDelayed, "concurrency_capacity_exhausted"
			continue
		}
		request := Request{SchemaName: RequestSchema, SchemaVersion: SchemaVersion, ConfigurationID: config.ID, ScheduleID: value.schedule.ID, OccurrenceID: value.occurrence.ID, ScheduledAt: value.occurrence.ScheduledAt, Attempt: value.attempt, Priority: value.schedule.Priority, CommandProfile: value.schedule.CommandProfile, CheckIDs: append([]string{}, value.schedule.CheckIDs...), ExecutionTimeoutNS: value.schedule.ExecutionTimeoutNS, RetryPolicyID: value.schedule.RetryPolicyID}
		evaluation.Requests = append(evaluation.Requests, request)
		reserved[value.schedule.ID] = true
		available--
	}
	sort.Slice(evaluation.Records, func(i, j int) bool { return evaluation.Records[i].ScheduleID < evaluation.Records[j].ScheduleID })
	evaluation.ID = evaluationID(evaluation)
	for i := range evaluation.Requests {
		evaluation.Requests[i].EvaluationID = evaluation.ID
		evaluation.Requests[i].ID = requestID(evaluation.Requests[i])
		item := findScheduleState(&state, evaluation.Requests[i].ScheduleID)
		if item.Pending != nil && item.Pending.ID == evaluation.Requests[i].OccurrenceID {
			item.Pending = nil
		}
		item.Active = append(item.Active, ActiveRun{RequestID: evaluation.Requests[i].ID, OccurrenceID: evaluation.Requests[i].OccurrenceID, ScheduledAt: evaluation.Requests[i].ScheduledAt, Attempt: evaluation.Requests[i].Attempt, ReservedAt: observation.WallTime})
		record := findRecord(evaluation.Records, evaluation.Requests[i].ScheduleID)
		record.RequestIDs = append(record.RequestIDs, evaluation.Requests[i].ID)
		record.Decision, record.Reason = DecisionDue, "execution_request_created"
		evaluation.Events = append(evaluation.Events, newEvent(EventRequestReserved, observation.WallTime, evaluation.Requests[i].ScheduleID, evaluation.Requests[i].ID, "execution_request_reserved", map[string]string{}))
	}
	for i := range evaluation.Records {
		sort.Strings(evaluation.Records[i].OccurrenceIDs)
		evaluation.Records[i].OccurrenceIDs = unique(evaluation.Records[i].OccurrenceIDs)
		sort.Strings(evaluation.Records[i].RequestIDs)
		evaluation.Records[i].ID = recordID(evaluation.Records[i])
	}
	state.SessionID = observation.SessionID
	state.LastWallTime = observation.WallTime
	state.LastMonotonicNS = observation.MonotonicNS
	sortState(&state)
	evaluation.NextState = state
	evaluation.ID = evaluationID(evaluation)
	// Evaluation identity excludes derived cross-reference IDs. Rebind requests
	// once to the final identity and then stabilize records and state.
	for i := range evaluation.Requests {
		evaluation.Requests[i].EvaluationID = evaluation.ID
		evaluation.Requests[i].ID = requestID(evaluation.Requests[i])
	}
	rebindActive(&evaluation)
	rebindEvents(&evaluation)
	for i := range evaluation.Records {
		evaluation.Records[i].ID = recordID(evaluation.Records[i])
	}
	sortState(&evaluation.NextState)
	evaluation.ID = evaluationID(evaluation)
	for i := range evaluation.Requests {
		evaluation.Requests[i].EvaluationID = evaluation.ID
		evaluation.Requests[i].ID = requestID(evaluation.Requests[i])
	}
	rebindActive(&evaluation)
	rebindEvents(&evaluation)
	sortState(&evaluation.NextState)
	for i := range evaluation.Records {
		evaluation.Records[i].ID = recordID(evaluation.Records[i])
	}
	for i := range evaluation.Events {
		evaluation.Events[i].ID = eventID(evaluation.Events[i])
	}
	if err := ValidateEvaluation(evaluation); err != nil {
		return Evaluation{}, err
	}
	return evaluation, nil
}

func ApplyCompletions(config configuration.Effective, state State, completions []Completion) (State, []ExecutionResult, []Event, error) {
	if err := configuration.ValidateEffective(config); err != nil {
		return State{}, nil, nil, err
	}
	if err := ValidateState(state); err != nil {
		return State{}, nil, nil, err
	}
	state = cloneState(state)
	sort.Slice(completions, func(i, j int) bool { return completions[i].Request.ID < completions[j].Request.ID })
	results := []ExecutionResult{}
	events := []Event{}
	for _, completion := range completions {
		if err := validateRequest(completion.Request); err != nil {
			return State{}, nil, nil, err
		}
		item := findScheduleState(&state, completion.Request.ScheduleID)
		found := false
		for i, active := range item.Active {
			if active.RequestID == completion.Request.ID {
				item.Active = append(item.Active[:i], item.Active[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return State{}, nil, nil, fmt.Errorf("completion references a non-active request")
		}
		outcome := ExecutionSucceeded
		if completion.FailureCode != "" {
			outcome = ExecutionFailed
		} else if !completion.Execution.Complete {
			outcome = ExecutionIncomplete
		}
		result := ExecutionResult{SchemaName: ResultSchema, SchemaVersion: SchemaVersion, RequestID: completion.Request.ID, ScheduleID: completion.Request.ScheduleID, OccurrenceID: completion.Request.OccurrenceID, ScheduledAt: completion.Request.ScheduledAt, Attempt: completion.Request.Attempt, Outcome: outcome, StartedAt: normalizeTime(completion.StartedAt), CompletedAt: normalizeTime(completion.CompletedAt), CommandExecutionID: completion.Execution.ID, CommandComplete: completion.Execution.Complete, StageContracts: []string{}, PolicyEvaluationIDs: []string{}, PolicyOutcomes: []string{}, FailureCode: completion.FailureCode}
		captureExecution(&result, completion.Execution)
		if outcome != ExecutionSucceeded {
			if retry, ok := retryPolicy(config, completion.Request.RetryPolicyID); ok && completion.Request.Attempt < retry.MaxAttempts {
				delay := retryDelay(retry, completion.Request.Attempt)
				result.NextRetryAt = result.CompletedAt.Add(delay)
			}
		}
		sort.Strings(result.StageContracts)
		sort.Strings(result.PolicyEvaluationIDs)
		sort.Strings(result.PolicyOutcomes)
		result.StageContracts = unique(result.StageContracts)
		result.PolicyEvaluationIDs = unique(result.PolicyEvaluationIDs)
		result.PolicyOutcomes = unique(result.PolicyOutcomes)
		result.ID = resultID(result)
		results = append(results, result)
		state.Results = append(state.Results, result)
		events = append(events, newEvent(EventExecutionCompleted, result.CompletedAt, result.ScheduleID, result.RequestID, "command_execution_recorded", map[string]string{"outcome": string(result.Outcome)}))
	}
	sortState(&state)
	return state, results, events, nil
}

func captureExecution(result *ExecutionResult, execution command.Execution) {
	for _, stage := range execution.Stages {
		result.StageContracts = append(result.StageContracts, string(stage.Stage)+":"+stage.ContractName+"/"+stage.Version)
		if governed, ok := stage.Value.(policy.Result); ok {
			for _, record := range governed.Records {
				result.PolicyEvaluationIDs = append(result.PolicyEvaluationIDs, record.ID)
				result.PolicyOutcomes = append(result.PolicyOutcomes, string(record.Outcome))
			}
		}
	}
}

func dueOccurrences(schedule configuration.Schedule, state ScheduleState, now time.Time, zones TimeZoneResolver) ([]Occurrence, error) {
	start := state.LastEvaluatedAt
	if start.IsZero() {
		start = state.AnchorAt
	}
	if !now.After(start) {
		return []Occurrence{}, nil
	}
	if schedule.Trigger == configuration.IntervalTrigger {
		base := state.LastScheduledAt
		if base.IsZero() {
			base = state.AnchorAt
		}
		values := []Occurrence{}
		for at := base.Add(time.Duration(schedule.IntervalNS)); !at.After(now) && len(values) < MaxOccurrences; at = at.Add(time.Duration(schedule.IntervalNS)) {
			if at.After(start) {
				values = append(values, newOccurrence(schedule, at))
			}
		}
		return values, nil
	}
	location, err := zones.Resolve(schedule.TimeZone)
	if err != nil {
		return nil, err
	}
	values := []Occurrence{}
	cursor := start.UTC().Truncate(time.Minute)
	limit := now
	if limit.Sub(cursor) > MaxLookahead {
		cursor = limit.Add(-MaxLookahead)
	}
	for at := cursor.Add(time.Minute); !at.After(limit) && len(values) < MaxOccurrences; at = at.Add(time.Minute) {
		local := at.In(location)
		if calendarMatches(schedule.Calendar, local) && dstSelected(schedule, at, location) {
			values = append(values, newOccurrence(schedule, at))
		}
	}
	return values, nil
}

func nextRun(schedule configuration.Schedule, state ScheduleState, now time.Time, zones TimeZoneResolver) time.Time {
	if !schedule.Enabled {
		return time.Time{}
	}
	if schedule.Trigger == configuration.IntervalTrigger {
		base := state.LastScheduledAt
		if base.IsZero() {
			base = state.AnchorAt
		}
		if base.IsZero() {
			base = now
		}
		for next := base.Add(time.Duration(schedule.IntervalNS)); ; next = next.Add(time.Duration(schedule.IntervalNS)) {
			if next.After(now) {
				return normalizeTime(next)
			}
		}
	}
	location, err := zones.Resolve(schedule.TimeZone)
	if err != nil {
		return time.Time{}
	}
	cursor := now.UTC().Truncate(time.Minute)
	for at := cursor.Add(time.Minute); at.Sub(cursor) <= MaxLookahead; at = at.Add(time.Minute) {
		if calendarMatches(schedule.Calendar, at.In(location)) && dstSelected(schedule, at, location) {
			return normalizeTime(at)
		}
	}
	return time.Time{}
}

func calendarMatches(value configuration.Calendar, local time.Time) bool {
	return containsInt(value.Minutes, local.Minute()) && containsInt(value.Hours, local.Hour()) && (len(value.MonthDays) == 0 || containsInt(value.MonthDays, local.Day())) && (len(value.Months) == 0 || containsInt(value.Months, int(local.Month()))) && (len(value.Weekdays) == 0 || containsInt(value.Weekdays, int(local.Weekday())))
}
func dstSelected(schedule configuration.Schedule, at time.Time, location *time.Location) bool {
	key := at.In(location).Format("2006-01-02T15:04")
	other := time.Time{}
	for _, delta := range []time.Duration{-2 * time.Hour, -time.Hour, time.Hour, 2 * time.Hour} {
		candidate := at.Add(delta)
		if candidate.In(location).Format("2006-01-02T15:04") == key {
			other = candidate
			break
		}
	}
	if other.IsZero() {
		return true
	}
	first := at.Before(other)
	switch schedule.DSTPolicy {
	case configuration.DSTSecondOccurrence:
		return !first
	case configuration.DSTFirstOccurrence, configuration.DSTSkipNonexistent:
		return first
	}
	return false
}
func newOccurrence(schedule configuration.Schedule, at time.Time) Occurrence {
	at = normalizeTime(at)
	trigger := string(schedule.Trigger)
	return Occurrence{ID: occurrenceID(schedule.ID, at, trigger), ScheduleID: schedule.ID, ScheduledAt: at, Trigger: trigger}
}
func selectMissed(schedule configuration.Schedule, values []Occurrence) (*Occurrence, Decision, string) {
	if len(values) == 1 {
		return &values[0], DecisionDue, "occurrence_due"
	}
	switch schedule.MisfirePolicy {
	case configuration.MisfireRunOnce:
		value := values[len(values)-1]
		return &value, DecisionDue, "missed_occurrences_coalesced"
	case configuration.MisfireSkip:
		return nil, DecisionSkipped, "missed_occurrences_skipped"
	default:
		return nil, DecisionIndeterminate, "missed_occurrences_indeterminate"
	}
}
func retryCandidate(config configuration.Effective, state State, schedule configuration.Schedule, now time.Time) (candidate, bool) {
	retry, ok := retryPolicy(config, schedule.RetryPolicyID)
	if !ok {
		return candidate{}, false
	}
	seenOccurrences := map[string]bool{}
	for i := len(state.Results) - 1; i >= 0; i-- {
		result := state.Results[i]
		if result.ScheduleID != schedule.ID || seenOccurrences[result.OccurrenceID] {
			continue
		}
		seenOccurrences[result.OccurrenceID] = true
		if result.Attempt >= retry.MaxAttempts || result.NextRetryAt.IsZero() || result.NextRetryAt.After(now) {
			continue
		}
		for _, item := range state.Schedules {
			for _, active := range item.Active {
				if active.OccurrenceID == result.OccurrenceID {
					return candidate{}, false
				}
			}
		}
		occ := Occurrence{ID: result.OccurrenceID, ScheduleID: schedule.ID, ScheduledAt: result.ScheduledAt, Trigger: "retry"}
		return candidate{schedule: schedule, occurrence: occ, attempt: result.Attempt + 1, retry: true}, true
	}
	return candidate{}, false
}
func scheduleByID(config configuration.Effective, id string) (configuration.Schedule, bool) {
	for _, schedule := range config.Values.Schedules {
		if schedule.ID == id {
			return schedule, true
		}
	}
	return configuration.Schedule{}, false
}
func retryDelay(value configuration.RetryPolicy, attempt int) time.Duration {
	delay := time.Duration(value.InitialDelayNS)
	for i := 1; i < attempt; i++ {
		if delay >= time.Duration(value.MaxDelayNS)/2 {
			return time.Duration(value.MaxDelayNS)
		}
		delay *= 2
	}
	if delay > time.Duration(value.MaxDelayNS) {
		delay = time.Duration(value.MaxDelayNS)
	}
	return delay
}
func findRecord(values []Record, id string) *Record {
	for i := range values {
		if values[i].ScheduleID == id {
			return &values[i]
		}
	}
	panic("record missing")
}
func rebindActive(value *Evaluation) {
	byOccurrence := map[string]Request{}
	for _, request := range value.Requests {
		byOccurrence[request.OccurrenceID] = request
	}
	for i := range value.NextState.Schedules {
		for j := range value.NextState.Schedules[i].Active {
			if request, ok := byOccurrence[value.NextState.Schedules[i].Active[j].OccurrenceID]; ok {
				value.NextState.Schedules[i].Active[j].RequestID = request.ID
			}
		}
	}
	for i := range value.Records {
		value.Records[i].RequestIDs = []string{}
		for _, request := range value.Requests {
			if request.ScheduleID == value.Records[i].ScheduleID {
				value.Records[i].RequestIDs = append(value.Records[i].RequestIDs, request.ID)
			}
		}
	}
}
func rebindEvents(value *Evaluation) {
	bySchedule := map[string]Request{}
	for _, request := range value.Requests {
		bySchedule[request.ScheduleID] = request
	}
	for i := range value.Events {
		if value.Events[i].Kind != EventRequestReserved {
			continue
		}
		if request, ok := bySchedule[value.Events[i].ScheduleID]; ok {
			value.Events[i].RequestID = request.ID
			value.Events[i].ID = eventID(value.Events[i])
		}
	}
}
func newEvent(kind EventKind, at time.Time, scheduleID, requestID, reason string, metadata map[string]string) Event {
	event := Event{SchemaName: EventSchema, SchemaVersion: SchemaVersion, Kind: kind, At: normalizeTime(at), ScheduleID: scheduleID, RequestID: requestID, Reason: reason, Metadata: metadata}
	event.ID = eventID(event)
	return event
}
func containsInt(values []int, wanted int) bool {
	index := sort.SearchInts(values, wanted)
	return index < len(values) && values[index] == wanted
}
func absDuration(value time.Duration) time.Duration {
	if value < 0 {
		return -value
	}
	return value
}
func unique(values []string) []string {
	if len(values) == 0 {
		return []string{}
	}
	result := values[:1]
	for _, value := range values[1:] {
		if value != result[len(result)-1] {
			result = append(result, value)
		}
	}
	return result
}
