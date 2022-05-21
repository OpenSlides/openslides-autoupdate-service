package datastore

import (
	"context"
	"fmt"
	"os"
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
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	config.Password = password

	return &SourcePostgres{config: config, SourcePosition: positioner}, nil
}

// Get fetches the keys from postgres.
func (p *SourcePostgres) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	conn, err := pgx.ConnectConfig(ctx, p.config)
	if err != nil {
		return nil, fmt.Errorf("connecting to postgres: %w", err)
	}
	defer conn.Close(ctx)

	uniqueFQIDSet := make(map[string]struct{})
	uniqueFieldsSet := make(map[string]struct{})
	for _, k := range keys {
		uniqueFieldsSet[fmt.Sprintf("data->%s as %s", k.Field, k.Field)] = struct{}{}
		uniqueFQIDSet[k.FQID()] = struct{}{}
	}

	uniqueFields := make([]string, 0, len(uniqueFieldsSet))
	fieldIndex := make(map[string]int, len(uniqueFieldsSet))
	for k := range uniqueFieldsSet {
		uniqueFields = append(uniqueFields, k)
		fieldIndex[k] = len(uniqueFields) - 1
	}

	uniqueFQID := make([]string, 0, len(uniqueFQIDSet))
	for k := range uniqueFQIDSet {
		uniqueFQID = append(uniqueFQID, k)
	}
	uniqueFieldsStr := strings.Join(uniqueFields, ",")

	sql := fmt.Sprintf(`SELECT fqfield, %s from models where fqid in $1 AND deleted=false`, uniqueFieldsStr)

	rows, err := conn.Query(ctx, sql, uniqueFQID)
	if err != nil {
		return nil, fmt.Errorf("sending query: %w", err)
	}
	defer rows.Close()

	table := make(map[string][][]byte)
	for rows.Next() {
		r := rows.RawValues()
		table[string(r[0])] = r[1:]
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("reading postgres result: %w", rows.Err())
	}

	values := make(map[Key][]byte, len(keys))
	for _, k := range keys {
		var value []byte
		fields, ok := table[k.FQID()]
		if ok {
			value = fields[fieldIndex[k.Field]]
		}
		values[k] = value
	}

	return values, nil
}
