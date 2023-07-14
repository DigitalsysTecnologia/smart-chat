package service

import (
	"context"
	"smart-chat/internal/model"
)

type ChatService struct {
	chatRepository chatRepository
}

func NewChatService(chatRepository chatRepository) *ChatService {
	return &ChatService{
		chatRepository: chatRepository,
	}
}

func (c *ChatService) CreateChat(ctx context.Context) (*model.Chat, error) {
	return c.chatRepository.CreateChat(ctx)
}

func (c *ChatService) UpdateChat(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	return c.chatRepository.UpdateChat(ctx, chat)
}
