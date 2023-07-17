package dto

import (
	"smart-chat/internal/model"
	"time"
)

type ChatMessageRequest struct {
	UserID   string `json:"user_id"`
	ChatID   uint64 `json:"chat_id"`
	Question string `json:"question"`
}

func (c *ChatMessageRequest) ParseFromChatMessageRequestAndAnswer() *model.ChatMessage {

	return &model.ChatMessage{
		UserID:       c.UserID,
		ChatID:       c.ChatID,
		Question:     c.Question,
		QuestionDate: time.Now(),
		ResponseDate: time.Now(),
	}
}
