package structcs_methods_and_interfaces

import (
	"go-fundamentals/utils"
	"testing"
)

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	utils.AssertCorrectFloatMessage(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	utils.AssertCorrectFloatMessage(t, got, want)
}
