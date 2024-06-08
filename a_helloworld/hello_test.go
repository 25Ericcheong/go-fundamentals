package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Hi Eric")
	want := "Hi Eric"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
