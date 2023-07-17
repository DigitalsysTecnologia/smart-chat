package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type chatController struct {
	chatFacade chatFacade
}

func NewChatController(chatFacade chatFacade) *chatController {
	return &chatController{
		chatFacade: chatFacade,
	}
}

// Create - create a chat
// @Summary - create a chat
// @Description - create a chat
// @Tags - Chat
// @Accept json
// @Produce json
// @Success 201 {object} dto.ChatResponse
// @Router /smart-chat/v1/chat [post]
// @Security ApiKeyAuth
func (c *chatController) Create(g *gin.Context) {

	ctx := context.Background()

	chat, err := c.chatFacade.CreateChat(ctx)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	g.JSON(http.StatusCreated, chat.ChatID)
}
