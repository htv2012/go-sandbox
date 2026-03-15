package main

import (
	"fmt"
	"os"
)

type MyError struct {
	Status  int
	Message string
	Inner   error
}

func (me MyError) Error() string {
	return fmt.Sprintf("[%d] %s", me.Status, me.Message)
}

func (me MyError) Unwrap() error {
	return me.Inner
}

func FileCheck(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return MyError{
			Status:  1001,
			Message: "CHECK",
			Inner:   err,
		}
	}

	f.Close()
	return nil
}
