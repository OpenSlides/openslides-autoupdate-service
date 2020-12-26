package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Group handels permissions for groups.
type Group struct {
	dp dataprovider.DataProvider
}

// NewGroup initializes a Group.
func NewGroup(dp dataprovider.DataProvider) *Group {
	return &Group{
		dp: dp,
	}
}

// Connect creates the routes.
func (g *Group) Connect(s perm.HandlerStore) {
	s.RegisterReadHandler("group", perm.ReadeCheckerFunc(g.read))
}

func (g *Group) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		fqid := fmt.Sprintf("group/%d", fqfield.ID)
		meetingID, err := g.dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
		}

		return g.dp.InMeeting(ctx, userID, meetingID)
	})
}
