package models

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement;"`
	UserID    uuid.UUID `json:"user_id" gorm:"index"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (Todo) TableName() string {
	return "todos"
}
