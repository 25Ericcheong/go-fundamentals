# Go Fundamentals
This repo is added for me to learn Go. I am currently going through this ~~[Go by Example](https://gobyexample.com/) tutorial~~ [Learn Go with Tests](https://github.com/quii/learn-go-with-tests) and will continue to update the Wiki (which then leads to new pages added and new links just to better organize the information I have gathered)

## Learn Go with Tests

### Go Hello World

- When writing go will have `main` package wiwth a `main` func in it. Constants should improve performance since it saves time having to create a specific type instance whenever code is called. However, simple example is negligible.


## About Tests

### Go Hello World

- Each test file will need to be in the following format `xxx_test.go` and test function must start with the word `Test`. Test function only takes one argument `t *testing.T`
and in order to use `*testing.T` type, will need to `import "testing"`
- It is also possible to run subset of tests for the same function (different arguments supplied) wit the use of `t.Run("Message to display when error occurs", func(t *testing.T) { testing code ehere })`
