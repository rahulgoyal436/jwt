package jwt

import (
	"encoding/json"
	"errors"
	"testing"
	"github.com/golang-jwt/jwt/v5"
)

// TestSignToken is a unit test for signToken function
func TestSignToken(t *testing.T) {
	tests := []struct {
		name         string
		alg          string
		claims       ArgList
		key          string
		sign         string
		expectError  bool
	}{
		{
			name:        "Valid JWT Token Generation",
			alg:         "HS256",
			claims:      ArgList{"foo": "bar"},
			key:         "testKey",
			sign:        "testSign",
			expectError: false,
		},
		{
			name:        "Invalid Claims",
			alg:         "HS256",
			claims:      ArgList{},
			key:         "testKey",
			sign:        "testSign",
			expectError: true,
		},
		{
			name:        "Invalid Key",
			alg:         "HS256",
			claims:      ArgList{"foo": "bar"},
			key:         "",
			sign:        "testSign",
			expectError: true,
		},
		{
			name:        "Unsupported Signing Method",
			alg:         "Unsupported",
			claims:      ArgList{"foo": "bar"},
			key:         "testKey",
			sign:        "testSign",
			expectError: true,
		},
		{
			name:        "Test with 'None' Signing Method",
			alg:         "none",
			claims:      ArgList{"foo": "bar"},
			key:         "testKey",
			sign:        "testSign",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			*flagAlg = &tt.alg
			flagClaims = tt.claims
			*flagKey = &tt.key
			*flagSign = &tt.sign

			// Act
			err := signToken()

			// Assert
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				} else {
					t.Logf("Expected error and got error: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got error: %v", err)
				} else {
					t.Logf("Expected no error and got no error")
				}
			}
		})
	}
}
