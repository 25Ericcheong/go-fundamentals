package dictionary

import (
	"go-fundamentals/utils"
	"testing"
)

func TestSearch(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dict := Dictionary{word: definition}

	t.Run("known word", func(t *testing.T) {
		assertDefinition(t, dict, word, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		utils.AssertCorrectErrorMessage(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	word := "test"
	definition := "this is just a test"

	dict.Add(word, definition)

	assertDefinition(t, dict, word, definition)
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	utils.AssertCorrectStringsMessage(t, got, definition)
}
