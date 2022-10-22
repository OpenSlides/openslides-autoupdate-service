package datastore

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SourcePostgres uses postgres to get the connections.
type SourcePostgres struct {
	pool *pgxpool.Pool
	SourcePosition
}

// NewSourcePostgres initializes a SourcePostgres.
func NewSourcePostgres(ctx context.Context, addr string, password string, positioner SourcePosition) (*SourcePostgres, error) {
	config, err := pgxpool.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.ConnConfig.Password = password
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("creating connection pool: %w", err)
	}

	return &SourcePostgres{pool: pool, SourcePosition: positioner}, nil
}

// Get fetches the keys from postgres.
func (p *SourcePostgres) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	uniqueFieldsStr, fieldIndex, uniqueFQID := prepareQuery(keys)

	// For very big SQL Queries, split them in part
	if len(fieldIndex) > maxFieldsOnQuery {
		keysList := splitFieldKeys(keys)
		result := make(map[Key][]byte, len(keys))
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

	values := make(map[Key][]byte, len(keys))
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

func prepareQuery(keys []Key) (uniqueFieldsStr string, fieldIndex map[string]int, uniqueFQID []string) {
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
func splitFieldKeys(keys []Key) [][]Key {
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Field < keys[j].Field
	})

	var out [][]Key
	keyCount := 0
	var nextList []Key
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
