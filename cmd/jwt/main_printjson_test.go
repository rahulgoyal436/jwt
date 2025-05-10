package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)



type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestprintJSON(t *testing.T) {

	testCases := []struct {
		name      string
		input     interface{}
		compact   bool
		expectErr bool
	}{
		{
			name: "Successful JSON Marshalling with Compact Flag Off",
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			compact:   false,
			expectErr: false,
		},
		{
			name: "Successful JSON Marshalling with Compact Flag On",
			input: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			compact:   true,
			expectErr: false,
		},
		{
			name:      "Unsuccessful JSON Marshalling due to Invalid Input",
			input:     make(chan int),
			compact:   false,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			*flagCompact = tc.compact
			var buf bytes.Buffer
			print := fmt.Println
			fmt.Println = func(args ...interface{}) (n int, err error) {
				return fmt.Fprintln(&buf, args...)
			}
			defer func() {
				fmt.Println = print
			}()

			err := printJSON(tc.input)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected an error but didn't get one")
				}
			} else {
				if err != nil {
					t.Errorf("Didn't expect an error but got one: %v", err)
				} else {

					var jsonData interface{}
					if err := json.Unmarshal(buf.Bytes(), &jsonData); err != nil {
						t.Errorf("Output is not valid JSON: %v", err)
					}
				}
			}
		})
	}
}

