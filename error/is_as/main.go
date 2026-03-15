package main

import "fmt"

func main() {
	fmt.Println()
	IsDemo("main.go")
	IsDemo("foo.bar")
	IsDemo("/etc/sudoers")

	fmt.Println()
	CustomIsDemo(101, 102)
	CustomIsDemo(102, 101)

	fmt.Println()
}
