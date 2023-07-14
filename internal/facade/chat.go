package facade

import (
	"context"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
)

type ChatFacade struct {
	chatService chatService
}

func NewChatFacade(chatService chatService) *ChatFacade {
	return &ChatFacade{
		chatService: chatService,
	}
}

func (c *ChatFacade) CreateChat(ctx context.Context) (*dto.ChatResponse, error) {

	chatCreated, err := c.chatService.CreateChat(ctx)
	if err != nil {
		return nil, err
	}

	chatResponse := &dto.ChatResponse{}

	chatResponse.ParseFromChatVO(chatCreated)

	return chatResponse, nil
}

func (c *ChatFacade) UpdateChat(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	return c.chatService.UpdateChat(ctx, chat)
}
