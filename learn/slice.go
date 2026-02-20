// array.go
package main

import "fmt"
import "slices"

func show(label string, ar []int) {
	fmt.Printf("%s = %v, is nill? %v (len = %d, cap = %d)\n", label, ar, ar == nil, len(ar), cap(ar))
}

func main() {
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
	fmt.Println("s1 == s2?", slices.Equal(s1, s2))

	fmt.Println("\n# Create an empty slice with capacity of 5")
	s3 := make([]int, 0, 5)
	show("s3", s3)
}
