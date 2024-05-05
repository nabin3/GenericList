package genericlist

import "errors"

func (list *GenericList[T]) DeleteByIndex(index int) (T, error) {
	// This variable is declared for returning defaultValue
	var defaultValule T

	// Checking if given index is out of bound
	if index > len(list.listItems)-1 {
		return defaultValule, errors.New("given index is out of bound")
	}

	// Storing item that eill be deleted
	willBeDeleted := list.listItems[index]

	// Deleting maintaing order
	list.listItems = append(list.listItems[:index], list.listItems[index+1:]...)

	return willBeDeleted, nil
}

// Delete a given itgem from list if item exists and return given item's index
func (list *GenericList[T]) DeleteByItem(item T) (int, error) {
	// Checking if given item exists in list
	for iterator := 0; iterator < len(list.listItems); iterator++ {
		if list.listItems[iterator] == item {
			// Deleting the item maintaing order
			list.listItems = append(list.listItems[:iterator], list.listItems[iterator+1:]...)
			return iterator, nil
		}
	}

	// In case item not found
	return -1, errors.New("given item doesn't exist in list")
}
