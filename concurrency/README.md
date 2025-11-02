# Concurrency
- To start a goroutine put `go` in front of function call. Often paired with anonymous functions to start one.
- In concurrency, can get race condition. So in tests, can you `go test -race` to enable race detector built-in feature.
- To solve, we leverage `channels`. Channels are used to coordinate goroutines and can receive and send values.
- Use `go test .bench=.` to benchmark tests if needed