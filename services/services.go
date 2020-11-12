package services

import "../domain"

func GetItem(itemName string, marketplace string) (*domain.Item, error) {
	item, err := domain.GetItem(itemName, marketplace)
	return item, err
}
