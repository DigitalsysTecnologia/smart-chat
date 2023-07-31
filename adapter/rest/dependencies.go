package rest

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

type Controllers struct {
	ChatController        chatController
	ChatMessageController chatMessageController
	HeathCheckController  heathCheckController
}

type Middlewares struct {
	LoggerMiddleware loggerMiddleware
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

type loggerMiddleware interface {
	ValidateAccess(c *gin.Context)
}
