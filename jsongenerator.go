package jsonspec

import (
	"github.com/corbym/gogiven/generator"
	"encoding/json"
	"bytes"
	"github.com/corbym/gogiven/base"
	"fmt"
)

//TestOutputGenerator is an implementation of the GoGivensOutputGenerator that generates an html file per
// test. It is thread safe between goroutines.
type TestOutputGenerator struct {
	generator.GoGivensOutputGenerator
}

//NewTestOutputGenerator creates a template that is used to generate the html output.
func NewTestOutputGenerator() *TestOutputGenerator {
	outputGenerator := new(TestOutputGenerator)
	return outputGenerator
}

// FileExtension for the output generated.
func (outputGenerator *TestOutputGenerator) FileExtension() string {
	return ".json"
}

type TestResults struct {
	Id         string `json:"id"`
	Failed     bool   `json:"failed"`
	Skipped    bool   `json:"skipped"`
	TestOutput string `json:"test_output"`
}

type TestState struct {
	TestResults       *TestResults           `json:"test_results"`
	TestTitle         string                 `json:"test_title"`
	InterestingGivens map[string]interface{} `json:"interesting_givens"`
	CapturedIO        map[string]interface{} `json:"captured_io"`
	GivenWhenThen     string                 `json:"given_when_then"`
}

type JsonData struct {
	Title     string                `json:"title"`
	TestState map[string]*TestState `json:"test_state"`
}

func NewJsonData(pageData *generator.PageData) (jsonData *JsonData) {
	jsonData = new(JsonData)
	jsonData.Title = pageData.Title
	jsonData.TestState = make(map[string]*TestState, len(*pageData.SomeMap))
	for k, v := range *pageData.SomeMap {
		jsonData.TestState[k] = NewTestState(v)
	}
	return
}

func NewTestState(some *base.Some) *TestState {
	return &TestState{
		TestResults:       NewTestingT(some.TestingT),
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

func NewTestingT(data *base.TestMetaData) *TestResults {
	return &TestResults{Failed: data.Failed(),
		TestOutput: data.TestOutput(),
		Id: data.TestId,
		Skipped: data.Skipped(),
	}
}

// Generate generates json output for a test. The return string contains the html
// that goes into the output file generated in gogivens.GenerateTestOutput().
// The function panics if the template cannot be generated.
func (outputGenerator *TestOutputGenerator) Generate(pageData *generator.PageData) string {

	b, err := json.Marshal(NewJsonData(pageData))
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	if err != nil {
		panic(err.Error())
	}
	return out.String()
}
