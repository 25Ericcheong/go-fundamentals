package dictionary

import (
	"go-fundamentals/utils"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		utils.AssertCorrectStringsMessage(t, got, want)
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
	dict.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dict.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	utils.AssertCorrectStringsMessage(t, got, want)
}
