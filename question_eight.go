package main

import (
	"testing"
)

// TestSumEvenNumbers tests the sumEvenNumbers function.
func TestSumEvenNumbers(t *testing.T) {
	// Define test cases with different input slices and expected results.
	tests := []struct {
		name     string // Name of the test case for identification.
		numbers  []int  // Input slice of numbers.
		expected int    // Expected sum of even numbers in the input slice.
	}{
		{"All Even Numbers", []int{2, 4, 6, 8}, 20},
		{"All Odd Numbers", []int{1, 3, 5, 7}, 0},
		{"Mixed Numbers", []int{1, 2, 3, 4, 5, 6}, 12},
		{"Empty Slice", []int{}, 0},
		{"Single Element (Even)", []int{10}, 10},
		{"Single Element (Odd)", []int{7}, 0},
		{"Negative Numbers", []int{-2, -4, -6, -8}, -20},
	}

	// Iterate over each test case.
	for _, test := range tests {
		// Run each test case.
		t.Run(test.name, func(t *testing.T) {
			// Call the sumEvenNumbers function with the test case input.
			actual := sumEvenNumbers(test.numbers)
			// Compare the actual result with the expected result.
			if actual != test.expected {
				// If the actual result doesn't match the expected result, report the failure.
				t.Errorf("Test case %q failed: expected %d, got %d", test.name, test.expected, actual)
			}
		})
	}
}

// TestMain is a test function for the main function.
func TestMain(t *testing.T) {
	// Main function testing strategy: Instead of capturing stdout and checking its contents,
	// we ensure that the main function runs without error.
	main()
}
