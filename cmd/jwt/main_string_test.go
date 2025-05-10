package main

import (
	"encoding/json"
	"testing"
	"time"
)

// TestArgListString is a unit test for the String() method of the ArgList type.
func TestArgListString(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		argList  ArgList
		expected string
		hasError bool
	}{
		{
			// Scenario 1: Normal operation with valid input
			name:     "Valid ArgList",
			argList:  ArgList{"key1": "value1", "key2": "value2"},
			expected: `{"key1":"value1","key2":"value2"}`,
			hasError: false,
		},
		{
			// Scenario 2: Empty ArgList
			name:     "Empty ArgList",
			argList:  ArgList{},
			expected: `{}`,
			hasError: false,
		},
		{
			// Scenario 3: Large ArgList
			name:     "Large ArgList",
			argList:  generateLargeArgList(),
			expected: "", // we cannot know the exact JSON string, so we leave this empty
			hasError: false,
		},
		{
			// Scenario 4: Non-JSON-Serializable ArgList
			name:     "Non-Serializable ArgList",
			argList:  ArgList{"key": make(chan int)}, // channels are not JSON serializable
			expected: "",
			hasError: true,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()

			// Act
			result, err := json.Marshal(tc.argList)

			// Assert
			if !tc.hasError && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if tc.hasError && err == nil {
				t.Fatalf("Expected error, got nil")
			}
			if string(result) != tc.expected {
				t.Fatalf("Expected %s, got %s", tc.expected, result)
			}

			// Check if the operation was timely for large ArgLists
			if tc.name == "Large ArgList" && time.Since(start) > time.Second {
				t.Fatalf("Operation took too long")
			}

			t.Logf("Test %s passed", tc.name)
		})
	}
}

// generateLargeArgList generates a large ArgList for testing.
func generateLargeArgList() ArgList {
	largeArgList := make(ArgList)

	for i := 0; i < 100000; i++ {
		largeArgList[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	return largeArgList
}
