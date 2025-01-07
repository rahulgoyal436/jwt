package main

import (
	"flag"
	"testing"
)







func TestprintJSON(t *testing.T) {

	type testData struct {
		name      string
		input     interface{}
		wantError bool
		compact   bool
	}

	tests := []testData{
		{
			name:      "Successful JSON Marshalling with Compact Flag Off",
			input:     map[string]string{"test": "data"},
			wantError: false,
			compact:   false,
		},
		{
			name:      "Successful JSON Marshalling with Compact Flag On",
			input:     map[string]string{"test": "data"},
			wantError: false,
			compact:   true,
		},
		{
			name:      "Unsuccessful JSON Marshalling due to Invalid Input",
			input:     make(chan int),
			wantError: true,
			compact:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			flagCompact = flag.Bool("compact", tc.compact, "output compact JSON")

			err := printJSON(tc.input)

			if (err != nil) != tc.wantError {
				t.Errorf("printJSON() error = %v, wantError %v", err, tc.wantError)
			}
		})
	}
}

