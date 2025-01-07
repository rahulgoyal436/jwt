package main

import (
	"io/ioutil"
	"os"
	"testing"
)







func TestloadData(t *testing.T) {

	tests := []struct {
		name    string
		path    string
		content string
		wantErr bool
	}{
		{
			name:    "Load data from a specified file path",
			path:    "testfile",
			content: "dummy data",
			wantErr: false,
		},
		{
			name:    "Return empty JSON object when path is +",
			path:    "+",
			content: "{}",
			wantErr: false,
		},
		{
			name:    "Return error when no path is specified",
			path:    "",
			content: "",
			wantErr: true,
		},
		{
			name:    "Return error when file does not exist",
			path:    "nonexistentfile",
			content: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.path != "" && tt.path != "+" && tt.path != "-" {
				err := ioutil.WriteFile(tt.path, []byte(tt.content), 0644)
				if err != nil {
					t.Fatal(err)
				}
				defer os.Remove(tt.path)
			}

			got, err := loadData(tt.path)

			if (err != nil) != tt.wantErr {
				t.Errorf("loadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && string(got) != tt.content {
				t.Errorf("loadData() = %v, want %v", string(got), tt.content)
			}
		})
	}
}
