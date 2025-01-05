package main

import (
	"fmt"
)

// recursiveFunction finds perfect squares up to n (excluding 1).
func recursiveFunction(n, current int) []int {
	// Base case: if current^2 is greater than n, terminate the function
	if current*current > n {
		return []int{}
	}

	// Collect the square if it's not 1
	result := []int{}
	if current != 1 {
		result = append(result, current*current)
	}

	// Recursively call the function and append the results
	return append(result, recursiveFunction(n, current+1)...)
}

// testRecursiveFunction tests the recursiveFunction.
func testRecursiveFunction() {
	testCases := []struct {
		input    int
		expected []int
		reason   string // Explanation of why the test is performed
	}{
		{
			input:    9,
			expected: []int{4, 9},
			reason:   "Test finding perfect squares up to 9 (excluding 1).",
		},
		{
			input:    16,
			expected: []int{4, 9, 16},
			reason:   "Test finding perfect squares up to 16 (excluding 1).",
		},
		{
			input:    1,
			expected: []int{},
			reason:   "Test finding perfect squares up to 1 (no output expected).",
		},
		{
			input:    0,
			expected: []int{},
			reason:   "Test finding perfect squares up to 0 (no output expected).",
		},
		{
			input:    25,
			expected: []int{4, 9, 16, 25},
			reason:   "Test finding perfect squares up to 25 (excluding 1).",
		},
	}

	for i, tc := range testCases {
		fmt.Printf("Test %d\n", i+1)
		fmt.Printf("Reason: %s\n", tc.reason)
		fmt.Printf("Input: %d\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)

		// Call the recursive function and get the result
		result := recursiveFunction(tc.input, 1)
		fmt.Printf("Result: %v\n", result)

		// Check if the result matches the expected output
		if equalSlices(result, tc.expected) {
			fmt.Printf("✅ Test %d passed\n\n", i+1)
		} else {
			fmt.Printf("❌ Test %d failed\n\n", i+1)
		}
	}
}

// equalSlices checks if two slices of integers are the same.
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	// Example input
	n := 9

	// Call the recursive function
	fmt.Println("Running recursiveFunction with input:", n)
	fmt.Println("Result:", recursiveFunction(n, 1))

	// Run the tests
	fmt.Println("\nRunning Tests...")
	testRecursiveFunction()
}
