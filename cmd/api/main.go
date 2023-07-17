package main

import (
	"fmt"
	"smart-chat/adapter/database"
	"smart-chat/adapter/provider"
	"smart-chat/adapter/rest"
	"smart-chat/internal/config"
	v1 "smart-chat/internal/controller/v1"
	"smart-chat/internal/facade"
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
	cfg := config.NewConfigService().GetConfig()

	db, err := database.NewDatabaseProvider(cfg).GetConnection()
	if err != nil {
		panic(err)
	}

	chatRepository := repository.NewChatRepository(db)
	chatMessageRepository := repository.NewChatMessageRepository(db)

	chatService := service.NewChatService(chatRepository)
	chatMessageService := service.NewChatMessageService(chatMessageRepository)

	deepAiProvider := provider.NewDeepAiProvider(cfg)

	chatFacade := facade.NewChatFacade(chatService)
	chatMessageFacade := facade.NewChatMessageFacade(chatMessageService, deepAiProvider)

	chatController := v1.NewChatController(chatFacade)
	chatMessageController := v1.NewChatMessageController(chatMessageFacade)

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			ChatController:        chatController,
			ChatMessageController: chatMessageController,
			HeathCheckController:  v1.NewHealthCheckController(),
		},
	)
	fmt.Println("Server is running on port: ", cfg.RestPort)
	serverRest.StartListening()

}
