package domain

import (
	"encoding/json"
	"fmt"

	"../../config"
)

// Item struct defines an item
type Item struct {
	ItemID      string `json:"itemid"`
	ItemName    string `json:"itemname"`
	Price       string `json:"price"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	Platform    string `json:"platform"`
}

type FlipkartItem struct {
	ItemID      string `json:"itemid"`
	ItemName    string `json:"itemname"`
	Price       string `json:"price"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	Platform    string `json:"platform"`
}

func (f FlipkartItem) GetItems(itemname string) (string, error) {
	rows, err := config.DB.Query("select * from items join marketplace using(itemid) where marketplace = 'flipkart';")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var items []Item
	for rows.Next() {
		item := Item{}
		err = rows.Scan(&item.ItemID, &item.ItemName, &item.Price, &item.Brand, &item.Description, &item.Rating, &item.Platform)
		if err != nil {
			fmt.Println("flipkart item-dto", err)
			return "", err
		}
		items = append(items, item)
	}

	itemJson, _ := json.Marshal(items)

	return string(itemJson), nil
}
