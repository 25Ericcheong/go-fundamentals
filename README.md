# Go Fundamentals
This repo is added for me to learn Go. I am currently going through ~~[Go by Example](https://gobyexample.com/) tutorial~~ [Learn Go with Tests](https://github.com/quii/learn-go-with-tests) and will continue to update the Wiki (which then leads to new pages added and new links just to better organize the information I have gathered)

I am beginning to transfer my notes into [Wiki](https://github.com/25Ericcheong/go-fundamentals/wiki) now. This will also serve as an opportunity to review that I have learned in the past.

## Learn Go with Tests

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
- Concurrency examples or patterns which I think is quite informative and useful - https://go.dev/talks/2012/concurrency.slide#14 . Helpful point. When trying to acquire values from channels in a for loop, it is important to understand this statement; when the main function executes <– c, it will wait for a value to be sent. Similarly, when the boring function executes c <– value, it waits for a receiver to be ready.
