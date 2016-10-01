package main

import (
	"fmt"
)

func (p *Point) invert() {
	p.X, p.Y = p.Y, p.X
}

// Point coordindate
type Point struct {
	X int
	Y int
}

func main() {

	center := Point{X: 5, Y: 2}
	center.invert()

	fmt.Println(center)
}
