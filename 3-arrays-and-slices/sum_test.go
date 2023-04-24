package sum

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	CheckValue(t, got, want, numbers)
}

func CheckValue(t testing.TB, got, want int, input [5]int) {
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, input)
	}
}
