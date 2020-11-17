package services

import (
	"fmt"

	"../domain"
)

func GetItem(item string) (*domain.Item, error) {
	itemData, err := domain.GetItemData(item)
	return itemData, err
}

// GetItems From DB
func GetItems(item string) ([]domain.Item, error) {
	itemData, err := domain.GetItemsFromDB(item)
	fmt.Println(itemData)
	return itemData, err
}
