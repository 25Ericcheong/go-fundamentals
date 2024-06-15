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

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6}
		got := rect.Area()
		want := 72.0

		utils.AssertCorrectFloatMoreDecimalsMessage(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		utils.AssertCorrectFloatMoreDecimalsMessage(t, got, want)
	})
}
