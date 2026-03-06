package main

import "math"
import "fmt"

type Point2D struct {
	X, Y int
}

type Point3D struct {
	X, Y, Z int
}

func (p Point2D) GoString() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}
func (p Point2D) String() string {
	return p.GoString()
}

func (p1 Point2D) Diff(p2 Point2D) float64 {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	d := math.Sqrt((dx * dx) + (dy * dy))
	return d
}
