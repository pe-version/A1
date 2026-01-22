FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go module file
COPY go.mod ./

# Copy source code
COPY main.go .

# Download dependencies and generate go.sum
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o sensor-service .

# Use a minimal image for the final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/sensor-service .

# Expose the port the service will run on
EXPOSE 8080

# Run the service
CMD ["./sensor-service"]
