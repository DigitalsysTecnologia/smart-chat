package facade

import (
	"context"
	"smart-chat/adapter/provider"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
)

type ChatFacade struct {
	chatService chatService
	logger      *provider.SystemLogger
}

func NewChatFacade(chatService chatService, logger *provider.SystemLogger) *ChatFacade {
	return &ChatFacade{
		chatService: chatService,
		logger:      logger,
	}
}

func (c *ChatFacade) CreateChat(ctx context.Context) (*dto.ChatResponse, error) {
	req := ctx.Value("requestID").(string)
	c.logger.NewLog("CreateChat", "requestID", req).
		Debug().
		Phase("Facade").
		Exe()

	chatCreated, err := c.chatService.CreateChat(ctx)
	if err != nil {
		c.logger.NewLog("CreateChat: error in the create chat in the facade", "requestID", req).
			Error().
			Description("error creating chat: " + err.Error()).
			Phase("Facade").
			Exe()
		return nil, err
	}

	c.logger.NewLog("Chat was created", "requestID", req,
		"Chat", chatCreated).
		Debug().
		Phase("Facade").
		Exe()

	chatResponse := &dto.ChatResponse{}
	chatResponse.ParseFromChatVO(chatCreated)

	c.logger.NewLog("ChatResponse was created", "requestID", req,
		"ChatResponse", chatResponse).
		Debug().
		Phase("Facade").
		Exe()

	return chatResponse, nil
}

func (c *ChatFacade) UpdateChat(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	return c.chatService.UpdateChat(ctx, chat)
}
