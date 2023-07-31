package v1_test

import (
	"go.uber.org/zap"
	"smart-chat/adapter/rest"
	v1 "smart-chat/internal/controller/v1"
	"smart-chat/internal/middleware"
	facadeMocks "smart-chat/internal/mocks"
	providerMocks "smart-chat/internal/mocks"
	"smart-chat/internal/model"
	"testing"
)

type Facade struct {
	ChatFacadeMock        *facadeMocks.ChatFacadeMock
	ChatMessageFacadeMock *facadeMocks.ChatMessageFacadeMock
	TokenProviderMock     *providerMocks.TokenProviderMock
}

func setupTestRouter(t *testing.T) (*rest.ServerRest, Facade) {
	t.Helper()

	facades := Facade{
		ChatFacadeMock:        &facadeMocks.ChatFacadeMock{},
		ChatMessageFacadeMock: &facadeMocks.ChatMessageFacadeMock{},
		TokenProviderMock:     &providerMocks.TokenProviderMock{},
	}

	cfg := &model.Config{}

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			ChatController:        v1.NewChatController(facades.ChatFacadeMock, zap.NewExample()),
			ChatMessageController: v1.NewChatMessageController(facades.ChatMessageFacadeMock, zap.NewExample()),
			HeathCheckController:  v1.NewHealthCheckController(),
		},
		&rest.Middlewares{
			LoggerMiddleware: middleware.NewLoggerMiddleware(facades.TokenProviderMock),
		},
	)

	return serverRest, facades

}
