package facade

import (
	"context"
	"smart-chat/adapter/provider"
	"smart-chat/internal/dto"
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
	c.logger.NewLog("Create", "requestID", req).
		Debug().
		Phase("Facade").
		Exe()

	chatCreated, err := c.chatService.Create(ctx)
	if err != nil {
		c.logger.NewLog("Create: error in the create chat in the facade", "requestID", req).
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
