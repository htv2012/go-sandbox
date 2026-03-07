package main

import (
	"fmt"
)

type Pair[T fmt.Stringer] struct {
	Val1 T
	Val2 T
}

func (p Pair[T]) String() string {
	return fmt.Sprintf("%v-%v", p.Val1, p.Val2)
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pair1, pair2 Pair[T]) Pair[T] {
	d1 := pair1.Val1.Diff(pair1.Val2)
	d2 := pair2.Val1.Diff(pair2.Val2)

	if d1 <= d2 {
		return pair1
	}
	return pair2
}
