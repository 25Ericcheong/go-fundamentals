package main

import "fmt"

const helloPrefix = "Hello"

func Hello(name string) string {
	return fmt.Sprintf("%v %v, from a Function!", helloPrefix, name)
}

func main() {
	fmt.Println(Hello("Eric"))
}
