package main

import (
	"strings"
	"testing"
	"github.com/golang-jwt/jwt/v5"
)

// TestalgHelp is a test function for algHelp function
func TestalgHelp(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		expected string
	}{
		{
			name:     "Normal Operation",
			expected: "HS256, HS384, HS512, RS256, RS384, RS512, ES256,\nES384, ES512, PS256, PS384, PS512, EdDSA",
		},
	}

	for _, tc := range testCases {
		// Run the test case
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := algHelp()

			// Assert
			if result != tc.expected {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}

			// Testing Formatting
			if len(strings.Split(result, ",\n")) != 2 {
				t.Errorf("Formatting error: Expected comma and newline after every 7th identifier")
			}

			// Testing Sorting
			algs := strings.Split(strings.ReplaceAll(result, ",\n", ", "), ", ")
			for i := 0; i < len(algs)-1; i++ {
				if algs[i] > algs[i+1] {
					t.Errorf("Sorting error: Identifiers are not sorted")
					break
				}
			}
		})
	}
}
