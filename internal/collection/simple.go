package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// SimpleRead checks that the user has a permission.
type SimpleRead struct {
	dp         dataprovider.DataProvider
	collection string
	perm       string
}

// NewSimpleRead initializes a SimpleRead.
func NewSimpleRead(dp dataprovider.DataProvider, collection string, perm string) *SimpleRead {
	return &SimpleRead{
		dp:         dp,
		collection: collection,
		perm:       perm,
	}
}

// Connect creates the read handler.
func (r *SimpleRead) Connect(s perm.HandlerStore) {
	s.RegisterReadHandler(r.collection, perm.ReadeCheckerFunc(r.read))
}

func (r *SimpleRead) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		fqid := fmt.Sprintf("%s/%d", r.collection, fqfield.ID)
		meetingID, err := r.dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
		}

		allowed, err := perm.IsAllowed(perm.EnsurePerm(ctx, r.dp, userID, meetingID, r.perm))
		if err != nil {
			return false, fmt.Errorf("ensuring perm %w", err)
		}
		return allowed, nil
	})
}
