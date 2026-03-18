// Package main drives the whole project.
package main

import "fmt"

// main drives the whole project. This is an entry point.
func main() {
	src := Money{Value: 100.0, Currency: "USD"}
	dest, err := Convert(src, "CDN")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v = %v\n", src, dest)
	}

	dest, err = Convert(src, "VND")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v = %v\n", src, dest)
	}
}
