// Package setupflow derives resumable setup progress from canonical readiness
// evidence. It stores no wizard state and performs no actions.
package setupflow

import (
	"fmt"
	"time"

	"quantumwizard.hu/qwsg/internal/assessment"
)

const (
	SchemaName = "qwsg.setup-flow"
	Version    = "1.0"
)

type ActionKind string

const (
	ActionNone         ActionKind = "none"
	ActionConfigure    ActionKind = "operator_input"
	ActionRecommend    ActionKind = "recommended_command"
	ActionVerify       ActionKind = "manual_verification"
	ActionTest         ActionKind = "notification_test"
	ActionActivate     ActionKind = "guardian_activation"
	ActionWaitEvidence ActionKind = "wait_for_evidence"
)

type Step struct {
	Phase                string                  `json:"phase"`
	State                assessment.SummaryState `json:"state"`
	Action               ActionKind              `json:"action"`
	ReasonToken          string                  `json:"reason_token"`
	RequiresInput        bool                    `json:"requires_input"`
	RequiresConfirmation bool                    `json:"requires_confirmation"`
}

type Plan struct {
	SchemaName    string                     `json:"schema_name"`
	SchemaVersion string                     `json:"schema_version"`
	AssessedAt    time.Time                  `json:"assessed_at"`
	Steps         []Step                     `json:"steps"`
	NextAction    Step                       `json:"next_action"`
	Readiness     []assessment.DomainSummary `json:"readiness"`
}

func Build(report assessment.Report) (Plan, error) {
	if err := assessment.ValidateReport(report); err != nil {
		return Plan{}, err
	}
	domains := map[string]assessment.SummaryState{}
	for _, d := range report.Domains {
		domains[d.Domain] = d.State
	}
	steps := []Step{
		step("environment", domains["environment_dependencies"], ActionVerify, "verify_environment", false, false),
		step("configuration", domains["configuration"], ActionConfigure, "configure_qwsg", true, false),
		step("notification", domains["notification"], ActionTest, "configure_and_test_notification", true, true),
		step("guardian_activation", domains["guardian_service"], ActionActivate, "activate_guardian", false, true),
	}
	for i := range steps {
		if steps[i].State == assessment.Ready {
			steps[i].Action, steps[i].ReasonToken = ActionNone, "already_satisfied"
		}
	}
	next := Step{Phase: "complete", State: assessment.Ready, Action: ActionNone, ReasonToken: "no_action_required"}
	for _, s := range steps {
		if s.State != assessment.Ready {
			next = s
			break
		}
	}
	return Plan{SchemaName: SchemaName, SchemaVersion: Version, AssessedAt: report.AssessedAt, Steps: steps, NextAction: next, Readiness: report.Domains}, nil
}

func step(phase string, state assessment.SummaryState, action ActionKind, reason string, input, confirm bool) Step {
	if state == "" {
		state = assessment.Unknown
	}
	return Step{Phase: phase, State: state, Action: action, ReasonToken: reason, RequiresInput: input, RequiresConfirmation: confirm}
}

func Validate(plan Plan) error {
	if plan.SchemaName != SchemaName || plan.SchemaVersion != Version || plan.AssessedAt.IsZero() || len(plan.Steps) != 4 {
		return fmt.Errorf("invalid setup plan")
	}
	return nil
}
