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

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{Width: 12, Height: 6}
		got := rect.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
