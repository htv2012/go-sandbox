package main

import "fmt"

type Stringer interface {
	String() string
}

type User struct {
	Uid   int
	Alias string
}

type Group struct {
	Gid  int
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("%s(%d)", u.Alias, u.Uid)
}

func (g Group) String() string {
	return fmt.Sprintf("%s[%d]", g.Name, g.Gid)
}

func Print(label string, obj Stringer) {
	fmt.Printf("%s: %s\n", label, obj.String())
}

func main() {
	user := User{Uid: 501, Alias: "Anna"}
	group := Group{Gid: 1001, Name: "Admin"}

	Print("User", user)
	Print("Group", group)
}
