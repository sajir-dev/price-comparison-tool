package services

import (
	"fmt"

	"../domain"
)

func GetItem(itemName string, marketplace string) ([]domain.Item, error) {
	item, err := domain.GetItem(itemName, marketplace)
	fmt.Println("from app services", item)
	return item, err
}
