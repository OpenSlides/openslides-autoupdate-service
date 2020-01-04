FROM golang:1.13
LABEL maintainer="OpenSlides Team <info@openslides.com>"
WORKDIR /root/

# Preload dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy everything from the current directory
COPY . .

# Test the Go app
CMD go test ./...