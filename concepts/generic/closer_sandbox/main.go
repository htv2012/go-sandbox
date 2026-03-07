package main

import "fmt"

func main() {
	p1 := Point2D{X: 3, Y: 7}
	p2 := Point2D{X: 6, Y: 5}
	origin := Point2D{X: 0, Y: 0}

	pair1 := Pair[Point2D]{p1, origin}
	pair2 := Pair[Point2D]{p2, origin}
	closer := FindCloser[Point2D](pair1, pair2)

	fmt.Printf("Pair1: %v\n", pair1)
	fmt.Printf("Pair2: %v\n", pair2)
	fmt.Printf("Which one is closer to the origin? %v\n", closer)
}
