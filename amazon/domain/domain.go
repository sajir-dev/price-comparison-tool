package domain

import (
	"errors"
	"fmt"

	"../../config"
)

var Items map[string]Item = map[string]Item{"MacBookAir": {"897465", "MacBookAir", "$860", "Apple", "Amazon", "Flagship business laptop for presentation and meeting purposes from Apple", "4.45"}}

// Items = map["MacBookAir"]{"MacBookAir", "$800", "Apple", "Flipkart", "4.3"}

func GetItemData(item string) (*Item, error) {
	itemData := Items[item]
	if itemData.Brand == "" {
		return nil, errors.New("No item found")
	}
	return &itemData, nil
}

func GetItemsFromDB(item string) ([]Item, error) {
	rows, err := config.DB.Query("select * from items join marketplace using(itemid) where marketplace = 'amazon';")
	if err != nil {
		fmt.Println(err)
		return nil, err
		panic(err)
	}

	var Items []Item
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.ItemID, &item.ItemName, &item.Price, &item.Brand, &item.Platform, &item.Description, &item.Rating)
		if err != nil {
			fmt.Println(err)
			panic(err)
			return nil, err
		}
		Items = append(Items, item)
	}
	return Items, nil
}

// select * from items join marketplace using(itemid) where marketplace = 'flipkart';
// select * from items join marketplace using(itemid) where marketplace = 'amazon';
