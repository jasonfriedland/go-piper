package piper

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_Read(t *testing.T) {
	// Tests
	tests := []struct {
		name    string
		stdIn   string
		want    string
		wantErr bool
	}{
		// Test cases
		{"Pass stdin, simple case", "x\n😀\nx", "x\n😀\nx", false},
		{"Pass stdin, trailing newline", "x\n\nx\n", "x\n\nx", false},
		{"Pass long stdin, err true", strings.Repeat("x", 10e5), "", true},
		{"No stdin, empty case", "", "", false},
		{"No stdin, multiple newlines", "\n\n\n\n", "", false},
	}

	for _, tt := range tests {
		// Back up STDIN
		oldStdin := os.Stdin
		if tt.stdIn != "" {
			tmp := makeTempFile(tt.stdIn)
			defer os.Remove(tmp.Name())
			os.Stdin = tmp
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
		// Restore STDIN
		os.Stdin = oldStdin
	}
}

func makeTempFile(input string) *os.File {
	content := []byte(input)
	tmpFile, err := ioutil.TempFile("", "main_test")
	if err != nil {
		return nil
	}
	if _, err := tmpFile.Write(content); err != nil {
		return nil
	}
	if _, err := tmpFile.Seek(0, 0); err != nil {
		return nil
	}
	return tmpFile
}
