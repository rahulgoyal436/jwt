```go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

// TestshowToken is a unit test for showToken function
func TestshowToken(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name        string
		inputToken  string
		inputDebug  bool
		expectedErr string
	}{
		{
			name:        "Successful token reading and printing",
			inputToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			inputDebug:  false,
			expectedErr: "",
		},
		{
			name:        "Invalid token input",
			inputToken:  "invalid-token",
			inputDebug:  false,
			expectedErr: "malformed token",
		},
		{
			name:        "Token reading error",
			inputToken:  "",
			inputDebug:  false,
			expectedErr: "