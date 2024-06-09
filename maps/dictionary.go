package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("this word was not found")
	ErrExistentWord = errors.New("it's not possible add this word, because this word already exist in our dictionary")
)

func (d Dictionary) Find(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Find(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrExistentWord
	default:
		return err

	}

	return nil
}