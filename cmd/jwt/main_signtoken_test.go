package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
)







type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestsignToken(t *testing.T) {
	tests := []struct {
		name    string
		alg     string
		claims  string
		key     string
		wantErr bool
	}{
		{
			name:    "Successful Token Signing",
			alg:     "HS256",
			claims:  `{"sub":"1234567890","name":"John Doe","iat":1516239022}`,
			key:     "secret",
			wantErr: false,
		},
		{
			name:    "Invalid Token Data",
			alg:     "HS256",
			claims:  `{"sub":"1234567890","name":"John Doe","iat":"invalid"}`,
			key:     "secret",
			wantErr: true,
		},
		{
			name:    "Invalid Key",
			alg:     "HS256",
			claims:  `{"sub":"1234567890","name":"John Doe","iat":1516239022}`,
			key:     "",
			wantErr: true,
		},
		{
			name:    "Invalid Signing Method",
			alg:     "INVALID",
			claims:  `{"sub":"1234567890","name":"John Doe","iat":1516239022}`,
			key:     "secret",
			wantErr: true,
		},
		{
			name:    "Invalid Claims JSON",
			alg:     "HS256",
			claims:  `{"sub":"1234567890","name":"John Doe","iat":invalid}`,
			key:     "secret",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*flagAlg = tt.alg
			*flagSign = tt.claims
			*flagKey = tt.key

			err := signToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("signToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				t.Logf("signToken() = success")
			}
		})
	}
}






