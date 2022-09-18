package main

import "fmt"

const helloEnglishPrefix = "Hello"
const helloMandarinPrefix = "Ni Hao"
const helloFrenchPrefix = "Bonjour"
const english = "English"
const french = "French"

func Hello(name string, language string) string {
	if name != "" {

		if language == english {
			return fmt.Sprintf("%v %v, from a Function!", helloEnglishPrefix, name)
		} else if language == french {
			return fmt.Sprintf("%v %v, from a Function!", helloFrenchPrefix, name)
		} else {
			return fmt.Sprintf("%v %v, from a Function!", helloMandarinPrefix, name)
		}
	}
	return "Hello Nobody!"
}

func main() {
	fmt.Println(Hello("Eric", ""))
}
