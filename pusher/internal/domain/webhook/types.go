package webhook

type WebhookMessage struct {
	Id       string         `json:"id"`
	Event    string         `json:"event"`
	Target   string         `json:"target"`
	Payload  map[string]any `json:"payload"`
	Throttle int            `json:"throttle"`
	Retry    int            `json:"retry"`
	Fallback string         `json:"fallback"`
	Sentry   bool           `json:"sentry"`
}

type RetryConfig struct {
	MaxRetry  uint
	CurrRetry uint
}

type FallbackConfig struct {
	Target string
}

type BackoffHandler interface {
	Calculate()
}

type BackoffConfig struct {
	Step uint
}

func (b *BackoffConfig) Calculate() {

}

type ThrottleConfig struct {
	Interval uint
}

type Webhook struct {
	Uuid     string
	Event    string
	Target   string
	Payload  map[string]any
	Retry    RetryConfig
	Fallback FallbackConfig
	Backoff  BackoffHandler
	Throttle ThrottleConfig
}

func NewWebhook() *Webhook {
	return &Webhook{}
}

/* func NewWebhookFromRequest(uuid string, req *handlers.SendRequest) *Webhook {
	retry := RetryConfig{MaxRetry: uint(req.Retry)}
	fallback := FallbackConfig{Target: req.Fallback}
	backoff := BackoffConfig{}
	throttle := ThrottleConfig{Interval: uint(req.Throttle)}

	return &Webhook{
		Uuid:     uuid,
		Event:    req.Event,
		Target:   req.Target,
		Payload:  req.Payload,
		Retry:    retry,
		Fallback: fallback,
		Backoff:  &backoff,
		Throttle: throttle,
	}
} */
