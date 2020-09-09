package learngowithtests

import "testing"
import "fmt"
import "reflect"

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output: 15
}

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// if got != want {
	// This won't work, have to use reflect.DeepEqual instead
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
