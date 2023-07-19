package service

import (
	"context"
	"smart-chat/adapter/provider"
	"smart-chat/internal/model"
)

type ChatMessageService struct {
	chatMessageRepository chatMessageRepository
	logger                *provider.SystemLogger
}

func NewChatMessageService(chatMessageRepository chatMessageRepository, logger *provider.SystemLogger) *ChatMessageService {
	return &ChatMessageService{
		chatMessageRepository: chatMessageRepository,
		logger:                logger,
	}
}

func (c *ChatMessageService) Create(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
	requestID := ctx.Value("requestID").(string)
	c.logger.NewLog("Create", "requestID", requestID,
		"ChatMessage", chatMessage).
		Debug().
		Phase("Service").
		Exe()
	return c.chatMessageRepository.Create(ctx, chatMessage)
}
