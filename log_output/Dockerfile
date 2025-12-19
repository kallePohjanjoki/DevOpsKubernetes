FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init random-gen && go build -o main .

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/main /main
CMD ["/main"]
