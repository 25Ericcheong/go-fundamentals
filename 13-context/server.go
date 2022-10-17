package main

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		// Fetch ran in a goroutine whcih will write the result into a new channel
		go func() {
			data <- store.Fetch()
		}()

		// use select to race processes and we either write a response or cancel
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}

		fmt.Fprint(w, store.Fetch())
	}
}
