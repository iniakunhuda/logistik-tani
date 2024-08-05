# Use the official Golang image to create a build artifact.
FROM golang:1.22.2-alpine AS builder
RUN adduser -D -g '' elf
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o sales_service .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/sales_service .
COPY .env .
EXPOSE 4002
CMD ["./sales_service"]