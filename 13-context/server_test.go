package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello, world"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	// ServeHTTP calls return function of Server function
	// Interestingly if a type implements this function - it will be treated as a HandlerFunc type
	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
