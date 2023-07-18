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

func (c *ChatMessageService) CreateChatMessage(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
	requestID := ctx.Value("requestID").(string)
	c.logger.NewLog("CreateChatMessage", "requestID", requestID,
		"ChatMessage", chatMessage).
		Debug().
		Phase("Service").
		Exe()
	return c.chatMessageRepository.CreateChatMessage(ctx, chatMessage)
}
