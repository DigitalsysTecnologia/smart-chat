package main

import (
	"fmt"
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

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

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
	serverRest.StartListening()

}
