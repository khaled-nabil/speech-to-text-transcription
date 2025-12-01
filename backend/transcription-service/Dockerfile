# Build stage
FROM golang:1.25.3-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o transcription-service ./cmd/main.go

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add curl

WORKDIR /app

COPY --from=builder /app/transcription-service .

EXPOSE 8080

CMD ["./transcription-service"]

