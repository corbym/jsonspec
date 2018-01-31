package model

import "github.com/corbym/gogiven/generator"

type jsonData struct {
	Title     string                `json:"title"`
	TestState map[string]*testState `json:"test_state"`
}

//newJSONData is internal and creates a new json data object for marshalling test data
func NewJSONData(pageData generator.PageData) (jsonPageData *jsonData) {
	jsonPageData = new(jsonData)
	jsonPageData.Title = pageData.Title
	jsonPageData.TestState = make(map[string]*testState, len(pageData.TestResults))
	for k, v := range pageData.TestResults {
		jsonPageData.TestState[k] = newTestState(v)
	}
	return
}
