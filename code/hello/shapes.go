package learngowithtests

import "math"

type Shape interface {
	Area() float64
}

type Triangle struct {
	base   float64
        height float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.width + rectangle.height)
	// test
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}

func (circle Circle) Area() float64 {
	return circle.radius * circle.radius * math.Pi
}

func (circle Circle) Perimeter() float64 {
	return circle.radius * 2 * math.Pi
}

func (triangle Triangle) Area() float64 {
	return (triangle.base * triangle.height)/2 
}
