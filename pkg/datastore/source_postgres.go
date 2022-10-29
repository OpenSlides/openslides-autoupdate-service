package datastore

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	envPostgresHost     = environment.NewVariable("DATASTORE_DATABASE_HOST", "localhost", "Postgres Host.")
	envPostgresPort     = environment.NewVariable("DATASTORE_DATABASE_PORT", "5432", "Postgres Post.")
	envPostgresUser     = environment.NewVariable("DATASTORE_DATABASE_USER", "openslides", "Postgres User.")
	envPostgresDatabase = environment.NewVariable("DATASTORE_DATABASE_NAME", "openslides", "Postgres Database.")
	envPostgresPassword = environment.NewSecret("postgres_password", "Postgres Password.")
)

// SourcePostgres uses postgres to get the connections.
//
// TODO: This should be unexported, but there is an import cycle in the tests.
type SourcePostgres struct {
	pool    *pgxpool.Pool
	updater Updater
}

// NewSourcePostgres initializes a SourcePostgres.
//
// TODO: This should be unexported, but there is an import cycle in the tests.
func NewSourcePostgres(lookup environment.Environmenter, updater Updater) (*SourcePostgres, error) {
	addr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		envPostgresUser.Value(lookup),
		envPostgresPassword.Value(lookup),
		envPostgresHost.Value(lookup),
		envPostgresPort.Value(lookup),
		envPostgresDatabase.Value(lookup),
	)

	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	source := SourcePostgres{pool: pool, updater: updater}

	return &source, nil
}

// Get fetches the keys from postgres.
func (p *SourcePostgres) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	uniqueFieldsStr, fieldIndex, uniqueFQID := prepareQuery(keys)

	// For very big SQL Queries, split them in part
	if len(fieldIndex) > maxFieldsOnQuery {
		keysList := splitFieldKeys(keys)
		result := make(map[dskey.Key][]byte, len(keys))
		for _, keys := range keysList {
			resultPart, err := p.Get(ctx, keys...)
			if err != nil {
				return nil, fmt.Errorf("get key list: %w", err)
			}

			for k, v := range resultPart {
				result[k] = v
			}
		}
		return result, nil
	}

	sql := fmt.Sprintf(`SELECT fqid, %s from models where fqid = ANY ($1) AND deleted=false;`, uniqueFieldsStr)

	rows, err := p.pool.Query(ctx, sql, uniqueFQID)
	if err != nil {
		return nil, fmt.Errorf("sending query: %w", err)
	}
	defer rows.Close()

	table := make(map[string][][]byte)
	for rows.Next() {
		r := rows.RawValues()
		copied := make([][]byte, len(r)-1)
		for i := 1; i < len(r); i++ {
			copied[i-1] = r[i]
			if r[i] == nil {
				continue
			}
			copied[i-1] = make([]byte, len(r[i]))
			copy(copied[i-1], r[i])
		}
		table[string(r[0])] = copied
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("reading postgres result: %w", rows.Err())
	}

	values := make(map[dskey.Key][]byte, len(keys))
	for _, k := range keys {
		var value []byte
		element, ok := table[k.FQID()]
		if ok {
			idx, ok := fieldIndex[k.Field]
			if ok {
				value = element[idx]
			}
		}
		values[k] = value
	}

	return values, nil
}

// Update calls the updater.
func (p *SourcePostgres) Update(ctx context.Context) (map[dskey.Key][]byte, error) {
	return p.updater.Update(ctx)
}

func prepareQuery(keys []dskey.Key) (uniqueFieldsStr string, fieldIndex map[string]int, uniqueFQID []string) {
	uniqueFQIDSet := make(map[string]struct{})
	uniqueFieldsSet := make(map[string]struct{})
	for _, k := range keys {
		uniqueFieldsSet[k.Field] = struct{}{}
		uniqueFQIDSet[k.FQID()] = struct{}{}
	}

	uniqueFields := make([]string, 0, len(uniqueFieldsSet))
	fieldIndex = make(map[string]int, len(uniqueFieldsSet))
	for field := range uniqueFieldsSet {
		uniqueFields = append(uniqueFields, fmt.Sprintf("data->'%s'", field))
		fieldIndex[field] = len(uniqueFields) - 1
	}

	uniqueFQID = make([]string, 0, len(uniqueFQIDSet))
	for k := range uniqueFQIDSet {
		uniqueFQID = append(uniqueFQID, k)
	}
	uniqueFieldsStr = strings.Join(uniqueFields, ",")
	return uniqueFieldsStr, fieldIndex, uniqueFQID
}

const maxFieldsOnQuery = 1_500

// splitFieldKeys splits a list of keys to many lists where any list has a
// maximum of different fields.
func splitFieldKeys(keys []dskey.Key) [][]dskey.Key {
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Field < keys[j].Field
	})

	var out [][]dskey.Key
	keyCount := 0
	var nextList []dskey.Key
	var lastField string
	for _, k := range keys {
		nextList = append(nextList, k)

		if k.Field != lastField {
			keyCount++
			if keyCount >= maxFieldsOnQuery {
				out = append(out, nextList)
				nextList = nil
				keyCount = 0
			}
		}
		lastField = k.Field
	}
	out = append(out, nextList)

	return out
}
