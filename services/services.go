package services

import (
	"fmt"

	"../domain"
)

func GetItem(itemName string, marketplace string) ([]domain.Item, error) {
	items, err := domain.GetItem(itemName, marketplace)
	fmt.Println("from app services", items)
	return items, err
}
