package jsonspec_test

import (
	"bytes"
	"encoding/json"
	"github.com/corbym/gocrest"
	"github.com/corbym/gocrest/is"
	. "github.com/corbym/gocrest/then"
	"github.com/corbym/gogiven/base"
	"github.com/corbym/gogiven/generator"
	"github.com/corbym/gogiven/testdata"
	"github.com/corbym/jsonspec"
	"testing"
)

var jsonString string
var underTest *jsonspec.TestOutputGenerator

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
				"foo": {
					"test_results": {
						"id": "foo",
						"failed": false,
						"skipped": false,
						"test_output": ""
					},
					"test_title": "Generator Test",
					"interesting_givens": {
						"faff": "flap"
					},
					"captured_io": {
						"foob": "barb"
					},
					"given_when_then": "GivenWhenThen"
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
			data := newPageData(newSomeMap())

			html := underTest.Generate(data)
			AssertThat(testing, html, is.ValueContaining("Generator Test"))
		}()
	}
}

func TestTestOutputGenerator_FileExtension(t *testing.T) {
	AssertThat(t, underTest.FileExtension(), is.EqualTo(".json"))
}

func TestTestOutputGenerator_Panics(t *testing.T) {
	defer func() {
		recovered := recover()
		AssertThat(t, recovered, is.Not(is.Nil()))
	}()
	data := newPageData(newSomeMap())
	data.SomeMap = nil

	underTest.Generate(data)
}

func fileIsConverted() {
	jsonString = underTest.Generate(newPageData(newSomeMap()))
}

func newSomeMap() *base.SomeMap {
	testingT := new(base.TestMetaData)
	some := base.NewSome(testingT,
		"Generator Test",
		base.NewTestMetaData("foo"),
		"GivenWhenThen",
		func(givens testdata.InterestingGivens) {
			givens["faff"] = "flap"
		})
	some.CapturedIO()["foob"] = "barb"
	someMap := &base.SomeMap{"foo": some}
	return someMap
}

func newPageData(someMap *base.SomeMap) *generator.PageData {
	return &generator.PageData{
		Title:   "Generator Test",
		SomeMap: someMap,
	}
}
