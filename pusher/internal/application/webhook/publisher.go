package webhook

import "github.com/quentpilot/webhook-pusher/internal/domain/webhook"

type PublishWebhook struct {
	publisher WebhookPublisher
}

func NewPublishWebhook(publisher WebhookPublisher) *PublishWebhook {
	return &PublishWebhook{
		publisher: publisher,
	}
}

func (p *PublishWebhook) Publish(webhook *webhook.WebhookMessage) error {
	return p.publisher.Publish(webhook)
}
