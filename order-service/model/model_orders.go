package model

type Order struct {
	OrderID     int    `json:"order_id" gorm:"primaryKey;autoIncrement"`
	UserID      int    `json:"user_id" gorm:"not null"`
	ProductName string `json:"product_name" gorm:"not null"`
}
