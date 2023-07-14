package service

import (
	"context"
	"smart-chat/internal/model"
)

type ChatMessageService struct {
	chatMessageRepository chatMessageRepository
}

func NewChatMessageService(chatMessageRepository chatMessageRepository) *ChatMessageService {
	return &ChatMessageService{
		chatMessageRepository: chatMessageRepository,
	}
}

func (c *ChatMessageService) CreateChatMessage(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
	return c.chatMessageRepository.CreateChatMessage(ctx, chatMessage)
}
