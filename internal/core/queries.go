package core

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"
)

var Queries = map[string]allowed.IsAllowed{
	"topic.create": topic.Create,
	"topic.update": topic.Update,
	"topic.delete": topic.Delete,
}
