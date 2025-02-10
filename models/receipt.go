package models

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price           string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

type ReceiptPoints struct {
	ID     string `json:"id"`
	Points int    `json:"points"`
}