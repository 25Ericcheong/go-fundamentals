package structs_methods_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 10, Height: 10}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rect := Rectangle{Width: 12, Height: 6}
	got := Area(rect)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
