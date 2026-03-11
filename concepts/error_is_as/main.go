package main

import (
	"errors"
	"fmt"
	"os"
)

func FileCheck(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("%s -> %w", name, err)
	}
	f.Close()
	return nil
}

func DemoFileCheck(name string) {
	err := FileCheck(name)
	if err == nil {
		fmt.Println(name, "-> OK")
		return
	}

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println(name, "-> does not exist")
	} else if errors.Is(err, os.ErrPermission) {
		fmt.Println(name, "-> permission denied")
	}
}
func main() {
	DemoFileCheck("main.go")
	DemoFileCheck("foo.bar")
	DemoFileCheck("/etc/sudoers")
}
