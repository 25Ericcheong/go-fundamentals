package main

import (
	"os"
	"time"

	clockface "15-maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}