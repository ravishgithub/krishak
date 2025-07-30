
# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o khetiapp main.go

# Runtime Stage
FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/khetiapp .
COPY ./configs/config.json ./configs/config.json

EXPOSE 8080

CMD ["./khetiapp"]
