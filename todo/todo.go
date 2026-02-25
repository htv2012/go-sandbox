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

func (todos *Todos) print() {
	table := NewTable()
	table.AddHeader([]string{"#", "Task", "Completed", "Created On", "Completed On"})
	for index, t := range *todos {
		completed := "No"
		completedAt := ""

		if t.Completed {
			completed = "Yes"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow([]string{strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt})
	}

	table.Print()
}
