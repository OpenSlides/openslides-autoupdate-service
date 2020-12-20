// Package permission provides tells, if a user has the permission to see or
// write an object.
package permission

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-permission-service/internal/collection"
)

// Permission impelements the permission.Permission interface.
type Permission struct {
	connecters   []collection.Connecter
	writeHandler map[string]collection.WriteChecker
	readHandler  map[string]collection.ReadeChecker
}

// New returns a new permission service.
func New(dp DataProvider, os ...Option) *Permission {
	p := &Permission{
		writeHandler: make(map[string]collection.WriteChecker),
		readHandler:  make(map[string]collection.ReadeChecker),
	}

	for _, o := range os {
		o(p)
	}

	if p.connecters == nil {
		p.connecters = openSlidesCollections(dp)
	}

	for _, con := range p.connecters {
		con.Connect(p)
	}

	return p
}

// IsAllowed tells, if something is allowed.
func (ps *Permission) IsAllowed(ctx context.Context, name string, userID int, dataList []map[string]json.RawMessage) ([]map[string]interface{}, error) {
	handler, ok := ps.writeHandler[name]
	if !ok {
		return nil, clientError{fmt.Sprintf("unknown collection: `%s`", name)}
	}

	additions := make([]map[string]interface{}, len(dataList))
	for i, data := range dataList {
		addition, err := handler.IsAllowed(ctx, userID, data)
		if err != nil {
			return nil, isAllowedError{name: name, index: i, err: err}
		}

		additions[i] = addition
	}

	return additions, nil
}

// RestrictFQFields tells, if the given user can see the fqfields.
func (ps Permission) RestrictFQFields(ctx context.Context, userID int, fqfields []string) (map[string]bool, error) {
	grouped := groupFQFields(fqfields)

	data := make(map[string]bool, len(fqfields))

	for name, fqfields := range grouped {
		handler, ok := ps.readHandler[name]
		if !ok {
			return nil, clientError{fmt.Sprintf("unknown collection: `%s`", name)}
		}

		if err := handler.RestrictFQFields(ctx, userID, fqfields, data); err != nil {
			return nil, fmt.Errorf("restrict for collection %s: %w", name, err)
		}
	}
	return data, nil
}

// AdditionalUpdate TODO
func (ps *Permission) AdditionalUpdate(ctx context.Context, updated map[string]json.RawMessage) ([]int, error) {
	return nil, nil
}

// RegisterReadHandler registers a reader.
func (ps *Permission) RegisterReadHandler(name string, reader collection.ReadeChecker) {
	if _, ok := ps.readHandler[name]; ok {
		panic(fmt.Sprintf("Read handler with name `%s` allready exists", name))
	}
	ps.readHandler[name] = reader
}

// RegisterWriteHandler registers a writer.
func (ps *Permission) RegisterWriteHandler(name string, writer collection.WriteChecker) {
	if _, ok := ps.writeHandler[name]; ok {
		panic(fmt.Sprintf("Write handler with name `%s` allready exists", name))
	}
	ps.writeHandler[name] = writer
}

func groupFQFields(fqfields []string) map[string][]string {
	grouped := make(map[string][]string)
	for _, fqfield := range fqfields {
		parts := strings.Split(fqfield, "/")
		grouped[parts[0]] = append(grouped[parts[0]], fqfield)
	}
	return grouped
}
