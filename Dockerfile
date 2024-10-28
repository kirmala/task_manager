FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o task_manager .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/task_manager .
EXPOSE 8080
CMD ["./task_manager"]