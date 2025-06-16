ARG CONTEXT=prod
ARG GOLANG_IMAGE_VERSION=1.24.3

FROM golang:${GOLANG_IMAGE_VERSION}-alpine as base

## Setup
ARG CONTEXT
ARG GOLANG_IMAGE_VERSION

WORKDIR /root/openslides-autoupdate-service
ENV ${CONTEXT}=1

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

## Command
COPY ./dev/command.sh ./
RUN chmod +x command.sh
CMD ["./command.sh"]
HEALTHCHECK CMD ["/openslides-autoupdate-service", "health"]



# Development Image
FROM base as dev

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]



# Test Image
FROM base as tests

RUN apk add build-base --no-cache



# Production Image

FROM base as builder
RUN go build



FROM scratch as prod

WORKDIR /

LABEL org.opencontainers.image.title="OpenSlides Autoupdate Service"
LABEL org.opencontainers.image.description="The Autoupdate Service is a http endpoint where the clients can connect to get the current data and also updates."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-autoupdate-service"

COPY --from=builder /root/openslides-autoupdate-service/openslides-autoupdate-service .

EXPOSE 9012
ENTRYPOINT ["/openslides-autoupdate-service"]
