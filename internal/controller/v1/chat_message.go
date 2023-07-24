package v1

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"smart-chat/internal/constants"
	"smart-chat/internal/dto"
)

type chatMessageController struct {
	chatMessageFacade chatMessageFacade
	logger            *zap.Logger
}

func NewChatMessageController(chatMessageFacade chatMessageFacade, logger *zap.Logger) *chatMessageController {
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

	c.logger.Debug("Call route Create chat-message",
		zap.String("requestID", requestID.(string)),
		zap.String("phase", "Controller"))
	ctx := context.WithValue(context.Background(), "requestID", requestID)
	chatMessageRequest := &dto.ChatMessageRequest{}

	if err := g.BindJSON(chatMessageRequest); err != nil {
		c.logger.Error("Error bind json", zap.String("requestID", requestID.(string)),
			zap.Error(err),
			zap.String("phase", "Controller"))
		g.JSON(http.StatusBadRequest, &dto.Error{Message: err.Error()})
		return
	}

	c.logger.Debug("Success bind json", zap.String("requestID", requestID.(string)),
		zap.Any("chatMessageRequest", chatMessageRequest),
		zap.String("phase", "Controller"))

	chatMessage, err := c.chatMessageFacade.CreateChatMessage(ctx, chatMessageRequest)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrChatNotFound):
			c.logger.Error("Error create chat-message", zap.String("requestID", requestID.(string)),
				zap.Error(err),
				zap.String("phase", "Controller"))
			g.JSON(http.StatusNotFound, &dto.Error{Message: err.Error()})
			return
		default:
			c.logger.Error("Error create chat-message", zap.String("requestID", requestID.(string)),
				zap.Error(err),
				zap.String("phase", "Controller"))
			g.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		}
		return
	}

	c.logger.Debug("Success create chat-message", zap.String("requestID", requestID.(string)),
		zap.Any("chatMessage", chatMessage),
		zap.String("phase", "Controller"))

	g.JSON(http.StatusCreated, chatMessage)
}
