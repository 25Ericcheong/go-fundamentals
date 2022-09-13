package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// zero value for an int is 0
	var e int
	fmt.Println(e)

	// shorthand way for declaring and initializng a variable
	// in this case this is equivalent to
	// var f string = "aooke"
	f := "apple"
	fmt.Println(f)
}
