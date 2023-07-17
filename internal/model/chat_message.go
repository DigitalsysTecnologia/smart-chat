package model

import (
	"time"
)

type ChatMessage struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	UserID       string    `json:"user_id" gorm:"type:varchar(191)"`
	Question     string    `json:"question" gorm:"type:varchar(300)"`
	ResponseID   string    `json:"response_id" gorm:"type:varchar(191)"`
	Response     string    `json:"response" gorm:"type:varchar(300)"`
	QuestionDate time.Time `json:"question_date" gorm:"type:datetime"`
	ResponseDate time.Time `json:"response_date" gorm:"type:datetime"`
	ChatID       uint64    `json:"chat_id"`
	Chat         Chat      `json:"chat" gorm:"foreignKey:ChatID"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (ChatMessage) TableName() string {
	return "chat_message"
}
