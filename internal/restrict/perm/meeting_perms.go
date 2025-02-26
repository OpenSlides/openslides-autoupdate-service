package perm

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// MeetingPermission is a cache for different Permission objects for each
// meeting.
//
// Can be used if fields from different meetings are checked.
type meetingPermission struct {
	perms map[int]*Permission
	ds    *dsfetch.Fetch
	uid   int
}

// NewMeetingPermission initializes a new MeetingPermission.
func newMeetingPermission(ds *dsfetch.Fetch, uid int) *meetingPermission {
	p := meetingPermission{
		perms: make(map[int]*Permission),
		ds:    ds,
		uid:   uid,
	}
	return &p
}

// Meeting returns the permission object for the meeting.
func (p *meetingPermission) Meeting(ctx context.Context, meetingID int) (*Permission, error) {
	perms, ok := p.perms[meetingID]
	if ok {
		return perms, nil
	}

	perms, err := New(ctx, p.ds, p.uid, meetingID)
	if err != nil {
		return nil, err
	}
	p.perms[meetingID] = perms
	return perms, nil
}

// UserID returns the user id the object was initialized with.
func (p *meetingPermission) UserID() int {
	return p.uid
}

type contextKeyType string

const contextKey contextKeyType = "meeting_permission"

// ContextWithPermissionCache adds a permission cache to the context.
func ContextWithPermissionCache(ctx context.Context, getter datastore.Getter, uid int) context.Context {
	fetcher := dsfetch.New(getter)
	return context.WithValue(ctx, contextKey, newMeetingPermission(fetcher, uid))
}

// FromContext gets a meeting specific permission object from a context.
//
// Make sure to generate the context with 'ContextWithPermissionCache.
func FromContext(ctx context.Context, meetingID int) (*Permission, error) {
	v := ctx.Value(contextKey)
	if v == nil {
		return nil, fmt.Errorf("context does not contain a meeting permission. Make sure to create the context with 'ContextWithPermissionCache'")
	}

	meetingPermission, ok := v.(*meetingPermission)
	if !ok {
		return nil, fmt.Errorf("meeting permission has wrong type: %T", v)
	}

	return meetingPermission.Meeting(ctx, meetingID)
}

// RequestUserFromContext returns the request user from the context.
//
// Make sure to generate the context with 'ContextWithPermissionCache.
func RequestUserFromContext(ctx context.Context) (int, error) {
	v := ctx.Value(contextKey)
	if v == nil {
		return 0, fmt.Errorf("context does not contain a meeting permission. Make sure to create the context with 'ContextWithPermissionCache'")
	}

	meetingPermission, ok := v.(*meetingPermission)
	if !ok {
		return 0, fmt.Errorf("meeting permission has wrong type: %T", v)
	}

	return meetingPermission.uid, nil
}
