FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init ping-pong && go build -o ping-pong-app .

FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/ping-pong-app /app/ping-pong-app
ENV PORT=8080
CMD ["/app/ping-pong-app"]
