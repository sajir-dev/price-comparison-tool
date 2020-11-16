package domain

// Item struct defines an item
type Item struct {
	ItemID      string `json:"itemid"`
	ItemName    string `json:"item"`
	Price       string `json:"price"`
	Brand       string `json:"brand"`
	Platform    string `json:"platform"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
}
