package slice

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %d | Want: %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v | Want: %v", got, want)
		}
	}

	t.Run("Sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSum(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})
}
