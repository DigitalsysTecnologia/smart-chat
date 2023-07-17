package dto

import (
	"smart-chat/internal/model"
	"time"
)

type ChatMessageRequest struct {
	UserID       string    `json:"user_id"`
	ChatID       uint64    `json:"chat_id"`
	Question     string    `json:"question"`
	ResponseID   string    `json:"response_id"`
	Response     string    `json:"response"`
	QuestionDate time.Time `json:"question_date"`
	ResponseDate time.Time `json:"response_date"`
}

func (c *ChatMessageRequest) ParseFromChatMessageRequestAndAnswer() *model.ChatMessage {

	return &model.ChatMessage{
		UserID:       c.UserID,
		ChatID:       c.ChatID,
		Question:     c.Question,
		ResponseID:   c.ResponseID,
		Response:     c.Response,
		QuestionDate: time.Now(),
		ResponseDate: time.Now(),
	}
}
