package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"testing"
)

// TestSignToken will test the function signToken
func TestSignToken(t *testing.T) {
	tests := []struct {
		name       string
		flagSign   string
		flagKey    string
		flagAlg    string
		wantErr    bool
	}{
		{
			name:     "Successful JWT Token Generation",
			flagSign: "validClaims.json",
			flagKey:  "validKey.pem",
			flagAlg:  "RS256",
			wantErr:  false,
		},
		{
			name:     "Invalid flagSign Input",
			flagSign: "invalidClaims.json",
			flagKey:  "validKey.pem",
			flagAlg:  "RS256",
			wantErr:  true,
		},
		{
			name:     "Invalid flagKey Input",
			flagSign: "validClaims.json",
			flagKey:  "invalidKey.pem",
			flagAlg:  "RS256",
			wantErr:  true,
		},
		{
			name:     "Unsupported flagAlg",
			flagSign: "validClaims.json",
			flagKey:  "validKey.pem",
			flagAlg:  "unsupportedAlg",
			wantErr:  true,
		},
		{
			name:     "Invalid Key Type",
			flagSign: "validClaims.json",
			flagKey:  "validKey.pem",
			flagAlg:  "HS256",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			*flagSign = tt.flagSign
			*flagKey = tt.flagKey
			*flagAlg = tt.flagAlg

			// Act
			err := signToken()

			// Assert
			if (err != nil) != tt.wantErr {
				t.Errorf("signToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
