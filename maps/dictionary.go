package maps

import (
	"errors"
)

// Search Before introducing custom type
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

// Search After introducing custom type
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrWordExists
	default:
		return err
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
