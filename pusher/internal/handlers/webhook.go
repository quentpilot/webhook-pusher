package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/quentpilot/webhook-pusher/internal/middleware"
)

func SendWebhook() gin.HandlerFunc {
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
		res := &SendResponse{
			Message: "webhook accepted",
			Uuid:    reqId,
		}

		c.JSON(http.StatusAccepted, res)
	}
}
