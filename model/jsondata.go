package model

import "github.com/corbym/gogiven/generator"

type jsonData struct {
	Title     string                `json:"title"`
	TestState map[string]*testState `json:"test_state"`
}

//NewJsonData is internal and creates a new json data object for marshalling test data
func NewJsonData(pageData *generator.PageData) (jsonPageData *jsonData) {
	jsonPageData = new(jsonData)
	jsonPageData.Title = pageData.Title
	jsonPageData.TestState = make(map[string]*testState, len(pageData.TestResults))
	for k, v := range pageData.TestResults {
		jsonPageData.TestState[k] = NewTestState(*v)
	}
	return
}
