package models

type Order struct {
	ID       int    `json:id gorm:"primaryKey AUTO_INCREMENT"`
	User_id  string `json:"user_id"`
	Menu_id  int    `json:"menu_id" gorm:"references:ID"`
	Optional string `json:"optional"`
	Lat      float64
	Lng      float64
}
