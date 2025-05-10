package main

import (
	"testing"
	"flag"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestisEs(t *testing.T) {

	testCases := []struct {
		name     string
		flagAlg  string
		expected bool
	}{
		{
			name:     "Normal operation with an algorithm that starts with 'ES'",
			flagAlg:  "ES256",
			expected: true,
		},
		{
			name:     "Normal operation with an algorithm that does not start with 'ES'",
			flagAlg:  "HS256",
			expected: false,
		},
		{
			name:     "Edge case with an empty algorithm string",
			flagAlg:  "",
			expected: false,
		},

		{
			name:     "Edge case with a nil algorithm string",
			flagAlg:  "",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			*flagAlg = tc.flagAlg

			result := isEs()

			if result != tc.expected {
				t.Fatalf("Expected %v but got %v", tc.expected, result)
			} else {
				t.Logf("Success: Expected %v and got %v", tc.expected, result)
			}
		})
	}
}

