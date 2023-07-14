package repository

import (
	"context"
	"gorm.io/gorm"
	"smart-chat/internal/model"
)

type chatRepository struct {
	*BaseRepository
}

func NewChatRepository(db *gorm.DB) *chatRepository {
	baseRepo := NewBaseRepository(db)
	return &chatRepository{
		baseRepo,
	}
}

func (c *chatRepository) CreateChat(ctx context.Context) (*model.Chat, error) {
	db, err := c.GetConnection(ctx)
	if err != nil {
		return nil, err
	}

	chat := &model.Chat{}

	if err = db.Create(chat).Error; err != nil {
		return nil, err
	}

	return chat, nil
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
