package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"smart-chat/internal/model"
)

type (
	ChatRepositoryMock struct {
		mock.Mock
	}
	ChatMessageRepositoryMock struct {
		mock.Mock
	}
)

func (c *ChatRepositoryMock) Create(ctx context.Context, chat *model.Chat) (*model.Chat, error) {
	args := c.Called(ctx, chat)

	chatCreated := &model.Chat{}
	var err error

	if args.Get(0) != nil {
		chatCreated = args.Get(0).(*model.Chat)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(error)
	}

	return chatCreated, err
}

func (c *ChatRepositoryMock) GetByID(ctx context.Context, chatID uint64) (bool, *model.Chat, error) {
	args := c.Called(ctx, chatID)

	chat := &model.Chat{}
	var err error

	if args.Get(1) != nil {
		chat = args.Get(1).(*model.Chat)
	}

	if args.Get(2) != nil {
		err = args.Get(2).(error)
	}

	return args.Bool(0), chat, err
}

func (c *ChatMessageRepositoryMock) Create(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error) {
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
