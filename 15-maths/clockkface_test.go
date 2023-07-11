package clockface_test

import (
	clockface "15-maths"
	"testing"
	"time"
)

// note to explain where values are coming from
// every clock has a centre of (150, 150)
// the hour hand is 50 long
// the minute hand is 80 long
// the second hand is 90 long

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 - 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}