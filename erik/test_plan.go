package erik

type TestPlan struct {
	TestCases []ITestCase
}

func (tp TestPlan) Run () {
	for _, tc := range tp.TestCases {
		tc.Run()
	}
}
