package repository

import (
	"context"
	"gorm.io/gorm"
	"smart-chat/internal/model"
)

type ChatMessage struct {
	*BaseRepository
}

func NewChatMessageRepository(db *gorm.DB) *ChatMessage {
	baseRepo := NewBaseRepository(db)
	return &ChatMessage{
		baseRepo,
	}
}

func (c *ChatMessage) CreateChatMessage(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
	db, err := c.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = db.Create(chatMessage).Error; err != nil {
		return nil, err
	}

	return chatMessage, nil
}
