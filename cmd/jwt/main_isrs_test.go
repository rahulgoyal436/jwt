package main

import (
	"flag"
	"testing"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestisRs(t *testing.T) {

	testCases := []struct {
		name    string
		flagAlg string
		want    bool
	}{
		{"Algorithm Prefix is 'RS'", "RS256", true},
		{"Algorithm Prefix is 'PS'", "PS384", true},
		{"Algorithm Prefix is neither 'RS' nor 'PS'", "HS256", false},
		{"Algorithm Prefix is empty", "", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {

			*flagAlg = tt.flagAlg

			got := isRs()

			if got != tt.want {
				t.Errorf("isRs() = %v, want %v", got, tt.want)
			} else {
				t.Logf("isRs() passed for case: %v", tt.name)
			}
		})
	}
}

