package redis

// Connection is the raw connection to a redis server.
type Connection interface {
	XREAD(count, stream, lastID string) (interface{}, error)
}
