package learngowithtests

import "fmt"

func ExamplePerimeterRectangle() {
	r := Rectangle{10.0, 5}
	fmt.Println(r.Perimeter())
	// Output: 30
}

func ExamplePerimeterCircle() {
	c := Circle{10}
	fmt.Println(c.Perimeter())
	// Output: 62.83185307179586
}

func checkArea(shape Shape) float64 {
	return shape.Area()
}

func ExampleAreaCircle() {
	fmt.Println(checkArea(Circle{10}))
	// Output: 314.1592653589793
}

func ExampleAreaRectangle() {
	fmt.Println(checkArea(Rectangle{12, 6}))
	// Output: 72
}
