package domain

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name             string `json:"name"`
    Category         string `json:"category"`
    Quantity         int    `json:"quantity"`
    Price            int    `json:"price"`
    Description      string `json:"description"`
    StyleNotes       string `json:"style_notes"`
    SizeAndFit       string `json:"size_and_fit"`
    Material         string `json:"material"`
    Specifications   string `json:"specifications"`
    SellerInformation string `json:"seller_information"`
    Image            string `json:"image"`
}
