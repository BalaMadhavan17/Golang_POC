# Stage 1: Build
FROM golang:1.23-alpine AS builder

# Install dependencies for build
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Download Go modules
RUN go mod tidy
RUN go get github.com/gorilla/mux \
    github.com/rs/cors \
    github.com/go-sql-driver/mysql
RUN go mod tidy

# Copy the rest of the application
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary ./main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates mysql-client

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app-binary .

# Expose port and run
EXPOSE 8080
CMD ["./app-binary"]
