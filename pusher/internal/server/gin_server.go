package server

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quentpilot/webhook-pusher/internal/application/webhook"
	"github.com/quentpilot/webhook-pusher/internal/config"
	"github.com/quentpilot/webhook-pusher/internal/handlers"
	wi "github.com/quentpilot/webhook-pusher/internal/infra/amqp/webhook"
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

	config := &config.AmqpPubConfig{
		Host:  "amqp://guest:guest@queue-rabbit",
		Port:  5672,
		Queue: "webhook-sender",
	}

	pub := wi.NewAmqpPublisher(config)

	publisher := webhook.NewPublishWebhook(pub)
	handler := handlers.NewWebhookHandler(publisher)

	auth := s.Engine.Group("/",
		middleware.Timeout(5*time.Second),
		middleware.BindAndValidate[handlers.SendRequest](),
	)

	auth.POST("/send", handler.Send())
}

func (s *GinServer) Run() {
	port := ":" + strconv.Itoa(int(s.Config.Port))
	url := s.Config.Host + port

	slog.Info("API Server runs on " + url)

	s.Engine.Run(port)
}
