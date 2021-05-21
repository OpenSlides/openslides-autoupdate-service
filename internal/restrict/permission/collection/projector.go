package collection

import "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"

func canSeeProjection(p *perm.Permission) bool {
	return p.Has(perm.ProjectorCanSee)
}

func canSeeProjector(p *perm.Permission) bool {
	return p.Has(perm.ProjectorCanSee)
}
