package models

type User struct {
	ID     int    `json:"id"  gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
	Phone  string `json:"phone"`
}
