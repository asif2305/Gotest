package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	unmarshal := []byte(s)
	var unPeople []person
	fmt.Println("unPeople type %T", unPeople)
	err2 := json.Unmarshal(unmarshal, &unPeople)
	if err != nil {
		fmt.Println(err2)
	}
	fmt.Println("all of the data", unPeople)
	for i, v := range unPeople {
		fmt.Println("\nPerson number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
