package InterfaceSample

import (
	"encoding/json"
	"fmt"

	"../config"
)

type FlipkartItem struct {
	ItemID      string  `json:"itemid"`
	ItemName    string  `json:"itemname"`
	Price       float32 `json:"price"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Rating      string  `json:"Rating"`
	MarketPlace string  `json:"marketplace"`
}

type AmazonItem struct {
	ItemID      string  `json:"itemid"`
	ItemName    string  `json:"itemname"`
	Price       float32 `json:"price"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Rating      string  `json:"Rating"`
	MarketPlace string  `json:"marketplace"`
}

func (f FlipkartItem) GetItem(itemname string) (string, error) {
	rows, err := config.DB.Query("select * from items join marketplace using(itemid) where marketplace = 'flipkart';")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var items []FlipkartItem
	for rows.Next() {
		item := FlipkartItem{}
		err = rows.Scan(&item.ItemID, &item.ItemName, &item.Price, &item.Brand, &item.Description, &item.Rating, &item.MarketPlace)
		if err != nil {
			fmt.Println("flipkart item-dto", err)
			return "", err
		}
		items = append(items, item)
	}

	itemJson, _ := json.Marshal(items)

	return string(itemJson), nil
}

func (f AmazonItem) GetItem(itemname string) (string, error) {
	rows, err := config.DB.Query("select * from items join marketplace using(itemid) where marketplace = 'amazon';")
	if err != nil {
		fmt.Println("#1. amazon item-dto", err)
		return "", err
	}

	var items []FlipkartItem
	for rows.Next() {
		item := FlipkartItem{}
		err = rows.Scan(&item.ItemID, &item.ItemName, &item.Price, &item.Brand, &item.Description, &item.Rating, &item.MarketPlace)
		if err != nil {
			fmt.Println("#2. amazon item-dto", err)
			return "", err
		}
		items = append(items, item)
	}

	itemJson, _ := json.Marshal(items)

	return string(itemJson), nil
}
