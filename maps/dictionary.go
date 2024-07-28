package dictionary

type DictionaryError string

var (
	ErrWordNotFound     = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("the word already exists")
	ErrWordDoesNotExist = DictionaryError("the word does not exist")
)

func (e DictionaryError) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definintion, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return definintion, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err

	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
