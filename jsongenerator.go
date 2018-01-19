package jsonspec

import (
	"github.com/corbym/gogiven/generator"
	"encoding/json"
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

// Generate generates json output for a test. The return string contains the html
// that goes into the output file generated in gogivens.GenerateTestOutput().
// The function panics if the template cannot be generated.
func (outputGenerator *TestOutputGenerator) Generate(pageData *generator.PageData) string {
	b, err := json.Marshal(pageData)
	if err != nil {
		panic(err.Error())
	}
	return string(b[:])
}
