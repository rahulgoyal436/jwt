package main

import (
	"testing"
	"github.com/golang-jwt/jwt/v5/cmd/jwt"
	"os"
	"bytes"
)

// Teststart tests the start function
func Teststart(t *testing.T) {
	tests := []struct {
		name      string
		flagSign  string
		flagVerify string
		flagShow  string
		wantErr   bool
	}{
		{
			"Test for flagSign with non-empty value",
			"signToken",
			"",
			"",
			false,
		},
		{
			"Test for flagVerify with non-empty value",
			"",
			"verifyToken",
			"",
			false,
		},
		{
			"Test for flagShow with non-empty value",
			"",
			"",
			"showToken",
			false,
		},
		{
			"Test for all flags empty",
			"",
			"",
			"",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*jwt.flagSign = tt.flagSign
			*jwt.flagVerify = tt.flagVerify
			*jwt.flagShow = tt.flagShow

			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w
			
			err := start()

			// Stop capturing stdout
			w.Close()
			os.Stdout = oldStdout

			out, _ := io.ReadAll(r)

			if (err != nil) != tt.wantErr {
				t.Errorf("start() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !strings.Contains(string(out), "none of the required flags are present. What do you want me to do?") {
					t.Errorf("start() expected error message not found in output")
				}
			} else {
				switch {
				case tt.flagSign != "":
					if !strings.Contains(string(out), "signToken") {
						t.Errorf("start() expected signToken function to be called")
					}
				case tt.flagVerify != "":
					if !strings.Contains(string(out), "verifyToken") {
						t.Errorf("start() expected verifyToken function to be called")
					}
				case tt.flagShow != "":
					if !strings.Contains(string(out), "showToken") {
						t.Errorf("start() expected showToken function to be called")
					}
				}
			}
		})
	}
}
