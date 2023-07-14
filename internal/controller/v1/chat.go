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

// @Summary create chat
// @Description create chat
// @Tags Action
// @Accept json
// @Produce json
// @Success 201 {object} dto.ChatResponse
// @Failure 500 {object} error
// @Router /smart-chat/v1/chat [post]
func (c *chatController) Create(g *gin.Context) {

	ctx := context.Background()

	chat, err := c.chatFacade.CreateChat(ctx)
	if err != nil {

	}

	g.JSON(http.StatusCreated, chat.ChatID)
}
