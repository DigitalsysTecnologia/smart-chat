package main

import (
	"fmt"
	"github.com/apex/gateway"
	"go.uber.org/zap"
	"log"
	"smart-chat/adapter/database"
	"smart-chat/adapter/provider"
	"smart-chat/adapter/provider/medpass"
	"smart-chat/adapter/rest"
	"smart-chat/internal/config"
	v1 "smart-chat/internal/controller/v1"
	"smart-chat/internal/facade"
	"smart-chat/internal/middleware"
	"smart-chat/internal/repository"
	"smart-chat/internal/service"
)

func main() {
	addr := ":80"
	logger := zap.NewExample()

	cfg := config.NewConfigService().GetConfig()

	db, err := database.NewDatabaseProvider(cfg).GetConnection()
	if err != nil {
		panic(err)
	}

	chatRepository := repository.NewChatRepository(db, logger)
	chatMessageRepository := repository.NewChatMessageRepository(db, logger)

	chatService := service.NewChatService(chatRepository, logger)
	chatMessageService := service.NewChatMessageService(chatMessageRepository, logger)

	deepAiProvider := provider.NewDeepAiProvider(cfg, logger)

	chatFacade := facade.NewChatFacade(chatService, logger)
	chatMessageFacade := facade.NewChatMessageFacade(chatMessageService, chatService, deepAiProvider, logger)

	chatController := v1.NewChatController(chatFacade, logger)
	chatMessageController := v1.NewChatMessageController(chatMessageFacade, logger)

	tokenProvider := medpass.NewAuthorizerGateway()

	middlewareLogger := middleware.NewLoggerMiddleware(tokenProvider)

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			ChatController:        chatController,
			ChatMessageController: chatMessageController,
			HeathCheckController:  v1.NewHealthCheckController(),
		},
		&rest.Middlewares{
			LoggerMiddleware: middlewareLogger,
		},
	)
	fmt.Println("Server is running on port: ", cfg.RestPort)

	log.Fatal(gateway.ListenAndServe(addr, serverRest.Engine))

}
