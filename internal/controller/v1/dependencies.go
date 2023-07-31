package v1

import (
	"context"
	"smart-chat/internal/dto"
)

type chatFacade interface {
	CreateChat(ctx context.Context, request *dto.ChatRequest) (*dto.ChatResponse, error)
}

type chatMessageFacade interface {
	CreateChatMessage(ctx context.Context, chatMessageRequest *dto.ChatMessageRequest) (*dto.ChatMessageResponse, error)
}
