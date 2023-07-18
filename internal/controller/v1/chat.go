package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"smart-chat/adapter/provider"
	"smart-chat/internal/dto"
)

type chatController struct {
	chatFacade chatFacade
	logger     *provider.SystemLogger
}

func NewChatController(chatFacade chatFacade, logger *provider.SystemLogger) *chatController {
	return &chatController{
		chatFacade: chatFacade,
		logger:     logger,
	}
}

// Create - create a chat
// @Summary - create a chat
// @Description - create a chat
// @Tags - Chat
// @Accept json
// @Produce json
// @Success 201 {object} dto.ChatResponse
// @Failure 500 {object} dto.Error
// @Router /smart-chat/v1/chat [post]
// @Security ApiKeyAuth
func (c *chatController) Create(g *gin.Context) {
	requestID, found := g.Get("loggerID")
	if !found {
		requestID = uuid.New().String()
	}
	c.logger.NewLog("Call route create chat", "requestID", requestID.(string)).
		Debug().
		Description("Call route create chat").
		Phase("Controller").
		Request()

	ctx := context.WithValue(context.Background(), "requestID", requestID)

	chat, err := c.chatFacade.CreateChat(ctx)
	if err != nil {
		c.logger.NewLog("Error create chat", "requestID", requestID.(string)).
			Error().
			Description("Error create chat: " + err.Error()).
			Phase("Controller").
			Response()
		g.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}

	c.logger.NewLog("Success create chat", "requestID", requestID.(string)).
		Debug().
		Description("Success create chat").
		Phase("Controller").
		Response()

	g.JSON(http.StatusCreated, chat.ChatID)
}
