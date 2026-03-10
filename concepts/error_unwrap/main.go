package main

import (
	"errors"
	"fmt"
)

func demoFileCheck(name string) {
	err := FileCheck(name)
	if err != nil {
		fmt.Printf("%s -> %s\n", name, err)
		if inner := errors.Unwrap(err); inner != nil {
			fmt.Printf("  %s\n", inner)
		}
	} else {
		fmt.Printf("%s -> OK\n", name)
	}
}
func main() {
	demoFileCheck("foo.bar")
	demoFileCheck("/etc/sudoers")
	demoFileCheck("main.go")
}
