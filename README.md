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
docker run --network host redis
```


### With Golang

```
export DATABASE_PASSWORD_FILE=secrets/postgres_password
export AUTH_TOKEN_KEY_FILE=secrets/auth_token_key
export AUTH_COOKIE_KEY_FILE=secrets/auth_cookie_key
go build
./openslides-autoupdate-service
```


### With Docker

Make sure the service inside the docker container can connect to the auth
service, postgres and redis, for example with the docker argument `--network
host`.

```
docker build . --tag openslides-autoupdate
docker run --network host -v $PWD/secrets:/run/secrets openslides-autoupdate
```


### With Auto Restart

To restart the service when ever a source file has shanged, the tool
[CompileDaemon](https://github.com/githubnemo/CompileDaemon) can help.

```
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -log-prefix=false -build "go build" -command "./openslides-autoupdate-service"
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

To see a list of possible json-strings see the file [internal/keysbuilder/keysbuilder_test.go](internal/keysbuilder/keysbuilder_test.go).

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


### Updates via redis

Values are updated via redis:

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


### Connection Count

The autoupdate services saves how many connections are currently open to each user.
The save interval can be defined with the environment variable
`METRIC_SAVE_INTERVAL`. The default is 5 minutes.

The values are saved for each instance of the autoupdate service. So it is
possible to access all open connection for every instance of the autoupdate service in
the same cloud.

`curl "localhost:9012/service/autoupdate/connection_count"`

It returns a JSON dictonary like this:

`{"0":15,"1":4,"2":3}`

The key is a user ID and the value is the amount of currently open connections.
User ID `0` is for connections, that are not logged in (public access). It the
example above, there are 15 open connections from public access, the user with
the ID 1 has 4 open connections and the user with the ID 2 has 3 open
connection.

Users can only access this page if they have the organization management level
or higher.


## Metric

The autoupdate service logs some metric values. The interval can be set with the
environment variable `METRIC_INTERVAL`.

The logged metric is a json dictonary like:

```json
{
    "connections_longpolling_connected_connections_public_access": 0,
    "connections_longpolling_connected_users_average_connections": 3,
    "connections_longpolling_connected_users_current": 1,
    "connections_longpolling_connected_users_current_local": 1,
    "connections_longpolling_connected_users_total": 1,
    "connections_longpolling_connected_users_total_local": 1,
    "connections_longpolling_current_connections": 3,
    "connections_longpolling_current_connections_local": 3,
    "connections_stream_connected_connections_public_access": 0,
    "connections_stream_connected_users_average_connections": 6,
    "connections_stream_connected_users_current": 2,
    "connections_stream_connected_users_current_local": 2,
    "connections_stream_connected_users_total": 3,
    "connections_stream_connected_users_total_local": 3,
    "connections_stream_current_connections": 13,
    "connections_stream_current_connections_local": 13,
    "datastore_cache_key_len": 236478,
    "datastore_cache_size": 1722987,
    "runtime_goroutines": 68
}
```

The prefix `connections_stream` are for "normal" connections.
`connections_longpolling` are for connections, that use the longpolling
fallback.


* `connections_stream_connected_connections_public_access`: Number of
  connections from non logged in users (public access) from all autoupdate instances.
* `connections_stream_connected_users_average_connections`: Average connection
  count for each logged in user (no public access).
* `connections_stream_connected_users_current`: Amount of connected users that
  have at least one open connection.
* `connections_stream_connected_users_current_local`: Amount of connected users
  that have at least one open connection of this instance.
* `connections_stream_connected_users_total`: Amount of different users that are
  currently connected or were connected since the autoupdate service was
  started.
* `connections_stream_connected_users_total_local`: Same as
`connected_users_total`, but only for this instance.
* `connections_stream_current_connections`: Amount of all connections.
* `connections_stream_current_connections_local`: Amount of all connections of this instance.
* `datastore_cache_key_len`: Amount of keys in the cache.
* `datastore_cache_size`: Combined size of all values in the cache.
* `runtime_goroutines`: Current goroutines used by the instance.


## Configuration

The service is configurated with environment variables. See [all environment varialbes](environment.md).
