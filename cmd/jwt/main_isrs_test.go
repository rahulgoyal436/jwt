package main

import (
	"strings"
	"testing"
	"flag"
)

func TestisRs(t *testing.T) {

	testCases := []struct {
		name   string
		prefix string
		want   bool
	}{
		{
			name:   "Algorithm Prefix is RS",
			prefix: "RS",
			want:   true,
		},
		{
			name:   "Algorithm Prefix is PS",
			prefix: "PS",
			want:   true,
		},
		{
			name:   "Algorithm Prefix is neither RS nor PS",
			prefix: "ES",
			want:   false,
		},
		{
			name:   "Algorithm Prefix is empty",
			prefix: "",
			want:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			*flagAlg = tc.prefix

			got := isRs()

			if got != tc.want {
				t.Errorf("isRs() = %v, want %v", got, tc.want)
			} else {
				t.Logf("Successful test run for scenario: %s", tc.name)
			}
		})
	}
}
