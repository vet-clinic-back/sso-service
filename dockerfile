# Используем базовый образ Go
FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /main ./cmd/sso/main.go

CMD ["/main"]