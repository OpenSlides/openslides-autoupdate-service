# OpenSlides Autoupdate Service

The Autoupdate Service is part of the OpenSlides environment. It is a http
endpoint where the clients can connect to get the actual data and also get
updates, when the requested data changes.

IMPORTANT: The data are sent via an open http-connection. All browsers limit the
amount of open http1.1 connections to a domain. For this service to work, the
browser has to connect to the service with http2 and therefore needs https.


## Start

### With Golang

```
go build ./cmd/autoupdate
./autoupdate
```


### With Docker

The docker build uses the redis messaging service, the auth token and the real
datastore service as default. Either configure it to use the fake services (see
environment variables below) or make sure the service inside the docker
container can connect to redis and the datastore-reader. For example with the
docker argument --network host.

```
docker build . --tag openslides-autoupdate
docker run --network host openslides-autoupdate
```

It uses the host network to connect to redis.


### With Auto Restart

To restart the service when ever a source file has shanged, the tool
[CompileDaemon](https://github.com/githubnemo/CompileDaemon) can help.

```
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -log-prefix=false -build "go build ./cmd/autoupdate" -command "./autoupdate"
```

The make target `build-dev` creates a docker image that uses this tool:

```
make build-dev
docker run -v $(pwd)/cert:/root/cert --network host openslides-autoupdate-dev
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


## Examples

Curl needs the flag `-N / --no-buffer` or it can happen, that the output is not
printed immediately.


### Without redis

When the server is started, clients can listen for keys to do so, they have to
send a keyrequest in the body of the request. Currently, all method-types (POST,
GET, etc) are supported. An example request is:

`curl -N  localhost:9012/system/autoupdate -d '[{"ids": [1], "collection": "user", "fields": {"username": null}}]'`

To see a list of possible json-strings see the file
internal/autoupdate/keysbuilder/keysbuilder_test.go

There is a simpler method to request keys:

`curl -N localhost:9012/system/autoupdate/keys?user/1/username,user/2/username`

With this simpler method, it is not possible to request related keys.

After the request is send, the values to the keys are returned as a json-object
without a newline:
```
{"user/1/name":"value","user/2/name":"value"}
```

To "update" keys, you can send them to the server via stdin with a value or
without a value in the form:

```
user/5/name
user/6/name="Emanuel"
user/1/group_ids=[1,2,3]
user/1/name="foo" user/2/name="bar"
```

If the value is skipped, the current time is used as value. If you give a value,
it has to be valid json without any spaces.

All clients that listen for the keys get an update for that key.


### With datastore-service

To connect the autoupdate-service with the datastore service, the following
environment variables can be used:

`DATASTORE=service MESSAGING=redis ./autoupdate`


### With redis

When redis is installed, it can be used to update keys. Start the autoupdate
service with the envirnmentvariable `MESSAGING_SERVICE=redis`. Afterwards it is
possible to update keys by sending the following command to redis:

`xadd field_changed * updated user/1/username updated user/1/password`


### Projector

To get projector data, you can use:

`curl -N localhost:9012/system/autoupdate/projector?projector_ids=1,2,3`


## Environment

The Service uses the following environment variables:

* `AUTOUPDATE_PORT`: Lets the service listen on port 9012. The default is
  `9012`.
* `AUTOUPDATE_HOST`: The device where the service starts. The default is am
  empty string which starts the service on any device.
* `DATASTORE`: Sets the datastore service. `fake` (default) or `service`.
* `DATASTORE_READER_HOST`: Host of the datastore reader. The default is
  `localhost`.
* `DATASTORE_READER_PORT`: Port of the datastore reader. The default is `9010`.
* `DATASTORE_READER_PROTOCOL`: Protocol of the datastore reader. The default is
  `http`.
* `MESSAGING`: Sets the type of messaging service. `fake`(default) or
  `redis`.
* `MESSAGE_BUS_HOST`: Host of the redis server. The default is `localhost`.
* `MESSAGE_BUS_PORT`: Port of the redis server. The default is `6379`.
* `REDIS_TEST_CONN`: Test the redis connection on startup. Disable on the cloud
  if redis needs more time to start then this service. The default is `true`.
* `AUTH`: Sets the type of the auth service. `fake` (default) or `ticket`.
* `AUTH_KEY_TOKEN`: Key to sign the JWT auth tocken. Default `auth-dev-key`.
* `AUTH_KEY_COOKIE`: Key to sign the JWT auth cookie. Default `auth-dev-key`.
* `AUTH_HOST`: Host of the auth service. The default is `localhost`.
* `AUTH_PORT`: Port of the auth service. The default is `9004`.
* `AUTH_PROTOCOL`: Protocol of the auth servicer. The default is `http`.
* `DEACTIVATE_PERMISSION`: Deactivate requests to the permission service. The
  result is, that every user can see everything (Defaullt: `false`)
