package repository

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smart-chat/internal/model"
)

type ChatMessage struct {
	*BaseRepository
	logger *zap.Logger
}

func NewChatMessageRepository(db *gorm.DB, logger *zap.Logger) *ChatMessage {
	baseRepo := NewBaseRepository(db)
	return &ChatMessage{
		baseRepo,
		logger,
	}
}

func (c *ChatMessage) Create(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
	requestId := ctx.Value("requestID").(string)

	c.logger.Debug("Create chat-message",
		zap.String("requestID", requestId),
		zap.String("phase", "Repository"))

	db, err := c.GetConnection(ctx)
	if err != nil {
		c.logger.Error("GetConnection: get connection in the repository",
			zap.String("requestID", requestId),
			zap.Error(err),
			zap.String("phase", "Repository"))
		return nil, err
	}

	if err = db.Create(chatMessage).Error; err != nil {
		c.logger.Error("Create chat-message: error in the create chat-message in the repository",
			zap.String("requestID", requestId),
			zap.Error(err),
			zap.String("phase", "Repository"))
		return nil, err
	}

	c.logger.Debug("Chat-message created",
		zap.String("requestID", requestId),
		zap.String("phase", "Repository"),
		zap.Any("chat-message", chatMessage))

	return chatMessage, nil
}
