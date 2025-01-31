package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSignToken tests the signToken function.
func TestSignToken(t *testing.T) {
	// Define the test cases
	tests := []struct {
		name        string
		alg         string
		claims      ArgList
		key         string
		sign        string
		expectedErr error
	}{
		{
			name:        "Successfully Signing a Token with Valid Parameters",
			alg:         "HS256",
			claims:      ArgList{"name": "John Doe"},
			key:         "valid-key",
			sign:        "valid-sign",
			expectedErr: nil,
		},
		{
			name:        "SignToken Function with Invalid Token Data",
			alg:         "HS256",
			claims:      ArgList{"name": "John Doe"},
			key:         "valid-key",
			sign:        "invalid-sign",
			expectedErr: errors.New("couldn't read token: invalid-sign"),
		},
		{
			name:        "SignToken Function with Invalid Key Data",
			alg:         "HS256",
			claims:      ArgList{"name": "John Doe"},
			key:         "invalid-key",
			sign:        "valid-sign",
			expectedErr: errors.New("couldn't read key: invalid-key"),
		},
		{
			name:        "SignToken Function with Nonexistent Signing Method",
			alg:         "NonexistentAlg",
			claims:      ArgList{"name": "John Doe"},
			key:         "valid-key",
			sign:        "valid-sign",
			expectedErr: errors.New("couldn't find signing method: NonexistentAlg"),
		},
		{
			name:        "SignToken Function with Incompatible Key and Signing Method",
			alg:         "HS256",
			claims:      ArgList{"name": "John Doe"},
			key:         "incompatible-key",
			sign:        "valid-sign",
			expectedErr: errors.New("couldn't convert key data to key"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			*flagAlg = tt.alg
			flagClaims = tt.claims
			*flagKey = tt.key
			*flagSign = tt.sign

			// Act
			err := signToken()

			// Assert
			if tt.expectedErr != nil {
				require.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
