package datastore

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4"
)

// SourcePostgres uses postgres to get the connections.
type SourcePostgres struct {
	config *pgx.ConnConfig
	SourcePosition
}

// NewSourcePostgres initializes a SourcePostgres.
func NewSourcePostgres(ctx context.Context, addr string, password string, positioner SourcePosition) (*SourcePostgres, error) {
	config, err := pgx.ParseConfig(addr)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.Password = password
	config.PreferSimpleProtocol = true

	return &SourcePostgres{config: config, SourcePosition: positioner}, nil
}

// Get fetches the keys from postgres.
func (p *SourcePostgres) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	conn, err := pgx.ConnectConfig(ctx, p.config)
	if err != nil {
		return nil, fmt.Errorf("connecting to postgres: %w", err)
	}
	defer conn.Close(ctx)

	uniqueFieldsStr, fieldIndex, uniqueFQID := prepareQuery(keys)

	sql := fmt.Sprintf(`SELECT fqid, %s from models where fqid = ANY ($1) AND deleted=false;`, uniqueFieldsStr)

	rows, err := conn.Query(ctx, sql, uniqueFQID)
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
