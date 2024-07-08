package main

import (
	"bytes"
	"go-fundamentals/utils"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	utils.AssertCorrectStringsMessage(t, got, want)
}
