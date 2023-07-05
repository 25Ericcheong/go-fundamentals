# Go Fundamentals
This repo is added for me to learn Go. I am currently going through ~~[Go by Example](https://gobyexample.com/) tutorial~~ [Learn Go with Tests](https://github.com/quii/learn-go-with-tests) and will continue to update Wiki (which then leads to new pages added and new links just to better organize the information I have gathered)

I am beginning to transfer my notes into [Wiki](https://github.com/25Ericcheong/go-fundamentals/wiki) now. This will also serve as an opportunity to review what I have learned in the past.

## Learn Go with Tests

- Everything can be found in the [Wiki](https://github.com/25Ericcheong/go-fundamentals/wiki). I used to have all my notes in this file but soon realized that it is not manageable with some chapters requiring more revision (like channels).

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
