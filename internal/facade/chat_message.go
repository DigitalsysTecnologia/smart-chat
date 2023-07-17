package facade

import (
	"context"
	"smart-chat/internal/dto"
)

type chatMessageFacade struct {
	chatMessageService chatMessageService
	deepAiProvider     deepAiProvider
}

func NewChatMessageFacade(chatMessageService chatMessageService, deepAiProvider deepAiProvider) *chatMessageFacade {
	return &chatMessageFacade{
		chatMessageService: chatMessageService,
		deepAiProvider:     deepAiProvider,
	}
}

func (c *chatMessageFacade) CreateChatMessage(ctx context.Context, chatMessageRequest *dto.ChatMessageRequest) (*dto.ChatMessageResponse, error) {
	ask := &dto.Ask{}

	ask.ParseFromChatMessageRequest(chatMessageRequest)

	answer, err := c.deepAiProvider.CallIA(ask)
	if err != nil {
		return nil, err
	}

	chatMessage := chatMessageRequest.ParseFromChatMessageRequestAndAnswer()

	chatMessage.ResponseID = answer.ID
	chatMessage.Response = answer.Output

	_, err = c.chatMessageService.CreateChatMessage(ctx, chatMessage)
	if err != nil {
		return nil, err
	}

	chatMessageResponse := &dto.ChatMessageResponse{}
	chatMessageResponse.ParseFromChatMessageResponse(answer)

	return chatMessageResponse, nil
}
