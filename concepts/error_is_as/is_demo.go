package main

import (
	"errors"
	"fmt"
	"os"
)

func IsDemo(name string) {
	f, err := os.Open(name)
	if err == nil {
		f.Close()
		fmt.Println(name, "-> exists")
		return
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println(name, "-> does not exist")
	} else if errors.Is(err, os.ErrPermission) {
		fmt.Println(name, "-> permission denied")
	}
}
