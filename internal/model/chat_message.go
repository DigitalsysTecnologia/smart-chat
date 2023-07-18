package model

import (
	"time"
)

type ChatMessage struct {
	ID           int64     `json:"id" gorm:"primaryKey;column:id"`
	UserID       string    `json:"user_id" gorm:"column:user_id"`
	Question     string    `json:"question" gorm:"column:question"`
	ResponseID   string    `json:"response_id" gorm:"column:response_id"`
	Response     string    `json:"response" gorm:"column:response"`
	QuestionDate time.Time `json:"question_date" gorm:"column:question_date"`
	ResponseDate time.Time `json:"response_date" gorm:"column:response_date"`
	ChatID       uint64    `json:"chat_id" gorm:"column:chat_id"`
	Chat         Chat      `json:"chat" gorm:"foreignKey:ChatID"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:updated_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (ChatMessage) TableName() string {
	return "chat_message"
}
