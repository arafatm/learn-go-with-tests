package learngowithtests

import "testing"
import "fmt"

// TODO: Why does this work but not the Example in adder_test.go
func ExampleSum() {
	sum := Sum([5]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output 13
}

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
