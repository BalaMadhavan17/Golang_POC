FROM golang:1.25.1-alpine

WORKDIR /app

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

EXPOSE 8080

# Run the compiled binary instead of go run
CMD ["./main"]
