package repository

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smart-chat/internal/model"
)

type chatRepository struct {
	*BaseRepository
	logger *zap.Logger
}

func NewChatRepository(db *gorm.DB, logger *zap.Logger) *chatRepository {
	baseRepo := NewBaseRepository(db)
	return &chatRepository{
		baseRepo,
		logger,
	}
}

func (c *chatRepository) Create(ctx context.Context) (*model.Chat, error) {
	requestId := ctx.Value("requestID").(string)

	c.logger.Debug("Create chat",
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

	chat := &model.Chat{}
	if err = db.Create(chat).Error; err != nil {
		c.logger.Error("Create chat: error in the create chat in the repository",
			zap.String("requestID", requestId),
			zap.Error(err),
			zap.String("phase", "Repository"))
		return nil, err
	}

	c.logger.Debug("Chat created",
		zap.String("requestID", requestId),
		zap.String("phase", "Repository"),
		zap.Any("chat", chat))

	return chat, nil
}

func (c *chatRepository) GetByID(ctx context.Context, chatID uint64) (bool, *model.Chat, error) {
	requestId := ctx.Value("requestID").(string)

	c.logger.Debug("GetByID chat",
		zap.String("requestID", requestId),
		zap.String("phase", "Repository"))
	db, err := c.GetConnection(ctx)
	if err != nil {
		c.logger.Error("GetConnection: get connection in the repository",
			zap.String("requestID", requestId),
			zap.Error(err),
			zap.String("phase", "Repository"))
		return false, nil, err
	}

	chat := &model.Chat{}

	if err = db.Model(&model.Chat{}).Where(&model.Chat{ID: chatID}).First(chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.logger.Error("GetByID chat: chat not found",
				zap.String("requestID", requestId),
				zap.Error(err),
				zap.String("phase", "Repository"))
			return false, nil, nil
		}
		c.logger.Error("GetByID chat: error in the get chat in the repository",
			zap.String("requestID", requestId),
			zap.Error(err),
			zap.String("phase", "Repository"))
		return false, nil, err
	}
	c.logger.Debug("Chat found",
		zap.String("requestID", requestId),
		zap.String("phase", "Repository"),
		zap.Any("chat", chat))
	return true, chat, nil
}
