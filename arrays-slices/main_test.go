package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 ints", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		expect := 10

		result := Sum(numbers)

		if result != expect {
			notEqualErr(t, expect, result)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should sum up elements", func(t *testing.T) {
		expect := []int{3, 9}
		result := SumAll([]int{1, 2}, []int{0, 9})

		if !reflect.DeepEqual(expect, result) {
			notEqualErr(t, expect, result)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("should sum up all except first element of array", func(t *testing.T) {
		expect := []int{2, 9}
		result := SumAllTails([]int{1, 2}, []int{0, 9})

		if !reflect.DeepEqual(expect, result) {
			notEqualErr(t, expect, result)
		}
	})

	t.Run("should sum up empty slices", func(t *testing.T) {
		expect := []int{0, 9}
		result := SumAllTails([]int{}, []int{3, 4, 5})

		if !reflect.DeepEqual(expect, result) {
			notEqualErr(t, expect, result)
		}
	})
}

func notEqualErr[T any](t testing.TB, expect, result T) {
	t.Helper()
	t.Errorf("Expected: %v, result: %v", expect, result)
}
