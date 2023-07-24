package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"smart-chat/internal/dto"
)

type chatController struct {
	chatFacade chatFacade
	logger     *zap.Logger
}

func NewChatController(chatFacade chatFacade, logger *zap.Logger) *chatController {
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

	c.logger.Debug("Call route Create chat",
		zap.String("requestID", requestID.(string)),
		zap.String("phase", "Controller"))

	ctx := context.WithValue(context.Background(), "requestID", requestID)

	chat, err := c.chatFacade.CreateChat(ctx)
	if err != nil {
		c.logger.Error("Error create chat",
			zap.String("requestID", requestID.(string)),
			zap.Error(err),
			zap.String("phase", "Controller"))
		g.JSON(http.StatusInternalServerError, &dto.Error{Message: err.Error()})
		return
	}

	c.logger.Debug("Success create chat",
		zap.String("requestID", requestID.(string)),
		zap.Any("chat", chat),
		zap.String("phase", "Controller"))

	g.JSON(http.StatusCreated, chat.ChatID)
}
