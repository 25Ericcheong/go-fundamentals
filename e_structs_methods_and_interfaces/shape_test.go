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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		utils.AssertCorrectFloatMoreDecimalsMessage(t, got, want)
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.want)
	}
}
