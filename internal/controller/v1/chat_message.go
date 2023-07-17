package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"smart-chat/internal/dto"
)

type chatMessageController struct {
	chatMessageFacade chatMessageFacade
}

func NewChatMessageController(chatMessageFacade chatMessageFacade) *chatMessageController {
	return &chatMessageController{
		chatMessageFacade: chatMessageFacade,
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
// @Router /smart-chat/v1/chat-message [post]
// @Security ApiKeyAuth
func (c *chatMessageController) Create(g *gin.Context) {

	ctx := context.Background()

	chatMessageRequest := &dto.ChatMessageRequest{}

	if err := g.BindJSON(chatMessageRequest); err != nil {
		g.JSON(http.StatusBadRequest, err)
		return
	}

	chatMessage, err := c.chatMessageFacade.CreateChatMessage(ctx, chatMessageRequest)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	g.JSON(http.StatusCreated, chatMessage.Answer)
}
