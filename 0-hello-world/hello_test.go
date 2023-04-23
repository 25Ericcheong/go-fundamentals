package main

import (
	"0-hello-world/assert"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people in English", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assert.CheckCorrectValue(t, got, want)
	})

	t.Run(("in Spanish"), func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assert.CheckCorrectValue(t, got, want)
	})

	t.Run(("in French"), func(t *testing.T) {
		got := Hello("Eric", "French")
		want := "Bonjour, Eric"

		assert.CheckCorrectValue(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assert.CheckCorrectValue(t, got, want)
	})
}
