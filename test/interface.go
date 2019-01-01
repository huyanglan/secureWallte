package test

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}
//For our example we’ll implement this interface on Rect and Circle types.

type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}
//To implement an interface in Go, we just need to implement all the methods in the interface. Here we implement Geometry on rects.

func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perim() float64 {
	return 2*r.width + 2*r.height
}
//The implementation for circles.

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.radius
}
//If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic Measure function taking advantage of this to work on any Geometry.

func Measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}
