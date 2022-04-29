package server

import (
	"fmt"
)

func add(a int, b int) int {

	fmt.Printf("Adding : %d and %d", a, b)
	return a + b
}

func subtract(a int, b int) int {
	fmt.Printf("Subtracting : %d from %d ", a, b)
	return a - b
}

func multiply(a int, b int) int {
	fmt.Printf("Multiplying : %d and %d ", a, b)
	return a * b
}
