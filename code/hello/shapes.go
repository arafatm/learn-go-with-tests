package learngowithtests

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
	// test
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}

func Area(circle Circle) float64 {
	return circle.radius * circle.radius * math.Pi
}

func Perimeter(circle Circle) float64 {
	return circle.radius * 2 * math.Pi
}
