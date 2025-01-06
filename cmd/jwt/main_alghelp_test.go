package main

import (
	"strings"
	"testing"
	"sort"
	"github.com/golang-jwt/jwt/v5"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestalgHelp(t *testing.T) {

	testCases := []struct {
		name     string
		expected string
	}{
		{
			name:     "Normal Operation",
			expected: "signing algorithm identifier, one of\nHS256, HS384, HS512, RS256, RS384, RS512, ES256,\nES384, ES512, PS256, PS384, PS512, EdDSA",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			result := algHelp()

			if result != tc.expected {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}

			if len(strings.Split(result, ",\n")) != 2 {
				t.Errorf("Formatting error: Expected comma and newline after every 7th identifier")
			}

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

