# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy everything
COPY . .

# Build client
RUN cd microservice_a && go build -o /client main.go

# Runtime stage
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /client .

CMD ["./client"]
