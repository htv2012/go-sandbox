// array.go
package main

import "fmt"
import "slices"

func show[S any, T any](label S, ar []T) {
	fmt.Printf("  %v = %v, (len = %d, cap = %d)\n", label, ar, len(ar), cap(ar))
}

func create_demo() {
	fmt.Println("\n# Slice: use [] for slice, [...] for array")
	var s1 = []int{1, 2, 3}
	show("s1", s1)

	fmt.Println("\n# Empty slice")
	var s2 []int
	show("s2", s2)

	fmt.Println("\n# Append 1, 2, 3")
	s2 = append(s2, 1, 2, 3)
	show("s2", s2)

	fmt.Println("\n# Compare, s1 == s2 won't work")
	fmt.Println("  s1 == s2?", slices.Equal(s1, s2))
}

func make_demo() {
	fmt.Println("\n# s := make([]int, 5)")
	s := make([]int, 5)
	show("s", s)

	s2 := make([]int, 5, 25)
	show("s2", s2)
}

func clear_demo() {
	fmt.Println("\n# Create a slice with 3 strings")
	s := []string{"Peter", "Paul", "Mary"}
	show("s", s)

	fmt.Println("# clear(s)")
	clear(s)
	show("s", s)
}

func slicing_demo() {
	fmt.Println("\n# Slicing Demo")
	s := []string{"zero", "one", "two", "three", "four"}
	show("s", s)
	show("s[:2]", s[:2])
	show("s[1:]", s[1:])
	show("s[1:3]", s[1:3])
	show("s[:]", s[:])
}

func slice_and_append() {
	fmt.Println("\n# Slice and Append Demo")
	a := []int{0, 1, 2, 3}
	b := a[:2]
	c := a[2:]

	show("a", a)
	show("b = a[:2]", b)
	show("c = a[2:]", c)

	fmt.Println("# b = append(b, 20)")
	b = append(b, 20)
	show("a", a)
	show("b", b)
	show("c", c)

	fmt.Println("# c = append(c, 40, 50)")
	c = append(c, 40, 50)
	show("a", a)
	show("b", b)
	show("c", c)
}

func main() {
	create_demo()
	make_demo()
	clear_demo()
	slicing_demo()
	slice_and_append()
}
