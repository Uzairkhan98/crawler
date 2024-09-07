package main

import (
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		expectedError string
		expectedOut   string
	}{
		{
			name:          "No arguments provided",
			args:          []string{},
			expectedError: "no website provided",
		},
		{
			name:          "Too many arguments provided",
			args:          []string{"https://example.com", "extra-arg"},
			expectedError: "too many arguments provided",
		},
		{
			name:        "Valid single argument",
			args:        []string{"https://example.com"},
			expectedOut: "starting crawl of: https://example.com",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Mocking os.Args
			os.Args = append([]string{"crawler"}, tc.args...)

			// Capture stdout or stderr (you can use libraries like "github.com/stretchr/testify" to capture output)
			var output string
			defer func() {
				if r := recover(); r != nil {
					output = r.(string)
				}
			}()

			// Running the main function to test behavior
			main()

			// Checking for expected error
			if tc.expectedError != "" && output != tc.expectedError {
				t.Errorf("Expected error: %v, got: %v", tc.expectedError, output)
			}

			// Checking for expected output
			if tc.expectedOut != "" && output != tc.expectedOut {
				t.Errorf("Expected output: %v, got: %v", tc.expectedOut, output)
			}
		})
	}
}
