# Use the official Go image as the base image for building
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Ensure dependencies are resolved
RUN go mod tidy

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary ./main.go

# Use a minimal alpine image for the final stage
FROM alpine:latest

# Install ca-certificates and mysql-client
RUN apk --no-cache add ca-certificates mysql-client && update-ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app-binary .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./app-binary"]
