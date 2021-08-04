package test

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MeetingPermissionStub lets you define permissions for meetings.
type MeetingPermissionStub struct {
	UID         int
	Permissions map[int][]perm.TPermission
}

// Meeting returns the permission for one meeting.
func (mps MeetingPermissionStub) Meeting(ctx context.Context, id int) (perm.Permission, error) {
	ps, ok := mps.Permissions[id]
	if !ok {
		return nil, datastore.DoesNotExistError(fmt.Sprintf("meeting/%d", id))
	}

	return PermissionStub{
		Perms:  ps,
		Admin:  false,
		Groups: nil,
	}, nil
}

// UserID returns the User id the stub was created with.
func (mps MeetingPermissionStub) UserID() int {
	return mps.UID
}

// PermissionStub defines permissions for one meeting.
type PermissionStub struct {
	Perms  []perm.TPermission
	Admin  bool
	Groups []int
}

// Has tels, if the given permission was defined.
func (ps PermissionStub) Has(prm perm.TPermission) bool {
	for _, p := range ps.Perms {
		if p == prm {
			return true
		}
	}
	return false
}

// IsAdmin returns, if the stub was initialized with admin.
// TODO: is this realy needed?
func (ps PermissionStub) IsAdmin() bool {
	return ps.Admin
}

// InGroup tells, if the given gid is in the list of groupIDs.
func (ps PermissionStub) InGroup(gid int) bool {
	for _, g := range ps.Groups {
		if g == gid {
			return true
		}
	}
	return false
}
