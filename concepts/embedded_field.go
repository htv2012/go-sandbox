package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func (e Employee) Print() {
	fmt.Printf("%s(%d)\n", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) Print() {
	m.Employee.Print()
	for _, e := range m.Reports {
		fmt.Print("- ")
		e.Print()
	}
}
func main() {
	m := Manager{
		Employee: Employee{ID: 1003, Name: "Anna"},
		Reports: []Employee{
			Employee{ID: 1001, Name: "Karen"},
			Employee{ID: 1002, Name: "Ken"},
		},
	}
	m.Print()
}
