package main

import (
	"fmt"
	"go.uber.org/zap"
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

	middlewareLogger := middleware.NewLoggerMiddleware()

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
	serverRest.StartListening()

}
