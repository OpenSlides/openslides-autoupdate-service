package perm

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MeetingPermission is a cache for different Permission objects for each
// meeting.
//
// Can be used if fields from different meetings are checked.
type MeetingPermission struct {
	perms map[int]*Permission
	fetch *datastore.Fetcher
	uid   int
}

// NewMeetingPermission initializes a new MeetingPermission.
func NewMeetingPermission(fetch *datastore.Fetcher, uid int) *MeetingPermission {
	p := MeetingPermission{
		perms: make(map[int]*Permission),
		fetch: fetch,
		uid:   uid,
	}
	return &p
}

// Meeting returns the permission object for the meeting.
func (p *MeetingPermission) Meeting(ctx context.Context, id int) (*Permission, error) {
	perms, ok := p.perms[id]
	if ok {
		return perms, nil
	}

	perms, err := New(ctx, p.fetch, p.uid, id)
	if err != nil {
		return nil, err
	}
	return perms, nil
}
