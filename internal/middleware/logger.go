package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type loggerMiddleware struct {
	tokenProvider tokenProvider
}

func NewLoggerMiddleware(tokenProvider tokenProvider) *loggerMiddleware {
	return &loggerMiddleware{
		tokenProvider,
	}
}

func (l *loggerMiddleware) ValidateAccess(c *gin.Context) {

	token := c.GetHeader("AuthToken")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token cannot be empty"})
		return
	}

	if err := l.tokenProvider.ValidateToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	requestUUID := uuid.New()
	c.Set("loggerID", requestUUID.String())
	return
}
