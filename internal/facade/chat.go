package facade

import (
	"context"
	"go.uber.org/zap"
	"smart-chat/internal/dto"
)

type ChatFacade struct {
	chatService chatService
	logger      *zap.Logger
}

func NewChatFacade(chatService chatService, logger *zap.Logger) *ChatFacade {
	return &ChatFacade{
		chatService: chatService,
		logger:      logger,
	}
}

func (c *ChatFacade) CreateChat(ctx context.Context, chatRequest *dto.ChatRequest) (*dto.ChatResponse, error) {
	chatVO := chatRequest.ParseFromChatRequest()

	req := ctx.Value("requestID").(string)
	c.logger.Debug(
		"CreateChat",
		zap.String("requestID", req),
		zap.String("phase", "Facade"))

	chatCreated, err := c.chatService.Create(ctx, chatVO)
	if err != nil {
		c.logger.Error("CreateChat: error in the create chat in the service",
			zap.String("requestID", req),
			zap.Error(err),
			zap.String("phase", "Facade"))
		return nil, err
	}

	c.logger.Debug("ChatCreated",
		zap.String("requestID", req),
		zap.String("phase", "Facade"))

	chatResponse := &dto.ChatResponse{}
	chatResponse.ParseFromChatVO(chatCreated)

	c.logger.Debug("ChatResponseCreated",
		zap.String("requestID", req),
		zap.String("phase", "Facade"))

	return chatResponse, nil
}
