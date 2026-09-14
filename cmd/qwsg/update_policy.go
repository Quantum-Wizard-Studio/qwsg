package main

import (
	"fmt"
	"io"
	"os"

	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/configurationstore"
	"quantumwizard.hu/qwsg/internal/productcapability"
	"quantumwizard.hu/qwsg/internal/updatepolicy"
)

// Fixed production authority. Only test-linked code injects Pro declarations.
// Future authenticated entitlement adapters replace this composition seam, not
// the common update engine. Configuration/flags/environment cannot grant Pro.
var installationCapabilities = func() (productcapability.Set, error) {
	return productcapability.Resolve(nil)
}

func effectiveUpdatePolicy(effective configuration.Effective) (updatepolicy.State, error) {
	request, err := configuration.UpdatePolicy(effective.Values)
	if err != nil {
		return updatepolicy.State{}, err
	}
	capabilities, err := installationCapabilities()
	if err != nil {
		return updatepolicy.State{}, err
	}
	return updatepolicy.Evaluate(request, capabilities)
}

func writeUpdatePolicy(out, errout io.Writer) bool {
	path, err := configurationstore.DefaultPath(os.Getenv)
	if err != nil {
		fmt.Fprintln(errout, "Update policy unavailable: configuration path invalid.")
		return false
	}
	source, found, err := configurationstore.Load(path)
	if err != nil {
		fmt.Fprintln(errout, "Update policy unavailable: configuration invalid.")
		return false
	}
	effective, err := resolveLocalConfiguration(source, found, nil)
	if err != nil {
		configFailure(errout, err)
		return false
	}
	state, err := effectiveUpdatePolicy(effective)
	if err != nil {
		configFailure(errout, err)
		return false
	}
	automatic := "unavailable (not entitled)"
	if state.AutomaticAllowed {
		automatic = "permitted (policy only)"
	}
	fmt.Fprintf(out, "Update policy: %s\nProduct: %s\nManual update capability: allowed\nAutomatic update policy: %s\n", state.Effective, state.Product, automatic)
	return true
}
