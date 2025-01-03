package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"testing"
	"os"
	"io/ioutil"
)


func TestprintJSON(t *testing.T) {
	flag.Parse()

	tests := []struct {
		name      string
		j         interface{}
		compact   bool
		wantError bool
	}{
		{
			name:      "Successful JSON Marshalling with Compact Flag Off",
			j:         map[string]string{"key": "value"},
			compact:   false,
			wantError: false,
		},
		{
			name:      "Successful JSON Marshalling with Compact Flag On",
			j:         map[string]string{"key": "value"},
			compact:   true,
			wantError: false,
		},
		{
			name:      "Unsuccessful JSON Marshalling due to Invalid Input",
			j:         make(chan int),
			compact:   false,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			*flagCompact = tt.compact

			err := printJSON(tt.j)

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = old

			if (err != nil) != tt.wantError {
				t.Errorf("printJSON() error = %v, wantError %v", err, tt.wantError)
				return
			}

			if err == nil {
				var expected []byte
				if tt.compact {
					expected, _ = json.Marshal(tt.j)
				} else {
					expected, _ = json.MarshalIndent(tt.j, "", "    ")
				}

				if tt.compact && bytes.Contains(out, []byte("\n")) {
					t.Errorf("printJSON() = %s, want %s", out, expected)
				} else if !tt.compact && !bytes.Contains(out, []byte("\n")) {
					t.Errorf("printJSON() = %s, want %s", out, expected)
				}
			}
		})
	}
}

