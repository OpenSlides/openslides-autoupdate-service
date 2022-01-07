FROM golang:1.17-alpine as base
WORKDIR /root/

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY internal internal
COPY pkg pkg

# Build service in seperate stage.
FROM base as builder
RUN CGO_ENABLED=0 go build ./cmd/autoupdate


# Test build.
FROM base as testing

RUN apk add build-base

CMD go vet ./... && go test ./...


# Development build.
FROM base as development

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
EXPOSE 9012
ENV MESSAGING redis
ENV AUTH ticket

CMD CompileDaemon -log-prefix=false -build="go build ./cmd/autoupdate" -command="./autoupdate"


# Productive build
FROM scratch

LABEL org.opencontainers.image.title="OpenSlides Autoupdate Service"
LABEL org.opencontainers.image.description="The Autoupdate Service is a http endpoint where the clients can connect to get the current data and also updates."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-autoupdate-service"

COPY --from=builder /root/autoupdate .
EXPOSE 9012
ENV MESSAGING redis
ENV AUTH ticket
ENTRYPOINT ["/autoupdate"]
