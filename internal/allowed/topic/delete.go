package topic

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// Delete TODO
var Delete = allowed.BuildModify([]string{"id"}, "topic", "agenda.can_manage")
