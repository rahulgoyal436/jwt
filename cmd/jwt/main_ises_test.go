package main

import (
	"testing"
	"flag"
)







func TestIsEs(t *testing.T) {

	tests := []struct {
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			*flagAlg = tc.flagAlg

			result := isEs()

			if result != tc.expected {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}

