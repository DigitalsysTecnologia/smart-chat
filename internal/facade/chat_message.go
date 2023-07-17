package facade

import (
	"context"
	"smart-chat/internal/dto"
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
	chatMessageRequest.ResponseID = answer.ID
	chatMessageRequest.Response = answer.Output

	chatMessage := chatMessageRequest.ParseFromChatMessageRequestAndAnswer()

	_, err = c.chatMessageService.CreateChatMessage(ctx, chatMessage)
	if err != nil {
		return nil, err
	}

	chatMessageResponse := &dto.ChatMessageResponse{}
	chatMessageResponse.ParseFromChatMessageResponse(answer)

	return chatMessageResponse, nil
}
