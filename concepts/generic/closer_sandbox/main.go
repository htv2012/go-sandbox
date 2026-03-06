package main

import "fmt"

func main() {
	fmt.Println("Hello")
	p1 := Point2D{X: 3, Y: 7}
	p2 := Point2D{X: 6, Y: 2}
	origin := Point2D{X: 0, Y: 0}

	fmt.Printf("origin: %#v\n", origin)
	fmt.Printf("p1:     %#v\n", p1)
	fmt.Printf("p2:     %#v\n", p2)

	pair1 := Pair[Point2D]{p1, origin}
	pair2 := Pair[Point2D]{p2, origin}
	closer := FindCloser[Point2D](pair1, pair2)
	fmt.Printf("Which one is closer to the origin? %v\n", closer)
}
