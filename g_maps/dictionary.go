package dictionary

import "errors"

type Dictionary map[string]string

func Search(dict Dictionary, word string) (string, error) {
	val, isFound := dict[word]

	if !isFound {
		return "", errors.New("could not find the word you were looking for")
	}

	return val, nil
}
