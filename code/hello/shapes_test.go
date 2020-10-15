package learngowithtests

import "fmt"

func ExamplePerimeterRectangle() {
	r := Rectangle{10.0, 5}
	fmt.Println(r.Perimeter())
	// Output: 30
}

func ExampleAreaRectangle() {
	r := Rectangle{12, 6}
	fmt.Println(r.Area())
	// Output: 72
}
func ExamplePerimeterCircle() {
	c := Circle{10}
	fmt.Println(c.Perimeter())
	// Output: 62.83185307179586
}

func ExampleAreaCircle() {
	c := Circle{10}
	fmt.Println(c.Area())
	// Output: 314.1592653589793
}
