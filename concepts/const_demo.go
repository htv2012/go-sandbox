package main

import "fmt"

const x = 10

const (
	idKey   = "id"
	nameKey = "name"
)

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)
}
