package v1_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
	"net/http"
	"net/http/httptest"
	"smart-chat/internal/constants"
	"smart-chat/internal/dto"
	"testing"
)

func TestChatMessageController_Create(t *testing.T) {

	testCases := []struct {
		name                       string
		request                    interface{}
		urlRequested               string
		expectedHttpStatusCode     int
		FacadeChatMessageResponse  *dto.ChatMessageResponse
		FacadeChatMessageError     error
		TokenProviderRequest       string
		TokenProviderResponseError error
		expectedError              error
	}{
		{
			name: "OK",
			request: &dto.ChatMessageRequest{
				ChatID:   1,
				Question: "Hello",
			},
			TokenProviderRequest:       "XPTO",
			TokenProviderResponseError: nil,
			urlRequested:               "/smart-chat/v1/chat-message",
			expectedHttpStatusCode:     http.StatusCreated,
			FacadeChatMessageResponse: &dto.ChatMessageResponse{
				Answer: "Hi",
			},
			FacadeChatMessageError: nil,
			expectedError:          nil,
		},
		{
			name:                       "BindJsonError",
			request:                    "error_bind_json",
			TokenProviderRequest:       "XPTO",
			TokenProviderResponseError: nil,
			urlRequested:               "/smart-chat/v1/chat-message",
			expectedHttpStatusCode:     http.StatusBadRequest,
			FacadeChatMessageResponse: &dto.ChatMessageResponse{
				Answer: "Hi",
			},
			FacadeChatMessageError: errors.New("json: cannot unmarshal string into Go value of type dto.ChatMessageRequest"),
			expectedError:          errors.New("json: cannot unmarshal string into Go value of type dto.ChatMessageRequest"),
		},
		{
			name: "ErrorOnCreateChatMessage_notfound",
			request: &dto.ChatMessageRequest{
				ChatID:   1,
				Question: "Hello",
			},
			TokenProviderRequest:       "XPTO",
			TokenProviderResponseError: nil,
			urlRequested:               "/smart-chat/v1/chat-message",
			expectedHttpStatusCode:     http.StatusNotFound,
			FacadeChatMessageResponse: &dto.ChatMessageResponse{
				Answer: "Hi",
			},
			FacadeChatMessageError: constants.ErrChatNotFound,
			expectedError:          constants.ErrChatNotFound,
		},
		{
			name: "ErrorOnCreateChatMessage_InternalServerError",
			request: &dto.ChatMessageRequest{
				ChatID:   1,
				Question: "Hello",
			},
			TokenProviderRequest:       "XPTO",
			TokenProviderResponseError: nil,
			urlRequested:               "/smart-chat/v1/chat-message",
			expectedHttpStatusCode:     http.StatusInternalServerError,
			FacadeChatMessageResponse: &dto.ChatMessageResponse{
				Answer: "Hi",
			},
			FacadeChatMessageError: errors.New("internal server error"),
			expectedError:          errors.New("internal server error"),
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ChatMessageFacadeMock.On("CreateChatMessage", mock.Anything, mock.Anything).
				Return(tc.FacadeChatMessageResponse, tc.FacadeChatMessageError)

			facade.TokenProviderMock.On("ValidateToken", mock.Anything).Return(nil)

			data, err := json.Marshal(tc.request)
			assert.NoError(t, err)
			reader := bytes.NewReader(data)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, tc.urlRequested, reader)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("AuthToken", tc.TokenProviderRequest)
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedHttpStatusCode, w.Code)
			responseString := w.Body.String()

			if tc.expectedError == nil {
				getParsing := &dto.ChatMessageResponse{}
				err := json.Unmarshal([]byte(responseString), &getParsing)
				assert.NoError(t, err)
				assert.Equal(t, tc.FacadeChatMessageResponse, getParsing)
				return
			}
			errorResponse := &dto.Error{}
			err = json.Unmarshal([]byte(responseString), &errorResponse)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedError.Error(), errorResponse.Message)

		}
		t.Run(tc.name, f)
	}

}
