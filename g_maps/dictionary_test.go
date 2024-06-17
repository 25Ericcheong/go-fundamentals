package dictionary

import (
	"go-fundamentals/utils"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "this is just a test"}

	got := Search(dict, "test")
	want := "this is just a test"

	utils.AssertCorrectStringsMessage(t, got, want)
}
