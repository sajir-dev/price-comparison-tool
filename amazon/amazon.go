package amazon

import "./controllers"

// GetItem ...
func GetItem(item string) (string, error) {
	itemData, err := controllers.GetItem(item)
	return itemData, err
}
