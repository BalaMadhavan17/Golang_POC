# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app-binary ./main.go

# Stage 2: Run minimal image
FROM alpine:3.18

# Install required tools
RUN apk --no-cache add ca-certificates mysql-client && update-ca-certificates

WORKDIR /root/

# Copy binary
COPY --from=builder /app/app-binary .

# Expose app port
EXPOSE 8080

CMD ["./app-binary"]
