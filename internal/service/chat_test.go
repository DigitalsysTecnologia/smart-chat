package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
	"smart-chat/adapter/provider"
	mockRepository "smart-chat/internal/mocks"

	"smart-chat/internal/model"
	"testing"
	"time"
)

func TestChatService_CreateChat(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatRepositoryMock := &mockRepository.ChatRepositoryMock{}
	newChatRepository := NewChatService(chatRepositoryMock, provider.NewLogger())

	chatRepositoryMock.On("Create", ctx).
		Return(&model.Chat{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil)

	chat, err := newChatRepository.Create(ctx)
	assert.NoError(t, err)
	assert.True(t, chat.ID > 0)
}

func TestChatService_GetByID(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatRepositoryMock := &mockRepository.ChatRepositoryMock{}
	newChatService := NewChatService(chatRepositoryMock, provider.NewLogger())

	chatRepositoryMock.On("GetByID", ctx, mock.Anything).
		Return(
			true,
			&model.Chat{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	chat, err := newChatService.GetByID(ctx, 1)
	assert.NoError(t, err)
	assert.True(t, chat.ID > 0)

}

func TestChatService_GetByID_NotFound(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatRepositoryMock := &mockRepository.ChatRepositoryMock{}
	newChatService := NewChatService(chatRepositoryMock, provider.NewLogger())

	chatRepositoryMock.On("GetByID", ctx, mock.Anything).
		Return(
			false,
			&model.Chat{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}, nil)

	_, err := newChatService.GetByID(ctx, 1)
	assert.Error(t, err)

}

func TestChatService_GetByID_InternalServerError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatRepositoryMock := &mockRepository.ChatRepositoryMock{}
	newChatService := NewChatService(chatRepositoryMock, provider.NewLogger())

	chatRepositoryMock.On("GetByID", ctx, mock.Anything).
		Return(
			true,
			nil,
			errors.New("internal server error"))

	_, err := newChatService.GetByID(ctx, 1)
	assert.Error(t, err)

}
