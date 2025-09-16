# Stage 1: Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git (needed for go mod) and build tools
RUN apk add --no-cache git

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Ensure go.sum is up-to-date
RUN go mod tidy && go mod download

# Copy the rest of the source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary ./main.go

# Stage 2: Minimal runtime
FROM alpine:3.18

RUN apk --no-cache add ca-certificates mysql-client

WORKDIR /root/
COPY --from=builder /app-binary .

CMD ["./app-binary"]
