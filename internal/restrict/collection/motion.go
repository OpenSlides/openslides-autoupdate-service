package collection

import (
	"context"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
	"github.com/OpenSlides/openslides-go/set"
)

// Motion handels restrictions of the collection motion.
//
// The user can see a motion if:
//
//	The user has motion.can_see in the meeting or
//	for one of the motions in `motion/all_derived_motion_ids` the user has motion.can_see_origin in the corresponding meeting
//
//	and for one `restriction` in the motion's state `state/restriction` field:
//	    If: `restriction` is `is_submitter`: The user needs to be a submitter of the motion
//	    Else: (a permission string): The user needs the permission
//
//	and - for amendments (lead_motion_id != null) - the user can also see the lead motion.
//
//	Admins in a meeting can see all motions.
//
// Mode A: The user can see the motion or the containing meeting is not closed and the user
// can see a referenced motion in motion/all_origin_ids and motion/all_derived_motion_ids.
//
// Mode B: The user has the permission motion.can_manage_metadata in the
// motion's meeting.
//
// Mode C: The user can see the motion.
//
// Mode D: Never published to any user.
//
// Mode E: If the motion states is_internal is true the user needs the
// permission motion.can_manage_metadata otherwise same as Mode C
type Motion struct{}

// Name returns the collection name.
func (m Motion) Name() string {
	return "motion"
}

// MeetingID returns the meetingID for the object.
func (m Motion) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Motion_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching meeting_id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m Motion) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.modeA
	case "B":
		return m.modeB
	case "C":
		return m.see
	case "D":
		return never
	case "E":
		return m.modeE
	}
	return nil
}

func (m Motion) see(ctx context.Context, ds *dsfetch.Fetch, motionIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return filterCanSeeLeadMotion(ctx, ds, motionIDs, func(motionIDs []int) ([]int, error) {
		return eachMeeting(ctx, ds, m, motionIDs, func(meetingID int, ids []int) ([]int, error) {
			perms, err := perm.FromContext(ctx, meetingID)
			if err != nil {
				return nil, fmt.Errorf("getting permissions: %w", err)
			}

			if perms.IsAdmin() {
				return ids, nil
			}

			lockedMeeting, err := ds.Meeting_LockedFromInside(meetingID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("fetching locked_from_inside: %w", err)
			}

			step1Allowed, err := eachCondition(ids, func(id int) (bool, error) {
				if perms.Has(perm.MotionCanSee) {
					return true, nil
				}

				if lockedMeeting {
					return false, nil
				}

				derivedIDs, err := ds.Motion_AllDerivedMotionIDs(id).Value(ctx)
				if err != nil {
					return false, fmt.Errorf("fetching all_derived_ids: %w", err)
				}

				meetingIDs := make([]int, len(derivedIDs))
				for i, derivedID := range derivedIDs {
					ds.Motion_MeetingID(derivedID).Lazy(&meetingIDs[i])
				}

				if err := ds.Execute(ctx); err != nil {
					return false, fmt.Errorf("fetching meeting ids from derived motions: %w", err)
				}

				for _, meetingID := range meetingIDs {
					perms, err := perm.FromContext(ctx, meetingID)
					if err != nil {
						return false, fmt.Errorf("getting permission from derived motion meeting: %w", err)
					}

					if perms.Has(perm.MotionCanSeeOrigin) {
						return true, nil
					}
				}

				return false, nil
			})

			if err != nil {
				return nil, fmt.Errorf("checking condition 1: %w", err)
			}

			return eachRelationField(ctx, ds.Motion_StateID, step1Allowed, func(stateID int, ids []int) ([]int, error) {
				restrictions, err := ds.MotionState_Restrictions(stateID).Value(ctx)
				if err != nil {
					return nil, fmt.Errorf("getting restrictions: %w", err)
				}

				if len(restrictions) == 0 {
					return ids, nil
				}

				hasIsSubmitterRestriction := false
				for _, restriction := range restrictions {
					if restriction == "is_submitter" {
						hasIsSubmitterRestriction = true
						continue
					}

					if perms.Has(perm.TPermission(restriction)) {
						return ids, nil
					}
				}
				if hasIsSubmitterRestriction {
					allowed, err := eachCondition(ids, func(motionID int) (bool, error) {
						submitter, err := isSubmitter(ctx, ds, requestUser, motionID)
						if err != nil {
							return false, fmt.Errorf("checking for motion submitter of motion %d: %w", motionID, err)
						}

						return submitter, nil
					})
					if err != nil {
						return nil, fmt.Errorf("checking if user is submitter: %w", err)
					}

					return allowed, nil
				}

				return nil, nil
			})
		})
	})
}

func (m Motion) modeB(ctx context.Context, ds *dsfetch.Fetch, motionIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, m, motionIDs, perm.MotionCanManageMetadata)
}

// leadMotionIndex creates an index from a motionID to its lead motion id. It
// also contains pairs for each found lead motion id.
//
// So each value in the index can also be found in the keys.
//
// motions without a lead motion are added with value 0
func leadMotionIndex(ctx context.Context, ds *dsfetch.Fetch, motionIDs []int) (map[int]int, error) {
	index := make(map[int]int, len(motionIDs))

	for len(motionIDs) > 0 {
		leadMotionIDs := make([]dsfetch.Maybe[int], len(motionIDs))
		for i, motionID := range motionIDs {
			ds.Motion_LeadMotionID(motionID).Lazy(&leadMotionIDs[i])
		}

		if err := ds.Execute(ctx); err != nil {
			return nil, fmt.Errorf("fetching lead motion ids: %w", err)
		}

		var nextMotionIDs []int
		for i := range leadMotionIDs {
			if _, ok := index[motionIDs[i]]; ok {
				continue
			}

			index[motionIDs[i]] = 0

			if v, ok := leadMotionIDs[i].Value(); ok {
				nextMotionIDs = append(nextMotionIDs, v)
				index[motionIDs[i]] = v
			}
		}
		motionIDs = nextMotionIDs
	}

	return index, nil
}

// isAllowedByLead returns true if the lead motion and its lead motion and so on is allowed
func isAllowedByLead(motionID int, allowedIDs set.Set[int], index map[int]int) bool {
	leadMotion := index[motionID]
	for {
		if leadMotion == 0 || leadMotion == motionID {
			return true
		}

		if !allowedIDs.Has(leadMotion) {
			return false
		}

		leadMotion = index[leadMotion]
	}
}

// filterCanSeeLeadMotion calls the given function by adding the lead motions to
// the motionIDs list.
//
// It only returns motions, where the user can also see the lead motion. This is
// done recursive, so for a lead_motion that also has a lead motion, the user
// must see all of them.
func filterCanSeeLeadMotion(ctx context.Context, ds *dsfetch.Fetch, motionIDs []int, fn func([]int) ([]int, error)) ([]int, error) {
	index, err := leadMotionIndex(ctx, ds, motionIDs)
	if err != nil {
		return nil, fmt.Errorf("create lead motion index: %w", err)
	}

	allIDs := make([]int, 0, len(index))
	for k := range index {
		allIDs = append(allIDs, k)
	}

	allowedIDs, err := fn(allIDs)
	if err != nil {
		return nil, fmt.Errorf("checking motions with lead motions: %w", err)
	}

	allowedSet := set.New(allowedIDs...)

	var filtered []int
	for _, motionID := range motionIDs {
		if !allowedSet.Has(motionID) {
			continue
		}

		if isAllowedByLead(motionID, allowedSet, index) {
			filtered = append(filtered, motionID)
		}
	}

	return filtered, nil
}

func isSubmitter(ctx context.Context, ds *dsfetch.Fetch, uid int, motionID int) (bool, error) {
	submitterIDs, err := ds.Motion_SubmitterIDs(motionID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting submitter ids: %w", err)
	}

	for _, submitterID := range submitterIDs {
		meetingUser, err := ds.MotionSubmitter_MeetingUserID(submitterID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting_user for submitter %d: %w", submitterID, err)
		}

		submitter, err := ds.MeetingUser_UserID(meetingUser).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting user id for submitter %d: %w", submitterID, err)
		}

		if submitter == uid {
			return true, nil
		}
	}

	return false, nil
}

func (m Motion) modeA(ctx context.Context, ds *dsfetch.Fetch, motionIDs ...int) ([]int, error) {
	allowed, err := m.see(ctx, ds, motionIDs...)
	if err != nil {
		return nil, fmt.Errorf("see motion: %w", err)
	}

	if len(allowed) == len(motionIDs) {
		return allowed, nil
	}

	notAllowed := set.New(motionIDs...)
	notAllowed.Remove(allowed...)

	allowed2, err := eachCondition(notAllowed.List(), func(motionID int) (bool, error) {
		allOriginIDs, err := ds.Motion_AllOriginIDs(motionID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting all origin ids: %w", err)
		}

		allDerivedMotionIDs, err := ds.Motion_AllDerivedMotionIDs(motionID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting all derived ids: %w", err)
		}

		motionIDs := make(map[int]struct{}, len(allOriginIDs)+len(allDerivedMotionIDs))
		for _, l := range [][]int{allOriginIDs, allDerivedMotionIDs} {
			for _, id := range l {
				motionIDs[id] = struct{}{}
			}
		}

		for referenceID := range motionIDs {
			// Check each motion as it own. It is enough when one motion returns
			// true. To call m.see with all motions at once would be slower.
			see, err := m.see(ctx, ds, referenceID)
			if err != nil {
				var errDoesNotExist dsfetch.DoesNotExistError
				if errors.As(err, &errDoesNotExist) {
					// The ids in all_derived_motion_ids and all_origin_ids can
					// contain motion, that were deleted. Ignore them.
					continue
				}
				return false, fmt.Errorf("see motion %d: %w", referenceID, err)
			}

			if len(see) == 1 {
				return true, nil
			}
		}
		return false, nil
	})
	if err != nil {
		return nil, fmt.Errorf("checkinging for referenced motions: %w", err)
	}

	return append(allowed, allowed2...), nil
}

func (m Motion) modeE(ctx context.Context, ds *dsfetch.Fetch, motionIDs ...int) ([]int, error) {
	allowed, err := m.see(ctx, ds, motionIDs...)
	if err != nil {
		return nil, fmt.Errorf("see motion: %w", err)
	}

	return eachMeeting(ctx, ds, m, allowed, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		if perms.Has(perm.MotionCanManageMetadata) {
			return ids, nil
		}

		return eachCondition(ids, func(motionID int) (bool, error) {
			motionStateID, err := ds.Motion_StateID(motionID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting motionStateID: %w", err)
			}

			isInternal, err := ds.MotionState_IsInternal(motionStateID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting motion state isInternal: %w", err)
			}

			return !isInternal, nil
		})
	})
}
