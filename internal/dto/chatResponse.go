package dto

import "smart-chat/internal/model"

type ChatResponse struct {
	ChatID uint64 `json:"chat_id"`
}

func (c *ChatResponse) ParseFromChatVO(chat *model.Chat) {
	c.ChatID = chat.ID
}
