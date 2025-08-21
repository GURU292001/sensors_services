# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project (including proto and sensorpb)
COPY . .

# Build server
RUN cd microservice_b && go build -o /server main.go

# Runtime stage
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /server .

EXPOSE 1323
CMD ["./server"]
