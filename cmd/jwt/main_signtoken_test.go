package jwt

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"os"
	"strings"
	"testing"
)

// TestSignToken tests the signToken function.
func TestSignToken(t *testing.T) {
	tests := []struct {
		name        string
		setup       func()
		expectedErr error
	}{
		{
			name: "Successful token signing",
			setup: func() {
				*flagSign = "+"
				*flagAlg = "HS256"
				*flagKey = "testdata/valid_key.pem"
			},
			expectedErr: nil,
		},
		{
			name: "Invalid token data",
			setup: func() {
				*flagSign = "testdata/invalid_token.json"
				*flagAlg = "HS256"
				*flagKey = "testdata/valid_key.pem"
			},
			expectedErr: errors.New("couldn't parse claims JSON: invalid character 'i' looking for beginning of value"),
		},
		{
			name: "Invalid key data",
			setup: func() {
				*flagSign = "+"
				*flagAlg = "HS256"
				*flagKey = "testdata/invalid_key.pem"
			},
			expectedErr: errors.New("couldn't read key: open testdata/invalid_key.pem: no such file or directory"),
		},
		{
			name: "No signing method",
			setup: func() {
				*flagSign = "+"
				*flagAlg = "INVALID"
				*flagKey = "testdata/valid_key.pem"
			},
			expectedErr: errors.New("couldn't find signing method: INVALID"),
		},
		{
			name: "Error signing token",
			setup: func() {
				*flagSign = "+"
				*flagAlg = "none"
				*flagKey = "testdata/valid_key.pem"
			},
			expectedErr: errors.New("error signing token: signing method (none) does not support key type: []uint8"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			// Capture stdout
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := signToken()

			// Stop capturing stdout
			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = old

			if tt.expectedErr == nil {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}

				if !jwtRegex.Match(out) {
					t.Errorf("Output is not a valid JWT: %s", out)
					return
				}

				t.Logf("Success: %s", out)
			} else {
				if err == nil || !strings.Contains(err.Error(), tt.expectedErr.Error()) {
					t.Errorf("Expected error: %v, got: %v", tt.expectedErr, err)
					return
				}

				t.Logf("Expected error: %v", err)
			}
		})
	}
}

// jwtRegex matches a JWT.
var jwtRegex = regexp.MustCompile(`^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.?[A-Za-z0-9-_.+/=]*$`)
