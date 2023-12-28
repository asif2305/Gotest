package main

import "fmt"

// dereferencing
func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}
func mapDelta(md map[string]int, s string) {
	md[s] = 100
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	sliceDelta(xi)
	fmt.Println(xi)

	// map
	m := make(map[string]int)
	m["James"] = 30
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])

}
