# Use full Go toolchain
FROM golang:1.23

# Set working directory
WORKDIR /app/microservice_b

# Copy module files first to cache dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# Copy the full service code
COPY . .

# Expose ports
EXPOSE 50051 8080

# Run server.go directly
CMD ["go", "run", "client.go"]
