package domain

// Item struct defines an item
type Item struct {
	ItemID      string `json:"itemid"`
	ItemName    string `json:"item"`
	Price       string `json:"price"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	Platform    string `json:"platform"`
}

type ItemInterface interface {
	GetItems(string) (string, error)
}

func GetItems(item ItemInterface, itemname string) (string, error) {
	itemJson, err := item.GetItems(itemname)
	if err != nil {
		return "", err
	}

	return itemJson, nil
}
