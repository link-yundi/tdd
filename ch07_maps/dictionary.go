package ch07_maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (definition string, err error) {
	v, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	// 1. word not exits
	case ErrNotFound:
		d[word] = definition
	// 2. word exits
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, new_definition string) error {
	// 1. word exits -> change definition
	// 2. word not exits -> ErrWordDoesNotExist
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = new_definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
