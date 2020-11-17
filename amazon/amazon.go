package amazon

import (
	"fmt"

	"./controllers"
)

// GetItem ...
func GetItem(item string) (string, error) {
	itemData, err := controllers.GetItem(item)
	fmt.Println("from amazon main", itemData)
	return itemData, err
}
