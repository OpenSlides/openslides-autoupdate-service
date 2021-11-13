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
docker argument --network host. The auth-secrets have to given as a file.

```
docker build . --tag openslides-autoupdate
printf "my_token_key" > auth_token_key 
printf "my_cookie_key" > auth_cookie_key
docker run --network host -v $PWD/auth_token_key:/run/secrets/auth_token_key -v $PWD/auth_cookie_key:/run/secrets/auth_cookie_key openslides-autoupdate
```

It uses the host network to connect to redis.


### With Auto Restart

To restart the service when ever a source file has shanged, the tool
[CompileDaemon](https://github.com/githubnemo/CompileDaemon) can help.

```
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -log-prefix=false -build "go build ./cmd/autoupdate" -command "./autoupdate"
```

The make target `build-dev` creates a docker image that uses this tool. The
environment varialbe `OPENSLIDES_DEVELOPMENT` is used to use default auth keys.

```
make build-dev
docker run --network host --env OPENSLIDES_DEVELOPMENT=true openslides-autoupdate-dev
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


### HTTP requests

When the server is started, clients can listen for keys to do so, they have to
send a keyrequest in the body of the request. An example request is:

`curl -N localhost:9012/system/autoupdate -d '[{"ids": [1], "collection": "user", "fields": {"username": null}}]'`

To see a list of possible json-strings see the file
internal/autoupdate/keysbuilder/keysbuilder_test.go

Keys can also defined with the query parameter `k`:

`curl -N localhost:9012/system/autoupdate?k=user/1/username,user/2/username`

With this query method, it is not possible to request related keys.

A request can have a body and the `k`-query parameter.

After the request is send, the values to the keys are returned as a json-object
without a newline:
```
{"user/1/name":"value","user/2/name":"value"}
```


### Updates via redis

When redis is installed, it can be used to update keys. Start the autoupdate
service with the envirnmentvariable `MESSAGING_SERVICE=redis`. Afterwards it is
possible to update keys by sending the following command to redis:

`xadd field_changed * updated user/1/username updated user/1/password`


### Projector

The data for a projector can be accessed with autoupdate requests. For example use:


```
curl -N localhost:9012/system/autoupdate -d '
[
  {
    "ids": [1],
    "collection": "projector",
    "fields": {
      "current_projection_ids": {
        "type": "relation-list",
        "collection": "projection",
        "fields": {
          "content": null,
          "content_object_id": null,
          "stable": null,
          "type": null,
          "options": null
        }
      }
    }
  }
]'
```

## Configuration

### Environment variables

The Service uses the following environment variables:

* `AUTOUPDATE_PORT`: Lets the service listen on port 9012. The default is
  `9012`.
* `AUTOUPDATE_HOST`: The device where the service starts. The default is am
  empty string which starts the service on any device.
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
* `VOTE_HOST`: Host of the vote-service. The default is `localhost`.
* `VOTE_PORT`: Port of the vote-service. The default is `9013`.
* `VOTE_PROTOCAL`: Protocol of the vote-service. The default is `http`.
* `AUTH`: Sets the type of the auth service. `fake` (default) or `ticket`.
* `AUTH_HOST`: Host of the auth service. The default is `localhost`.
* `AUTH_PORT`: Port of the auth service. The default is `9004`.
* `AUTH_PROTOCOL`: Protocol of the auth servicer. The default is `http`.
* `OPENSLIDES_DEVELOPMENT`: If set, the service starts, even when secrets (see
  below) are not given. The default is `false`.


### Secrets

Secrets are filenames in `/run/secrets/`. The service only starts if it can find
each secret file and read its content. The default values are only used, if the
environment variable `OPENSLIDES_DEVELOPMENT` is set.

* `auth_token_key`: Key to sign the JWT auth tocken. Default `auth-dev-key`.
* `auth_cookie_key`: Key to sign the JWT auth cookie. Default `auth-dev-key`.


## Update models.yml

To use a new models.yml update the value in the file `models-version`.
Afterwards call `go generate ./...` to update the generated files.
