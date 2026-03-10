package main

import "fmt"

type Status int

const (
	EmptyName = iota + 1
	EmptyPassword
	LoginError
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func login(name string, password string) error {
	if name == "" {
		return StatusErr{Status: EmptyName, Message: "name cannot be empty"}
	} else if password == "" {
		return StatusErr{Status: EmptyPassword, Message: "password cannot be empty"}
	}

	if name == "foo" && password == "bar" {
		return nil
	}
	return StatusErr{Status: LoginError, Message: "login error"}
}

func loginDemo(name string, password string) {
	fmt.Printf("login(%#v, %#v) -> ", name, password)
	err := login(name, password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("OK")
	}
}

func main() {
	loginDemo("foo", "bar")
	loginDemo("bar", "foo")
	loginDemo("", "bar")
	loginDemo("foo", "")
}
