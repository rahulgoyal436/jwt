package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestloadData(t *testing.T) {
	// Test Scenarios
	tests := []struct {
		name        string
		filePath    string
		expected    []byte
		expectError bool
	}{
		{
			"Load data from a specified file path",
			"testdata/testfile.txt",
			[]byte("test data"),
			false,
		},
		{
			"Return empty JSON object when path is '+'",
			"+",
			[]byte("{}"),
			false,
		},
		{
			"Return error when no path is specified",
			"",
			nil,
			true,
		},
		{
			"Return error when file does not exist",
			"testdata/non_existent.txt",
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			data, err := loadData(tt.filePath)

			// Assert
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if string(data) != string(tt.expected) {
					t.Errorf("Expected %s but got %s", tt.expected, data)
				}
			}
		})
	}
}

// Test setup
func TestMain(m *testing.M) {
	// Arrange
	err := ioutil.WriteFile("testdata/testfile.txt", []byte("test data"), 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
		os.Exit(1)
	}
	// Run the tests
	code := m.Run()

	// Clean up
	err = os.Remove("testdata/testfile.txt")
	if err != nil {
		fmt.Printf("Unable to remove file: %v", err)
		os.Exit(1)
	}

	os.Exit(code)
}
