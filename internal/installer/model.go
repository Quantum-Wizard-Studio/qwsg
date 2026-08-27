// Package installer owns the interface-neutral guided installation contract.
package installer

import "fmt"

type PhaseID string

const (
	PhasePreflight     PhaseID = "preflight"
	PhasePlan          PhaseID = "plan"
	PhaseInstall       PhaseID = "install"
	PhaseConfiguration PhaseID = "configuration"
	PhaseNotification  PhaseID = "notification"
	PhaseUpdatePolicy  PhaseID = "update_policy"
	PhaseActivation    PhaseID = "activation"
	PhaseReadiness     PhaseID = "readiness"
	PhaseCompletion    PhaseID = "completion"
)

type Phase struct {
	ID     PhaseID
	Weight int
}

var Phases = []Phase{
	{PhasePreflight, 12}, {PhasePlan, 10}, {PhaseInstall, 20},
	{PhaseConfiguration, 15}, {PhaseNotification, 10}, {PhaseUpdatePolicy, 8},
	{PhaseActivation, 12}, {PhaseReadiness, 10}, {PhaseCompletion, 3},
}

type State string

const (
	Pending   State = "pending"
	Active    State = "active"
	Completed State = "completed"
	Failed    State = "failed"
	Restored  State = "restored"
)

type Progress struct {
	States map[PhaseID]State
	Active PhaseID
}

// Percent counts only completed phases. Active or failed work is never
// presented as complete; completion is therefore exactly 100 percent.
func (p Progress) Percent() int {
	total := 0
	for _, phase := range Phases {
		if p.States[phase.ID] == Completed {
			total += phase.Weight
		}
	}
	return total
}

func (p *Progress) Start(id PhaseID) error {
	if !knownPhase(id) {
		return fmt.Errorf("unknown installation phase %q", id)
	}
	p.Active = id
	p.States[id] = Active
	return nil
}

func (p *Progress) Complete(id PhaseID) error {
	if p.States[id] != Active {
		return fmt.Errorf("phase %q is not active", id)
	}
	p.States[id] = Completed
	p.Active = ""
	return nil
}

func (p *Progress) Fail(id PhaseID, restored bool) {
	if restored {
		p.States[id] = Restored
	} else {
		p.States[id] = Failed
	}
	p.Active = id
}

func NewProgress() Progress {
	states := make(map[PhaseID]State, len(Phases))
	for _, phase := range Phases {
		states[phase.ID] = Pending
	}
	return Progress{States: states}
}

func knownPhase(id PhaseID) bool {
	for _, phase := range Phases {
		if phase.ID == id {
			return true
		}
	}
	return false
}

type InstallState string

const (
	Fresh    InstallState = "fresh"
	Existing InstallState = "existing"
	Upgrade  InstallState = "upgrade"
)

type UpdatePolicy string

const (
	UpdateManual UpdatePolicy = "manual"
	UpdateNotify UpdatePolicy = "notify"
)

type Plan struct {
	State        InstallState
	PackageWrite bool
	ConfigWrite  bool
	ServiceStart bool
	Paths        []string
}
