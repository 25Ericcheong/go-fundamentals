package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func Search(dict Dictionary, word string) (string, error) {
	val, isFound := dict[word]

	if !isFound {
		return "", ErrNotFound
	}

	return val, nil
}
