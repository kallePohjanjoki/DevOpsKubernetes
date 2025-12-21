FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init provider && go build -o provider .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/provider .
EXPOSE 8000
CMD ["./provider"]
