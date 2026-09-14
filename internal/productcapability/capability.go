// Package productcapability owns product authority, independently of presentation,
// operator configuration and release/migration authenticity. It performs no I/O.
package productcapability

import "errors"

const Schema = "qwsg.product-capabilities/1"

type Capability string

const (
	UpdateManual    Capability = "update.manual"
	UpdateAutomatic Capability = "update.automatic"
)

type Product string

const (
	Community Product = "community"
	Pro       Product = "pro"
)

var ErrAuthority = errors.New("invalid or unsupported product capability authority")

// Declaration is input from a trusted entitlement adapter, NOT operator config
// or a license proof. Task 083 has no production Pro adapter or transport.
// A future adapter must authenticate its input before calling Resolve.
type Declaration struct {
	Schema  string
	Product Product
	Grants  []Capability
}

// Set is immutable validated authority. Its zero value grants nothing.
type Set struct {
	product   Product
	manual    bool
	automatic bool
}

func (s Set) Valid() bool      { return s.product == Community || s.product == Pro }
func (s Set) Product() Product { return s.product }
func (s Set) Has(c Capability) bool {
	if !s.Valid() {
		return false
	}
	switch c {
	case UpdateManual:
		return s.manual
	case UpdateAutomatic:
		return s.automatic
	default:
		return false
	}
}

// Resolve treats absent Pro authority as the built-in Community baseline.
// Invalid supplied authority returns a zero Set and never falls back to Pro.
func Resolve(d *Declaration) (Set, error) {
	if d == nil {
		return Set{product: Community, manual: true}, nil
	}
	if d.Schema != Schema || (d.Product != Community && d.Product != Pro) || len(d.Grants) < 1 || len(d.Grants) > 2 {
		return Set{}, ErrAuthority
	}
	s := Set{product: d.Product}
	for _, c := range d.Grants {
		switch c {
		case UpdateManual:
			if s.manual {
				return Set{}, ErrAuthority
			}
			s.manual = true
		case UpdateAutomatic:
			if s.automatic || s.product != Pro {
				return Set{}, ErrAuthority
			}
			s.automatic = true
		default:
			return Set{}, ErrAuthority
		}
	}
	if !s.manual {
		return Set{}, ErrAuthority
	}
	return s, nil
}
