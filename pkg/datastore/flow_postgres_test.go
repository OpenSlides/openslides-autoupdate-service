package datastore_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/jackc/pgx/v5"
	"github.com/ory/dockertest/v3"
)

func TestSourcePostgresGetSomeData(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("Postgres Test")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp, err := newTestPostgres(ctx)
	if err != nil {
		t.Fatalf("starting postgres: %v", err)
	}
	defer tp.Close()

	for _, tt := range []struct {
		name   string            // Name of the test
		data   string            // Data inserted into postgres in yaml format
		expect map[string][]byte // expected data. Uses a get request on all keys of the expect map
	}{
		{
			"Same fqid",
			`---
			user/1:
				username: hugo
				first_name: Hugo
			`,
			map[string][]byte{
				"user/1/username":   []byte(`"hugo"`),
				"user/1/first_name": []byte(`"Hugo"`),
			},
		},
		{
			"different fqid",
			`---
			user/1:
				username: hugo
				first_name: Hugo

			motion/42:
				title: antrag
				text: beschluss
			`,
			map[string][]byte{
				"user/1/username":   []byte(`"hugo"`),
				"user/1/first_name": []byte(`"Hugo"`),
				"motion/42/title":   []byte(`"antrag"`),
				"motion/42/text":    []byte(`"beschluss"`),
			},
		},
		{
			"Empty Data",
			`---
			user/1:
				username: hugo
			`,
			map[string][]byte{
				"user/1/username": []byte(`"hugo"`),
				"motion/2/title":  nil,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			source, err := datastore.NewFlowPostgres(environment.ForTests(tp.Env), nil)
			if err != nil {
				t.Fatalf("NewSource(): %v", err)
			}

			data := dsmock.YAMLData(tt.data)
			if err := tp.addTestData(ctx, data); err != nil {
				t.Fatalf("adding test data: %v", err)
			}
			defer tp.dropData(ctx)

			keys := make([]dskey.Key, 0, len(tt.data))
			for k := range tt.expect {
				keys = append(keys, dskey.MustKey(k))
			}

			got, err := source.Get(ctx, keys...)
			if err != nil {
				t.Fatalf("Get: %v", err)
			}

			expect := make(map[dskey.Key][]byte)
			for k, v := range tt.expect {
				expect[dskey.MustKey(k)] = v
			}

			if !reflect.DeepEqual(got, expect) {
				t.Errorf("\nGot\t\t%v\nexpect\t%v", got, expect)
			}
		})
	}
}

func TestBigQuery(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("Postgres Test")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp, err := newTestPostgres(ctx)
	if err != nil {
		t.Fatalf("starting postgres: %v", err)
	}
	defer tp.Close()

	source, err := datastore.NewFlowPostgres(environment.ForTests(tp.Env), nil)
	if err != nil {
		t.Fatalf("NewSource(): %v", err)
	}

	count := 2_000

	keys := make([]dskey.Key, count)
	for i := 0; i < count; i++ {
		keys[i], _ = dskey.FromParts("user", 1, fmt.Sprintf("f%d", i))
	}

	testData := make(map[dskey.Key][]byte)
	for _, key := range keys {
		testData[key] = []byte(fmt.Sprintf(`"%s"`, key.String()))
	}

	if err := tp.addTestData(ctx, testData); err != nil {
		t.Fatalf("Writing test data: %v", err)
	}

	got, err := source.Get(ctx, keys...)
	if err != nil {
		t.Errorf("Sending request with %d fields returns: %v", count, err)
	}

	if !reflect.DeepEqual(got, testData) {
		t.Errorf("testdata is diffrent then the result: for key %s got('%s') expect ('%s')", keys[1600], got[keys[1600]], testData[keys[1600]])
	}
}

type testPostgres struct {
	dockerPool     *dockertest.Pool
	dockerResource *dockertest.Resource

	Env map[string]string

	pgxConfig *pgx.ConnConfig
}

func newTestPostgres(ctx context.Context) (tp *testPostgres, err error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("connect to docker: %w", err)
	}

	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=openslides",
			"POSTGRES_DB=database",
		},
	}

	resource, err := pool.RunWithOptions(&runOpts)
	if err != nil {
		return nil, fmt.Errorf("start postgres container: %w", err)
	}

	port := resource.GetPort("5432/tcp")
	addr := fmt.Sprintf(`user=postgres password='openslides' host=localhost port=%s dbname=database`, port)
	config, err := pgx.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	tp = &testPostgres{
		dockerPool:     pool,
		dockerResource: resource,
		pgxConfig:      config,

		Env: map[string]string{
			"DATABASE_HOST": "localhost",
			"DATABASE_PORT": port,
			"DATABASE_NAME": "database",
			"DATABASE_USER": "postgres",
		},
	}

	defer func() {
		if err != nil {
			if err := tp.Close(); err != nil {
				log.Println("Closing postgres: %w", err)
			}
		}
	}()

	if err := tp.addSchema(ctx); err != nil {
		return nil, fmt.Errorf("add schema: %w", err)
	}

	return tp, nil
}

func (tp *testPostgres) Close() error {
	if err := tp.dockerPool.Purge(tp.dockerResource); err != nil {
		return fmt.Errorf("purge postgres container: %w", err)
	}
	return nil
}

func (tp *testPostgres) conn(ctx context.Context) (*pgx.Conn, error) {
	var conn *pgx.Conn

	for {
		var err error
		conn, err = pgx.ConnectConfig(ctx, tp.pgxConfig)
		if err == nil {
			return conn, nil
		}

		select {
		case <-time.After(200 * time.Millisecond):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func (tp *testPostgres) addSchema(ctx context.Context) error {
	// Schema from datastore-repo
	schema := `
	CREATE TABLE IF NOT EXISTS models (
		fqid VARCHAR(48) PRIMARY KEY,
		data JSONB NOT NULL,
		deleted BOOLEAN NOT NULL
	);`
	conn, err := tp.conn(ctx)
	if err != nil {
		return fmt.Errorf("creating connection: %w", err)
	}

	if _, err := conn.Exec(ctx, schema); err != nil {
		return fmt.Errorf("adding schema: %w", err)
	}
	return nil
}

func (tp *testPostgres) addTestData(ctx context.Context, data map[dskey.Key][]byte) error {
	objects := make(map[string]map[string]json.RawMessage)
	for k, v := range data {
		fqid := k.FQID()
		if _, ok := objects[fqid]; !ok {
			objects[fqid] = make(map[string]json.RawMessage)
		}
		objects[fqid][k.Field()] = v
	}

	conn, err := tp.conn(ctx)
	if err != nil {
		return fmt.Errorf("creating connection: %w", err)
	}

	for fqid, data := range objects {
		encoded, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("encode %v: %v", data, err)
		}

		sql := fmt.Sprintf(`INSERT INTO models (fqid, data, deleted) VALUES ('%s', '%s', false);`, fqid, encoded)
		if _, err := conn.Exec(ctx, sql); err != nil {
			return fmt.Errorf("executing psql `%s`: %w", sql, err)
		}
	}

	return nil
}

func (tp *testPostgres) dropData(ctx context.Context) error {
	conn, err := tp.conn(ctx)
	if err != nil {
		return fmt.Errorf("creating connection: %w", err)
	}

	sql := `TRUNCATE models;`
	if _, err := conn.Exec(ctx, sql); err != nil {
		return fmt.Errorf("executing psql `%s`: %w", sql, err)
	}

	return nil
}
