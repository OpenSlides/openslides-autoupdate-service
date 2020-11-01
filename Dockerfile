FROM golang:1.15.3-alpine3.12 as basis
LABEL maintainer="OpenSlides Team <info@openslides.com>"
WORKDIR /root/

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY pkg pkg

# Build service in seperate stage.
FROM basis as builder
RUN go build ./cmd/permission


# Test build.
From basis as testing

RUN apk add build-base

CMD go vet ./... && go test ./...


# Development build.
FROM basis as development

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
EXPOSE 9005
ENV DATASTORE service

CMD CompileDaemon -log-prefix=false -build="go build ./cmd/permission" -command="./permission"


# Productive build
FROM alpine:3.12.1
WORKDIR /root/

COPY --from=builder /root/permission .
EXPOSE 9012
ENV DATASTORE service

CMD ./permission
