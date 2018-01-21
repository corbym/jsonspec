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

## Example Output
```json
{
	"title": "Clock Test",
	"test_state": {
		"github.com/corbym/berlinclock_test.TestGivenAClockWhenTimeIsEnteredThenCorrectClock": {
			"test_results": {
				"id": "github.com/corbym/berlinclock_test.TestGivenAClockWhenTimeIsEnteredThenCorrectClock",
				"failed": false,
				"skipped": false,
				"test_output": ""
			},
			"test_title": "Given A Clock When Time Is Entered Then Correct Clock",
			"interesting_givens": {
				"expected": "YOOOOOOOOOOOOOOOOOOOOOOO",
				"time": "00:00:00"
			},
			"captured_io": {},
			"given_when_then": [
				"Given testing clock parameters under test",
				"When",
				"Clock err = berlinclock clock test time",
				"Then",
				"Then assert that testing err is nil",
				"Then assert that testing clock is",
				"Equal to test expected",
				"Reasonf \"time incorrect for % s\" test time"
			]
		},
		"github.com/corbym/berlinclock_test.TestGivenAClockWhenTimeIsEnteredThenCorrectClock_1": {
			"test_results": {
				"id": "github.com/corbym/berlinclock_test.TestGivenAClockWhenTimeIsEnteredThenCorrectClock_1",
				"failed": false,
				"skipped": false,
				"test_output": ""
			},
			"test_title": "Given A Clock When Time Is Entered Then Correct Clock",
			"interesting_givens": {
				"expected": "ORRRRRRROYYRYYRYYRYYYYYY",
				"time": "23:59:59"
			},
			"captured_io": {},
			"given_when_then": [
				"Given testing clock parameters under test",
				"When",
				"Clock err = berlinclock clock test time",
				"Then",
				"Then assert that testing err is nil",
				"Then assert that testing clock is",
				"Equal to test expected",
				"Reasonf \"time incorrect for % s\" test time"
			]
		},
		"github.com/corbym/berlinclock_test.TestGivenAClockWhenTimeIsEnteredThenCorrectClock_1_1": {
			"test_results": {
				"id": "github.com/corbym/berlinclock_test.TestGivenAClockWhenTimeIsEnteredThenCorrectClock_1_1",
				"failed": false,
				"skipped": false,
				"test_output": ""
			},
			"test_title": "Given A Clock When Time Is Entered Then Correct Clock",
			"interesting_givens": {
				"expected": "YRRROROOOYYRYYRYYRYOOOOO",
				"time": "16:50:06"
			},
			"captured_io": {},
			"given_when_then": [
				"Given testing clock parameters under test",
				"When",
				"Clock err = berlinclock clock test time",
				"Then",
				"Then assert that testing err is nil",
				"Then assert that testing clock is",
				"Equal to test expected",
				"Reasonf \"time incorrect for % s\" test time"
			]
		},
		"github.com/corbym/berlinclock_test.TestGivenAClockWhenTimeIsEnteredThenCorrectClock_1_1_1": {
			"test_results": {
				"id": "github.com/corbym/berlinclock_test.TestGivenAClockWhenTimeIsEnteredThenCorrectClock_1_1_1",
				"failed": false,
				"skipped": false,
				"test_output": ""
			},
			"test_title": "Given A Clock When Time Is Entered Then Correct Clock",
			"interesting_givens": {
				"expected": "ORROOROOOYYRYYRYOOOOYYOO",
				"time": "11:37:01"
			},
			"captured_io": {},
			"given_when_then": [
				"Given testing clock parameters under test",
				"When",
				"Clock err = berlinclock clock test time",
				"Then",
				"Then assert that testing err is nil",
				"Then assert that testing clock is",
				"Equal to test expected",
				"Reasonf \"time incorrect for % s\" test time"
			]
		}
	}
}```
