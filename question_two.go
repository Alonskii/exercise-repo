package main

import (
	"fmt"
)

// sumEvenNumbers calculates the sum of even numbers in a slice of integers.
// It iterates through the slice and adds each even number to the sum.
func sumEvenNumbers(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number%2 == 0 {
			sum += number
		}
	}
	return sum
}

func main() {
	// Sample slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6}
	// Calculate the sum of even numbers
	evenSum := sumEvenNumbers(numbers)
	// Print the result
	fmt.Println("Sum of even numbers:", evenSum)
}
