package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"smart-chat/internal/dto"
)

type (
	ChatFacadeMock struct {
		mock.Mock
	}
	ChatMessageFacadeMock struct {
		mock.Mock
	}
)

func (a *ChatFacadeMock) CreateChat(ctx context.Context, request *dto.ChatRequest) (*dto.ChatResponse, error) {
	args := a.Called(ctx, request)

	chatResponse := &dto.ChatResponse{}
	var err error

	if args.Get(0) != nil {
		chatResponse = args.Get(0).(*dto.ChatResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return chatResponse, err
}

func (c *ChatMessageFacadeMock) CreateChatMessage(ctx context.Context, chatMessageRequest *dto.ChatMessageRequest) (*dto.ChatMessageResponse, error) {
	args := c.Called(ctx, chatMessageRequest)

	chatMessageResponse := &dto.ChatMessageResponse{}
	var err error

	if args.Get(0) != nil {
		chatMessageResponse = args.Get(0).(*dto.ChatMessageResponse)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return chatMessageResponse, err
}
