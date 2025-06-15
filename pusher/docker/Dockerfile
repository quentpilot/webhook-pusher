FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o webhook-pusher .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/webhook-pusher .

EXPOSE 42001

CMD ["./webhook-pusher"]