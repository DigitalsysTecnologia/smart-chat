package service

import (
	"context"
	"smart-chat/internal/model"
)

type chatRepository interface {
	Create(ctx context.Context) (*model.Chat, error)
	GetByID(ctx context.Context, chatID uint64) (bool, *model.Chat, error)
}

type chatMessageRepository interface {
	Create(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error)
}
