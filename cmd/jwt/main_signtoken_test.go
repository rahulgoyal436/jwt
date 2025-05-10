package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"testing"
)

var (
	flagAlg = flag.String("alg", "", "")
	flagClaims = make(ArgList)
	flagKey = flag.String("key", "", "")
	flagSign = flag.String("sign", "", "")
)

type testCase struct {
	name          string
	flagKey       string
	flagAlg       string
	flagClaims    ArgList
	flagSign      string
	expectedError error
}

func TestsignToken(t *testing.T) {
	tests := []testCase{
		{
			name:          "Successful Token Signing",
			flagKey:       "validKey",
			flagAlg:       "RS256",
			flagClaims:    ArgList{"exp": "3600", "iss": "testIssuer"},
			flagSign:      "validTokenData",
			expectedError: nil,
		},
		{
			name:          "Invalid Token Data",
			flagKey:       "validKey",
			flagAlg:       "RS256",
			flagClaims:    ArgList{"exp": "3600", "iss": "testIssuer"},
			flagSign:      "invalidTokenData",
			expectedError: fmt.Errorf("couldn't read token: %w", errors.New("no path specified")),
		},
		{
			name:          "Invalid Key",
			flagKey:       "invalidKey",
			flagAlg:       "RS256",
			flagClaims:    ArgList{"exp": "3600", "iss": "testIssuer"},
			flagSign:      "validTokenData",
			expectedError: fmt.Errorf("couldn't read key: %w", errors.New("no path specified")),
		},
		{
			name:          "Invalid Signing Method",
			flagKey:       "validKey",
			flagAlg:       "invalidAlg",
			flagClaims:    ArgList{"exp": "3600", "iss": "testIssuer"},
			flagSign:      "validTokenData",
			expectedError: fmt.Errorf("couldn't find signing method: %v", "invalidAlg"),
		},
		{
			name:          "Invalid Claims JSON",
			flagKey:       "validKey",
			flagAlg:       "RS256",
			flagClaims:    ArgList{"exp": "invalidExp", "iss": "testIssuer"},
			flagSign:      "validTokenData",
			expectedError: fmt.Errorf("couldn't parse claims JSON: %w", errors.New("invalid character 'i' looking for beginning of value")),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			*flagAlg = test.flagAlg
			*flagKey = test.flagKey
			flagClaims = test.flagClaims
			*flagSign = test.flagSign

			var buf bytes.Buffer
			os.Stdout = &buf

			err := signToken()
			if !errors.Is(err, test.expectedError) {
				t.Logf("Expected error: %v, got: %v", test.expectedError, err)
				t.Fail()
			}

			if test.expectedError == nil {
				tokenString := strings.TrimSpace(buf.String())
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					return []byte(test.flagKey), nil
				})
				if err != nil {
					t.Logf("Error parsing token: %v", err)
					t.Fail()
				} else {
					if !token.Valid {
						t.Log("Token is not valid")
						t.Fail()
					}
				}
			}
		})
	}
}
