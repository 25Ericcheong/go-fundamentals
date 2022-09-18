package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Eric")
	want := "Hello Eric, from a Function!"

	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
