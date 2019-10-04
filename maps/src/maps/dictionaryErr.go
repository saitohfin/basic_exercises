package maps

//DictionaryErr message
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
