package autoupdate

import (
	"github.com/cespare/xxhash/v2"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysbuilder"
)

// Connection holds the state of an open connection to the autoupdate system
type Connection struct {
	user int
	tid  uint64
	b    *keysbuilder.Builder
	data map[string]uint64
}

// filter removes values from data, that are the same as before
func (c *Connection) filter(data map[string][]byte) {
	if c.data == nil {
		c.data = make(map[string]uint64)
	}
	for key, value := range data {
		new := xxhash.Sum64(value)
		old, ok := c.data[key]
		if ok && old == new {
			delete(data, key)
			continue
		}

		c.data[key] = new
	}
}
