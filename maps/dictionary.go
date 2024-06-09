package maps

type Dictionary map[string]string

const (
	ErrNotFound = ErrDictionary("this word was not found")
	ErrExistentWord = ErrDictionary("it's not possible add this word, because this word already exist in our dictionary")
)

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

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