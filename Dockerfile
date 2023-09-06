FROM golang:1.21.0-alpine as base
WORKDIR /root/

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY main.go main.go
COPY internal internal
COPY pkg pkg

# Build service in seperate stage.
FROM base as builder
RUN go build


# Test build.
FROM base as testing

RUN apk add build-base

CMD go vet ./... && go test -test.short ./...


# Development build.
FROM base as development

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
EXPOSE 9012

CMD CompileDaemon -log-prefix=false -build="go build" -command="./openslides-autoupdate-service"


# Productive build
FROM scratch

LABEL org.opencontainers.image.title="OpenSlides Autoupdate Service"
LABEL org.opencontainers.image.description="The Autoupdate Service is a http endpoint where the clients can connect to get the current data and also updates."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-autoupdate-service"

COPY --from=builder /root/openslides-autoupdate-service .
EXPOSE 9012
ENTRYPOINT ["/openslides-autoupdate-service"]
HEALTHCHECK CMD ["/openslides-autoupdate-service", "health"]
