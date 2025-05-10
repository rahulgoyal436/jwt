package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
	"github.com/golang-jwt/jwt/v5"
)



type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestshowToken(t *testing.T) {
	testCases := []struct {
		name        string
		inputToken  string
		inputDebug  bool
		expectedErr string
	}{
		{
			name:        "Successful token reading and printing",
			inputToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			inputDebug:  false,
			expectedErr: "",
		},
		{
			name:        "Invalid token input",
			inputToken:  "invalid-token",
			inputDebug:  false,
			expectedErr: "malformed token",
		},
		{
			name:        "Token reading error",
			inputToken:  "",
			inputDebug:  false,
			expectedErr: "no path specified",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			*flagShow = tc.inputToken
			*flagDebug = tc.inputDebug

			err := showToken()

			if err != nil && err.Error() != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}



