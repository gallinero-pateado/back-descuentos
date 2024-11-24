package models

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Category      string `json:"category"`
	Description   string `json:"description"`
	Price         string `json:"price"`
	PreviousPrice string `json:"previous_price,omitempty"`
	Image         string `json:"image"`
	Logo          string `json:"logo"`
	Type          string `json:"type"`
	Url           string `json:"url"`
}
