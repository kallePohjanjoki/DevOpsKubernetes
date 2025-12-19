FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY main.go .
RUN go mod init todo-app && go build -o todo-server .

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/todo-server /todo-server

ENV PORT=8080

CMD ["/todo-server"]
