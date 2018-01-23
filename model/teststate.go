package model

import (
	"fmt"
	"github.com/corbym/gogiven/generator"
)

type testState struct {
	TestResults       *testResults           `json:"test_results"`
	TestTitle         string                 `json:"test_title"`
	InterestingGivens map[string]interface{} `json:"interesting_givens"`
	CapturedIO        map[string]interface{} `json:"captured_io"`
	GivenWhenThen     []string               `json:"given_when_then"`
}

//NewTestState is internal and creates a new json data object for marshalling test data
func NewTestState(testData generator.TestData) *testState {
	return &testState{
		TestResults:       NewTestResults(testData.TestResult),
		TestTitle:         testData.TestTitle,
		InterestingGivens: convertToMapStringInterface(testData.InterestingGivens),
		CapturedIO:        convertToMapStringInterface(testData.CapturedIO),
		GivenWhenThen:     testData.GivenWhenThen,
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
