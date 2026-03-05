package main

import "fmt"

type Team struct {
	Name    string
	Members []string
}

func (t Team) Print() {
	// fmt.Println()
	fmt.Println("- Team", t.Name)
	for _, name := range t.Members {
		fmt.Println("  -", name)
	}

}
