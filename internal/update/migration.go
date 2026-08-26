package update

import "fmt"

type Migration struct {
	From, To, ConfigurationSchema, GuardianSchema, SchedulerSchema, OperatorState string
	Mutation                                                                      bool
}

func PlanMigration(from, to string) (Migration, error) {
	if from == "1.1.0" && to == "1.2.0-rc.1" {
		return Migration{From: from, To: to, ConfigurationSchema: "1.0", GuardianSchema: "1.0", SchedulerSchema: "1.0", OperatorState: "1.0-1.2", Mutation: false}, nil
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
	if m.From == "" || m.To == "" || m.ConfigurationSchema != "1.0" || m.GuardianSchema != "1.0" || m.SchedulerSchema != "1.0" || m.OperatorState != "1.0-1.2" || m.Mutation {
		return fmt.Errorf("invalid compatibility migration")
	}
	return nil
}
