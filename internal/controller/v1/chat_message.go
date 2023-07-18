package v1

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"smart-chat/adapter/provider"
	"smart-chat/internal/constants"
	"smart-chat/internal/dto"
)

type chatMessageController struct {
	chatMessageFacade chatMessageFacade
	logger            *provider.SystemLogger
}

func NewChatMessageController(chatMessageFacade chatMessageFacade, logger *provider.SystemLogger) *chatMessageController {
	return &chatMessageController{
		chatMessageFacade: chatMessageFacade,
		logger:            logger,
	}
}

// Create - create a chat-message
// @Summary - create a chat-message
// @Description - create a chat-message
// @Tags - Chat-Message
// @Accept json
// @Param chatMessageRequest body dto.ChatMessageRequest true "chatMessageRequest"
// @Produce json
// @Success 201 {object} dto.ChatMessageResponse
// @Failure 404 {object} dto.Error
// @Failure 500 {object} dto.Error
// @Router /smart-chat/v1/chat-message [post]
// @Security ApiKeyAuth
func (c *chatMessageController) Create(g *gin.Context) {
	requestID, found := g.Get("loggerID")
	if !found {
		requestID = uuid.New().String()
	}

	c.logger.NewLog("Call route create chat-message", "requestID", requestID.(string)).
		Debug().
		Description("Call route create chat").
		Phase("Controller").
		Request()
	ctx := context.WithValue(context.Background(), "requestID", requestID)
	chatMessageRequest := &dto.ChatMessageRequest{}

	if err := g.BindJSON(chatMessageRequest); err != nil {
		c.logger.NewLog("Error bind json", "requestID", requestID.(string)).
			Error().
			Description("Error bind json: " + err.Error()).
			Phase("Controller").
			Response()
		g.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
		return
	}

	c.logger.NewLog("Success bind json", "requestID", requestID.(string),
		"chatMessageRequest", chatMessageRequest).
		Debug().
		Description("Success bind json").
		Phase("Controller").
		Request()

	chatMessage, err := c.chatMessageFacade.CreateChatMessage(ctx, chatMessageRequest)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrChatNotFound):
			c.logger.NewLog("Error chat not found", "requestID", requestID.(string)).
				Error().
				Description("Error chat not found: " + err.Error()).
				Phase("Controller").
				Response()
			g.JSON(http.StatusNotFound, &dto.Error{Message: err.Error()})
			return
		default:
			c.logger.NewLog("Error create chat-message", "requestID", requestID.(string)).
				Error().
				Description("Error create chat-message: " + err.Error()).
				Phase("Controller").
				Response()
			g.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		}
		return
	}

	c.logger.NewLog("Success create chat-message", "requestID", requestID.(string)).
		Debug().
		Description("Success create chat-message").
		Phase("Controller").
		Response()

	g.JSON(http.StatusCreated, chatMessage)
}
