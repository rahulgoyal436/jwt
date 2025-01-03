package main

import (
	"strings"
	"testing"
)

// TestArgListSet is a unit test function for Set function of ArgList type
func TestArgListSet(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name          string
		arg           string
		expectedError string
		expectedKey   string
		expectedValue string
	}{
		{
			name:          "Valid Argument Test",
			arg:           "key=value",
			expectedError: "",
			expectedKey:   "key",
			expectedValue: "value",
		},
		{
			name:          "Invalid Argument Test",
			arg:           "invalidArg",
			expectedError: "invalid argument 'invalidArg'.  Must use format 'key=value'. [invalidArg]",
		},
		{
			name:          "Empty Argument Test",
			arg:           "",
			expectedError: "invalid argument ''.  Must use format 'key=value'. []",
		},
		{
			name:          "Argument with Extra Equals Signs Test",
			arg:           "key=value=value",
			expectedError: "",
			expectedKey:   "key",
			expectedValue: "value=value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			argList := make(ArgList)

			// Act
			err := argList.Set(tc.arg)

			// Assert
			if tc.expectedError == "" {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if argList[tc.expectedKey] != tc.expectedValue {
					t.Errorf("expected key-value pair to be '%v=%v', but got '%v=%v'", tc.expectedKey, tc.expectedValue, tc.expectedKey, argList[tc.expectedKey])
				}
			} else {
				if err == nil || !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("expected error to contain '%v', but got '%v'", tc.expectedError, err)
				}
			}
		})
	}
}
