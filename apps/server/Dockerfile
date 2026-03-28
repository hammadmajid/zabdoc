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
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/zabdoc ./zabdoc

# Do not change. Also hardcoded in go and ts code
EXPOSE 8080
CMD ["./zabdoc"]
