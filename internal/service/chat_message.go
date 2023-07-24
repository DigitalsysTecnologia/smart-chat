package service

import (
	"context"
	"go.uber.org/zap"
	"smart-chat/internal/model"
)

type ChatMessageService struct {
	chatMessageRepository chatMessageRepository
	logger                *zap.Logger
}

func NewChatMessageService(chatMessageRepository chatMessageRepository, logger *zap.Logger) *ChatMessageService {
	return &ChatMessageService{
		chatMessageRepository: chatMessageRepository,
		logger:                logger,
	}
}

func (c *ChatMessageService) Create(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
	requestID := ctx.Value("requestID").(string)
	c.logger.Debug("Create chat message",
		zap.String("requestID", requestID),
		zap.String("phase", "Service"))
	return c.chatMessageRepository.Create(ctx, chatMessage)
}
