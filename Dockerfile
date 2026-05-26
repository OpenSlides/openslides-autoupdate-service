ARG CONTEXT=prod

FROM golang:1.26.3-alpine AS base

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

## Command
COPY ./dev/command.sh ./
RUN chmod +x command.sh
CMD ["./command.sh"]

# Development Image
FROM base AS dev

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]

# Test Image
FROM base AS tests

COPY dev/container-tests.sh ./dev/container-tests.sh

RUN apk add --no-cache \
    build-base \
    docker && \
    go get -u github.com/ory/dockertest/v4 && \
    go install golang.org/x/lint/golint@latest && \
    chmod +x dev/container-tests.sh

## Command
STOPSIGNAL SIGKILL

# Production Image

FROM base AS base-gowork
COPY ./lib ../lib
COPY ./autoupdate.work ../go.work

FROM base-gowork AS builder-gowork
RUN go build

FROM base AS builder
RUN go build

FROM scratch AS pre-prod

## Setup
ARG CONTEXT
ENV APP_CONTEXT=prod

LABEL org.opencontainers.image.title="OpenSlides Autoupdate Service"
LABEL org.opencontainers.image.description="The Autoupdate Service is a http endpoint where the clients can connect to get the current data and also updates."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-autoupdate-service"

EXPOSE 9012
ENTRYPOINT ["/openslides-autoupdate-service"]

HEALTHCHECK CMD ["/openslides-autoupdate-service", "health"]

FROM pre-prod AS prod-gowork

COPY --from=builder-gowork /app/openslides-autoupdate-service/openslides-autoupdate-service /

FROM pre-prod AS prod

COPY --from=builder /app/openslides-autoupdate-service/openslides-autoupdate-service /
