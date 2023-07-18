package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type loggerMiddleware struct {
}

func NewLoggerMiddleware() *loggerMiddleware {
	return &loggerMiddleware{}
}

func (l *loggerMiddleware) GenerateLoggerID(c *gin.Context) {
	requestUUID := uuid.New()
	c.Set("loggerID", requestUUID.String())
}
