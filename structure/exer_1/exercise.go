package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favIC: []string{"Chocolate", "Banana", "passion fruit with mango ang apple"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		favIC: []string{"Strawberry", "Chocolate"},
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.favIC)
	fmt.Println(p2.favIC)

	for _, v := range p1.favIC {
		fmt.Println(p1.first, "favorite is", v)
	}

	for _, v := range p2.favIC {
		fmt.Println(p2.first, "favorite is", v)
	}

	// map with struct
	fmt.Println("................................")
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for key, v := range m {
		fmt.Println(key, v)
		for innerKey, v2 := range v.favIC {
			fmt.Println(innerKey, v.first, v.last, v2)
		}
	}

}
