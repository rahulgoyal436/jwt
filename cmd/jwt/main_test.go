package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
	"github.com/golang-jwt/jwt/v5"
)








/*
ROOST_METHOD_HASH=verifyToken_588ea162aa
ROOST_METHOD_SIG_HASH=verifyToken_d325b8424b

FUNCTION_DEF=func verifyToken() error 

 */
func TestVerifyToken(t *testing.T) {
	type testCase struct {
		name       string
		alg        string
		key        string
		verifyPath string
		expected   error
		compact    bool
		debug      bool
	}

	tests := []testCase{
		{
			name:       "Verifying a Valid Token Successfully",
			alg:        "RS256",
			key:        "testdata/valid-public.pem",
			verifyPath: "testdata/valid-token.jwt",
			expected:   nil,
			compact:    true,
			debug:      false,
		},
		{
			name:       "Handling an Invalid Token",
			alg:        "RS256",
			key:        "testdata/valid-public.pem",
			verifyPath: "testdata/invalid-token.jwt",
			expected:   fmt.Errorf("couldn't parse token: %w", jwt.NewValidationError("", jwt.ValidationErrorMalformed)),
			compact:    false,
			debug:      false,
		},
		{
			name:       "Missing Public Key",
			alg:        "RS256",
			key:        "",
			verifyPath: "testdata/valid-token.jwt",
			expected:   fmt.Errorf("couldn't parse token: %w", fmt.Errorf("no path specified")),
			compact:    false,
			debug:      false,
		},
		{
			name:       "Unsupported Algorithm",
			alg:        "HS256",
			key:        "testdata/valid-public.pem",
			verifyPath: "testdata/invalid-algorithm.jwt",
			expected:   fmt.Errorf("couldn't parse token: %w", jwt.NewValidationError("", jwt.ValidationErrorUnverifiable)),
			compact:    false,
			debug:      false,
		},
		{
			name:       "Debug Mode Output",
			alg:        "RS256",
			key:        "testdata/valid-public.pem",
			verifyPath: "testdata/valid-token.jwt",
			expected:   nil,
			compact:    false,
			debug:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			*flagAlg = tc.alg
			*flagKey = tc.key
			*flagVerify = tc.verifyPath
			*flagCompact = tc.compact
			*flagDebug = tc.debug

			var stdout, stderr bytes.Buffer
			stdoutBackup := os.Stdout
			stderrBackup := os.Stderr
			defer func() {
				os.Stdout = stdoutBackup
				os.Stderr = stderrBackup
			}()
			os.Stdout = &stdout
			os.Stderr = &stderr

			err := verifyToken()

			if (err != nil && tc.expected == nil) || (err == nil && tc.expected != nil) || (err != nil && tc.expected != nil && !strings.Contains(err.Error(), tc.expected.Error())) {
				t.Fatalf("Expected error: %v, got: %v", tc.expected, err)
			}

			if tc.debug {
				if stderr.Len() == 0 {
					t.Log("Debug data was not printed")
				}
			}
		})
	}
}

