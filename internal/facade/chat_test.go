package facade

import (
	"context"
	"errors"
	"github.com/tj/assert"
	"go.uber.org/zap"
	"smart-chat/internal/dto"
	serviceMock "smart-chat/internal/mocks"
	"smart-chat/internal/model"
	"testing"
	"time"
)

func TestChatFacade_CreateChat(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatServiceMock := &serviceMock.ChatServiceMock{}
	newChatService := NewChatFacade(chatServiceMock, zap.NewExample())

	chatServiceMock.On("Create", ctx).
		Return(
			&model.Chat{
				ID:        1,
				CreatedAt: time.Now().String(),
			},
			nil,
		)

	chatToCreated := &dto.ChatRequest{
		UserID: "XPTO",
	}

	chatCreated, err := newChatService.CreateChat(ctx, chatToCreated)
	assert.NoError(t, err)
	assert.True(t, chatCreated.ChatID > 0)

}

func TestChatFacade_CreateChat_InternalServerError(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatServiceMock := &serviceMock.ChatServiceMock{}
	newChatService := NewChatFacade(chatServiceMock, zap.NewExample())

	chatServiceMock.On("Create", ctx).
		Return(
			nil,
			errors.New("internal server error"),
		)

	chatToCreate := &dto.ChatRequest{
		UserID: "XPTO",
	}

	_, err := newChatService.CreateChat(ctx, chatToCreate)
	assert.Error(t, err)

}
