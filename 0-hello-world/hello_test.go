package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Function is saying hello to someone", func(t *testing.T) {
		got := Hello("Eric")
		want := "Hello Eric, from a Function!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Function saying Hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello Nobody!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
