package v1_test

import (
	"smart-chat/adapter/provider"
	"smart-chat/adapter/rest"
	v1 "smart-chat/internal/controller/v1"
	"smart-chat/internal/middleware"
	facadeMocks "smart-chat/internal/mocks"
	"smart-chat/internal/model"
	"testing"
)

type Facade struct {
	ChatFacadeMock        *facadeMocks.ChatFacadeMock
	ChatMessageFacadeMock *facadeMocks.ChatMessageFacadeMock
}

func setupTestRouter(t *testing.T) (*rest.ServerRest, Facade) {
	t.Helper()

	facades := Facade{
		ChatFacadeMock:        &facadeMocks.ChatFacadeMock{},
		ChatMessageFacadeMock: &facadeMocks.ChatMessageFacadeMock{},
	}

	cfg := &model.Config{}

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			ChatController:        v1.NewChatController(facades.ChatFacadeMock, provider.NewLogger()),
			ChatMessageController: v1.NewChatMessageController(facades.ChatMessageFacadeMock, provider.NewLogger()),
			HeathCheckController:  v1.NewHealthCheckController(),
		},
		&rest.Middlewares{
			LoggerMiddleware: middleware.NewLoggerMiddleware(),
		},
	)

	return serverRest, facades

}
