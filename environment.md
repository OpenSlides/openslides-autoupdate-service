<!--- Code generated with go generate ./... DO NOT EDIT. --->
# Configuration

## Environment Variables

The Service uses the following environment variables:

* `AUTOUPDATE_PORT`: Port on which the service listen on. The default is `9012`.
* `MESSAGE_BUS_HOST`: Host of the redis server. The default is `localhost`.
* `MESSAGE_BUS_PORT`: Port of the redis server. The default is `6379`.
* `VOTE_PROTOCOL`: Protocol of the vote-service. The default is `http`.
* `VOTE_HOST`: Host of the vote-service. The default is `localhost`.
* `VOTE_PORT`: Port of the vote-service. The default is `9013`.
* `DATASTORE_READER_PROTOCOL`: Protocol of the datastore reader. The default is `http`.
* `DATASTORE_READER_HOST`: Host of the datastore reader. The default is `localhost`.
* `DATASTORE_READER_PORT`: Port of the datastore reader. The default is `9010`.
* `DATASTORE_TIMEOUT`: Time until a request to the datastore times out. The default is `3s`.
* `DATASTORE_MAX_PARALLEL_KEYS`: Max keys that are send in one request to the datastore. The default is `1000`.
* `DATASTORE_DATABASE_USER`: Postgres User. The default is `openslides`.
* `DATASTORE_DATABASE_HOST`: Postgres Host. The default is `localhost`.
* `DATASTORE_DATABASE_PORT`: Postgres Post. The default is `5432`.
* `DATASTORE_DATABASE_NAME`: Postgres Database. The default is `openslides`.
* `OPENSLIDES_DEVELOPMENT`: If set, the service uses the default secrets. The default is `false`.
* `SECRETS_PATH`: Path where the secrets are stored. The default is `/run/secrets`.
* `AUTH_PROTOCOL`: Protocol of the auth service. The default is `http`.
* `AUTH_HOST`: Host of the auth service. The default is `localhost`.
* `AUTH_PORT`: Port of the auth service. The default is `9004`.
* `AUTH_Fake`: Use user id 1 for every request. Ignores all other auth environment variables. The default is `false`.
* `CONCURENT_WORKER`: Amount of clients that calculate there values at the same time. Default to GOMAXPROCS. The default is `0`.
* `METRIC_INTERVAL`: Time in how often the metrics are gathered. Zero disables the metrics. The default is `5m`.


## Secrets

Secrets are filenames in the directory `SECRETS_PATH` (default: `/run/secrets/`). 
The service only starts if it can find each secret file and read its content. 
The default values are only used, if the environment variable `OPENSLIDES_DEVELOPMENT` is set.

* `postgres_password`: Postgres Password. The default is `openslides`.
* `auth_token_key`: Key to sign the JWT auth tocken. The default is `auth-dev-token-key`.
* `auth_cookie_key`: Key to sign the JWT auth cookie. The default is `auth-dev-cookie-key`.
