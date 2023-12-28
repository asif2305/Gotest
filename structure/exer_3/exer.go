package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jennay": 27,
			"Q":      30,
			"Ian":    47,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends - ", k, v)
	}

	for k, v := range p1.favDrinks {
		fmt.Println(p1.first, "- drinks - ", k, v)
	}
}
