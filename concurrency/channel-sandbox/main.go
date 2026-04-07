package main

import "fmt"

func main() {
	que := make(chan string)

	go func() {
		que <- "Hello"
		que <- "world"
	}()

	word1 := <-que
	word2 := <-que
	fmt.Println(word1, word2)
}
