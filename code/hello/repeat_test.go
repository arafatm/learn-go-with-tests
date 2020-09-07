package learngowithtests

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	const repeatCount = 7
	repeated := Repeat("ab", repeatCount)
	expected := "ababababababab"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("ab", 3))
	// Output: ababab
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 9)
	}
}
