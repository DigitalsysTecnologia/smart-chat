package facade

import (
	"context"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
)

type chatService interface {
	Create(ctx context.Context) (*model.Chat, error)
	GetByID(ctx context.Context, chatID uint64) (*model.Chat, error)
}

type chatMessageService interface {
	Create(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error)
}

type deepAiProvider interface {
	CallIA(ctx context.Context, text string) (*dto.Answer, error)
}
