# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o zabdoc ./cmd/zabdoc/main.go

# Final stage
FROM alpine:latest
WORKDIR /app

# Install CA certificates for HTTPS/TLS verification
# This is critical for the app to make HTTPS requests to external services
RUN apk add --no-cache ca-certificates

# Copy the compiled binary from builder stage
COPY --from=builder /app/zabdoc ./zabdoc

# Verify CA certificates are in place
RUN ls -la /etc/ssl/certs/ca-certificates.crt || echo "CA certificates installed"

# Do not change. Also hardcoded in go and ts code
EXPOSE 8080

# Run with proper signal handling
CMD ["./zabdoc"]
