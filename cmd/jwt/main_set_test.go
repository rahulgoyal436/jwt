package main

import (
	"strings"
	"testing"
)


func TestArgListSet(t *testing.T) {

	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{
			name:    "Valid Argument Test",
			arg:     "key=value",
			wantErr: false,
		},
		{
			name:    "Invalid Argument Test",
			arg:     "invalid",
			wantErr: true,
		},
		{
			name:    "Empty Argument Test",
			arg:     "",
			wantErr: true,
		},
		{
			name:    "Argument with Extra Equals Signs Test",
			arg:     "key=value=value",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			al := make(ArgList)

			err := al.Set(tt.arg)

			if (err != nil) != tt.wantErr {
				t.Errorf("ArgList.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				parts := strings.SplitN(tt.arg, "=", 2)
				if al[parts[0]] != parts[1] {
					t.Errorf("ArgList.Set() = %v, want %v", al[parts[0]], parts[1])
				}
			}
		})
	}
}
