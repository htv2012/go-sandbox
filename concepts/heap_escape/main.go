package main

import "fmt"

type Person struct {
	first     string
	last      string
	birthYear uint16
}

func NewPerson(first string, last string, birthYear int) Person {
	// This Person variable will be moved (escaped) to the heap
	//because it is out-last the function creating it
	return Person{first: first, last: last, birthYear: uint16(birthYear)}
}

func NewPersonP(first string, last string, birthYear int) *Person {
	// So is this variable
	return &Person{first: first, last: last, birthYear: uint16(birthYear)}
}

func main() {
	p := NewPerson("Anna", "Karenina", 27)
	pp := NewPersonP("Ken", "Thompson", 83)

	fmt.Println("Person:", p)
	fmt.Println("Person2:", *pp)
}
