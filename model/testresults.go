package model

import "github.com/corbym/gogiven/base"

type testResults struct {
	Id         string `json:"id"`
	Failed     bool   `json:"failed"`
	Skipped    bool   `json:"skipped"`
	TestOutput string `json:"test_output"`
}

//NewTestResults is internal and creates a new json data object for marshalling test data
func NewTestResults(data *base.TestMetaData) *testResults {
	return &testResults{Failed: data.Failed(),
		TestOutput: data.TestOutput(),
		Id:         data.TestId,
		Skipped:    data.Skipped(),
	}
}
