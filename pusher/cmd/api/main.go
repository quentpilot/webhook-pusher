package main

import "github.com/quentpilot/webhook-pusher/internal/server"

// This package provides the main API to send webhook
// It loads HTTP server
// It handles user authorizations

func main() {
	config := &server.HttpConfig{
		Host: "http://localhost",
		Port: 42001,
	}

	s := server.NewGinServer(config)
	s.Load()
	s.Run()
}
