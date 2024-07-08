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
- Ability to create new types from existing ones. Example is `type Ringgit int` with `Ringgit` being the new type name
- Stringer Interface already defined in fmt package and essentially is called when printing using the %s format. Allows us to customize how our type should be printed
- `(t testing.TB).Fatal` stops test and prevents more unnecessary asserts from being performed
- `var` keyword allows us to define values global to the package
- `errcheck` a Golang  linter helps with identifying missing code blocks that do not check for `error` being returned

## Maps
- Maps can return 2 values - the value and whether key exists in map
- Maps can be `nil`. Ensure to never initialize an empty map variable like so `var m map[string]string`
- Do not need to pass `map` type as an address like usual with `*type` as Go makes a copy of just the pointer part already for us but not the underlying data structure that contains the data. Maps are pointers to runtime types
- Create a specific `error type` by ensuring type implements the `error` interface with the implementation of the `Error()` method

## Dependency Injection
- Learnt that `httpResponseWriter` implements `io.Writer` which is a clear indication of how handlers can be tested with the use of dependency injection and testing the output of the handlers via its response writers

## Mocking
- A [great example of mocking](https://github.com/25Ericcheong/go-fundamentals/commit/7902b04b1e1ecf2f249de3d6f2daa8a56365e4fe). This allow us inject a mock dependency for testing purposes and also still ensure code work as intended in prod (specific to time in this instance). This is important as it promotes loose coupling and allow us to test our dependency
- Be aware that [an existing spy object](https://github.com/25Ericcheong/go-fundamentals/commit/987fa702ad3c858521998b71fa0fa6aa939ffd32) can be further extended to implement multiple interfaces as long as it implements the method signature listed in the interface. This allow us to test multiple methods in this scenario

## Concurrency
- Note when using concurrency, "not enough time to add their results", when for looping each variable; important to ensure that a new variable is use for each iteration (each goroutine spawned via `go func(variableNameInGoRoutine){}(variablePassedIntoGoRoutine)` if not done same variable is replaced with a new variable for new goroutine and the error `fatal error: concurrent map writes` which is due to multiple `goroutines` writing to the results map at exactly the same time (specific to `map` in Go)
- If multiple things attempting to write to one thing, this is a race condition. To avoid such case, can use the command `go test -race`.
- This can be resolve appropriately with the use of `channels`. Channels coordinate the goroutines which can both receive and send values. This ability allows it to communicate between different processes. Remember: think about parent process and each of the `goroutines` that it makes to do the work of running a function.
- By sending results into the channel (receiver expression), we can control timing of each write into the results map, ensuring that it happens one at a time. Function calls in a goroutine and sending results into/from channel is happening concurrently inside its own process and finally each result is dealt with one at a time.

## Select
- `httptest` used in this chapter. Worth referring back to this for practice to test handlers or mock a server later on. `select` allows you to wait for multiple channels and use whichever that returns a value first.
- `time.After` will be useful to timeout in `select` block (since it returns a `chan`) to ensure that if none of the channels return anything we create a timeout instead

## Sync
- `sync.WaitGroup` is used to synchronise concurrent processes. It waits for a collection of goroutines to finish.
- The main goroutine calls `Add` to set the number of goroutines to wait for. Then each of the goroutines runs and calls `Done` when finished. At the same time, `Wait` can be used to block until all goroutines have finished.
- `Mutex` can be used to ensure that something is locked which ensures that only one goroutine can modify something one at a time. Do not directly embed `mutex` into struct and remember to not create a copy of mutex when passing to other functions which golang automatically does
- Common mistake is to overuse `channels` and `goroutines`. Use channels when passing ownership of data and use mutexes for managing state. Use `go vet` to catch subtle bugs. 
