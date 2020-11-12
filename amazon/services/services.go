package services

import "../domain"

func GetItem(item string) (*domain.Item, error) {
	itemData, err := domain.GetItemData(item)
	return itemData, err
}
