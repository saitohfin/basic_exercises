package maps

//Dictionary is a string map
type Dictionary map[string]string

//Search find inside a dictionary
func (d Dictionary) Search(key string) string {
	return d[key]
}
