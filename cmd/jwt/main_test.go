package main

import (
	"bytes"
	"errors"
	"flag"
	"strings"
	"testing"
	"github.com/golang-jwt/jwt/v5"
)








/*
ROOST_METHOD_HASH=verifyToken_588ea162aa
ROOST_METHOD_SIG_HASH=verifyToken_d325b8424b

FUNCTION_DEF=func verifyToken() error 

 */
func TestverifyToken(t *testing.T) {
	tests := []struct {
		name           string
		alg            string
		token          string
		key            string
		compact        bool
		debug          bool
		expectedError  string
		expectedOutput string
	}{
		{
			name:           "Scenario1: ValidToken",
			alg:            "RS256",
			token:          "valid.jwt.token",
			key:            "path/to/public.key",
			compact:        false,
			debug:          false,
			expectedOutput: `{"claim": "value"}`,
		},
		{
			name:          "Scenario2: InvalidToken",
			alg:           "RS256",
			token:         "invalid.jwt.token",
			key:           "path/to/public.key",
			expectedError: "couldn't parse token: jwt: parsing error",
		},
		{
			name:          "Scenario3: MissingPublicKey",
			alg:           "RS256",
			token:         "valid.jwt.token",
			key:           "",
			expectedError: "couldn't parse token: open : no such file or directory",
		},
		{
			name:          "Scenario4: UnsupportedAlgorithm",
			alg:           "UNSUPPORTED_ALG",
			token:         "valid.jwt.token",
			key:           "path/to/public.key",
			expectedError: "couldn't parse token: jwt: unsupported signing method: UNSUPPORTED_ALG",
		},
		{
			name:           "Scenario5: DebugModeOutput",
			alg:            "RS256",
			token:          "valid.jwt.token",
			key:            "path/to/public.key",
			debug:          true,
			expectedOutput: "Token len: 11 bytes\nHeader\nClaims",
		},
		{
			name:          "Scenario6: InvalidKeyData",
			alg:           "RS256",
			token:         "valid.jwt.token",
			key:           "invalid.key.data",
			expectedError: "couldn't parse token: key is of invalid type",
		},
		{
			name:           "Scenario7: CompactOutputHandling",
			alg:            "RS256",
			token:          "valid.jwt.token",
			key:            "path/to/public.key",
			compact:        true,
			expectedOutput: `{"claim":"value"}`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			*flagAlg = tt.alg
			*flagKey = tt.key
			*flagVerify = tt.token
			*flagCompact = tt.compact
			*flagDebug = tt.debug

			var stdout, stderr bytes.Buffer
			originalStdout := os.Stdout
			originalStderr := os.Stderr
			defer func() {
				os.Stdout = originalStdout
				os.Stderr = originalStderr
			}()
			os.Stdout = &stdout
			os.Stderr = &stderr

			err := verifyToken()

			if tt.expectedError == "" {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				if !strings.Contains(stdout.String(), tt.expectedOutput) {
					t.Fatalf("expected output %q, got %q", tt.expectedOutput, stdout.String())
				}
			} else {
				if err == nil {
					t.Fatalf("expected error %v, got none", tt.expectedError)
				}
				if !errors.Is(err, jwt.ErrSignatureInvalid) {
					t.Fatalf("expected error %q, got %v", tt.expectedError, err)
				}
			}

			t.Logf("Scenario: %s, Output: %s, Error: %v", tt.name, stdout.String(), err)
		})
	}
}

