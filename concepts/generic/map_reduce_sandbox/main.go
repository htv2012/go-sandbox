package main

import "fmt"

func main() {
	words := []string{"One", "Potato", "Two", "Potato", "BiCauSanKy"}
	filtered := Filter[string](words, func(word string) bool {
		return word != "Potato"
	})
	lengths := Map[string, int](filtered, func(word string) int {
		return len(word)
	})
	longest := Reduce[int, int](lengths, 0, func(a int, b int) int {
		return max(a, b)
	})

	fmt.Println("Words:", words)
	fmt.Println("Filtered:", filtered)
	fmt.Println("Lengths:", lengths)
	fmt.Println("Longest:", longest)
}
