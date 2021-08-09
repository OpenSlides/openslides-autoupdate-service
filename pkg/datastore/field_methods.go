// Code generated with go generate DO NOT EDIT.
package datastore

import (
	"context"
	"encoding/json"
)

func (f Fields) AgendaItem_ChildIDs(ctx context.Context, agendaItemID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/child_ids", agendaItemID)
	return v
}

func (f Fields) AgendaItem_Closed(ctx context.Context, agendaItemID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/closed", agendaItemID)
	return v
}

func (f Fields) AgendaItem_Comment(ctx context.Context, agendaItemID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/comment", agendaItemID)
	return v
}

func (f Fields) AgendaItem_ContentObjectID(ctx context.Context, agendaItemID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/content_object_id", agendaItemID)
	return v
}

func (f Fields) AgendaItem_Duration(ctx context.Context, agendaItemID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/duration", agendaItemID)
	return v
}

func (f Fields) AgendaItem_ID(ctx context.Context, agendaItemID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/id", agendaItemID)
	return v
}

func (f Fields) AgendaItem_IsHidden(ctx context.Context, agendaItemID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/is_hidden", agendaItemID)
	return v
}

func (f Fields) AgendaItem_IsInternal(ctx context.Context, agendaItemID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/is_internal", agendaItemID)
	return v
}

func (f Fields) AgendaItem_ItemNumber(ctx context.Context, agendaItemID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/item_number", agendaItemID)
	return v
}

func (f Fields) AgendaItem_Level(ctx context.Context, agendaItemID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/level", agendaItemID)
	return v
}

func (f Fields) AgendaItem_MeetingID(ctx context.Context, agendaItemID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/meeting_id", agendaItemID)
	return v
}

func (f Fields) AgendaItem_ParentID(ctx context.Context, agendaItemID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/parent_id", agendaItemID)
	return v
}

func (f Fields) AgendaItem_ProjectionIDs(ctx context.Context, agendaItemID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/projection_ids", agendaItemID)
	return v
}

func (f Fields) AgendaItem_TagIDs(ctx context.Context, agendaItemID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/tag_ids", agendaItemID)
	return v
}

func (f Fields) AgendaItem_Type(ctx context.Context, agendaItemID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/type", agendaItemID)
	return v
}

func (f Fields) AgendaItem_Weight(ctx context.Context, agendaItemID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "agenda_item/%d/weight", agendaItemID)
	return v
}

func (f Fields) AssignmentCandidate_AssignmentID(ctx context.Context, assignmentCandidateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/assignment_id", assignmentCandidateID)
	return v
}

func (f Fields) AssignmentCandidate_ID(ctx context.Context, assignmentCandidateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/id", assignmentCandidateID)
	return v
}

func (f Fields) AssignmentCandidate_MeetingID(ctx context.Context, assignmentCandidateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/meeting_id", assignmentCandidateID)
	return v
}

func (f Fields) AssignmentCandidate_UserID(ctx context.Context, assignmentCandidateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/user_id", assignmentCandidateID)
	return v
}

func (f Fields) AssignmentCandidate_Weight(ctx context.Context, assignmentCandidateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment_candidate/%d/weight", assignmentCandidateID)
	return v
}

func (f Fields) Assignment_AgendaItemID(ctx context.Context, assignmentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/agenda_item_id", assignmentID)
	return v
}

func (f Fields) Assignment_AttachmentIDs(ctx context.Context, assignmentID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/attachment_ids", assignmentID)
	return v
}

func (f Fields) Assignment_CandidateIDs(ctx context.Context, assignmentID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/candidate_ids", assignmentID)
	return v
}

func (f Fields) Assignment_DefaultPollDescription(ctx context.Context, assignmentID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/default_poll_description", assignmentID)
	return v
}

func (f Fields) Assignment_Description(ctx context.Context, assignmentID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/description", assignmentID)
	return v
}

func (f Fields) Assignment_ID(ctx context.Context, assignmentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/id", assignmentID)
	return v
}

func (f Fields) Assignment_ListOfSpeakersID(ctx context.Context, assignmentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/list_of_speakers_id", assignmentID)
	return v
}

func (f Fields) Assignment_MeetingID(ctx context.Context, assignmentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/meeting_id", assignmentID)
	return v
}

func (f Fields) Assignment_NumberPollCandidates(ctx context.Context, assignmentID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/number_poll_candidates", assignmentID)
	return v
}

func (f Fields) Assignment_OpenPosts(ctx context.Context, assignmentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/open_posts", assignmentID)
	return v
}

func (f Fields) Assignment_Phase(ctx context.Context, assignmentID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/phase", assignmentID)
	return v
}

func (f Fields) Assignment_PollIDs(ctx context.Context, assignmentID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/poll_ids", assignmentID)
	return v
}

func (f Fields) Assignment_ProjectionIDs(ctx context.Context, assignmentID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/projection_ids", assignmentID)
	return v
}

func (f Fields) Assignment_TagIDs(ctx context.Context, assignmentID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/tag_ids", assignmentID)
	return v
}

func (f Fields) Assignment_Title(ctx context.Context, assignmentID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "assignment/%d/title", assignmentID)
	return v
}

func (f Fields) ChatGroup_ID(ctx context.Context, chatGroupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/id", chatGroupID)
	return v
}

func (f Fields) ChatGroup_MeetingID(ctx context.Context, chatGroupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/meeting_id", chatGroupID)
	return v
}

func (f Fields) ChatGroup_Name(ctx context.Context, chatGroupID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/name", chatGroupID)
	return v
}

func (f Fields) ChatGroup_ReadGroupIDs(ctx context.Context, chatGroupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/read_group_ids", chatGroupID)
	return v
}

func (f Fields) ChatGroup_Weight(ctx context.Context, chatGroupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/weight", chatGroupID)
	return v
}

func (f Fields) ChatGroup_WriteGroupIDs(ctx context.Context, chatGroupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "chat_group/%d/write_group_ids", chatGroupID)
	return v
}

func (f Fields) Committee_DefaultMeetingID(ctx context.Context, committeeID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/default_meeting_id", committeeID)
	return v
}

func (f Fields) Committee_Description(ctx context.Context, committeeID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/description", committeeID)
	return v
}

func (f Fields) Committee_ForwardToCommitteeIDs(ctx context.Context, committeeID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/forward_to_committee_ids", committeeID)
	return v
}

func (f Fields) Committee_ID(ctx context.Context, committeeID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/id", committeeID)
	return v
}

func (f Fields) Committee_MeetingIDs(ctx context.Context, committeeID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/meeting_ids", committeeID)
	return v
}

func (f Fields) Committee_Name(ctx context.Context, committeeID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/name", committeeID)
	return v
}

func (f Fields) Committee_OrganizationID(ctx context.Context, committeeID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/organization_id", committeeID)
	return v
}

func (f Fields) Committee_OrganizationTagIDs(ctx context.Context, committeeID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/organization_tag_ids", committeeID)
	return v
}

func (f Fields) Committee_ReceiveForwardingsFromCommitteeIDs(ctx context.Context, committeeID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/receive_forwardings_from_committee_ids", committeeID)
	return v
}

func (f Fields) Committee_TemplateMeetingID(ctx context.Context, committeeID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/template_meeting_id", committeeID)
	return v
}

func (f Fields) Committee_UserIDs(ctx context.Context, committeeID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "committee/%d/user_ids", committeeID)
	return v
}

func (f Fields) Group_AdminGroupForMeetingID(ctx context.Context, groupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/admin_group_for_meeting_id", groupID)
	return v
}

func (f Fields) Group_DefaultGroupForMeetingID(ctx context.Context, groupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/default_group_for_meeting_id", groupID)
	return v
}

func (f Fields) Group_ID(ctx context.Context, groupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/id", groupID)
	return v
}

func (f Fields) Group_MediafileAccessGroupIDs(ctx context.Context, groupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/mediafile_access_group_ids", groupID)
	return v
}

func (f Fields) Group_MediafileInheritedAccessGroupIDs(ctx context.Context, groupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/mediafile_inherited_access_group_ids", groupID)
	return v
}

func (f Fields) Group_MeetingID(ctx context.Context, groupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/meeting_id", groupID)
	return v
}

func (f Fields) Group_Name(ctx context.Context, groupID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "group/%d/name", groupID)
	return v
}

func (f Fields) Group_Permissions(ctx context.Context, groupID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "group/%d/permissions", groupID)
	return v
}

func (f Fields) Group_PollIDs(ctx context.Context, groupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/poll_ids", groupID)
	return v
}

func (f Fields) Group_ReadChatGroupIDs(ctx context.Context, groupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/read_chat_group_ids", groupID)
	return v
}

func (f Fields) Group_ReadCommentSectionIDs(ctx context.Context, groupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/read_comment_section_ids", groupID)
	return v
}

func (f Fields) Group_UsedAsAssignmentPollDefaultID(ctx context.Context, groupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/used_as_assignment_poll_default_id", groupID)
	return v
}

func (f Fields) Group_UsedAsMotionPollDefaultID(ctx context.Context, groupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/used_as_motion_poll_default_id", groupID)
	return v
}

func (f Fields) Group_UsedAsPollDefaultID(ctx context.Context, groupID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/used_as_poll_default_id", groupID)
	return v
}

func (f Fields) Group_UserIDs(ctx context.Context, groupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/user_ids", groupID)
	return v
}

func (f Fields) Group_WriteChatGroupIDs(ctx context.Context, groupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/write_chat_group_ids", groupID)
	return v
}

func (f Fields) Group_WriteCommentSectionIDs(ctx context.Context, groupID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "group/%d/write_comment_section_ids", groupID)
	return v
}

func (f Fields) ListOfSpeakers_Closed(ctx context.Context, listOfSpeakersID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/closed", listOfSpeakersID)
	return v
}

func (f Fields) ListOfSpeakers_ContentObjectID(ctx context.Context, listOfSpeakersID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/content_object_id", listOfSpeakersID)
	return v
}

func (f Fields) ListOfSpeakers_ID(ctx context.Context, listOfSpeakersID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/id", listOfSpeakersID)
	return v
}

func (f Fields) ListOfSpeakers_MeetingID(ctx context.Context, listOfSpeakersID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/meeting_id", listOfSpeakersID)
	return v
}

func (f Fields) ListOfSpeakers_ProjectionIDs(ctx context.Context, listOfSpeakersID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/projection_ids", listOfSpeakersID)
	return v
}

func (f Fields) ListOfSpeakers_SpeakerIDs(ctx context.Context, listOfSpeakersID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "list_of_speakers/%d/speaker_ids", listOfSpeakersID)
	return v
}

func (f Fields) Mediafile_AccessGroupIDs(ctx context.Context, mediafileID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/access_group_ids", mediafileID)
	return v
}

func (f Fields) Mediafile_AttachmentIDs(ctx context.Context, mediafileID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/attachment_ids", mediafileID)
	return v
}

func (f Fields) Mediafile_ChildIDs(ctx context.Context, mediafileID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/child_ids", mediafileID)
	return v
}

func (f Fields) Mediafile_CreateTimestamp(ctx context.Context, mediafileID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/create_timestamp", mediafileID)
	return v
}

func (f Fields) Mediafile_Filename(ctx context.Context, mediafileID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/filename", mediafileID)
	return v
}

func (f Fields) Mediafile_Filesize(ctx context.Context, mediafileID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/filesize", mediafileID)
	return v
}

func (f Fields) Mediafile_ID(ctx context.Context, mediafileID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/id", mediafileID)
	return v
}

func (f Fields) Mediafile_InheritedAccessGroupIDs(ctx context.Context, mediafileID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/inherited_access_group_ids", mediafileID)
	return v
}

func (f Fields) Mediafile_IsDirectory(ctx context.Context, mediafileID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/is_directory", mediafileID)
	return v
}

func (f Fields) Mediafile_IsPublic(ctx context.Context, mediafileID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/is_public", mediafileID)
	return v
}

func (f Fields) Mediafile_ListOfSpeakersID(ctx context.Context, mediafileID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/list_of_speakers_id", mediafileID)
	return v
}

func (f Fields) Mediafile_MeetingID(ctx context.Context, mediafileID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/meeting_id", mediafileID)
	return v
}

func (f Fields) Mediafile_Mimetype(ctx context.Context, mediafileID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/mimetype", mediafileID)
	return v
}

func (f Fields) Mediafile_ParentID(ctx context.Context, mediafileID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/parent_id", mediafileID)
	return v
}

func (f Fields) Mediafile_PdfInformation(ctx context.Context, mediafileID int) json.RawMessage {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/pdf_information", mediafileID)
	return v
}

func (f Fields) Mediafile_ProjectionIDs(ctx context.Context, mediafileID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/projection_ids", mediafileID)
	return v
}

func (f Fields) Mediafile_Title(ctx context.Context, mediafileID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/title", mediafileID)
	return v
}

func (f Fields) Mediafile_UsedAsFontInMeetingIDTmpl(ctx context.Context, mediafileID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/used_as_font_$_in_meeting_id", mediafileID)
	return v
}

func (f Fields) Mediafile_UsedAsFontInMeetingID(ctx context.Context, mediafileID int, replacement string) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/used_as_font_$%s_in_meeting_id", mediafileID, replacement)
	return v
}

func (f Fields) Mediafile_UsedAsLogoInMeetingIDTmpl(ctx context.Context, mediafileID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/used_as_logo_$_in_meeting_id", mediafileID)
	return v
}

func (f Fields) Mediafile_UsedAsLogoInMeetingID(ctx context.Context, mediafileID int, replacement string) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "mediafile/%d/used_as_logo_$%s_in_meeting_id", mediafileID, replacement)
	return v
}

func (f Fields) Meeting_AdminGroupID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/admin_group_id", meetingID)
	return v
}

func (f Fields) Meeting_AgendaEnableNumbering(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_enable_numbering", meetingID)
	return v
}

func (f Fields) Meeting_AgendaItemCreation(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_item_creation", meetingID)
	return v
}

func (f Fields) Meeting_AgendaItemIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_item_ids", meetingID)
	return v
}

func (f Fields) Meeting_AgendaNewItemsDefaultVisibility(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_new_items_default_visibility", meetingID)
	return v
}

func (f Fields) Meeting_AgendaNumberPrefix(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_number_prefix", meetingID)
	return v
}

func (f Fields) Meeting_AgendaNumeralSystem(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_numeral_system", meetingID)
	return v
}

func (f Fields) Meeting_AgendaShowInternalItemsOnProjector(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_show_internal_items_on_projector", meetingID)
	return v
}

func (f Fields) Meeting_AgendaShowSubtitles(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/agenda_show_subtitles", meetingID)
	return v
}

func (f Fields) Meeting_AllProjectionIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/all_projection_ids", meetingID)
	return v
}

func (f Fields) Meeting_ApplauseEnable(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_enable", meetingID)
	return v
}

func (f Fields) Meeting_ApplauseMaxAmount(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_max_amount", meetingID)
	return v
}

func (f Fields) Meeting_ApplauseMinAmount(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_min_amount", meetingID)
	return v
}

func (f Fields) Meeting_ApplauseParticleImageUrl(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_particle_image_url", meetingID)
	return v
}

func (f Fields) Meeting_ApplauseShowLevel(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_show_level", meetingID)
	return v
}

func (f Fields) Meeting_ApplauseTimeout(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_timeout", meetingID)
	return v
}

func (f Fields) Meeting_ApplauseType(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/applause_type", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentCandidateIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_candidate_ids", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_ids", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentPollAddCandidatesToListOfSpeakers(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_add_candidates_to_list_of_speakers", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentPollBallotPaperNumber(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_ballot_paper_number", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentPollBallotPaperSelection(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_ballot_paper_selection", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentPollDefault100PercentBase(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_default_100_percent_base", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentPollDefaultGroupIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_default_group_ids", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentPollDefaultMethod(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_default_method", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentPollDefaultType(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_default_type", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentPollSortPollResultByVotes(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignment_poll_sort_poll_result_by_votes", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentsExportPreamble(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignments_export_preamble", meetingID)
	return v
}

func (f Fields) Meeting_AssignmentsExportTitle(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/assignments_export_title", meetingID)
	return v
}

func (f Fields) Meeting_ChatGroupIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/chat_group_ids", meetingID)
	return v
}

func (f Fields) Meeting_CommitteeID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/committee_id", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceAutoConnect(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_auto_connect", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceAutoConnectNextSpeakers(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_auto_connect_next_speakers", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceEnableHelpdesk(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_enable_helpdesk", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceLosRestriction(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_los_restriction", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceOpenMicrophone(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_open_microphone", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceOpenVideo(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_open_video", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceShow(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_show", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceStreamPosterUrl(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_stream_poster_url", meetingID)
	return v
}

func (f Fields) Meeting_ConferenceStreamUrl(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/conference_stream_url", meetingID)
	return v
}

func (f Fields) Meeting_CustomTranslations(ctx context.Context, meetingID int) json.RawMessage {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/custom_translations", meetingID)
	return v
}

func (f Fields) Meeting_DefaultGroupID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/default_group_id", meetingID)
	return v
}

func (f Fields) Meeting_DefaultMeetingForCommitteeID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/default_meeting_for_committee_id", meetingID)
	return v
}

func (f Fields) Meeting_DefaultProjectorIDTmpl(ctx context.Context, meetingID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/default_projector_$_id", meetingID)
	return v
}

func (f Fields) Meeting_DefaultProjectorID(ctx context.Context, meetingID int, replacement string) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/default_projector_$%s_id", meetingID, replacement)
	return v
}

func (f Fields) Meeting_Description(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/description", meetingID)
	return v
}

func (f Fields) Meeting_EnableAnonymous(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/enable_anonymous", meetingID)
	return v
}

func (f Fields) Meeting_EnableChat(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/enable_chat", meetingID)
	return v
}

func (f Fields) Meeting_EndTime(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/end_time", meetingID)
	return v
}

func (f Fields) Meeting_ExportCsvEncoding(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_csv_encoding", meetingID)
	return v
}

func (f Fields) Meeting_ExportCsvSeparator(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_csv_separator", meetingID)
	return v
}

func (f Fields) Meeting_ExportPdfFontsize(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_pdf_fontsize", meetingID)
	return v
}

func (f Fields) Meeting_ExportPdfPagenumberAlignment(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_pdf_pagenumber_alignment", meetingID)
	return v
}

func (f Fields) Meeting_ExportPdfPagesize(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/export_pdf_pagesize", meetingID)
	return v
}

func (f Fields) Meeting_FontIDTmpl(ctx context.Context, meetingID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/font_$_id", meetingID)
	return v
}

func (f Fields) Meeting_FontID(ctx context.Context, meetingID int, replacement string) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/font_$%s_id", meetingID, replacement)
	return v
}

func (f Fields) Meeting_GroupIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/group_ids", meetingID)
	return v
}

func (f Fields) Meeting_ID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/id", meetingID)
	return v
}

func (f Fields) Meeting_ImportedAt(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/imported_at", meetingID)
	return v
}

func (f Fields) Meeting_JitsiDomain(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/jitsi_domain", meetingID)
	return v
}

func (f Fields) Meeting_JitsiRoomName(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/jitsi_room_name", meetingID)
	return v
}

func (f Fields) Meeting_JitsiRoomPassword(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/jitsi_room_password", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersAmountLastOnProjector(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_amount_last_on_projector", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersAmountNextOnProjector(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_amount_next_on_projector", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersCanSetContributionSelf(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_can_set_contribution_self", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersCountdownID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_countdown_id", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersCoupleCountdown(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_couple_countdown", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_enable_point_of_order_speakers", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersEnableProContraSpeech(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_enable_pro_contra_speech", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_ids", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersInitiallyClosed(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_initially_closed", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersPresentUsersOnly(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_present_users_only", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_show_amount_of_speakers_on_slide", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersShowFirstContribution(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_show_first_contribution", meetingID)
	return v
}

func (f Fields) Meeting_ListOfSpeakersSpeakerNoteForEveryone(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/list_of_speakers_speaker_note_for_everyone", meetingID)
	return v
}

func (f Fields) Meeting_Location(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/location", meetingID)
	return v
}

func (f Fields) Meeting_LogoIDTmpl(ctx context.Context, meetingID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/logo_$_id", meetingID)
	return v
}

func (f Fields) Meeting_LogoID(ctx context.Context, meetingID int, replacement string) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/logo_$%s_id", meetingID, replacement)
	return v
}

func (f Fields) Meeting_MediafileIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/mediafile_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionBlockIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_block_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionCategoryIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_category_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionChangeRecommendationIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_change_recommendation_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionCommentIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_comment_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionCommentSectionIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_comment_section_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionPollBallotPaperNumber(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_ballot_paper_number", meetingID)
	return v
}

func (f Fields) Meeting_MotionPollBallotPaperSelection(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_ballot_paper_selection", meetingID)
	return v
}

func (f Fields) Meeting_MotionPollDefault100PercentBase(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_default_100_percent_base", meetingID)
	return v
}

func (f Fields) Meeting_MotionPollDefaultGroupIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_default_group_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionPollDefaultType(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_poll_default_type", meetingID)
	return v
}

func (f Fields) Meeting_MotionStateIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_state_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionStatuteParagraphIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_statute_paragraph_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionSubmitterIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_submitter_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionWorkflowIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motion_workflow_ids", meetingID)
	return v
}

func (f Fields) Meeting_MotionsAmendmentsEnabled(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_enabled", meetingID)
	return v
}

func (f Fields) Meeting_MotionsAmendmentsInMainList(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_in_main_list", meetingID)
	return v
}

func (f Fields) Meeting_MotionsAmendmentsMultipleParagraphs(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_multiple_paragraphs", meetingID)
	return v
}

func (f Fields) Meeting_MotionsAmendmentsOfAmendments(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_of_amendments", meetingID)
	return v
}

func (f Fields) Meeting_MotionsAmendmentsPrefix(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_prefix", meetingID)
	return v
}

func (f Fields) Meeting_MotionsAmendmentsTextMode(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_amendments_text_mode", meetingID)
	return v
}

func (f Fields) Meeting_MotionsDefaultAmendmentWorkflowID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_amendment_workflow_id", meetingID)
	return v
}

func (f Fields) Meeting_MotionsDefaultLineNumbering(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_line_numbering", meetingID)
	return v
}

func (f Fields) Meeting_MotionsDefaultSorting(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_sorting", meetingID)
	return v
}

func (f Fields) Meeting_MotionsDefaultStatuteAmendmentWorkflowID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_statute_amendment_workflow_id", meetingID)
	return v
}

func (f Fields) Meeting_MotionsDefaultWorkflowID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_default_workflow_id", meetingID)
	return v
}

func (f Fields) Meeting_MotionsEnableReasonOnProjector(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_enable_reason_on_projector", meetingID)
	return v
}

func (f Fields) Meeting_MotionsEnableRecommendationOnProjector(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_enable_recommendation_on_projector", meetingID)
	return v
}

func (f Fields) Meeting_MotionsEnableSideboxOnProjector(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_enable_sidebox_on_projector", meetingID)
	return v
}

func (f Fields) Meeting_MotionsEnableTextOnProjector(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_enable_text_on_projector", meetingID)
	return v
}

func (f Fields) Meeting_MotionsExportFollowRecommendation(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_export_follow_recommendation", meetingID)
	return v
}

func (f Fields) Meeting_MotionsExportPreamble(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_export_preamble", meetingID)
	return v
}

func (f Fields) Meeting_MotionsExportSubmitterRecommendation(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_export_submitter_recommendation", meetingID)
	return v
}

func (f Fields) Meeting_MotionsExportTitle(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_export_title", meetingID)
	return v
}

func (f Fields) Meeting_MotionsLineLength(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_line_length", meetingID)
	return v
}

func (f Fields) Meeting_MotionsNumberMinDigits(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_number_min_digits", meetingID)
	return v
}

func (f Fields) Meeting_MotionsNumberType(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_number_type", meetingID)
	return v
}

func (f Fields) Meeting_MotionsNumberWithBlank(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_number_with_blank", meetingID)
	return v
}

func (f Fields) Meeting_MotionsPreamble(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_preamble", meetingID)
	return v
}

func (f Fields) Meeting_MotionsReasonRequired(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_reason_required", meetingID)
	return v
}

func (f Fields) Meeting_MotionsRecommendationTextMode(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_recommendation_text_mode", meetingID)
	return v
}

func (f Fields) Meeting_MotionsRecommendationsBy(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_recommendations_by", meetingID)
	return v
}

func (f Fields) Meeting_MotionsShowReferringMotions(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_show_referring_motions", meetingID)
	return v
}

func (f Fields) Meeting_MotionsShowSequentialNumber(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_show_sequential_number", meetingID)
	return v
}

func (f Fields) Meeting_MotionsStatuteRecommendationsBy(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_statute_recommendations_by", meetingID)
	return v
}

func (f Fields) Meeting_MotionsStatutesEnabled(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_statutes_enabled", meetingID)
	return v
}

func (f Fields) Meeting_MotionsSupportersMinAmount(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/motions_supporters_min_amount", meetingID)
	return v
}

func (f Fields) Meeting_Name(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/name", meetingID)
	return v
}

func (f Fields) Meeting_OptionIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/option_ids", meetingID)
	return v
}

func (f Fields) Meeting_OrganizationTagIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/organization_tag_ids", meetingID)
	return v
}

func (f Fields) Meeting_PersonalNoteIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/personal_note_ids", meetingID)
	return v
}

func (f Fields) Meeting_PollBallotPaperNumber(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_ballot_paper_number", meetingID)
	return v
}

func (f Fields) Meeting_PollBallotPaperSelection(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_ballot_paper_selection", meetingID)
	return v
}

func (f Fields) Meeting_PollCountdownID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_countdown_id", meetingID)
	return v
}

func (f Fields) Meeting_PollCoupleCountdown(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_couple_countdown", meetingID)
	return v
}

func (f Fields) Meeting_PollDefault100PercentBase(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_default_100_percent_base", meetingID)
	return v
}

func (f Fields) Meeting_PollDefaultGroupIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_default_group_ids", meetingID)
	return v
}

func (f Fields) Meeting_PollDefaultMethod(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_default_method", meetingID)
	return v
}

func (f Fields) Meeting_PollDefaultType(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_default_type", meetingID)
	return v
}

func (f Fields) Meeting_PollIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_ids", meetingID)
	return v
}

func (f Fields) Meeting_PollSortPollResultByVotes(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/poll_sort_poll_result_by_votes", meetingID)
	return v
}

func (f Fields) Meeting_PresentUserIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/present_user_ids", meetingID)
	return v
}

func (f Fields) Meeting_ProjectionIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projection_ids", meetingID)
	return v
}

func (f Fields) Meeting_ProjectorCountdownDefaultTime(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_countdown_default_time", meetingID)
	return v
}

func (f Fields) Meeting_ProjectorCountdownIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_countdown_ids", meetingID)
	return v
}

func (f Fields) Meeting_ProjectorCountdownWarningTime(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_countdown_warning_time", meetingID)
	return v
}

func (f Fields) Meeting_ProjectorIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_ids", meetingID)
	return v
}

func (f Fields) Meeting_ProjectorMessageIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/projector_message_ids", meetingID)
	return v
}

func (f Fields) Meeting_ReferenceProjectorID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/reference_projector_id", meetingID)
	return v
}

func (f Fields) Meeting_SpeakerIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/speaker_ids", meetingID)
	return v
}

func (f Fields) Meeting_StartTime(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/start_time", meetingID)
	return v
}

func (f Fields) Meeting_TagIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/tag_ids", meetingID)
	return v
}

func (f Fields) Meeting_TemplateForCommitteeID(ctx context.Context, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/template_for_committee_id", meetingID)
	return v
}

func (f Fields) Meeting_TopicIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/topic_ids", meetingID)
	return v
}

func (f Fields) Meeting_UrlName(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/url_name", meetingID)
	return v
}

func (f Fields) Meeting_UserIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/user_ids", meetingID)
	return v
}

func (f Fields) Meeting_UsersAllowSelfSetPresent(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_allow_self_set_present", meetingID)
	return v
}

func (f Fields) Meeting_UsersEmailBody(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_email_body", meetingID)
	return v
}

func (f Fields) Meeting_UsersEmailReplyto(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_email_replyto", meetingID)
	return v
}

func (f Fields) Meeting_UsersEmailSender(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_email_sender", meetingID)
	return v
}

func (f Fields) Meeting_UsersEmailSubject(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_email_subject", meetingID)
	return v
}

func (f Fields) Meeting_UsersEnablePresenceView(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_enable_presence_view", meetingID)
	return v
}

func (f Fields) Meeting_UsersEnableVoteWeight(ctx context.Context, meetingID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_enable_vote_weight", meetingID)
	return v
}

func (f Fields) Meeting_UsersPdfUrl(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_url", meetingID)
	return v
}

func (f Fields) Meeting_UsersPdfWelcometext(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_welcometext", meetingID)
	return v
}

func (f Fields) Meeting_UsersPdfWelcometitle(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_welcometitle", meetingID)
	return v
}

func (f Fields) Meeting_UsersPdfWlanEncryption(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_wlan_encryption", meetingID)
	return v
}

func (f Fields) Meeting_UsersPdfWlanPassword(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_wlan_password", meetingID)
	return v
}

func (f Fields) Meeting_UsersPdfWlanSsid(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_pdf_wlan_ssid", meetingID)
	return v
}

func (f Fields) Meeting_UsersSortBy(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/users_sort_by", meetingID)
	return v
}

func (f Fields) Meeting_VoteIDs(ctx context.Context, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/vote_ids", meetingID)
	return v
}

func (f Fields) Meeting_WelcomeText(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/welcome_text", meetingID)
	return v
}

func (f Fields) Meeting_WelcomeTitle(ctx context.Context, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "meeting/%d/welcome_title", meetingID)
	return v
}

func (f Fields) MotionBlock_AgendaItemID(ctx context.Context, motionBlockID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/agenda_item_id", motionBlockID)
	return v
}

func (f Fields) MotionBlock_ID(ctx context.Context, motionBlockID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/id", motionBlockID)
	return v
}

func (f Fields) MotionBlock_Internal(ctx context.Context, motionBlockID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/internal", motionBlockID)
	return v
}

func (f Fields) MotionBlock_ListOfSpeakersID(ctx context.Context, motionBlockID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/list_of_speakers_id", motionBlockID)
	return v
}

func (f Fields) MotionBlock_MeetingID(ctx context.Context, motionBlockID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/meeting_id", motionBlockID)
	return v
}

func (f Fields) MotionBlock_MotionIDs(ctx context.Context, motionBlockID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/motion_ids", motionBlockID)
	return v
}

func (f Fields) MotionBlock_ProjectionIDs(ctx context.Context, motionBlockID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/projection_ids", motionBlockID)
	return v
}

func (f Fields) MotionBlock_Title(ctx context.Context, motionBlockID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_block/%d/title", motionBlockID)
	return v
}

func (f Fields) MotionCategory_ChildIDs(ctx context.Context, motionCategoryID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/child_ids", motionCategoryID)
	return v
}

func (f Fields) MotionCategory_ID(ctx context.Context, motionCategoryID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/id", motionCategoryID)
	return v
}

func (f Fields) MotionCategory_Level(ctx context.Context, motionCategoryID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/level", motionCategoryID)
	return v
}

func (f Fields) MotionCategory_MeetingID(ctx context.Context, motionCategoryID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/meeting_id", motionCategoryID)
	return v
}

func (f Fields) MotionCategory_MotionIDs(ctx context.Context, motionCategoryID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/motion_ids", motionCategoryID)
	return v
}

func (f Fields) MotionCategory_Name(ctx context.Context, motionCategoryID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/name", motionCategoryID)
	return v
}

func (f Fields) MotionCategory_ParentID(ctx context.Context, motionCategoryID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/parent_id", motionCategoryID)
	return v
}

func (f Fields) MotionCategory_Prefix(ctx context.Context, motionCategoryID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/prefix", motionCategoryID)
	return v
}

func (f Fields) MotionCategory_Weight(ctx context.Context, motionCategoryID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_category/%d/weight", motionCategoryID)
	return v
}

func (f Fields) MotionChangeRecommendation_CreationTime(ctx context.Context, motionChangeRecommendationID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/creation_time", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_ID(ctx context.Context, motionChangeRecommendationID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/id", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_Internal(ctx context.Context, motionChangeRecommendationID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/internal", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_LineFrom(ctx context.Context, motionChangeRecommendationID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/line_from", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_LineTo(ctx context.Context, motionChangeRecommendationID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/line_to", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_MeetingID(ctx context.Context, motionChangeRecommendationID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/meeting_id", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_MotionID(ctx context.Context, motionChangeRecommendationID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/motion_id", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_OtherDescription(ctx context.Context, motionChangeRecommendationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/other_description", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_Rejected(ctx context.Context, motionChangeRecommendationID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/rejected", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_Text(ctx context.Context, motionChangeRecommendationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/text", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionChangeRecommendation_Type(ctx context.Context, motionChangeRecommendationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_change_recommendation/%d/type", motionChangeRecommendationID)
	return v
}

func (f Fields) MotionCommentSection_CommentIDs(ctx context.Context, motionCommentSectionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/comment_ids", motionCommentSectionID)
	return v
}

func (f Fields) MotionCommentSection_ID(ctx context.Context, motionCommentSectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/id", motionCommentSectionID)
	return v
}

func (f Fields) MotionCommentSection_MeetingID(ctx context.Context, motionCommentSectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/meeting_id", motionCommentSectionID)
	return v
}

func (f Fields) MotionCommentSection_Name(ctx context.Context, motionCommentSectionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/name", motionCommentSectionID)
	return v
}

func (f Fields) MotionCommentSection_ReadGroupIDs(ctx context.Context, motionCommentSectionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/read_group_ids", motionCommentSectionID)
	return v
}

func (f Fields) MotionCommentSection_Weight(ctx context.Context, motionCommentSectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/weight", motionCommentSectionID)
	return v
}

func (f Fields) MotionCommentSection_WriteGroupIDs(ctx context.Context, motionCommentSectionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment_section/%d/write_group_ids", motionCommentSectionID)
	return v
}

func (f Fields) MotionComment_Comment(ctx context.Context, motionCommentID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/comment", motionCommentID)
	return v
}

func (f Fields) MotionComment_ID(ctx context.Context, motionCommentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/id", motionCommentID)
	return v
}

func (f Fields) MotionComment_MeetingID(ctx context.Context, motionCommentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/meeting_id", motionCommentID)
	return v
}

func (f Fields) MotionComment_MotionID(ctx context.Context, motionCommentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/motion_id", motionCommentID)
	return v
}

func (f Fields) MotionComment_SectionID(ctx context.Context, motionCommentID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_comment/%d/section_id", motionCommentID)
	return v
}

func (f Fields) MotionState_AllowCreatePoll(ctx context.Context, motionStateID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/allow_create_poll", motionStateID)
	return v
}

func (f Fields) MotionState_AllowSubmitterEdit(ctx context.Context, motionStateID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/allow_submitter_edit", motionStateID)
	return v
}

func (f Fields) MotionState_AllowSupport(ctx context.Context, motionStateID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/allow_support", motionStateID)
	return v
}

func (f Fields) MotionState_CssClass(ctx context.Context, motionStateID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/css_class", motionStateID)
	return v
}

func (f Fields) MotionState_FirstStateOfWorkflowID(ctx context.Context, motionStateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/first_state_of_workflow_id", motionStateID)
	return v
}

func (f Fields) MotionState_ID(ctx context.Context, motionStateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/id", motionStateID)
	return v
}

func (f Fields) MotionState_MeetingID(ctx context.Context, motionStateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/meeting_id", motionStateID)
	return v
}

func (f Fields) MotionState_MergeAmendmentIntoFinal(ctx context.Context, motionStateID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/merge_amendment_into_final", motionStateID)
	return v
}

func (f Fields) MotionState_MotionIDs(ctx context.Context, motionStateID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/motion_ids", motionStateID)
	return v
}

func (f Fields) MotionState_MotionRecommendationIDs(ctx context.Context, motionStateID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/motion_recommendation_ids", motionStateID)
	return v
}

func (f Fields) MotionState_Name(ctx context.Context, motionStateID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/name", motionStateID)
	return v
}

func (f Fields) MotionState_NextStateIDs(ctx context.Context, motionStateID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/next_state_ids", motionStateID)
	return v
}

func (f Fields) MotionState_PreviousStateIDs(ctx context.Context, motionStateID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/previous_state_ids", motionStateID)
	return v
}

func (f Fields) MotionState_RecommendationLabel(ctx context.Context, motionStateID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/recommendation_label", motionStateID)
	return v
}

func (f Fields) MotionState_Restrictions(ctx context.Context, motionStateID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/restrictions", motionStateID)
	return v
}

func (f Fields) MotionState_SetNumber(ctx context.Context, motionStateID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/set_number", motionStateID)
	return v
}

func (f Fields) MotionState_ShowRecommendationExtensionField(ctx context.Context, motionStateID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/show_recommendation_extension_field", motionStateID)
	return v
}

func (f Fields) MotionState_ShowStateExtensionField(ctx context.Context, motionStateID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/show_state_extension_field", motionStateID)
	return v
}

func (f Fields) MotionState_WorkflowID(ctx context.Context, motionStateID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_state/%d/workflow_id", motionStateID)
	return v
}

func (f Fields) MotionStatuteParagraph_ID(ctx context.Context, motionStatuteParagraphID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/id", motionStatuteParagraphID)
	return v
}

func (f Fields) MotionStatuteParagraph_MeetingID(ctx context.Context, motionStatuteParagraphID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/meeting_id", motionStatuteParagraphID)
	return v
}

func (f Fields) MotionStatuteParagraph_MotionIDs(ctx context.Context, motionStatuteParagraphID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/motion_ids", motionStatuteParagraphID)
	return v
}

func (f Fields) MotionStatuteParagraph_Text(ctx context.Context, motionStatuteParagraphID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/text", motionStatuteParagraphID)
	return v
}

func (f Fields) MotionStatuteParagraph_Title(ctx context.Context, motionStatuteParagraphID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/title", motionStatuteParagraphID)
	return v
}

func (f Fields) MotionStatuteParagraph_Weight(ctx context.Context, motionStatuteParagraphID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_statute_paragraph/%d/weight", motionStatuteParagraphID)
	return v
}

func (f Fields) MotionSubmitter_ID(ctx context.Context, motionSubmitterID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/id", motionSubmitterID)
	return v
}

func (f Fields) MotionSubmitter_MeetingID(ctx context.Context, motionSubmitterID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/meeting_id", motionSubmitterID)
	return v
}

func (f Fields) MotionSubmitter_MotionID(ctx context.Context, motionSubmitterID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/motion_id", motionSubmitterID)
	return v
}

func (f Fields) MotionSubmitter_UserID(ctx context.Context, motionSubmitterID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/user_id", motionSubmitterID)
	return v
}

func (f Fields) MotionSubmitter_Weight(ctx context.Context, motionSubmitterID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_submitter/%d/weight", motionSubmitterID)
	return v
}

func (f Fields) MotionWorkflow_DefaultAmendmentWorkflowMeetingID(ctx context.Context, motionWorkflowID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/default_amendment_workflow_meeting_id", motionWorkflowID)
	return v
}

func (f Fields) MotionWorkflow_DefaultStatuteAmendmentWorkflowMeetingID(ctx context.Context, motionWorkflowID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/default_statute_amendment_workflow_meeting_id", motionWorkflowID)
	return v
}

func (f Fields) MotionWorkflow_DefaultWorkflowMeetingID(ctx context.Context, motionWorkflowID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/default_workflow_meeting_id", motionWorkflowID)
	return v
}

func (f Fields) MotionWorkflow_FirstStateID(ctx context.Context, motionWorkflowID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/first_state_id", motionWorkflowID)
	return v
}

func (f Fields) MotionWorkflow_ID(ctx context.Context, motionWorkflowID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/id", motionWorkflowID)
	return v
}

func (f Fields) MotionWorkflow_MeetingID(ctx context.Context, motionWorkflowID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/meeting_id", motionWorkflowID)
	return v
}

func (f Fields) MotionWorkflow_Name(ctx context.Context, motionWorkflowID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/name", motionWorkflowID)
	return v
}

func (f Fields) MotionWorkflow_StateIDs(ctx context.Context, motionWorkflowID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion_workflow/%d/state_ids", motionWorkflowID)
	return v
}

func (f Fields) Motion_AgendaItemID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/agenda_item_id", motionID)
	return v
}

func (f Fields) Motion_AllDerivedMotionIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/all_derived_motion_ids", motionID)
	return v
}

func (f Fields) Motion_AllOriginIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/all_origin_ids", motionID)
	return v
}

func (f Fields) Motion_AmendmentIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/amendment_ids", motionID)
	return v
}

func (f Fields) Motion_AmendmentParagraphTmpl(ctx context.Context, motionID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/amendment_paragraph_$", motionID)
	return v
}

func (f Fields) Motion_AmendmentParagraph(ctx context.Context, motionID int, replacement string) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/amendment_paragraph_$%s", motionID, replacement)
	return v
}

func (f Fields) Motion_AttachmentIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/attachment_ids", motionID)
	return v
}

func (f Fields) Motion_BlockID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/block_id", motionID)
	return v
}

func (f Fields) Motion_CategoryID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/category_id", motionID)
	return v
}

func (f Fields) Motion_CategoryWeight(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/category_weight", motionID)
	return v
}

func (f Fields) Motion_ChangeRecommendationIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/change_recommendation_ids", motionID)
	return v
}

func (f Fields) Motion_CommentIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/comment_ids", motionID)
	return v
}

func (f Fields) Motion_Created(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/created", motionID)
	return v
}

func (f Fields) Motion_DerivedMotionIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/derived_motion_ids", motionID)
	return v
}

func (f Fields) Motion_ID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/id", motionID)
	return v
}

func (f Fields) Motion_LastModified(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/last_modified", motionID)
	return v
}

func (f Fields) Motion_LeadMotionID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/lead_motion_id", motionID)
	return v
}

func (f Fields) Motion_ListOfSpeakersID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/list_of_speakers_id", motionID)
	return v
}

func (f Fields) Motion_MeetingID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/meeting_id", motionID)
	return v
}

func (f Fields) Motion_ModifiedFinalVersion(ctx context.Context, motionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/modified_final_version", motionID)
	return v
}

func (f Fields) Motion_Number(ctx context.Context, motionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/number", motionID)
	return v
}

func (f Fields) Motion_NumberValue(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/number_value", motionID)
	return v
}

func (f Fields) Motion_OptionIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/option_ids", motionID)
	return v
}

func (f Fields) Motion_OriginID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/origin_id", motionID)
	return v
}

func (f Fields) Motion_PersonalNoteIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/personal_note_ids", motionID)
	return v
}

func (f Fields) Motion_PollIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/poll_ids", motionID)
	return v
}

func (f Fields) Motion_ProjectionIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/projection_ids", motionID)
	return v
}

func (f Fields) Motion_Reason(ctx context.Context, motionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/reason", motionID)
	return v
}

func (f Fields) Motion_RecommendationExtension(ctx context.Context, motionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/recommendation_extension", motionID)
	return v
}

func (f Fields) Motion_RecommendationExtensionReferenceIDs(ctx context.Context, motionID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/recommendation_extension_reference_ids", motionID)
	return v
}

func (f Fields) Motion_RecommendationID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/recommendation_id", motionID)
	return v
}

func (f Fields) Motion_ReferencedInMotionRecommendationExtensionIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/referenced_in_motion_recommendation_extension_ids", motionID)
	return v
}

func (f Fields) Motion_SequentialNumber(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/sequential_number", motionID)
	return v
}

func (f Fields) Motion_SortChildIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/sort_child_ids", motionID)
	return v
}

func (f Fields) Motion_SortParentID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/sort_parent_id", motionID)
	return v
}

func (f Fields) Motion_SortWeight(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/sort_weight", motionID)
	return v
}

func (f Fields) Motion_StateExtension(ctx context.Context, motionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/state_extension", motionID)
	return v
}

func (f Fields) Motion_StateID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/state_id", motionID)
	return v
}

func (f Fields) Motion_StatuteParagraphID(ctx context.Context, motionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/statute_paragraph_id", motionID)
	return v
}

func (f Fields) Motion_SubmitterIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/submitter_ids", motionID)
	return v
}

func (f Fields) Motion_SupporterIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/supporter_ids", motionID)
	return v
}

func (f Fields) Motion_TagIDs(ctx context.Context, motionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/tag_ids", motionID)
	return v
}

func (f Fields) Motion_Text(ctx context.Context, motionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/text", motionID)
	return v
}

func (f Fields) Motion_Title(ctx context.Context, motionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "motion/%d/title", motionID)
	return v
}

func (f Fields) Option_Abstain(ctx context.Context, optionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/abstain", optionID)
	return v
}

func (f Fields) Option_ContentObjectID(ctx context.Context, optionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "option/%d/content_object_id", optionID)
	return v
}

func (f Fields) Option_ID(ctx context.Context, optionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/id", optionID)
	return v
}

func (f Fields) Option_MeetingID(ctx context.Context, optionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/meeting_id", optionID)
	return v
}

func (f Fields) Option_No(ctx context.Context, optionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/no", optionID)
	return v
}

func (f Fields) Option_PollID(ctx context.Context, optionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/poll_id", optionID)
	return v
}

func (f Fields) Option_Text(ctx context.Context, optionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "option/%d/text", optionID)
	return v
}

func (f Fields) Option_UsedAsGlobalOptionInPollID(ctx context.Context, optionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/used_as_global_option_in_poll_id", optionID)
	return v
}

func (f Fields) Option_VoteIDs(ctx context.Context, optionID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/vote_ids", optionID)
	return v
}

func (f Fields) Option_Weight(ctx context.Context, optionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/weight", optionID)
	return v
}

func (f Fields) Option_Yes(ctx context.Context, optionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "option/%d/yes", optionID)
	return v
}

func (f Fields) OrganizationTag_Color(ctx context.Context, organizationTagID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/color", organizationTagID)
	return v
}

func (f Fields) OrganizationTag_ID(ctx context.Context, organizationTagID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/id", organizationTagID)
	return v
}

func (f Fields) OrganizationTag_Name(ctx context.Context, organizationTagID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/name", organizationTagID)
	return v
}

func (f Fields) OrganizationTag_OrganizationID(ctx context.Context, organizationTagID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/organization_id", organizationTagID)
	return v
}

func (f Fields) OrganizationTag_TaggedIDs(ctx context.Context, organizationTagID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "organization_tag/%d/tagged_ids", organizationTagID)
	return v
}

func (f Fields) Organization_CommitteeIDs(ctx context.Context, organizationID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/committee_ids", organizationID)
	return v
}

func (f Fields) Organization_Description(ctx context.Context, organizationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/description", organizationID)
	return v
}

func (f Fields) Organization_EnableElectronicVoting(ctx context.Context, organizationID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/enable_electronic_voting", organizationID)
	return v
}

func (f Fields) Organization_ID(ctx context.Context, organizationID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/id", organizationID)
	return v
}

func (f Fields) Organization_LegalNotice(ctx context.Context, organizationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/legal_notice", organizationID)
	return v
}

func (f Fields) Organization_LoginText(ctx context.Context, organizationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/login_text", organizationID)
	return v
}

func (f Fields) Organization_Name(ctx context.Context, organizationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/name", organizationID)
	return v
}

func (f Fields) Organization_OrganizationTagIDs(ctx context.Context, organizationID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/organization_tag_ids", organizationID)
	return v
}

func (f Fields) Organization_PrivacyPolicy(ctx context.Context, organizationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/privacy_policy", organizationID)
	return v
}

func (f Fields) Organization_ResetPasswordVerboseErrors(ctx context.Context, organizationID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/reset_password_verbose_errors", organizationID)
	return v
}

func (f Fields) Organization_ResourceIDs(ctx context.Context, organizationID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/resource_ids", organizationID)
	return v
}

func (f Fields) Organization_Theme(ctx context.Context, organizationID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "organization/%d/theme", organizationID)
	return v
}

func (f Fields) PersonalNote_ContentObjectID(ctx context.Context, personalNoteID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/content_object_id", personalNoteID)
	return v
}

func (f Fields) PersonalNote_ID(ctx context.Context, personalNoteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/id", personalNoteID)
	return v
}

func (f Fields) PersonalNote_MeetingID(ctx context.Context, personalNoteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/meeting_id", personalNoteID)
	return v
}

func (f Fields) PersonalNote_Note(ctx context.Context, personalNoteID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/note", personalNoteID)
	return v
}

func (f Fields) PersonalNote_Star(ctx context.Context, personalNoteID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/star", personalNoteID)
	return v
}

func (f Fields) PersonalNote_UserID(ctx context.Context, personalNoteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "personal_note/%d/user_id", personalNoteID)
	return v
}

func (f Fields) Poll_Backend(ctx context.Context, pollID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/backend", pollID)
	return v
}

func (f Fields) Poll_ContentObjectID(ctx context.Context, pollID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/content_object_id", pollID)
	return v
}

func (f Fields) Poll_Description(ctx context.Context, pollID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/description", pollID)
	return v
}

func (f Fields) Poll_EntitledGroupIDs(ctx context.Context, pollID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/entitled_group_ids", pollID)
	return v
}

func (f Fields) Poll_EntitledUsersAtStop(ctx context.Context, pollID int) json.RawMessage {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/entitled_users_at_stop", pollID)
	return v
}

func (f Fields) Poll_GlobalAbstain(ctx context.Context, pollID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/global_abstain", pollID)
	return v
}

func (f Fields) Poll_GlobalNo(ctx context.Context, pollID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/global_no", pollID)
	return v
}

func (f Fields) Poll_GlobalOptionID(ctx context.Context, pollID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/global_option_id", pollID)
	return v
}

func (f Fields) Poll_GlobalYes(ctx context.Context, pollID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/global_yes", pollID)
	return v
}

func (f Fields) Poll_ID(ctx context.Context, pollID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/id", pollID)
	return v
}

func (f Fields) Poll_IsPseudoanonymized(ctx context.Context, pollID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/is_pseudoanonymized", pollID)
	return v
}

func (f Fields) Poll_MaxVotesAmount(ctx context.Context, pollID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/max_votes_amount", pollID)
	return v
}

func (f Fields) Poll_MeetingID(ctx context.Context, pollID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/meeting_id", pollID)
	return v
}

func (f Fields) Poll_MinVotesAmount(ctx context.Context, pollID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/min_votes_amount", pollID)
	return v
}

func (f Fields) Poll_OnehundredPercentBase(ctx context.Context, pollID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/onehundred_percent_base", pollID)
	return v
}

func (f Fields) Poll_OptionIDs(ctx context.Context, pollID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/option_ids", pollID)
	return v
}

func (f Fields) Poll_Pollmethod(ctx context.Context, pollID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/pollmethod", pollID)
	return v
}

func (f Fields) Poll_ProjectionIDs(ctx context.Context, pollID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/projection_ids", pollID)
	return v
}

func (f Fields) Poll_State(ctx context.Context, pollID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/state", pollID)
	return v
}

func (f Fields) Poll_Title(ctx context.Context, pollID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/title", pollID)
	return v
}

func (f Fields) Poll_Type(ctx context.Context, pollID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/type", pollID)
	return v
}

func (f Fields) Poll_VotedIDs(ctx context.Context, pollID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/voted_ids", pollID)
	return v
}

func (f Fields) Poll_Votescast(ctx context.Context, pollID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/votescast", pollID)
	return v
}

func (f Fields) Poll_Votesinvalid(ctx context.Context, pollID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/votesinvalid", pollID)
	return v
}

func (f Fields) Poll_Votesvalid(ctx context.Context, pollID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "poll/%d/votesvalid", pollID)
	return v
}

func (f Fields) Projection_Content(ctx context.Context, projectionID int) json.RawMessage {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/content", projectionID)
	return v
}

func (f Fields) Projection_ContentObjectID(ctx context.Context, projectionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/content_object_id", projectionID)
	return v
}

func (f Fields) Projection_CurrentProjectorID(ctx context.Context, projectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/current_projector_id", projectionID)
	return v
}

func (f Fields) Projection_HistoryProjectorID(ctx context.Context, projectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/history_projector_id", projectionID)
	return v
}

func (f Fields) Projection_ID(ctx context.Context, projectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/id", projectionID)
	return v
}

func (f Fields) Projection_MeetingID(ctx context.Context, projectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/meeting_id", projectionID)
	return v
}

func (f Fields) Projection_Options(ctx context.Context, projectionID int) json.RawMessage {
	var v json.RawMessage
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/options", projectionID)
	return v
}

func (f Fields) Projection_PreviewProjectorID(ctx context.Context, projectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/preview_projector_id", projectionID)
	return v
}

func (f Fields) Projection_Stable(ctx context.Context, projectionID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/stable", projectionID)
	return v
}

func (f Fields) Projection_Type(ctx context.Context, projectionID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/type", projectionID)
	return v
}

func (f Fields) Projection_Weight(ctx context.Context, projectionID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projection/%d/weight", projectionID)
	return v
}

func (f Fields) ProjectorCountdown_CountdownTime(ctx context.Context, projectorCountdownID int) float32 {
	var v float32
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/countdown_time", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_DefaultTime(ctx context.Context, projectorCountdownID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/default_time", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_Description(ctx context.Context, projectorCountdownID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/description", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_ID(ctx context.Context, projectorCountdownID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/id", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_MeetingID(ctx context.Context, projectorCountdownID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/meeting_id", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_ProjectionIDs(ctx context.Context, projectorCountdownID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/projection_ids", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_Running(ctx context.Context, projectorCountdownID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/running", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_Title(ctx context.Context, projectorCountdownID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/title", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_UsedAsListOfSpeakerCountdownMeetingID(ctx context.Context, projectorCountdownID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/used_as_list_of_speaker_countdown_meeting_id", projectorCountdownID)
	return v
}

func (f Fields) ProjectorCountdown_UsedAsPollCountdownMeetingID(ctx context.Context, projectorCountdownID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_countdown/%d/used_as_poll_countdown_meeting_id", projectorCountdownID)
	return v
}

func (f Fields) ProjectorMessage_ID(ctx context.Context, projectorMessageID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_message/%d/id", projectorMessageID)
	return v
}

func (f Fields) ProjectorMessage_MeetingID(ctx context.Context, projectorMessageID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector_message/%d/meeting_id", projectorMessageID)
	return v
}

func (f Fields) ProjectorMessage_Message(ctx context.Context, projectorMessageID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector_message/%d/message", projectorMessageID)
	return v
}

func (f Fields) ProjectorMessage_ProjectionIDs(ctx context.Context, projectorMessageID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector_message/%d/projection_ids", projectorMessageID)
	return v
}

func (f Fields) Projector_AspectRatioDenominator(ctx context.Context, projectorID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/aspect_ratio_denominator", projectorID)
	return v
}

func (f Fields) Projector_AspectRatioNumerator(ctx context.Context, projectorID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/aspect_ratio_numerator", projectorID)
	return v
}

func (f Fields) Projector_BackgroundColor(ctx context.Context, projectorID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/background_color", projectorID)
	return v
}

func (f Fields) Projector_ChyronBackgroundColor(ctx context.Context, projectorID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/chyron_background_color", projectorID)
	return v
}

func (f Fields) Projector_ChyronFontColor(ctx context.Context, projectorID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/chyron_font_color", projectorID)
	return v
}

func (f Fields) Projector_Color(ctx context.Context, projectorID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/color", projectorID)
	return v
}

func (f Fields) Projector_CurrentProjectionIDs(ctx context.Context, projectorID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/current_projection_ids", projectorID)
	return v
}

func (f Fields) Projector_HeaderBackgroundColor(ctx context.Context, projectorID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/header_background_color", projectorID)
	return v
}

func (f Fields) Projector_HeaderFontColor(ctx context.Context, projectorID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/header_font_color", projectorID)
	return v
}

func (f Fields) Projector_HeaderH1Color(ctx context.Context, projectorID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/header_h1_color", projectorID)
	return v
}

func (f Fields) Projector_HistoryProjectionIDs(ctx context.Context, projectorID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/history_projection_ids", projectorID)
	return v
}

func (f Fields) Projector_ID(ctx context.Context, projectorID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/id", projectorID)
	return v
}

func (f Fields) Projector_MeetingID(ctx context.Context, projectorID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/meeting_id", projectorID)
	return v
}

func (f Fields) Projector_Name(ctx context.Context, projectorID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/name", projectorID)
	return v
}

func (f Fields) Projector_PreviewProjectionIDs(ctx context.Context, projectorID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/preview_projection_ids", projectorID)
	return v
}

func (f Fields) Projector_Scale(ctx context.Context, projectorID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/scale", projectorID)
	return v
}

func (f Fields) Projector_Scroll(ctx context.Context, projectorID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/scroll", projectorID)
	return v
}

func (f Fields) Projector_ShowClock(ctx context.Context, projectorID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/show_clock", projectorID)
	return v
}

func (f Fields) Projector_ShowHeaderFooter(ctx context.Context, projectorID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/show_header_footer", projectorID)
	return v
}

func (f Fields) Projector_ShowLogo(ctx context.Context, projectorID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/show_logo", projectorID)
	return v
}

func (f Fields) Projector_ShowTitle(ctx context.Context, projectorID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/show_title", projectorID)
	return v
}

func (f Fields) Projector_UsedAsDefaultInMeetingIDTmpl(ctx context.Context, projectorID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/used_as_default_$_in_meeting_id", projectorID)
	return v
}

func (f Fields) Projector_UsedAsDefaultInMeetingID(ctx context.Context, projectorID int, replacement string) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/used_as_default_$%s_in_meeting_id", projectorID, replacement)
	return v
}

func (f Fields) Projector_UsedAsReferenceProjectorMeetingID(ctx context.Context, projectorID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/used_as_reference_projector_meeting_id", projectorID)
	return v
}

func (f Fields) Projector_Width(ctx context.Context, projectorID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "projector/%d/width", projectorID)
	return v
}

func (f Fields) Resource_Filesize(ctx context.Context, resourceID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/filesize", resourceID)
	return v
}

func (f Fields) Resource_ID(ctx context.Context, resourceID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/id", resourceID)
	return v
}

func (f Fields) Resource_Mimetype(ctx context.Context, resourceID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/mimetype", resourceID)
	return v
}

func (f Fields) Resource_OrganizationID(ctx context.Context, resourceID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/organization_id", resourceID)
	return v
}

func (f Fields) Resource_Token(ctx context.Context, resourceID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "resource/%d/token", resourceID)
	return v
}

func (f Fields) Speaker_BeginTime(ctx context.Context, speakerID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/begin_time", speakerID)
	return v
}

func (f Fields) Speaker_EndTime(ctx context.Context, speakerID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/end_time", speakerID)
	return v
}

func (f Fields) Speaker_ID(ctx context.Context, speakerID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/id", speakerID)
	return v
}

func (f Fields) Speaker_ListOfSpeakersID(ctx context.Context, speakerID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/list_of_speakers_id", speakerID)
	return v
}

func (f Fields) Speaker_MeetingID(ctx context.Context, speakerID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/meeting_id", speakerID)
	return v
}

func (f Fields) Speaker_Note(ctx context.Context, speakerID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/note", speakerID)
	return v
}

func (f Fields) Speaker_PointOfOrder(ctx context.Context, speakerID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/point_of_order", speakerID)
	return v
}

func (f Fields) Speaker_SpeechState(ctx context.Context, speakerID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/speech_state", speakerID)
	return v
}

func (f Fields) Speaker_UserID(ctx context.Context, speakerID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/user_id", speakerID)
	return v
}

func (f Fields) Speaker_Weight(ctx context.Context, speakerID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "speaker/%d/weight", speakerID)
	return v
}

func (f Fields) Tag_ID(ctx context.Context, tagID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "tag/%d/id", tagID)
	return v
}

func (f Fields) Tag_MeetingID(ctx context.Context, tagID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "tag/%d/meeting_id", tagID)
	return v
}

func (f Fields) Tag_Name(ctx context.Context, tagID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "tag/%d/name", tagID)
	return v
}

func (f Fields) Tag_TaggedIDs(ctx context.Context, tagID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "tag/%d/tagged_ids", tagID)
	return v
}

func (f Fields) Topic_AgendaItemID(ctx context.Context, topicID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/agenda_item_id", topicID)
	return v
}

func (f Fields) Topic_AttachmentIDs(ctx context.Context, topicID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/attachment_ids", topicID)
	return v
}

func (f Fields) Topic_ID(ctx context.Context, topicID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/id", topicID)
	return v
}

func (f Fields) Topic_ListOfSpeakersID(ctx context.Context, topicID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/list_of_speakers_id", topicID)
	return v
}

func (f Fields) Topic_MeetingID(ctx context.Context, topicID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/meeting_id", topicID)
	return v
}

func (f Fields) Topic_OptionIDs(ctx context.Context, topicID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/option_ids", topicID)
	return v
}

func (f Fields) Topic_ProjectionIDs(ctx context.Context, topicID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/projection_ids", topicID)
	return v
}

func (f Fields) Topic_TagIDs(ctx context.Context, topicID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/tag_ids", topicID)
	return v
}

func (f Fields) Topic_Text(ctx context.Context, topicID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/text", topicID)
	return v
}

func (f Fields) Topic_Title(ctx context.Context, topicID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "topic/%d/title", topicID)
	return v
}

func (f Fields) User_AboutMeTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/about_me_$", userID)
	return v
}

func (f Fields) User_AboutMe(ctx context.Context, userID int, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/about_me_$%d", userID, meetingID)
	return v
}

func (f Fields) User_AssignmentCandidateIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/assignment_candidate_$_ids", userID)
	return v
}

func (f Fields) User_AssignmentCandidateIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/assignment_candidate_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_CanChangeOwnPassword(ctx context.Context, userID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "user/%d/can_change_own_password", userID)
	return v
}

func (f Fields) User_CommentTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/comment_$", userID)
	return v
}

func (f Fields) User_Comment(ctx context.Context, userID int, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/comment_$%d", userID, meetingID)
	return v
}

func (f Fields) User_CommitteeIDs(ctx context.Context, userID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/committee_ids", userID)
	return v
}

func (f Fields) User_CommitteeManagementLevelTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/committee_$_management_level", userID)
	return v
}

func (f Fields) User_CommitteeManagementLevel(ctx context.Context, userID int, committeeID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/committee_$%d_management_level", userID, committeeID)
	return v
}

func (f Fields) User_DefaultNumber(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/default_number", userID)
	return v
}

func (f Fields) User_DefaultPassword(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/default_password", userID)
	return v
}

func (f Fields) User_DefaultStructureLevel(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/default_structure_level", userID)
	return v
}

func (f Fields) User_DefaultVoteWeight(ctx context.Context, userID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/default_vote_weight", userID)
	return v
}

func (f Fields) User_Email(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/email", userID)
	return v
}

func (f Fields) User_FirstName(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/first_name", userID)
	return v
}

func (f Fields) User_Gender(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/gender", userID)
	return v
}

func (f Fields) User_GroupIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/group_$_ids", userID)
	return v
}

func (f Fields) User_GroupIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/group_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_ID(ctx context.Context, userID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/id", userID)
	return v
}

func (f Fields) User_IsActive(ctx context.Context, userID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "user/%d/is_active", userID)
	return v
}

func (f Fields) User_IsDemoUser(ctx context.Context, userID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "user/%d/is_demo_user", userID)
	return v
}

func (f Fields) User_IsPhysicalPerson(ctx context.Context, userID int) bool {
	var v bool
	f.fetch.FetchIfExist(ctx, &v, "user/%d/is_physical_person", userID)
	return v
}

func (f Fields) User_IsPresentInMeetingIDs(ctx context.Context, userID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/is_present_in_meeting_ids", userID)
	return v
}

func (f Fields) User_LastEmailSend(ctx context.Context, userID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/last_email_send", userID)
	return v
}

func (f Fields) User_LastName(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/last_name", userID)
	return v
}

func (f Fields) User_MeetingIDs(ctx context.Context, userID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/meeting_ids", userID)
	return v
}

func (f Fields) User_NumberTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/number_$", userID)
	return v
}

func (f Fields) User_Number(ctx context.Context, userID int, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/number_$%d", userID, meetingID)
	return v
}

func (f Fields) User_OptionIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/option_$_ids", userID)
	return v
}

func (f Fields) User_OptionIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/option_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_OrganizationManagementLevel(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/organization_management_level", userID)
	return v
}

func (f Fields) User_Password(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/password", userID)
	return v
}

func (f Fields) User_PersonalNoteIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/personal_note_$_ids", userID)
	return v
}

func (f Fields) User_PersonalNoteIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/personal_note_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_PollVotedIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/poll_voted_$_ids", userID)
	return v
}

func (f Fields) User_PollVotedIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/poll_voted_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_ProjectionIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/projection_$_ids", userID)
	return v
}

func (f Fields) User_ProjectionIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/projection_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_SpeakerIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/speaker_$_ids", userID)
	return v
}

func (f Fields) User_SpeakerIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/speaker_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_StructureLevelTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/structure_level_$", userID)
	return v
}

func (f Fields) User_StructureLevel(ctx context.Context, userID int, meetingID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/structure_level_$%d", userID, meetingID)
	return v
}

func (f Fields) User_SubmittedMotionIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/submitted_motion_$_ids", userID)
	return v
}

func (f Fields) User_SubmittedMotionIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/submitted_motion_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_SupportedMotionIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/supported_motion_$_ids", userID)
	return v
}

func (f Fields) User_SupportedMotionIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/supported_motion_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_Title(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/title", userID)
	return v
}

func (f Fields) User_Username(ctx context.Context, userID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/username", userID)
	return v
}

func (f Fields) User_VoteDelegatedToIDTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegated_$_to_id", userID)
	return v
}

func (f Fields) User_VoteDelegatedToID(ctx context.Context, userID int, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegated_$%d_to_id", userID, meetingID)
	return v
}

func (f Fields) User_VoteDelegatedVoteIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegated_vote_$_ids", userID)
	return v
}

func (f Fields) User_VoteDelegatedVoteIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegated_vote_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_VoteDelegationsFromIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegations_$_from_ids", userID)
	return v
}

func (f Fields) User_VoteDelegationsFromIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_delegations_$%d_from_ids", userID, meetingID)
	return v
}

func (f Fields) User_VoteIDsTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_$_ids", userID)
	return v
}

func (f Fields) User_VoteIDs(ctx context.Context, userID int, meetingID int) []int {
	var v []int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_$%d_ids", userID, meetingID)
	return v
}

func (f Fields) User_VoteWeightTmpl(ctx context.Context, userID int) []string {
	var v []string
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_weight_$", userID)
	return v
}

func (f Fields) User_VoteWeight(ctx context.Context, userID int, meetingID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "user/%d/vote_weight_$%d", userID, meetingID)
	return v
}

func (f Fields) Vote_DelegatedUserID(ctx context.Context, voteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/delegated_user_id", voteID)
	return v
}

func (f Fields) Vote_ID(ctx context.Context, voteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/id", voteID)
	return v
}

func (f Fields) Vote_MeetingID(ctx context.Context, voteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/meeting_id", voteID)
	return v
}

func (f Fields) Vote_OptionID(ctx context.Context, voteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/option_id", voteID)
	return v
}

func (f Fields) Vote_UserID(ctx context.Context, voteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/user_id", voteID)
	return v
}

func (f Fields) Vote_UserToken(ctx context.Context, voteID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/user_token", voteID)
	return v
}

func (f Fields) Vote_Value(ctx context.Context, voteID int) string {
	var v string
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/value", voteID)
	return v
}

func (f Fields) Vote_Weight(ctx context.Context, voteID int) int {
	var v int
	f.fetch.FetchIfExist(ctx, &v, "vote/%d/weight", voteID)
	return v
}
