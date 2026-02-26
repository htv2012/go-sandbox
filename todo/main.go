package main

import (
	"fmt"
	"os"
)

func main() {
	dataPath, err := GetDataPath()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	todos := Todos{}
	storage := NewStorage[Todos](dataPath)
	storage.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
