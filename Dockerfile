# Stage 1: Build the Go application
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Stage 2: Create a minimal final runtime container
FROM alpine:latest

# Set the working directory for the runtime container
WORKDIR /root/

# Install Redis client for debugging (optional)
RUN apk add --no-cache redis

ENV REDIS_ADDR=redis-server:6379

# Copy the built binary from the builder stage
COPY --from=builder /main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]