package models

import (
	"time"
)

// type Receipt struct {
// 	ID          uint      `gorm:"primaryKey"`
// 	ShopName    string    `json:"shop_name"`
// 	ShopAddress string    `json:"shop_address"`
// 	Items       []Item    `json:"items" gorm:"foreignKey:ReceiptID"`
// 	TotalAmount float64   `json:"total_amount"`
// 	DateOfPurchase time.Time `json:"date_of_purchase"`
// 	CreatedAt   time.Time `json:"created_at"`
// 	UpdatedAt   time.Time `json:"updated_at"`
// }

type Receipt struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Date      string    `json:"date"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Zip       string    `json:"zip"`
	Phone     string    `json:"phone"`
	Charge    string    `json:"charge"`
	Author    string    `json:"author"`
	SalesTax  string    `json:"sales_tax"`
	Total     string    `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// type Item struct {
// 	ID         uint    `gorm:"primaryKey"`
// 	ReceiptID  uint    `json:"receipt_id"`
// 	ItemName   string  `json:"item_name"`
// 	Quantity   int     `json:"quantity"`
// 	Price      float64 `json:"price"`
// 	TotalPrice float64 `json:"total_price"`
// }
