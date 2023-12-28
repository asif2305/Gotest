package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func addI(a, b int) int {
	return a + b
}
func addF(a, b float64) float64 {
	return a + b
}

// type contraint
type myNumber interface {
	//~int | ~float64 // for type alias
	constraints.Integer | constraints.Float
}

// type alice
type myAlice int

// generic function

func addT[T myNumber](a, b T) T {
	return a + b
}

func main() {
	var n myAlice = 1
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.2, 2.2))
}
