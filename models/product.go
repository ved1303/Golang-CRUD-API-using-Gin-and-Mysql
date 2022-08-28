package models

type Product struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
