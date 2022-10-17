package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// now a spy as we want to observe the way it is called
type SpyStore struct {
	response  string
	cancelled bool
}

// to simulate a long process
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// cancellingCtx is derived from request which returns a cancel function
		// we then schedule function to be called in 5 ms
		// finally we use context in our request by calling request.WithContext
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}
