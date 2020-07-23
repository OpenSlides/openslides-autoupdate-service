# OpenSlides Autoupdate Service

The Autoupdate Service is part of the OpenSlides environment. It is a http
endpoint where the clients can connect to get the actual data and also get
updates, when the requested data changes.


## Start

The service requires https and therefore needs https. If no certificat is given,
the service creates and uses an inmemory self signed certificat. To create a
valid certificat for development, the tool
[mkcert](https://github.com/FiloSottile/mkcert) can be used. If `mkcert` is
installed, the make target `make dev-cert` can be used to create a certivicate
for the autoupdate-service on localhost.


### With Golang

```
go build ./cmd/autoupdate
./autoupdate
```


### With Docker

The docker build uses the redis messaging service and the real datastore service
as default. Either configure it to use the fake services (see environment
variables below) or make sure the service inside the docker container can
connect to redis and the datastore-reader. For example with the docker argument

```
docker build . --tag openslides-autoupdate
docker run --network host openslides-autoupdate
```

It uses the host network to connect to redis.


### With Auto Restart

To restart the service when ever a source file has shanged, the tool
[CompileDaemon](https://github.com/githubnemo/CompileDaemon) can help.

```
go get github.com/githubnemo/CompileDaemon
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
printed immediately. With a self signed certificat (the default of the
autoupdate-service) is also needs the flag `-k / --insecure`.


### Without redis

When the server is started, clients can listen for keys to do so, they have to
send a keyrequest in the body of the request. Currently, all method-types (POST,
GET, etc) are supported. An example request is:

`curl -Nk  https://localhost:9012/system/autoupdate -d '[{"ids": [5], "collection": "user", "fields": {"name": null}}]'`

To see a list of possible json-strings see the file
internal/autoupdate/keysbuilder/keysbuilder_test.go

There is a simpler method to request keys:

`curl -Nk https://localhost:9012/system/autoupdate/keys?user/1/name,user/2/name`

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

`xadd field_changed * updated user/5/name updated user/5/password`


## Environment

The Service uses the following environment variables:

* `AUTOUPDATE_PORT`: Lets the service listen on port 9012. The default is
  `9012`.
* `AUTOUPDATE_HOST`: The device where the service starts. The default is am
  empty string which starts the service on any device.
* `CERT_DIR`: Path where the tls certificates and the keys are. If emtpy, the
  server creates a self signed inmemory certificat. The default is empty.
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
