package configuration

import "quantumwizard.hu/qwsg/internal/updatepolicy"

// UpdatePolicy extracts canonical intent. Absence preserves existing manual
// behavior; duplicates and malformed declarations refuse even on raw Models.
// Entitlement is checked separately by updatepolicy.Evaluate.
func UpdatePolicy(model Model) (updatepolicy.Request, error) {
	result := updatepolicy.Request{Mode: updatepolicy.Manual}
	found := false
	for _, extension := range model.Extensions {
		if extension.ID != updatepolicy.ExtensionID {
			continue
		}
		if found {
			return updatepolicy.Request{}, updatepolicy.ErrPolicy
		}
		found = true
		var err error
		result, err = updatepolicy.Parse(extension.Version, extension.Required, extension.Fields)
		if err != nil {
			return updatepolicy.Request{}, err
		}
	}
	return result, nil
}
