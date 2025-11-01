package main

import (
	"go-fundamentals/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, InnerSleep: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
