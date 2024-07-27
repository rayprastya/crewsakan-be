package models

type Wishlist struct {
	ID          int    `json:id gorm:"primaryKey AUTO_INCREMENT"`
	Merchant_id int    `json:"merchant_id" gorm:"references:ID"`
	name        string `json:"name"`
	recipes     string `json:"recipes"`
	steps       string `json:"steps"`
}
