package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"testing"
	"encoding/json"
	"io"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)








/*
ROOST_METHOD_HASH=verifyToken_588ea162aa
ROOST_METHOD_SIG_HASH=verifyToken_d325b8424b

FUNCTION_DEF=func verifyToken() error 

 */
func TestVerifyToken(t *testing.T) {

	loadData = loadDataMock

	tests := []struct {
		name       string
		verifyPath string
		keyPath    string
		alg        string
		debug      bool
		expected   string
	}{
		{
			name:       "Successfully Verify RSA Signed JWT Token",
			verifyPath: "validToken",
			keyPath:    "validKey",
			alg:        "RS256",
			debug:      false,
			expected:   "",
		},
		{
			name:       "Fail Verification with Invalid RSA Key",
			verifyPath: "validToken",
			keyPath:    "invalidKey",
			alg:        "RS256",
			debug:      false,
			expected:   "couldn't parse token: key is of invalid type",
		},
		{
			name:       "Handle None Algorithm Safely",
			verifyPath: "noneAlgToken",
			keyPath:    "",
			alg:        "none",
			debug:      false,
			expected:   "",
		},
		{
			name:       "Display Debug Information for Valid Token",
			verifyPath: "validToken",
			keyPath:    "validKey",
			alg:        "RS256",
			debug:      true,
			expected:   "",
		},
		{
			name:       "Fail Verification with Invalid Token Format",
			verifyPath: "badFormatToken",
			keyPath:    "validKey",
			alg:        "RS256",
			debug:      false,
			expected:   "couldn't parse token",
		},
		{
			name:       "Edge Case - Empty Token File",
			verifyPath: "emptyToken",
			keyPath:    "validKey",
			alg:        "RS256",
			debug:      false,
			expected:   "couldn't read token",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			*flagVerify = tc.verifyPath
			*flagKey = tc.keyPath
			*flagAlg = tc.alg
			*flagDebug = tc.debug

			var buf bytes.Buffer
			oldStderr := os.Stderr
			defer func() {
				os.Stderr = oldStderr
			}()
			os.Stderr = &buf

			err := verifyToken()
			output := buf.String()

			if (err != nil && err.Error() != tc.expected) || (err == nil && tc.expected != "") {
				t.Errorf("expected error '%v', got '%v'", tc.expected, err)
			}

			if tc.debug && err == nil && !strings.Contains(output, "Header:") {
				t.Errorf("expected debug output to include 'Header:', got '%v'", output)
			}

			t.Logf("Completed test case: %s", tc.name)
		})
	}
}

func loadDataMock(path string) ([]byte, error) {
	switch path {
	case "validToken":
		return []byte(`valid.jwt.token`), nil
	case "validKey":
		return []byte(`valid RSA key data`), nil
	case "invalidKey":
		return []byte(`invalid RSA key data`), nil
	case "noneAlgToken":
		return []byte(`none.algorithm.token`), nil
	case "badFormatToken":
		return []byte(`bad.format.token`), nil
	case "emptyToken":
		return []byte(``), nil
	default:
		return nil, errors.New("loadDataMock: unknown path")
	}
}


/*
ROOST_METHOD_HASH=signToken_23bc8d51bd
ROOST_METHOD_SIG_HASH=signToken_ead4688e18

FUNCTION_DEF=func signToken() error 

 */
func TestSignToken(t *testing.T) {

	flag.Set("compact", "true")
	flag.Set("debug", "false")

	tests := []struct {
		name       string
		flagSign   string
		flagKey    string
		flagAlg    string
		flagClaims ArgList
		flagDebug  bool
		wantErr    string
	}{
		{
			name:     "Successfully Sign a Token with Valid Claims and Key",
			flagSign: `{"sub": "1234567890", "name": "John Doe", "iat": 1516239022}`,
			flagKey:  "validRSAPrivateKey",
			flagAlg:  "RS256",
			wantErr:  "",
		},
		{
			name:     "Fail to Sign a Token due to Missing Claims",
			flagSign: "",
			flagKey:  "validRSAPrivateKey",
			flagAlg:  "RS256",
			wantErr:  "couldn't read token",
		},
		{
			name:     "Fail to Sign a Token due to Invalid Claims JSON",
			flagSign: `{"sub": "1234567890", "name": "John Doe", iat: 1516239022}`,
			flagKey:  "validRSAPrivateKey",
			flagAlg:  "RS256",
			wantErr:  "couldn't parse claims JSON",
		},
		{
			name:     "Fail to Sign a Token due to Unsupported Signing Method",
			flagSign: `{"sub": "1234567890", "name": "John Doe", "iat": 1516239022}`,
			flagKey:  "validRSAPrivateKey",
			flagAlg:  "foobar",
			wantErr:  "couldn't find signing method",
		},
		{
			name:       "Successfully Merge Additional Claims",
			flagSign:   `{"sub": "1234567890", "name": "John Doe", "iat": 1516239022}`,
			flagKey:    "validRSAPrivateKey",
			flagAlg:    "RS256",
			flagClaims: ArgList{"role": "admin"},
			wantErr:    "",
		},
		{
			name:     "Handle Unsupported Key Data for Signing",
			flagSign: `{"sub": "1234567890", "name": "John Doe", "iat": 1516239022}`,
			flagKey:  "incorrectFormatKey",
			flagAlg:  "RS256",
			wantErr:  "couldn't convert key data to key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			flag.Set("sign", tt.flagSign)
			flag.Set("key", tt.flagKey)
			flag.Set("alg", tt.flagAlg)
			for k, v := range tt.flagClaims {
				flagClaims[k] = v
			}

			var outBuf bytes.Buffer
			fmtPrintln = func(a ...interface{}) (n int, err error) {
				return fmt.Fprintln(&outBuf, a...)
			}

			err := signToken()

			if tt.wantErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				assert.NoError(t, err)

				tokenParts := strings.Split(outBuf.String(), ".")
				assert.Len(t, tokenParts, 3, "Expecting a valid JWT with 3 parts")
			}
		})
	}
}

