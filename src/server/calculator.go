package server

import (
	"fmt"
)

func add(a int, b int) int {

	fmt.Printf("Adding : %d and %d", a, b)
	return a + b
}
