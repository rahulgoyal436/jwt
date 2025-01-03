package main

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"testing"
)

// Test data structure for signToken function
type signTokenTestData struct {
	name           string
	tokenData      string
	claims         map[string]string
	signingMethod  string
	key            string
	expectedOutput string
	expectError    bool
}

func TestsignToken(t *testing.T) {
	// Prepare the test data
	testData := []signTokenTestData{
		{
			name:          "Successful Token Signing",
			tokenData:     "validTokenData",
			claims:        map[string]string{"sub": "1234567890", "name": "John Doe", "admin": "true"},
			signingMethod: "HS256",
			key:           "validKey",
			expectError:   false,
		},
		{
			name:          "Invalid Token Data",
			tokenData:     "invalidTokenData",
			claims:        map[string]string{"sub": "1234567890", "name": "John Doe", "admin": "true"},
			signingMethod: "HS256",
			key:           "validKey",
			expectError:   true,
		},
		{
			name:          "Invalid Key",
			tokenData:     "validTokenData",
			claims:        map[string]string{"sub": "1234567890", "name": "John Doe", "admin": "true"},
			signingMethod: "HS256",
			key:           "invalidKey",
			expectError:   true,
		},
		{
			name:          "Invalid Signing Method",
			tokenData:     "validTokenData",
			claims:        map[string]string{"sub": "1234567890", "name": "John Doe", "admin": "true"},
			signingMethod: "InvalidMethod",
			key:           "validKey",
			expectError:   true,
		},
		{
			name:          "Invalid Claims JSON",
			tokenData:     "validTokenData",
			claims:        map[string]string{"sub": "1234567890", "name": "John Doe", "admin": "invalidJson"},
			signingMethod: "HS256",
			key:           "validKey",
			expectError:   true,
		},
	}

	// Run the tests
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Set up the environment for the test
			*flagAlg = tt.signingMethod
			*flagKey = tt.key
			*flagSign = tt.tokenData
			for k, v := range tt.claims {
				flagClaims[k] = v
			}

			// Call the function and check the result
			err := signToken()
			if (err != nil) != tt.expectError {
				t.Errorf("signToken() error = %v, expectError %v", err, tt.expectError)
				return
			}

			// TODO: Validate the output
			// This is left as an exercise for the user as the signToken function doesn't return the token
		})
	}
}
