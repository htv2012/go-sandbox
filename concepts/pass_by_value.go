package main

import "fmt"

type User struct {
	Uid   int
	Alias string
}

// won't work because user is passed by value
func SetAliasThatFailed(user User, newAlias string) {
	user.Alias = newAlias
}

// work, preferred
func SetAlias(user User, newAlias string) User {
	user.Alias = newAlias
	return user
}

// also work
func SetAlias2(user *User, newAlias string) {
	*&user.Alias = newAlias
}

func main() {
	user := User{Uid: 501, Alias: "anna"}
	fmt.Println("User =", user)

	SetAliasThatFailed(user, "karen")
	fmt.Println("User =", user)

	user = SetAlias(user, "karen")
	fmt.Println("User =", user)

	SetAlias2(&user, "kim")
	fmt.Println("User =", user)
}
