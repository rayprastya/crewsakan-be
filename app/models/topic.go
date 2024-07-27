package models

import (
	"time"
)

type Topic struct {
	ID        int       `json:id gorm:"primaryKey"`
	Name      string    `json:"Name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
