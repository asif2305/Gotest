package main

import "fmt"

func main() {

	x := 42
	y := &x
	fmt.Println(*y)

	fmt.Println(x)
	fmt.Println(&x)
	*y = 10

	fmt.Println(x)
	fmt.Println(&x)
}
