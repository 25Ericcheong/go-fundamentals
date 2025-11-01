package maps

import "errors"

// Search Before introducing custom type
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

// Search After introducing custom type
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
