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

// @Summary chat-message
// @Description chat-message
// @Tags Action
// @Accept json
// @Param chatMessageRequest body dto.ChatMessageRequest true "chatMessageRequest"
// @Produce json
// @Success 201 {object} dto.ChatMessageResponse
// @Failure 500 {object} error
// @Router /smart-chat/v1/chat-message [post]
func (c *chatMessageController) CreateChatMessage(g *gin.Context) {

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
