package model

import "time"

type ChatMessage struct {
	ID           int64     `json:"id"`
	UserID       string    `json:"user_id"`
	Question     string    `json:"question"`
	ResponseID   string    `json:"response_id"`
	Response     string    `json:"response"`
	QuestionDate time.Time `json:"question_date"`
	ResponseDate time.Time `json:"response_date"`
	ChatID       int64     `json:"chat_id"`
	Chat         Chat      `json:"chat"`
}
