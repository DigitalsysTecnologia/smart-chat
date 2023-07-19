package service

import (
	"context"
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
