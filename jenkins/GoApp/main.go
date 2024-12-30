package main

import (
	"fmt"
	"pipe/ops"
)

func main() {
	// fmt.Println("hello")
	// Hardcoded example values
	a, b := 10, 5

	// // Calling Add and Subtract functions from operations.go
	fmt.Printf("Adding %d and %d gives: %d\n", a, b, ops.Add(a, b))
	fmt.Printf("Subtracting %d from %d gives: %d\n", b, a, ops.Subtract(a, b))
}
