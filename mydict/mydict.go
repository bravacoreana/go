package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string // type can have methods

var errNotFound = errors.New("Not Found")
var errWordExist = errors.New("Already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil
	// if err==errNotFound {
	// 	return d[word] = def
	// } else if err == nil{
	// 	return errWordExist
	// }

}
