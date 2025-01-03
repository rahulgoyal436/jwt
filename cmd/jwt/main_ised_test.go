package main

import (
	"flag"
	"testing"
)

func TestisEd(t *testing.T) {

	testCases := []struct {
		name     string
		flagAlg  string
		expected bool
	}{
		{
			name:     "Testing for EdDSA algorithm",
			flagAlg:  "EdDSA",
			expected: true,
		},
		{
			name:     "Testing for non-EdDSA algorithm",
			flagAlg:  "RS256",
			expected: false,
		},
		{
			name:     "Testing for empty algorithm",
			flagAlg:  "",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			*flagAlg = tc.flagAlg

			result := isEd()

			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

