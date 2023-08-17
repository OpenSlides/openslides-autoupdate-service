package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Speaker handels restrictions of the collection speaker.
//
// The user can see a speaker if he has list_of_speakers.can_see or if user_id is the request_user.
//
// Mode A: The user can see the speaker.
type Speaker struct{}

// Name returns the collection name.
func (s Speaker) Name() string {
	return "speaker"
}

// MeetingID returns the meetingID for the object.
func (s Speaker) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Speaker_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (s Speaker) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return s.see
	}
	return nil
}

func (s Speaker) see(ctx context.Context, ds *dsfetch.Fetch, speakerIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, s, speakerIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting perms for meetind %d: %w", meetingID, err)
		}

		if canSee := perms.Has(perm.ListOfSpeakersCanSee); canSee {
			return ids, nil
		}

		if canBeSpeaker := perms.Has(perm.ListOfSpeakersCanBeSpeaker); !canBeSpeaker {
			return nil, nil
		}

		requestUser, err := perm.RequestUserFromContext(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting request user: %w", err)
		}

		meetingUserIDs := make([]int, len(ids))
		for i, speakerID := range ids {
			ds.Speaker_MeetingUserID(speakerID).Lazy(&meetingUserIDs[i])
		}

		if err := ds.Execute(ctx); err != nil {
			return nil, fmt.Errorf("getting meeting-user ids of speakers: %w", err)
		}

		speakerUserIDs := make([]int, len(ids))
		for i, muID := range meetingUserIDs {
			ds.MeetingUser_UserID(muID).Lazy(&speakerUserIDs[i])
		}

		if err := ds.Execute(ctx); err != nil {
			return nil, fmt.Errorf("getting user ids of meeting-user ids: %w", err)
		}

		var allowed []int
		for i, uid := range speakerUserIDs {
			if uid == requestUser {
				allowed = append(allowed, ids[i])
			}
		}

		return allowed, nil
	})
}
