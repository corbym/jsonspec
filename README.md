# jsonspec
[![CI](https://github.com/corbym/jsonspec/actions/workflows/ci.yml/badge.svg)](https://github.com/corbym/jsonspec/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/corbym/jsonspec.svg)](https://pkg.go.dev/github.com/corbym/jsonspec)
[![Go Report Card](https://goreportcard.com/badge/github.com/corbym/jsonspec)](https://goreportcard.com/report/github.com/corbym/jsonspec)

JSON output generator for the BDD framework [GoGiven](https://github.com/corbym/gogiven)

## Installation

```sh
go get github.com/corbym/jsonspec
```

## Usage

```go
package foo

import (
	"os"
	"testing"

	"github.com/corbym/gogiven"
	"github.com/corbym/jsonspec"
)

func TestMain(testmain *testing.M) {
	gogiven.Generator = jsonspec.NewTestOutputGenerator()
	runOutput := testmain.Run()
	gogiven.GenerateTestOutput()
	os.Exit(runOutput)
}

// ... actual tests ...
```

## Example Output
```json
{
  "title": "Generator Test",
  "test_state": [
    {
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
  ]
}
```
