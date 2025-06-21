package handlers

type SendRequest struct {
	Event    string         `json:"event" binding:"required" message_required:"is required"`
	Target   string         `json:"target" binding:"required,url" message_required:"is required" message_url:"must be a valid url format"`
	Payload  map[string]any `json:"payload" binding:"required,omitempty" message_required:"is required"`
	Throttle int            `json:"throttle" binding:"omitempty,min=0" message_min:"muse be greater or equal to zero"`
	Retry    int            `json:"retry" binding:"omitempty,min=0" message_min:"must be greater or equal to zero"`
	Fallback string         `json:"fallback" binding:"omitempty,url" message_url:"must be a valid url format"`
	Sentry   bool           `json:"sentry" binding:"omitempty"`
}

type SendResponse struct {
	Message string `json:"message"`
	Uuid    string `json:"uuid,omitempty"`
}
