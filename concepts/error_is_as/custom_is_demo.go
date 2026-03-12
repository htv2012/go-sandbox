package main

import (
	"errors"
	"fmt"
	"slices"
)

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("MyErr%v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return slices.Equal(me.Codes, me2.Codes)
	}
	return false
}

func CustomIsDemo(code1, code2 int) {
	ref := MyErr{[]int{101, 102}}
	err := MyErr{[]int{code1, code2}}

	res := "does not match"
	if errors.Is(err, ref) {
		res = "matched"
	}

	fmt.Println(err, res, ref)

}
