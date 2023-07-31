package dto

import "smart-chat/internal/model"

type ChatRequest struct {
	UserID string `json:"user_id"`
}

func (c *ChatRequest) ParseFromChatRequest() *model.Chat {
	return &model.Chat{
		UserID: c.UserID,
	}
}
