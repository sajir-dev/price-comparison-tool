package domain

import (
	"encoding/json"
	"errors"
	"fmt"

	"../amazon"
	"../flipkart"
)

// GetOneItem queries an item from db
func GetItem(itemName string, marketplace string) ([]Item, error) {
	// TODO query an item from db

	// item0, err0 := flipkart.GetItem(itemName)
	// item1, err1 := amazon.GetItem(itemName)

	var OneItem []Item

	// err0 = json.Unmarshal([]byte(item0), Items[0])
	// err1 = json.Unmarshal([]byte(item1), Items[1])

	switch marketplace {
	case "flipkart":
		{
			item, err := flipkart.GetItem(itemName)
			if err != nil {
				return nil, err
			}

			json.Unmarshal([]byte(item), &OneItem)

			return OneItem, nil
		}

	case "amazon":
		{
			item, err := amazon.GetItem(itemName)
			if err != nil {
				return nil, err
			}

			json.Unmarshal([]byte(item), OneItem)
			fmt.Println(item)
			return OneItem, nil
		}

	default:
		{
			return nil, errors.New("marketplace not available")
		}
	}
}

// // GetAllItems returns a slice of items from db
// func GetAllItems() ([]Item, error) {
// 	// TODO query an item from db
// 	return nil, nil
// }
