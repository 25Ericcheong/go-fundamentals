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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		utils.AssertCorrectFloatMoreDecimalsMessage(t, got, want)
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{12.0, 6}
		checkArea(t, rect, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
