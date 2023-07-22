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

func (m Motion) see(ctx context.Context, fetcher *dsfetch.Fetch, motionIDs []int) ([]Tuple, error) {
	// TODO: Filter lead motion
	return byMeeting(ctx, fetcher, m, motionIDs, func(meetingID int, motionIDs []int) ([]Tuple, error) {
		// TODO: What about the super admin
		attrMotionCanSee := attribute.FuncPerm(meetingID, perm.MotionCanSee)

		return byRelationField(ctx, fetcher, m, fetcher.Motion_StateID, motionIDs, func(stateID int, motionIDs []int) ([]Tuple, error) {
			restrictions, err := fetcher.MotionState_Restrictions(stateID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting restrictions: %w", err)
			}

			if len(restrictions) == 0 {
				return TupleFromModeKeys(m, motionIDs, "C", attrMotionCanSee), nil
			}

			var hasIsSubmitterRestriction bool
			restrictPerms := make([]attribute.Func, 0, len(restrictions))
			for _, restriction := range restrictions {
				if restriction == "is_submitter" {
					hasIsSubmitterRestriction = true
					continue
				}
				restrictPerms = append(restrictPerms, attribute.FuncPerm(meetingID, perm.TPermission(restriction)))
			}

			if !hasIsSubmitterRestriction {
				attr := attribute.FuncAnd(
					attrMotionCanSee,
					attribute.FuncOr(restrictPerms...),
				)
				return TupleFromModeKeys(m, motionIDs, "C", attr), nil
			}

			motionAttr, err := submitterFunc(ctx, fetcher, motionIDs)
			if err != nil {
				return nil, fmt.Errorf("calculate submitter attributes: %w", err)
			}

			result := make([]Tuple, len(motionAttr))
			for i, isSubmitter := range motionAttr {
				motionID := motionIDs[i]

				result[i] = Tuple{
					Key: modeKey(m, motionID, "C"),
					Value: attribute.FuncAnd(
						attrMotionCanSee,
						attribute.FuncOr(
							attribute.FuncOr(restrictPerms...),
							isSubmitter,
						),
					),
				}

			}

			return result, nil
		})
	})
}

// leadMotionIndex creates an index from a motionID to its lead motion id. It
// also contains pairs for each found lead motion id.
//
// So each value in the index can also be found in the keys.
//
// motions without a lead motion are added with value 0
func leadMotionIndex(ctx context.Context, fetcher *dsfetch.Fetch, motionIDs []int) (map[int]int, error) {
	index := make(map[int]int, len(motionIDs))

	for len(motionIDs) > 0 {
		leadMotionIDs := make([]int, len(motionIDs))
		for i, motionID := range motionIDs {
			fetcher.Motion_LeadMotionID(motionID).Lazy(&leadMotionIDs[i])
		}

		if err := fetcher.Execute(ctx); err != nil {
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
func filterCanSeeLeadMotion(ctx context.Context, fetcher *dsfetch.Fetch, motionIDs []int, fn func([]int) ([]int, error)) ([]int, error) {
	index, err := leadMotionIndex(ctx, fetcher, motionIDs)
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

func (m Motion) modeA(ctx context.Context, fetcher *dsfetch.Fetch, motionIDs []int) ([]Tuple, error) {
	originIDs := make([][]int, len(motionIDs))
	derivedIDs := make([][]int, len(motionIDs))
	motionIDToRelatedIdx := make(map[int]int, len(motionIDs))
	for i, motionID := range motionIDs {
		fetcher.Motion_AllOriginIDs(motionID).Lazy(&originIDs[i])
		fetcher.Motion_AllDerivedMotionIDs(motionID).Lazy(&derivedIDs[i])
		motionIDToRelatedIdx[motionID] = i
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting origin and derived ids: %w", err)
	}

	allIDSet := set.New(motionIDs...)
	for i := range motionIDs {
		allIDSet.Add(originIDs[i]...)
		allIDSet.Add(derivedIDs[i]...)
	}
	allMotionIDs := allIDSet.List()

	tuples, err := m.see(ctx, fetcher, allMotionIDs)
	if err != nil {
		return nil, fmt.Errorf("see motion: %w", err)
	}

	motionIDToFunc := make(map[int]attribute.Func, len(allMotionIDs))
	for _, tuple := range tuples {
		motionID := tuple.Key.ID
		motionIDToFunc[motionID] = tuple.Value
	}

	result := make([]Tuple, len(motionIDs))
	for i, motionID := range motionIDs {
		relatedIdx := motionIDToRelatedIdx[motionID]
		originMotionIDs := originIDs[relatedIdx]
		derivedMotionIDs := derivedIDs[relatedIdx]

		funcList := make([]attribute.Func, 0, len(originMotionIDs)+len(derivedMotionIDs)+1)
		funcList = append(funcList, motionIDToFunc[motionID])
		for _, motionID := range originMotionIDs {
			funcList = append(funcList, motionIDToFunc[motionID])
		}
		for _, motionID := range derivedMotionIDs {
			funcList = append(funcList, motionIDToFunc[motionID])
		}

		result[i].Key = modeKey(m, motionID, "A")
		result[i].Value = attribute.FuncOr(funcList...)
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

	userIDsList := make([][]int, len(motionIDs))
	for i, submitterIDs := range submitterIDsList {
		userIDsList[i] = make([]int, len(submitterIDs))
		for j, submitterID := range submitterIDs {
			fetcher.MotionSubmitter_UserID(submitterID).Lazy(&userIDsList[i][j])
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
