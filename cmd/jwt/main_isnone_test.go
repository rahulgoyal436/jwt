package main

import (
	"flag"
	"testing"
)

func TestisNone(t *testing.T) {

	tests := []struct {
		name string
		flag string
		want bool
	}{
		{
			name: "Scenario 1: FlagAlg is set to 'none'",
			flag: "none",
			want: true,
		},
		{
			name: "Scenario 2: FlagAlg is set to a value other than 'none'",
			flag: "something",
			want: false,
		},
		{
			name: "Scenario 3: FlagAlg is not set",
			flag: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			flagAlg = flag.String("alg", tt.flag, algHelp())

			got := isNone()

			if got != tt.want {
				t.Errorf("isNone() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Success: %s", tt.name)
			}
		})
	}
}
