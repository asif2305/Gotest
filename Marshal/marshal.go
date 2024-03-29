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
		Age:   30,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   25,
	}

	people := []person{p1, p2}
	fmt.Println(people)
	/*
	   In Go programming language, a "marshal" typically refers to the process of converting data
	   from its in-memory representation to a byte stream or other serializable format.
	   This is often done when you want to transmit data over a network, store it in a file,
	   or send it between different parts of a program.
	*/
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	// unarshal
	ubs := []byte((bs))
	var people2 []person
	err2 := json.Unmarshal(ubs, &people2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("all of the data", people2)
	for i, v := range people2 {
		fmt.Println("--- person numer are: ", i)
		fmt.Println(v.First, v.Last, v.Age)

	}

}
