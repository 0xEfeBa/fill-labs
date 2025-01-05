package main

import (
	"fmt"
	"math"
	"strconv"
)

// This program takes user input, validates it, and prints a sequence of perfect squares up to the input number.
func getValidInput(input string) (int, string) {
	if input == "c" {
		return 0, "Exiting the program...\n"
	}

	number, err := strconv.ParseFloat(input, 64)
	if err != nil || number < 2 {
		return -1, "Error: Invalid input! Please enter a number greater than or equal to 2.\n"
	}

	roundedNumber := int(math.Round(number))
	return roundedNumber, ""
}

// Recursively generates and returns perfect squares up to the given number n, starting from i.
func generateOutputRecursive(n, i int) string {
	if i*i > n {
		return ""
	}
	return fmt.Sprintf("%d ", i*i) + generateOutputRecursive(n, i+1)
}

// Handles the main logic of the program, including input validation and output generation.
func MainRecursive(input string) string {
	number, message := getValidInput(input)
	if message != "" {
		return message
	}
	if number == -1 {
		return "Error: Invalid input! Please enter a number greater than or equal to 2.\n"
	}
	output := fmt.Sprintf("Input: %d\n", number)
	output += "Output: 2 " + generateOutputRecursive(number, 2) + "\n"
	return output
}

func main() {
	// Run the program interactively
	for {
		var input string
		fmt.Print("Enter a number greater than or equal to 2 (type 'c' to exit): ")
		fmt.Scan(&input)

		// Call MainRecursive and print the output
		output := MainRecursive(input)
		fmt.Print(output)

		// Exit the program if the input is 'c'
		if input == "c" {
			break
		}
	}
}
