package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Function is saying Hello to someone", func(t *testing.T) {
		got := Hello("Eric", "English")
		want := "Hello Eric, from a Function!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Function saying Hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello Nobody!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Function saying Hello to someone in mandarin", func(t *testing.T) {
		got := Hello("Eric", "Mandarin")
		want := "Ni Hao Eric, from a Function!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Function saying Hello to someone in french", func(t *testing.T) {
		got := Hello("Eric", "French")
		want := "Bonjour Eric, from a Function!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
