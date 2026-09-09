package update

import "fmt"

type Migration struct {
	ID, From, To, ConfigurationSchema, GuardianSchema, SchedulerSchema, OperatorState string
	Mutation                                                                          bool
	PreserveConfiguration, PreserveCredentials, PreserveState, ReplaceGuardianUnit    bool
}

var migrationPaths = []Migration{
	compatibilityPath("compat-1.3.0-to-1.3.1", "1.3.0", "1.3.1"),
	compatibilityPath("compat-1.2.0-to-1.3.0", "1.2.0", "1.3.0"),
	compatibilityPath("compat-1.1.0-to-1.2.0-rc.1", "1.1.0", "1.2.0-rc.1"),
	compatibilityPath("compat-1.1.0-to-1.2.0-rc.2", "1.1.0", "1.2.0-rc.2"),
	compatibilityPath("compat-1.2.0-rc.1-to-1.2.0-rc.2", "1.2.0-rc.1", "1.2.0-rc.2"),
	compatibilityPath("compat-1.2.0-rc.2-to-1.2.0-rc.7", "1.2.0-rc.2", "1.2.0-rc.7"),
	compatibilityPath("compat-1.2.0-rc.2-to-1.2.0", "1.2.0-rc.2", "1.2.0"),
}

func compatibilityPath(id, from, to string) Migration {
	return Migration{ID: id, From: from, To: to, ConfigurationSchema: "1.0", GuardianSchema: "1.0", SchedulerSchema: "1.0", OperatorState: "1.0-1.2", PreserveConfiguration: true, PreserveCredentials: true, PreserveState: true, ReplaceGuardianUnit: true}
}

func PlanMigration(from, to string) (Migration, error) {
	for _, path := range migrationPaths {
		if path.From == from && path.To == to {
			return path, nil
		}
	}
	fromV, e1 := ParseVersion(from)
	toV, e2 := ParseVersion(to)
	if e1 != nil || e2 != nil || fromV.Major != 1 || toV.Major != 1 {
		return Migration{}, fmt.Errorf("unsupported migration identity")
	}
	if Compare(toV, fromV) <= 0 {
		return Migration{}, fmt.Errorf("migration target is not newer")
	}
	return Migration{}, fmt.Errorf("no deterministic migration path")
}

func (m Migration) Validate() error {
	if m.ID == "" || m.From == "" || m.To == "" || m.ConfigurationSchema != "1.0" || m.GuardianSchema != "1.0" || m.SchedulerSchema != "1.0" || m.OperatorState != "1.0-1.2" || m.Mutation || !m.PreserveConfiguration || !m.PreserveCredentials || !m.PreserveState || !m.ReplaceGuardianUnit {
		return fmt.Errorf("invalid compatibility migration")
	}
	return nil
}

// PreservePackageV1 is the existing bounded package replacement transaction.
// It never interprets metadata as commands or changes configuration/state schemas.
const PreservePackageV1 = "preserve-package-v1"
const MigrationSchema = "qwsg.migration/1"

type CompatibilityDeclaration struct {
	Schema              string `json:"schema"`
	SourceVersion       string `json:"source_version"`
	TargetVersion       string `json:"target_version"`
	Platform            string `json:"platform"`
	Capability          string `json:"capability"`
	ConfigurationSchema string `json:"configuration_schema"`
	GuardianSchema      string `json:"guardian_schema"`
	SchedulerSchema     string `json:"scheduler_schema"`
	OperatorState       string `json:"operator_state"`
}

// PlanCapability validates local executable authority only. Callers must first
// authenticate the containing release and bind its artifact and provenance.
func PlanCapability(from, to, platform string, d CompatibilityDeclaration) (Migration, error) {
	if d.Schema != MigrationSchema || d.Capability != PreservePackageV1 ||
		d.SourceVersion != from || d.TargetVersion != to || d.Platform != platform || platform != "linux-amd64" ||
		d.ConfigurationSchema != "1.0" || d.GuardianSchema != "1.0" || d.SchedulerSchema != "1.0" || d.OperatorState != "1.0-1.2" || Classify(from, to) != Newer {
		return Migration{}, fmt.Errorf("unsupported migration authority or capability")
	}
	return compatibilityPath(PreservePackageV1, from, to), nil
}
