package main

import (
	"fmt"
	"math/rand"
	"time"
)

// producer generates random integers and sends them to the provided channel.
// It generates 5 random numbers and sends them with a delay of 500 milliseconds.
func producer(ch chan<- int) {
	defer close(ch)
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize a new random number generator
	for i := 0; i < 5; i++ {
		num := r.Intn(100)
		fmt.Println("Producing:", num)
		ch <- num
		time.Sleep(time.Millisecond * 500)
	}
}

// consumer reads integers from the provided channel and calculates their squares.
// It prints the squares of received numbers.
func consumer(ch <-chan int) {
	for num := range ch {
		square := num * num
		fmt.Println("Consuming:", square)
	}
}

func main() {
	// Create a channel to communicate between producer and consumer.
	ch := make(chan int)

	// Start the producer in a separate goroutine to generate random numbers.
	go producer(ch)

	// Start the consumer in the main goroutine to consume generated numbers.
	consumer(ch)
}
