package structs_methods_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 10}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}

	// using table driven tests below with struct anonymous type and adding name to identify each test name uniquely
	//t.Run("rectangles", func(t *testing.T) {
	//	rect := Rectangle{Width: 12, Height: 6}
	//	checkArea(t, rect, 72.0)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t, circle, 314.1592653589793)
	//})

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{Width: 12, Height: 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		// use tt.name from case to use as a test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
