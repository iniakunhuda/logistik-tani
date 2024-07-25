# # base image
# FROM golang:1.19.3-alpine AS builder
# # create appuser
# RUN adduser -D -g '' elf
# # create workspace
# WORKDIR /opt/app/
# COPY go.mod go.sum ./
# # fetch dependancies
# RUN go mod download && \
#     go mod verify
# # copy the source code as the last step
# COPY . .
# # build binary
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/logistik-api-users ./cmd/app


# # build a small image
# FROM alpine:3.17.3
# LABEL language="golang"
# LABEL org.opencontainers.image.source https://github.com/mmorejon/microservices-docker-go-mongodb
# # import the user and group files from the builder
# COPY --from=builder /etc/passwd /etc/passwd
# # copy the static executable
# COPY --from=builder --chown=elf:1000 /go/bin/logistik-api-users /logistik-api-users
# # use a non-root user
# USER elf
# # run app
# ENTRYPOINT ["./logistik-api-users"]



###
# Development Only
###
# base image
FROM golang:1.22.2-alpine
# Install necessary build tools
RUN apk add --no-cache git

# create appuser
RUN adduser -D -g '' elf

# After creating the elf user and before setting the WORKDIR
RUN go install github.com/air-verse/air@latest

# create workspace
WORKDIR /opt/app/

# Use go mod download as a separate step to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Set the user
USER elf

# Run the application
CMD ["air", "-c", ".air.toml"]