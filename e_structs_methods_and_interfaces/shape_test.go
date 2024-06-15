package structs_methods_and_interfaces

import (
	"go-fundamentals/utils"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	utils.AssertCorrectFloatMessage(t, got, want)
}

func TestArea(t *testing.T) {
	rect := Rectangle{12.0, 6}
	got := Area(rect)
	want := 72.0

	utils.AssertCorrectFloatMessage(t, got, want)
}
