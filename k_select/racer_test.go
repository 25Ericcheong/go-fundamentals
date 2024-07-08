package racer

import (
	"go-fundamentals/utils"
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	utils.AssertCorrectStringsMessage(t, got, want)
}
