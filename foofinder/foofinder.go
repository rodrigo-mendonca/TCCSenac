package foofinder

import "errors"

func IsItFoo(word string) (bool, error) {
	switch word {
	case "foo":
		return true, nil
	case "bar":
		err := errors.New("Oh noooooo, it's bar!")
		return false, err
	}
	return false, nil
}
