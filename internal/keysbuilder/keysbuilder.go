// Package keysbuilder ...
package keysbuilder

import (
	"context"
	"fmt"
	"strconv"

	"github.com/openslides/openslides-autoupdate-service/internal/keysrequest"
)

const keySep = "/"

// Builder ...
type Builder struct {
	User    int
	Restr   Restricter
	Request keysrequest.KeysRequest
}

// Keys returns the keys
func (b *Builder) Keys() ([]string, error) {
	ids := b.Request.IDs
	if ids == nil {
		var err error
		ids, err = b.Restr.IDsFromCollection(context.TODO(), b.User, b.Request.MeetingID, b.Request.Collection)
		if err != nil {
			return nil, fmt.Errorf("can not get all ids for collection \"%s\": %w", b.Request.Collection, err)
		}
	}
	return b.run(ids, b.Request.FieldDescription)
}

func (b *Builder) run(ids []int, fd keysrequest.FieldDescription) ([]string, error) {
	out := []string{}
	for _, id := range ids {
		for field, description := range fd.Fields {
			out = append(out, buildKey(fd.Collection, id, field))
			if description.Null() {
				// field is not a reference
				continue
			}
			ids, err := b.Restr.IDsFromKey(context.TODO(), b.User, b.Request.MeetingID, buildKey(b.Request.Collection, id, field))
			if err != nil {
				return nil, err
			}
			if len(ids) == 0 {
				continue
			}
			keys, err := b.run(ids, description)
			if err != nil {
				return nil, err
			}
			out = append(out, keys...)
		}
	}
	return out, nil
}

func buildKey(collection string, id int, field string) string {
	return collection + keySep + strconv.Itoa(id) + keySep + field
}
