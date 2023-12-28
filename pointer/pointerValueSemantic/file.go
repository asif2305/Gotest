package main

import "fmt"

// value semantics
func addOne(v int) int {
	return v + 1
}

// poiner semantics
func addOnePointer(v *int) {
	*v += 1
}

func main() {
	// value semantics
	a := 1
	fmt.Println(a)         // 1
	fmt.Println(addOne(1)) // 2
	fmt.Println(a)         // 1
	// pointer semantics
	b := 1
	fmt.Println("Value semantics")
	fmt.Println(b) // 1
	addOnePointer(&b)
	fmt.Println(b) //2

}
