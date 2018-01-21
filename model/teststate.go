package model

import (
	"fmt"
	"github.com/corbym/gogiven/base"
)

type testState struct {
	TestResults       *testResults           `json:"test_results"`
	TestTitle         string                 `json:"test_title"`
	InterestingGivens map[string]interface{} `json:"interesting_givens"`
	CapturedIO        map[string]interface{} `json:"captured_io"`
	GivenWhenThen     []string               `json:"given_when_then"`
}

//NewTestState is internal and creates a new json data object for marshalling test data
func NewTestState(some *base.Some) *testState {
	return &testState{
		TestResults:       NewTestResults(some.TestingT),
		TestTitle:         some.TestTitle,
		InterestingGivens: convertToMapStringInterface(some.InterestingGivens()),
		CapturedIO:        convertToMapStringInterface(some.CapturedIO()),
		GivenWhenThen:     some.GivenWhenThen,
	}
}

func convertToMapStringInterface(someMap map[interface{}]interface{}) (converted map[string]interface{}) {
	converted = make(map[string]interface{}, len(someMap))
	for k, v := range someMap {
		newKey := fmt.Sprintf("%v", k)
		converted[newKey] = v
	}
	return
}
