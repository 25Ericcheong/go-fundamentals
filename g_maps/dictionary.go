package dictionary

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Search(word string) (string, error) {
	val, isFound := d[word]

	if !isFound {
		return "", ErrNotFound
	}

	return val, nil
}
