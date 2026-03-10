package main

import (
	"errors"
	"fmt"
)

type User struct {
	Uid   int
	Alias string
}

func (u User) Validate() error {
	var errs []error

	if u.Uid < 500 || u.Uid > 600 {
		errs = append(errs, fmt.Errorf("uid must be in range (500-600), got %d", u.Uid))
	}

	if u.Alias == "" {
		errs = append(errs, errors.New("alias must not be empty"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
