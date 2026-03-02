package main

import (
	"fmt"
	"os"
)

func PrintNode(value int) {
	fmt.Print(value, " ")
}

func InsertNode(t *Tree, value int) *Tree {
	t, err := t.Insert(value)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return t
}

func main() {
	var tree *Tree
	values := []int{50, 25, 75, 100, 12}
	for _, value := range values {
		tree = InsertNode(tree, value)
	}

	fmt.Println("\n# In order traversal")
	tree.InOrderTraverse(PrintNode)
	fmt.Println()

	fmt.Println("\n# Pre order traversal")
	tree.PreOrderTraverse(PrintNode)
	fmt.Println()

	fmt.Println("\n# Post order traversal")
	tree.PostOrderTraverse(PrintNode)
	fmt.Println()

	fmt.Println("\n# Container test")
	candidates := []int{1, 25, 99}
	for _, candidate := range candidates {
		fmt.Println("-", candidate, "->", tree.Contain((candidate)))
	}
}
