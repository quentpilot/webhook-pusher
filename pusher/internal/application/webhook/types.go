package webhook

import "github.com/quentpilot/webhook-pusher/internal/domain/webhook"

type WebhookPublisher interface {
	Publish(message *webhook.WebhookMessage) error
}
