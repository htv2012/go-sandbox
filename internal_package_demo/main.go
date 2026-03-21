package main

import (
	"fmt"
	"internal_package_demo/cow"
	"internal_package_demo/internal"
	"internal_package_demo/internal/bar"
	"internal_package_demo/internal/foo"
)

func init() {
	fmt.Println("main.init")
}

func main() {
	fmt.Println("main")
	cow.Hello()
	internal.Hello()
	foo.Hello()
	bar.Hello()
}
