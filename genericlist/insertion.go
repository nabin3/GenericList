package genericlist

// Insert new item and return index where insertion happened
func (list *GenericList[T]) InsertNewItem(item T) int {
	list.listItems = append(list.listItems, item)

	// Returning index where given item inserted
	return len(list.listItems) - 1
}
