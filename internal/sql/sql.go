package sql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

//go:generate sh -c "go run generate/main.go  schema.sql def.go"

type fqID struct {
	collection string
	id         int
}

// Insert creates sql insert statements from data.
func Insert(data map[dskey.Key][]byte) string {
	grouped := make(map[fqID]map[string][]byte)
	for key, value := range data {
		fqid := fqID{
			collection: key.Collection,
			id:         key.ID,
		}

		if _, ok := grouped[fqid]; !ok {
			grouped[fqid] = make(map[string][]byte)
		}

		grouped[fqid][key.Field] = value
	}

	var statements []string
	for fqid, fieldValue := range grouped {
		fieldValue["id"] = []byte(strconv.Itoa(fqid.id))

		var fields []string
		var values []string

		for field, v := range fieldValue {
			fields = append(fields, field)

			collectionField := fmt.Sprintf("%s/%s", fqid.collection, field)
			values = append(values, ValueToSQL(v, sqlTypes[collectionField]))
		}

		stmColumns := strings.Join(fields, ",")
		stmValues := strings.Join(values, ",")

		statements = append(statements, fmt.Sprintf(`INSERT INTO "%s" (%s) VALUES (%s);`, fqid.collection, stmColumns, stmValues))
	}

	return strings.Join(statements, "\n")
}

// ValueToSQL converts a json []byte value to the postgres representation.
func ValueToSQL(value []byte, sqlType string) string {
	if strings.HasSuffix(sqlType, "[]") {
		value = value[1 : len(value)-1]
		value = []byte(fmt.Sprintf("'{%s}'", value))
	}

	switch sqlType {
	case "TEXT", "VARCHAR(255)":
		value[0] = '\''
		value[len(value)-1] = '\''
		return string(value)

	default:
		return string(value)
	}
}
