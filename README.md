# jsonspec
[![Build status](https://travis-ci.org/corbym/jsonspec.svg?branch=master)](https://github.com/corbym/jsonspec)
[![GoDoc](https://godoc.org/github.com/corbym/jsonspec?status.svg)](http://godoc.org/github.com/corbym/jsonspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/corbym/jsonspec)](https://goreportcard.com/report/github.com/corbym/jsonspec)
[![Coverage Status](https://coveralls.io/repos/github/corbym/jsonspec/badge.svg?branch=master)](https://coveralls.io/github/corbym/jsonspec?branch=master)

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

## Example Output
```json
{
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
      "given_when_then": {
        "given_when_then": [
          "given",
          "when",
          "then"
        ],
        "comment": [
          "Fooing is best",
          "done with friends"
        ]
      }
    }
  }
}
```
