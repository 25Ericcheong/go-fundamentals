package dictionary

import (
	"go-fundamentals/utils"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := Search(dict, "test")
		want := "this is just a test"

		utils.AssertCorrectStringsMessage(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := Search(dict, "unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		utils.AssertCorrectStringsMessage(t, err.Error(), want)
	})
}
