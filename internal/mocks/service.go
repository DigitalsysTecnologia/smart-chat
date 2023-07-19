package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"smart-chat/internal/model"
)

type (
	ChatServiceMock struct {
		mock.Mock
	}
	ChatMessageServiceMock struct {
		mock.Mock
	}
)

func (c *ChatServiceMock) Create(ctx context.Context) (*model.Chat, error) {
	args := c.Called(ctx)

	chat := &model.Chat{}
	var err error

	if args.Get(0) != nil {
		chat = args.Get(0).(*model.Chat)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return chat, err
}

func (c *ChatServiceMock) GetByID(ctx context.Context, chatID uint64) (*model.Chat, error) {
	args := c.Called(ctx, chatID)

	chat := &model.Chat{}
	var err error

	if args.Get(0) != nil {
		chat = args.Get(0).(*model.Chat)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return chat, err
}

func (c *ChatMessageServiceMock) Create(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
	args := c.Called(ctx, chatMessage)

	chatMessageCreated := &model.ChatMessage{}
	var err error

	if args.Get(0) != nil {
		chatMessageCreated = args.Get(0).(*model.ChatMessage)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return chatMessageCreated, err
}
