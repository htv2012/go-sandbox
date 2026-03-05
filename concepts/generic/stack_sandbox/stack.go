package main

type Stack[T any] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}

	last := len(s.vals) - 1
	top := s.vals[last]
	s.vals = s.vals[:last]
	return top, true
}
