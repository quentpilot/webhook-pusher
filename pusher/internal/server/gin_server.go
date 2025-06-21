package server

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quentpilot/webhook-pusher/internal/handlers"
	"github.com/quentpilot/webhook-pusher/internal/middleware"
)

type GinServer struct {
	Engine *gin.Engine // Current HTTP server (http, Gin, Fiber,...)
	Config *HttpConfig // Server configuration
}

func NewGinServer(c *HttpConfig) *GinServer {

	return &GinServer{
		Engine: gin.Default(),
		Config: c,
	}
}

func (s *GinServer) Load() {
	slog.Info("Loading API endpoints")

	auth := s.Engine.Group("/",
		middleware.Timeout(5*time.Second),
		middleware.BindAndValidate[handlers.SendRequest](),
	)

	auth.POST("/send", handlers.SendWebhook())
}

func (s *GinServer) Run() {
	port := ":" + strconv.Itoa(int(s.Config.Port))
	url := s.Config.Host + port

	slog.Info("API Server runs on " + url)

	s.Engine.Run(port)
}
