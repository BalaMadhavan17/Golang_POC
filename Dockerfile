# Use the same Go version as your developer's laptop
FROM golang:1.25.1-alpine

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod go.sum ./

# Download Go modules
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Expose the port your app runs on
EXPOSE 8080

# Command to run the application (like npm start)
CMD ["go", "run", "main.go"]
