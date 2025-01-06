package main

import (
	"io/ioutil"
	"testing"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestloadData(t *testing.T) {

	testCases := []struct {
		name     string
		path     string
		setup    func() error
		expected string
		hasError bool
	}{
		{
			name: "Load data from a specified file path",
			path: "test.txt",
			setup: func() error {
				return ioutil.WriteFile("test.txt", []byte("Hello, World!"), 0644)
			},
			expected: "Hello, World!",
			hasError: false,
		},
		{
			name:     "Return error when no path is specified",
			path:     "",
			expected: "",
			hasError: true,
		},
		{
			name:     "Return empty JSON object when path is '+'",
			path:     "+",
			expected: "{}",
			hasError: false,
		},
		{
			name:     "Return error when file does not exist",
			path:     "non_existent.txt",
			expected: "",
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.setup != nil {
				err := tc.setup()
				if err != nil {
					t.Fatalf("Failed to setup test: %v", err)
				}
			}

			result, err := loadData(tc.path)

			if (err != nil) != tc.hasError {
				t.Fatalf("loadData() error = %v, wantErr %v", err, tc.hasError)
			}
			if string(result) != tc.expected {
				t.Errorf("loadData() = %v, want %v", string(result), tc.expected)
			}
		})
	}
}
