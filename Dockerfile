FROM golang:1.23-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod tidy && go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary ./main.go

FROM alpine:3.18
RUN apk --no-cache add ca-certificates mysql-client

WORKDIR /root/
COPY --from=builder /app-binary .

EXPOSE 8080
CMD ["./app-binary"]
