package main

import (
	"flag"
	"testing"
)







func TestIsRs(t *testing.T) {
	tests := []struct {
		name    string
		flagAlg string
		want    bool
	}{
		{
			name:    "Scenario 1: Algorithm Prefix is 'RS'",
			flagAlg: "RS256",
			want:    true,
		},
		{
			name:    "Scenario 2: Algorithm Prefix is 'PS'",
			flagAlg: "PS256",
			want:    true,
		},
		{
			name:    "Scenario 3: Algorithm Prefix is neither 'RS' nor 'PS'",
			flagAlg: "HS256",
			want:    false,
		},
		{
			name:    "Scenario 4: Algorithm Prefix is empty",
			flagAlg: "",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			*flagAlg = tt.flagAlg

			got := isRs()

			if got != tt.want {
				t.Errorf("isRs() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Success: %s", tt.name)
			}
		})
	}
}

