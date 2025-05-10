package main

import (
	"flag"
	"testing"
)







func TestisNone(t *testing.T) {
	tests := []struct {
		name    string
		flagAlg string
		want    bool
	}{
		{
			name:    "Scenario 1: FlagAlg is set to 'none'",
			flagAlg: "none",
			want:    true,
		},
		{
			name:    "Scenario 2: FlagAlg is set to a value other than 'none'",
			flagAlg: "other",
			want:    false,
		},
		{
			name:    "Scenario 3: FlagAlg is not set",
			flagAlg: "",
			want:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			flagAlg = flag.String("alg", tc.flagAlg, algHelp())

			got := isNone()

			if got != tc.want {
				t.Errorf("isNone() = %v, want %v", got, tc.want)
			}
		})
	}
}
