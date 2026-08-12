package presentationmodel

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/command"
	"quantumwizard.hu/qwsg/internal/comparison"
	"quantumwizard.hu/qwsg/internal/drift"
	"quantumwizard.hu/qwsg/internal/health"
	"quantumwizard.hu/qwsg/internal/inventory"
	"quantumwizard.hu/qwsg/internal/policy"
	"quantumwizard.hu/qwsg/internal/report"
	"quantumwizard.hu/qwsg/internal/rule"
	"quantumwizard.hu/qwsg/internal/runtime"
	"quantumwizard.hu/qwsg/internal/runtimeservice"
)

var tokenPattern = regexp.MustCompile(`^[a-z0-9]+([._:-][a-z0-9]+)*$`)

func Project(input Input) (Overview, error) {
	if err := ValidateInput(input); err != nil {
		return Overview{}, err
	}
	b := builder{input: input, overview: Overview{
		SchemaName: SchemaName, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion,
		ObservedAt: input.ObservedAt.UTC(), Condition: Unknown, Attention: AttentionUnknown,
		Guardian: GuardianNotObserved, Freshness: FreshnessNotObserved,
		Completeness: CompletenessNotObserved, Changes: []ChangeSummary{},
		AttentionItems: []AttentionItem{}, Recommendations: []Recommendation{}, Sources: []SourceReference{},
	}}
	b.consume()
	b.finish()
	if err := Validate(b.overview); err != nil {
		return Overview{}, err
	}
	return b.overview, nil
}

// RequalifyFreshness ages a validated stored Overview without adding or
// reinterpreting evidence. The exclusive freshness deadline is supplied by the
// Current Operator State contract; only presentationmodel owns this transition.
func RequalifyFreshness(overview Overview, evaluatedAt, freshUntil time.Time) (Overview, error) {
	if err := Validate(overview); err != nil {
		return Overview{}, err
	}
	if !validTime(evaluatedAt) || !validTime(freshUntil) || freshUntil.Before(overview.ObservedAt) || evaluatedAt.Before(overview.ObservedAt) {
		return Overview{}, fmt.Errorf("invalid freshness requalification")
	}
	if evaluatedAt.Before(freshUntil) {
		return overview, nil
	}
	next := overview
	next.Freshness = FreshnessStale
	if next.Guardian != GuardianNotObserved {
		next.Guardian = GuardianUnavailable
	}
	if next.Completeness == CompletenessComplete {
		next.Completeness = CompletenessPartial
	}
	if next.Condition == Healthy || next.Condition == Unknown {
		next.Condition = Degraded
	}
	next.Recommendations = []Recommendation{}
	b := builder{overview: next}
	b.recommend()
	next = b.overview
	next.ID = overviewID(next)
	if err := Validate(next); err != nil {
		return Overview{}, err
	}
	return next, nil
}

// TransitionGuardian preserves qualified engineering facts while applying a
// current, externally correlated operational lifecycle observation.
func TransitionGuardian(overview Overview, guardian Guardian, observedAt time.Time) (Overview, error) {
	if err := Validate(overview); err != nil || !validTime(observedAt) || observedAt.Before(overview.ObservedAt) || (guardian != GuardianStopped && guardian != GuardianDegraded && guardian != GuardianUnavailable) {
		return Overview{}, fmt.Errorf("invalid Guardian transition")
	}
	overview.ObservedAt = observedAt
	overview.Guardian = guardian
	overview.Freshness = FreshnessCurrent
	if guardian != GuardianStopped || overview.Condition == Healthy {
		overview.Completeness = CompletenessPartial
		overview.Condition = Degraded
	}
	overview.Recommendations = []Recommendation{}
	b := builder{overview: overview}
	b.recommend()
	overview = b.overview
	overview.ID = overviewID(overview)
	return overview, Validate(overview)
}

type builder struct {
	input                       Input
	overview                    Overview
	stale, unsupported, partial bool
	attentionCandidates         []attentionCandidate
	policyRuleIDs               map[string]bool
}

type attentionCandidate struct {
	item       AttentionItem
	importance int
}

const (
	importanceRule = iota + 1
	importancePolicy
	importanceDirect
)

func (b *builder) source(kind, contract, version, id string, at time.Time) SourceReference {
	r := SourceReference{Kind: kind, Contract: contract, Version: version, RecordID: id, ObservedAt: at.UTC()}
	b.overview.Sources = append(b.overview.Sources, r)
	if b.input.ObservedAt.Sub(at) > time.Duration(b.input.FreshForNS) {
		b.stale = true
	}
	return r
}

func (b *builder) consume() {
	if v := b.input.Command; v != nil {
		b.source("command", v.Value.SchemaName, v.Value.SchemaVersion, v.Value.ID, v.ObservedAt)
	}
	if v := b.input.Inventory; v != nil {
		b.source("inventory", inventory.CanonicalSchemaName, inventory.SchemaVersion, v.Value.SnapshotID, v.ObservedAt)
		if v.Value.Status != inventory.Complete {
			b.partial = true
		}
	}
	if v := b.input.Comparison; v != nil {
		b.source("comparison", v.Value.SchemaName, v.Value.SchemaVersion, v.Value.ComparisonID, v.ObservedAt)
		b.overview.Summary.Changes = v.Value.Counts.Added + v.Value.Counts.Removed + v.Value.Counts.Modified
	}
	if v := b.input.Drift; v != nil {
		b.source("drift", v.Value.SchemaName, v.Value.SchemaVersion, envelopeID(v.Value), v.ObservedAt)
		by := map[string]*ChangeSummary{}
		for _, r := range v.Value.Records {
			key := string(r.Category)
			if by[key] == nil {
				by[key] = &ChangeSummary{Category: key}
			}
			switch r.Classification {
			case drift.PresenceAdded:
				by[key].Added++
			case drift.PresenceRemoved:
				by[key].Removed++
			case drift.ValueModified:
				by[key].Modified++
			}
		}
		for _, v := range by {
			b.overview.Changes = append(b.overview.Changes, *v)
		}
	}
	if v := b.input.Health; v != nil {
		b.source("health", v.Value.SchemaName, v.Value.SchemaVersion, envelopeID(v.Value), v.ObservedAt)
		for _, r := range v.Value.Records {
			recordSource := SourceReference{Kind: "health_record", Contract: health.SchemaName, Version: health.SchemaVersion, RecordID: r.ID, ObservedAt: v.ObservedAt}
			switch r.Status {
			case health.Critical:
				b.overview.Summary.CriticalHealth++
				b.addAttention(AttentionUrgent, "health_critical", r.Reason, recordSource, importanceDirect)
			case health.Warning:
				b.overview.Summary.WarningHealth++
				b.addAttention(AttentionReview, "health_warning", r.Reason, recordSource, importanceDirect)
			case health.Unknown:
				b.partial = true
			case health.Unsupported:
				b.unsupported = true
			}
			if r.EvidenceState == health.EvidenceInsufficient {
				b.partial = true
			}
			if r.EvidenceState == health.EvidenceUnsupported {
				b.unsupported = true
			}
		}
		if v.Value.OverallStatus == health.Unknown {
			b.partial = true
		}
		if v.Value.OverallStatus == health.Unsupported {
			b.unsupported = true
		}
	}
	if v := b.input.Rule; v != nil {
		b.source("rule", v.Value.SchemaName, v.Value.SchemaVersion, envelopeID(v.Value), v.ObservedAt)
		for _, record := range v.Value.Records {
			if severity := ruleAttention(record.Outcome); severity != AttentionNone {
				if severity == AttentionUnknown {
					b.partial = true
					continue
				}
				b.addAttention(severity, "rule_"+string(record.Outcome), record.Explanation, SourceReference{Kind: "rule_evaluation", Contract: rule.SchemaName, Version: rule.SchemaVersion, RecordID: record.ID, ObservedAt: v.ObservedAt}, importanceRule)
			}
		}
	}
	if v := b.input.Policy; v != nil {
		b.source("policy", v.Value.SchemaName, v.Value.SchemaVersion, envelopeID(v.Value), v.ObservedAt)
		if b.policyRuleIDs == nil {
			b.policyRuleIDs = make(map[string]bool, len(v.Value.Records))
		}
		for _, record := range v.Value.Records {
			b.policyRuleIDs[record.RuleEvaluationID] = true
			if severity := policyAttention(record.Outcome); severity != AttentionNone {
				if record.Outcome == policy.Indeterminate {
					b.partial = true
				}
				b.addAttention(severity, "policy_"+string(record.Outcome), record.Explanation, SourceReference{Kind: "policy_evaluation", Contract: policy.SchemaName, Version: policy.SchemaVersion, RecordID: record.ID, ObservedAt: v.ObservedAt}, importancePolicy)
			}
		}
	}
	if v := b.input.Report; v != nil {
		b.source("report", v.Value.SchemaName, v.Value.SchemaVersion, v.Value.ID, v.ObservedAt)
		if v.Value.Completeness == report.Incomplete {
			b.partial = true
		}
	}
	if v := b.input.PolicyReport; v != nil {
		b.source("policy_report", v.Value.SchemaName, v.Value.SchemaVersion, v.Value.ID, v.ObservedAt)
	}
	if v := b.input.Runtime; v != nil {
		s := b.source("runtime", v.Value.SchemaName, v.Value.SchemaVersion, v.Value.ID, v.ObservedAt)
		if severity := runtimeAttention(v.Value.Outcome); severity != AttentionNone {
			b.partial = true
			failures := runtimeFailureTokens(v.Value)
			if len(failures) == 0 {
				failures = []string{runtimeOutcomeFailureToken(v.Value.Outcome)}
			}
			for _, failure := range failures {
				b.addAttention(severity, failure, string(v.Value.Outcome), s, importanceDirect)
			}
		}
		for _, ar := range v.Value.AlertResults {
			for _, r := range ar.Records {
				b.consumeAlert(r, v.ObservedAt)
			}
		}
	}
	if v := b.input.ServiceState; v != nil {
		s := b.source("runtime_service_state", v.Value.SchemaName, v.Value.SchemaVersion, v.Value.ID, v.ObservedAt)
		b.overview.Guardian = guardianFromLifecycle(v.Value.Lifecycle)
		if v.Value.Lifecycle == runtimeservice.Running && v.Value.LastRuntimeOutcome != "" && v.Value.LastRuntimeOutcome != runtime.Completed {
			b.overview.Guardian = GuardianDegraded
			b.partial = true
		}
		b.attendGuardian(v.Value.Lifecycle, "service_state_observed", s)
	}
	if v := b.input.ServiceResult; v != nil {
		s := b.source("runtime_service_result", v.Value.SchemaName, v.Value.SchemaVersion, v.Value.ID, v.ObservedAt)
		b.overview.Guardian = guardianFromLifecycle(v.Value.FinalState.Lifecycle)
		if v.Value.FinalState.Lifecycle == runtimeservice.Failed {
			b.overview.Guardian = GuardianDegraded
			b.partial = true
		}
		b.attendGuardian(v.Value.FinalState.Lifecycle, v.Value.TerminalReason, s)
	}
}

func runtimeFailureTokens(result runtime.Result) []string {
	tokens := []string{}
	for _, component := range result.Components {
		if component.Status == runtime.ComponentFailed {
			tokens = append(tokens, component.FailureToken)
		}
	}
	sort.Strings(tokens)
	return tokens
}

func runtimeOutcomeFailureToken(outcome runtime.Outcome) string {
	switch outcome {
	case runtime.Cancelled:
		return "runtime_cancelled"
	case runtime.TimedOut:
		return "runtime_timeout"
	default:
		return "runtime_not_completed"
	}
}

func (b *builder) consumeAlert(r alert.Record, observedAt time.Time) {
	source := SourceReference{Kind: "alert_record", Contract: alert.RecordSchema, Version: alert.SchemaVersion, RecordID: r.ID, ObservedAt: observedAt}
	switch r.LifecycleState {
	case alert.StateActive:
		b.overview.Summary.ActiveAlerts++
	case alert.StateAcknowledged:
		b.overview.Summary.AcknowledgedAlerts++
	case alert.StateSuppressed:
		b.overview.Summary.SuppressedAlerts++
	case alert.StateResolved:
		b.overview.Summary.RecoveredAlerts++
	case alert.StateExpired:
		b.overview.Summary.ExpiredAlerts++
	}
	if r.LifecycleState == alert.StateResolved || r.LifecycleState == alert.StateExpired || r.LifecycleState == alert.StateSuppressed {
		return
	}
	b.addAttention(alertAttention(r.Severity), "alert_"+string(r.LifecycleState), r.ReasonToken, source, importanceDirect)
}

func (b *builder) attendGuardian(lifecycle runtimeservice.Lifecycle, reason string, source SourceReference) {
	switch lifecycle {
	case runtimeservice.Failed:
		b.addAttention(AttentionUrgent, "guardian_failed", reason, source, importanceDirect)
	case runtimeservice.Created, runtimeservice.Starting, runtimeservice.Stopping, runtimeservice.Stopped:
		b.addAttention(AttentionReview, "guardian_not_running", reason, source, importanceDirect)
	}
}

func runtimeAttention(outcome runtime.Outcome) Attention {
	switch outcome {
	case runtime.Completed:
		return AttentionNone
	case runtime.Failed, runtime.TimedOut:
		return AttentionUrgent
	case runtime.Partial, runtime.Cancelled:
		return AttentionReview
	default:
		return AttentionUnknown
	}
}

func ruleAttention(outcome rule.Outcome) Attention {
	switch outcome {
	case rule.Matched:
		return AttentionReview
	case rule.NotMatched, rule.DisabledRule:
		return AttentionNone
	case rule.InsufficientEvidence, rule.UnsupportedRule, rule.InvalidRule, rule.EvaluationError:
		return AttentionUnknown
	default:
		return AttentionUnknown
	}
}

func policyAttention(outcome policy.Outcome) Attention {
	switch outcome {
	case policy.Escalated, policy.Conflict:
		return AttentionUrgent
	case policy.Observe, policy.Indeterminate:
		return AttentionReview
	case policy.Accepted, policy.Suppressed, policy.NotApplicable:
		return AttentionNone
	default:
		return AttentionUnknown
	}
}

func (b *builder) addAttention(severity Attention, title, reason string, source SourceReference, importance int) {
	b.attentionCandidates = append(b.attentionCandidates, attentionCandidate{item: AttentionItem{Severity: severity, TitleToken: title, ReasonToken: reason, Source: source}, importance: importance})
}

func (b *builder) finish() {
	b.reduceAttention()
	sort.Slice(b.overview.Sources, func(i, j int) bool {
		a, c := b.overview.Sources[i], b.overview.Sources[j]
		if a.Kind != c.Kind {
			return a.Kind < c.Kind
		}
		if a.RecordID != c.RecordID {
			return a.RecordID < c.RecordID
		}
		return a.ObservedAt.Before(c.ObservedAt)
	})
	sort.Slice(b.overview.Changes, func(i, j int) bool { return b.overview.Changes[i].Category < b.overview.Changes[j].Category })
	if len(b.overview.Sources) == 0 {
		b.overview.Completeness = CompletenessMissing
		b.overview.Freshness = FreshnessNotObserved
		b.overview.Condition = Unavailable
		b.overview.Attention = AttentionUnknown
	} else {
		b.overview.Freshness = FreshnessCurrent
		if b.stale {
			b.overview.Freshness = FreshnessStale
		}
		b.overview.Completeness = CompletenessComplete
		if b.unsupported {
			b.overview.Completeness = CompletenessUnsupported
		} else if b.partial || b.stale {
			b.overview.Completeness = CompletenessPartial
		}
		b.overview.Attention = AttentionNone
		for _, i := range b.overview.AttentionItems {
			if attentionRank(i.Severity) > attentionRank(b.overview.Attention) {
				b.overview.Attention = i.Severity
			}
		}
		switch {
		case b.overview.Attention == AttentionUrgent:
			b.overview.Condition = Critical
		case b.overview.Attention == AttentionReview || b.overview.Completeness != CompletenessComplete:
			b.overview.Condition = Degraded
		case b.input.Health != nil && b.input.Health.Value.OverallStatus == health.Healthy:
			b.overview.Condition = Healthy
		default:
			b.overview.Condition = Unknown
			b.overview.Attention = AttentionUnknown
		}
	}
	b.recommend()
	b.overview.ID = overviewID(b.overview)
}

func (b *builder) reduceAttention() {
	total := len(b.attentionCandidates)
	correlated := 0
	remaining := b.attentionCandidates[:0]
	for _, candidate := range b.attentionCandidates {
		if candidate.item.Source.Kind == "rule_evaluation" && b.policyRuleIDs[candidate.item.Source.RecordID] {
			correlated++
			continue
		}
		remaining = append(remaining, candidate)
	}
	sort.Slice(remaining, func(i, j int) bool { return attentionCandidateLess(remaining[i], remaining[j]) })
	projected := remaining[:0]
	for _, candidate := range remaining {
		if len(projected) > 0 && sameAttentionProjection(projected[len(projected)-1], candidate) {
			correlated++
			continue
		}
		projected = append(projected, candidate)
	}
	remaining = projected
	represented := len(remaining)
	if represented > MaxAttention {
		represented = MaxAttention
	}
	b.overview.AttentionItems = make([]AttentionItem, represented)
	for index := 0; index < represented; index++ {
		b.overview.AttentionItems[index] = remaining[index].item
	}
	omitted := len(remaining) - represented
	if correlated > 0 || omitted > 0 {
		b.overview.AttentionSummary = &AttentionSummary{TotalCandidates: total, Represented: represented, CorrelatedDuplicates: correlated, Omitted: omitted}
	}
}

func sameAttentionProjection(a, b attentionCandidate) bool {
	return a.importance == b.importance &&
		a.item.Severity == b.item.Severity &&
		a.item.TitleToken == b.item.TitleToken &&
		a.item.ReasonToken == b.item.ReasonToken &&
		a.item.Source.Kind == b.item.Source.Kind &&
		a.item.Source.Contract == b.item.Source.Contract &&
		a.item.Source.Version == b.item.Source.Version
}

func attentionCandidateLess(a, b attentionCandidate) bool {
	if attentionRank(a.item.Severity) != attentionRank(b.item.Severity) {
		return attentionRank(a.item.Severity) > attentionRank(b.item.Severity)
	}
	if a.importance != b.importance {
		return a.importance > b.importance
	}
	if a.item.TitleToken != b.item.TitleToken {
		return a.item.TitleToken < b.item.TitleToken
	}
	if a.item.ReasonToken != b.item.ReasonToken {
		return a.item.ReasonToken < b.item.ReasonToken
	}
	if a.item.Source.Kind != b.item.Source.Kind {
		return a.item.Source.Kind < b.item.Source.Kind
	}
	if a.item.Source.Contract != b.item.Source.Contract {
		return a.item.Source.Contract < b.item.Source.Contract
	}
	if a.item.Source.Version != b.item.Source.Version {
		return a.item.Source.Version < b.item.Source.Version
	}
	if a.item.Source.RecordID != b.item.Source.RecordID {
		return a.item.Source.RecordID < b.item.Source.RecordID
	}
	return a.item.Source.ObservedAt.Before(b.item.Source.ObservedAt)
}

func attentionImportance(kind string) int {
	switch kind {
	case "rule_evaluation":
		return importanceRule
	case "policy_evaluation":
		return importancePolicy
	default:
		return importanceDirect
	}
}

func (b *builder) recommend() {
	seen := map[RecommendationToken]bool{}
	add := func(t RecommendationToken) {
		if !seen[t] && len(b.overview.Recommendations) < MaxRecommendations {
			seen[t] = true
			b.overview.Recommendations = append(b.overview.Recommendations, Recommendation{Token: t})
		}
	}
	if b.overview.Attention == AttentionUrgent || b.overview.Attention == AttentionReview {
		add(RecommendInspectAttention)
	}
	if b.overview.Summary.Changes > 0 {
		add(RecommendReviewChanges)
	}
	if b.overview.Freshness == FreshnessStale || b.overview.Completeness == CompletenessMissing {
		add(RecommendRunFreshCheck)
	}
	if b.overview.Freshness != FreshnessStale && (b.overview.Completeness == CompletenessPartial || b.overview.Completeness == CompletenessUnsupported) {
		add(RecommendInspectEvidence)
	}
	if b.input.Runtime != nil && b.input.Runtime.Value.Outcome != runtime.Completed {
		add(RecommendInspectFailedOperation)
	}
	if b.overview.Guardian == GuardianFailed || b.overview.Guardian == GuardianDegraded || b.overview.Guardian == GuardianUnavailable || b.overview.Guardian == GuardianStopped || b.overview.Guardian == GuardianNotObserved {
		add(RecommendVerifyGuardian)
	}
	if len(b.overview.Recommendations) == 0 {
		add(RecommendNoAction)
	}
}

func ValidateInput(v Input) error {
	if v.SchemaName != InputSchema || v.SchemaVersion != SchemaVersion || v.ObservedAt.IsZero() || v.FreshForNS <= 0 || time.Duration(v.FreshForNS) > MaxFreshFor {
		return fmt.Errorf("invalid operator projection input")
	}
	if v.ServiceState != nil && v.ServiceResult != nil {
		return fmt.Errorf("service state and terminal result are mutually exclusive")
	}
	observed := []time.Time{}
	add := func(at time.Time) error {
		if at.IsZero() || at.After(v.ObservedAt) {
			return fmt.Errorf("invalid source observation time")
		}
		observed = append(observed, at)
		return nil
	}
	if x := v.Command; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := validateCommand(x.Value); err != nil {
			return err
		}
	}
	if x := v.Inventory; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := inventory.Validate(x.Value); err != nil {
			return fmt.Errorf("invalid inventory: %w", err)
		}
	}
	if x := v.Comparison; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := comparison.Validate(x.Value); err != nil {
			return fmt.Errorf("invalid comparison: %w", err)
		}
	}
	if x := v.Drift; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := drift.Validate(x.Value); err != nil {
			return fmt.Errorf("invalid drift: %w", err)
		}
	}
	if x := v.Health; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := health.Validate(x.Value); err != nil {
			return fmt.Errorf("invalid health: %w", err)
		}
	}
	if x := v.Rule; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := rule.Validate(x.Value); err != nil {
			return fmt.Errorf("invalid rule: %w", err)
		}
	}
	if x := v.Policy; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := policy.Validate(x.Value); err != nil {
			return fmt.Errorf("invalid policy: %w", err)
		}
	}
	if x := v.Report; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := report.Validate(x.Value); err != nil {
			return fmt.Errorf("invalid report: %w", err)
		}
	}
	if x := v.PolicyReport; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := report.ValidatePolicyReport(x.Value); err != nil {
			return fmt.Errorf("invalid policy report: %w", err)
		}
	}
	if x := v.Runtime; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := runtime.ValidateResult(x.Value); err != nil {
			return fmt.Errorf("invalid runtime: %w", err)
		}
	}
	if x := v.ServiceState; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := runtimeservice.ValidateState(x.Value); err != nil {
			return fmt.Errorf("invalid service state: %w", err)
		}
	}
	if x := v.ServiceResult; x != nil {
		if err := add(x.ObservedAt); err != nil {
			return err
		}
		if err := runtimeservice.ValidateResult(x.Value); err != nil {
			return fmt.Errorf("invalid service result: %w", err)
		}
	}
	if len(observed) > MaxSources {
		return fmt.Errorf("operator source limit exceeded")
	}
	return validateCorrelations(v)
}

func validateCommand(v command.Execution) error {
	if v.SchemaName != command.ExecutionSchema || v.SchemaVersion != command.SchemaVersion || v.ID == "" || v.CommandID == "" || v.PlanID == "" || v.Stages == nil || v.View.Rows == nil || v.View.Groups == nil || v.Diagnostics == nil {
		return fmt.Errorf("invalid command execution")
	}
	seen := map[command.Stage]bool{}
	for _, s := range v.Stages {
		if seen[s.Stage] || s.ContractName == "" || s.Version == "" || s.RecordCount < 0 {
			return fmt.Errorf("invalid command stage")
		}
		seen[s.Stage] = true
	}
	return nil
}

func validateCorrelations(v Input) error {
	if v.Comparison != nil && v.Drift != nil {
		changes := map[string]bool{}
		for _, r := range v.Comparison.Value.Changes {
			changes[r.ID] = true
		}
		for _, r := range v.Drift.Value.Records {
			if !changes[r.ChangeID] {
				return fmt.Errorf("drift source is not correlated to comparison")
			}
		}
	}
	if v.Drift != nil && v.Health != nil {
		drifts := map[string]bool{}
		for _, r := range v.Drift.Value.Records {
			drifts[r.ID] = true
		}
		for _, r := range v.Health.Value.Records {
			if !drifts[r.DriftID] {
				return fmt.Errorf("health source is not correlated to drift")
			}
		}
	}
	if v.Rule != nil && v.Policy != nil {
		rules := map[string]bool{}
		for _, record := range v.Rule.Value.Records {
			rules[record.ID] = true
		}
		for _, record := range v.Policy.Value.Records {
			if !rules[record.RuleEvaluationID] {
				return fmt.Errorf("policy source is not correlated to rule")
			}
		}
	}
	return nil
}

func Validate(v Overview) error {
	currentVersion := v.SchemaVersion == SchemaVersion && v.ModelVersion == ModelVersion
	legacyVersion := (v.SchemaVersion == LegacySchemaVersion && v.ModelVersion == LegacyModelVersion && v.AttentionSummary == nil) || (v.SchemaVersion == LegacySchemaV11 && v.ModelVersion == LegacyModelV11)
	if v.SchemaName != SchemaName || (!currentVersion && !legacyVersion) || !validTime(v.ObservedAt) || !validCondition(v.Condition) || !validAttention(v.Attention) || !validGuardian(v.Guardian) || (legacyVersion && (v.Guardian == GuardianDegraded || v.Guardian == GuardianUnavailable)) || !validFreshness(v.Freshness) || !validCompleteness(v.Completeness) || v.Changes == nil || v.AttentionItems == nil || v.Recommendations == nil || v.Sources == nil || len(v.Sources) > MaxSources || len(v.AttentionItems) > MaxAttention || len(v.Recommendations) == 0 || len(v.Recommendations) > MaxRecommendations {
		return fmt.Errorf("invalid operator overview")
	}
	if v.AttentionSummary != nil {
		s := v.AttentionSummary
		if s.TotalCandidates < 0 || s.TotalCandidates > 1<<20 || s.Represented != len(v.AttentionItems) || s.CorrelatedDuplicates < 0 || s.Omitted < 0 || s.CorrelatedDuplicates+s.Omitted == 0 || s.TotalCandidates != s.Represented+s.CorrelatedDuplicates+s.Omitted {
			return fmt.Errorf("invalid attention summary")
		}
	}
	counts := []int{v.Summary.Changes, v.Summary.CriticalHealth, v.Summary.WarningHealth, v.Summary.ActiveAlerts, v.Summary.AcknowledgedAlerts, v.Summary.SuppressedAlerts, v.Summary.RecoveredAlerts, v.Summary.ExpiredAlerts}
	for _, count := range counts {
		if count < 0 || count > 1<<20 {
			return fmt.Errorf("invalid operator summary count")
		}
	}
	last := ""
	for _, s := range v.Sources {
		key := s.Kind + "\x00" + s.RecordID + "\x00" + s.ObservedAt.Format(time.RFC3339Nano)
		if key <= last || !validReference(s) || s.ObservedAt.After(v.ObservedAt) {
			return fmt.Errorf("invalid operator source")
		}
		last = key
	}
	last = ""
	for _, c := range v.Changes {
		if !token(c.Category) || c.Added < 0 || c.Removed < 0 || c.Modified < 0 || c.Category <= last {
			return fmt.Errorf("invalid change summary")
		}
		last = c.Category
	}
	for index, i := range v.AttentionItems {
		if i.Severity != AttentionReview && i.Severity != AttentionUrgent || !token(i.TitleToken) || !token(i.ReasonToken) || !validReference(i.Source) {
			return fmt.Errorf("invalid attention item")
		}
		if index > 0 {
			previous := v.AttentionItems[index-1]
			if !attentionCandidateLess(attentionCandidate{item: previous, importance: attentionImportance(previous.Source.Kind)}, attentionCandidate{item: i, importance: attentionImportance(i.Source.Kind)}) {
				return fmt.Errorf("attention items are not canonically ordered")
			}
		}
	}
	seen := map[RecommendationToken]bool{}
	for _, r := range v.Recommendations {
		if seen[r.Token] || !validRecommendation(r.Token) {
			return fmt.Errorf("invalid recommendation")
		}
		seen[r.Token] = true
	}
	if seen[RecommendNoAction] && len(v.Recommendations) != 1 {
		return fmt.Errorf("no-action recommendation must stand alone")
	}
	if len(v.Sources) == 0 && (v.Condition != Unavailable || v.Attention != AttentionUnknown || v.Freshness != FreshnessNotObserved || v.Completeness != CompletenessMissing) {
		return fmt.Errorf("missing evidence is not explicit")
	}
	if v.Condition == Healthy && (v.Attention != AttentionNone || v.Freshness != FreshnessCurrent || v.Completeness != CompletenessComplete) {
		return fmt.Errorf("invalid healthy overview")
	}
	if v.Condition == Critical && v.Attention != AttentionUrgent {
		return fmt.Errorf("critical overview is not urgent")
	}
	if v.Guardian == GuardianFailed && v.Attention != AttentionUrgent {
		return fmt.Errorf("failed Guardian is not urgent")
	}
	maxAttention := AttentionNone
	for _, item := range v.AttentionItems {
		if attentionRank(item.Severity) > attentionRank(maxAttention) {
			maxAttention = item.Severity
		}
	}
	if len(v.AttentionItems) > 0 && attentionRank(v.Attention) < attentionRank(maxAttention) {
		return fmt.Errorf("attention summary masks an item")
	}
	if v.ID != overviewID(v) {
		return fmt.Errorf("invalid operator overview identity")
	}
	return nil
}

func MarshalCanonical(v Overview) ([]byte, error) {
	if err := Validate(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}

func DecodeInput(data []byte) (Input, error) {
	var v Input
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&v); err != nil {
		return Input{}, err
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return Input{}, fmt.Errorf("trailing operator input data")
	}
	return v, ValidateInput(v)
}
func Decode(data []byte) (Overview, error) {
	var v Overview
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&v); err != nil {
		return Overview{}, err
	}
	if err := d.Decode(&struct{}{}); err != io.EOF {
		return Overview{}, fmt.Errorf("trailing operator overview data")
	}
	return v, Validate(v)
}

func guardianFromLifecycle(v runtimeservice.Lifecycle) Guardian {
	switch v {
	case runtimeservice.Running:
		return GuardianRunning
	case runtimeservice.Created, runtimeservice.Starting:
		return GuardianStarting
	case runtimeservice.Stopping:
		return GuardianStopping
	case runtimeservice.Stopped:
		return GuardianStopped
	case runtimeservice.Failed:
		return GuardianDegraded
	default:
		return GuardianNotObserved
	}
}
func attentionRank(v Attention) int {
	switch v {
	case AttentionUrgent:
		return 3
	case AttentionReview:
		return 2
	case AttentionNone:
		return 1
	default:
		return 0
	}
}
func envelopeID(v any) string {
	d, _ := json.Marshal(v)
	s := sha256.Sum256(d)
	return hex.EncodeToString(s[:])
}
func overviewID(v Overview) string {
	v.ID = ""
	d, _ := json.Marshal(v)
	s := sha256.Sum256(append([]byte(SchemaName+"/"+v.SchemaVersion+"\x00"), d...))
	return hex.EncodeToString(s[:])
}
func validTime(v time.Time) bool { return !v.IsZero() && v.Location() == time.UTC }
func validReference(v SourceReference) bool {
	return token(v.Kind) && len(v.Contract) > 0 && len(v.Contract) <= MaxTokenLength && len(v.Version) > 0 && len(v.Version) <= MaxTokenLength && len(v.RecordID) > 0 && len(v.RecordID) <= 512 && validTime(v.ObservedAt)
}
func token(v string) bool {
	return len(v) > 0 && len(v) <= MaxTokenLength && tokenPattern.MatchString(v)
}
func validCondition(v Condition) bool {
	return v == Healthy || v == Degraded || v == Critical || v == Unknown || v == Unavailable
}
func validAttention(v Attention) bool {
	return v == AttentionNone || v == AttentionReview || v == AttentionUrgent || v == AttentionUnknown
}
func validGuardian(v Guardian) bool {
	return v == GuardianRunning || v == GuardianStarting || v == GuardianStopping || v == GuardianStopped || v == GuardianFailed || v == GuardianDegraded || v == GuardianUnavailable || v == GuardianNotObserved
}
func validFreshness(v Freshness) bool {
	return v == FreshnessCurrent || v == FreshnessStale || v == FreshnessInvalid || v == FreshnessNotObserved
}
func validCompleteness(v Completeness) bool {
	return v == CompletenessComplete || v == CompletenessPartial || v == CompletenessUnsupported || v == CompletenessMissing || v == CompletenessInvalid || v == CompletenessNotObserved
}
func validRecommendation(v RecommendationToken) bool {
	switch v {
	case RecommendInspectAttention, RecommendReviewChanges, RecommendRunFreshCheck, RecommendInspectEvidence, RecommendInspectFailedOperation, RecommendVerifyGuardian, RecommendNoAction:
		return true
	}
	return false
}
