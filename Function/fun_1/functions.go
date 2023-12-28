package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...) // unfurling slice
	fmt.Println("The sum is ", x)
	defer foo()
	barCopy("Todo")

	s := aloha("Asif")
	fmt.Println(s)

	s1, n := dogYears("Shepard", 40)
	fmt.Println(s1, n)

}

// variadic parameter
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

// without parameter, receiver and return type
func foo() {
	fmt.Println("This is a function with no return")
}

// without receiver and return type
func barCopy(s string) {
	fmt.Println("My name is ", s)
}

// without receiver type
func aloha(s string) string {
	return fmt.Sprint("My name is Mr. ", s)
}
func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, "is this dog in dig years "), age
}
