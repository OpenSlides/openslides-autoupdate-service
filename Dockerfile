FROM golang:1.14-alpine as builder
LABEL maintainer="OpenSlides Team <info@openslides.com>"
WORKDIR /root/

# Install git, it is needed to install the dependencies
RUN apk --no-cache add git

# Preload dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy everything from the current directory
COPY . .

# Build the Go app
RUN go build ./cmd/autoupdate


######## Start a new stage from scratch #######
FROM alpine:latest
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /root/autoupdate .

EXPOSE 8002

CMD ./autoupdate
