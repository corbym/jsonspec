# jsonspec
[![Build status](https://travis-ci.org/corbym/jsonspec.svg?branch=master)](https://github.com/corbym/jsonspec)
[![GoDoc](https://godoc.org/github.com/corbym/jsonspec?status.svg)](http://godoc.org/github.com/corbym/jsonspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/corbym/jsonspec)](https://goreportcard.com/report/github.com/corbym/jsonspec)

JSON output generator for the BDD framework [GoGiven](https://github.com/corbym/gogiven)

Import:

```go
import github.com/corbym/jsonspec
```

Usage:

```go
package foo
import (
	"testing"
	"github.com/corbym/gogiven"
	"github.com/corbym/jsonspec"
	"os"
)

func TestMain(testmain *testing.M) {
	gogiven.Generator = jsonspec.NewTestOutputGenerator()
	runOutput := testmain.Run()
	gogiven.GenerateTestOutput()
	os.Exit(runOutput)
}

... actual tests...

```