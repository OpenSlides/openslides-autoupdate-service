package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"text/template"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/models"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/sql"
)

func main() {
	schemaFile, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer schemaFile.Close()

	defFile, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer defFile.Close()

	if err := run(schemaFile, defFile); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run(schemaOut io.Writer, defOut io.Writer) error {
	tables, fqidTypes, err := parse()
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	if err := writeSchema(schemaOut, tables); err != nil {
		return fmt.Errorf("wirite schema: %w", err)
	}

	if err := writeDef(defOut, fqidTypes); err != nil {
		return fmt.Errorf("wirite type defs: %w", err)
	}

	return nil
}

type sqlTable struct {
	TableName string
	Columns   []sqlColumn
}

type sqlColumn struct {
	Name      string
	SQLType   string
	Constrain string
}

func parse() ([]sqlTable, map[string]string, error) {
	modelsContent, err := loadDefition()
	if err != nil {
		return nil, nil, fmt.Errorf("Can not load models defition: %w", err)
	}
	defer modelsContent.Close()

	collections, err := models.Unmarshal(modelsContent)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	tables := make([]sqlTable, 0, len(collections))
	fqidTypes := make(map[string]string)
	for collectionName, collection := range collections {
		table := sqlTable{
			TableName: collectionName,
			Columns:   make([]sqlColumn, 0, len(collection.Fields)),
		}

		for fieldName, field := range collection.Fields {
			sqltype, constrain, err := modelsTypeToSQLType(fieldName, field)
			if err != nil {
				return nil, nil, fmt.Errorf("converting type to SQL: %w", err)
			}

			table.Columns = append(table.Columns, sqlColumn{fieldName, sqltype, constrain})
			fqidTypes[collectionName+"/"+fieldName] = sqltype
		}

		tables = append(tables, table)
	}

	return tables, fqidTypes, nil
}

func loadDefition() (io.ReadCloser, error) {
	r, err := http.Get(models.URLModelsYML())
	if err != nil {
		return nil, fmt.Errorf("request defition: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned status %s", r.Status)
	}
	return r.Body, nil
}

func modelsTypeToSQLType(fieldName string, field *models.Field) (string, string, error) {
	if fieldName == "id" {
		return "INTEGER", "PRIMARY KEY", nil
	}

	var sqlType string
	switch field.Type {
	case "number", "relation", "timestamp":
		sqlType = "INTEGER"

	case "string", "color", "generic-relation", "decimal(6)":
		sqlType = "VARCHAR(255)"

	case "text", "HTMLStrict", "HTMLPermissive", "JSON":
		sqlType = "TEXT"

	case "boolean":
		sqlType = "BOOLEAN"

	case "float":
		sqlType = "REAL"

	case "relation-list", "number[]":
		sqlType = "INTEGER[]"

	case "string[]", "generic-relation-list":
		sqlType = "TEXT[]"

	default:
		return "", "", fmt.Errorf("Unknown type %q", field.Type)
	}

	constrain := ""
	if field.Required {
		constrain = "NOT Null"
	}

	if field.Default != nil {
		bs, err := json.Marshal(field.Default)
		if err != nil {
			return "", "", fmt.Errorf("converting default to json: %w", err)
		}

		constrain += fmt.Sprintf(" DEFAULT %s", sql.ValueToSQL(bs, sqlType))
	}

	return sqlType, constrain, nil
}

const tmplCreateTable = `
CREATE TABLE "{{.TableName}}" (
    {{- range $i, $column := .Columns}}
    {{$column.Name}} {{$column.SQLType}} {{$column.Constrain}}{{if last $i $.Columns | not}},{{end}}
    {{- end}}
);
`

var fns = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}

func writeSchema(w io.Writer, tables []sqlTable) error {
	tmpl, err := template.New("create_table.sql").Funcs(fns).Parse(tmplCreateTable)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	for _, table := range tables {
		if err := tmpl.Execute(w, table); err != nil {
			return fmt.Errorf("executing template: %w", err)
		}
	}

	return nil
}

const tmplDef = `// Code generated with models.yml DO NOT EDIT.
package sql

// sqlTypes is a map from fqid to sql type
var sqlTypes = map[string]string{
	{{- range $fqid, $sqltype := .}}
		"{{$fqid}}": "{{$sqltype}}",
	{{- end}}
}
`

func writeDef(w io.Writer, types map[string]string) error {
	tmpl, err := template.New("def.go").Parse(tmplDef)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, types); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	formated, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formating code: %w", err)
	}

	if _, err := w.Write(formated); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}
