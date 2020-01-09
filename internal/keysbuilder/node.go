package keysbuilder

import "github.com/openslides/openslides-autoupdate-service/internal/keysrequest"

type node struct {
	name   string
	fd     keysrequest.FieldDescription
	fields []*node
}

func (n *node) run() []string {
	if n == nil {
		return nil
	}
	out := make([]string, 0)
	for _, field := range n.fields {
		out = append(out, field.name)
		out = append(out, field.run()...)
	}
	return out
}
