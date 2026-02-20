package main

import "fmt"

func exercise01() {
	fmt.Println("\n# Exercise 1")

	var i int32 = 20
	fmt.Println("i =", i)

	var f float32 = float32(i)
	fmt.Println("f =", f)
}

func exercise02() {
	fmt.Println("\n# Excercise 2")

	const value = 113

	var i int32 = value
	fmt.Println("i =", i)

	var f float32 = value
	fmt.Println("f =", f)
}

func exercise03() {
	fmt.Println("\n# Excercise 3")

	var tiny byte = 255
	fmt.Println("tiny =", tiny)
	fmt.Println("tiny + 1 =", tiny+1)
	fmt.Println()

	var small int32 = 2_147_483_647
	fmt.Println("small =", small)
	fmt.Println("small + 1 =", small+1)
	fmt.Println()

	var big uint64 = 18_446_744_073_709_551_615
	fmt.Println("big =", big)
	fmt.Println("big + 1 =", big+1)
	fmt.Println()
}

func main() {
	exercise01()
	exercise02()
	exercise03()
}
