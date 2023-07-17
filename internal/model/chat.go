package model

import "time"

type Chat struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Chat) TableName() string {
	return "chat"
}
