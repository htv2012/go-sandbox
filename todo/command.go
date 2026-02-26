package main

import "strings"
import "os"
import "strconv"
import "fmt"
import "flag"

type CmdFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new task specify the title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a new task, specify a new title")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a task by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle by index")
	flag.BoolVar(&cf.List, "list", false, "List all tasks")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error invalid format for edit, expect id:title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index")
		}

		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Delete != -1:
		todos.delete(cf.Delete)

	default:
		fmt.Println("Invalid command")
	}
}
