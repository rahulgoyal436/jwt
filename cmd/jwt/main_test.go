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
	"encoding/json"
	"errors"
	"io/ioutil"
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


/*
ROOST_METHOD_HASH=signToken_23bc8d51bd
ROOST_METHOD_SIG_HASH=signToken_ead4688e18

FUNCTION_DEF=func signToken() error 

 */
func TestSignToken(t *testing.T) {
	type scenario struct {
		name            string
		alg             string
		claims          map[string]interface{}
		keyFileContent  string
		expectError     bool
		expectedErrMsg  string
		headers         map[string]interface{}
		outputValidator func(string) error
	}

	scenarios := []scenario{
		{
			name: "Successfully signing a token with a valid key and claims",
			alg:  "HS256",
			claims: map[string]interface{}{
				"user": "testUser",
			},
			keyFileContent: "mysecretkey",
			expectError:    false,
			outputValidator: func(output string) error {

				if output == "" {
					return errors.New("empty token generated")
				}
				return nil
			},
		},
		{
			name:           "Fail to sign token due to missing signing method",
			alg:            "UNSUPPORTED",
			expectError:    true,
			expectedErrMsg: "couldn't find signing method: UNSUPPORTED",
		},
		{
			name:           "Fail to sign token due to invalid JSON in claims",
			alg:            "HS256",
			expectError:    true,
			expectedErrMsg: "couldn't parse claims JSON",
		},
		{
			name:           "Fail to sign token due to invalid key data",
			alg:            "HS256",
			keyFileContent: "invalidKeyData",
			expectError:    true,
			expectedErrMsg: "couldn't read key",
		},
		{
			name: "Successfully sign a token with additional header fields",
			alg:  "HS256",
			claims: map[string]interface{}{
				"user": "testUser",
			},
			keyFileContent: "mysecretkey",
			headers: map[string]interface{}{
				"extra": "headerInfo",
			},
			expectError: false,
			outputValidator: func(output string) error {

				if output == "" {
					return errors.New("empty token generated")
				}
				return nil
			},
		},
		{
			name: "Ensure None method is handled correctly",
			alg:  "none",
			claims: map[string]interface{}{
				"user": "testUser",
			},
			expectError: false,
			outputValidator: func(output string) error {

				if output == "" {
					return errors.New("empty token generated with 'none' signing method")
				}
				return nil
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			var stdout bytes.Buffer
			old := os.Stdout
			defer func() { os.Stdout = old }()
			os.Stdout = &stdout

			flag.Set("alg", s.alg)
			tmpKeyFile, err := ioutil.TempFile("", "keyfile")
			if err != nil {
				t.Fatalf("error creating temp key file: %v", err)
			}
			defer os.Remove(tmpKeyFile.Name())

			if s.keyFileContent != "" {
				if _, err := tmpKeyFile.Write([]byte(s.keyFileContent)); err != nil {
					t.Fatalf("error writing to temp key file: %v", err)
				}
				flag.Set("key", tmpKeyFile.Name())
			} else {
				flag.Set("key", "")
			}

			if s.claims != nil {
				tmpSignFile, err := ioutil.TempFile("", "signfile")
				if err != nil {
					t.Fatalf("error creating temp sign file: %v", err)
				}
				defer os.Remove(tmpSignFile.Name())

				claimsData, _ := json.Marshal(s.claims)
				if _, err := tmpSignFile.Write(claimsData); err != nil {
					t.Fatalf("error writing to temp sign file: %v", err)
				}
				flag.Set("sign", tmpSignFile.Name())
			} else {
				flag.Set("sign", "invalid_path")
			}

			if s.headers != nil {
				for k, v := range s.headers {
					flagHead.Set(k, v.(string))
				}
			}

			err = signToken()
			gotOutput := stdout.String()

			if s.expectError {
				t.Log("Expecting error scenario")
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				if !strings.Contains(err.Error(), s.expectedErrMsg) {
					t.Fatalf("expected error message to contain: %v got: %v", s.expectedErrMsg, err.Error())
				}
			} else {
				t.Log("Expecting successful signing scenario")
				if err != nil {
					t.Fatalf("expected no error, but got: %v", err)
				}
				if validateErr := s.outputValidator(gotOutput); validateErr != nil {
					t.Fatalf("validation of token output failed: %v", validateErr)
				}
			}

			t.Logf("Scenario %s passed", s.name)
		})
	}
}

func init() {

	flag.Parse()
}

