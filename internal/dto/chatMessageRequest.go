package dto

import (
	"smart-chat/internal/model"
	"time"
)

type ChatMessageRequest struct {
	ChatID   uint64 `json:"chat_id"`
	Question string `json:"question"`
}

func (c *ChatMessageRequest) ParseFromChatMessageRequestAndAnswer() *model.ChatMessage {

	return &model.ChatMessage{
		ChatID:       c.ChatID,
		Question:     c.Question,
		QuestionDate: time.Now(),
		ResponseDate: time.Now(),
	}
}
