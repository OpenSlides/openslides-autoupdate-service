package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Motion handels restrictions of the collection motion.
//
// The user can see a motion if:
//
//		The user has motion.can_see in the meeting, and
//		For one `restriction` in the motion's state `state/restriction` field:
//		    If: `restriction` is `is_submitter`: The user needs to be a submitter of the motion
//		    Else: (a permission string): The user needs the permission
//	 And - for amendments (lead_motion_id != null) - the user can also see the lead motion.
//
// Mode A: The user can see the motion or can see a referenced motion in motion/all_origin_ids and motion/all_derived_motion_ids.
//
// Mode C: The user can see the motion.
//
// Mode D: Never published to any user.
type Motion struct{}

// Name returns the collection name.
func (m Motion) Name() string {
	return "motion"
}

// MeetingID returns the meetingID for the object.
func (m Motion) MeetingID(ctx context.Context, fetcher *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := fetcher.Motion_MeetingID(id).Value(ctx)
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
	case "C":
		return m.see
	case "D":
		return never
	}
	return nil
}

func (m Motion) see(ctx context.Context, fetcher *dsfetch.Fetch, motionIDs []int) ([]attribute.Func, error) {
	return filterCanSeeLeadMotion(ctx, fetcher, motionIDs, func(motionIDs []int) ([]attribute.Func, error) {
		return byMeeting(ctx, fetcher, m, motionIDs, func(meetingID int, motionIDs []int) ([]attribute.Func, error) {
			groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
			if err != nil {
				return nil, fmt.Errorf("getting group map: %w", err)
			}

			// TODO: What about the super admin
			attrMotionCanSee := attribute.FuncInGroup(groupMap[perm.MotionCanSee])

			return byRelationField(ctx, fetcher.Motion_StateID, motionIDs, func(stateID int, motionIDs []int) ([]attribute.Func, error) {
				restrictions, err := fetcher.MotionState_Restrictions(stateID).Value(ctx)
				if err != nil {
					return nil, fmt.Errorf("getting restrictions: %w", err)
				}

				if len(restrictions) == 0 {
					return attributeFuncList(len(motionIDs), attrMotionCanSee), nil
				}

				var hasIsSubmitterRestriction bool
				var restrictGroups []int
				for _, restriction := range restrictions {
					if restriction == "is_submitter" {
						hasIsSubmitterRestriction = true
						continue
					}
					restrictGroups = append(restrictGroups, groupMap[perm.TPermission(restriction)]...)
				}

				if !hasIsSubmitterRestriction {
					attr := attribute.FuncAnd(
						attrMotionCanSee,
						attribute.FuncOr(attribute.FuncInGroup(restrictGroups)),
					)
					return attributeFuncList(len(motionIDs), attr), nil
				}

				motionAttr, err := submitterFunc(ctx, fetcher, motionIDs)
				if err != nil {
					return nil, fmt.Errorf("calculate submitter attributes: %w", err)
				}

				result := make([]attribute.Func, len(motionIDs))
				for i := range motionIDs {
					isSubmitter := motionAttr[i]

					result[i] = attribute.FuncAnd(
						attrMotionCanSee,
						attribute.FuncOr(
							attribute.FuncInGroup(restrictGroups),
							isSubmitter,
						),
					)
				}

				return result, nil
			})
		})
	})
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
		leadMotionIDs := make([]int, len(motionIDs))
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

			index[motionIDs[i]] = leadMotionIDs[i]

			if leadMotionIDs[i] != 0 {
				nextMotionIDs = append(nextMotionIDs, leadMotionIDs[i])
			}
		}
		motionIDs = nextMotionIDs
	}

	return index, nil
}

// filterCanSeeLeadMotion calls the given function by adding the lead motions to
// the motionIDs list.
//
// The returned attributes require, that the user can see the motion and its
// leed motion. If this leed motion has also a leed motion, then this goes on.
func filterCanSeeLeadMotion(ctx context.Context, fetcher *dsfetch.Fetch, motionIDs []int, fn func([]int) ([]attribute.Func, error)) ([]attribute.Func, error) {
	index, err := leadMotionIndex(ctx, fetcher, motionIDs)
	if err != nil {
		return nil, fmt.Errorf("create lead motion index: %w", err)
	}

	// TODO: add a shortcut if no requested motion has a lead motion

	allMotionIDs := make([]int, 0, len(index))
	motionIDIdx := make(map[int]int, len(index)) // Index from motionID to index in allMotionIDs
	for motionID := range index {
		motionIDIdx[motionID] = len(allMotionIDs)
		allMotionIDs = append(allMotionIDs, motionID)
	}

	attrFuncs, err := fn(allMotionIDs)
	if err != nil {
		return nil, fmt.Errorf("checking motions with lead motions: %w", err)
	}

	result := make([]attribute.Func, len(motionIDs))
	for i, motionID := range motionIDs {
		var funcs []attribute.Func
		seen := set.New[int]()
		for motionID != 0 && !seen.Has(motionID) {
			funcs = append(funcs, attrFuncs[motionIDIdx[motionID]])
			seen.Add(motionID)
			motionID = index[motionID]
		}
		result[i] = attribute.FuncAnd(funcs...)
	}
	return result, nil
}

func (m Motion) modeA(ctx context.Context, fetcher *dsfetch.Fetch, motionIDs []int) ([]attribute.Func, error) {
	originIDs := make([][]int, len(motionIDs))
	derivedIDs := make([][]int, len(motionIDs))

	for i, motionID := range motionIDs {
		fetcher.Motion_AllOriginIDs(motionID).Lazy(&originIDs[i])
		fetcher.Motion_AllDerivedMotionIDs(motionID).Lazy(&derivedIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting origin and derived ids: %w", err)
	}

	relatedIdxFrom := make([]int, len(motionIDs))
	relatedIdxTo := make([]int, len(motionIDs))
	var allMotionIDs []int
	for i, motionID := range motionIDs {
		allMotionIDs = append(allMotionIDs, motionID)
		relatedIDs := append([]int{motionID}, append(originIDs[i], derivedIDs[i]...)...)
		relatedIdxFrom[i] = len(allMotionIDs)
		allMotionIDs = append(allMotionIDs, relatedIDs...)
		relatedIdxTo[i] = len(allMotionIDs)
	}

	attrFunc, err := Collection(ctx, m.Name()).Modes("C")(ctx, fetcher, allMotionIDs)
	if err != nil {
		return nil, fmt.Errorf("see motion: %w", err)
	}

	result := make([]attribute.Func, len(motionIDs))
	for i := range motionIDs {
		size := relatedIdxTo[i] - relatedIdxFrom[i]

		funcList := make([]attribute.Func, size)
		for j := 0; j < size; j++ {
			funcList[j] = attrFunc[relatedIdxFrom[i]+j]
		}

		result[i] = attribute.FuncOr(funcList...)
	}

	return result, nil
}

// submitterFunc returns for a list of motions for each motion an
// attribute.Func, that returns true, if the request user is a submitter of that
// motion.
func submitterFunc(ctx context.Context, fetcher *dsfetch.Fetch, motionIDs []int) ([]attribute.Func, error) {
	submitterIDsList := make([][]int, len(motionIDs))
	for i := range motionIDs {
		fetcher.Motion_SubmitterIDs(motionIDs[i]).Lazy(&submitterIDsList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting submitter ids: %w", err)
	}

	meetingUserIDsList := make([][]int, len(motionIDs))
	for i, submitterIDs := range submitterIDsList {
		meetingUserIDsList[i] = make([]int, len(submitterIDs))
		for j, submitterID := range submitterIDs {
			fetcher.MotionSubmitter_MeetingUserID(submitterID).Lazy(&meetingUserIDsList[i][j])
		}
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting meeting user ids: %w", err)
	}

	userIDsList := make([][]int, len(motionIDs))
	for i, meetingUserIDs := range meetingUserIDsList {
		userIDsList[i] = make([]int, len(meetingUserIDs))
		for j, meetingUserID := range meetingUserIDs {
			fetcher.MeetingUser_UserID(meetingUserID).Lazy(&userIDsList[i][j])
		}
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting user ids: %w", err)
	}

	out := make([]attribute.Func, len(motionIDs))
	for i, userIDs := range userIDsList {
		out[i] = attribute.FuncUserIDs(userIDs)
	}
	return out, nil
}
