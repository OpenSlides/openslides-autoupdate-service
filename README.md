# OpenSlides Autoupdate Service

The Autoupdate Service is part of the OpenSlides environment. It is a http
endpoint where the clients can connect to get the actual data and also get
updates, when the requested data changes.

## Start

### With Go

```
go build ./cmd/autoupdate
./autoupdate
```

### With Docker

```
docker build . --tag autoupdate
docker run -ip 9012:9012 autoupdate
```
The i argument is important for fake key changes.

## Test

### With Go

```
go test ./...
```

### With Docker

```
docker build . -f tests/Dockerfile --tag autoupdate-test
docker run autoupdate-test
```

## Examples

### Without redis

When the server is started, clients can listen for keys to do so, they have to
send a keyrequest in the body of the request. Currently, all method-types (POST,
GET, etc) are supported. An example request

`curl localhost:9012/system/autoupdate -d '[{"ids": [5], "collection": "user", "fields": {"name": null}}]'`

To see a list of possible json-strings see the file
internal/autoupdate/keysbuilder/keysbuilder_test.go

There is a simpler method to request keys:

`curl localhost:9012/system/autoupdate/keys?user/1/name,user/2/name`

With this simpler method, it is not possible to request related keys.

After the request is send, the values to the keys are returned as a json-object
without a newline:
```
{"user/1/name": "value", "user/2/name":"value"}
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

`DATASTORE=service MESSAGING_SERVICE=redis ./autoupdate`


### With redis

When redis is installed, it can be used to update keys. Start the autoupdate
service with the envirnmentvariable `MESSAGING_SERVICE=redis`. Afterwards it is
possible to update keys by sending the following command to redis:

`xadd field_changed * updated user/5/name updated user/5/password`


## Environment

The Service uses the following environment variables:

* `LISTEN_HTTP_ADDR=:9012`: Lets the service listen on port 9012 on any device.
  The default is `:9012`.
* `DATASTORE=fake`: Sets the datastore service. `fake` (default) or `service`.
* `DATASTORE_URL`: Sets the url for the datastore service. The default is
  `http://localhost:9010`.
* `MESSAGING_SERVICE=fake`: Sets the type of messaging service. `fake`(default)
  or `redis`.
* `REDIS_ADDR=localhost:6379`: The address to redis.
* `REDIS_TEST_CONN=true`: Test the redis connection on startup. Disable on the
  cloud if redis needs more time to start then this service.
