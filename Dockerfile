# Stage 1: Build Go binary
FROM golang:1.23-alpine AS builder

# Install git (needed for downloading Go modules)
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./

# Download dependencies (ensure go.sum is updated properly)
RUN go mod tidy && go mod download

# Force install missing dependencies explicitly
RUN go get github.com/gorilla/mux \
    github.com/rs/cors \
    github.com/go-sql-driver/mysql && \
    go mod tidy && go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application into a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary ./main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

# Install ca-certificates and mysql-client for DB connectivity
RUN apk --no-cache add ca-certificates mysql-client

# Set working directory
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app-binary .

# Expose application port
EXPOSE 8080

# Run the binary
CMD ["./app-binary"]
