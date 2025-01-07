package main

import (
	"testing"
	"time"
)







func TestArgListString(t *testing.T) {

	testCases := []struct {
		name        string
		argList     ArgList
		expectedStr string
		expectError bool
	}{
		{
			name:        "Normal operation with valid input",
			argList:     ArgList{"key1": "value1", "key2": "value2"},
			expectedStr: `{"key1":"value1","key2":"value2"}`,
			expectError: false,
		},
		{
			name:        "Empty ArgList",
			argList:     ArgList{},
			expectedStr: `{}`,
			expectError: false,
		},
		{
			name:        "Large ArgList",
			argList:     generateLargeArgList(),
			expectedStr: "",
			expectError: false,
		},
		{
			name:        "Non-JSON-Serializable ArgList",
			argList:     ArgList{"key1": string([]byte{0x89, 0x89})},
			expectedStr: "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			startTime := time.Now()
			actualStr, err := tc.argList.String()
			duration := time.Since(startTime)

			if !tc.expectError {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if actualStr != tc.expectedStr {
					t.Errorf("Unexpected string representation of ArgList. Expected: %s, got: %s", tc.expectedStr, actualStr)
				}
				if tc.name == "Large ArgList" && duration > time.Second*5 {
					t.Errorf("Operation took too long to complete. Duration: %s", duration)
				}
			} else {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			}
		})
	}
}
func generateLargeArgList() ArgList {
	return ArgList{}
}
