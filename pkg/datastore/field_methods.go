// Code generated with go generate DO NOT EDIT.
package datastore

import (
	"context"
	"encoding/json"
	"fmt"
)

func (f Fields) AgendaItem_ChildIDs(ctx context.Context, AgendaItemID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/child_ids", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"agenda_item/%d/child_ids\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_Closed(ctx context.Context, AgendaItemID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/closed", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"agenda_item/%d/closed\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_Comment(ctx context.Context, AgendaItemID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/comment", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"agenda_item/%d/comment\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_ContentObjectID(ctx context.Context, AgendaItemID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/content_object_id", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"agenda_item/%d/content_object_id\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_Duration(ctx context.Context, AgendaItemID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/duration", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"agenda_item/%d/duration\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_ID(ctx context.Context, AgendaItemID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/id", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"agenda_item/%d/id\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_IsHidden(ctx context.Context, AgendaItemID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/is_hidden", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"agenda_item/%d/is_hidden\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_IsInternal(ctx context.Context, AgendaItemID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/is_internal", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"agenda_item/%d/is_internal\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_ItemNumber(ctx context.Context, AgendaItemID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/item_number", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"agenda_item/%d/item_number\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_Level(ctx context.Context, AgendaItemID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/level", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"agenda_item/%d/level\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_MeetingID(ctx context.Context, AgendaItemID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/meeting_id", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"agenda_item/%d/meeting_id\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_ParentID(ctx context.Context, AgendaItemID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/parent_id", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"agenda_item/%d/parent_id\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_ProjectionIDs(ctx context.Context, AgendaItemID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/projection_ids", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"agenda_item/%d/projection_ids\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_TagIDs(ctx context.Context, AgendaItemID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/tag_ids", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"agenda_item/%d/tag_ids\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_Type(ctx context.Context, AgendaItemID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/type", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"agenda_item/%d/type\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AgendaItem_Weight(ctx context.Context, AgendaItemID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/weight", AgendaItemID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"agenda_item/%d/weight\": %w", AgendaItemID, err)
	}

	return v, nil
}

func (f Fields) AssignmentCandidate_AssignmentID(ctx context.Context, AssignmentCandidateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/assignment_id", AssignmentCandidateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment_candidate/%d/assignment_id\": %w", AssignmentCandidateID, err)
	}

	return v, nil
}

func (f Fields) AssignmentCandidate_ID(ctx context.Context, AssignmentCandidateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/id", AssignmentCandidateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment_candidate/%d/id\": %w", AssignmentCandidateID, err)
	}

	return v, nil
}

func (f Fields) AssignmentCandidate_MeetingID(ctx context.Context, AssignmentCandidateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/meeting_id", AssignmentCandidateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment_candidate/%d/meeting_id\": %w", AssignmentCandidateID, err)
	}

	return v, nil
}

func (f Fields) AssignmentCandidate_UserID(ctx context.Context, AssignmentCandidateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/user_id", AssignmentCandidateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment_candidate/%d/user_id\": %w", AssignmentCandidateID, err)
	}

	return v, nil
}

func (f Fields) AssignmentCandidate_Weight(ctx context.Context, AssignmentCandidateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/weight", AssignmentCandidateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment_candidate/%d/weight\": %w", AssignmentCandidateID, err)
	}

	return v, nil
}

func (f Fields) Assignment_AgendaItemID(ctx context.Context, AssignmentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/agenda_item_id", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment/%d/agenda_item_id\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_AttachmentIDs(ctx context.Context, AssignmentID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/attachment_ids", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"assignment/%d/attachment_ids\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_CandidateIDs(ctx context.Context, AssignmentID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/candidate_ids", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"assignment/%d/candidate_ids\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_DefaultPollDescription(ctx context.Context, AssignmentID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/default_poll_description", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"assignment/%d/default_poll_description\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_Description(ctx context.Context, AssignmentID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/description", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"assignment/%d/description\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_ID(ctx context.Context, AssignmentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/id", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment/%d/id\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_ListOfSpeakersID(ctx context.Context, AssignmentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/list_of_speakers_id", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment/%d/list_of_speakers_id\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_MeetingID(ctx context.Context, AssignmentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/meeting_id", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment/%d/meeting_id\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_NumberPollCandidates(ctx context.Context, AssignmentID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/number_poll_candidates", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"assignment/%d/number_poll_candidates\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_OpenPosts(ctx context.Context, AssignmentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/open_posts", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"assignment/%d/open_posts\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_Phase(ctx context.Context, AssignmentID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/phase", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"assignment/%d/phase\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_PollIDs(ctx context.Context, AssignmentID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/poll_ids", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"assignment/%d/poll_ids\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_ProjectionIDs(ctx context.Context, AssignmentID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/projection_ids", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"assignment/%d/projection_ids\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_TagIDs(ctx context.Context, AssignmentID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/tag_ids", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"assignment/%d/tag_ids\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) Assignment_Title(ctx context.Context, AssignmentID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/title", AssignmentID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"assignment/%d/title\": %w", AssignmentID, err)
	}

	return v, nil
}

func (f Fields) ChatGroup_ID(ctx context.Context, ChatGroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/id", ChatGroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"chat_group/%d/id\": %w", ChatGroupID, err)
	}

	return v, nil
}

func (f Fields) ChatGroup_MeetingID(ctx context.Context, ChatGroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/meeting_id", ChatGroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"chat_group/%d/meeting_id\": %w", ChatGroupID, err)
	}

	return v, nil
}

func (f Fields) ChatGroup_Name(ctx context.Context, ChatGroupID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/name", ChatGroupID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"chat_group/%d/name\": %w", ChatGroupID, err)
	}

	return v, nil
}

func (f Fields) ChatGroup_ReadGroupIDs(ctx context.Context, ChatGroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/read_group_ids", ChatGroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"chat_group/%d/read_group_ids\": %w", ChatGroupID, err)
	}

	return v, nil
}

func (f Fields) ChatGroup_Weight(ctx context.Context, ChatGroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/weight", ChatGroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"chat_group/%d/weight\": %w", ChatGroupID, err)
	}

	return v, nil
}

func (f Fields) ChatGroup_WriteGroupIDs(ctx context.Context, ChatGroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/write_group_ids", ChatGroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"chat_group/%d/write_group_ids\": %w", ChatGroupID, err)
	}

	return v, nil
}

func (f Fields) Committee_DefaultMeetingID(ctx context.Context, CommitteeID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/default_meeting_id", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"committee/%d/default_meeting_id\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_Description(ctx context.Context, CommitteeID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/description", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"committee/%d/description\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_ForwardToCommitteeIDs(ctx context.Context, CommitteeID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/forward_to_committee_ids", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"committee/%d/forward_to_committee_ids\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_ID(ctx context.Context, CommitteeID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/id", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"committee/%d/id\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_MeetingIDs(ctx context.Context, CommitteeID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/meeting_ids", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"committee/%d/meeting_ids\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_Name(ctx context.Context, CommitteeID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/name", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"committee/%d/name\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_OrganizationID(ctx context.Context, CommitteeID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/organization_id", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"committee/%d/organization_id\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_OrganizationTagIDs(ctx context.Context, CommitteeID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/organization_tag_ids", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"committee/%d/organization_tag_ids\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_ReceiveForwardingsFromCommitteeIDs(ctx context.Context, CommitteeID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/receive_forwardings_from_committee_ids", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"committee/%d/receive_forwardings_from_committee_ids\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_TemplateMeetingID(ctx context.Context, CommitteeID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/template_meeting_id", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"committee/%d/template_meeting_id\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Committee_UserIDs(ctx context.Context, CommitteeID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/user_ids", CommitteeID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"committee/%d/user_ids\": %w", CommitteeID, err)
	}

	return v, nil
}

func (f Fields) Group_AdminGroupForMeetingID(ctx context.Context, GroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/admin_group_for_meeting_id", GroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"group/%d/admin_group_for_meeting_id\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_DefaultGroupForMeetingID(ctx context.Context, GroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/default_group_for_meeting_id", GroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"group/%d/default_group_for_meeting_id\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_ID(ctx context.Context, GroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/id", GroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"group/%d/id\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_MediafileAccessGroupIDs(ctx context.Context, GroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/mediafile_access_group_ids", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/mediafile_access_group_ids\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_MediafileInheritedAccessGroupIDs(ctx context.Context, GroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/mediafile_inherited_access_group_ids", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/mediafile_inherited_access_group_ids\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_MeetingID(ctx context.Context, GroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/meeting_id", GroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"group/%d/meeting_id\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_Name(ctx context.Context, GroupID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "group/%d/name", GroupID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"group/%d/name\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_Permissions(ctx context.Context, GroupID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "group/%d/permissions", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/permissions\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_PollIDs(ctx context.Context, GroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/poll_ids", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/poll_ids\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_ReadChatGroupIDs(ctx context.Context, GroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/read_chat_group_ids", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/read_chat_group_ids\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_ReadCommentSectionIDs(ctx context.Context, GroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/read_comment_section_ids", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/read_comment_section_ids\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_UsedAsAssignmentPollDefaultID(ctx context.Context, GroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/used_as_assignment_poll_default_id", GroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"group/%d/used_as_assignment_poll_default_id\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_UsedAsMotionPollDefaultID(ctx context.Context, GroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/used_as_motion_poll_default_id", GroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"group/%d/used_as_motion_poll_default_id\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_UsedAsPollDefaultID(ctx context.Context, GroupID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/used_as_poll_default_id", GroupID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"group/%d/used_as_poll_default_id\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_UserIDs(ctx context.Context, GroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/user_ids", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/user_ids\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_WriteChatGroupIDs(ctx context.Context, GroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/write_chat_group_ids", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/write_chat_group_ids\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) Group_WriteCommentSectionIDs(ctx context.Context, GroupID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/write_comment_section_ids", GroupID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"group/%d/write_comment_section_ids\": %w", GroupID, err)
	}

	return v, nil
}

func (f Fields) ListOfSpeakers_Closed(ctx context.Context, ListOfSpeakersID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/closed", ListOfSpeakersID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"list_of_speakers/%d/closed\": %w", ListOfSpeakersID, err)
	}

	return v, nil
}

func (f Fields) ListOfSpeakers_ContentObjectID(ctx context.Context, ListOfSpeakersID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/content_object_id", ListOfSpeakersID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"list_of_speakers/%d/content_object_id\": %w", ListOfSpeakersID, err)
	}

	return v, nil
}

func (f Fields) ListOfSpeakers_ID(ctx context.Context, ListOfSpeakersID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/id", ListOfSpeakersID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"list_of_speakers/%d/id\": %w", ListOfSpeakersID, err)
	}

	return v, nil
}

func (f Fields) ListOfSpeakers_MeetingID(ctx context.Context, ListOfSpeakersID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/meeting_id", ListOfSpeakersID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"list_of_speakers/%d/meeting_id\": %w", ListOfSpeakersID, err)
	}

	return v, nil
}

func (f Fields) ListOfSpeakers_ProjectionIDs(ctx context.Context, ListOfSpeakersID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/projection_ids", ListOfSpeakersID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"list_of_speakers/%d/projection_ids\": %w", ListOfSpeakersID, err)
	}

	return v, nil
}

func (f Fields) ListOfSpeakers_SpeakerIDs(ctx context.Context, ListOfSpeakersID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/speaker_ids", ListOfSpeakersID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"list_of_speakers/%d/speaker_ids\": %w", ListOfSpeakersID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_AccessGroupIDs(ctx context.Context, MediafileID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/access_group_ids", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"mediafile/%d/access_group_ids\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_AttachmentIDs(ctx context.Context, MediafileID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/attachment_ids", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"mediafile/%d/attachment_ids\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_ChildIDs(ctx context.Context, MediafileID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/child_ids", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"mediafile/%d/child_ids\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_CreateTimestamp(ctx context.Context, MediafileID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/create_timestamp", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"mediafile/%d/create_timestamp\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_Filename(ctx context.Context, MediafileID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/filename", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"mediafile/%d/filename\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_Filesize(ctx context.Context, MediafileID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/filesize", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"mediafile/%d/filesize\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_ID(ctx context.Context, MediafileID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/id", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"mediafile/%d/id\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_InheritedAccessGroupIDs(ctx context.Context, MediafileID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/inherited_access_group_ids", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"mediafile/%d/inherited_access_group_ids\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_IsDirectory(ctx context.Context, MediafileID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/is_directory", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"mediafile/%d/is_directory\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_IsPublic(ctx context.Context, MediafileID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/is_public", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"mediafile/%d/is_public\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_ListOfSpeakersID(ctx context.Context, MediafileID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/list_of_speakers_id", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"mediafile/%d/list_of_speakers_id\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_MeetingID(ctx context.Context, MediafileID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/meeting_id", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"mediafile/%d/meeting_id\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_Mimetype(ctx context.Context, MediafileID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/mimetype", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"mediafile/%d/mimetype\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_ParentID(ctx context.Context, MediafileID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/parent_id", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"mediafile/%d/parent_id\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_PdfInformation(ctx context.Context, MediafileID int) (json.RawMessage, error) {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/pdf_information", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"mediafile/%d/pdf_information\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_ProjectionIDs(ctx context.Context, MediafileID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/projection_ids", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"mediafile/%d/projection_ids\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_Title(ctx context.Context, MediafileID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/title", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"mediafile/%d/title\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_UsedAsFontInMeetingIDTmpl(ctx context.Context, MediafileID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/used_as_font_$_in_meeting_id", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"mediafile/%d/used_as_font_$_in_meeting_id\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_UsedAsFontInMeetingID(ctx context.Context, MediafileID int, replacement string) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/used_as_font_$%s_in_meeting_id", MediafileID, replacement)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"mediafile/%d/used_as_font_$%s_in_meeting_id\": %w", MediafileID, replacement, err)
	}

	return v, nil
}

func (f Fields) Mediafile_UsedAsLogoInMeetingIDTmpl(ctx context.Context, MediafileID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/used_as_logo_$_in_meeting_id", MediafileID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"mediafile/%d/used_as_logo_$_in_meeting_id\": %w", MediafileID, err)
	}

	return v, nil
}

func (f Fields) Mediafile_UsedAsLogoInMeetingID(ctx context.Context, MediafileID int, replacement string) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/used_as_logo_$%s_in_meeting_id", MediafileID, replacement)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"mediafile/%d/used_as_logo_$%s_in_meeting_id\": %w", MediafileID, replacement, err)
	}

	return v, nil
}

func (f Fields) Meeting_AdminGroupID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/admin_group_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/admin_group_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AgendaEnableNumbering(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_enable_numbering", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/agenda_enable_numbering\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AgendaItemCreation(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_item_creation", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/agenda_item_creation\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AgendaItemIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_item_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/agenda_item_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AgendaNewItemsDefaultVisibility(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_new_items_default_visibility", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/agenda_new_items_default_visibility\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AgendaNumberPrefix(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_number_prefix", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/agenda_number_prefix\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AgendaNumeralSystem(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_numeral_system", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/agenda_numeral_system\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AgendaShowInternalItemsOnProjector(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_show_internal_items_on_projector", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/agenda_show_internal_items_on_projector\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AgendaShowSubtitles(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_show_subtitles", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/agenda_show_subtitles\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AllProjectionIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/all_projection_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/all_projection_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ApplauseEnable(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_enable", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/applause_enable\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ApplauseMaxAmount(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_max_amount", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/applause_max_amount\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ApplauseMinAmount(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_min_amount", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/applause_min_amount\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ApplauseParticleImageUrl(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_particle_image_url", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/applause_particle_image_url\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ApplauseShowLevel(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_show_level", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/applause_show_level\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ApplauseTimeout(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_timeout", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/applause_timeout\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ApplauseType(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_type", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/applause_type\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentCandidateIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_candidate_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/assignment_candidate_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/assignment_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentPollAddCandidatesToListOfSpeakers(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_add_candidates_to_list_of_speakers", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/assignment_poll_add_candidates_to_list_of_speakers\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentPollBallotPaperNumber(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_ballot_paper_number", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/assignment_poll_ballot_paper_number\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentPollBallotPaperSelection(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_ballot_paper_selection", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/assignment_poll_ballot_paper_selection\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentPollDefault100PercentBase(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_default_100_percent_base", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/assignment_poll_default_100_percent_base\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentPollDefaultGroupIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_default_group_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/assignment_poll_default_group_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentPollDefaultMethod(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_default_method", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/assignment_poll_default_method\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentPollDefaultType(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_default_type", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/assignment_poll_default_type\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentPollSortPollResultByVotes(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_sort_poll_result_by_votes", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/assignment_poll_sort_poll_result_by_votes\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentsExportPreamble(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignments_export_preamble", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/assignments_export_preamble\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_AssignmentsExportTitle(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignments_export_title", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/assignments_export_title\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ChatGroupIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/chat_group_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/chat_group_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_CommitteeID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/committee_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/committee_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceAutoConnect(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_auto_connect", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/conference_auto_connect\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceAutoConnectNextSpeakers(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_auto_connect_next_speakers", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/conference_auto_connect_next_speakers\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceEnableHelpdesk(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_enable_helpdesk", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/conference_enable_helpdesk\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceLosRestriction(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_los_restriction", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/conference_los_restriction\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceOpenMicrophone(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_open_microphone", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/conference_open_microphone\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceOpenVideo(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_open_video", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/conference_open_video\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceShow(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_show", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/conference_show\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceStreamPosterUrl(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_stream_poster_url", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/conference_stream_poster_url\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ConferenceStreamUrl(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_stream_url", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/conference_stream_url\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_CustomTranslations(ctx context.Context, MeetingID int) (json.RawMessage, error) {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/custom_translations", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/custom_translations\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_DefaultGroupID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/default_group_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/default_group_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_DefaultMeetingForCommitteeID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/default_meeting_for_committee_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/default_meeting_for_committee_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_DefaultProjectorIDTmpl(ctx context.Context, MeetingID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/default_projector_$_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/default_projector_$_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_DefaultProjectorID(ctx context.Context, MeetingID int, replacement string) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/default_projector_$%s_id", MeetingID, replacement)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/default_projector_$%s_id\": %w", MeetingID, replacement, err)
	}

	return v, nil
}

func (f Fields) Meeting_Description(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/description", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/description\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_EnableAnonymous(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/enable_anonymous", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/enable_anonymous\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_EnableChat(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/enable_chat", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/enable_chat\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_EndTime(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/end_time", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/end_time\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ExportCsvEncoding(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_csv_encoding", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/export_csv_encoding\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ExportCsvSeparator(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_csv_separator", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/export_csv_separator\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ExportPdfFontsize(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_pdf_fontsize", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/export_pdf_fontsize\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ExportPdfPagenumberAlignment(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_pdf_pagenumber_alignment", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/export_pdf_pagenumber_alignment\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ExportPdfPagesize(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_pdf_pagesize", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/export_pdf_pagesize\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_FontIDTmpl(ctx context.Context, MeetingID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/font_$_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/font_$_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_FontID(ctx context.Context, MeetingID int, replacement string) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/font_$%s_id", MeetingID, replacement)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/font_$%s_id\": %w", MeetingID, replacement, err)
	}

	return v, nil
}

func (f Fields) Meeting_GroupIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/group_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/group_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ImportedAt(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/imported_at", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/imported_at\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_JitsiDomain(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/jitsi_domain", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/jitsi_domain\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_JitsiRoomName(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/jitsi_room_name", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/jitsi_room_name\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_JitsiRoomPassword(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/jitsi_room_password", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/jitsi_room_password\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersAmountLastOnProjector(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_amount_last_on_projector", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_amount_last_on_projector\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersAmountNextOnProjector(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_amount_next_on_projector", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_amount_next_on_projector\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersCanSetContributionSelf(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_can_set_contribution_self", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_can_set_contribution_self\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersCountdownID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_countdown_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_countdown_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersCoupleCountdown(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_couple_countdown", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_couple_countdown\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_enable_point_of_order_speakers", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_enable_point_of_order_speakers\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersEnableProContraSpeech(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_enable_pro_contra_speech", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_enable_pro_contra_speech\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersInitiallyClosed(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_initially_closed", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_initially_closed\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersPresentUsersOnly(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_present_users_only", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_present_users_only\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_show_amount_of_speakers_on_slide", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_show_amount_of_speakers_on_slide\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersShowFirstContribution(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_show_first_contribution", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_show_first_contribution\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ListOfSpeakersSpeakerNoteForEveryone(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_speaker_note_for_everyone", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/list_of_speakers_speaker_note_for_everyone\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_Location(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/location", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/location\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_LogoIDTmpl(ctx context.Context, MeetingID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/logo_$_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/logo_$_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_LogoID(ctx context.Context, MeetingID int, replacement string) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/logo_$%s_id", MeetingID, replacement)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/logo_$%s_id\": %w", MeetingID, replacement, err)
	}

	return v, nil
}

func (f Fields) Meeting_MediafileIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/mediafile_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/mediafile_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionBlockIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_block_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_block_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionCategoryIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_category_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_category_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionChangeRecommendationIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_change_recommendation_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_change_recommendation_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionCommentIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_comment_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_comment_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionCommentSectionIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_comment_section_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_comment_section_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionPollBallotPaperNumber(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_ballot_paper_number", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/motion_poll_ballot_paper_number\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionPollBallotPaperSelection(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_ballot_paper_selection", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motion_poll_ballot_paper_selection\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionPollDefault100PercentBase(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_default_100_percent_base", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motion_poll_default_100_percent_base\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionPollDefaultGroupIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_default_group_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_poll_default_group_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionPollDefaultType(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_default_type", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motion_poll_default_type\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionStateIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_state_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_state_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionStatuteParagraphIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_statute_paragraph_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_statute_paragraph_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionSubmitterIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_submitter_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_submitter_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionWorkflowIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_workflow_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/motion_workflow_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsAmendmentsEnabled(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_enabled", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_amendments_enabled\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsAmendmentsInMainList(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_in_main_list", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_amendments_in_main_list\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsAmendmentsMultipleParagraphs(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_multiple_paragraphs", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_amendments_multiple_paragraphs\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsAmendmentsOfAmendments(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_of_amendments", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_amendments_of_amendments\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsAmendmentsPrefix(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_prefix", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_amendments_prefix\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsAmendmentsTextMode(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_text_mode", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_amendments_text_mode\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsDefaultAmendmentWorkflowID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_amendment_workflow_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/motions_default_amendment_workflow_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsDefaultLineNumbering(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_line_numbering", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_default_line_numbering\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsDefaultSorting(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_sorting", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_default_sorting\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsDefaultStatuteAmendmentWorkflowID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_statute_amendment_workflow_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/motions_default_statute_amendment_workflow_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsDefaultWorkflowID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_workflow_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/motions_default_workflow_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsEnableReasonOnProjector(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_enable_reason_on_projector", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_enable_reason_on_projector\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsEnableRecommendationOnProjector(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_enable_recommendation_on_projector", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_enable_recommendation_on_projector\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsEnableSideboxOnProjector(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_enable_sidebox_on_projector", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_enable_sidebox_on_projector\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsEnableTextOnProjector(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_enable_text_on_projector", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_enable_text_on_projector\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsExportFollowRecommendation(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_export_follow_recommendation", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_export_follow_recommendation\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsExportPreamble(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_export_preamble", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_export_preamble\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsExportSubmitterRecommendation(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_export_submitter_recommendation", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_export_submitter_recommendation\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsExportTitle(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_export_title", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_export_title\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsLineLength(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_line_length", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/motions_line_length\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsNumberMinDigits(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_number_min_digits", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/motions_number_min_digits\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsNumberType(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_number_type", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_number_type\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsNumberWithBlank(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_number_with_blank", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_number_with_blank\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsPreamble(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_preamble", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_preamble\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsReasonRequired(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_reason_required", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_reason_required\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsRecommendationTextMode(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_recommendation_text_mode", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_recommendation_text_mode\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsRecommendationsBy(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_recommendations_by", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_recommendations_by\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsShowReferringMotions(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_show_referring_motions", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_show_referring_motions\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsShowSequentialNumber(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_show_sequential_number", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_show_sequential_number\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsStatuteRecommendationsBy(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_statute_recommendations_by", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/motions_statute_recommendations_by\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsStatutesEnabled(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_statutes_enabled", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/motions_statutes_enabled\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_MotionsSupportersMinAmount(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_supporters_min_amount", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/motions_supporters_min_amount\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_Name(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/name", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/name\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_OptionIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/option_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/option_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_OrganizationTagIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/organization_tag_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/organization_tag_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PersonalNoteIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/personal_note_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/personal_note_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollBallotPaperNumber(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_ballot_paper_number", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/poll_ballot_paper_number\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollBallotPaperSelection(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_ballot_paper_selection", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/poll_ballot_paper_selection\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollCountdownID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_countdown_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/poll_countdown_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollCoupleCountdown(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_couple_countdown", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/poll_couple_countdown\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollDefault100PercentBase(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_default_100_percent_base", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/poll_default_100_percent_base\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollDefaultGroupIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_default_group_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/poll_default_group_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollDefaultMethod(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_default_method", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/poll_default_method\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollDefaultType(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_default_type", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/poll_default_type\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/poll_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PollSortPollResultByVotes(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_sort_poll_result_by_votes", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/poll_sort_poll_result_by_votes\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_PresentUserIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/present_user_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/present_user_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ProjectionIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projection_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/projection_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ProjectorCountdownDefaultTime(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_countdown_default_time", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/projector_countdown_default_time\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ProjectorCountdownIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_countdown_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/projector_countdown_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ProjectorCountdownWarningTime(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_countdown_warning_time", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/projector_countdown_warning_time\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ProjectorIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/projector_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ProjectorMessageIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_message_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/projector_message_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_ReferenceProjectorID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/reference_projector_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/reference_projector_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_SpeakerIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/speaker_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/speaker_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_StartTime(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/start_time", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/start_time\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_TagIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/tag_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/tag_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_TemplateForCommitteeID(ctx context.Context, MeetingID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/template_for_committee_id", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"meeting/%d/template_for_committee_id\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_TopicIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/topic_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/topic_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UrlName(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/url_name", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/url_name\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UserIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/user_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/user_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersAllowSelfSetPresent(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_allow_self_set_present", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/users_allow_self_set_present\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersEmailBody(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_email_body", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_email_body\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersEmailReplyto(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_email_replyto", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_email_replyto\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersEmailSender(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_email_sender", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_email_sender\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersEmailSubject(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_email_subject", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_email_subject\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersEnablePresenceView(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_enable_presence_view", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/users_enable_presence_view\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersEnableVoteWeight(ctx context.Context, MeetingID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_enable_vote_weight", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"meeting/%d/users_enable_vote_weight\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersPdfUrl(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_url", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_pdf_url\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersPdfWelcometext(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_welcometext", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_pdf_welcometext\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersPdfWelcometitle(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_welcometitle", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_pdf_welcometitle\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersPdfWlanEncryption(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_wlan_encryption", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_pdf_wlan_encryption\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersPdfWlanPassword(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_wlan_password", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_pdf_wlan_password\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersPdfWlanSsid(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_wlan_ssid", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_pdf_wlan_ssid\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_UsersSortBy(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_sort_by", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/users_sort_by\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_VoteIDs(ctx context.Context, MeetingID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/vote_ids", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"meeting/%d/vote_ids\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_WelcomeText(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/welcome_text", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/welcome_text\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) Meeting_WelcomeTitle(ctx context.Context, MeetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/welcome_title", MeetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"meeting/%d/welcome_title\": %w", MeetingID, err)
	}

	return v, nil
}

func (f Fields) MotionBlock_AgendaItemID(ctx context.Context, MotionBlockID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/agenda_item_id", MotionBlockID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_block/%d/agenda_item_id\": %w", MotionBlockID, err)
	}

	return v, nil
}

func (f Fields) MotionBlock_ID(ctx context.Context, MotionBlockID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/id", MotionBlockID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_block/%d/id\": %w", MotionBlockID, err)
	}

	return v, nil
}

func (f Fields) MotionBlock_Internal(ctx context.Context, MotionBlockID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/internal", MotionBlockID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_block/%d/internal\": %w", MotionBlockID, err)
	}

	return v, nil
}

func (f Fields) MotionBlock_ListOfSpeakersID(ctx context.Context, MotionBlockID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/list_of_speakers_id", MotionBlockID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_block/%d/list_of_speakers_id\": %w", MotionBlockID, err)
	}

	return v, nil
}

func (f Fields) MotionBlock_MeetingID(ctx context.Context, MotionBlockID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/meeting_id", MotionBlockID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_block/%d/meeting_id\": %w", MotionBlockID, err)
	}

	return v, nil
}

func (f Fields) MotionBlock_MotionIDs(ctx context.Context, MotionBlockID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/motion_ids", MotionBlockID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_block/%d/motion_ids\": %w", MotionBlockID, err)
	}

	return v, nil
}

func (f Fields) MotionBlock_ProjectionIDs(ctx context.Context, MotionBlockID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/projection_ids", MotionBlockID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_block/%d/projection_ids\": %w", MotionBlockID, err)
	}

	return v, nil
}

func (f Fields) MotionBlock_Title(ctx context.Context, MotionBlockID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/title", MotionBlockID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_block/%d/title\": %w", MotionBlockID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_ChildIDs(ctx context.Context, MotionCategoryID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/child_ids", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_category/%d/child_ids\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_ID(ctx context.Context, MotionCategoryID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/id", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_category/%d/id\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_Level(ctx context.Context, MotionCategoryID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/level", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_category/%d/level\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_MeetingID(ctx context.Context, MotionCategoryID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/meeting_id", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_category/%d/meeting_id\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_MotionIDs(ctx context.Context, MotionCategoryID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/motion_ids", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_category/%d/motion_ids\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_Name(ctx context.Context, MotionCategoryID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/name", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_category/%d/name\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_ParentID(ctx context.Context, MotionCategoryID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/parent_id", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_category/%d/parent_id\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_Prefix(ctx context.Context, MotionCategoryID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/prefix", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_category/%d/prefix\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionCategory_Weight(ctx context.Context, MotionCategoryID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/weight", MotionCategoryID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_category/%d/weight\": %w", MotionCategoryID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_CreationTime(ctx context.Context, MotionChangeRecommendationID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/creation_time", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_change_recommendation/%d/creation_time\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_ID(ctx context.Context, MotionChangeRecommendationID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/id", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_change_recommendation/%d/id\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_Internal(ctx context.Context, MotionChangeRecommendationID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/internal", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_change_recommendation/%d/internal\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_LineFrom(ctx context.Context, MotionChangeRecommendationID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/line_from", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_change_recommendation/%d/line_from\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_LineTo(ctx context.Context, MotionChangeRecommendationID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/line_to", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_change_recommendation/%d/line_to\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_MeetingID(ctx context.Context, MotionChangeRecommendationID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/meeting_id", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_change_recommendation/%d/meeting_id\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_MotionID(ctx context.Context, MotionChangeRecommendationID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/motion_id", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_change_recommendation/%d/motion_id\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_OtherDescription(ctx context.Context, MotionChangeRecommendationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/other_description", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_change_recommendation/%d/other_description\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_Rejected(ctx context.Context, MotionChangeRecommendationID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/rejected", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_change_recommendation/%d/rejected\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_Text(ctx context.Context, MotionChangeRecommendationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/text", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_change_recommendation/%d/text\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionChangeRecommendation_Type(ctx context.Context, MotionChangeRecommendationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/type", MotionChangeRecommendationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_change_recommendation/%d/type\": %w", MotionChangeRecommendationID, err)
	}

	return v, nil
}

func (f Fields) MotionCommentSection_CommentIDs(ctx context.Context, MotionCommentSectionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/comment_ids", MotionCommentSectionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_comment_section/%d/comment_ids\": %w", MotionCommentSectionID, err)
	}

	return v, nil
}

func (f Fields) MotionCommentSection_ID(ctx context.Context, MotionCommentSectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/id", MotionCommentSectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_comment_section/%d/id\": %w", MotionCommentSectionID, err)
	}

	return v, nil
}

func (f Fields) MotionCommentSection_MeetingID(ctx context.Context, MotionCommentSectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/meeting_id", MotionCommentSectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_comment_section/%d/meeting_id\": %w", MotionCommentSectionID, err)
	}

	return v, nil
}

func (f Fields) MotionCommentSection_Name(ctx context.Context, MotionCommentSectionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/name", MotionCommentSectionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_comment_section/%d/name\": %w", MotionCommentSectionID, err)
	}

	return v, nil
}

func (f Fields) MotionCommentSection_ReadGroupIDs(ctx context.Context, MotionCommentSectionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/read_group_ids", MotionCommentSectionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_comment_section/%d/read_group_ids\": %w", MotionCommentSectionID, err)
	}

	return v, nil
}

func (f Fields) MotionCommentSection_Weight(ctx context.Context, MotionCommentSectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/weight", MotionCommentSectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_comment_section/%d/weight\": %w", MotionCommentSectionID, err)
	}

	return v, nil
}

func (f Fields) MotionCommentSection_WriteGroupIDs(ctx context.Context, MotionCommentSectionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/write_group_ids", MotionCommentSectionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_comment_section/%d/write_group_ids\": %w", MotionCommentSectionID, err)
	}

	return v, nil
}

func (f Fields) MotionComment_Comment(ctx context.Context, MotionCommentID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/comment", MotionCommentID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_comment/%d/comment\": %w", MotionCommentID, err)
	}

	return v, nil
}

func (f Fields) MotionComment_ID(ctx context.Context, MotionCommentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/id", MotionCommentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_comment/%d/id\": %w", MotionCommentID, err)
	}

	return v, nil
}

func (f Fields) MotionComment_MeetingID(ctx context.Context, MotionCommentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/meeting_id", MotionCommentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_comment/%d/meeting_id\": %w", MotionCommentID, err)
	}

	return v, nil
}

func (f Fields) MotionComment_MotionID(ctx context.Context, MotionCommentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/motion_id", MotionCommentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_comment/%d/motion_id\": %w", MotionCommentID, err)
	}

	return v, nil
}

func (f Fields) MotionComment_SectionID(ctx context.Context, MotionCommentID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/section_id", MotionCommentID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_comment/%d/section_id\": %w", MotionCommentID, err)
	}

	return v, nil
}

func (f Fields) MotionState_AllowCreatePoll(ctx context.Context, MotionStateID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/allow_create_poll", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_state/%d/allow_create_poll\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_AllowSubmitterEdit(ctx context.Context, MotionStateID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/allow_submitter_edit", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_state/%d/allow_submitter_edit\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_AllowSupport(ctx context.Context, MotionStateID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/allow_support", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_state/%d/allow_support\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_CssClass(ctx context.Context, MotionStateID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/css_class", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_state/%d/css_class\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_FirstStateOfWorkflowID(ctx context.Context, MotionStateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/first_state_of_workflow_id", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_state/%d/first_state_of_workflow_id\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_ID(ctx context.Context, MotionStateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/id", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_state/%d/id\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_MeetingID(ctx context.Context, MotionStateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/meeting_id", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_state/%d/meeting_id\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_MergeAmendmentIntoFinal(ctx context.Context, MotionStateID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/merge_amendment_into_final", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_state/%d/merge_amendment_into_final\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_MotionIDs(ctx context.Context, MotionStateID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/motion_ids", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_state/%d/motion_ids\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_MotionRecommendationIDs(ctx context.Context, MotionStateID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/motion_recommendation_ids", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_state/%d/motion_recommendation_ids\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_Name(ctx context.Context, MotionStateID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/name", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_state/%d/name\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_NextStateIDs(ctx context.Context, MotionStateID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/next_state_ids", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_state/%d/next_state_ids\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_PreviousStateIDs(ctx context.Context, MotionStateID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/previous_state_ids", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_state/%d/previous_state_ids\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_RecommendationLabel(ctx context.Context, MotionStateID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/recommendation_label", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_state/%d/recommendation_label\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_Restrictions(ctx context.Context, MotionStateID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/restrictions", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_state/%d/restrictions\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_SetNumber(ctx context.Context, MotionStateID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/set_number", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_state/%d/set_number\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_ShowRecommendationExtensionField(ctx context.Context, MotionStateID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/show_recommendation_extension_field", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_state/%d/show_recommendation_extension_field\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_ShowStateExtensionField(ctx context.Context, MotionStateID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/show_state_extension_field", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"motion_state/%d/show_state_extension_field\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionState_WorkflowID(ctx context.Context, MotionStateID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/workflow_id", MotionStateID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_state/%d/workflow_id\": %w", MotionStateID, err)
	}

	return v, nil
}

func (f Fields) MotionStatuteParagraph_ID(ctx context.Context, MotionStatuteParagraphID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/id", MotionStatuteParagraphID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_statute_paragraph/%d/id\": %w", MotionStatuteParagraphID, err)
	}

	return v, nil
}

func (f Fields) MotionStatuteParagraph_MeetingID(ctx context.Context, MotionStatuteParagraphID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/meeting_id", MotionStatuteParagraphID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_statute_paragraph/%d/meeting_id\": %w", MotionStatuteParagraphID, err)
	}

	return v, nil
}

func (f Fields) MotionStatuteParagraph_MotionIDs(ctx context.Context, MotionStatuteParagraphID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/motion_ids", MotionStatuteParagraphID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_statute_paragraph/%d/motion_ids\": %w", MotionStatuteParagraphID, err)
	}

	return v, nil
}

func (f Fields) MotionStatuteParagraph_Text(ctx context.Context, MotionStatuteParagraphID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/text", MotionStatuteParagraphID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_statute_paragraph/%d/text\": %w", MotionStatuteParagraphID, err)
	}

	return v, nil
}

func (f Fields) MotionStatuteParagraph_Title(ctx context.Context, MotionStatuteParagraphID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/title", MotionStatuteParagraphID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_statute_paragraph/%d/title\": %w", MotionStatuteParagraphID, err)
	}

	return v, nil
}

func (f Fields) MotionStatuteParagraph_Weight(ctx context.Context, MotionStatuteParagraphID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/weight", MotionStatuteParagraphID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_statute_paragraph/%d/weight\": %w", MotionStatuteParagraphID, err)
	}

	return v, nil
}

func (f Fields) MotionSubmitter_ID(ctx context.Context, MotionSubmitterID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/id", MotionSubmitterID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_submitter/%d/id\": %w", MotionSubmitterID, err)
	}

	return v, nil
}

func (f Fields) MotionSubmitter_MeetingID(ctx context.Context, MotionSubmitterID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/meeting_id", MotionSubmitterID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_submitter/%d/meeting_id\": %w", MotionSubmitterID, err)
	}

	return v, nil
}

func (f Fields) MotionSubmitter_MotionID(ctx context.Context, MotionSubmitterID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/motion_id", MotionSubmitterID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_submitter/%d/motion_id\": %w", MotionSubmitterID, err)
	}

	return v, nil
}

func (f Fields) MotionSubmitter_UserID(ctx context.Context, MotionSubmitterID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/user_id", MotionSubmitterID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_submitter/%d/user_id\": %w", MotionSubmitterID, err)
	}

	return v, nil
}

func (f Fields) MotionSubmitter_Weight(ctx context.Context, MotionSubmitterID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/weight", MotionSubmitterID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_submitter/%d/weight\": %w", MotionSubmitterID, err)
	}

	return v, nil
}

func (f Fields) MotionWorkflow_DefaultAmendmentWorkflowMeetingID(ctx context.Context, MotionWorkflowID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/default_amendment_workflow_meeting_id", MotionWorkflowID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_workflow/%d/default_amendment_workflow_meeting_id\": %w", MotionWorkflowID, err)
	}

	return v, nil
}

func (f Fields) MotionWorkflow_DefaultStatuteAmendmentWorkflowMeetingID(ctx context.Context, MotionWorkflowID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/default_statute_amendment_workflow_meeting_id", MotionWorkflowID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_workflow/%d/default_statute_amendment_workflow_meeting_id\": %w", MotionWorkflowID, err)
	}

	return v, nil
}

func (f Fields) MotionWorkflow_DefaultWorkflowMeetingID(ctx context.Context, MotionWorkflowID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/default_workflow_meeting_id", MotionWorkflowID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_workflow/%d/default_workflow_meeting_id\": %w", MotionWorkflowID, err)
	}

	return v, nil
}

func (f Fields) MotionWorkflow_FirstStateID(ctx context.Context, MotionWorkflowID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/first_state_id", MotionWorkflowID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_workflow/%d/first_state_id\": %w", MotionWorkflowID, err)
	}

	return v, nil
}

func (f Fields) MotionWorkflow_ID(ctx context.Context, MotionWorkflowID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/id", MotionWorkflowID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_workflow/%d/id\": %w", MotionWorkflowID, err)
	}

	return v, nil
}

func (f Fields) MotionWorkflow_MeetingID(ctx context.Context, MotionWorkflowID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/meeting_id", MotionWorkflowID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion_workflow/%d/meeting_id\": %w", MotionWorkflowID, err)
	}

	return v, nil
}

func (f Fields) MotionWorkflow_Name(ctx context.Context, MotionWorkflowID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/name", MotionWorkflowID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion_workflow/%d/name\": %w", MotionWorkflowID, err)
	}

	return v, nil
}

func (f Fields) MotionWorkflow_StateIDs(ctx context.Context, MotionWorkflowID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/state_ids", MotionWorkflowID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion_workflow/%d/state_ids\": %w", MotionWorkflowID, err)
	}

	return v, nil
}

func (f Fields) Motion_AgendaItemID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/agenda_item_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/agenda_item_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_AllDerivedMotionIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/all_derived_motion_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/all_derived_motion_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_AllOriginIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/all_origin_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/all_origin_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_AmendmentIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/amendment_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/amendment_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_AmendmentParagraphTmpl(ctx context.Context, MotionID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/amendment_paragraph_$", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/amendment_paragraph_$\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_AmendmentParagraph(ctx context.Context, MotionID int, replacement string) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/amendment_paragraph_$%s", MotionID, replacement)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion/%d/amendment_paragraph_$%s\": %w", MotionID, replacement, err)
	}

	return v, nil
}

func (f Fields) Motion_AttachmentIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/attachment_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/attachment_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_BlockID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/block_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/block_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_CategoryID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/category_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/category_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_CategoryWeight(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/category_weight", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/category_weight\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_ChangeRecommendationIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/change_recommendation_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/change_recommendation_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_CommentIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/comment_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/comment_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_Created(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/created", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/created\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_DerivedMotionIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/derived_motion_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/derived_motion_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_ID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_LastModified(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/last_modified", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/last_modified\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_LeadMotionID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/lead_motion_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/lead_motion_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_ListOfSpeakersID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/list_of_speakers_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/list_of_speakers_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_MeetingID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/meeting_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/meeting_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_ModifiedFinalVersion(ctx context.Context, MotionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/modified_final_version", MotionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion/%d/modified_final_version\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_Number(ctx context.Context, MotionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/number", MotionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion/%d/number\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_NumberValue(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/number_value", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/number_value\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_OptionIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/option_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/option_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_OriginID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/origin_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/origin_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_PersonalNoteIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/personal_note_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/personal_note_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_PollIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/poll_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/poll_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_ProjectionIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/projection_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/projection_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_Reason(ctx context.Context, MotionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/reason", MotionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion/%d/reason\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_RecommendationExtension(ctx context.Context, MotionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/recommendation_extension", MotionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion/%d/recommendation_extension\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_RecommendationExtensionReferenceIDs(ctx context.Context, MotionID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/recommendation_extension_reference_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/recommendation_extension_reference_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_RecommendationID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/recommendation_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/recommendation_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_ReferencedInMotionRecommendationExtensionIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/referenced_in_motion_recommendation_extension_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/referenced_in_motion_recommendation_extension_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_SequentialNumber(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/sequential_number", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/sequential_number\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_SortChildIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/sort_child_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/sort_child_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_SortParentID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/sort_parent_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/sort_parent_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_SortWeight(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/sort_weight", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/sort_weight\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_StateExtension(ctx context.Context, MotionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/state_extension", MotionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion/%d/state_extension\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_StateID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/state_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/state_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_StatuteParagraphID(ctx context.Context, MotionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/statute_paragraph_id", MotionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"motion/%d/statute_paragraph_id\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_SubmitterIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/submitter_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/submitter_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_SupporterIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/supporter_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/supporter_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_TagIDs(ctx context.Context, MotionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/tag_ids", MotionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"motion/%d/tag_ids\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_Text(ctx context.Context, MotionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/text", MotionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion/%d/text\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Motion_Title(ctx context.Context, MotionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/title", MotionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"motion/%d/title\": %w", MotionID, err)
	}

	return v, nil
}

func (f Fields) Option_Abstain(ctx context.Context, OptionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/abstain", OptionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"option/%d/abstain\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_ContentObjectID(ctx context.Context, OptionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "option/%d/content_object_id", OptionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"option/%d/content_object_id\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_ID(ctx context.Context, OptionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/id", OptionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"option/%d/id\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_MeetingID(ctx context.Context, OptionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/meeting_id", OptionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"option/%d/meeting_id\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_No(ctx context.Context, OptionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/no", OptionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"option/%d/no\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_PollID(ctx context.Context, OptionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/poll_id", OptionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"option/%d/poll_id\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_Text(ctx context.Context, OptionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "option/%d/text", OptionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"option/%d/text\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_UsedAsGlobalOptionInPollID(ctx context.Context, OptionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/used_as_global_option_in_poll_id", OptionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"option/%d/used_as_global_option_in_poll_id\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_VoteIDs(ctx context.Context, OptionID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/vote_ids", OptionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"option/%d/vote_ids\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_Weight(ctx context.Context, OptionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/weight", OptionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"option/%d/weight\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) Option_Yes(ctx context.Context, OptionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/yes", OptionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"option/%d/yes\": %w", OptionID, err)
	}

	return v, nil
}

func (f Fields) OrganizationTag_Color(ctx context.Context, OrganizationTagID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/color", OrganizationTagID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"organization_tag/%d/color\": %w", OrganizationTagID, err)
	}

	return v, nil
}

func (f Fields) OrganizationTag_ID(ctx context.Context, OrganizationTagID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/id", OrganizationTagID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"organization_tag/%d/id\": %w", OrganizationTagID, err)
	}

	return v, nil
}

func (f Fields) OrganizationTag_Name(ctx context.Context, OrganizationTagID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/name", OrganizationTagID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"organization_tag/%d/name\": %w", OrganizationTagID, err)
	}

	return v, nil
}

func (f Fields) OrganizationTag_OrganizationID(ctx context.Context, OrganizationTagID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/organization_id", OrganizationTagID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"organization_tag/%d/organization_id\": %w", OrganizationTagID, err)
	}

	return v, nil
}

func (f Fields) OrganizationTag_TaggedIDs(ctx context.Context, OrganizationTagID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/tagged_ids", OrganizationTagID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"organization_tag/%d/tagged_ids\": %w", OrganizationTagID, err)
	}

	return v, nil
}

func (f Fields) Organization_CommitteeIDs(ctx context.Context, OrganizationID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/committee_ids", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"organization/%d/committee_ids\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_Description(ctx context.Context, OrganizationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/description", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"organization/%d/description\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_EnableElectronicVoting(ctx context.Context, OrganizationID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/enable_electronic_voting", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"organization/%d/enable_electronic_voting\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_ID(ctx context.Context, OrganizationID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/id", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"organization/%d/id\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_LegalNotice(ctx context.Context, OrganizationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/legal_notice", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"organization/%d/legal_notice\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_LoginText(ctx context.Context, OrganizationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/login_text", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"organization/%d/login_text\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_Name(ctx context.Context, OrganizationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/name", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"organization/%d/name\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_OrganizationTagIDs(ctx context.Context, OrganizationID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/organization_tag_ids", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"organization/%d/organization_tag_ids\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_PrivacyPolicy(ctx context.Context, OrganizationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/privacy_policy", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"organization/%d/privacy_policy\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_ResetPasswordVerboseErrors(ctx context.Context, OrganizationID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/reset_password_verbose_errors", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"organization/%d/reset_password_verbose_errors\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_ResourceIDs(ctx context.Context, OrganizationID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/resource_ids", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"organization/%d/resource_ids\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) Organization_Theme(ctx context.Context, OrganizationID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/theme", OrganizationID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"organization/%d/theme\": %w", OrganizationID, err)
	}

	return v, nil
}

func (f Fields) PersonalNote_ContentObjectID(ctx context.Context, PersonalNoteID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/content_object_id", PersonalNoteID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"personal_note/%d/content_object_id\": %w", PersonalNoteID, err)
	}

	return v, nil
}

func (f Fields) PersonalNote_ID(ctx context.Context, PersonalNoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/id", PersonalNoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"personal_note/%d/id\": %w", PersonalNoteID, err)
	}

	return v, nil
}

func (f Fields) PersonalNote_MeetingID(ctx context.Context, PersonalNoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/meeting_id", PersonalNoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"personal_note/%d/meeting_id\": %w", PersonalNoteID, err)
	}

	return v, nil
}

func (f Fields) PersonalNote_Note(ctx context.Context, PersonalNoteID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/note", PersonalNoteID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"personal_note/%d/note\": %w", PersonalNoteID, err)
	}

	return v, nil
}

func (f Fields) PersonalNote_Star(ctx context.Context, PersonalNoteID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/star", PersonalNoteID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"personal_note/%d/star\": %w", PersonalNoteID, err)
	}

	return v, nil
}

func (f Fields) PersonalNote_UserID(ctx context.Context, PersonalNoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/user_id", PersonalNoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"personal_note/%d/user_id\": %w", PersonalNoteID, err)
	}

	return v, nil
}

func (f Fields) Poll_Backend(ctx context.Context, PollID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/backend", PollID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"poll/%d/backend\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_ContentObjectID(ctx context.Context, PollID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/content_object_id", PollID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"poll/%d/content_object_id\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_Description(ctx context.Context, PollID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/description", PollID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"poll/%d/description\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_EntitledGroupIDs(ctx context.Context, PollID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/entitled_group_ids", PollID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"poll/%d/entitled_group_ids\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_EntitledUsersAtStop(ctx context.Context, PollID int) (json.RawMessage, error) {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/entitled_users_at_stop", PollID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"poll/%d/entitled_users_at_stop\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_GlobalAbstain(ctx context.Context, PollID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/global_abstain", PollID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"poll/%d/global_abstain\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_GlobalNo(ctx context.Context, PollID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/global_no", PollID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"poll/%d/global_no\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_GlobalOptionID(ctx context.Context, PollID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/global_option_id", PollID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"poll/%d/global_option_id\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_GlobalYes(ctx context.Context, PollID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/global_yes", PollID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"poll/%d/global_yes\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_ID(ctx context.Context, PollID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/id", PollID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"poll/%d/id\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_IsPseudoanonymized(ctx context.Context, PollID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/is_pseudoanonymized", PollID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"poll/%d/is_pseudoanonymized\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_MaxVotesAmount(ctx context.Context, PollID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/max_votes_amount", PollID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"poll/%d/max_votes_amount\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_MeetingID(ctx context.Context, PollID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/meeting_id", PollID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"poll/%d/meeting_id\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_MinVotesAmount(ctx context.Context, PollID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/min_votes_amount", PollID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"poll/%d/min_votes_amount\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_OnehundredPercentBase(ctx context.Context, PollID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/onehundred_percent_base", PollID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"poll/%d/onehundred_percent_base\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_OptionIDs(ctx context.Context, PollID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/option_ids", PollID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"poll/%d/option_ids\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_Pollmethod(ctx context.Context, PollID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/pollmethod", PollID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"poll/%d/pollmethod\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_ProjectionIDs(ctx context.Context, PollID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/projection_ids", PollID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"poll/%d/projection_ids\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_State(ctx context.Context, PollID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/state", PollID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"poll/%d/state\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_Title(ctx context.Context, PollID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/title", PollID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"poll/%d/title\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_Type(ctx context.Context, PollID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/type", PollID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"poll/%d/type\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_VotedIDs(ctx context.Context, PollID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/voted_ids", PollID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"poll/%d/voted_ids\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_Votescast(ctx context.Context, PollID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/votescast", PollID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"poll/%d/votescast\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_Votesinvalid(ctx context.Context, PollID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/votesinvalid", PollID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"poll/%d/votesinvalid\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Poll_Votesvalid(ctx context.Context, PollID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/votesvalid", PollID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"poll/%d/votesvalid\": %w", PollID, err)
	}

	return v, nil
}

func (f Fields) Projection_Content(ctx context.Context, ProjectionID int) (json.RawMessage, error) {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/content", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"projection/%d/content\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_ContentObjectID(ctx context.Context, ProjectionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/content_object_id", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projection/%d/content_object_id\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_CurrentProjectorID(ctx context.Context, ProjectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/current_projector_id", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projection/%d/current_projector_id\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_HistoryProjectorID(ctx context.Context, ProjectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/history_projector_id", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projection/%d/history_projector_id\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_ID(ctx context.Context, ProjectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/id", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projection/%d/id\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_MeetingID(ctx context.Context, ProjectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/meeting_id", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projection/%d/meeting_id\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_Options(ctx context.Context, ProjectionID int) (json.RawMessage, error) {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/options", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"projection/%d/options\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_PreviewProjectorID(ctx context.Context, ProjectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/preview_projector_id", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projection/%d/preview_projector_id\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_Stable(ctx context.Context, ProjectionID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/stable", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"projection/%d/stable\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_Type(ctx context.Context, ProjectionID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/type", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projection/%d/type\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) Projection_Weight(ctx context.Context, ProjectionID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/weight", ProjectionID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projection/%d/weight\": %w", ProjectionID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_CountdownTime(ctx context.Context, ProjectorCountdownID int) (float32, error) {
	var v float32
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/countdown_time", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector_countdown/%d/countdown_time\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_DefaultTime(ctx context.Context, ProjectorCountdownID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/default_time", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector_countdown/%d/default_time\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_Description(ctx context.Context, ProjectorCountdownID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/description", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector_countdown/%d/description\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_ID(ctx context.Context, ProjectorCountdownID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/id", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector_countdown/%d/id\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_MeetingID(ctx context.Context, ProjectorCountdownID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/meeting_id", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector_countdown/%d/meeting_id\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_ProjectionIDs(ctx context.Context, ProjectorCountdownID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/projection_ids", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"projector_countdown/%d/projection_ids\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_Running(ctx context.Context, ProjectorCountdownID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/running", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"projector_countdown/%d/running\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_Title(ctx context.Context, ProjectorCountdownID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/title", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector_countdown/%d/title\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_UsedAsListOfSpeakerCountdownMeetingID(ctx context.Context, ProjectorCountdownID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/used_as_list_of_speaker_countdown_meeting_id", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector_countdown/%d/used_as_list_of_speaker_countdown_meeting_id\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorCountdown_UsedAsPollCountdownMeetingID(ctx context.Context, ProjectorCountdownID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/used_as_poll_countdown_meeting_id", ProjectorCountdownID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector_countdown/%d/used_as_poll_countdown_meeting_id\": %w", ProjectorCountdownID, err)
	}

	return v, nil
}

func (f Fields) ProjectorMessage_ID(ctx context.Context, ProjectorMessageID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_message/%d/id", ProjectorMessageID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector_message/%d/id\": %w", ProjectorMessageID, err)
	}

	return v, nil
}

func (f Fields) ProjectorMessage_MeetingID(ctx context.Context, ProjectorMessageID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_message/%d/meeting_id", ProjectorMessageID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector_message/%d/meeting_id\": %w", ProjectorMessageID, err)
	}

	return v, nil
}

func (f Fields) ProjectorMessage_Message(ctx context.Context, ProjectorMessageID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector_message/%d/message", ProjectorMessageID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector_message/%d/message\": %w", ProjectorMessageID, err)
	}

	return v, nil
}

func (f Fields) ProjectorMessage_ProjectionIDs(ctx context.Context, ProjectorMessageID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector_message/%d/projection_ids", ProjectorMessageID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"projector_message/%d/projection_ids\": %w", ProjectorMessageID, err)
	}

	return v, nil
}

func (f Fields) Projector_AspectRatioDenominator(ctx context.Context, ProjectorID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/aspect_ratio_denominator", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector/%d/aspect_ratio_denominator\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_AspectRatioNumerator(ctx context.Context, ProjectorID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/aspect_ratio_numerator", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector/%d/aspect_ratio_numerator\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_BackgroundColor(ctx context.Context, ProjectorID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/background_color", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/background_color\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_ChyronBackgroundColor(ctx context.Context, ProjectorID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/chyron_background_color", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/chyron_background_color\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_ChyronFontColor(ctx context.Context, ProjectorID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/chyron_font_color", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/chyron_font_color\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_Color(ctx context.Context, ProjectorID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/color", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/color\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_CurrentProjectionIDs(ctx context.Context, ProjectorID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/current_projection_ids", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"projector/%d/current_projection_ids\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_HeaderBackgroundColor(ctx context.Context, ProjectorID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/header_background_color", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/header_background_color\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_HeaderFontColor(ctx context.Context, ProjectorID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/header_font_color", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/header_font_color\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_HeaderH1Color(ctx context.Context, ProjectorID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/header_h1_color", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/header_h1_color\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_HistoryProjectionIDs(ctx context.Context, ProjectorID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/history_projection_ids", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"projector/%d/history_projection_ids\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_ID(ctx context.Context, ProjectorID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/id", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector/%d/id\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_MeetingID(ctx context.Context, ProjectorID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/meeting_id", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector/%d/meeting_id\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_Name(ctx context.Context, ProjectorID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/name", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/name\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_PreviewProjectionIDs(ctx context.Context, ProjectorID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/preview_projection_ids", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"projector/%d/preview_projection_ids\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_Scale(ctx context.Context, ProjectorID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/scale", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector/%d/scale\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_Scroll(ctx context.Context, ProjectorID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/scroll", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector/%d/scroll\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_ShowClock(ctx context.Context, ProjectorID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/show_clock", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"projector/%d/show_clock\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_ShowHeaderFooter(ctx context.Context, ProjectorID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/show_header_footer", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"projector/%d/show_header_footer\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_ShowLogo(ctx context.Context, ProjectorID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/show_logo", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"projector/%d/show_logo\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_ShowTitle(ctx context.Context, ProjectorID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/show_title", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"projector/%d/show_title\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_UsedAsDefaultInMeetingIDTmpl(ctx context.Context, ProjectorID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/used_as_default_$_in_meeting_id", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"projector/%d/used_as_default_$_in_meeting_id\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_UsedAsDefaultInMeetingID(ctx context.Context, ProjectorID int, replacement string) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/used_as_default_$%s_in_meeting_id", ProjectorID, replacement)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"projector/%d/used_as_default_$%s_in_meeting_id\": %w", ProjectorID, replacement, err)
	}

	return v, nil
}

func (f Fields) Projector_UsedAsReferenceProjectorMeetingID(ctx context.Context, ProjectorID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/used_as_reference_projector_meeting_id", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector/%d/used_as_reference_projector_meeting_id\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Projector_Width(ctx context.Context, ProjectorID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/width", ProjectorID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"projector/%d/width\": %w", ProjectorID, err)
	}

	return v, nil
}

func (f Fields) Resource_Filesize(ctx context.Context, ResourceID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/filesize", ResourceID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"resource/%d/filesize\": %w", ResourceID, err)
	}

	return v, nil
}

func (f Fields) Resource_ID(ctx context.Context, ResourceID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/id", ResourceID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"resource/%d/id\": %w", ResourceID, err)
	}

	return v, nil
}

func (f Fields) Resource_Mimetype(ctx context.Context, ResourceID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/mimetype", ResourceID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"resource/%d/mimetype\": %w", ResourceID, err)
	}

	return v, nil
}

func (f Fields) Resource_OrganizationID(ctx context.Context, ResourceID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/organization_id", ResourceID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"resource/%d/organization_id\": %w", ResourceID, err)
	}

	return v, nil
}

func (f Fields) Resource_Token(ctx context.Context, ResourceID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/token", ResourceID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"resource/%d/token\": %w", ResourceID, err)
	}

	return v, nil
}

func (f Fields) Speaker_BeginTime(ctx context.Context, SpeakerID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/begin_time", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"speaker/%d/begin_time\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_EndTime(ctx context.Context, SpeakerID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/end_time", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"speaker/%d/end_time\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_ID(ctx context.Context, SpeakerID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/id", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"speaker/%d/id\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_ListOfSpeakersID(ctx context.Context, SpeakerID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/list_of_speakers_id", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"speaker/%d/list_of_speakers_id\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_MeetingID(ctx context.Context, SpeakerID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/meeting_id", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"speaker/%d/meeting_id\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_Note(ctx context.Context, SpeakerID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/note", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"speaker/%d/note\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_PointOfOrder(ctx context.Context, SpeakerID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/point_of_order", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"speaker/%d/point_of_order\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_SpeechState(ctx context.Context, SpeakerID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/speech_state", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"speaker/%d/speech_state\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_UserID(ctx context.Context, SpeakerID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/user_id", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"speaker/%d/user_id\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Speaker_Weight(ctx context.Context, SpeakerID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/weight", SpeakerID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"speaker/%d/weight\": %w", SpeakerID, err)
	}

	return v, nil
}

func (f Fields) Tag_ID(ctx context.Context, TagID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "tag/%d/id", TagID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"tag/%d/id\": %w", TagID, err)
	}

	return v, nil
}

func (f Fields) Tag_MeetingID(ctx context.Context, TagID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "tag/%d/meeting_id", TagID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"tag/%d/meeting_id\": %w", TagID, err)
	}

	return v, nil
}

func (f Fields) Tag_Name(ctx context.Context, TagID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "tag/%d/name", TagID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"tag/%d/name\": %w", TagID, err)
	}

	return v, nil
}

func (f Fields) Tag_TaggedIDs(ctx context.Context, TagID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "tag/%d/tagged_ids", TagID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"tag/%d/tagged_ids\": %w", TagID, err)
	}

	return v, nil
}

func (f Fields) Topic_AgendaItemID(ctx context.Context, TopicID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/agenda_item_id", TopicID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"topic/%d/agenda_item_id\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_AttachmentIDs(ctx context.Context, TopicID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/attachment_ids", TopicID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"topic/%d/attachment_ids\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_ID(ctx context.Context, TopicID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/id", TopicID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"topic/%d/id\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_ListOfSpeakersID(ctx context.Context, TopicID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/list_of_speakers_id", TopicID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"topic/%d/list_of_speakers_id\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_MeetingID(ctx context.Context, TopicID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/meeting_id", TopicID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"topic/%d/meeting_id\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_OptionIDs(ctx context.Context, TopicID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/option_ids", TopicID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"topic/%d/option_ids\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_ProjectionIDs(ctx context.Context, TopicID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/projection_ids", TopicID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"topic/%d/projection_ids\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_TagIDs(ctx context.Context, TopicID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/tag_ids", TopicID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"topic/%d/tag_ids\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_Text(ctx context.Context, TopicID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/text", TopicID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"topic/%d/text\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) Topic_Title(ctx context.Context, TopicID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/title", TopicID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"topic/%d/title\": %w", TopicID, err)
	}

	return v, nil
}

func (f Fields) User_AboutMeTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/about_me_$", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/about_me_$\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_AboutMe(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/about_me_$%d", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/about_me_$%d\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_AssignmentCandidateIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/assignment_candidate_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/assignment_candidate_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_AssignmentCandidateIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/assignment_candidate_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/assignment_candidate_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_CanChangeOwnPassword(ctx context.Context, UserID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "user/%d/can_change_own_password", UserID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"user/%d/can_change_own_password\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_CommentTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/comment_$", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/comment_$\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_Comment(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/comment_$%d", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/comment_$%d\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_CommitteeIDs(ctx context.Context, UserID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/committee_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/committee_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_CommitteeManagementLevelTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/committee_$_management_level", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/committee_$_management_level\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_CommitteeManagementLevel(ctx context.Context, UserID int, committeeID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/committee_$%d_management_level", UserID, committeeID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/committee_$%d_management_level\": %w", UserID, committeeID, err)
	}

	return v, nil
}

func (f Fields) User_DefaultNumber(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/default_number", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/default_number\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_DefaultPassword(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/default_password", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/default_password\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_DefaultStructureLevel(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/default_structure_level", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/default_structure_level\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_DefaultVoteWeight(ctx context.Context, UserID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/default_vote_weight", UserID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"user/%d/default_vote_weight\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_Email(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/email", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/email\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_FirstName(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/first_name", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/first_name\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_Gender(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/gender", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/gender\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_GroupIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/group_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/group_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_GroupIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/group_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/group_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_ID(ctx context.Context, UserID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/id", UserID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"user/%d/id\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_IsActive(ctx context.Context, UserID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "user/%d/is_active", UserID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"user/%d/is_active\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_IsDemoUser(ctx context.Context, UserID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "user/%d/is_demo_user", UserID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"user/%d/is_demo_user\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_IsPhysicalPerson(ctx context.Context, UserID int) (bool, error) {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "user/%d/is_physical_person", UserID)
	if err := f.fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching \"user/%d/is_physical_person\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_IsPresentInMeetingIDs(ctx context.Context, UserID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/is_present_in_meeting_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/is_present_in_meeting_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_LastEmailSend(ctx context.Context, UserID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/last_email_send", UserID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"user/%d/last_email_send\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_LastName(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/last_name", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/last_name\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_MeetingIDs(ctx context.Context, UserID int) ([]int, error) {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/meeting_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/meeting_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_NumberTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/number_$", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/number_$\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_Number(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/number_$%d", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/number_$%d\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_OptionIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/option_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/option_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_OptionIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/option_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/option_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_OrganizationManagementLevel(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/organization_management_level", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/organization_management_level\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_Password(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/password", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/password\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_PersonalNoteIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/personal_note_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/personal_note_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_PersonalNoteIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/personal_note_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/personal_note_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_PollVotedIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/poll_voted_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/poll_voted_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_PollVotedIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/poll_voted_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/poll_voted_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_ProjectionIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/projection_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/projection_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_ProjectionIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/projection_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/projection_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_SpeakerIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/speaker_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/speaker_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_SpeakerIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/speaker_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/speaker_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_StructureLevelTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/structure_level_$", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/structure_level_$\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_StructureLevel(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/structure_level_$%d", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/structure_level_$%d\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_SubmittedMotionIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/submitted_motion_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/submitted_motion_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_SubmittedMotionIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/submitted_motion_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/submitted_motion_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_SupportedMotionIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/supported_motion_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/supported_motion_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_SupportedMotionIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/supported_motion_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/supported_motion_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_Title(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/title", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/title\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_Username(ctx context.Context, UserID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/username", UserID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/username\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_VoteDelegatedToIDTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegated_$_to_id", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/vote_delegated_$_to_id\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_VoteDelegatedToID(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegated_$%d_to_id", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/vote_delegated_$%d_to_id\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_VoteDelegatedVoteIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegated_vote_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/vote_delegated_vote_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_VoteDelegatedVoteIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegated_vote_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/vote_delegated_vote_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_VoteDelegationsFromIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegations_$_from_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/vote_delegations_$_from_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_VoteDelegationsFromIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegations_$%d_from_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/vote_delegations_$%d_from_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_VoteIDsTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_$_ids", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/vote_$_ids\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_VoteIDs(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_$%d_ids", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/vote_$%d_ids\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) User_VoteWeightTmpl(ctx context.Context, UserID int) ([]string, error) {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_weight_$", UserID)
	if err := f.fetch.Err(); err != nil {
		return nil, fmt.Errorf("fetching \"user/%d/vote_weight_$\": %w", UserID, err)
	}

	return v, nil
}

func (f Fields) User_VoteWeight(ctx context.Context, UserID int, meetingID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_weight_$%d", UserID, meetingID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"user/%d/vote_weight_$%d\": %w", UserID, meetingID, err)
	}

	return v, nil
}

func (f Fields) Vote_DelegatedUserID(ctx context.Context, VoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/delegated_user_id", VoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"vote/%d/delegated_user_id\": %w", VoteID, err)
	}

	return v, nil
}

func (f Fields) Vote_ID(ctx context.Context, VoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/id", VoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"vote/%d/id\": %w", VoteID, err)
	}

	return v, nil
}

func (f Fields) Vote_MeetingID(ctx context.Context, VoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/meeting_id", VoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"vote/%d/meeting_id\": %w", VoteID, err)
	}

	return v, nil
}

func (f Fields) Vote_OptionID(ctx context.Context, VoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/option_id", VoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"vote/%d/option_id\": %w", VoteID, err)
	}

	return v, nil
}

func (f Fields) Vote_UserID(ctx context.Context, VoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/user_id", VoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"vote/%d/user_id\": %w", VoteID, err)
	}

	return v, nil
}

func (f Fields) Vote_UserToken(ctx context.Context, VoteID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/user_token", VoteID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"vote/%d/user_token\": %w", VoteID, err)
	}

	return v, nil
}

func (f Fields) Vote_Value(ctx context.Context, VoteID int) (string, error) {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/value", VoteID)
	if err := f.fetch.Err(); err != nil {
		return "", fmt.Errorf("fetching \"vote/%d/value\": %w", VoteID, err)
	}

	return v, nil
}

func (f Fields) Vote_Weight(ctx context.Context, VoteID int) (int, error) {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/weight", VoteID)
	if err := f.fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching \"vote/%d/weight\": %w", VoteID, err)
	}

	return v, nil
}
