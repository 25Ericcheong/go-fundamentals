package main

import "fmt"

const helloEnglishPrefix = "Hello"
const helloMandarinPrefix = "Ni Hao"
const helloFrenchPrefix = "Bonjour"
const english = "English"
const french = "French"

func Hello(name string, language string) string {
	if name != "" {

		helloPrefix := helloMandarinPrefix
		switch language {
		case french:
			helloPrefix = helloFrenchPrefix
		case english:
			helloPrefix = helloEnglishPrefix
		}

		return fmt.Sprintf("%v %v, from a Function!", helloPrefix, name)
	}
	return "Hello Nobody!"
}

func main() {
	fmt.Println(Hello("Eric", ""))
}
