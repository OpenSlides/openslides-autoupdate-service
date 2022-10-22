package main

var defaultEnvironment = map[string]string{
	"AUTOUPDATE_PORT": "9012",
	"METRIC_INTERVAL": "5m",

	// Connection to postgres
	"DATASTORE_DATABASE_HOST": "localhost",
	"DATASTORE_DATABASE_PORT": "5432",
	"DATASTORE_DATABASE_USER": "openslides",
	"DATASTORE_DATABASE_NAME": "openslides",

	// The datastore-reader is only used for history requests
	"DATASTORE_READER_HOST":              "localhost",
	"DATASTORE_READER_PORT":              "9010",
	"DATASTORE_READER_PROTOCOL":          "http",
	"DATASTORE_READER_MAX_PARALLEL_KEYS": "1000",
	"DATASTORE_READER_TIMEOUT":           "3s",

	"MESSAGE_BUS_HOST": "localhost",
	"MESSAGE_BUS_PORT": "6379",

	// Connection to the vote-service for the field poll/vote_count
	"VOTE_HOST":     "localhost",
	"VOTE_PORT":     "9013",
	"VOTE_PROTOCOL": "http",

	"AUTH":          "fake",
	"AUTH_PROTOCOL": "http",
	"AUTH_HOST":     "localhost",
	"AUTH_PORT":     "9004",

	"OPENSLIDES_DEVELOPMENT": "false",
	"SECRETS_PATH":           "/run/secrets",
}
