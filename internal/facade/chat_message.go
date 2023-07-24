package facade

import (
	"context"
	"go.uber.org/zap"
	"smart-chat/internal/dto"
)

type chatMessageFacade struct {
	chatMessageService chatMessageService
	chatService        chatService
	deepAiProvider     deepAiProvider
	logger             *zap.Logger
}

func NewChatMessageFacade(chatMessageService chatMessageService, chatService chatService, deepAiProvider deepAiProvider, logger *zap.Logger) *chatMessageFacade {
	return &chatMessageFacade{
		chatMessageService: chatMessageService,
		chatService:        chatService,
		deepAiProvider:     deepAiProvider,
		logger:             logger,
	}
}

func (c *chatMessageFacade) CreateChatMessage(ctx context.Context, chatMessageRequest *dto.ChatMessageRequest) (*dto.ChatMessageResponse, error) {
	requestID := ctx.Value("requestID").(string)

	c.logger.Debug("CreateChatMessage",
		zap.String("requestID", requestID),
		zap.String("phase", "Facade"))

	_, err := c.chatService.GetByID(ctx, chatMessageRequest.ChatID)
	if err != nil {
		c.logger.Error("CreateChatMessage: error in the get chat in the service",
			zap.String("requestID", requestID),
			zap.Error(err),
			zap.String("phase", "Facade"))
		return nil, err
	}

	c.logger.Debug("Chat was found",
		zap.String("requestID", requestID),
		zap.String("phase", "Facade"))

	answer, err := c.deepAiProvider.CallIA(ctx, chatMessageRequest.Question)
	if err != nil {
		c.logger.Error("CreateChatMessage: error in the call IA in the provider",
			zap.String("requestID", requestID),
			zap.Error(err),
			zap.String("phase", "Facade"))
		return nil, err
	}

	chatMessage := chatMessageRequest.ParseFromChatMessageRequestAndAnswer()

	chatMessage.ResponseID = answer.ID
	chatMessage.Response = answer.Output
	chatMessage.QuestionDate = answer.QuestionDate
	chatMessage.ResponseDate = answer.ResponseDate

	_, err = c.chatMessageService.Create(ctx, chatMessage)
	if err != nil {
		c.logger.Error("CreateChatMessage: error in the create chat message in the service",
			zap.String("requestID", requestID),
			zap.Error(err),
			zap.String("phase", "Facade"))
		return nil, err
	}

	c.logger.Debug("ChatMessage was created",
		zap.String("requestID", requestID),
		zap.String("phase", "Facade"))

	chatMessageResponse := &dto.ChatMessageResponse{}
	chatMessageResponse.ParseFromChatMessageResponse(answer)

	c.logger.Debug("ChatMessageResponse was created",
		zap.String("requestID", requestID),
		zap.String("phase", "Facade"))

	return chatMessageResponse, nil
}
