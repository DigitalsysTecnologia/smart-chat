package service

import (
	"context"
	"smart-chat/internal/model"
)

type chatRepository interface {
	CreateChat(ctx context.Context) (*model.Chat, error)
	UpdateChat(ctx context.Context, chat *model.Chat) (*model.Chat, error)
}

type chatMessageRepository interface {
	CreateChatMessage(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error)
}
