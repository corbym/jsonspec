package jsonspec_test

import (
	"bytes"
	"encoding/json"
	"github.com/corbym/gocrest"
	"github.com/corbym/gocrest/is"
	. "github.com/corbym/gocrest/then"
	"github.com/corbym/gogiven/generator"
	"github.com/corbym/jsonspec"
	"testing"
)

var jsonString string
var underTest generator.GoGivensOutputGenerator

func init() {
	underTest = jsonspec.NewTestOutputGenerator()
}

func TestTestOutputGenerator_Generate(testing *testing.T) {
	fileIsConverted()
	AssertThat(testing, jsonString, isValidJson())
	AssertThat(testing, jsonString, is.EqualToIgnoringWhitespace(
		`{
			"title": "Generator Test",
			"test_state": {
				"test title": {
					"test_results": {
						"id": "abc2124",
						"failed": true,
						"skipped": true,
						"test_output": "well alrighty then"
					},
					"test_title": "test title",
					"interesting_givens": {
						"faff": "flap"
					},
					"captured_io": {
						"foob": "barb"
					},
					"given_when_then": [
						"given",
						"when",
						"then"
					]
				}
			}
		}`))
}

func isValidJson() *gocrest.Matcher {
	matcher := &gocrest.Matcher{Describe: "valid jsonString"}
	matcher.Matches = func(actual interface{}) bool {
		buffer := &bytes.Buffer{}
		buffer.WriteString(actual.(string))
		var f interface{}
		error := json.Unmarshal(buffer.Bytes(), &f)
		return error == nil
	}
	return matcher
}

func TestTestOutputGenerator_GenerateConcurrently(testing *testing.T) {
	for i := 0; i < 15; i++ {
		go func() {
			data := newPageData(false, false)

			json := underTest.Generate(data)
			buffer := new(bytes.Buffer)
			buffer.ReadFrom(json)
			AssertThat(testing, buffer.String(), is.ValueContaining("Generator Test"))
		}()
	}
}

func TestTestOutputGenerator_FileExtension(t *testing.T) {
	AssertThat(t, underTest.ContentType(), is.EqualTo("text/json"))
}

func TestTestOutputGenerator_Panics(t *testing.T) {
	defer func() {
		recovered := recover()
		AssertThat(t, recovered, is.Not(is.Nil()))
	}()
	underTest.Generate(nil)
}

func fileIsConverted() {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(underTest.Generate(newPageData(true, true)))
	jsonString = buffer.String()
}

func newPageData(skipped bool, failed bool) *generator.PageData {
	testData := make(map[string]*generator.TestData)
	capturedIO := make(map[interface{}]interface{})
	capturedIO["foob"] = "barb"
	interestingGivens := make(map[interface{}]interface{})
	interestingGivens["faff"] = "flap"
	testData["test title"] = &generator.TestData{
		TestTitle:         "test title",
		GivenWhenThen:     []string{"given", "when", "then"},
		CapturedIO:        capturedIO,
		InterestingGivens: interestingGivens,
		TestResult: &generator.TestResult{
			Failed:     failed,
			Skipped:    skipped,
			TestOutput: "well alrighty then",
			TestId:     "abc2124",
		},
	}
	return &generator.PageData{
		TestResults: testData,
		Title:       "Generator Test",
	}
}

