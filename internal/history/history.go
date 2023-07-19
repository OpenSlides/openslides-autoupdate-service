package history

import (
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

var reValidKeys = regexp.MustCompile(`^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*`)

// Datastore is the source for the data.
type Datastore interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
	GetPosition(ctx context.Context, position int, keys ...dskey.Key) (map[dskey.Key][]byte, error)
	HistoryInformation(ctx context.Context, fqid string, w io.Writer) error
}

// History get access to old data int the datastore
type History struct {
	source *sourceDatastore
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update(ctx context.Context, ds flow.Getter) ([]dskey.Key, error)
}

// New initializes a new history.
func New(lookup environment.Environmenter) (*History, error) {
	source, err := newSourceDatastore(lookup)
	if err != nil {
		return nil, fmt.Errorf("initializing datastore: %w", err)
	}
	return &History{
		source: source,
	}, nil
}

// HistoryInformation returns the histrory information for an fqid.
func (h History) HistoryInformation(ctx context.Context, uid int, fqid string, w io.Writer) error {
	if !reValidKeys.MatchString(fqid) {
		// TODO Client Error
		return invalidInputError{fmt.Sprintf("fqid %s is invalid", fqid)}
	}

	coll, rawID, _ := strings.Cut(fqid, "/")
	id, _ := strconv.Atoi(rawID)

	ds := dsfetch.New(h.source)

	meetingID, hasMeeting, err := collection.Collection(ctx, coll).MeetingID(ctx, ds, id)
	if err != nil {
		var errNotExist dsfetch.DoesNotExistError
		if errors.As(err, &errNotExist) {
			// TODO Client Error
			return notExistError{dskey.Key(errNotExist)}
		}
		return fmt.Errorf("getting meeting id for collection %s id %d: %w", coll, id, err)
	}

	if !hasMeeting {
		hasOML, err := perm.HasOrganizationManagementLevel(ctx, ds, uid, perm.OMLCanManageOrganization)
		if err != nil {
			return fmt.Errorf("getting organization management level: %w", err)
		}

		if !hasOML {
			// TODO Client Error
			return permissionDeniedError{fmt.Errorf("you are not allowed to use history information on %s", fqid)}
		}
	} else {
		p, err := perm.New(ctx, ds, uid, meetingID)
		if err != nil {
			return fmt.Errorf("getting meeting permissions: %w", err)
		}

		if !p.Has(perm.MeetingCanSeeHistory) {
			// TODO Client Error
			return permissionDeniedError{fmt.Errorf("you are not allowed to use history information on %s", fqid)}
		}
	}

	if err := h.source.HistoryInformation(ctx, fqid, w); err != nil {
		return fmt.Errorf("getting history information: %w", err)
	}

	fmt.Fprintln(w)

	return nil
}

// Data returns old data from the datastore
func (h History) Data(ctx context.Context, userID int, kb KeysBuilder, position int) (map[dskey.Key][]byte, error) {
	getter := newGetPosition(h.source, position)
	restricter := newRestricter(h.source, getter, userID)

	keys, err := kb.Update(ctx, restricter)
	if err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	data, err := restricter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("get restricted data: %w", err)
	}

	for k, v := range data {
		if len(v) == 0 {
			delete(data, k)
		}
	}

	return data, nil
}

type permissionDeniedError struct {
	err error
}

func (e permissionDeniedError) Error() string {
	return fmt.Sprintf("permissoin denied: %v", e.err)
}

func (e permissionDeniedError) Type() string {
	return "permission_denied"
}

type invalidInputError struct {
	msg string
}

func (e invalidInputError) Error() string {
	return e.msg
}

func (e invalidInputError) Type() string {
	return "invalid_input"
}

type notExistError struct {
	key dskey.Key
}

func (e notExistError) Error() string {
	return fmt.Sprintf("%s does not exist", e.key)
}

func (e notExistError) Type() string {
	return "not_exist"
}
