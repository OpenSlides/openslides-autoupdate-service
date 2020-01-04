# OpenSlides Autoupdate Service

The Autoupdate Service is part of the OpenSlides environment. It is a
http endpoint where the clients can connect to get the actual data and
also get updates, when the requested data changes.

## Start

### With Go

```
go build ./cmd/autoupdate
./autoupdate
```

### With Docker

```
docker build . --tag autoupdate
docker run -ip 8002:8002 autoupdate
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

When the server is started, clients can listen for keys to do so, they have to send a keyrequest in the body
of the request. Currently, all method-types are supported. An example request

`curl localhost:8002/autoupdate/ -d '{"ids": [5], "collection": "user", "fields": {"name": null}}'`

To see a list of possible json-strings see the file internal/keysbuilder/keysbuilder_test.go

After the request is send, the values to the keys are returned in the form
```
key1: value
key2: value
key3: value
```

To "update" keys, you can send them to the server via stdin in the form:
```
user/5/name
user/6/name=Emanuel
```

All clients that listen for the keys get an update in the same form then the initial form.


## Environment

The Service uses the following environment variables:

* `LISTEN_HTTP_ADDR=:8080`: Lets the service listen on port 8080 on any device. The default is `:8002`.
* `MESSAGIN_SERVICE=fake`: Tells the service what kind of messagin service is used. Currently, there is only the default: `fake`.
* `AUTH_SERVICE=fake`: The same for the auth service.
* `RESTRICTER_SERVICE=fakse`: The same for the restricter service.

