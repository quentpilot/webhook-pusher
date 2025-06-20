package server

import (
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
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
}

func (s *GinServer) Run() {
	port := ":" + strconv.Itoa(int(s.Config.Port))
	url := s.Config.Host + port

	slog.Info("API Server runs on " + url)

	s.Engine.Run(port)
}
