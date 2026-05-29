FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -mod=vendor -o scanner .

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/scanner .
EXPOSE 8080
CMD ["./scanner", "--web"]
