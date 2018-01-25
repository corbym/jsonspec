package jsonspec

import (
	"bytes"
	"encoding/json"
	"github.com/corbym/gogiven/generator"
	"github.com/corbym/jsonspec/model"
	"io"
)

//TestOutputGenerator is an implementation of the GoGivensOutputGenerator that generates a JSON file per
// test. It is thread safe between goroutines.
type TestOutputGenerator struct {
	generator.GoGivensOutputGenerator
	MarshalJson func(v interface{}) ([]byte, error)
}

//NewTestOutputGenerator creates a template that is used to generate the json output.
func NewTestOutputGenerator() *TestOutputGenerator {
	outputGenerator := new(TestOutputGenerator)
	outputGenerator.MarshalJson = json.Marshal
	return outputGenerator
}

// FileExtension for the output generated.
func (outputGenerator *TestOutputGenerator) ContentType() string {
	return "application/json"
}

// Generate generates json output for a test. The return string contains the html
// that goes into the output file generated in gogivens.GenerateTestOutput().
// The function panics if the template cannot be generated.
func (outputGenerator *TestOutputGenerator) Generate(pageData *generator.PageData) (io.Reader) {
	jsonBytes, err := outputGenerator.MarshalJson(model.NewJsonData(pageData))
	if err != nil {
		panic("Could not marshal pageData to json")
	}
	var out = new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "\t")
	return out
}
