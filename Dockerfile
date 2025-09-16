FROM --platform=linux/amd64 golang:1.23-alpine AS builder
# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY . .

# Explicitly add dependencies to ensure go.sum is populated
RUN go install

EXPOSE 8080
# Command to run the application
CMD ["go","run","main.go"]
