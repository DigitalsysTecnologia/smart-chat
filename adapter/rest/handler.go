package rest

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"smart-chat/docs"
	"smart-chat/internal/model"
	"time"
)

type ServerRest struct {
	httpServer  *http.Server
	Engine      *gin.Engine
	config      *model.Config
	controllers *Controllers
	middlewares *Middlewares
}

func NewRestServer(cfg *model.Config, controllers *Controllers, middlewares *Middlewares) *ServerRest {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())

	docs.SwaggerInfo.Title = "SmartChat - API"
	docs.SwaggerInfo.Description = "API para comunicação com o sistema DeepAI"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	server := &ServerRest{
		Engine:      engine,
		config:      cfg,
		controllers: controllers,
		middlewares: middlewares,
	}

	server.registerRoutes()

	return server
}

func (s *ServerRest) registerRoutes() {
	v1 := s.Engine.Group("smart-chat/v1", s.middlewares.LoggerMiddleware.GenerateLoggerID)
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1.GET("/health", s.controllers.HeathCheckController.HealthCheck)

		chatGroup := v1.Group("/chat")
		{
			chatGroup.POST("", s.controllers.ChatController.Create)
		}

		chatMessageGroup := v1.Group("/chat-message")
		{
			chatMessageGroup.POST("", s.controllers.ChatMessageController.Create)
		}
	}

}

func (s *ServerRest) StartListening() {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.RestPort),
		Handler: s.Engine,
	}

	fmt.Println("Listening on port", s.config.RestPort)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err.Error())
	}
}

func (s *ServerRest) StopListening(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	err := s.httpServer.Shutdown(ctxWithTimeout)
	if err != nil {
		fmt.Println("http server forced to shutdown due to timeout")
	}

	fmt.Println("http server was gracefully stopped")
}
