package main

import "fmt"

type person struct {
	Name string
	age  int
}

func updateAgeNot(p person, newAge int) {
	p.age = newAge
}
func updateAgeWithPointer(p *person, newAge int) {
	p.age = newAge
}

func main() {
	person := person{
		Name: "Alice",
		age:  20,
	}
	updateAgeNot(person, 25)
	fmt.Println(person)
	updateAgeWithPointer(&person, 25)
	fmt.Println(person.age)

}
