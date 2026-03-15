package main

import "fmt"

func DemoUserValidation(uid int, alias string) {
	user := User{uid, alias}
	err := user.Validate()
	if err != nil {
		fmt.Printf("%+v -> %s\n", user, err)
	} else {
		fmt.Printf("%+v -> OK\n", user)
	}
}

func main() {
	fmt.Println()
	DemoUserValidation(501, "anna")
	DemoUserValidation(1501, "anna")
	DemoUserValidation(501, "")
	DemoUserValidation(1501, "")
	fmt.Println()
}
