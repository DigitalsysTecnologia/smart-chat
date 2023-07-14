package facade

import (
	"context"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
)

type chatMessageFacade struct {
	chatMessageService chatMessageService
	deepAiProvider     deepAiProvider
}

func NewChatMessageFacade(chatMessageService chatMessageService) *chatMessageFacade {
	return &chatMessageFacade{
		chatMessageService: chatMessageService,
	}
}

func (c *chatMessageFacade) CreateChatMessage(ctx context.Context, chatMessageRequest *dto.ChatMessageRequest) (*dto.ChatMessageResponse, error) {
	ask := &dto.Ask{}

	ask.ParseFromChatMessageRequest(chatMessageRequest)

	answer, err := c.deepAiProvider.CallIA(ask)
	if err != nil {
		return nil, err
	}

	chatMessage := &model.ChatMessage{}
	chatMessage.ParseFromChatMessageRequestAndAnswer(chatMessageRequest, answer)

	_, err = c.chatMessageService.CreateChatMessage(ctx, chatMessage)
	if err != nil {
		return nil, err
	}

	chatMessageResponse := &dto.ChatMessageResponse{}
	chatMessageResponse.ParseFromChatMessageResponse(answer)

	return chatMessageResponse, nil
}
