# Go Fundamentals
This repo is added for me to learn Go. I am currently going through this ~~[Go by Example](https://gobyexample.com/) tutorial~~ [Learn Go with Tests](https://github.com/quii/learn-go-with-tests) and will continue to update the Wiki (which then leads to new pages added and new links just to better organize the information I have gathered)

## Learn Go with Tests

### Go Hello World

- When writing go will have `main` package wiwth a `main` func in it. Constants should improve performance since it saves time having to create a specific type instance whenever code is called. However, simple example is negligible.

```
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
```

- The above is a cool way of declaring a variable that will be assigned and return within a function (just my opinion - never seen such a syntax). The `prefix` will be assigned a `"zero"` value on the type (for a string it is `""` and for an int type it would be `0`). `return` will return the `prefix` variable declared in the method's signature.
- Function names starting with a lowercase letter would be private and anything that starts with an uppercase will be used publicly. 

### Integers

```
func Add(x, y int) int {
	return x + y
}
```

- We can define arguments with the same type with the following code above. A shortcut way of doing it.

### Iteration

- Go only has `for` loop used for anything that involves iteration. 

## About Tests

### Go Hello World

- Each test file will need to be in the following format `xxx_test.go` and test function must start with the word `Test`. Test function only takes one argument `t *testing.T`
and in order to use `*testing.T` type, will need to `import "testing"`
- It is also possible to run subset of tests for the same function (different arguments supplied) wit the use of `t.Run("Message to display when error occurs", func(t *testing.T) { testing code ehere })`
- When including helper functions, good to pass `t testing.TB` as an argument which is also an interface of `*testing.T` and `*testing.B`. Within the helper function, `t.Helper()` is used to tell the test suite that this method is used specifically as a helper. This is important since if something fafils, the lline number reported will be in the function call rather than inside the test helper function
