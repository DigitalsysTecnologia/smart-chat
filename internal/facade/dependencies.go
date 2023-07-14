package facade

import (
	"context"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
)

type chatService interface {
	CreateChat(ctx context.Context) (*model.Chat, error)
	UpdateChat(ctx context.Context, chat *model.Chat) (*model.Chat, error)
}

type chatMessageService interface {
	CreateChatMessage(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error)
}

type deepAiProvider interface {
	CallIA(ask *dto.Ask) (*dto.Answer, error)
}
