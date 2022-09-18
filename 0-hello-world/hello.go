package main

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello %v, from a Function!", name)
}

func main() {
	fmt.Println(Hello("Eric"))
}
