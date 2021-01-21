package collection

import "github.com/OpenSlides/openslides-permission-service/internal/perm"

func canSeeProjection(p *perm.Permission) bool {
	return p.Has("projector.can_see")
}

func canSeeProjector(p *perm.Permission) bool {
	return p.Has("projector.can_see")
}
