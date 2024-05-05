package genericlist

import (
	"errors"
)

// Retrieving item by index
func (list *GenericList[T]) Get(index int) (T, error) {
	// This variable is declared for returning defaultValue
	var defaultValule T

	// Checking if given index is out of bound
	if index > len(list.listItems)-1 {
		return defaultValule, errors.New("given index is out of bound")
	}

	return list.listItems[index], nil
}
