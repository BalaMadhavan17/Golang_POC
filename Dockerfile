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

# Ensure required dependencies are fetched
RUN go get github.com/gorilla/mux \
    && go get github.com/rs/cors \
    && go get github.com/go-sql-driver/mysql

# Optional: Build a binary to avoid go run issues
RUN go build -o main .

# Expose the port your app runs on
EXPOSE 8080

# Run the compiled binary instead of go run
CMD ["./main"]

