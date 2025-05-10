package main

import (
	"strings"
	"testing"
)

func TestisEs(t *testing.T) {

	tests := []struct {
		name     string
		arg      string
		expected bool
	}{
		{
			name:     "Scenario 1: Normal operation with an algorithm that starts with 'ES'",
			arg:      "ES256",
			expected: true,
		},
		{
			name:     "Scenario 2: Normal operation with an algorithm that does not start with 'ES'",
			arg:      "RS256",
			expected: false,
		},
		{
			name:     "Scenario 3: Edge case with an empty algorithm string",
			arg:      "",
			expected: false,
		},
	}

	for _, test := range tests {

		flagAlg = &test.arg

		result := isEs()

		if result != test.expected {
			t.Errorf("Test case %s failed: got %v, expected %v", test.name, result, test.expected)
		} else {
			t.Logf("Test case %s passed", test.name)
		}
	}
}
