package main

import "fmt"

type Person struct {
	first     string
	last      string
	birthYear uint16
}

func NewPerson(first string, last string, birthYear int) Person {
	return Person{first: first, last: last, birthYear: uint16(birthYear)}
}

func NewPersonP(first string, last string, birthYear int) *Person {
	return &Person{first: first, last: last, birthYear: uint16(birthYear)}
}

func main() {
	p := NewPerson("Anna", "Karenina", 27)
	fmt.Println("Person:", p)

	pp := NewPersonP("Ken", "Thompson", 81)
	fmt.Println("Person2:", *pp)
}
