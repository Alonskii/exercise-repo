package main

import (
	"fmt"
	"sync"
)

// Function to calculate the sum of elements in a sub-array
func sum(arr []int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	result := 0
	for _, num := range arr {
		result += num
	}
	ch <- result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Your large array
	numGoroutines := 4                          // Number of goroutines to use
	ch := make(chan int, numGoroutines)
	var wg sync.WaitGroup

	// Calculate the length of sub-arrays
	subArrayLen := len(arr) / numGoroutines
	wg.Add(numGoroutines)

	// Launch goroutines to calculate sum of sub-arrays
	for i := 0; i < numGoroutines; i++ {
		start := i * subArrayLen
		end := (i + 1) * subArrayLen
		if i == numGoroutines-1 { // If it's the last goroutine, include the remaining elements
			end = len(arr)
		}
		go sum(arr[start:end], &wg, ch)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Receive results from goroutines and calculate total sum
	totalSum := 0
	for partialSum := range ch {
		totalSum += partialSum
	}

	fmt.Println("Total sum:", totalSum)
}
