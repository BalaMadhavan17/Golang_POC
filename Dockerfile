FROM golang:1.25.1-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go get github.com/gorilla/mux \
    && go get github.com/rs/cors \
    && go get github.com/go-sql-driver/mysql
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
