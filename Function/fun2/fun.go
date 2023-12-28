package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
	//first string
}

// interface
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
	//fmt.Println("I am a secret agent", sa.person.first)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
		//first: "ahmed",
	}
	p2 := person{
		first: "Jenny",
	}
	//sa1.speak()
	//p2.speak()

	// interface with polymorphism
	saySomething(sa1)
	saySomething(p2)
}
