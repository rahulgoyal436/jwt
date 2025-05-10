package main

import (
	"bytes"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

type testCase struct {
	name          string
	tokenPath     string
	debugFlag     bool
	expectedError string
}
func TestshowToken(t *testing.T) {

	testCases := []testCase{
		{
			name:          "Successful token reading and printing",
			tokenPath:     "path_to_valid_token",
			debugFlag:     false,
			expectedError: "",
		},
		{
			name:          "Invalid token input",
			tokenPath:     "path_to_invalid_token",
			debugFlag:     false,
			expectedError: "malformed token",
		},
		{
			name:          "Token reading error",
			tokenPath:     "non_existent_path",
			debugFlag:     false,
			expectedError: "couldn't read token",
		},
		{
			name:          "Debug flag is true",
			tokenPath:     "path_to_valid_token",
			debugFlag:     true,
			expectedError: "",
		},
		{
			name:          "Error in printing JSON",
			tokenPath:     "path_to_token_with_unprintable_json",
			debugFlag:     false,
			expectedError: "failed to output claims",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			*flagShow = tc.tokenPath
			*flagDebug = tc.debugFlag
			err := showToken()
			if err != nil {
				if tc.expectedError == "" || !strings.Contains(err.Error(), tc.expectedError) {
					t.Errorf("unexpected error: got %v, want %v", err, tc.expectedError)
				}
			} else if tc.expectedError != "" {
				t.Errorf("expected error: got %v, want %v", err, tc.expectedError)
			}
		})
	}
}
