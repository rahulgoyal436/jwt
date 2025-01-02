package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type ArgList map[string]string


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

			result, err := json.Marshal(tc.argList)

			if !tc.hasError && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if tc.hasError && err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if string(result) != tc.expected {
				t.Fatalf("Expected %s, got %s", tc.expected, result)
			}

			if tc.name == "Large ArgList" && time.Since(start) > time.Second {
				t.Fatalf("Operation took too long")
			}

			t.Logf("Test %s passed", tc.name)
		})
	}
}
func generateLargeArgList() ArgList {
	largeArgList := make(ArgList)

	for i := 0; i < 100000; i++ {
		largeArgList[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	return largeArgList
}
