# Stage 1: Build the Go application
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy Go module files and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o /app/receipt-processor . && ls -lah /app/receipt-processor

# Stage 2: Create a minimal final runtime container
FROM alpine:latest

# Set the working directory for the runtime container
WORKDIR /app

# Install Redis client for debugging (optional)
RUN apk add --no-cache redis

# Copy the built binary from the builder stage
COPY --from=builder /app/receipt-processor /app/receipt-processor

# Set environment variables for the application
ENV SERVER_PORT=:8080
ENV REDIS_ADDR=redis:6379

# Expose the application port
EXPOSE 8080

# Ensure binary has execute permission
RUN chmod +x /app/receipt-processor

# Command to run the application
CMD ["/app/receipt-processor"]