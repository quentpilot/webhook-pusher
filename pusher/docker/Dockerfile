FROM golang:1.24-alpine

RUN go install github.com/cespare/reflex@latest

WORKDIR /app

COPY . .


RUN go mod tidy && go build -o webhook-pusher .

COPY reflex.conf .

EXPOSE 42001