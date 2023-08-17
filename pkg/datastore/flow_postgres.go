package datastore

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const maxFieldsOnQuery = 1_500

var (
	envPostgresHost         = environment.NewVariable("DATABASE_HOST", "localhost", "Postgres Host.")
	envPostgresPort         = environment.NewVariable("DATABASE_PORT", "5432", "Postgres Post.")
	envPostgresDatabase     = environment.NewVariable("DATABASE_NAME", "openslides", "Postgres User.")
	envPostgresUser         = environment.NewVariable("DATABASE_USER", "openslides", "Postgres Database.")
	envPostgresPasswordFile = environment.NewVariable("DATABASE_PASSWORD_FILE", "/run/secrets/postgres_password", "Postgres Password.")
)

// FlowPostgres uses postgres to get the connections.
type FlowPostgres struct {
	pool    *pgxpool.Pool
	updater flow.Updater
}

// encodePostgresConfig encodes a string to be used in the postgres key value style.
//
// See: https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-CONNSTRING
func encodePostgresConfig(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `'`, `\'`)
	return s
}

// NewFlowPostgres initializes a SourcePostgres.
//
// TODO: This should be unexported, but there is an import cycle in the tests.
func NewFlowPostgres(lookup environment.Environmenter, updater flow.Updater) (*FlowPostgres, error) {
	password, err := environment.ReadSecret(lookup, envPostgresPasswordFile)
	if err != nil {
		return nil, fmt.Errorf("reading postgres password: %w", err)
	}

	addr := fmt.Sprintf(
		`user='%s' password='%s' host='%s' port='%s' dbname='%s'`,
		encodePostgresConfig(envPostgresUser.Value(lookup)),
		encodePostgresConfig(password),
		encodePostgresConfig(envPostgresHost.Value(lookup)),
		encodePostgresConfig(envPostgresPort.Value(lookup)),
		encodePostgresConfig(envPostgresDatabase.Value(lookup)),
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

	flow := FlowPostgres{pool: pool, updater: updater}

	return &flow, nil
}

// Get fetches the keys from postgres.
func (p *FlowPostgres) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
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
			idx, ok := fieldIndex[k.Field()]
			if ok {
				value = element[idx]
			}
		}

		if string(value) == "null" {
			value = nil
		}

		values[k] = value
	}

	return values, nil
}

// Update calls the updater.
func (p *FlowPostgres) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	p.updater.Update(ctx, updateFn)
}

func prepareQuery(keys []dskey.Key) (uniqueFieldsStr string, fieldIndex map[string]int, uniqueFQID []string) {
	uniqueFQIDSet := make(map[string]struct{})
	uniqueFieldsSet := make(map[string]struct{})
	for _, k := range keys {
		uniqueFieldsSet[k.Field()] = struct{}{}
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

// splitFieldKeys splits a list of keys to many lists where any list has a
// maximum of different fields.
func splitFieldKeys(keys []dskey.Key) [][]dskey.Key {
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Field() < keys[j].Field()
	})

	var out [][]dskey.Key
	keyCount := 0
	var nextList []dskey.Key
	var lastField string
	for _, k := range keys {
		nextList = append(nextList, k)

		if k.Field() != lastField {
			keyCount++
			if keyCount >= maxFieldsOnQuery {
				out = append(out, nextList)
				nextList = nil
				keyCount = 0
			}
		}
		lastField = k.Field()
	}
	out = append(out, nextList)

	return out
}
