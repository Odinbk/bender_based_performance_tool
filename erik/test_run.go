package erik

func NewTestRun(plans []TestPlan) *TestRun {
	return &TestRun{
		Plans: plans,
	}
}

type TestRun struct {
	Plans []TestPlan
}

func (tr TestRun) Run() {
	for _, tp := range tr.Plans {
		tp.Run()
	}
}