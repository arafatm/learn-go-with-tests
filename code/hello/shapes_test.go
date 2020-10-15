package learngowithtests

import "fmt"

func ExamplePerimeterRectangle() {
	fmt.Println(Perimeter(Rectangle{10.0, 5}))
	// Output: 30
}

func ExampleAreaRectangle() {
	fmt.Println(Area(Rectangle{12.0, 6.0}))
	// Output: 72
}
func ExamplePerimeterCircle() {
	fmt.Println(Perimeter(Circle{10.0}))
	// Output: 62.8
}

func ExampleAreaCircle() {
	fmt.Println(Area(Circle{10}))
	// Output: 314
}
