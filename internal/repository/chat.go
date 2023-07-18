package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"smart-chat/adapter/provider"
	"smart-chat/internal/model"
)

type chatRepository struct {
	*BaseRepository
	logger *provider.SystemLogger
}

func NewChatRepository(db *gorm.DB, logger *provider.SystemLogger) *chatRepository {
	baseRepo := NewBaseRepository(db)
	return &chatRepository{
		baseRepo,
		logger,
	}
}

func (c *chatRepository) CreateChat(ctx context.Context) (*model.Chat, error) {
	requestId := ctx.Value("requestID").(string)

	c.logger.NewLog("CreateChat", "requestID", requestId).
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

	chat := &model.Chat{}
	if err = db.Create(chat).Error; err != nil {
		c.logger.NewLog("Create chat: error in the create chat in the repository", "requestID", requestId).
			Error().
			Description("error creating chat").
			Phase("Repository").
			Exe()
		return nil, err
	}

	c.logger.NewLog("CreateChat", "requestID", requestId,
		"Model chat", chat).
		Debug().
		Phase("Repository").
		Exe()

	return chat, nil
}

func (c *chatRepository) GetChatByID(ctx context.Context, chatID uint64) (bool, *model.Chat, error) {
	requestId := ctx.Value("requestID").(string)

	c.logger.NewLog("GetChatByID", "requestID", requestId).
		Debug().
		Phase("Repository").
		Exe()
	db, err := c.GetConnection(ctx)
	if err != nil {
		c.logger.NewLog("GetConnection: get connection in the repository", "requestID", requestId).
			Error().
			Phase("Repository").
			Exe()
		return false, nil, err
	}

	chat := &model.Chat{}

	if err = db.Where(&model.Chat{
		ID: chatID,
	}).First(chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.logger.NewLog("Chat not found: error in the found chat in the repository", "requestID", requestId).
				Warn().
				Description("error chat not found").
				Phase("Repository").
				Exe()
			return false, nil, nil
		}
		c.logger.NewLog("Error on get chat: error in the get chat in the repository", "requestID", requestId).
			Error().
			Description("error chat on get chat").
			Phase("Repository").
			Exe()
		return false, nil, err
	}
	c.logger.NewLog("Chat was found", "requestID", requestId,
		"Model chat", chat).
		Debug().
		Phase("Repository").
		Exe()
	return true, chat, nil
}

func (c *chatRepository) UpdateChat(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	db, err := c.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	if err = db.Save(chat).Error; err != nil {
		return nil, err
	}

	return chat, nil
}
