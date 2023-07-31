package facade

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"reflect"
	"smart-chat/internal/dto"
	layerMock "smart-chat/internal/mocks"
	"smart-chat/internal/model"
	"testing"
	"time"
)

func TestChatMessageFacade_CreateChatMessage_OK(t *testing.T) {
	ctx := context.WithValue(context.Background(), "requestID", "123")
	responseID := uuid.New().String()

	testCases := []struct {
		name                           string
		getByIDResponse                *model.Chat
		getByIDResponseError           error
		deepAiProviderResponse         *dto.Answer
		deepAiProviderResponseError    error
		createChatMessageResponse      *model.ChatMessage
		createChatMessageResponseError error
		chatMessageRequest             *dto.ChatMessageRequest
		expectedError                  error
	}{
		{
			name: "OK",
			getByIDResponse: &model.Chat{
				ID:        1,
				CreatedAt: time.Now().String(),
			},
			getByIDResponseError: nil,
			deepAiProviderResponse: &dto.Answer{
				ID:           responseID,
				Output:       "Azul",
				QuestionDate: time.Now().Add(-time.Second * 10),
				ResponseDate: time.Now(),
			},
			deepAiProviderResponseError: nil,
			createChatMessageResponse: &model.ChatMessage{
				ID:           1,
				Question:     "Qual é a cor do céu?",
				ResponseID:   responseID,
				Response:     "Azul",
				QuestionDate: time.Now().Add(time.Second - 10),
				ResponseDate: time.Now(),
				ChatID:       1,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			createChatMessageResponseError: nil,
			chatMessageRequest: &dto.ChatMessageRequest{
				Question: "Qual é a cor do céu?",
				ChatID:   1,
			},
			expectedError: nil,
		},
		{
			name:                        "GetByID returns error",
			getByIDResponse:             &model.Chat{},
			getByIDResponseError:        errors.New("chat not found"),
			deepAiProviderResponse:      &dto.Answer{},
			deepAiProviderResponseError: errors.New("internal server error"),
			createChatMessageResponse: &model.ChatMessage{
				ID:           1,
				Question:     "Qual é a cor do céu?",
				ResponseID:   responseID,
				Response:     "Azul",
				QuestionDate: time.Now().Add(time.Second - 10),
				ResponseDate: time.Now(),
				ChatID:       1,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			createChatMessageResponseError: nil,
			chatMessageRequest: &dto.ChatMessageRequest{
				Question: "Qual é a cor do céu?",
				ChatID:   1,
			},
			expectedError: errors.New("chat not found"),
		},
		{
			name: "DeepAiProvider returns error",
			getByIDResponse: &model.Chat{
				ID:        1,
				CreatedAt: time.Now().String(),
			},
			getByIDResponseError:        nil,
			deepAiProviderResponse:      &dto.Answer{},
			deepAiProviderResponseError: errors.New("internal server error"),
			createChatMessageResponse: &model.ChatMessage{
				ID:           1,
				Question:     "Qual é a cor do céu?",
				ResponseID:   responseID,
				Response:     "Azul",
				QuestionDate: time.Now().Add(time.Second - 10),
				ResponseDate: time.Now(),
				ChatID:       1,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			createChatMessageResponseError: nil,
			chatMessageRequest: &dto.ChatMessageRequest{
				Question: "Qual é a cor do céu?",
				ChatID:   1,
			},
			expectedError: errors.New("internal server error"),
		},
		{
			name: "CreateChatMessage returns error",
			getByIDResponse: &model.Chat{
				ID:        1,
				CreatedAt: time.Now().String(),
			},
			getByIDResponseError: nil,
			deepAiProviderResponse: &dto.Answer{
				ID:           responseID,
				Output:       "Azul",
				QuestionDate: time.Now().Add(-time.Second * 10),
				ResponseDate: time.Now(),
			},
			deepAiProviderResponseError:    nil,
			createChatMessageResponse:      &model.ChatMessage{},
			createChatMessageResponseError: errors.New("internal server error"),
			chatMessageRequest: &dto.ChatMessageRequest{
				Question: "Qual é a cor do céu?",
				ChatID:   1,
			},
			expectedError: errors.New("internal server error"),
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			chatMessageMock := &layerMock.ChatMessageServiceMock{}
			deepAiProviderMock := &layerMock.DeepAiProviderMock{}
			chatMock := &layerMock.ChatServiceMock{}

			newChatMessage := NewChatMessageFacade(chatMessageMock, chatMock, deepAiProviderMock, zap.NewExample())

			chatMock.On("GetByID", ctx, mock.Anything).
				Return(
					tc.getByIDResponse,
					tc.getByIDResponseError,
				)

			deepAiProviderMock.On("CallIA", ctx, mock.Anything).
				Return(
					tc.deepAiProviderResponse,
					tc.deepAiProviderResponseError,
				)

			chatMessageMock.On("Create", ctx, mock.Anything).
				Return(
					tc.createChatMessageResponse,
					tc.createChatMessageResponseError,
				)

			_, err := newChatMessage.CreateChatMessage(ctx, tc.chatMessageRequest)
			reflect.DeepEqual(err, tc.expectedError)
		}
		t.Run(tc.name, f)
	}

}
