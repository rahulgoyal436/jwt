package main

import (
	"strings"
	"testing"
)








func TestArgListSet(t *testing.T) {

	testCases := []struct {
		name          string
		arg           string
		expectedKey   string
		expectedValue string
		expectedError bool
	}{
		{
			name:          "Valid Argument Test",
			arg:           "key=value",
			expectedKey:   "key",
			expectedValue: "value",
			expectedError: false,
		},
		{
			name:          "Invalid Argument Test",
			arg:           "invalid",
			expectedError: true,
		},
		{
			name:          "Empty Argument Test",
			arg:           "",
			expectedError: true,
		},
		{
			name:          "Argument with Extra Equals Signs Test",
			arg:           "key=value=value",
			expectedKey:   "key",
			expectedValue: "value=value",
			expectedError: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			argList := make(ArgList)

			err := argList.Set(testCase.arg)

			if testCase.expectedError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
				if err != nil && !strings.Contains(err.Error(), "invalid argument") {
					t.Errorf("expected 'invalid argument' error but got '%v'", err)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got '%v'", err)
				}

				if val, ok := argList[testCase.expectedKey]; !ok {
					t.Errorf("expected key '%v' not found in ArgList", testCase.expectedKey)
				} else if val != testCase.expectedValue {
					t.Errorf("expected value '%v' but got '%v'", testCase.expectedValue, val)
				}
			}
		})
	}
}
