package core

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"
)

// Queries is a list of all possible queries.
var Queries = map[string]allowed.IsAllowed{
	"topic.create": topic.Create,
	"topic.update": topic.Update,
	"topic.delete": topic.Delete,
}
