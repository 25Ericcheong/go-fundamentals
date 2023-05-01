# Go Fundamentals
This repo is added for me to learn Go. I am currently going through ~~[Go by Example](https://gobyexample.com/) tutorial~~ [Learn Go with Tests](https://github.com/quii/learn-go-with-tests) and will continue to update the Wiki (which then leads to new pages added and new links just to better organize the information I have gathered)

I am beginning to transfer my notes into [Wiki](https://github.com/25Ericcheong/go-fundamentals/wiki) now. This will also serve as an opportunity to review that I have learned in the past.

## Learn Go with Tests

### Dependency Injection

- Helps with testing and allows us to write general-purpose functions. Worth noting that the `Writer` interface will be seen quite regularly. It is a great general purpose interface for "put this data somewhere".
- How to test something that prints to console (`stdout`). Will need to inject (meaning - pass in) the dependency of printing. In example, we try to test something that prints (`Printf`).
- Upon inspection, `Printf` implementation is as such:

```
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
```

- `Printf` calls `Fprintf` while passing in `os.Stdout`. What is `os.Stdout`? As such, further inspection of `Fprintf` then leads to the following implementation:

```
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

- We can see that `os.Stdout` implements `io.Writer`. Note, `io.Writer` interface could be seen quite often because it is a great general purpose interface for "put this data somewhere". The interface can be found to look like the following:

```
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

- The `os.Stdout` implements the method - `Write` which explains why it implements the `Writer` interface. Also worth noting that `Buffer` type from the `bytes` package also implement the `Writer` interface because it implements the same method from the `Writer` interface - `Write`.
- Worth noting that `Writer` interface is general purpose and is used in a variety of ways.

- Hard-wired dependencies or global state would make testing hard and slow. Example, global database connection pool used by a service later will make testing difficult and slow because it needs to be initialized and isolated for proper testing. Dependency injection will allow us to inject in a database dependency with an interface so that we can mock out with something we can control.
- Simnilar to example shown in the dependency example, test was refactored so that we could control where data was written by `injecting a dependency` (with an interface).
- `Separate your concerns`, decoupling where data goes from how to generate it. If a function has too many responsibilities (generating and writing to db or handling HTTP requests and doing domain level logic) then dependency injection is what we need.
- Code can be reused in different contexts - new dependencies can be used with our function (as long as same interface is used / same dependency is injected) 

### Concurrency

- Having more than one thing in progress. An operation that does not block code in Go will run in a seaprate process called `goroutine`. To start a new `goroutine` we turn a function into a `go` statement by putting the keyword `go` in front of the function.
- Only way to start a `goroutine` is to have `go` keyword in front of function call, usually use anonymous functions.
- Useful because they can be executed at the sane time they are declared - by having `()` after the ending curly brace. Next, they maintain access to the lexical scope they are defined in - all variables available at the point when anonymous function is declared are also avaialble in the body of the function. 
- If not handled correctly, will be hard to predict what is going to happen. Which is why tests are written, to know we are handling concurrency predictably. 
- Ensure that each `goroutine` doesn't reference the same item when a new `goroutine` process is created. Getting the last item only means that the processes before that is referring to the same item of the `map`. We have to make sure that the variable within the `goroutine` is fixed and we do this by defining a parameter for the anonymous function.
- Possible for `fatal error: concurrent map writes` to occur. This is due to  multiple `goroutine` writing to the map. As such, we need to spot race conditions with the built in `race detector`.
- This can be solved by coordinating goroutines using `channels`. Channels are Go data structure that can both receive and send values. This will allow communication between different processes. `chan` is the type - which stands for channel.
- A `send statement` would be `resultChannel <- result{u, wc(u)}` which in this case - we are sending `result` struct to the `resultChannel`.
- On the other hand, a `receive expression` is as such `r := <-resultChannel` where the channel is now on the right and the variable that we are assigninng to is on the left. 

### Select 

- Example included is a scenario where we try to test with real websites to test our logic.
- `defer` is a keyword that can be used to as a prefix on a function call so that the call of the function will be done at the end of the containing function. Most commonly used to cleanup resources - closing a file or closing a server so that it does not continue to listen to a port (in example)
- This is used early on but called at the end to improve readability.
- Always `make` channels. `ch := make(chan struct{})` is better over `var ch chan struct{}` because with `var`, variable will be initialised with the zero value of the type. `string` will be `""` and `int` will be `0`. For channels, the zero value is `nil` and sending it with `<-` will block forever because we can't send `nil` channels.
- Recall that you can wait for values to be sent to a channel with `myVar := <-ch`. This is a blocking call, as you are waiting for a value. `select` lets us wait on multiple channels. The first one to send a value "wins" and the code underneath the `case` is executed. 
- `time.After` returns a `chan` and will send a signal down it after amount of time defined. 

### Reflection

- `interface{}` used for when we don't know what the type is at compile time. Generally less performant (since need to do additional checks at runtime). As such, should only use reflection if you really need to.
- `reflect` used to attempt to look at the `interface{}` any type variable and look at its properties. The `reflect package` then uses the `ValueOf` which returns the `Value` of a given variable. We then have to make an assumption of the type.
- Since compiler will not help us with identifying if something is the right type or not, we will need to check the type of specific fields to ensure that they are of the right type at runtime.
- If `interface{}`, is a pointer. We will need to check for `Ptr` with the reflect package as we will need to acquire the underlyning value with `Elem()` before accessing its fields.
- If it is a `slice` type with `struct` type as its items, we can't call `NumField` on it directly. As `NumField` can only be called on `struct` type not a `slice`, as such, we will need to access each item of the `slice` and call `NumField` on the item instead.
- Note, `map` is similar to `struct` but keys are unknown at compile time

### Sync

- `sync.WaitGroup` is a way of synchronising concurrent processes: `WaitGroup` waits for a collection of goroutines to finish. The main goroutine calls `Add` to set the number of goroutines to wait for. Then each of the goroutines runs and calls `Done` when finished. At the same time, `Wait` can be used to block until all goroutines have finished.
- If we use `Wait` before any of our assertions in testing environment, we can be sure that all `goroutines` have been attempted.
- Test will fail when we try to allow multiple `goroutines` to mutate a value at the same time. As such we need to use a lock - with a `Mutex`. 
- Meaning, any `goroutine` that calls a method that calls the `Lock` method will prevent other goroutines from calling the same method - the rest will need to wait till the method has been `Unlock`ed.
- Should not embed `sync.Mutex` into struct because that will mean that the methods of this type will be part of the public interface - meaning we are allowing other code to couple themselves to it. 

### Context

- Used to manage long-running processes. We need to be able to handle a long process that could get cancelled unexpectingly. 
- Important to derive contexts so that cancellations are propagated throughout call stack for given request. 
- `context` has a method `Done()` which returns a channel which gets sent a signal when context is "done" or "cancelled". 
- Incoming requests to a server should create a `Context`, and outgoing calls to servers should accept a Context. The chain of function calls between them must propogate the Context so that when a Context is cancelled, all Contexts derived from it are also canceled.
- Context must be passed down to next responsible function in order for cancellation to propagate back up the chain if a cancellation does occur
- NEVER pass values through `context`. It is an untyped map so we do not have type-safety 

### Maths Lib

- Random. SVGs allow us to manipulate programmatically as they are described in XML. 

## About Tests

### Go Hello World

- Each test file will need to be in the following format `xxx_test.go` and test function must start with the word `Test`. Test function only takes one argument `t *testing.T`
and in order to use `*testing.T` type, will need to `import "testing"`
- It is also possible to run subset of tests for the same function (different arguments supplied) wit the use of `t.Run("Message to display when error occurs", func(t *testing.T) { testing code ehere })`
- When including helper functions, good to pass `t testing.TB` as an argument which is also an interface of `*testing.T` and `*testing.B`. Within the helper function, `t.Helper()` is used to tell the test suite that this method is used specifically as a helper. This is important since if something fafils, the lline number reported will be in the function call rather than inside the test helper function

### Go Iterations

- Benchmarks in Go is used to measure how long a block of code takes to execute. Framework will determine what is a "good" value. Example, for reference and keep in mind.

```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

- To run the benchmark - `go test -bench="."`. 

### Go Arrays and Slices

- `go test -cover` is used to check how much of code has been tested.

### Structs, Methods and Interfaces

- Table driven tests are useful when wanting to build a list of test cases that can be tested in the same manner.
- When creating table driven tests, it is important to name each anonymous struct type within the slice since it may have multiple fields and this will help with readability.
- Instead of the following error logged below when a test case fails:

```
--- FAIL: TestArea (0.00s)
    shape_test.go:30: got 36 want 31
```

- A simple change to the format within the `Errorf` from `t.Errorf("got %g want %g", got, tt.hasArea)` to `t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)` will allow us to find out the struct involved in the failed test as well. As such, the output log will be this instead (which would be more ideal):

```
--- FAIL: TestArea (0.00s)
    shape_test.go:30: shape.Triangle{Width:12, Height:6} got 36 want 31
```

- With proper definitions of each test cases, we can specifically test something within our table. `go test -run TestArea/Rectangle` will test the `Rectangle` test case only.

### Maps

- `Error` type can be converted ot string with the `.Error()`` method.

### Mocking

- Tests with time intervals (say ensuring that we want 1 second sleep between each second) will slow down developer productivity. We would not want tests to be dependent on dependencies like `Sleep`. 
- Instead, we could define dependency as `interface`. Allowing us to use real `Sleep` in `main` and a `spy sleeper` in tests. 
- `Spies` are a kind of mock which can record how a dependency is used, arguments sent in, how many times it has been called, etc.

### Select

- `http.HandlerFunc` is a type that looks like `type HandlerFunc func(ResponseWriter, *Request)`. The handler takes a `ResponseWriter` and a `Request`. Note, this is also how a real HTTP server is written in Go. For testing purposes, the `http.HandlerFunc` is currently being wrapped in the `httptest.NewServer` for testing purposes.
- The `httptest.NewServer` finds an open port to listen on for testing purposes which can then be closed after tested.

### Reflection

- Worth noting if we are testing and storing our test case values in an array of something. The order may also affect the results of our test when it comes to utilizing `map` types. We need to ensure order should not matter and the test case accounts for that as well.

### Property Based Testing

- Some interesting use of string builder and interesting point about indexing a string would give us a byte. Note this really isn't relevant to ptoperty based testing but it was found while I was going through this chapter
- This is part of the chapter. Property based tests helps us throw random data at code and verifying rules described are always true. Property based testing is ensuring we have a good understanding of domain so we can write these properties.
- `quick.Check` standard library is a great way of acquiring random inputs to test our understanding of our domain. 
- Random note, `unit16` are unsigned integers - meaning integers that cannot be negative. 16 bit integer can store a max number of `65535`. 

## Additional Investigation

- Get more familiar with Go Standard Library (may help with writing dependency injection for testing purposes)
- Understanding the purpose of `Buffer` and `Writer`
- Look into the term Test Doubles (https://martinfowler.com/bliki/TestDouble.html)
- To mock HTTP server, look at standard library - `net/http/httptest`; preferred to do more investigation in along with the methods that can be called like `NewRequest` and `NewRecorder`
- To read more about reflection - https://go.dev/blog/laws-of-reflection
- To read more about context - https://go.dev/blog/context. Look at examples and try to better understand the use of context and how to manage cancellations and how a function would accept `context` and use it to cancel itself with `goroutines`, `select` and `channels`. These are worth practicing and understanding
- More on `context` and examples here - https://blog.golang.org/context
- A deeper understanding of `http.HandlerFunc` and how that works with servers 
- Handling errors gracefully - https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
