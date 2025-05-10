package main

import (
	"testing"
	"strings"
)

var testTable = []struct {
	name		string
	input		string
	expected	bool
}{
	{
		name:		"Scenario 1: Algorithm Prefix is RS",
		input:		"RS256",
		expected:	true,
	},
	{
		name:		"Scenario 2: Algorithm Prefix is PS",
		input:		"PS384",
		expected:	true,
	},
	{
		name:		"Scenario 3: Algorithm Prefix is neither RS nor PS",
		input:		"HS256",
		expected:	false,
	},
	{
		name:		"Scenario 4: Algorithm Prefix is empty",
		input:		"",
		expected:	false,
	},
}func TestisRs(t *testing.T) {
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {

			*flagAlg = tt.input

			result := isRs()

			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
