package main

import (
	"encoding/json"
	"testing"
	"time"
	"fmt"
)

type ArgList map[string]string


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}

type Time struct {
	// wall and ext encode the wall time seconds, wall time nanoseconds,
	// and optional monotonic clock reading in nanoseconds.
	//
	// From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
	// a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
	// The nanoseconds field is in the range [0, 999999999].
	// If the hasMonotonic bit is 0, then the 33-bit field must be zero
	// and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
	// If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
	// unsigned wall seconds since Jan 1 year 1885, and ext holds a
	// signed 64-bit monotonic clock reading, nanoseconds since process start.
	wall uint64
	ext  int64

	// loc specifies the Location that should be used to
	// determine the minute, hour, month, day, and year
	// that correspond to this Time.
	// The nil location means UTC.
	// All UTC times are represented with loc==nil, never loc==&utcLoc.
	loc *Location
}

func TestArgListString(t *testing.T) {

	testCases := []struct {
		name     string
		argList  ArgList
		expected string
		hasError bool
	}{
		{

			name:     "Valid ArgList",
			argList:  ArgList{"key1": "value1", "key2": "value2"},
			expected: `{"key1":"value1","key2":"value2"}`,
			hasError: false,
		},
		{

			name:     "Empty ArgList",
			argList:  ArgList{},
			expected: `{}`,
			hasError: false,
		},
		{

			name:     "Large ArgList",
			argList:  generateLargeArgList(),
			expected: "",
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			start := time.Now()

			result := tc.argList.String()

			if tc.hasError {
				if json.Valid([]byte(result)) {
					t.Errorf("expected non-serializable ArgList to produce invalid JSON, got: %s", result)
				}
			} else {
				if result != tc.expected {
					t.Errorf("expected: %s, got: %s", tc.expected, result)
				}

				if tc.name == "Large ArgList" && time.Since(start) > time.Second {
					t.Errorf("operation took too long")
				}
			}
		})
	}
}
func generateLargeArgList() ArgList {
	argList := make(ArgList)
	for i := 0; i < 1000000; i++ {
		argList[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}
	return argList
}
