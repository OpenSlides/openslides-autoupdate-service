ARG CONTEXT=prod

FROM golang:1.24.5-alpine as base

## Setup
ARG CONTEXT

WORKDIR /app/openslides-autoupdate-service
ENV APP_CONTEXT=${CONTEXT}

## Install
RUN apk add git --no-cache

COPY go.mod go.sum ./
RUN go mod download

COPY main.go main.go
COPY internal internal

## External Information
EXPOSE 9012

## Healthcheck
HEALTHCHECK CMD ["/app/openslides-autoupdate-service/openslides-autoupdate-service", "health"]

# Development Image
FROM base as dev

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]

CMD CompileDaemon -log-prefix=false -build="go build" -command="./openslides-autoupdate-service"

# Test Image
FROM base as tests

COPY dev/container-tests.sh ./dev/container-tests.sh

RUN apk add --no-cache \
    build-base \
    docker && \
    go get -u github.com/ory/dockertest/v3 && \
    go install golang.org/x/lint/golint@latest && \
    chmod +x dev/container-tests.sh

## Command
STOPSIGNAL SIGKILL
CMD ["sleep", "inf"]

# Production Image

FROM base as builder
RUN go build

FROM scratch as prod

## Setup
ARG CONTEXT
ENV APP_CONTEXT=prod

LABEL org.opencontainers.image.title="OpenSlides Autoupdate Service"
LABEL org.opencontainers.image.description="The Autoupdate Service is a http endpoint where the clients can connect to get the current data and also updates."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-autoupdate-service"

COPY --from=builder /app/openslides-autoupdate-service/openslides-autoupdate-service /

EXPOSE 9012
ENTRYPOINT ["/openslides-autoupdate-service"]

HEALTHCHECK CMD ["/openslides-autoupdate-service", "health"]
