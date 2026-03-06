package main

import "fmt"

func main() {
	words := []string{"One", "Potato", "Two", "Potato", "BiCauSanKy"}
	wordLens := Map[string, int](words, func(word string) int {
		return len(word)
	})
	longest := Reduce[int, int](wordLens, 0, func(a int, b int) int {
		return max(a, b)
	})

	fmt.Println("Words:", words)
	fmt.Println("Lengths:", wordLens)
	fmt.Println("Longest:", longest)
}
