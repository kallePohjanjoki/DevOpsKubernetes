FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init todo-app && go build -o todo-server .

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/todo-server /app/todo-server

COPY index.html /app/index.html

ENV PORT=8080

CMD ["/app/todo-server"]
