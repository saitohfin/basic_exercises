package maps

import "errors"

//Dictionary is a string map
type Dictionary map[string]string

//ErrNotFound error triggered when a word does not exist in the dictionary
var ErrNotFound = errors.New("could not find the word you were looking for")

//Search find inside a dictionary
func (d Dictionary) Search(key string) (word string, err error) {
	word, ok := d[key]
	if !ok {
		err = ErrNotFound
	}

	return word, err
}
