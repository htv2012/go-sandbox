package main

import "fmt"

func loginDemo(user, password string) {
	fmt.Printf("login(%#v, %#v) -> ", user, password)

	err := login(user, password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}

func main() {
	loginDemo("foo", "bar")
	loginDemo("bar", "foo")
	loginDemo("", "foo")
	loginDemo("bar", "")
}
