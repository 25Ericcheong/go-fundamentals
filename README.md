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

- Above, is a way of declaring a variable that will be assigned and return within a function (just my opinion - never seen such a syntax). The `prefix` will be assigned a `"zero"` value on the type (for a string it is `""` and for an int type it would be `0`). `return` will return the `prefix` variable declared in the method's signature.
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
- Apart from the example included in the Iterations directory. This is a `for` loop variation example that I thought was quite cool and would like to be mindful of.

```
i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}
```

### Array & Slices

- When iterating through an array, we can use `range` to iterate through array instead. An example is as shown below.

```
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
```

- `range` returns two values; the index and value. In this case, `_` from the code is ignored since it is not used
- Interestingly, a function expecting a `[5]int` as an argument type and is passed with a `[4]int` will not compile. This can be the same as passing an `int` into a function as an argument that expects a `string`. 
- Due to that, `slices` would usually be used instead because it does not encode the size eof the collection and instead can have any size.
- Creating slices would be as simple as `[]int` and `make([]int, sizeOfSlice)` - allows us to create a new slice of the type inserted as the first argument along with the size of the slice as the second argument.
- Cannot use equality operators with slices. Instead, should use `reflect.DeepEqual` which is useful for seeing if any 2 variables are the same. 
- Worth noting that `DeepEqual` doesn't type safe check values. 
- Instead of `make` which restricts the capacity of a slice, we can use the `append` function instead which appends new values to a slice - effectively allowing slice to grow in capacity as needed and will return a new slice along with the new appended value.
- Example of a good practice to slice and taking only subset of the array by utilizing  the `copy` function and prevent memory wastage as shown in the example below

```
a := make([]int, 1e6) // slice "a" with len = 1 million
b := a[:2] // even though "b" len = 2, it points to the same the underlying array "a" points to

c := make([]int, len(b)) // create a copy of the slice so "a" can be garbage collected
copy(c, b)
fmt.Println(c)
```

- An example of properly copying values from one slice type to another slice type

```
x := [3]string{"Лайка", "Белка", "Стрелка"}

y := x[:] // slice "y" points to the underlying array "x"

z := make([]string, len(x))
copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

y[1] = "Belka" // the value at index 1 is now "Belka" for both "y" and "x"

fmt.Printf("%T %v\n", x, x)
fmt.Printf("%T %v\n", y, y)
fmt.Printf("%T %v\n", z, z)
```

### Structs, Methods and Interfaces

- Struct is a named collection of fields where we can store data. Worth noting that the syntax for writing a method for a struct is as such `func (receiverName ReceiverType) MethodName(args)`.
- In Go, there is no need to explicitly have a `struct` implement an `interace`. It implicitly ensures structs implements the interfaces with the same method signature. 
- Anonymous struct can be used in table driven test to inclcude fields for the struct itself and the expected value which will be looped through for test cases.

### Pointers and Errors

- If a symbol (variable, types, functions, etc.) starts with a lowercase symbol then it is private.
- Calling a function or method - the arguments are copied. Example when calling `func (w Wallet) Deposit(amount int)` the `w` is a copy of whatever we call the method from. If we have differe methods being called from different parts of our code, the arguments called within the methods will have its own memory address. The `struct` defined with a variable within a different script would also have its own designated memory address for its variable as well. This would mean that manipulating a variable within a method from a different script would not affect a variable defined within the `struct` that has been initialized in a different script.
- This is where `pointers` come into play. Pointers would allow us to point to some values and then let us to change them. So, rather than having a copy, we would take a pointer to that variable / object so that we can change the original values within it.
- Go allow us to create a different type on top of existing ones to be more descriptive and explicit about why we have a specific `struct` to represent something. The syntax would be `type MyName OriginalType`. 
- With a new type defined, we can now declare methods on them. This allow us to be add domain specific functionality on top of existing types. In this case, `stringer` interface was implemented on our type (called Bitcoin). This interface will allow us to define how our type will be printed when used with the `%s` format string in prints. Example of code is as such.

```
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
```

- `nil` is the same as `null`. Errors can be `nil` because a return type `error` is an interface. A function that takes arguments or returns values that are interfaces, can be nillable.
- Important to have assertion for when error is not supposed to be returned as well. This can be checked with the following `go packages` - `go install github.com/kisielk/errcheck@latest` and runing `errcheck .` within the directory to check if we have missed anything out. 

### Maps

- Store items by `key` and look them up quicky. Similar to a dictionary. Declaring `map` will require two types. First would be the key type (written within the `[]`) and the second would be the value type, which goes after the `[]`. Note that the `key` can only be a comparable type.
- Similar to how we can create a `custom type` for readability from previous example. We can also do the same and create a `custom type` called `Dictionary` that thinly wraps around `map`. 
- With this, we can then create a method that can be used to search up for values with keys when/if the `custom type` is created. Note that, the method can only be used when the `custom type` has been instantiated.
- When attempting to access a value, the `map` returns 2 values. First would ne the value that corresponds to the key amd the second is a `boolean` that indicates if the key was found successfully.
- Interestingly, do not need to ensure pointers for `map` type. Because, a map value is a pointer to a runtime.hmap structure. As such, we should never initialize an empty map variable like this `var m map[string]string`. Instead, we can initialize an empty map like this `var dictionary = map[string]string{}` or with the `make` like this `var dictionary = make(map[string]string)`. This prevent a runtime panic as we will create an empty `hash map`.

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
- Incoming requests to a server should create a `Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propogate the Context so that when a Context is cancelled, all Contexts derived from it are also canceled.

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

## Additional Investigation

- Get more familiar with Go Standard Library (may help with writing dependency injection for testing purposes)
- Understanding the purpose of `Buffer` and `Writer`
- Look into the term Test Doubles (https://martinfowler.com/bliki/TestDouble.html)
- To mock HTTP server, look at standard library - `net/http/httptest`
- To read more about reflection - https://go.dev/blog/laws-of-reflection
- To read more about context - https://go.dev/blog/context. Look at examples and try to better understand the use of context
