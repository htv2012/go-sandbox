package main

import (
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("%s -> %s", name, err)
	}
	f.Close()
	return nil
}

func demoFileChecker(name string) {
	err := fileChecker(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s -> OK\n", name)
	}

}
func main() {
	demoFileChecker("foo.bar")
	demoFileChecker("/etc/sudoers")
	demoFileChecker("main.go")
}
