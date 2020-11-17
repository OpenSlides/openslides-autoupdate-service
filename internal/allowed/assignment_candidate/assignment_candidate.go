package assignment_candidate

import (
	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

var selfCreate = allowed.BuildCreateThroughId([]string{
	"assignment_id",
	"user_id",
}, "assignment", "assignment_id", "assignments.can_nominate_self")
var otherCreate = allowed.BuildCreateThroughId([]string{
	"assignment_id",
	"user_id",
}, "assignment", "assignment_id", "assignments.can_nominate_other")

func Create(params *allowed.IsAllowedParams) (map[string]interface{}, error) {
	userID, err := allowed.GetId(params.Data, "user_id")
	if err != nil {
		return nil, err
	}

	if userID == params.UserID {
		return selfCreate(params)
	} else {
		return otherCreate(params)
	}
}

var Sort = allowed.BuildModifyThroughId([]string{
	"assignment_id",
	"candidate_ids",
}, "assignment_candidate", "assignment", "assignment_id", "assignments.can_manage")

func Delete(params *allowed.IsAllowedParams) (map[string]interface{}, error) {
	if err := allowed.ValidateFields(params.Data, allowed.MakeSet([]string{"id"})); err != nil {
		return nil, err
	}

	isAllowed, err := allowed.CheckUser(params)
	if err != nil {
		return nil, err
	}
	if isAllowed {
		return nil, nil
	}

	id, err := allowed.GetId(params.Data, "id")
	if err != nil {
		return nil, err
	}
	fqid := definitions.FqidFromCollectionAndId("assignment_candidate", id)
	exists, err := allowed.DoesModelExists(fqid, params.DataProvider)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, allowed.NotAllowedf("The assignment_candidate with id %d does not exist", id)
	}
	userFqfield := definitions.FqfieldFromFqidAndField(fqid, "user_id")
	assignmentCandidateUserID, err := params.DataProvider.GetInt(userFqfield)
	if err != nil {
		return nil, err
	}

	meetingID, err := allowed.GetMeetingIDFromModel(fqid, params.DataProvider)
	if err != nil {
		return nil, err
	}

	var permission string
	if assignmentCandidateUserID == params.UserID {
		permission = "assignments.can_nominate_self"
	} else {
		permission = "assignments.can_nominate_other"
	}

	err = allowed.CheckCommitteeMeetingPermissions(params, meetingID, permission)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
