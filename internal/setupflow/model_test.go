package setupflow

import (
	"quantumwizard.hu/qwsg/internal/assessment"
	"testing"
	"time"
)

func TestPlanIsDerivedAndResumable(t *testing.T) {
	r := assessment.Report{SchemaName: assessment.SchemaName, SchemaVersion: assessment.SchemaVersion, ModelVersion: assessment.ModelVersion, RegistryVersion: assessment.RegistryVersion, AssessedAt: time.Now().UTC(), Phase: "operational", Domains: []assessment.DomainSummary{{Domain: "environment_dependencies", State: assessment.Ready}, {Domain: "configuration", State: assessment.NotReady}, {Domain: "notification", State: assessment.Partial}, {Domain: "guardian_service", State: assessment.NotReady}}}
	p, err := Build(r)
	if err != nil || Validate(p) != nil {
		t.Fatalf("plan=%+v err=%v", p, err)
	}
	if p.NextAction.Phase != "configuration" || !p.NextAction.RequiresInput {
		t.Fatalf("next=%+v", p.NextAction)
	}
}
