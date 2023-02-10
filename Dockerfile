FROM golang:latest

WORKDIR /app
COPY . .

RUN go build -o server server.go
CMD ["./server"]

