package maps

//Dictionary is a string map
type Dictionary map[string]string

//ErrNotFound error triggered when a word does not exist in the dictionary
var ErrNotFound = DictionaryErr("could not find the word you were looking for")

//ErrWordDoesNotExist not found word
var ErrWordDoesNotExist = DictionaryErr("word does not exists")

//ErrWordExists triggere when we try to add a word which already exists in the dictionary.
var ErrWordExists = DictionaryErr("word already exists")

//NotErrorFound indicate in returns that have had any error
var NotErrorFound error = nil

//Search find inside a dictionary
func (d Dictionary) Search(key string) (word string, err error) {
	word, ok := d[key]
	if !ok {
		err = ErrNotFound
	}

	return word, err
}

//Add a word and his decription to the dictionary
func (d Dictionary) Add(word, description string) error {
	_, err := d.Search(word)
	if err != ErrNotFound {
		return ErrWordExists
	}

	d[word] = description
	return NotErrorFound
}

//Update word which already exists
func (d Dictionary) Update(word, description string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}

	d[word] = description
	return NotErrorFound
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
