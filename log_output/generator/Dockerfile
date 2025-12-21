FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init generator && go build -o generator .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/generator .
CMD ["./generator"]
