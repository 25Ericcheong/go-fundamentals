package main

import "fmt"

const helloEnglishPrefix = "Hello"
const helloMandarinPrefix = "Ni Hao"
const helloFrenchPrefix = "Bonjour"
const english = "English"
const french = "French"

func Hello(name string, language string) string {
	if name != "" {

		helloPrefix := greetingPrefix(language)
		return fmt.Sprintf("%v %v, from a Function!", helloPrefix, name)
	}
	return "Hello Nobody!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = helloFrenchPrefix
	case english:
		prefix = helloEnglishPrefix
	default:
		prefix = helloMandarinPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Eric", ""))
}
