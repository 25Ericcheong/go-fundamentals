package iteration

import (
	"go-fundamentals/utils"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaaa"

	utils.AssertCorrectStringsMessage(t, got, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
