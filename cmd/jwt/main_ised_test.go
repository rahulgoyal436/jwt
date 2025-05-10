package main

import (
	"flag"
	"testing"
)







func TestisEd(t *testing.T) {

	tests := []struct {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			flagAlg = flag.String("alg", tt.flagAlg, algHelp())

			result := isEd()

			if result != tt.expected {
				t.Errorf("isEd() = %v, want %v", result, tt.expected)
			} else {
				t.Logf("Success: Expected output %v and got %v", tt.expected, result)
			}
		})
	}
}
