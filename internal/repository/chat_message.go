package repository

import (
	"context"
	"gorm.io/gorm"
	"smart-chat/adapter/provider"
	"smart-chat/internal/model"
)

type ChatMessage struct {
	*BaseRepository
	logger *provider.SystemLogger
}

func NewChatMessageRepository(db *gorm.DB, logger *provider.SystemLogger) *ChatMessage {
	baseRepo := NewBaseRepository(db)
	return &ChatMessage{
		baseRepo,
		logger,
	}
}

func (c *ChatMessage) Create(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
	requestId := ctx.Value("requestID").(string)

	c.logger.NewLog("Create", "requestID", requestId).
		Debug().
		Phase("Repository").
		Exe()

	db, err := c.GetConnection(ctx)
	if err != nil {
		c.logger.NewLog("GetConnection: get connection in the repository", "requestID", requestId).
			Error().
			Phase("Repository").
			Exe()
		return nil, err
	}

	if err = db.Create(chatMessage).Error; err != nil {
		c.logger.NewLog("Create chat-message: error in the create chat-message in the repository", "requestID", requestId).
			Error().
			Description("error creating chat-message: " + err.Error()).
			Phase("Repository").
			Exe()
		return nil, err
	}

	c.logger.NewLog("ChatMessageCreated", "requestID", requestId,
		"Model chat-message", chatMessage).
		Debug().
		Phase("Repository").
		Exe()

	return chatMessage, nil
}
