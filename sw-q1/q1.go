package main

import (
	"fmt"
	"sort"
	"strings"
)

// countA: Counts how many times the letter "a" appears in a word (case insensitive)
func countA(word string) int {
	return strings.Count(strings.ToLower(word), "a")
}

// byAThenLength: Helps sort words by "a" count, then by length, and finally by original order.
type byAThenLength struct {
	words []string
	index []int // Keeps the original index of words
}

func (s byAThenLength) Len() int {
	return len(s.words)
}

func (s byAThenLength) Swap(i, j int) {
	s.words[i], s.words[j] = s.words[j], s.words[i]
	s.index[i], s.index[j] = s.index[j], s.index[i]
}

func (s byAThenLength) Less(i, j int) bool {
	// Compare by "a" count
	countA_i := countA(s.words[i])
	countA_j := countA(s.words[j])

	if countA_i != countA_j {
		return countA_i > countA_j // Words with more "a"s come first
	}

	// If "a" counts are the same, compare by word length
	if len(s.words[i]) != len(s.words[j]) {
		return len(s.words[i]) > len(s.words[j]) // Longer words come first
	}

	// If tied, keep the original order
	return s.index[i] < s.index[j]
}

// sortByA: Sorts words by "a" count, then by length, and keeps the original order for ties.
func sortByA(words []string) []string {
	// Save the original indices of words
	index := make([]int, len(words))
	for i := range words {
		index[i] = i
	}

	// Sort the words using custom logic
	sorter := byAThenLength{words: words, index: index}
	sort.Stable(sorter)

	return sorter.words
}

// equalSlices: Checks if two slices of strings are the same.
func equalSlices(a, b []string) bool {
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

// testSortByA: Tests the sortByA function with various test cases.
func testSortByA() {
	testCases := []struct {
		description string   // What the test case is testing
		input       []string // Input for the test
		expected    []string // Expected output
	}{
		{
			description: "Sort by 'a' count, then length, simple example",
			input:       []string{"apple", "banana", "cherry", "date", "fig"},
			expected:    []string{"banana", "apple", "date", "cherry", "fig"},
		},
		{
			description: "Words with only 'a', sorted by count and length",
			input:       []string{"aa", "a", "aaa", "aaaa", "b"},
			expected:    []string{"aaaa", "aaa", "aa", "a", "b"},
		},
		{
			description: "Mixed words with various 'a' counts and lengths",
			input:       []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"},
			expected:    []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"},
		},
		{
			description: "Empty list of words",
			input:       []string{},
			expected:    []string{},
		},
		{
			description: "All identical words",
			input:       []string{"aa", "aa", "aa"},
			expected:    []string{"aa", "aa", "aa"},
		},
		{
			description: "Words without 'a', original order maintained",
			input:       []string{"hello", "world", "test"},
			expected:    []string{"hello", "world", "test"},
		},
		{
			description: "Words with uppercase and lowercase 'a'",
			input:       []string{"Apple", "banana", "Cherry", "date", "Fig"},
			expected:    []string{"banana", "Apple", "date", "Cherry", "Fig"},
		},
	}

	for i, tc := range testCases {
		fmt.Printf("Test %d: %s\n", i+1, tc.description)
		result := sortByA(tc.input)
		fmt.Printf("Input: %v\n", tc.input)
		fmt.Printf("Expected: %v\n", tc.expected)
		fmt.Printf("Result: %v\n", result)
		if !equalSlices(result, tc.expected) {
			fmt.Printf("❌ Test %d failed\n\n", i+1)
		} else {
			fmt.Printf("✅ Test %d passed\n\n", i+1)
		}
	}
}

func main() {
	// Example input
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	// Sort words and display the result
	fmt.Println("Sorted Words:", sortByA(words))

	// Run tests
	fmt.Println("\nRunning Tests...")
	testSortByA()
}
