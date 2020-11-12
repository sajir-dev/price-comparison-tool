package domain

// Item struct defines an item
type Item struct {
	ItemName    string `json:"item"`
	Price       string `json:"price"`
	Brand       string `json:"brand"`
	Platform    string `json:"platform"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
}
