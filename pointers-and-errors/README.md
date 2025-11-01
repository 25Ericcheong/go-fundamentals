# Pointers and Errors
- Arguments are copied when they are being passed into function. Use receiver type `*w` and will auto deference when called in method.
- Can create simple type to standard known types for others to understand context
- Creating global error variables in package would help give context to type of error
- `go install github.com/kisielk/errcheck@latest` for linter check to ensure we did not miss anything
- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully - more on error