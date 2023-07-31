package v1_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"smart-chat/internal/dto"
	"testing"
)

func TestChatController_Create(t *testing.T) {

	testCases := []struct {
		name                       string
		urlRequested               string
		expectedHttpStatusCode     int
		TokenProviderRequest       string
		TokenProviderResponseError error
		Request                    *dto.ChatRequest
		FacadeChatResponse         *dto.ChatResponse
		FacadeChatError            error
		expectedError              error
	}{
		{
			name:                       "OK",
			urlRequested:               "/smart-chat/v1/chat",
			expectedHttpStatusCode:     http.StatusCreated,
			TokenProviderRequest:       "XPTO",
			TokenProviderResponseError: nil,
			Request: &dto.ChatRequest{
				UserID: "XPTO",
			},
			FacadeChatResponse: &dto.ChatResponse{
				ChatID: 1,
			},
			FacadeChatError: nil,
			expectedError:   nil,
		},
		{
			name:                       "FacadeChatResponseError",
			urlRequested:               "/smart-chat/v1/chat",
			expectedHttpStatusCode:     http.StatusInternalServerError,
			TokenProviderRequest:       "XPTO",
			TokenProviderResponseError: nil,
			Request: &dto.ChatRequest{
				UserID: "XPTO",
			},
			FacadeChatResponse: &dto.ChatResponse{},
			FacadeChatError:    errors.New("internal server error"),
			expectedError:      errors.New("internal server error"),
		},
	}

	for _, tc := range testCases {
		f := func(t *testing.T) {
			server, facade := setupTestRouter(t)
			router := server.Engine

			facade.ChatFacadeMock.On("CreateChat", mock.Anything, mock.Anything).
				Return(tc.FacadeChatResponse, tc.FacadeChatError)

			facade.TokenProviderMock.On("ValidateToken", mock.Anything).Return(nil)

			data, err := json.Marshal(tc.Request)
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
				getParsing := &dto.ChatResponse{}
				err := json.Unmarshal([]byte(responseString), &getParsing.ChatID)
				assert.NoError(t, err)
				assert.Equal(t, tc.FacadeChatResponse, getParsing)
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
