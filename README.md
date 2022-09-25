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
