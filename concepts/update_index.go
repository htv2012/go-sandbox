package main

import "fmt"

func update(array []int, index int, value int) int {
	old := array[index]
	array[index] = value
	return old
}

func main() {
	ar := []int{0, 10, 20}
	old := update(ar, 0, 100)
	fmt.Println("Old:", old)
	fmt.Println("ar:", ar)
}
