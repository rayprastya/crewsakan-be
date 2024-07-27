package models

type Menu struct {
	ID          int     `json:id gorm:"primaryKey AUTO_INCREMENT"`
	Name        string  `json:"Name"`
	Merchant_id int     `json:"merchant_id" gorm:"references:ID"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
