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

func Filter[T any](s []T, f func(s T) bool) []T {
	var out []T
	for _, v := range s {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}
