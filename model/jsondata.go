package model

import "github.com/corbym/gogiven/generator"

//JSONData is the JSON model for the test contemt
type JSONData struct {
	Title     string                `json:"title"`
	TestState map[string]*testState `json:"test_state"`
}

//NewJSONData is internal and creates a new json data object for marshalling test data
func NewJSONData(pageData generator.PageData) *JSONData {
	jsonPageData := new(JSONData)
	jsonPageData.Title = pageData.Title
	jsonPageData.TestState = make(map[string]*testState, len(pageData.TestResults))
	for k, v := range pageData.TestResults {
		jsonPageData.TestState[k] = newTestState(v)
	}
	return jsonPageData
}
