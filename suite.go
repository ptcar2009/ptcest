package pkg

import (
	"bytes"
	"fmt"
)

type TestSuite struct {
	beforeEach func(t *T)
	afterEach  func(t *T)
	beforeAll  func(t *T)
	afterAll   func(t *T)
	cases      []TestCase
	format     string
}

type TestCase interface {
	run(*T)
}

type Case struct {
	runf   func(*T)
	format string
}

func Then(what string, check func(*T)) *Case {
	return &Case{
		format: fmt.Sprintf("->\033[34;1m Then\033[0m %v", what),
		runf:   check,
	}
}
func (c *Case) run(t *T) {
	t.printHeader(c.format)
	c.runf(t)
	if t.failed {
		t.failHeader()
	} else {
		t.okHeader()
	}
}

func When(what string, cases ...TestCase) *TestSuite {

	ret := &TestSuite{
		format: fmt.Sprintf("->\033[34;1m When\033[0m %v", what),
		cases:  cases,
	}
	return ret
}
func Given(what string, cases ...TestCase) *TestSuite {

	ret := &TestSuite{
		format: fmt.Sprintf("->\033[34;1m Given\033[0m %v", what),
		cases:  cases,
	}
	return ret
}

func (t *T) printHeader(format string) {
	errStr := ""
	for i := 0; i < t.level; i++ {
		errStr += "    "
	}
	fmt.Fprintf(t.writer, "%v%v", errStr, format)
}

func (s *TestSuite) run(t *T) {
	t.printHeader(s.format)
	first := true
	rwba := bytes.NewBufferString("")
	if s.beforeAll != nil {
		s.beforeAll(&T{
			level:      t.level,
			writerLogs: rwba,
		})
	}
	if s.afterAll != nil {
		defer s.afterAll(t)
	}

	for _, c := range s.cases {
		// create a different function so that defer works properly
		//
		rw := bytes.NewBufferString("")
		rwl := bytes.NewBufferString("")
		rwb := bytes.NewBufferString("")
		tt := &T{
			level:      t.level + 1,
			writer:     rw,
			writerBody: rwb,
			writerLogs: rwl,
		}
		func() {
			if s.beforeEach != nil {
				s.beforeEach(
					&T{
						level:      t.level + 1,
						writerLogs: rwl,
					},
				)
			}
			if s.afterEach != nil {
				defer s.afterEach(&T{
					level:      t.level + 1,
					writerLogs: rwl,
				})
			}

			c.run(tt)
		}()

		if tt.failed {
			if first {
				t.failHeader()
				t.writerBody.Write(rwba.Bytes())
				first = false
			}
			t.writerBody.Write(rw.Bytes())
			t.writerBody.Write(rwb.Bytes())
			t.writerBody.Write(rwl.Bytes())
		} else {
			if first {
				t.okHeader()
				first = false
			}
			t.writerBody.Write(rw.Bytes())
			t.writerBody.Write(rwb.Bytes())

		}
		t.failed = t.failed || tt.failed
	}
}

var baseSuite TestSuite = TestSuite{}

func Scenario(what string, cases ...TestCase) *TestSuite {

	ret := &TestSuite{
		format: fmt.Sprintf("->\033[34;1m Scenario\033[0m %v", what),
		cases:  cases,
	}
	baseSuite.cases = append(baseSuite.cases, ret)
	return ret
}

var _ = Scenario("a running testing environment",
	Given("a scenario",
		When("that scenario runs",

			Then("it shall succeed", func(t *T) {
				t.Equals(true, true)
			}),
			Then("this is a then", func(t *T) {
				t.Equals(true, true)
			}),
			Then("it shall fail", func(t *T) {

				t.IsNotNil(nil)
			}),
		).BeforeEach(func(t *T) {
			t.Infof("before each")
		}).AfterEach(func(t *T) {
			t.Infof("after each")
		}),
	),
).BeforeAll(func(t *T) {
	t.Infof("this is before")
}).AfterAll(func(t *T) {
	t.Infof("this is after")
})

var _ = Scenario("this is a scenario",
	Given("this is a given",
		When("this is a when",
			Then("this is a then", func(t *T) {
				t.Equals(true, true)
			}),
			Then("this is a then", func(t *T) {
				t.Equals(true, true)
			}),
			Then("this is a then", func(t *T) {
				t.Equals(true, true)
			}),
		).BeforeEach(func(t *T) {
			t.Infof("before each")
		}).AfterEach(func(t *T) {
			t.Infof("after each")
		}),
	),
).BeforeAll(func(t *T) {
	t.Infof("this is before")
}).AfterAll(func(t *T) {
	t.Infof("this is after")
})

func (s *TestSuite) BeforeEach(f func(t *T)) *TestSuite {
	s.beforeEach = f
	return s
}

func (s *TestSuite) AfterEach(f func(t *T)) *TestSuite {
	s.afterEach = f
	return s
}

func (s *TestSuite) BeforeAll(f func(t *T)) *TestSuite {
	s.beforeAll = f
	return s
}

func (s *TestSuite) AfterAll(f func(t *T)) *TestSuite {
	s.afterAll = f
	return s
}
