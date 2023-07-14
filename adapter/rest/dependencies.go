package rest

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type Controllers struct {
	chatMessageController chatMessageController
	chatController        chatController
	heathCheckController  heathCheckController
}

type chatMessageController interface {
	Create(c *gin.Context)
}

type chatController interface {
	Create(c *gin.Context)
}

type heathCheckController interface {
	HealthCheck(c *gin.Context)
}
