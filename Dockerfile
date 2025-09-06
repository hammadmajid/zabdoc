# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o zabcover ./cmd/zabcover/main.go

# Final stage
FROM alpine:latest
WORKDIR /app

# Install Chromium and dependencies
RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    libstdc++

# Copy built binary
COPY --from=builder /app/zabcover ./zabcover

# Copy web assets
COPY web ./web

# Heroku requires the port to be set via the PORT env var
#EXPOSE 8080

# Set environment for chromedp
ENV CHROME_PATH=/usr/bin/chromium-browser
ENV CHROMEDP_HEADLESS=true

# Run the app
CMD ["./zabcover"]
