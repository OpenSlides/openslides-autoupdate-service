// Code generated from models.yml DO NOT EDIT.
package perm

const (
	AgendaItemCanManage               TPermission = "agenda_item.can_manage"
	AgendaItemCanManageModeratorNotes TPermission = "agenda_item.can_manage_moderator_notes"
	AgendaItemCanSee                  TPermission = "agenda_item.can_see"
	AgendaItemCanSeeInternal          TPermission = "agenda_item.can_see_internal"
	AgendaItemCanSeeModeratorNotes    TPermission = "agenda_item.can_see_moderator_notes"
	AssignmentCanManage               TPermission = "assignment.can_manage"
	AssignmentCanNominateOther        TPermission = "assignment.can_nominate_other"
	AssignmentCanNominateSelf         TPermission = "assignment.can_nominate_self"
	AssignmentCanSee                  TPermission = "assignment.can_see"
	ChatCanManage                     TPermission = "chat.can_manage"
	ListOfSpeakersCanBeSpeaker        TPermission = "list_of_speakers.can_be_speaker"
	ListOfSpeakersCanManage           TPermission = "list_of_speakers.can_manage"
	ListOfSpeakersCanSee              TPermission = "list_of_speakers.can_see"
	MediafileCanManage                TPermission = "mediafile.can_manage"
	MediafileCanSee                   TPermission = "mediafile.can_see"
	MeetingCanManageLogosAndFonts     TPermission = "meeting.can_manage_logos_and_fonts"
	MeetingCanManageSettings          TPermission = "meeting.can_manage_settings"
	MeetingCanSeeAutopilot            TPermission = "meeting.can_see_autopilot"
	MeetingCanSeeFrontpage            TPermission = "meeting.can_see_frontpage"
	MeetingCanSeeHistory              TPermission = "meeting.can_see_history"
	MeetingCanSeeLivestream           TPermission = "meeting.can_see_livestream"
	MotionCanCreate                   TPermission = "motion.can_create"
	MotionCanCreateAmendments         TPermission = "motion.can_create_amendments"
	MotionCanForward                  TPermission = "motion.can_forward"
	MotionCanManage                   TPermission = "motion.can_manage"
	MotionCanManageMetadata           TPermission = "motion.can_manage_metadata"
	MotionCanManagePolls              TPermission = "motion.can_manage_polls"
	MotionCanSee                      TPermission = "motion.can_see"
	MotionCanSeeInternal              TPermission = "motion.can_see_internal"
	MotionCanSupport                  TPermission = "motion.can_support"
	PollCanManage                     TPermission = "poll.can_manage"
	ProjectorCanManage                TPermission = "projector.can_manage"
	ProjectorCanSee                   TPermission = "projector.can_see"
	TagCanManage                      TPermission = "tag.can_manage"
	UserCanManage                     TPermission = "user.can_manage"
	UserCanManagePresence             TPermission = "user.can_manage_presence"
	UserCanSee                        TPermission = "user.can_see"
	UserCanSeeSensitiveData           TPermission = "user.can_see_sensitive_data"
	UserCanUpdate                     TPermission = "user.can_update"
)

var derivatePerms = map[TPermission][]TPermission{
	"agenda_item.can_manage":                 {"agenda_item.can_see", "agenda_item.can_see_internal"},
	"agenda_item.can_manage_moderator_notes": {"agenda_item.can_see", "agenda_item.can_see_moderator_notes"},
	"agenda_item.can_see":                    {},
	"agenda_item.can_see_internal":           {"agenda_item.can_see"},
	"agenda_item.can_see_moderator_notes":    {"agenda_item.can_see"},
	"assignment.can_manage":                  {"assignment.can_nominate_other", "assignment.can_see"},
	"assignment.can_nominate_other":          {"assignment.can_see"},
	"assignment.can_nominate_self":           {"assignment.can_see"},
	"assignment.can_see":                     {},
	"chat.can_manage":                        {},
	"list_of_speakers.can_be_speaker":        {},
	"list_of_speakers.can_manage":            {"list_of_speakers.can_see"},
	"list_of_speakers.can_see":               {},
	"mediafile.can_manage":                   {"mediafile.can_see"},
	"mediafile.can_see":                      {},
	"meeting.can_manage_logos_and_fonts":     {},
	"meeting.can_manage_settings":            {},
	"meeting.can_see_autopilot":              {},
	"meeting.can_see_frontpage":              {},
	"meeting.can_see_history":                {},
	"meeting.can_see_livestream":             {},
	"motion.can_create":                      {"motion.can_see"},
	"motion.can_create_amendments":           {"motion.can_see"},
	"motion.can_forward":                     {"motion.can_see"},
	"motion.can_manage":                      {"motion.can_create", "motion.can_create_amendments", "motion.can_forward", "motion.can_manage_metadata", "motion.can_manage_polls", "motion.can_see", "motion.can_see", "motion.can_see", "motion.can_see", "motion.can_see", "motion.can_see", "motion.can_see_internal"},
	"motion.can_manage_metadata":             {"motion.can_see"},
	"motion.can_manage_polls":                {"motion.can_see"},
	"motion.can_see":                         {},
	"motion.can_see_internal":                {"motion.can_see"},
	"motion.can_support":                     {"motion.can_see"},
	"poll.can_manage":                        {},
	"projector.can_manage":                   {"projector.can_see"},
	"projector.can_see":                      {},
	"tag.can_manage":                         {},
	"user.can_manage":                        {"user.can_manage_presence", "user.can_see", "user.can_see", "user.can_see_sensitive_data", "user.can_update"},
	"user.can_manage_presence":               {"user.can_see"},
	"user.can_see":                           {},
	"user.can_see_sensitive_data":            {"user.can_see"},
	"user.can_update":                        {"user.can_see", "user.can_see_sensitive_data"},
}
