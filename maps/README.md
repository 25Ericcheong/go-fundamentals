# Maps
- Example of with and without custom type for reference
- Can modify `map` without passing an address. Gotcha is that they can be `nil` value. `nil` map will cause runtime panic.
- To make errors more usable - https://dave.cheney.net/2016/04/07/constant-errors
- Can define constant error to be used and easily modified for future
- Better to be specific with errors a prime example in web app `You can redirect the user when ErrNotFound is encountered, but display an error message when ErrWordDoesNotExist is encountered.`