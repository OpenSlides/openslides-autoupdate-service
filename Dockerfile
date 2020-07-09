FROM golang:1.14-alpine as basis
LABEL maintainer="OpenSlides Team <info@openslides.com>"
WORKDIR /root/

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .


# Build service in seperate stage.
FROM basis as builder
RUN go build ./cmd/autoupdate


# Test build.
From basis as testing

RUN apk add build-base

CMD go vet ./... && go test ./...


# Development build.
FROM basis as development

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
EXPOSE 9012
ENV DATASTORE service
ENV MESSAGING redis

CMD CompileDaemon -log-prefix=false -build="go build ./cmd/autoupdate" -command="./autoupdate"


# Productive build
FROM alpine:latest
WORKDIR /root/

COPY --from=builder /root/autoupdate .
EXPOSE 9012
ENV DATASTORE service
ENV MESSAGING redis

CMD ./autoupdate
