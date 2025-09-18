package webhook

import (
	"fmt"
	"strconv"

	"github.com/quentpilot/webhook-pusher/internal/config"
	"github.com/quentpilot/webhook-pusher/internal/domain/webhook"
	"github.com/rabbitmq/amqp091-go"
)

type AmqpPublisher struct {
	config *config.AmqpPubConfig
	conn   *amqp091.Connection
}

func NewAmqpPublisher(config *config.AmqpPubConfig) *AmqpPublisher {
	url := config.Host + ":" + strconv.Itoa(config.Port)
	//url := config.Host

	conn, err := amqp091.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ [dsn=%s]: %s", url, err))
	}

	return &AmqpPublisher{
		config: config,
		conn:   conn,
	}
}

func (p *AmqpPublisher) Publish(message *webhook.WebhookMessage) error {

	_, err := p.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %v", err.Error())
	}

	return fmt.Errorf("publisher not implemented yet, config: %#v", p.config)
}
