# Builder
FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o binaryapp cmd/app/main.go

# Runner
FROM alpine AS runner
WORKDIR /app

COPY ./.env /app/.env
ENV CONFIG_FILE=/app/.env
COPY --from=builder /app/binaryapp .

CMD ["./binaryapp"]