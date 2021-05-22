package collection

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

func canSeeAssignmentCandidate(p *perm.Permission) bool {
	return p.Has(perm.AssignmentCanSee)
}
