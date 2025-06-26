ARG CONTEXT=prod

FROM golang:1.24.3-alpine as base

## Setup
ARG CONTEXT

WORKDIR /app/openslides-autoupdate-service
# Used for easy target differentiation
ARG ${CONTEXT}=1 
ENV APP_CONTEXT=${CONTEXT}

## Install
RUN apk add git --no-cache

COPY go.mod go.sum ./
RUN go mod download

COPY main.go main.go
COPY internal internal

## External Information
LABEL org.opencontainers.image.title="OpenSlides Autoupdate Service"
LABEL org.opencontainers.image.description="The Autoupdate Service is a http endpoint where the clients can connect to get the current data and also updates."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-autoupdate-service"

EXPOSE 9012

## Healthcheck
HEALTHCHECK CMD ["/app/openslides-autoupdate-service/openslides-autoupdate-service", "health"]



# Development Image
FROM base as dev

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]

CMD CompileDaemon -log-prefix=false -build="go build" -command="./openslides-autoupdate-service"


# Test Image
FROM base as tests

RUN apk add build-base --no-cache

CMD go vet ./... && go test -test.short ./...


# Production Image

FROM base as builder
RUN go build



FROM scratch as prod

WORKDIR /
ENV APP_CONTEXT=prod

LABEL org.opencontainers.image.title="OpenSlides Autoupdate Service"
LABEL org.opencontainers.image.description="The Autoupdate Service is a http endpoint where the clients can connect to get the current data and also updates."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-autoupdate-service"

COPY --from=builder /app/openslides-autoupdate-service/openslides-autoupdate-service .

EXPOSE 9012
ENTRYPOINT ["/openslides-autoupdate-service"]

HEALTHCHECK CMD ["/openslides-autoupdate-service", "health"]