package main

import "fmt"

func main() {
	var intStack Stack[int]
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println("Pop int stack: ")
	for _, trial := range []int{1, 2, 3, 4} {
		fmt.Printf("%d. ", trial)
		val, ok := intStack.Pop()
		if ok {
			fmt.Println(val)
		} else {
			fmt.Println("END")
		}
	}
	fmt.Println()

	var stringStack Stack[string]
	stringStack.Push("Mary")
	stringStack.Push("Paul")
	stringStack.Push("Peter")

	fmt.Println("Is Paul in stack?", stringStack.Contains("Paul"))
	fmt.Println("Is John in stack?", stringStack.Contains("John"))
	fmt.Println()

	fmt.Println("Pop string stack")
	for _, trial := range []int{1, 2, 3, 4} {
		fmt.Printf("%d. ", trial)
		val, ok := stringStack.Pop()
		if ok {
			fmt.Println(val)
		} else {
			fmt.Println("END")
		}
	}

}
