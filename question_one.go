package main

import "fmt"

// swap swaps the values of two integers.
func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	// Declare two integer variables.
	a := 5
	b := 10

	// Print the values of a and b before swapping.
	fmt.Println("Before swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// Call the swap function, passing the addresses of a and b.
	swap(&a, &b)

	// Print the values of a and b after swapping.
	fmt.Println("After swapping:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
