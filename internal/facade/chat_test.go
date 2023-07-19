package facade

import (
	"context"
	"errors"
	"github.com/tj/assert"
	"smart-chat/adapter/provider"
	serviceMock "smart-chat/internal/mocks"
	"smart-chat/internal/model"
	"testing"
	"time"
)

func TestChatFacade_CreateChat(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatServiceMock := &serviceMock.ChatServiceMock{}
	newChatService := NewChatFacade(chatServiceMock, provider.NewLogger())

	chatServiceMock.On("Create", ctx).
		Return(
			&model.Chat{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			nil,
		)

	chatCreated, err := newChatService.CreateChat(ctx)
	assert.NoError(t, err)
	assert.True(t, chatCreated.ChatID > 0)

}

func TestChatFacade_CreateChat_InternalServerError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatServiceMock := &serviceMock.ChatServiceMock{}
	newChatService := NewChatFacade(chatServiceMock, provider.NewLogger())

	chatServiceMock.On("Create", ctx).
		Return(
			nil,
			errors.New("internal server error"),
		)

	_, err := newChatService.CreateChat(ctx)
	assert.Error(t, err)

}
