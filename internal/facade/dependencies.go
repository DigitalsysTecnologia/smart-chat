package facade

import (
	"context"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
)

type chatService interface {
	CreateChat(ctx context.Context) (*model.Chat, error)
	UpdateChat(ctx context.Context, chat *model.Chat) (*model.Chat, error)
	GetChatByID(ctx context.Context, chatID uint64) (*model.Chat, error)
}

type chatMessageService interface {
	CreateChatMessage(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error)
}

type deepAiProvider interface {
	CallIA(ctx context.Context, text string) (*dto.Answer, error)
}
