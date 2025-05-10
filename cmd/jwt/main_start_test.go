package main

import (
	"testing"
	"flag"
	"bytes"
	"strings"
)




func TestStart(t *testing.T) {
	testCases := []struct {
		name          string
		flagSign      string
		flagVerify    string
		flagShow      string
		expectedError string
	}{
		{
			name:          "Test for flagSign with a non-empty value",
			flagSign:      "sign",
			flagVerify:    "",
			flagShow:      "",
			expectedError: "",
		},
		{
			name:          "Test for flagVerify with a non-empty value",
			flagSign:      "",
			flagVerify:    "verify",
			flagShow:      "",
			expectedError: "",
		},
		{
			name:          "Test for flagShow with a non-empty value",
			flagSign:      "",
			flagVerify:    "",
			flagShow:      "show",
			expectedError: "",
		},
		{
			name:          "Test for all flags empty",
			flagSign:      "",
			flagVerify:    "",
			flagShow:      "",
			expectedError: "none of the required flags are present. What do you want me to do?",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			*flagSign = test.flagSign
			*flagVerify = test.flagVerify
			*flagShow = test.flagShow

			var buf bytes.Buffer
			out = &buf

			err := start()

			if err != nil {
				if !strings.Contains(err.Error(), test.expectedError) {
					t.Fatalf("Expected error to contain: %q, but got: %v", test.expectedError, err)
				}
			} else if test.expectedError != "" {
				t.Fatalf("Expected error to contain: %q, but got: %v", test.expectedError, err)
			}
		})
	}
}

