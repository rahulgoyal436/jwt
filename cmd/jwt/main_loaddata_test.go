package main

import (
	"io"
	"os"
	"testing"
	"io/ioutil"
)


func TestloadData(t *testing.T) {

	type test struct {
		name string
		path string
		want string
		err  bool
	}

	tests := []test{
		{
			name: "Load data from a specified file path",
			path: "test.txt",
			want: "This is a test file",
			err:  false,
		},
		{
			name: "Load data from standard input",
			path: "-",
			want: "Standard input data",
			err:  false,
		},
		{
			name: "Return empty JSON object when path is '+'",
			path: "+",
			want: "{}",
			err:  false,
		},
		{
			name: "Return error when no path is specified",
			path: "",
			want: "",
			err:  true,
		},
		{
			name: "Return error when file does not exist",
			path: "non-existent.txt",
			want: "",
			err:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.path == "test.txt" {
				ioutil.WriteFile(tt.path, []byte(tt.want), 0644)
			}
			if tt.path == "-" {
				oldStdin := os.Stdin
				defer func() { os.Stdin = oldStdin }()
				r, w, _ := os.Pipe()
				os.Stdin = r
				io.WriteString(w, tt.want)
				w.Close()
			}

			got, err := loadData(tt.path)

			if (err != nil) != tt.err {
				t.Errorf("loadData() error = %v, wantErr %v", err, tt.err)
				return
			}
			if string(got) != tt.want {
				t.Errorf("loadData() = %v, want %v", string(got), tt.want)
			}

			if tt.path == "test.txt" {
				os.Remove(tt.path)
			}
		})
	}
}
