package main

import (
	"go-fundamentals/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout)
}
