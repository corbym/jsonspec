package model

import (
	"fmt"
	"github.com/corbym/gogiven/generator"
	"github.com/corbym/gogiven/base"
)

type ParsedTestContent struct {
	GivenWhenThen []string `json:"given_when_then"`
	// Comment contains each comment line from the tests' comment block
	Comment []string `json:"comment"`
}

type testState struct {
	TestResults       testResults            `json:"test_results"`
	TestTitle         string                 `json:"test_title"`
	InterestingGivens map[string]interface{} `json:"interesting_givens"`
	CapturedIO        map[string]interface{} `json:"captured_io"`
	GivenWhenThen     ParsedTestContent      `json:"given_when_then"`
}

//NewTestState is internal and creates a new json data object for marshalling test data
func NewTestState(testData generator.TestData) *testState {
	return &testState{
		TestResults:       NewTestResults(testData.TestResult),
		TestTitle:         testData.TestTitle,
		InterestingGivens: convertToMapStringInterface(testData.InterestingGivens),
		CapturedIO:        convertToMapStringInterface(testData.CapturedIO),
		GivenWhenThen:     NewParsedTestContent(testData.ParsedTestContent),
	}
}
func NewParsedTestContent(content base.ParsedTestContent) ParsedTestContent {
	return ParsedTestContent{
		GivenWhenThen: content.GivenWhenThen,
		Comment:       content.Comment,
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
