package hello

import (
	"go-fundamentals/utils"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello with name", func(t *testing.T) {
		got := Hello("Eric", "")
		want := "Hello, Eric"

		utils.AssertCorrectStringsMessage(t, got, want)
	})

	t.Run("saying hello without name should print 'Hello, World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		utils.AssertCorrectStringsMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		utils.AssertCorrectStringsMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"

		utils.AssertCorrectStringsMessage(t, got, want)
	})
}
