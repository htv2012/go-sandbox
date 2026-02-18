package main

import "fmt"

func main() {
	stats := map[string]int{"health": 100}

	// The idiom in action
	value, ok := stats["mana"]

	if ok {
		fmt.Println("Mana level:", value)
	} else {
		fmt.Println("Mana key does not exist!")
	}
}
