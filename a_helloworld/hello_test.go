package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello with name", func(t *testing.T) {
		got := Hello("Eric")
		want := "Hello, Eric"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying hello without name should print 'Hello, World'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
