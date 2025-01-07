package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

// This is the function under test
func showToken(flagShow *string, flagDebug *bool, flagCompact *bool) error {

	tokData, err := loadData(*flagShow)
	if err != nil {
		return fmt.Errorf("couldn't read token: %w", err)
	}

	tokData = regexp.MustCompile(`\s*`).ReplaceAll(tokData, []byte{})
	if *flagDebug {
		fmt.Fprintf(os.Stderr, "Token len: %v bytes\n", len(tokData))
	}

	token, _, err := jwt.NewParser().ParseUnverified(string(tokData), make(jwt.MapClaims))
	if err != nil {
		return fmt.Errorf("malformed token: %w", err)
	}

	fmt.Println("Header:")
	if err := printJSON(token.Header, flagCompact); err != nil {
		return fmt.Errorf("failed to output header: %w", err)
	}

	fmt.Println("Claims:")
	if err := printJSON(token.Claims, flagCompact); err != nil {
		return fmt.Errorf("failed to output claims: %w", err)
	}

	return nil
}

func TestshowToken(t *testing.T) {
	// define test cases
	tests := []struct {
		name        string
		flagShow    string
		flagDebug   bool
		flagCompact bool
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "Successful token reading and printing",
			flagShow:    "validTokenFile", // TODO: replace with actual valid token file path
			flagDebug:   false,
			flagCompact: false,
			wantErr:     false,
			errMsg:      "",
		},
		{
			name:        "Invalid token input",
			flagShow:    "invalidTokenFile", // TODO: replace with actual invalid token file path
			flagDebug:   false,
			flagCompact: false,
			wantErr:     true,
			errMsg:      "couldn't read token: open invalidTokenFile: no such file or directory",
		},
	}
	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := showToken(&tt.flagShow, &tt.flagDebug, &tt.flagCompact)
			if (err != nil) != tt.wantErr {
				t.Errorf("showToken() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("showToken() error message = %v, wantErrMsg %v", err.Error(), tt.errMsg)
			}
		})
	}
}
