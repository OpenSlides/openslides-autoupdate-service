package redis_test

import (
	"context"
	"fmt"
	"testing"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/ory/dockertest/v3"
)

type testRedis struct {
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource
	addr           string

	Env map[string]string
}

func newTestRedis(t *testing.T) *testRedis {
	t.Helper()

	if testing.Short() {
		t.Skip("Redis Test")
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("redis", "6.2", nil)
	if err != nil {
		t.Fatalf("Could not start redis container: %s", err)
	}

	addr := resource.GetPort("6379/tcp")

	tr := &testRedis{
		dockerPool:     pool,
		dockerResource: resource,
		addr:           addr,
		Env: map[string]string{
			"MESSAGE_BUS_PORT": addr,
		},
	}

	return tr
}

func (tp *testRedis) Close() error {
	if err := tp.dockerPool.Purge(tp.dockerResource); err != nil {
		return fmt.Errorf("purge redis container: %w", err)
	}
	return nil
}

func (tr *testRedis) conn(ctx context.Context) (redigo.Conn, error) {
	conn, err := redigo.DialContext(ctx, "tcp", ":"+tr.addr)
	if err != nil {
		return nil, fmt.Errorf("creating test connection: %w", err)
	}

	return conn, nil
}
