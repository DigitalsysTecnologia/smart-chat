package main

import (
	"github.com/apex/gateway"
	"log"
	"smart-chat/adapter/database"
	"smart-chat/adapter/rest"
	"smart-chat/internal/config"
	v1 "smart-chat/internal/controller/v1"
	"smart-chat/internal/facade"
	"smart-chat/internal/repository"
	"smart-chat/internal/service"
)

func main() {
	addr := ":80"
	cfg := config.NewConfigService().GetConfig()

	db, err := database.NewDatabaseProvider(cfg).GetConnection()
	if err != nil {
		panic(err)
	}

	chatRepository := repository.NewChatRepository(db)
	chatMessageRepository := repository.NewChatMessageRepository(db)

	chatService := service.NewChatService(chatRepository)
	chatMessageService := service.NewChatMessageService(chatMessageRepository)

	chatFacade := facade.NewChatFacade(chatService)
	chatMessageFacade := facade.NewChatMessageFacade(chatMessageService)

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

	log.Fatal(gateway.ListenAndServe(addr, serverRest.Engine))

}
