package main

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	out := make([]T2, len(s))
	for i, v := range s {
		out[i] = f(v)
	}
	return out
}

func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}
