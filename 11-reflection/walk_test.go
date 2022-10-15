package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	// to test for multiple different scenarios
	// an array containing an anon struct type with the following fields
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			// remember, walk has already been implemented which utilizes the anonymous func to acquire the any type - interface{}'s field called Name (assuming we only have one field) that has a string type (assuming it is a string type) on Runtime.
			// this anonymous func essentially takes the Field mentioned above as an argument/parameter called input (we need to test that the function has successfully acquire such string which is why we store the value to check if the right string has been sent through)
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
