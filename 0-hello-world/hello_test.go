package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello from a Function!"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
