# GenericList
List of comparable type with insertion, deletion, retrieval

## How to install
Inside a go module
```bash
go get github.com/nabin3/GenericList
```

## User guide demonstrated by example
```go

import (
	"fmt"
	"log"

	"github.com/nabin3/GenericList/genericlist" // importing this library
)

func main() {
	smartList := genericlist.NewGenericList[string]()

	// Insertion --
	smartList.InsertNewItem("foo")
	smartList.InsertNewItem("bar")
	smartList.InsertNewItem("buzz")
	fmt.Printf("smartList: %+v \n", smartList)
	insertionIndex := smartList.InsertNewItem("car")
	fmt.Printf("car inserted in %d index of our list \n", insertionIndex)
	fmt.Printf("After insertion, updated smartList: %+v \n", smartList)

	// Deletion --
	item, err := smartList.DeleteByIndex(1)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	fmt.Printf("%v deleted \n", item)
	fmt.Printf("after deletion, updated smartList: %+v \n", smartList)

	index, err := smartList.DeleteByItem("buzz")
	if err != nil {
		log.Printf("%v", err)
		return
	}
	fmt.Printf("buzz deleted from %d index", index)
	fmt.Printf("after deletion, updated smartList: %+v \n", smartList)

	// Retrieval
	retrievedItem, err := smartList.Get(1)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	fmt.Printf("Retrieved item: %v\n", retrievedItem)
}
```
That's it