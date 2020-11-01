# Permission Service

The permisson service is a service and a libary for OpenSlides4 that tells, if a
user can see specific content.

## Build and run

### With Golang

```
go build ./cmd/permission/ && ./permission
```


### With Docker

The docker build uses the datastore-reader-service as default. Either configure
it to use the fake services (see environment variables below) or make sure the
service inside the docker container can connect to the datastore-reader. For
example with the docker argument --network host.

```
docker build . --tag openslides-permission
docker run --network host openslides-permission
```


### With Auto Restart

To restart the service when ever a source file has shanged, the tool
[CompileDaemon](https://github.com/githubnemo/CompileDaemon) can help.

```
go get github.com/githubnemo/CompileDaemon
CompileDaemon -log-prefix=false -build "go build ./cmd/permission" -command "./permission"
```

The make target `build-dev` creates a docker image that uses this tool:

```
make build-dev
docker run -v $(pwd)/cert:/root/cert --network host openslides-permission-dev
```


## Example Request

```
curl http://localhost:9005/internal/permission/is_allowed -d '{"name":"topic.create","user_id":1}'
```


## Test

### With Golang

```
go test ./...
```


### With Make

There is a make target, that creates and runs the docker-test-container:

```
make run-tests
```


## Environment Variables

* `PERMISSION_HOST`: Host where the http service listens to. Default is an empty
  string which means all devices.
* `PERMISSION_PORT`: Port where the http services listens to. Default is 9005.
* `DATASTORE`: Sets the datastore service. `fake` (default) or `service`.
* `DATASTORE_READER_HOST`: Host of the datastore reader. The default is
  `localhost`.
* `DATASTORE_READER_PORT`: Port of the datastore reader. The default is `9010`.
* `DATASTORE_READER_PROTOCOL`: Protocol of the datastore reader. The default is
  `http`.