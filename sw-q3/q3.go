package main

import "fmt"

// findMostFrequent finds the most frequently occurring element in a slice.
// In case of ties, it returns the first encountered element.
func findMostFrequent(elements []string) string {
	// Create a map to store the frequency of each element
	frequency := make(map[string]int)

	// Iterate through the slice and update the frequency map
	for _, element := range elements {
		frequency[element]++
	}

	// Find the most frequent element
	var mostFrequentElement string
	highestFrequency := 0

	// Iterate through the original slice to maintain the first encountered order
	for _, element := range elements {
		if frequency[element] > highestFrequency {
			mostFrequentElement = element
			highestFrequency = frequency[element]
		}
	}

	return mostFrequentElement
}

func main() {
	// Test scenarios
	testScenarios := []struct {
		data     []string
		expected string
		reason   string // Explanation of what the test is checking
	}{
		{
			data:     []string{"apple", "pie", "apple", "red", "red", "red"},
			expected: "red",
			reason:   "Test finding the most frequent element in a slice with multiple occurrences.",
		},
		{
			data:     []string{"a", "b", "c", "a", "b", "a"},
			expected: "a",
			reason:   "Test finding the most frequent element in a slice with repeated elements.",
		},
		{
			data:     []string{"x", "y", "z", "z", "y"},
			expected: "y",
			reason:   "Test finding the most frequent element in a slice with ties (first encountered is returned).",
		},
		{
			data:     []string{},
			expected: "",
			reason:   "Test finding the most frequent element in an empty slice.",
		},
	}

	// Run test scenarios
	for index, scenario := range testScenarios {
		outcome := findMostFrequent(scenario.data)
		fmt.Printf("Test Scenario %d\n", index+1)
		fmt.Printf("Reason: %s\n", scenario.reason)
		fmt.Printf("Input: %v\n", scenario.data)
		fmt.Printf("Expected: %v\n", scenario.expected)
		fmt.Printf("Result: %v\n", outcome)

		// Check if the outcome matches the expected result
		if outcome != scenario.expected {
			fmt.Printf("Status: ❌ Failed\n\n")
		} else {
			fmt.Printf("Status: ✅ Passed\n\n")
		}
	}
}
