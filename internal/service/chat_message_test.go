package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
	"go.uber.org/zap"
	mockRepository "smart-chat/internal/mocks"
	"smart-chat/internal/model"
	"testing"
	"time"
)

func TestChatMessageService_CreateService(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")

	chatMessageRepositoryMock := &mockRepository.ChatMessageRepositoryMock{}
	newChatMessageRepository := NewChatMessageService(chatMessageRepositoryMock, zap.NewExample())

	responseID := uuid.New().String()

	chatMessageRepositoryMock.On("Create", ctx, mock.Anything).
		Return(&model.ChatMessage{
			ID:           1,
			Question:     "Qual é a cor do céu?",
			ResponseID:   responseID,
			Response:     "Azul",
			QuestionDate: time.Now().Add(time.Second - 10),
			ResponseDate: time.Now(),
			ChatID:       1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}, nil)

	chatMessage := &model.ChatMessage{
		Question:     "Qual é a cor do céu?",
		ResponseID:   responseID,
		Response:     "Azul",
		QuestionDate: time.Now().Add(time.Second - 10),
		ResponseDate: time.Now(),
		ChatID:       1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	chatMessageCreated, err := newChatMessageRepository.Create(ctx, chatMessage)
	assert.NoError(t, err)
	assert.True(t, chatMessageCreated.ID > 0)
}
