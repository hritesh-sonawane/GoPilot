package loop

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("x")
	expected := "xxxxx"

	if repeated != expected {
		t.Errorf("Got: %q | Want: %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("x")
	}
}

func ExampleRepeat() {
	rep := Repeat("x")
	fmt.Println(rep)
	// Output: xxxxx
}
