package main

import "fmt"

func increase(p *int, delta int) {
	*p += delta
}

func main() {
	x := 10
	px := &x

	fmt.Println("x =", x, "px =", px)
	increase(px, 5)
	fmt.Println("x =", x, "px =", px)
	increase(&x, 5)
	fmt.Println("x =", x, "px =", px)
}
