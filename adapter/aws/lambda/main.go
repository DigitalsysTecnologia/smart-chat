package main

import (
	"fmt"
	"github.com/apex/gateway"
	"log"
	"smart-chat/adapter/database"
	"smart-chat/adapter/provider"
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
	sysLogger := provider.NewLogger()
	defer sysLogger.ZapSync()

	cfg := config.NewConfigService().GetConfig()

	db, err := database.NewDatabaseProvider(cfg).GetConnection()
	if err != nil {
		panic(err)
	}

	chatRepository := repository.NewChatRepository(db, sysLogger)
	chatMessageRepository := repository.NewChatMessageRepository(db, sysLogger)

	chatService := service.NewChatService(chatRepository, sysLogger)
	chatMessageService := service.NewChatMessageService(chatMessageRepository, sysLogger)

	deepAiProvider := provider.NewDeepAiProvider(cfg, sysLogger)

	chatFacade := facade.NewChatFacade(chatService, sysLogger)
	chatMessageFacade := facade.NewChatMessageFacade(chatMessageService, chatService, deepAiProvider, sysLogger)

	chatController := v1.NewChatController(chatFacade, sysLogger)
	chatMessageController := v1.NewChatMessageController(chatMessageFacade, sysLogger)

	logger := middleware.NewLoggerMiddleware()

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			ChatController:        chatController,
			ChatMessageController: chatMessageController,
			HeathCheckController:  v1.NewHealthCheckController(),
		},
		&rest.Middlewares{
			LoggerMiddleware: logger,
		},
	)
	fmt.Println("Server is running on port: ", cfg.RestPort)

	log.Fatal(gateway.ListenAndServe(addr, serverRest.Engine))

}
