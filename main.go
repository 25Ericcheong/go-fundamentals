package main

import (
	"go-fundamentals/mocking"
	"os"
)

func main() {
	defaultSleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, defaultSleeper)
}
