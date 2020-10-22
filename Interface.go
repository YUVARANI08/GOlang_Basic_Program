package main

import "fmt"

func main() {
	r1 := Rect{10, 7}
	c1 := Circle{8}
	fmt.Println(getArea(r1))
	fmt.Println(getArea(c1))
}

type Shape interface {
	area() float64
}
type Rect struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}

func (r1 Rect) area() float64 {
	return r1.width * r1.height
}
func (c1 Circle) area() float64 {
	return 3.14 * c1.radius * c1.radius
}
func getArea(shape Shape) float64 {
	return shape.area()
}
