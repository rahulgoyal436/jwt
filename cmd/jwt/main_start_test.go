package main

import (
	"flag"
	"testing"
	"os"
	"io"
	"strings"
)


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func Teststart(t *testing.T) {
	tests := []struct {
		name       string
		flagSign   string
		flagVerify string
		flagShow   string
		wantErr    bool
	}{
		{"Test for flagSign with a non-empty value", "sign", "", "", false},
		{"Test for flagVerify with a non-empty value", "", "verify", "", false},
		{"Test for flagShow with a non-empty value", "", "", "show", false},
		{"Test for all flags empty", "", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			*flagSign = tt.flagSign
			*flagVerify = tt.flagVerify
			*flagShow = tt.flagShow

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := start()

			w.Close()
			out, _ := io.ReadAll(r)
			os.Stdout = old

			if (err != nil) != tt.wantErr {
				t.Errorf("start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {

				if !strings.Contains(string(out), "none of the required flags are present") {
					t.Errorf("Expected error message not found in output")
				}
			} else {

				var expectedOutput string

				switch {
				case *flagSign != "":
					expectedOutput = "signToken"
				case *flagVerify != "":
					expectedOutput = "verifyToken"
				case *flagShow != "":
					expectedOutput = "showToken"
				}

				if !strings.Contains(string(out), expectedOutput) {
					t.Errorf("Expected output of %v not found in stdout", expectedOutput)
				}
			}
		})
	}
}
