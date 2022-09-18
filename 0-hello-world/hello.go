package main

import "fmt"

const helloPrefix = "Hello"

func Hello(name string) string {
	if name != "" {
		return fmt.Sprintf("%v %v, from a Function!", helloPrefix, name)
	}
	return "Hello Nobody!"
}

func main() {
	fmt.Println(Hello("Eric"))
}
