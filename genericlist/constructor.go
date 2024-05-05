package genericlist

func NewGenericList[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		listItems: []T{}, // Creates an empty slice of type T
	}
}
