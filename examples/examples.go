package main

import (
	"fmt"

	"github.com/25Ericcheong/go-fundamentals/hello"
)

func main() {
	theme := "to learn Golang heh heh"
	toOutput := hello.HelloTheme(theme)
	fmt.Println(toOutput)
}
