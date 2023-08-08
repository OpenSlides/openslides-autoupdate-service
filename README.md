# OpenSlides Autoupdate Service

The Autoupdate Service is part of the OpenSlides environment. It is a http
endpoint where the clients can connect to get the actual data and also get
updates, when the requested data changes.

IMPORTANT: The data are sent via an open http-connection. All browsers limit the
amount of open http1.1 connections to a domain. For this service to work, the
browser has to connect to the service with http2 and therefore needs https.


## Start

The service needs some secrets to run. You can create them with:

```
mkdir secrets
printf "password" > secrets/postgres_password
printf "my_token_key" > secrets/auth_token_key 
printf "my_cookie_key" > secrets/auth_cookie_key
```

It also needs a running postgres and redis instance. You can start one with:

```
docker run  --network host -e POSTGRES_PASSWORD=password -e POSTGRES_USER=openslides -e POSTGRES_DB=openslides postgres:13
```

and

```
docker run --network host redis
```


### With Golang

```
export SECRETS_PATH=secrets
go build
./autoupdate
```


### With Docker

The docker build uses the auth token as default. Either configure it to use the
auth-fake services (see environment variables below) or make sure the service
inside the docker container can connect to the auth service. For example with
the docker argument --network host. The auth-secrets have to given as a file.

```
docker build . --tag openslides-autoupdate
docker run --network host -v $PWD/secrets:/run/secrets openslides-autoupdate
```

It uses the host network to connect to redis and postgres.


### With Auto Restart

To restart the service when ever a source file has shanged, the tool
[CompileDaemon](https://github.com/githubnemo/CompileDaemon) can help.

```
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -log-prefix=false -build "go build" -command "./autoupdate"
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
{"user/1/username":"value","user/2/username":"value"}
```

With the query parameter `single` the server writes the first response and
closes the request immediately. So there are not autoupdates:

`curl -N localhost:9012/system/autoupdate?k=user/1/username&single=1`

With the query parameter `position=XX` it is possible to request the data at a
specific position from the datastore. This implieds `single`:

`curl -N localhost:9012/system/autoupdate?k=user/1/username&position=42`


### Updates via redis

Keys are updated via redis:

`xadd ModifiedFields * user/1/username newName user/1/password newPassword`


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

### History Information

To get all history information for an fqid call:

`curl localhost:9012/system/autoupdate/history_information?fqid=motion/42`

It returns a list of all changes to the requested fqid. Each element in the list
is an object like this:

```
{
  "position": 23,
  "user_id": 5,
  "information": "motion was created",
  "timestamp: 1234567
}
```

To get the data at a position, use the normal autoupdate request with the
attribute `position`. See above.


### Internal Autoupdate

The autoupdate service provides an internal route to return fields for a defined user.

`curl "localhost:9012/internal/autoupdate?user_id=42&k=user/1/username"`

It also supports the attributes `single=1` and the normal autoupdate body.


## Configuration

The service is configurated with environment variables. See [all environment varialbes](environment.md).


## Update models.yml

To use a new models.yml update the value in the file `models-version`.
Afterwards call `go generate ./...` to update the generated files.
