package service

import (
	"context"
	"smart-chat/adapter/provider"
	"smart-chat/internal/constants"
	"smart-chat/internal/model"
)

type ChatService struct {
	chatRepository chatRepository
	logger         *provider.SystemLogger
}

func NewChatService(chatRepository chatRepository, logger *provider.SystemLogger) *ChatService {
	return &ChatService{
		chatRepository: chatRepository,
		logger:         logger,
	}
}

func (c *ChatService) Create(ctx context.Context) (*model.Chat, error) {
	requestID := ctx.Value("requestID").(string)
	c.logger.NewLog("Create", "requestID", requestID).
		Debug().
		Phase("Service").
		Exe()
	return c.chatRepository.Create(ctx)
}

func (c *ChatService) GetByID(ctx context.Context, chatID uint64) (*model.Chat, error) {
	requestID := ctx.Value("requestID").(string)

	c.logger.NewLog("GetByID", "requestID", requestID,
		"ChatID", chatID).
		Debug().
		Phase("Service").
		Exe()

	found, chat, err := c.chatRepository.GetByID(ctx, chatID)
	if err != nil {
		c.logger.NewLog("GetByID: error in the get chat by id in the service", "requestID", requestID).
			Error().
			Description("error getting chat by id: " + err.Error()).
			Phase("Service").
			Exe()
		return nil, err
	}

	if !found {
		c.logger.NewLog("GetByID: chat not found", "requestID", requestID).
			Warn().
			Description("chat not found").
			Phase("Service").
			Exe()
		return nil, constants.ErrChatNotFound
	}

	c.logger.NewLog("GetByID", "requestID", requestID,
		"Chat", chat).
		Debug().
		Phase("Service").
		Exe()

	return chat, nil
}
