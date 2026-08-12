// Package notification provides deterministic, provider-neutral delivery
// planning for immutable canonical Alert Records and an explicitly invoked
// one-cycle provider adapter. It owns no alert, persistence, daemon, transport,
// configuration, monitoring, remediation, or presentation behavior.
package notification

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"sort"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
)

const (
	SchemaVersion         = "1.0"
	ModelVersion          = "1.0"
	TaxonomyVersion       = "1.0"
	InputSchema           = "qwsg.notification-planning-input"
	PolicySchema          = "qwsg.notification-delivery-policy"
	QueueSchema           = "qwsg.notification-queue-state"
	PlanSchema            = "qwsg.notification-delivery-plan"
	RequestSchema         = "qwsg.notification-delivery-request"
	AttemptSchema         = "qwsg.notification-delivery-attempt"
	StatusSchema          = "qwsg.notification-delivery-status"
	AcknowledgementSchema = "qwsg.notification-delivery-acknowledgement"
	EvidenceSchema        = "qwsg.notification-delivery-evidence"
	ProviderSchema        = "qwsg.notification-provider-descriptor"
	ProviderResultSchema  = "qwsg.notification-provider-result"
	CycleResultSchema     = "qwsg.notification-cycle-result"
	MaxAlerts             = 4096
	MaxRoutes             = 256
	MaxEndpoints          = 1024
	MaxBindings           = 256
	MaxQueueEntries       = 8192
	MaxAttempts           = 10
	MaxBackoffSteps       = 10
	MaxStringLength       = 512
	MaxProviderEvidence   = 16
	MaxDeliveryWindow     = 24 * time.Hour
)

type Channel string

const (
	Email    Channel = "email"
	Webhook  Channel = "webhook"
	Slack    Channel = "slack"
	Discord  Channel = "discord"
	Telegram Channel = "telegram"
	SMS      Channel = "sms"
)

type DeliveryStatus string

const (
	StatusQueued           DeliveryStatus = "queued"
	StatusAccepted         DeliveryStatus = "accepted"
	StatusDelivered        DeliveryStatus = "delivered"
	StatusRetryableFailure DeliveryStatus = "retryable_failure"
	StatusRetryScheduled   DeliveryStatus = "retry_scheduled"
	StatusTerminalFailure  DeliveryStatus = "terminal_failure"
	StatusExhausted        DeliveryStatus = "exhausted"
	StatusIndeterminate    DeliveryStatus = "indeterminate"
)

type AcknowledgementKind string

const (
	AcknowledgedAccepted  AcknowledgementKind = "provider_accepted"
	AcknowledgedDelivered AcknowledgementKind = "provider_reported_delivered"
	AcknowledgedUnknown   AcknowledgementKind = "provider_outcome_unknown"
)

type FailureClass string

const (
	FailureNone                FailureClass = "none"
	FailureRetryable           FailureClass = "retryable"
	FailureRateLimited         FailureClass = "rate_limited"
	FailureIndeterminate       FailureClass = "indeterminate"
	FailureAuthentication      FailureClass = "authentication"
	FailureAuthorization       FailureClass = "authorization"
	FailureInvalidDestination  FailureClass = "invalid_destination"
	FailureRejectedPayload     FailureClass = "rejected_payload"
	FailureUnsupportedProvider FailureClass = "unsupported_provider"
	FailureTerminal            FailureClass = "terminal"
)

type RetryPolicy struct {
	MaxAttempts      int     `json:"max_attempts"`
	DeliveryWindowNS int64   `json:"delivery_window_ns"`
	BackoffNS        []int64 `json:"backoff_ns"`
}

type Route struct {
	ID          string            `json:"id"`
	Enabled     bool              `json:"enabled"`
	Severities  []alert.Severity  `json:"severities"`
	Categories  []alert.Category  `json:"categories"`
	Events      []alert.EventKind `json:"events"`
	EndpointIDs []string          `json:"endpoint_ids"`
}

type EndpointReference struct {
	ID             string  `json:"id"`
	Channel        Channel `json:"channel"`
	DestinationRef string  `json:"destination_ref"`
	SecretRef      string  `json:"secret_ref,omitempty"`
}

type ProviderBinding struct {
	ID         string  `json:"id"`
	Channel    Channel `json:"channel"`
	ProviderID string  `json:"provider_id"`
}

type Policy struct {
	SchemaName    string              `json:"schema_name"`
	SchemaVersion string              `json:"schema_version"`
	ID            string              `json:"id"`
	Retry         RetryPolicy         `json:"retry"`
	Routes        []Route             `json:"routes"`
	Endpoints     []EndpointReference `json:"endpoints"`
	Bindings      []ProviderBinding   `json:"bindings"`
}

type DeliveryEnvelope struct {
	AlertRecordID      string          `json:"alert_record_id"`
	LifecycleID        string          `json:"lifecycle_id"`
	ConditionKey       string          `json:"condition_key"`
	Event              alert.EventKind `json:"event"`
	Severity           alert.Severity  `json:"severity"`
	Category           alert.Category  `json:"category"`
	ReasonToken        string          `json:"reason_token"`
	EventTime          time.Time       `json:"event_time"`
	EvidenceReferences []string        `json:"evidence_references"`
}

type Request struct {
	SchemaName     string           `json:"schema_name"`
	SchemaVersion  string           `json:"schema_version"`
	ID             string           `json:"id"`
	DeliveryID     string           `json:"delivery_id"`
	IdempotencyKey string           `json:"idempotency_key"`
	AttemptNumber  int              `json:"attempt_number"`
	RouteID        string           `json:"route_id"`
	EndpointID     string           `json:"endpoint_id"`
	Channel        Channel          `json:"channel"`
	ProviderID     string           `json:"provider_id"`
	DestinationRef string           `json:"destination_ref"`
	SecretRef      string           `json:"secret_ref,omitempty"`
	CreatedAt      time.Time        `json:"created_at"`
	Deadline       time.Time        `json:"deadline"`
	Envelope       DeliveryEnvelope `json:"envelope"`
}

type EvidenceReference struct {
	SchemaName        string   `json:"schema_name"`
	SchemaVersion     string   `json:"schema_version"`
	AlertRecordID     string   `json:"alert_record_id"`
	RequestID         string   `json:"request_id"`
	ProviderID        string   `json:"provider_id"`
	ProviderReference string   `json:"provider_reference,omitempty"`
	Tokens            []string `json:"tokens"`
}

type Acknowledgement struct {
	SchemaName        string              `json:"schema_name"`
	SchemaVersion     string              `json:"schema_version"`
	ID                string              `json:"id"`
	RequestID         string              `json:"request_id"`
	Kind              AcknowledgementKind `json:"kind"`
	At                time.Time           `json:"at"`
	ProviderReference string              `json:"provider_reference,omitempty"`
}

type StatusRecord struct {
	SchemaName    string         `json:"schema_name"`
	SchemaVersion string         `json:"schema_version"`
	ID            string         `json:"id"`
	RequestID     string         `json:"request_id"`
	Status        DeliveryStatus `json:"status"`
	Failure       FailureClass   `json:"failure"`
	At            time.Time      `json:"at"`
}

type Attempt struct {
	SchemaName      string            `json:"schema_name"`
	SchemaVersion   string            `json:"schema_version"`
	ID              string            `json:"id"`
	Request         Request           `json:"request"`
	StartedAt       time.Time         `json:"started_at"`
	CompletedAt     time.Time         `json:"completed_at"`
	Status          StatusRecord      `json:"status"`
	Acknowledgement *Acknowledgement  `json:"acknowledgement,omitempty"`
	Evidence        EvidenceReference `json:"evidence"`
}

type QueueEntry struct {
	DeliveryID    string         `json:"delivery_id"`
	Request       Request        `json:"request"`
	Status        DeliveryStatus `json:"status"`
	Failure       FailureClass   `json:"failure"`
	NextAttemptAt time.Time      `json:"next_attempt_at,omitempty"`
	Attempts      []Attempt      `json:"attempts"`
}

type QueueState struct {
	SchemaName    string       `json:"schema_name"`
	SchemaVersion string       `json:"schema_version"`
	ID            string       `json:"id"`
	Entries       []QueueEntry `json:"entries"`
}

type PlanningInput struct {
	SchemaName    string         `json:"schema_name"`
	SchemaVersion string         `json:"schema_version"`
	EvaluatedAt   time.Time      `json:"evaluated_at"`
	Alerts        []alert.Record `json:"alerts"`
	Policy        Policy         `json:"policy"`
	PreviousQueue QueueState     `json:"previous_queue"`
}

type Eligibility struct {
	AlertRecordID string   `json:"alert_record_id"`
	Eligible      bool     `json:"eligible"`
	ReasonToken   string   `json:"reason_token"`
	RequestIDs    []string `json:"request_ids"`
}

type Plan struct {
	SchemaName      string        `json:"schema_name"`
	SchemaVersion   string        `json:"schema_version"`
	ModelVersion    string        `json:"model_version"`
	TaxonomyVersion string        `json:"taxonomy_version"`
	ID              string        `json:"id"`
	EvaluatedAt     time.Time     `json:"evaluated_at"`
	Eligibility     []Eligibility `json:"eligibility"`
	Requests        []Request     `json:"requests"`
	NextQueue       QueueState    `json:"next_queue"`
}

type ProviderDescriptor struct {
	SchemaName    string    `json:"schema_name"`
	SchemaVersion string    `json:"schema_version"`
	ID            string    `json:"id"`
	Channels      []Channel `json:"channels"`
}

type ProviderResult struct {
	SchemaName        string         `json:"schema_name"`
	SchemaVersion     string         `json:"schema_version"`
	Status            DeliveryStatus `json:"status"`
	Failure           FailureClass   `json:"failure"`
	CompletedAt       time.Time      `json:"completed_at"`
	ProviderReference string         `json:"provider_reference,omitempty"`
	EvidenceTokens    []string       `json:"evidence_tokens"`
}

type Provider interface {
	Descriptor() ProviderDescriptor
	Deliver(context.Context, Request) ProviderResult
}

type Registry struct{ providers map[string]Provider }

type CycleResult struct {
	SchemaName    string     `json:"schema_name"`
	SchemaVersion string     `json:"schema_version"`
	ID            string     `json:"id"`
	PlanID        string     `json:"plan_id"`
	ExecutedAt    time.Time  `json:"executed_at"`
	Attempts      []Attempt  `json:"attempts"`
	NextQueue     QueueState `json:"next_queue"`
}

var tokenPattern = regexp.MustCompile(`^[a-z0-9]+([._:-][a-z0-9]+)*$`)

func NewQueueState() QueueState {
	v := QueueState{SchemaName: QueueSchema, SchemaVersion: SchemaVersion, Entries: []QueueEntry{}}
	v.ID = queueIdentity(v)
	return v
}

func NewPolicy(retry RetryPolicy, routes []Route, endpoints []EndpointReference, bindings []ProviderBinding) (Policy, error) {
	v := Policy{SchemaName: PolicySchema, SchemaVersion: SchemaVersion, Retry: retry, Routes: cloneRoutes(routes), Endpoints: append([]EndpointReference{}, endpoints...), Bindings: append([]ProviderBinding{}, bindings...)}
	canonicalizePolicy(&v)
	v.ID = contentIdentity("notification-policy", v)
	return v, ValidatePolicy(v)
}

func NewRegistry(providers ...Provider) (Registry, error) {
	r := Registry{providers: map[string]Provider{}}
	if len(providers) > MaxBindings {
		return Registry{}, fmt.Errorf("provider limit exceeded")
	}
	for _, p := range providers {
		if p == nil {
			return Registry{}, fmt.Errorf("nil provider")
		}
		d := p.Descriptor()
		if err := ValidateProviderDescriptor(d); err != nil {
			return Registry{}, err
		}
		if _, exists := r.providers[d.ID]; exists {
			return Registry{}, fmt.Errorf("duplicate provider")
		}
		r.providers[d.ID] = p
	}
	return r, nil
}

func PlanDeliveries(input PlanningInput) (Plan, error) {
	if err := ValidatePlanningInput(input); err != nil {
		return Plan{}, err
	}
	endpointByID := map[string]EndpointReference{}
	providerByChannel := map[Channel]string{}
	for _, v := range input.Policy.Endpoints {
		endpointByID[v.ID] = v
	}
	for _, v := range input.Policy.Bindings {
		providerByChannel[v.Channel] = v.ProviderID
	}
	queue := cloneQueue(input.PreviousQueue)
	byDelivery := map[string]int{}
	for i := range queue.Entries {
		byDelivery[queue.Entries[i].DeliveryID] = i
	}
	eligibility := make([]Eligibility, 0, len(input.Alerts))
	requests := []Request{}
	for _, record := range input.Alerts {
		item := Eligibility{AlertRecordID: record.ID, RequestIDs: []string{}}
		if record.Decision == alert.DecisionSuppressed {
			item.ReasonToken = "alert_record_suppressed"
			eligibility = append(eligibility, item)
			continue
		}
		for _, route := range input.Policy.Routes {
			if !route.Enabled || !routeMatches(route, record) {
				continue
			}
			for _, endpointID := range route.EndpointIDs {
				endpoint := endpointByID[endpointID]
				deliveryID := stableID("delivery", record.ID, route.ID, endpoint.ID)
				if index, exists := byDelivery[deliveryID]; exists {
					entry := &queue.Entries[index]
					retryRequest, due, disposition := planRetry(*entry, input.Policy.Retry, input.EvaluatedAt)
					switch disposition {
					case StatusQueued:
						entry.Request = retryRequest
						entry.Status, entry.Failure, entry.NextAttemptAt = StatusQueued, FailureNone, time.Time{}
						requests = append(requests, retryRequest)
						item.RequestIDs = append(item.RequestIDs, retryRequest.ID)
					case StatusRetryScheduled:
						entry.Status, entry.NextAttemptAt = StatusRetryScheduled, due
					case StatusExhausted:
						entry.Status, entry.NextAttemptAt = StatusExhausted, time.Time{}
					}
					continue
				}
				req := newRequest(record, route.ID, endpoint, providerByChannel[endpoint.Channel], input.Policy.Retry, input.EvaluatedAt, 1)
				entry := QueueEntry{DeliveryID: deliveryID, Request: req, Status: StatusQueued, Failure: FailureNone, Attempts: []Attempt{}}
				queue.Entries = append(queue.Entries, entry)
				if len(queue.Entries) > MaxQueueEntries {
					return Plan{}, fmt.Errorf("notification queue entry limit exceeded")
				}
				byDelivery[deliveryID] = len(queue.Entries) - 1
				requests = append(requests, req)
				item.RequestIDs = append(item.RequestIDs, req.ID)
			}
		}
		item.Eligible = len(item.RequestIDs) > 0
		if item.Eligible {
			item.ReasonToken = "delivery_planned"
		} else {
			item.ReasonToken = "no_matching_route"
		}
		sort.Strings(item.RequestIDs)
		eligibility = append(eligibility, item)
	}
	sort.Slice(eligibility, func(i, j int) bool { return eligibility[i].AlertRecordID < eligibility[j].AlertRecordID })
	sortRequests(requests)
	sortQueue(&queue)
	queue.ID = queueIdentity(queue)
	plan := Plan{SchemaName: PlanSchema, SchemaVersion: SchemaVersion, ModelVersion: ModelVersion, TaxonomyVersion: TaxonomyVersion, EvaluatedAt: input.EvaluatedAt, Eligibility: eligibility, Requests: requests, NextQueue: queue}
	plan.ID = contentIdentity("notification-plan", plan)
	if err := ValidatePlan(plan); err != nil {
		return Plan{}, err
	}
	return plan, nil
}

func ExecuteCycle(ctx context.Context, plan Plan, registry Registry, executedAt time.Time) (CycleResult, error) {
	if err := ValidatePlan(plan); err != nil {
		return CycleResult{}, err
	}
	if !validTime(executedAt) || executedAt.Before(plan.EvaluatedAt) {
		return CycleResult{}, fmt.Errorf("invalid execution time")
	}
	queue := cloneQueue(plan.NextQueue)
	byDelivery := map[string]int{}
	for i := range queue.Entries {
		byDelivery[queue.Entries[i].DeliveryID] = i
	}
	attempts := make([]Attempt, 0, len(plan.Requests))
	for _, req := range plan.Requests {
		provider, ok := registry.providers[req.ProviderID]
		var outcome ProviderResult
		if !ok || !providerSupports(provider.Descriptor(), req.Channel) {
			outcome = ProviderResult{SchemaName: ProviderResultSchema, SchemaVersion: SchemaVersion, Status: StatusTerminalFailure, Failure: FailureUnsupportedProvider, CompletedAt: executedAt, EvidenceTokens: []string{"provider_unavailable"}}
		} else {
			bounded, cancel := context.WithDeadline(ctx, req.Deadline)
			outcome = provider.Deliver(bounded, req)
			cancel()
			if err := ValidateProviderResult(outcome, executedAt, req.Deadline); err != nil {
				return CycleResult{}, fmt.Errorf("invalid provider result for %s: %w", req.ID, err)
			}
		}
		attempt := makeAttempt(req, executedAt, outcome)
		attempts = append(attempts, attempt)
		index, exists := byDelivery[req.DeliveryID]
		if !exists {
			return CycleResult{}, fmt.Errorf("request missing from queue")
		}
		entry := &queue.Entries[index]
		entry.Attempts = append(entry.Attempts, attempt)
		entry.Status, entry.Failure = attempt.Status.Status, attempt.Status.Failure
	}
	sort.Slice(attempts, func(i, j int) bool { return attempts[i].ID < attempts[j].ID })
	sortQueue(&queue)
	queue.ID = queueIdentity(queue)
	result := CycleResult{SchemaName: CycleResultSchema, SchemaVersion: SchemaVersion, PlanID: plan.ID, ExecutedAt: executedAt, Attempts: attempts, NextQueue: queue}
	result.ID = contentIdentity("notification-cycle", result)
	if err := ValidateCycleResult(result); err != nil {
		return CycleResult{}, err
	}
	return result, nil
}

func ValidatePlanningInput(v PlanningInput) error {
	if v.SchemaName != InputSchema || v.SchemaVersion != SchemaVersion || !validTime(v.EvaluatedAt) || v.Alerts == nil || len(v.Alerts) > MaxAlerts {
		return fmt.Errorf("invalid notification planning input")
	}
	if err := ValidatePolicy(v.Policy); err != nil {
		return err
	}
	if err := ValidateQueue(v.PreviousQueue); err != nil {
		return err
	}
	last := ""
	for _, record := range v.Alerts {
		if err := alert.ValidateRecord(record); err != nil {
			return fmt.Errorf("invalid alert record: %w", err)
		}
		if record.ID <= last || record.EvaluationTime.After(v.EvaluatedAt) {
			return fmt.Errorf("duplicate, unordered, or future alert record")
		}
		last = record.ID
	}
	return nil
}

func ValidatePolicy(v Policy) error {
	if v.SchemaName != PolicySchema || v.SchemaVersion != SchemaVersion || v.Routes == nil || v.Endpoints == nil || v.Bindings == nil || len(v.Routes) > MaxRoutes || len(v.Endpoints) > MaxEndpoints || len(v.Bindings) > MaxBindings {
		return fmt.Errorf("invalid notification policy envelope")
	}
	if v.Retry.MaxAttempts < 1 || v.Retry.MaxAttempts > MaxAttempts || v.Retry.DeliveryWindowNS <= 0 || time.Duration(v.Retry.DeliveryWindowNS) > MaxDeliveryWindow || len(v.Retry.BackoffNS) != v.Retry.MaxAttempts-1 || len(v.Retry.BackoffNS) > MaxBackoffSteps {
		return fmt.Errorf("invalid retry policy")
	}
	for i, ns := range v.Retry.BackoffNS {
		if ns <= 0 || time.Duration(ns) >= time.Duration(v.Retry.DeliveryWindowNS) || (i > 0 && ns <= v.Retry.BackoffNS[i-1]) {
			return fmt.Errorf("invalid retry backoff")
		}
	}
	endpoints := map[string]EndpointReference{}
	last := ""
	for _, e := range v.Endpoints {
		if e.ID <= last || !token(e.ID) || !validChannel(e.Channel) || !reference(e.DestinationRef) || (e.SecretRef != "" && !reference(e.SecretRef)) {
			return fmt.Errorf("invalid endpoint")
		}
		last = e.ID
		endpoints[e.ID] = e
	}
	bindings := map[Channel]ProviderBinding{}
	last = ""
	for _, b := range v.Bindings {
		if b.ID <= last || !token(b.ID) || !validChannel(b.Channel) || !token(b.ProviderID) {
			return fmt.Errorf("invalid binding")
		}
		if _, ok := bindings[b.Channel]; ok {
			return fmt.Errorf("duplicate channel binding")
		}
		last = b.ID
		bindings[b.Channel] = b
	}
	last = ""
	totalFanout := 0
	for _, route := range v.Routes {
		if route.ID <= last || !token(route.ID) || route.EndpointIDs == nil || !sortedUnique(route.EndpointIDs) || !sortedAlertSeverities(route.Severities) || !sortedAlertCategories(route.Categories) || !sortedAlertEvents(route.Events) {
			return fmt.Errorf("invalid route")
		}
		for _, id := range route.EndpointIDs {
			e, ok := endpoints[id]
			if !ok {
				return fmt.Errorf("unknown route endpoint")
			}
			b, ok := bindings[e.Channel]
			if !ok || b.Channel != e.Channel {
				return fmt.Errorf("endpoint has no provider binding")
			}
		}
		totalFanout += len(route.EndpointIDs)
		if totalFanout > MaxQueueEntries {
			return fmt.Errorf("notification policy fanout limit exceeded")
		}
		last = route.ID
	}
	if v.ID != contentIdentity("notification-policy", v) {
		return fmt.Errorf("invalid notification policy identity")
	}
	return nil
}

func ValidateQueue(v QueueState) error {
	if v.SchemaName != QueueSchema || v.SchemaVersion != SchemaVersion || v.Entries == nil || len(v.Entries) > MaxQueueEntries {
		return fmt.Errorf("invalid notification queue")
	}
	var previous *QueueEntry
	for _, e := range v.Entries {
		if (previous != nil && !lessQueueEntry(*previous, e)) || e.DeliveryID != e.Request.DeliveryID || !validQueueOutcome(e.Status, e.Failure) || e.Attempts == nil || len(e.Attempts) > MaxAttempts || !validRequest(e.Request) {
			return fmt.Errorf("invalid notification queue entry")
		}
		if (e.Status == StatusRetryScheduled) != !e.NextAttemptAt.IsZero() || (!e.NextAttemptAt.IsZero() && !validTime(e.NextAttemptAt)) {
			return fmt.Errorf("invalid retry schedule")
		}
		for i, a := range e.Attempts {
			if !validAttempt(a) || a.Request.AttemptNumber != i+1 || a.Request.DeliveryID != e.DeliveryID {
				return fmt.Errorf("invalid queue attempt")
			}
		}
		copy := e
		previous = &copy
	}
	if v.ID != queueIdentity(v) {
		return fmt.Errorf("invalid notification queue identity")
	}
	return nil
}

func ValidatePlan(v Plan) error {
	if v.SchemaName != PlanSchema || v.SchemaVersion != SchemaVersion || v.ModelVersion != ModelVersion || v.TaxonomyVersion != TaxonomyVersion || !validTime(v.EvaluatedAt) || v.Eligibility == nil || v.Requests == nil {
		return fmt.Errorf("invalid notification plan")
	}
	if err := ValidateQueue(v.NextQueue); err != nil {
		return err
	}
	last := ""
	for _, e := range v.Eligibility {
		if e.AlertRecordID <= last || !token(e.AlertRecordID) || !token(e.ReasonToken) || e.RequestIDs == nil || !sortedUnique(e.RequestIDs) || e.Eligible != (len(e.RequestIDs) > 0) {
			return fmt.Errorf("invalid eligibility")
		}
		last = e.AlertRecordID
	}
	var previousRequest *Request
	for _, req := range v.Requests {
		if (previousRequest != nil && !lessRequest(*previousRequest, req)) || !validRequest(req) || req.CreatedAt != v.EvaluatedAt {
			return fmt.Errorf("invalid or unordered request")
		}
		copy := req
		previousRequest = &copy
	}
	if v.ID != contentIdentity("notification-plan", v) {
		return fmt.Errorf("invalid notification plan identity")
	}
	return nil
}

func ValidateProviderDescriptor(v ProviderDescriptor) error {
	if v.SchemaName != ProviderSchema || v.SchemaVersion != SchemaVersion || !token(v.ID) || v.Channels == nil || len(v.Channels) == 0 || !sortedChannels(v.Channels) {
		return fmt.Errorf("invalid provider descriptor")
	}
	return nil
}

func ValidateProviderResult(v ProviderResult, startedAt, deadline time.Time) error {
	if v.SchemaName != ProviderResultSchema || v.SchemaVersion != SchemaVersion || !validProviderOutcome(v.Status, v.Failure) || !validTime(v.CompletedAt) || v.CompletedAt.Before(startedAt) || v.CompletedAt.After(deadline) || (v.ProviderReference != "" && !reference(v.ProviderReference)) || v.EvidenceTokens == nil || len(v.EvidenceTokens) > MaxProviderEvidence || !sortedUnique(v.EvidenceTokens) {
		return fmt.Errorf("invalid provider result")
	}
	return nil
}

func ValidateCycleResult(v CycleResult) error {
	if v.SchemaName != CycleResultSchema || v.SchemaVersion != SchemaVersion || !token(v.PlanID) || !validTime(v.ExecutedAt) || v.Attempts == nil {
		return fmt.Errorf("invalid cycle result")
	}
	if err := ValidateQueue(v.NextQueue); err != nil {
		return err
	}
	last := ""
	for _, a := range v.Attempts {
		if a.ID <= last || !validAttempt(a) {
			return fmt.Errorf("invalid cycle attempt")
		}
		last = a.ID
	}
	if v.ID != contentIdentity("notification-cycle", v) {
		return fmt.Errorf("invalid cycle identity")
	}
	return nil
}

func MarshalCanonicalPlan(v Plan) ([]byte, error) {
	if err := ValidatePlan(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func MarshalCanonicalCycle(v CycleResult) ([]byte, error) {
	if err := ValidateCycleResult(v); err != nil {
		return nil, err
	}
	return json.Marshal(v)
}
func DecodePlanningInput(data []byte) (PlanningInput, error) {
	var v PlanningInput
	if err := strictDecode(data, &v); err != nil {
		return PlanningInput{}, err
	}
	return v, ValidatePlanningInput(v)
}
func DecodePlan(data []byte) (Plan, error) {
	var v Plan
	if err := strictDecode(data, &v); err != nil {
		return Plan{}, err
	}
	return v, ValidatePlan(v)
}

func newRequest(record alert.Record, routeID string, endpoint EndpointReference, providerID string, retry RetryPolicy, at time.Time, attempt int) Request {
	deliveryID := stableID("delivery", record.ID, routeID, endpoint.ID)
	evidence := append([]string{}, record.Source.EvidenceReferences...)
	envelope := DeliveryEnvelope{AlertRecordID: record.ID, LifecycleID: record.LifecycleID, ConditionKey: record.ConditionKey, Event: record.Event, Severity: record.Severity, Category: record.Category, ReasonToken: record.ReasonToken, EventTime: record.EventTime, EvidenceReferences: evidence}
	v := Request{SchemaName: RequestSchema, SchemaVersion: SchemaVersion, DeliveryID: deliveryID, IdempotencyKey: deliveryID, AttemptNumber: attempt, RouteID: routeID, EndpointID: endpoint.ID, Channel: endpoint.Channel, ProviderID: providerID, DestinationRef: endpoint.DestinationRef, SecretRef: endpoint.SecretRef, CreatedAt: at, Deadline: record.EventTime.Add(time.Duration(retry.DeliveryWindowNS)), Envelope: envelope}
	v.ID = stableID("request", deliveryID, fmt.Sprintf("%d", attempt))
	return v
}

func planRetry(entry QueueEntry, retry RetryPolicy, at time.Time) (Request, time.Time, DeliveryStatus) {
	if entry.Status != StatusRetryableFailure && entry.Status != StatusIndeterminate && entry.Status != StatusRetryScheduled {
		return Request{}, time.Time{}, ""
	}
	next := len(entry.Attempts) + 1
	if next > retry.MaxAttempts || !at.Before(entry.Request.Deadline) {
		return Request{}, time.Time{}, StatusExhausted
	}
	due := entry.Request.CreatedAt
	if len(entry.Attempts) > 0 {
		due = entry.Attempts[len(entry.Attempts)-1].CompletedAt.Add(time.Duration(retry.BackoffNS[len(entry.Attempts)-1]))
	}
	if !due.Before(entry.Request.Deadline) {
		return Request{}, time.Time{}, StatusExhausted
	}
	if at.Before(due) {
		return Request{}, due, StatusRetryScheduled
	}
	v := entry.Request
	v.AttemptNumber = next
	v.CreatedAt = at
	v.ID = stableID("request", v.DeliveryID, fmt.Sprintf("%d", next))
	return v, time.Time{}, StatusQueued
}

func makeAttempt(req Request, started time.Time, outcome ProviderResult) Attempt {
	status := StatusRecord{SchemaName: StatusSchema, SchemaVersion: SchemaVersion, RequestID: req.ID, Status: outcome.Status, Failure: outcome.Failure, At: outcome.CompletedAt}
	status.ID = contentIdentity("delivery-status", status)
	evidence := EvidenceReference{SchemaName: EvidenceSchema, SchemaVersion: SchemaVersion, AlertRecordID: req.Envelope.AlertRecordID, RequestID: req.ID, ProviderID: req.ProviderID, ProviderReference: outcome.ProviderReference, Tokens: append([]string{}, outcome.EvidenceTokens...)}
	var ack *Acknowledgement
	if outcome.Status == StatusAccepted || outcome.Status == StatusDelivered || outcome.Status == StatusIndeterminate {
		kind := AcknowledgedAccepted
		if outcome.Status == StatusDelivered {
			kind = AcknowledgedDelivered
		}
		if outcome.Status == StatusIndeterminate {
			kind = AcknowledgedUnknown
		}
		v := Acknowledgement{SchemaName: AcknowledgementSchema, SchemaVersion: SchemaVersion, RequestID: req.ID, Kind: kind, At: outcome.CompletedAt, ProviderReference: outcome.ProviderReference}
		v.ID = contentIdentity("delivery-ack", v)
		ack = &v
	}
	v := Attempt{SchemaName: AttemptSchema, SchemaVersion: SchemaVersion, Request: req, StartedAt: started, CompletedAt: outcome.CompletedAt, Status: status, Acknowledgement: ack, Evidence: evidence}
	v.ID = contentIdentity("delivery-attempt", v)
	return v
}

func validRequest(v Request) bool {
	return v.SchemaName == RequestSchema && v.SchemaVersion == SchemaVersion && token(v.ID) && token(v.DeliveryID) && v.IdempotencyKey == v.DeliveryID && v.AttemptNumber >= 1 && v.AttemptNumber <= MaxAttempts && token(v.RouteID) && token(v.EndpointID) && validChannel(v.Channel) && token(v.ProviderID) && reference(v.DestinationRef) && (v.SecretRef == "" || reference(v.SecretRef)) && validTime(v.CreatedAt) && validTime(v.Deadline) && v.CreatedAt.Before(v.Deadline) && validEnvelope(v.Envelope) && v.ID == stableID("request", v.DeliveryID, fmt.Sprintf("%d", v.AttemptNumber))
}
func validEnvelope(v DeliveryEnvelope) bool {
	return token(v.AlertRecordID) && token(v.LifecycleID) && token(v.ConditionKey) && validAlertEvent(v.Event) && validAlertSeverity(v.Severity) && validAlertCategory(v.Category) && token(v.ReasonToken) && validTime(v.EventTime) && v.EvidenceReferences != nil && len(v.EvidenceReferences) <= alert.MaxReferences && sortedUnique(v.EvidenceReferences)
}
func validAttempt(v Attempt) bool {
	if v.SchemaName != AttemptSchema || v.SchemaVersion != SchemaVersion || !validRequest(v.Request) || !validTime(v.StartedAt) || !validTime(v.CompletedAt) || v.CompletedAt.Before(v.StartedAt) || v.Status.RequestID != v.Request.ID || v.Status.SchemaName != StatusSchema || v.Status.SchemaVersion != SchemaVersion || v.Status.ID != contentIdentity("delivery-status", v.Status) || v.Status.At != v.CompletedAt || !validProviderOutcome(v.Status.Status, v.Status.Failure) || v.Evidence.SchemaName != EvidenceSchema || v.Evidence.SchemaVersion != SchemaVersion || v.Evidence.RequestID != v.Request.ID || v.Evidence.AlertRecordID != v.Request.Envelope.AlertRecordID || v.Evidence.ProviderID != v.Request.ProviderID || (v.Evidence.ProviderReference != "" && !reference(v.Evidence.ProviderReference)) || v.Evidence.Tokens == nil || len(v.Evidence.Tokens) > MaxProviderEvidence || !sortedUnique(v.Evidence.Tokens) {
		return false
	}
	if v.Acknowledgement != nil && (v.Acknowledgement.SchemaName != AcknowledgementSchema || v.Acknowledgement.SchemaVersion != SchemaVersion || v.Acknowledgement.RequestID != v.Request.ID || v.Acknowledgement.At != v.CompletedAt || v.Acknowledgement.ProviderReference != v.Evidence.ProviderReference || !validAcknowledgementKind(v.Acknowledgement.Kind) || v.Acknowledgement.ID != contentIdentity("delivery-ack", *v.Acknowledgement)) {
		return false
	}
	if (v.Status.Status == StatusAccepted || v.Status.Status == StatusDelivered || v.Status.Status == StatusIndeterminate) != (v.Acknowledgement != nil) {
		return false
	}
	return v.ID == contentIdentity("delivery-attempt", v)
}

func routeMatches(r Route, a alert.Record) bool {
	return containsSeverity(r.Severities, a.Severity) && containsCategory(r.Categories, a.Category) && containsEvent(r.Events, a.Event)
}
func providerSupports(d ProviderDescriptor, c Channel) bool {
	i := sort.Search(len(d.Channels), func(i int) bool { return d.Channels[i] >= c })
	return i < len(d.Channels) && d.Channels[i] == c
}
func canonicalizePolicy(v *Policy) {
	for i := range v.Routes {
		sort.Slice(v.Routes[i].Severities, func(a, b int) bool { return v.Routes[i].Severities[a] < v.Routes[i].Severities[b] })
		sort.Slice(v.Routes[i].Categories, func(a, b int) bool { return v.Routes[i].Categories[a] < v.Routes[i].Categories[b] })
		sort.Slice(v.Routes[i].Events, func(a, b int) bool { return v.Routes[i].Events[a] < v.Routes[i].Events[b] })
		sort.Strings(v.Routes[i].EndpointIDs)
	}
	sort.Slice(v.Routes, func(i, j int) bool { return v.Routes[i].ID < v.Routes[j].ID })
	sort.Slice(v.Endpoints, func(i, j int) bool { return v.Endpoints[i].ID < v.Endpoints[j].ID })
	sort.Slice(v.Bindings, func(i, j int) bool { return v.Bindings[i].ID < v.Bindings[j].ID })
}
func cloneRoutes(v []Route) []Route {
	out := make([]Route, len(v))
	for i := range v {
		out[i] = v[i]
		out[i].Severities = append([]alert.Severity{}, v[i].Severities...)
		out[i].Categories = append([]alert.Category{}, v[i].Categories...)
		out[i].Events = append([]alert.EventKind{}, v[i].Events...)
		out[i].EndpointIDs = append([]string{}, v[i].EndpointIDs...)
	}
	return out
}
func cloneQueue(v QueueState) QueueState {
	data, _ := json.Marshal(v)
	var out QueueState
	_ = json.Unmarshal(data, &out)
	return out
}
func sortRequests(v []Request) { sort.Slice(v, func(i, j int) bool { return lessRequest(v[i], v[j]) }) }
func sortQueue(v *QueueState) {
	sort.Slice(v.Entries, func(i, j int) bool { return lessQueueEntry(v.Entries[i], v.Entries[j]) })
}
func lessRequest(a, b Request) bool {
	if severityPriority(a.Envelope.Severity) != severityPriority(b.Envelope.Severity) {
		return severityPriority(a.Envelope.Severity) > severityPriority(b.Envelope.Severity)
	}
	if !a.Envelope.EventTime.Equal(b.Envelope.EventTime) {
		return a.Envelope.EventTime.Before(b.Envelope.EventTime)
	}
	return a.ID < b.ID
}
func lessQueueEntry(a, b QueueEntry) bool {
	if severityPriority(a.Request.Envelope.Severity) != severityPriority(b.Request.Envelope.Severity) {
		return severityPriority(a.Request.Envelope.Severity) > severityPriority(b.Request.Envelope.Severity)
	}
	if !a.Request.Envelope.EventTime.Equal(b.Request.Envelope.EventTime) {
		return a.Request.Envelope.EventTime.Before(b.Request.Envelope.EventTime)
	}
	return a.DeliveryID < b.DeliveryID
}
func queueIdentity(v QueueState) string {
	copy := v
	copy.ID = ""
	return identity("notification-queue", copy)
}
func identity(prefix string, v any) string {
	data, _ := json.Marshal(v)
	sum := sha256.Sum256(data)
	return prefix + ":" + hex.EncodeToString(sum[:])
}

func contentIdentity(prefix string, v any) string {
	data, _ := json.Marshal(v)
	var content map[string]any
	_ = json.Unmarshal(data, &content)
	delete(content, "id")
	return identity(prefix, content)
}
func stableID(prefix string, parts ...string) string {
	h := sha256.New()
	for _, p := range parts {
		h.Write([]byte{0})
		h.Write([]byte(p))
	}
	return prefix + ":" + hex.EncodeToString(h.Sum(nil))
}
func token(v string) bool {
	return len(v) > 0 && len(v) <= MaxStringLength && tokenPattern.MatchString(v)
}
func reference(v string) bool    { return token(v) && len(v) <= MaxStringLength }
func validTime(v time.Time) bool { return !v.IsZero() && v.Location() == time.UTC }
func validChannel(v Channel) bool {
	switch v {
	case Email, Webhook, Slack, Discord, Telegram, SMS:
		return true
	}
	return false
}
func validStatus(v DeliveryStatus) bool {
	switch v {
	case StatusQueued, StatusAccepted, StatusDelivered, StatusRetryableFailure, StatusRetryScheduled, StatusTerminalFailure, StatusExhausted, StatusIndeterminate:
		return true
	}
	return false
}
func validFailure(v FailureClass) bool {
	switch v {
	case FailureNone, FailureRetryable, FailureRateLimited, FailureIndeterminate, FailureAuthentication, FailureAuthorization, FailureInvalidDestination, FailureRejectedPayload, FailureUnsupportedProvider, FailureTerminal:
		return true
	}
	return false
}
func validProviderOutcome(s DeliveryStatus, f FailureClass) bool {
	switch s {
	case StatusAccepted, StatusDelivered:
		return f == FailureNone
	case StatusRetryableFailure:
		return f == FailureRetryable || f == FailureRateLimited
	case StatusIndeterminate:
		return f == FailureIndeterminate
	case StatusTerminalFailure:
		return f == FailureAuthentication || f == FailureAuthorization || f == FailureInvalidDestination || f == FailureRejectedPayload || f == FailureUnsupportedProvider || f == FailureTerminal
	}
	return false
}
func validQueueOutcome(s DeliveryStatus, f FailureClass) bool {
	switch s {
	case StatusQueued:
		return f == FailureNone
	case StatusAccepted, StatusDelivered, StatusRetryableFailure, StatusIndeterminate, StatusTerminalFailure:
		return validProviderOutcome(s, f)
	case StatusRetryScheduled, StatusExhausted:
		return f == FailureRetryable || f == FailureRateLimited || f == FailureIndeterminate
	}
	return false
}
func validAcknowledgementKind(v AcknowledgementKind) bool {
	return v == AcknowledgedAccepted || v == AcknowledgedDelivered || v == AcknowledgedUnknown
}
func sortedUnique(v []string) bool {
	last := ""
	for _, x := range v {
		if !token(x) || x <= last {
			return false
		}
		last = x
	}
	return true
}
func sortedChannels(v []Channel) bool {
	last := Channel("")
	for _, x := range v {
		if !validChannel(x) || x <= last {
			return false
		}
		last = x
	}
	return true
}
func sortedAlertSeverities(v []alert.Severity) bool {
	last := alert.Severity("")
	for _, x := range v {
		if !validAlertSeverity(x) || x <= last {
			return false
		}
		last = x
	}
	return true
}
func sortedAlertCategories(v []alert.Category) bool {
	last := alert.Category("")
	for _, x := range v {
		if !validAlertCategory(x) || x <= last {
			return false
		}
		last = x
	}
	return true
}
func sortedAlertEvents(v []alert.EventKind) bool {
	last := alert.EventKind("")
	for _, x := range v {
		if !validAlertEvent(x) || x <= last {
			return false
		}
		last = x
	}
	return true
}
func validAlertSeverity(v alert.Severity) bool {
	return v == alert.Informational || v == alert.Warning || v == alert.Critical || v == alert.Emergency
}
func severityPriority(v alert.Severity) int {
	switch v {
	case alert.Emergency:
		return 4
	case alert.Critical:
		return 3
	case alert.Warning:
		return 2
	case alert.Informational:
		return 1
	}
	return 0
}
func validAlertCategory(v alert.Category) bool {
	switch v {
	case alert.EngineeringCondition, alert.RuleMatch, alert.PolicyGovernance, alert.SchedulerOperation, alert.ReportCompleteness, alert.EvidenceLoss:
		return true
	}
	return false
}
func validAlertEvent(v alert.EventKind) bool {
	switch v {
	case alert.EventEntered, alert.EventEscalated, alert.EventDeescalated, alert.EventAcknowledged, alert.EventSuppressionStarted, alert.EventSuppressionEnded, alert.EventMaintenanceEnded, alert.EventReminder, alert.EventExpired, alert.EventRecovered, alert.EventIndeterminate:
		return true
	}
	return false
}
func containsSeverity(v []alert.Severity, x alert.Severity) bool {
	if len(v) == 0 {
		return true
	}
	i := sort.Search(len(v), func(i int) bool { return v[i] >= x })
	return i < len(v) && v[i] == x
}
func containsCategory(v []alert.Category, x alert.Category) bool {
	if len(v) == 0 {
		return true
	}
	i := sort.Search(len(v), func(i int) bool { return v[i] >= x })
	return i < len(v) && v[i] == x
}
func containsEvent(v []alert.EventKind, x alert.EventKind) bool {
	if len(v) == 0 {
		return true
	}
	i := sort.Search(len(v), func(i int) bool { return v[i] >= x })
	return i < len(v) && v[i] == x
}
func strictDecode(data []byte, target any) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(target); err != nil {
		return fmt.Errorf("invalid notification JSON: %w", err)
	}
	var extra any
	if err := d.Decode(&extra); err != io.EOF {
		if err == nil {
			return fmt.Errorf("unexpected trailing notification JSON")
		}
		return err
	}
	return nil
}
