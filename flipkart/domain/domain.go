package domain

import "errors"

var Items map[string]Item = map[string]Item{"MacBookAir": {"312431", "MacBookAir", "$800", "Apple", "Flagship business laptop for presentation and meeting purposes from Apple", "4.3", "Flipkart"}}

// Items = map["MacBookAir"]{"MacBookAir", "$800", "Apple", "Flipkart", "4.3"}

func GetItemData(item string) (*Item, error) {
	itemData := Items[item]
	if itemData.Brand == "" {
		return nil, errors.New("No item found")
	}
	return &itemData, nil
}
