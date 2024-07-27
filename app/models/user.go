package models

type User struct {
	ID    int    `json:id gorm:"primaryKey AUTO_INCREMENT"`
	Name  string `json:"Name" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
}
