# Use the official Golang image to create a build artifact.
FROM golang:1.22.2-alpine AS build
RUN adduser -D -g '' elf
WORKDIR /opt/web
COPY go.mod go.sum ./
RUN go mod download && \
    go mod verify
COPY . .
RUN go build -o /opt/web/logistik-web ./cmd/web

# Start a new stage from scratch using a lightweight alpine image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN adduser -D -g '' elf
WORKDIR /opt/web/
COPY --from=build /opt/web/logistik-web .
RUN chown elf:elf logistik-web
COPY ./ui ui
USER elf
EXPOSE 8000
CMD ["./logistik-web"]
