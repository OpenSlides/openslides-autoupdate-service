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
	ds    *datastore.Request
	uid   int
}

// NewMeetingPermission initializes a new MeetingPermission.
func NewMeetingPermission(ds *datastore.Request, uid int) *MeetingPermission {
	p := MeetingPermission{
		perms: make(map[int]*Permission),
		ds:    ds,
		uid:   uid,
	}
	return &p
}

// Meeting returns the permission object for the meeting.
func (p MeetingPermission) Meeting(ctx context.Context, meetingID int) (*Permission, error) {
	perms, ok := p.perms[meetingID]
	if ok {
		return perms, nil
	}

	perms, err := New(ctx, p.ds, p.uid, meetingID)
	if err != nil {
		return nil, err
	}
	return perms, nil
}

// UserID returns the user id the object was initialized with.
func (p MeetingPermission) UserID() int {
	return p.uid
}
