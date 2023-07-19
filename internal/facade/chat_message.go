package facade

import (
	"context"
	"smart-chat/adapter/provider"
	"smart-chat/internal/dto"
)

type chatMessageFacade struct {
	chatMessageService chatMessageService
	chatService        chatService
	deepAiProvider     deepAiProvider
	logger             *provider.SystemLogger
}

func NewChatMessageFacade(chatMessageService chatMessageService, chatService chatService, deepAiProvider deepAiProvider, logger *provider.SystemLogger) *chatMessageFacade {
	return &chatMessageFacade{
		chatMessageService: chatMessageService,
		chatService:        chatService,
		deepAiProvider:     deepAiProvider,
		logger:             logger,
	}
}

func (c *chatMessageFacade) CreateChatMessage(ctx context.Context, chatMessageRequest *dto.ChatMessageRequest) (*dto.ChatMessageResponse, error) {
	requestID := ctx.Value("requestID").(string)

	c.logger.NewLog("Create", "requestID", requestID,
		"ChatMessageRequest", chatMessageRequest).
		Debug().
		Phase("Facade").
		Exe()

	_, err := c.chatService.GetByID(ctx, chatMessageRequest.ChatID)
	if err != nil {
		c.logger.NewLog("Create: error in the get chat by id in the facade", "requestID", requestID).
			Error().
			Description("error getting chat: " + err.Error()).
			Phase("Facade").
			Exe()
		return nil, err
	}

	c.logger.NewLog("Chat was found", "requestID", requestID).
		Debug().
		Phase("Facade").
		Exe()

	answer, err := c.deepAiProvider.CallIA(ctx, chatMessageRequest.Question)
	if err != nil {
		c.logger.NewLog("Create: error in the call IA in the facade", "requestID", requestID).
			Error().
			Description("error calling IA: " + err.Error()).
			Phase("Facade").
			Exe()
		return nil, err
	}

	chatMessage := chatMessageRequest.ParseFromChatMessageRequestAndAnswer()

	chatMessage.ResponseID = answer.ID
	chatMessage.Response = answer.Output
	chatMessage.QuestionDate = answer.QuestionDate
	chatMessage.ResponseDate = answer.ResponseDate

	c.logger.NewLog("ChatMessage before creating", "requestID", requestID,
		"ChatMessage", chatMessage).
		Debug().
		Phase("Facade").
		Exe()

	_, err = c.chatMessageService.Create(ctx, chatMessage)
	if err != nil {
		c.logger.NewLog("Create: error in the create chat message in the facade", "requestID", requestID).
			Error().
			Description("error creating chat message: " + err.Error()).
			Phase("Facade").
			Exe()
		return nil, err
	}

	c.logger.NewLog("ChatMessage was created", "requestID", requestID,
		"ChatMessage", chatMessage).
		Debug().
		Phase("Facade").
		Exe()

	chatMessageResponse := &dto.ChatMessageResponse{}
	chatMessageResponse.ParseFromChatMessageResponse(answer)

	c.logger.NewLog("ChatMessageResponse was created", "requestID", requestID,
		"ChatMessageResponse", chatMessageResponse).
		Debug().
		Phase("Facade").
		Exe()

	return chatMessageResponse, nil
}
