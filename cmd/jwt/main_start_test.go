package main

import (
	"errors"
	"flag"
	"fmt"
	"testing"
)






var flagSign = flag.String("sign", "", "")flagVerify = flag.String("verify", "", "")flagShow = flag.String("show", "", "")signTokenCalled boolverifyTokenCalled boolshowTokenCalled bool
var flagSign = flag.String("sign", "", "")flagVerify = flag.String("verify", "", "")flagShow = flag.String("show", "", "")signTokenCalled bool
var flagSign = flag.String("sign", "", "")flagVerify = flag.String("verify", "", "")flagShow = flag.String("show", "", "")signTokenCalled boolverifyTokenCalled bool



func Teststart(t *testing.T) {
	tt := []struct {
		name      string
		setup     func()
		wantError bool
	}{
		{
			name: "Test for flagSign with a non-empty value",
			setup: func() {
				*flagSign = "test"
			},
			wantError: false,
		},
		{
			name: "Test for flagVerify with a non-empty value",
			setup: func() {
				*flagVerify = "test"
			},
			wantError: false,
		},
		{
			name: "Test for flagShow with a non-empty value",
			setup: func() {
				*flagShow = "test"
			},
			wantError: false,
		},
		{
			name: "Test for all flags empty",
			setup: func() {

			},
			wantError: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			resetFlags()
			tc.setup()
			err := start()
			if (err != nil) != tc.wantError {
				t.Errorf("start() error = %v, wantError %v", err, tc.wantError)
			}

			switch {
			case *flagSign != "":
				if !signTokenCalled {
					t.Error("signToken was not called")
				}
			case *flagVerify != "":
				if !verifyTokenCalled {
					t.Error("verifyToken was not called")
				}
			case *flagShow != "":
				if !showTokenCalled {
					t.Error("showToken was not called")
				}
			default:
				if err == nil || !errors.Is(err, fmt.Errorf("none of the required flags are present. What do you want me to do?")) {
					t.Error("expected error was not returned")
				}
			}
		})
	}
}
func resetFlags() {
	*flagSign = ""
	*flagVerify = ""
	*flagShow = ""
	signTokenCalled = false
	verifyTokenCalled = false
	showTokenCalled = false
}
