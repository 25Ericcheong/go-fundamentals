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