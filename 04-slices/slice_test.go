package slice

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numArr := [5]int{1, 2, 3, 4, 5}

		got := SumArr(numArr)
		want := 15

		if got != want {
			t.Errorf("Got: %d | Want: %d", got, want)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numSlice := []int{1, 2, 3, 4, 5}

		got := SumSlice(numSlice)
		want := 15

		if got != want {
			t.Errorf("Got: %d | Want: %d", got, want)
		}
	})
}
