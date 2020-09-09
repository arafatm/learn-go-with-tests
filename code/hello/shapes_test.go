package learngowithtests

import "fmt"

func ExamplePerimeter() {
	fmt.Println(Perimeter(Rectangle{10.0, 5}))
	// Output: 30
}

func ExampleArea() {
	fmt.Println(Area(Rectangle{12.0, 6.0}))
	// Output: 72
}
