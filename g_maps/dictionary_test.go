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

	t.Run("new word", func(t *testing.T) {
		err := dict.Add(word, definition)

		utils.AssertCorrectErrorMessage(t, err, nil)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{word: definition}
		err := dict.Add(word, "another definition")

		utils.AssertCorrectErrorMessage(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dict.Update(word, newDefinition)

		utils.AssertCorrectErrorMessage(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "new definition"
		dict := Dictionary{}

		err := dict.Update(word, definition)
		utils.AssertCorrectErrorMessage(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test definition"}

	dict.Delete(word)

	_, err := dict.Search(word)
	utils.AssertCorrectErrorMessage(t, err, ErrNotFound)
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	utils.AssertCorrectStringsMessage(t, got, definition)
}
