package domain

import "time"

type Product struct {
    ID                uint      `gorm:"primaryKey" json:"id"`
    Name              string    `json:"name"`
    Category          string    `json:"category"`
    Quantity          int       `json:"quantity"`
    Price             int       `json:"price"`
    Description       string    `json:"description"`
    StyleNotes        string    `json:"style_notes"`
    SizeAndFit        string    `json:"size_and_fit"`
    Material          string    `json:"material"`
    Specifications    string    `json:"specifications"`
    SellerInformation string    `json:"seller_information"`
    Image             string    `json:"image"`
    CreatedAt         time.Time `json:"created_at"`
    UpdatedAt         time.Time `json:"updated_at"`
}
