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

func multi(a int, b int) int {
	fmt.Printf("multi : %d from %d ", a, b)
	return a * b
}

func multiply(a int, b int) int {
	fmt.Printf("Multiplying : %d and %d ", a, b)

	return a * b
}

func divide(a int, b int) int {
	fmt.Printf("Dividing : %d by %d", a, b)
	return a / b
}
