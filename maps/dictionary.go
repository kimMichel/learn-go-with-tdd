package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("this word was not found")

func (d Dictionary) Find(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}