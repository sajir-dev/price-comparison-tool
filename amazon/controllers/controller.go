package controllers

import (
	"encoding/json"

	"../services"
)

// GetItem ...
func GetItem(item string) (string, error) {
	// itemData, err := services.GetItem(item)
	itemData, err := services.GetItems(item)
	if err != nil {
		return "", err
	}

	itemJson, _ := json.Marshal(itemData)
	return string(itemJson), nil
}
