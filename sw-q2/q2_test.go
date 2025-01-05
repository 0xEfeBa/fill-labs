package main

import (
	"fmt"
	"testing"
)

// Test function to check the MainRecursive function with various inputs.
func TestMainRecursive2(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5", "Input: 5\nOutput: 2 4 \n"},
		{"15", "Input: 15\nOutput: 2 4 9 \n"},
		{"16", "Input: 16\nOutput: 2 4 9 16 \n"},
		{"-1", "Error: Invalid input! Please enter a number greater than or equal to 2.\n"},
		{"abc", "Error: Invalid input! Please enter a number greater than or equal to 2.\n"},
		{"c", "Exiting the program...\n"},
	}

	for _, test := range tests {
		// Call MainRecursive with the test input
		actualOutput := MainRecursive(test.input)

		// Compare the actual output with the expected output
		if actualOutput == test.expected {
			fmt.Printf("✓ Test passed for input: %q\n", test.input)
		} else {
			fmt.Printf("✗ Test failed for input: %q\n", test.input)
			fmt.Printf("  Expected: %q\n", test.expected)
			fmt.Printf("  Got: %q\n", actualOutput)
		}
	}
}
