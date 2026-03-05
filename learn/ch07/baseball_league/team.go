package main

import (
	"fmt"
	"strings"
)

type Team struct {
	Name    string
	Members []string
}

func (t Team) Print() {
	members := strings.Join(t.Members, ", ")
	fmt.Printf("- Team %s: %s\n", t.Name, members)
}
