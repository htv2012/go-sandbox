package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todo Todo) CreateTableRow(index int) []string {
	row := []string{}
	row = append(row, strconv.Itoa(index))
	row = append(row, todo.Title)
	row = append(row, todo.CreatedAt.Format(time.RFC1123))
	if todo.Completed {
		completedAt := todo.CompletedAt.Format(time.RFC1123)
		row = append(row, "Yes", completedAt)
	} else {
		row = append(row, "No", "")
	}
	return row
}

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Completed = !t[index].Completed
	if t[index].Completed {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) Print() {
	table := NewTable()
	table.AddHeader([]string{"#", "Task", "Created On", "Completed", "Completed On"})
	for index, t := range *todos {
		table.AddRow(t.CreateTableRow(index))
	}

	table.Print()
}
