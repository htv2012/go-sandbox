package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("%s -> %w", name, err)
	}
	f.Close()
	return nil
}

func demoFileChecker(name string) {
	err := fileChecker(name)
	if err != nil {
		fmt.Print(err)

		// unwrap. For demo only. In practice, use errors.Is and errors.As
		if wrapped := errors.Unwrap(err); wrapped != nil {
			fmt.Printf(" [%s]", wrapped)
		}
		fmt.Println()
	} else {
		fmt.Printf("%s -> OK\n", name)
	}

}
func main() {
	demoFileChecker("foo.bar")
	demoFileChecker("/etc/sudoers")
	demoFileChecker("main.go")
}
