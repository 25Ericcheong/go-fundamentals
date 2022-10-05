package concurrency

type WebsiteChecker func(string) bool

// made to associate itself with the return value of WebsiteChecker
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// send statement
			// send values to the channel
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// for loop iterates once for each of the urls
	for i := 0; i < len(urls); i++ {
		// receive expression
		// assign values that we have received (inserted into the channel) to a variable
		// as we iterate through, each value gets assigned to variable (replacing previous assigned value)
		// then use the variable to update the results
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
