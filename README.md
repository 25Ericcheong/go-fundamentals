# go-fundamentals
My attempt to learn Go via testing
- A chapter a day
- A continuous integration in place via GitHub actions to visually see green ticks for fun

<br/>

# Notes for each chapter
## Integers
- Could include an additional `testable example` in collection of test. In this instance, running the command `go test -v` would example the example and will execute the unit test. Remember to include `// Output` in the unit test method so that the unit test runs. If it is not, the test will not run.

## Iteration
- Learned to use benchmark testing. Useful for testing performance of new code with the following command `go test -bench=.`

## Arrays and Slices
- Use `go test -cover` to display test coverage if needed. Great way to determining areas of code that need testing for.
- The use of `make`, `slices`, `array`, `append`

## Structs, methods and interfaces
- Implicit interfaces and the use of table driven tests to build a list of test cases (anonymous structs are use in the table of driven tests too)
- Note the ability to run single test case instead of everything with the following command `go test -run TestArea/Rectangle`

## Pointers and Errors
- A field that starts with a lowercase symbol is private to outside the package it is defined in.
- In a function or a method in Go, the arguments are copied. It is important to use pointers so we ensure we take the pointer of a type that has already been instantiated to update/modify rather than update/modify a copy of it when the method is called (the use of a `*`)
- Methods within structs perform automatic de-referencing of the pointers. Whereas, in a normal function you would need to. Note the differences and instances when using `&` (gets the memory address for this variable currently in the code) and `*` (before a variable is for de-referencing and after a `type` we're saying it's a pointer for this `type`)