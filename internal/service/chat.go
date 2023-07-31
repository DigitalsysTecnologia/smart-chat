package service

import (
	"context"
	"go.uber.org/zap"
	"smart-chat/internal/constants"
	"smart-chat/internal/model"
	"time"
)

type ChatService struct {
	chatRepository chatRepository
	logger         *zap.Logger
}

func NewChatService(chatRepository chatRepository, logger *zap.Logger) *ChatService {
	return &ChatService{
		chatRepository: chatRepository,
		logger:         logger,
	}
}

func (c *ChatService) Create(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	chat.CreatedAt = time.Now().UTC().String()
	chat.UpdatedAt = time.Now().UTC().String()

	requestID := ctx.Value("requestID").(string)
	c.logger.Debug("Create chat",
		zap.String("requestID", requestID),
		zap.String("phase", "Service"))

	return c.chatRepository.Create(ctx, chat)
}

func (c *ChatService) GetByID(ctx context.Context, chatID uint64) (*model.Chat, error) {
	requestID := ctx.Value("requestID").(string)

	c.logger.Debug("GetByID chat",
		zap.String("requestID", requestID),
		zap.String("phase", "Service"))

	found, chat, err := c.chatRepository.GetByID(ctx, chatID)
	if err != nil {
		c.logger.Error("GetByID: error in the get chat in the repository",
			zap.String("requestID", requestID),
			zap.String("phase", "Service"))
		return nil, err
	}

	if !found {
		c.logger.Warn("GetByID: chat not found",
			zap.String("requestID", requestID),
			zap.String("phase", "Service"))
		return nil, constants.ErrChatNotFound
	}

	c.logger.Debug("Chat found",
		zap.String("requestID", requestID),
		zap.String("phase", "Service"),
		zap.Any("chat", chat))

	return chat, nil
}
