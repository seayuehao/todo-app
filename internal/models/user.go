package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"type:char(36);primaryKey"`
	Email      string    `gorm:"uniqueIndex"`
	Username   string    `gorm:"index"`
	ProviderID string    `gorm:"index"`
	Provider   string
}

func (User) TableName() string {
	return "users"
}
