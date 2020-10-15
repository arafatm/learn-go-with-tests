package learngowithtests

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
	// test
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}

func (c Circle) Area() float64 {
	return circle.radius * circle.radius * 3.14
}

func (c Circle) Perimeter() float64 {
	return circle.radius * 2 * 3.14
}
