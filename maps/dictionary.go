package main



const (
	ErrWordNotFound = DictionaryOfErrors("sorry, could not found the word you requested")
	ErrWordExists = DictionaryOfErrors("sorry, the word already exists")
	ErrWordDoesNotExist = DictionaryOfErrors("sorry, the word does not exist")
)

type DictionaryOfErrors string

func (e DictionaryOfErrors) Error() string {
	return string(e)
}

type Dictionary map[string]string

func(d Dictionary) Search( searchText string) (string, error) {
	definition, ok := d[searchText]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition,nil
}

func (d Dictionary) Add(word, definition string)  error {
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

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
	
}