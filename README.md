# Go Fundamentals
This repo is added for me to learn Go. I am currently going through ~~[Go by Example](https://gobyexample.com/) tutorial~~ [Learn Go with Tests](https://github.com/quii/learn-go-with-tests) and will continue to update Wiki (which then leads to new pages added and new links just to better organize the information I have gathered)

I am beginning to transfer my notes into [Wiki](https://github.com/25Ericcheong/go-fundamentals/wiki) now. This will also serve as an opportunity to review what I have learned in the past.

## Learn Go with Tests

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
- Concurrency examples or patterns which I think is quite informative and useful - https://go.dev/talks/2012/concurrency.slide#14 . Helpful point. When trying to acquire values from channels in a for loop, it is important to understand this statement; when the main function executes <– c, it will wait for a value to be sent. Similarly, when the boring function executes c <– value, it waits for a receiver to be ready. In addition to what I have mentioned; https://go.dev/blog/pipelines
