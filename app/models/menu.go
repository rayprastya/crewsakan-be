package models

type Menu struct {
	ID      int    `json:id gorm:"primaryKey AUTO_INCREMENT"`
	Name    string `json:"Name"`
	User_id int    `json:"user_id" gorm:"references:ID"`
}
