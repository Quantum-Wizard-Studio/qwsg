// Package updatepolicy validates update intent against product capabilities.
// Policy eligibility never authenticates a release or executes an update.
package updatepolicy

import (
	"errors"
	"quantumwizard.hu/qwsg/internal/productcapability"
)

const ExtensionID = "installer.update-policy"
const Version = "1.0"

type Mode string

const (
	Manual    Mode = "manual"
	Automatic Mode = "automatic"
)

var (
	ErrPolicy               = errors.New("invalid or unsupported update policy")
	ErrAutomaticUnavailable = errors.New("automatic update policy unavailable: update.automatic capability required")
)

// Request preserves the legacy notify option as manual installation with notices.
type Request struct {
	Mode   Mode
	Notify bool
}

// Parse validates the existing configuration extension without granting authority.
func Parse(version string, required bool, fields map[string]string) (Request, error) {
	if version != Version || required || len(fields) != 1 {
		return Request{}, ErrPolicy
	}
	switch fields["policy"] {
	case "manual":
		return Request{Mode: Manual}, nil
	case "notify":
		return Request{Mode: Manual, Notify: true}, nil
	case "automatic":
		return Request{Mode: Automatic}, nil
	default:
		return Request{}, ErrPolicy
	}
}

// State is operator-visible eligibility only, never an executable authorization.
type State struct {
	Product          productcapability.Product `json:"product"`
	Effective        Mode                      `json:"effective_policy"`
	ManualAllowed    bool                      `json:"manual_allowed"`
	AutomaticAllowed bool                      `json:"automatic_allowed"`
	Notify           bool                      `json:"notify"`
}

func Evaluate(request Request, capabilities productcapability.Set) (State, error) {
	if !capabilities.Valid() || !capabilities.Has(productcapability.UpdateManual) {
		return State{}, productcapability.ErrAuthority
	}
	if request.Mode != Manual && request.Mode != Automatic || request.Notify && request.Mode != Manual {
		return State{}, ErrPolicy
	}
	if request.Mode == Automatic && !capabilities.Has(productcapability.UpdateAutomatic) {
		return State{}, ErrAutomaticUnavailable
	}
	return State{Product: capabilities.Product(), Effective: request.Mode, ManualAllowed: true, AutomaticAllowed: capabilities.Has(productcapability.UpdateAutomatic), Notify: request.Notify}, nil
}
