package greet

import (
	"bytes"
	"go-fundamentals/utils"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	utils.AssertCorrectStringsMessage(t, got, want)
}
