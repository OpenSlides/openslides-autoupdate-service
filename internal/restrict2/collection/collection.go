package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Restricter returns a fieldRestricter for a restriction_mode.
//
// The FieldRestricter is a function that tells, if a user can see fields in
// that mode.
type Restricter interface {
	Modes(mode string) FieldRestricter

	// MeetingID returns the meeting id for an object. Returns hasMeeting=false,
	// if the object does not belong to a meeting.
	MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (meetingID int, hasMeeting bool, err error)

	Name() string
}

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, ds *dsfetch.Fetch, attr *attribute.Map, agendaIDs ...int) error

func meetingPerm(ctx context.Context, ds *dsfetch.Fetch, r Restricter, mode string, ids []int, permission perm.TPermission, attrMap *attribute.Map) error {
	return mapMeeting(ctx, ds, r, ids, func(meetingID int, ids []int) error {
		groupMap, err := perm.GroupMapFromContext(ctx, ds, meetingID)
		if err != nil {
			return fmt.Errorf("getting permission: %w", err)
		}

		attr := attribute.Attribute{
			GlobalPermission: attribute.GlobalFromPerm(perm.OMLSuperadmin),
			GroupIDs:         groupMap[permission],
		}

		for _, id := range ids {
			attrMap.Add(dskey.Key{Collection: r.Name(), ID: id, Field: mode}, &attr)
		}
		return nil
	})
}

func mapMeeting(ctx context.Context, ds *dsfetch.Fetch, r Restricter, ids []int, fn func(meetingID int, ids []int) error) error {
	meetingToIDs := make(map[int][]int)
	for _, id := range ids {
		meetingID, hasMeeting, err := r.MeetingID(ctx, ds, id)
		if err != nil {
			return fmt.Errorf("getting meeting id of element %d: %w", id, err)
		}

		if !hasMeeting || meetingID == 0 {
			return fmt.Errorf("element with id %d has no meeting", id)
		}

		meetingToIDs[meetingID] = append(meetingToIDs[meetingID], id)
	}

	for meetingID, ids := range meetingToIDs {
		if err := fn(meetingID, ids); err != nil {
			return fmt.Errorf("restricting for meeting %d: %w", meetingID, err)
		}
	}

	return nil
}
