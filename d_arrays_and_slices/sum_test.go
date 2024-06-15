package arrays_and_slices

import (
	"go-fundamentals/utils"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	utils.AssertCorrectNumbersAndArrayMessage(t, got, want, numbers)
}
