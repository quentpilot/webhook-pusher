package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quentpilot/webhook-pusher/internal/application/webhook"
	wd "github.com/quentpilot/webhook-pusher/internal/domain/webhook"
	"github.com/quentpilot/webhook-pusher/internal/middleware"
)

type WebhookHandler struct {
	publisher webhook.WebhookPublisher
}

func NewWebhookHandler(publisher webhook.WebhookPublisher) *WebhookHandler {
	return &WebhookHandler{
		publisher: publisher,
	}
}

func (h *WebhookHandler) Send() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info("Handle /send")

		req, ok := middleware.GetValidated[SendRequest](c)
		if !ok {
			slog.Error("Cannot retrieve validated request")
			c.JSON(http.StatusInternalServerError, gin.H{"message": "unexpected error"})
			return
		}

		slog.Info(fmt.Sprintf("Receive webhook %#v", req))

		reqId := uuid.New().String()

		webhookMessage := &wd.WebhookMessage{
			Id:       reqId,
			Event:    req.Event,
			Target:   req.Target,
			Payload:  req.Payload,
			Throttle: req.Throttle,
			Retry:    req.Retry,
			Fallback: req.Fallback,
			Sentry:   req.Sentry,
		}

		if err := h.publisher.Publish(webhookMessage); err != nil {
			res := &SendResponse{
				Message: err.Error(),
			}

			c.JSON(http.StatusInternalServerError, res)
			return
		}

		res := &SendResponse{
			Message: "webhook accepted",
			Uuid:    reqId,
		}

		c.JSON(http.StatusAccepted, res)
	}
}
