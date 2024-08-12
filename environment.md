<!--- Code generated with go generate ./... DO NOT EDIT. --->
# Configuration

## Environment Variables

The Service uses the following environment variables:

* `AUTOUPDATE_PORT`: Port on which the service listen on. The default is `9012`.
* `MESSAGE_BUS_HOST`: Host of the redis server. The default is `localhost`.
* `MESSAGE_BUS_PORT`: Port of the redis server. The default is `6379`.
* `ANONYMOUS_ONLY`: Start for only anonymous users. Does not write to redis or connect to the vote-service. The default is `false`.
* `OPENSLIDES_DEVELOPMENT`: If set, the service uses the default secrets. The default is `false`.
* `DATABASE_PASSWORD_FILE`: Postgres Password. The default is `/run/secrets/postgres_password`.
* `DATABASE_USER`: Postgres Database. The default is `openslides`.
* `DATABASE_HOST`: Postgres Host. The default is `localhost`.
* `DATABASE_PORT`: Postgres Post. The default is `5432`.
* `DATABASE_NAME`: Postgres User. The default is `openslides`.
* `VOTE_PROTOCOL`: Protocol of the vote-service. The default is `http`.
* `VOTE_HOST`: Host of the vote-service. The default is `localhost`.
* `VOTE_PORT`: Port of the vote-service. The default is `9013`.
* `AUTH_PROTOCOL`: Protocol of the auth service. The default is `http`.
* `AUTH_HOST`: Host of the auth service. The default is `localhost`.
* `AUTH_PORT`: Port of the auth service. The default is `9004`.
* `AUTH_FAKE`: Use user id 1 for every request. Ignores all other auth environment variables. The default is `false`.
* `AUTH_TOKEN_KEY_FILE`: Key to sign the JWT auth tocken. The default is `/run/secrets/auth_token_key`.
* `AUTH_COOKIE_KEY_FILE`: Key to sign the JWT auth cookie. The default is `/run/secrets/auth_cookie_key`.
* `CONCURENT_WORKER`: Amount of clients that calculate there values at the same time. Default to GOMAXPROCS. The default is `0`.
* `CACHE_RESET`: Time to reset the cache. The default is `24h`.
* `METRIC_INTERVAL`: Time in how often the metrics are gathered. Zero disables the metrics. The default is `5m`.
* `METRIC_SAVE_INTERVAL`: Interval, how often the metric should be saved to redis. Redis will ignore entries, that are twice at old then the save interval. The default is `5m`.
* `DISABLE_CONNECTION_COUNT`: Do not count connections. The default is `false`.
