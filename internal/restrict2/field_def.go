// Code generated with models.yml DO NOT EDIT.
package restrict

var relationFields = map[string]string{
	"agenda_item/meeting_id":                                                         "meeting/agenda_item_ids",
	"agenda_item/parent_id":                                                          "agenda_item/child_ids",
	"assignment/agenda_item_id":                                                      "agenda_item/content_object_id",
	"assignment/list_of_speakers_id":                                                 "list_of_speakers/content_object_id",
	"assignment/meeting_id":                                                          "meeting/assignment_ids",
	"assignment_candidate/assignment_id":                                             "assignment/candidate_ids",
	"assignment_candidate/meeting_id":                                                "meeting/assignment_candidate_ids",
	"assignment_candidate/meeting_user_id":                                           "meeting_user/assignment_candidate_ids",
	"chat_group/meeting_id":                                                          "meeting/chat_group_ids",
	"chat_message/chat_group_id":                                                     "chat_group/chat_message_ids",
	"chat_message/meeting_id":                                                        "meeting/chat_message_ids",
	"chat_message/meeting_user_id":                                                   "meeting_user/chat_message_ids",
	"committee/default_meeting_id":                                                   "meeting/default_meeting_for_committee_id",
	"committee/forwarding_user_id":                                                   "user/forwarding_committee_ids",
	"committee/organization_id":                                                      "organization/committee_ids",
	"group/admin_group_for_meeting_id":                                               "meeting/admin_group_id",
	"group/default_group_for_meeting_id":                                             "meeting/default_group_id",
	"group/meeting_id":                                                               "meeting/group_ids",
	"group/used_as_assignment_poll_default_id":                                       "meeting/assignment_poll_default_group_ids",
	"group/used_as_motion_poll_default_id":                                           "meeting/motion_poll_default_group_ids",
	"group/used_as_poll_default_id":                                                  "meeting/poll_default_group_ids",
	"group/used_as_topic_poll_default_id":                                            "meeting/topic_poll_default_group_ids",
	"list_of_speakers/meeting_id":                                                    "meeting/list_of_speakers_ids",
	"mediafile/list_of_speakers_id":                                                  "list_of_speakers/content_object_id",
	"mediafile/parent_id":                                                            "mediafile/child_ids",
	"mediafile/used_as_font_bold_in_meeting_id":                                      "meeting/font_bold_id",
	"mediafile/used_as_font_bold_italic_in_meeting_id":                               "meeting/font_bold_italic_id",
	"mediafile/used_as_font_chyron_speaker_name_in_meeting_id":                       "meeting/font_chyron_speaker_name_id",
	"mediafile/used_as_font_italic_in_meeting_id":                                    "meeting/font_italic_id",
	"mediafile/used_as_font_monospace_in_meeting_id":                                 "meeting/font_monospace_id",
	"mediafile/used_as_font_projector_h1_in_meeting_id":                              "meeting/font_projector_h1_id",
	"mediafile/used_as_font_projector_h2_in_meeting_id":                              "meeting/font_projector_h2_id",
	"mediafile/used_as_font_regular_in_meeting_id":                                   "meeting/font_regular_id",
	"mediafile/used_as_logo_pdf_ballot_paper_in_meeting_id":                          "meeting/logo_pdf_ballot_paper_id",
	"mediafile/used_as_logo_pdf_footer_l_in_meeting_id":                              "meeting/logo_pdf_footer_l_id",
	"mediafile/used_as_logo_pdf_footer_r_in_meeting_id":                              "meeting/logo_pdf_footer_r_id",
	"mediafile/used_as_logo_pdf_header_l_in_meeting_id":                              "meeting/logo_pdf_header_l_id",
	"mediafile/used_as_logo_pdf_header_r_in_meeting_id":                              "meeting/logo_pdf_header_r_id",
	"mediafile/used_as_logo_projector_header_in_meeting_id":                          "meeting/logo_projector_header_id",
	"mediafile/used_as_logo_projector_main_in_meeting_id":                            "meeting/logo_projector_main_id",
	"mediafile/used_as_logo_web_header_in_meeting_id":                                "meeting/logo_web_header_id",
	"meeting/admin_group_id":                                                         "group/admin_group_for_meeting_id",
	"meeting/committee_id":                                                           "committee/meeting_ids",
	"meeting/default_group_id":                                                       "group/default_group_for_meeting_id",
	"meeting/default_meeting_for_committee_id":                                       "committee/default_meeting_id",
	"meeting/font_bold_id":                                                           "mediafile/used_as_font_bold_in_meeting_id",
	"meeting/font_bold_italic_id":                                                    "mediafile/used_as_font_bold_italic_in_meeting_id",
	"meeting/font_chyron_speaker_name_id":                                            "mediafile/used_as_font_chyron_speaker_name_in_meeting_id",
	"meeting/font_italic_id":                                                         "mediafile/used_as_font_italic_in_meeting_id",
	"meeting/font_monospace_id":                                                      "mediafile/used_as_font_monospace_in_meeting_id",
	"meeting/font_projector_h1_id":                                                   "mediafile/used_as_font_projector_h1_in_meeting_id",
	"meeting/font_projector_h2_id":                                                   "mediafile/used_as_font_projector_h2_in_meeting_id",
	"meeting/font_regular_id":                                                        "mediafile/used_as_font_regular_in_meeting_id",
	"meeting/is_active_in_organization_id":                                           "organization/active_meeting_ids",
	"meeting/is_archived_in_organization_id":                                         "organization/archived_meeting_ids",
	"meeting/list_of_speakers_countdown_id":                                          "projector_countdown/used_as_list_of_speakers_countdown_meeting_id",
	"meeting/logo_pdf_ballot_paper_id":                                               "mediafile/used_as_logo_pdf_ballot_paper_in_meeting_id",
	"meeting/logo_pdf_footer_l_id":                                                   "mediafile/used_as_logo_pdf_footer_l_in_meeting_id",
	"meeting/logo_pdf_footer_r_id":                                                   "mediafile/used_as_logo_pdf_footer_r_in_meeting_id",
	"meeting/logo_pdf_header_l_id":                                                   "mediafile/used_as_logo_pdf_header_l_in_meeting_id",
	"meeting/logo_pdf_header_r_id":                                                   "mediafile/used_as_logo_pdf_header_r_in_meeting_id",
	"meeting/logo_projector_header_id":                                               "mediafile/used_as_logo_projector_header_in_meeting_id",
	"meeting/logo_projector_main_id":                                                 "mediafile/used_as_logo_projector_main_in_meeting_id",
	"meeting/logo_web_header_id":                                                     "mediafile/used_as_logo_web_header_in_meeting_id",
	"meeting/motions_default_amendment_workflow_id":                                  "motion_workflow/default_amendment_workflow_meeting_id",
	"meeting/motions_default_statute_amendment_workflow_id":                          "motion_workflow/default_statute_amendment_workflow_meeting_id",
	"meeting/motions_default_workflow_id":                                            "motion_workflow/default_workflow_meeting_id",
	"meeting/poll_countdown_id":                                                      "projector_countdown/used_as_poll_countdown_meeting_id",
	"meeting/reference_projector_id":                                                 "projector/used_as_reference_projector_meeting_id",
	"meeting/template_for_organization_id":                                           "organization/template_meeting_ids",
	"meeting_user/meeting_id":                                                        "meeting/meeting_user_ids",
	"meeting_user/user_id":                                                           "user/meeting_user_ids",
	"meeting_user/vote_delegated_to_id":                                              "meeting_user/vote_delegations_from_ids",
	"motion/agenda_item_id":                                                          "agenda_item/content_object_id",
	"motion/block_id":                                                                "motion_block/motion_ids",
	"motion/category_id":                                                             "motion_category/motion_ids",
	"motion/lead_motion_id":                                                          "motion/amendment_ids",
	"motion/list_of_speakers_id":                                                     "list_of_speakers/content_object_id",
	"motion/meeting_id":                                                              "meeting/motion_ids",
	"motion/origin_id":                                                               "motion/derived_motion_ids",
	"motion/origin_meeting_id":                                                       "meeting/forwarded_motion_ids",
	"motion/recommendation_id":                                                       "motion_state/motion_recommendation_ids",
	"motion/sort_parent_id":                                                          "motion/sort_child_ids",
	"motion/state_id":                                                                "motion_state/motion_ids",
	"motion/statute_paragraph_id":                                                    "motion_statute_paragraph/motion_ids",
	"motion_block/agenda_item_id":                                                    "agenda_item/content_object_id",
	"motion_block/list_of_speakers_id":                                               "list_of_speakers/content_object_id",
	"motion_block/meeting_id":                                                        "meeting/motion_block_ids",
	"motion_category/meeting_id":                                                     "meeting/motion_category_ids",
	"motion_category/parent_id":                                                      "motion_category/child_ids",
	"motion_change_recommendation/meeting_id":                                        "meeting/motion_change_recommendation_ids",
	"motion_change_recommendation/motion_id":                                         "motion/change_recommendation_ids",
	"motion_comment/meeting_id":                                                      "meeting/motion_comment_ids",
	"motion_comment/motion_id":                                                       "motion/comment_ids",
	"motion_comment/section_id":                                                      "motion_comment_section/comment_ids",
	"motion_comment_section/meeting_id":                                              "meeting/motion_comment_section_ids",
	"motion_editor/meeting_id":                                                       "meeting/motion_editor_ids",
	"motion_editor/meeting_user_id":                                                  "meeting_user/motion_editor_ids",
	"motion_editor/motion_id":                                                        "motion/editor_ids",
	"motion_state/first_state_of_workflow_id":                                        "motion_workflow/first_state_id",
	"motion_state/meeting_id":                                                        "meeting/motion_state_ids",
	"motion_state/submitter_withdraw_state_id":                                       "motion_state/submitter_withdraw_back_ids",
	"motion_state/workflow_id":                                                       "motion_workflow/state_ids",
	"motion_statute_paragraph/meeting_id":                                            "meeting/motion_statute_paragraph_ids",
	"motion_submitter/meeting_id":                                                    "meeting/motion_submitter_ids",
	"motion_submitter/meeting_user_id":                                               "meeting_user/motion_submitter_ids",
	"motion_submitter/motion_id":                                                     "motion/submitter_ids",
	"motion_workflow/default_amendment_workflow_meeting_id":                          "meeting/motions_default_amendment_workflow_id",
	"motion_workflow/default_statute_amendment_workflow_meeting_id":                  "meeting/motions_default_statute_amendment_workflow_id",
	"motion_workflow/default_workflow_meeting_id":                                    "meeting/motions_default_workflow_id",
	"motion_workflow/first_state_id":                                                 "motion_state/first_state_of_workflow_id",
	"motion_workflow/meeting_id":                                                     "meeting/motion_workflow_ids",
	"motion_working_group_speaker/meeting_id":                                        "meeting/motion_working_group_speaker_ids",
	"motion_working_group_speaker/meeting_user_id":                                   "meeting_user/motion_working_group_speaker_ids",
	"motion_working_group_speaker/motion_id":                                         "motion/working_group_speaker_ids",
	"option/meeting_id":                                                              "meeting/option_ids",
	"option/poll_id":                                                                 "poll/option_ids",
	"option/used_as_global_option_in_poll_id":                                        "poll/global_option_id",
	"organization/theme_id":                                                          "theme/theme_for_organization_id",
	"organization_tag/organization_id":                                               "organization/organization_tag_ids",
	"personal_note/meeting_id":                                                       "meeting/personal_note_ids",
	"personal_note/meeting_user_id":                                                  "meeting_user/personal_note_ids",
	"point_of_order_category/meeting_id":                                             "meeting/point_of_order_category_ids",
	"poll/global_option_id":                                                          "option/used_as_global_option_in_poll_id",
	"poll/meeting_id":                                                                "meeting/poll_ids",
	"poll_candidate/meeting_id":                                                      "meeting/poll_candidate_ids",
	"poll_candidate/poll_candidate_list_id":                                          "poll_candidate_list/poll_candidate_ids",
	"poll_candidate/user_id":                                                         "user/poll_candidate_ids",
	"poll_candidate_list/meeting_id":                                                 "meeting/poll_candidate_list_ids",
	"poll_candidate_list/option_id":                                                  "option/content_object_id",
	"projection/current_projector_id":                                                "projector/current_projection_ids",
	"projection/history_projector_id":                                                "projector/history_projection_ids",
	"projection/meeting_id":                                                          "meeting/all_projection_ids",
	"projection/preview_projector_id":                                                "projector/preview_projection_ids",
	"projector/meeting_id":                                                           "meeting/projector_ids",
	"projector/used_as_default_projector_for_agenda_item_list_in_meeting_id":         "meeting/default_projector_agenda_item_list_ids",
	"projector/used_as_default_projector_for_amendment_in_meeting_id":                "meeting/default_projector_amendment_ids",
	"projector/used_as_default_projector_for_assignment_in_meeting_id":               "meeting/default_projector_assignment_ids",
	"projector/used_as_default_projector_for_assignment_poll_in_meeting_id":          "meeting/default_projector_assignment_poll_ids",
	"projector/used_as_default_projector_for_countdown_in_meeting_id":                "meeting/default_projector_countdown_ids",
	"projector/used_as_default_projector_for_current_list_of_speakers_in_meeting_id": "meeting/default_projector_current_list_of_speakers_ids",
	"projector/used_as_default_projector_for_list_of_speakers_in_meeting_id":         "meeting/default_projector_list_of_speakers_ids",
	"projector/used_as_default_projector_for_mediafile_in_meeting_id":                "meeting/default_projector_mediafile_ids",
	"projector/used_as_default_projector_for_message_in_meeting_id":                  "meeting/default_projector_message_ids",
	"projector/used_as_default_projector_for_motion_block_in_meeting_id":             "meeting/default_projector_motion_block_ids",
	"projector/used_as_default_projector_for_motion_in_meeting_id":                   "meeting/default_projector_motion_ids",
	"projector/used_as_default_projector_for_motion_poll_in_meeting_id":              "meeting/default_projector_motion_poll_ids",
	"projector/used_as_default_projector_for_poll_in_meeting_id":                     "meeting/default_projector_poll_ids",
	"projector/used_as_default_projector_for_topic_in_meeting_id":                    "meeting/default_projector_topic_ids",
	"projector/used_as_reference_projector_meeting_id":                               "meeting/reference_projector_id",
	"projector_countdown/meeting_id":                                                 "meeting/projector_countdown_ids",
	"projector_countdown/used_as_list_of_speakers_countdown_meeting_id":              "meeting/list_of_speakers_countdown_id",
	"projector_countdown/used_as_poll_countdown_meeting_id":                          "meeting/poll_countdown_id",
	"projector_message/meeting_id":                                                   "meeting/projector_message_ids",
	"speaker/list_of_speakers_id":                                                    "list_of_speakers/speaker_ids",
	"speaker/meeting_id":                                                             "meeting/speaker_ids",
	"speaker/meeting_user_id":                                                        "meeting_user/speaker_ids",
	"speaker/point_of_order_category_id":                                             "point_of_order_category/speaker_ids",
	"speaker/structure_level_list_of_speakers_id":                                    "structure_level_list_of_speakers/speaker_ids",
	"structure_level/meeting_id":                                                     "meeting/structure_level_ids",
	"structure_level_list_of_speakers/list_of_speakers_id":                           "list_of_speakers/structure_level_list_of_speakers_ids",
	"structure_level_list_of_speakers/meeting_id":                                    "meeting/structure_level_list_of_speakers_ids",
	"structure_level_list_of_speakers/structure_level_id":                            "structure_level/structure_level_list_of_speakers_ids",
	"tag/meeting_id":                                                                 "meeting/tag_ids",
	"theme/organization_id":                                                          "organization/theme_ids",
	"theme/theme_for_organization_id":                                                "organization/theme_id",
	"topic/agenda_item_id":                                                           "agenda_item/content_object_id",
	"topic/list_of_speakers_id":                                                      "list_of_speakers/content_object_id",
	"topic/meeting_id":                                                               "meeting/topic_ids",
	"user/organization_id":                                                           "organization/user_ids",
	"vote/delegated_user_id":                                                         "user/delegated_vote_ids",
	"vote/meeting_id":                                                                "meeting/vote_ids",
	"vote/option_id":                                                                 "option/vote_ids",
	"vote/user_id":                                                                   "user/vote_ids",
}

var relationListFields = map[string]string{
	"agenda_item/child_ids":                                    "agenda_item/parent_id",
	"agenda_item/projection_ids":                               "projection/content_object_id",
	"agenda_item/tag_ids":                                      "tag/tagged_ids",
	"assignment/attachment_ids":                                "mediafile/attachment_ids",
	"assignment/candidate_ids":                                 "assignment_candidate/assignment_id",
	"assignment/poll_ids":                                      "poll/content_object_id",
	"assignment/projection_ids":                                "projection/content_object_id",
	"assignment/tag_ids":                                       "tag/tagged_ids",
	"chat_group/chat_message_ids":                              "chat_message/chat_group_id",
	"chat_group/read_group_ids":                                "group/read_chat_group_ids",
	"chat_group/write_group_ids":                               "group/write_chat_group_ids",
	"committee/forward_to_committee_ids":                       "committee/receive_forwardings_from_committee_ids",
	"committee/manager_ids":                                    "user/committee_management_ids",
	"committee/meeting_ids":                                    "meeting/committee_id",
	"committee/organization_tag_ids":                           "organization_tag/tagged_ids",
	"committee/receive_forwardings_from_committee_ids":         "committee/forward_to_committee_ids",
	"committee/user_ids":                                       "user/committee_ids",
	"group/mediafile_access_group_ids":                         "mediafile/access_group_ids",
	"group/mediafile_inherited_access_group_ids":               "mediafile/inherited_access_group_ids",
	"group/meeting_user_ids":                                   "meeting_user/group_ids",
	"group/poll_ids":                                           "poll/entitled_group_ids",
	"group/read_chat_group_ids":                                "chat_group/read_group_ids",
	"group/read_comment_section_ids":                           "motion_comment_section/read_group_ids",
	"group/write_chat_group_ids":                               "chat_group/write_group_ids",
	"group/write_comment_section_ids":                          "motion_comment_section/write_group_ids",
	"list_of_speakers/projection_ids":                          "projection/content_object_id",
	"list_of_speakers/speaker_ids":                             "speaker/list_of_speakers_id",
	"list_of_speakers/structure_level_list_of_speakers_ids":    "structure_level_list_of_speakers/list_of_speakers_id",
	"mediafile/access_group_ids":                               "group/mediafile_access_group_ids",
	"mediafile/child_ids":                                      "mediafile/parent_id",
	"mediafile/inherited_access_group_ids":                     "group/mediafile_inherited_access_group_ids",
	"mediafile/projection_ids":                                 "projection/content_object_id",
	"meeting/agenda_item_ids":                                  "agenda_item/meeting_id",
	"meeting/all_projection_ids":                               "projection/meeting_id",
	"meeting/assignment_candidate_ids":                         "assignment_candidate/meeting_id",
	"meeting/assignment_ids":                                   "assignment/meeting_id",
	"meeting/assignment_poll_default_group_ids":                "group/used_as_assignment_poll_default_id",
	"meeting/chat_group_ids":                                   "chat_group/meeting_id",
	"meeting/chat_message_ids":                                 "chat_message/meeting_id",
	"meeting/default_projector_agenda_item_list_ids":           "projector/used_as_default_projector_for_agenda_item_list_in_meeting_id",
	"meeting/default_projector_amendment_ids":                  "projector/used_as_default_projector_for_amendment_in_meeting_id",
	"meeting/default_projector_assignment_ids":                 "projector/used_as_default_projector_for_assignment_in_meeting_id",
	"meeting/default_projector_assignment_poll_ids":            "projector/used_as_default_projector_for_assignment_poll_in_meeting_id",
	"meeting/default_projector_countdown_ids":                  "projector/used_as_default_projector_for_countdown_in_meeting_id",
	"meeting/default_projector_current_list_of_speakers_ids":   "projector/used_as_default_projector_for_current_list_of_speakers_in_meeting_id",
	"meeting/default_projector_list_of_speakers_ids":           "projector/used_as_default_projector_for_list_of_speakers_in_meeting_id",
	"meeting/default_projector_mediafile_ids":                  "projector/used_as_default_projector_for_mediafile_in_meeting_id",
	"meeting/default_projector_message_ids":                    "projector/used_as_default_projector_for_message_in_meeting_id",
	"meeting/default_projector_motion_block_ids":               "projector/used_as_default_projector_for_motion_block_in_meeting_id",
	"meeting/default_projector_motion_ids":                     "projector/used_as_default_projector_for_motion_in_meeting_id",
	"meeting/default_projector_motion_poll_ids":                "projector/used_as_default_projector_for_motion_poll_in_meeting_id",
	"meeting/default_projector_poll_ids":                       "projector/used_as_default_projector_for_poll_in_meeting_id",
	"meeting/default_projector_topic_ids":                      "projector/used_as_default_projector_for_topic_in_meeting_id",
	"meeting/forwarded_motion_ids":                             "motion/origin_meeting_id",
	"meeting/group_ids":                                        "group/meeting_id",
	"meeting/list_of_speakers_ids":                             "list_of_speakers/meeting_id",
	"meeting/mediafile_ids":                                    "mediafile/owner_id",
	"meeting/meeting_user_ids":                                 "meeting_user/meeting_id",
	"meeting/motion_block_ids":                                 "motion_block/meeting_id",
	"meeting/motion_category_ids":                              "motion_category/meeting_id",
	"meeting/motion_change_recommendation_ids":                 "motion_change_recommendation/meeting_id",
	"meeting/motion_comment_ids":                               "motion_comment/meeting_id",
	"meeting/motion_comment_section_ids":                       "motion_comment_section/meeting_id",
	"meeting/motion_editor_ids":                                "motion_editor/meeting_id",
	"meeting/motion_ids":                                       "motion/meeting_id",
	"meeting/motion_poll_default_group_ids":                    "group/used_as_motion_poll_default_id",
	"meeting/motion_state_ids":                                 "motion_state/meeting_id",
	"meeting/motion_statute_paragraph_ids":                     "motion_statute_paragraph/meeting_id",
	"meeting/motion_submitter_ids":                             "motion_submitter/meeting_id",
	"meeting/motion_workflow_ids":                              "motion_workflow/meeting_id",
	"meeting/motion_working_group_speaker_ids":                 "motion_working_group_speaker/meeting_id",
	"meeting/option_ids":                                       "option/meeting_id",
	"meeting/organization_tag_ids":                             "organization_tag/tagged_ids",
	"meeting/personal_note_ids":                                "personal_note/meeting_id",
	"meeting/point_of_order_category_ids":                      "point_of_order_category/meeting_id",
	"meeting/poll_candidate_ids":                               "poll_candidate/meeting_id",
	"meeting/poll_candidate_list_ids":                          "poll_candidate_list/meeting_id",
	"meeting/poll_default_group_ids":                           "group/used_as_poll_default_id",
	"meeting/poll_ids":                                         "poll/meeting_id",
	"meeting/present_user_ids":                                 "user/is_present_in_meeting_ids",
	"meeting/projection_ids":                                   "projection/content_object_id",
	"meeting/projector_countdown_ids":                          "projector_countdown/meeting_id",
	"meeting/projector_ids":                                    "projector/meeting_id",
	"meeting/projector_message_ids":                            "projector_message/meeting_id",
	"meeting/speaker_ids":                                      "speaker/meeting_id",
	"meeting/structure_level_ids":                              "structure_level/meeting_id",
	"meeting/structure_level_list_of_speakers_ids":             "structure_level_list_of_speakers/meeting_id",
	"meeting/tag_ids":                                          "tag/meeting_id",
	"meeting/topic_ids":                                        "topic/meeting_id",
	"meeting/topic_poll_default_group_ids":                     "group/used_as_topic_poll_default_id",
	"meeting/vote_ids":                                         "vote/meeting_id",
	"meeting_user/assignment_candidate_ids":                    "assignment_candidate/meeting_user_id",
	"meeting_user/chat_message_ids":                            "chat_message/meeting_user_id",
	"meeting_user/group_ids":                                   "group/meeting_user_ids",
	"meeting_user/motion_editor_ids":                           "motion_editor/meeting_user_id",
	"meeting_user/motion_submitter_ids":                        "motion_submitter/meeting_user_id",
	"meeting_user/motion_working_group_speaker_ids":            "motion_working_group_speaker/meeting_user_id",
	"meeting_user/personal_note_ids":                           "personal_note/meeting_user_id",
	"meeting_user/speaker_ids":                                 "speaker/meeting_user_id",
	"meeting_user/structure_level_ids":                         "structure_level/meeting_user_ids",
	"meeting_user/supported_motion_ids":                        "motion/supporter_meeting_user_ids",
	"meeting_user/vote_delegations_from_ids":                   "meeting_user/vote_delegated_to_id",
	"motion/all_derived_motion_ids":                            "motion/all_origin_ids",
	"motion/all_origin_ids":                                    "motion/all_derived_motion_ids",
	"motion/amendment_ids":                                     "motion/lead_motion_id",
	"motion/attachment_ids":                                    "mediafile/attachment_ids",
	"motion/change_recommendation_ids":                         "motion_change_recommendation/motion_id",
	"motion/comment_ids":                                       "motion_comment/motion_id",
	"motion/derived_motion_ids":                                "motion/origin_id",
	"motion/editor_ids":                                        "motion_editor/motion_id",
	"motion/identical_motion_ids":                              "motion/identical_motion_ids",
	"motion/option_ids":                                        "option/content_object_id",
	"motion/personal_note_ids":                                 "personal_note/content_object_id",
	"motion/poll_ids":                                          "poll/content_object_id",
	"motion/projection_ids":                                    "projection/content_object_id",
	"motion/referenced_in_motion_recommendation_extension_ids": "motion/recommendation_extension_reference_ids",
	"motion/referenced_in_motion_state_extension_ids":          "motion/state_extension_reference_ids",
	"motion/sort_child_ids":                                    "motion/sort_parent_id",
	"motion/submitter_ids":                                     "motion_submitter/motion_id",
	"motion/supporter_meeting_user_ids":                        "meeting_user/supported_motion_ids",
	"motion/tag_ids":                                           "tag/tagged_ids",
	"motion/working_group_speaker_ids":                         "motion_working_group_speaker/motion_id",
	"motion_block/motion_ids":                                  "motion/block_id",
	"motion_block/projection_ids":                              "projection/content_object_id",
	"motion_category/child_ids":                                "motion_category/parent_id",
	"motion_category/motion_ids":                               "motion/category_id",
	"motion_comment_section/comment_ids":                       "motion_comment/section_id",
	"motion_comment_section/read_group_ids":                    "group/read_comment_section_ids",
	"motion_comment_section/write_group_ids":                   "group/write_comment_section_ids",
	"motion_state/motion_ids":                                  "motion/state_id",
	"motion_state/motion_recommendation_ids":                   "motion/recommendation_id",
	"motion_state/next_state_ids":                              "motion_state/previous_state_ids",
	"motion_state/previous_state_ids":                          "motion_state/next_state_ids",
	"motion_state/submitter_withdraw_back_ids":                 "motion_state/submitter_withdraw_state_id",
	"motion_statute_paragraph/motion_ids":                      "motion/statute_paragraph_id",
	"motion_workflow/state_ids":                                "motion_state/workflow_id",
	"option/vote_ids":                                          "vote/option_id",
	"organization/active_meeting_ids":                          "meeting/is_active_in_organization_id",
	"organization/archived_meeting_ids":                        "meeting/is_archived_in_organization_id",
	"organization/committee_ids":                               "committee/organization_id",
	"organization/mediafile_ids":                               "mediafile/owner_id",
	"organization/organization_tag_ids":                        "organization_tag/organization_id",
	"organization/template_meeting_ids":                        "meeting/template_for_organization_id",
	"organization/theme_ids":                                   "theme/organization_id",
	"organization/user_ids":                                    "user/organization_id",
	"point_of_order_category/speaker_ids":                      "speaker/point_of_order_category_id",
	"poll/entitled_group_ids":                                  "group/poll_ids",
	"poll/option_ids":                                          "option/poll_id",
	"poll/projection_ids":                                      "projection/content_object_id",
	"poll/voted_ids":                                           "user/poll_voted_ids",
	"poll_candidate_list/poll_candidate_ids":                   "poll_candidate/poll_candidate_list_id",
	"projector/current_projection_ids":                         "projection/current_projector_id",
	"projector/history_projection_ids":                         "projection/history_projector_id",
	"projector/preview_projection_ids":                         "projection/preview_projector_id",
	"projector_countdown/projection_ids":                       "projection/content_object_id",
	"projector_message/projection_ids":                         "projection/content_object_id",
	"structure_level/meeting_user_ids":                         "meeting_user/structure_level_ids",
	"structure_level/structure_level_list_of_speakers_ids":     "structure_level_list_of_speakers/structure_level_id",
	"structure_level_list_of_speakers/speaker_ids":             "speaker/structure_level_list_of_speakers_id",
	"topic/attachment_ids":                                     "mediafile/attachment_ids",
	"topic/poll_ids":                                           "poll/content_object_id",
	"topic/projection_ids":                                     "projection/content_object_id",
	"user/committee_ids":                                       "committee/user_ids",
	"user/committee_management_ids":                            "committee/manager_ids",
	"user/delegated_vote_ids":                                  "vote/delegated_user_id",
	"user/forwarding_committee_ids":                            "committee/forwarding_user_id",
	"user/is_present_in_meeting_ids":                           "meeting/present_user_ids",
	"user/meeting_user_ids":                                    "meeting_user/user_id",
	"user/option_ids":                                          "option/content_object_id",
	"user/poll_candidate_ids":                                  "poll_candidate/user_id",
	"user/poll_voted_ids":                                      "poll/voted_ids",
	"user/vote_ids":                                            "vote/user_id",
}

var genericRelationFields = map[string]map[string]string{
	"agenda_item/content_object_id":      {"assignment": "agenda_item_id", "motion": "agenda_item_id", "motion_block": "agenda_item_id", "topic": "agenda_item_id"},
	"list_of_speakers/content_object_id": {"assignment": "list_of_speakers_id", "mediafile": "list_of_speakers_id", "motion": "list_of_speakers_id", "motion_block": "list_of_speakers_id", "topic": "list_of_speakers_id"},
	"mediafile/owner_id":                 {"meeting": "mediafile_ids", "organization": "mediafile_ids"},
	"option/content_object_id":           {"motion": "option_ids", "poll_candidate_list": "option_id", "user": "option_ids"},
	"personal_note/content_object_id":    {"motion": "personal_note_ids"},
	"poll/content_object_id":             {"assignment": "poll_ids", "motion": "poll_ids", "topic": "poll_ids"},
	"projection/content_object_id":       {"agenda_item": "projection_ids", "assignment": "projection_ids", "list_of_speakers": "projection_ids", "mediafile": "projection_ids", "meeting": "projection_ids", "motion": "projection_ids", "motion_block": "projection_ids", "poll": "projection_ids", "projector_countdown": "projection_ids", "projector_message": "projection_ids", "topic": "projection_ids"},
}

var genericRelationListFields = map[string]map[string]string{
	"mediafile/attachment_ids":                      {"assignment": "attachment_ids", "motion": "attachment_ids", "topic": "attachment_ids"},
	"motion/recommendation_extension_reference_ids": {"motion": "referenced_in_motion_recommendation_extension_ids"},
	"motion/state_extension_reference_ids":          {"motion": "referenced_in_motion_state_extension_ids"},
	"organization_tag/tagged_ids":                   {"committee": "organization_tag_ids", "meeting": "organization_tag_ids"},
	"tag/tagged_ids":                                {"agenda_item": "tag_ids", "assignment": "tag_ids", "motion": "tag_ids"},
}

// TODO: Do I still need this?
// // restrictionModes are all fields to there restriction_mode.
// var restrictionModes = map[string]string{
//
// 		// action_worker
//
// 			"action_worker/created": "A",
//
// 			"action_worker/id": "A",
//
// 			"action_worker/name": "A",
//
// 			"action_worker/result": "A",
//
// 			"action_worker/state": "A",
//
// 			"action_worker/timestamp": "A",
//
//
// 		// agenda_item
//
// 			"agenda_item/child_ids": "A",
//
// 			"agenda_item/closed": "A",
//
// 			"agenda_item/content_object_id": "A",
//
// 			"agenda_item/id": "A",
//
// 			"agenda_item/is_hidden": "A",
//
// 			"agenda_item/is_internal": "A",
//
// 			"agenda_item/item_number": "A",
//
// 			"agenda_item/level": "A",
//
// 			"agenda_item/meeting_id": "A",
//
// 			"agenda_item/parent_id": "A",
//
// 			"agenda_item/projection_ids": "A",
//
// 			"agenda_item/tag_ids": "A",
//
// 			"agenda_item/type": "A",
//
// 			"agenda_item/weight": "A",
//
// 			"agenda_item/duration": "B",
//
// 			"agenda_item/comment": "C",
//
// 			"agenda_item/moderator_notes": "D",
//
//
// 		// assignment
//
// 			"assignment/agenda_item_id": "A",
//
// 			"assignment/attachment_ids": "A",
//
// 			"assignment/candidate_ids": "A",
//
// 			"assignment/default_poll_description": "A",
//
// 			"assignment/description": "A",
//
// 			"assignment/id": "A",
//
// 			"assignment/list_of_speakers_id": "A",
//
// 			"assignment/meeting_id": "A",
//
// 			"assignment/number_poll_candidates": "A",
//
// 			"assignment/open_posts": "A",
//
// 			"assignment/phase": "A",
//
// 			"assignment/poll_ids": "A",
//
// 			"assignment/projection_ids": "A",
//
// 			"assignment/sequential_number": "A",
//
// 			"assignment/tag_ids": "A",
//
// 			"assignment/title": "A",
//
//
// 		// assignment_candidate
//
// 			"assignment_candidate/assignment_id": "A",
//
// 			"assignment_candidate/id": "A",
//
// 			"assignment_candidate/meeting_id": "A",
//
// 			"assignment_candidate/meeting_user_id": "A",
//
// 			"assignment_candidate/weight": "A",
//
//
// 		// chat_group
//
// 			"chat_group/chat_message_ids": "A",
//
// 			"chat_group/id": "A",
//
// 			"chat_group/meeting_id": "A",
//
// 			"chat_group/name": "A",
//
// 			"chat_group/read_group_ids": "A",
//
// 			"chat_group/weight": "A",
//
// 			"chat_group/write_group_ids": "A",
//
//
// 		// chat_message
//
// 			"chat_message/chat_group_id": "A",
//
// 			"chat_message/content": "A",
//
// 			"chat_message/created": "A",
//
// 			"chat_message/id": "A",
//
// 			"chat_message/meeting_id": "A",
//
// 			"chat_message/meeting_user_id": "A",
//
//
// 		// committee
//
// 			"committee/default_meeting_id": "A",
//
// 			"committee/description": "A",
//
// 			"committee/external_id": "A",
//
// 			"committee/id": "A",
//
// 			"committee/meeting_ids": "A",
//
// 			"committee/name": "A",
//
// 			"committee/organization_id": "A",
//
// 			"committee/organization_tag_ids": "A",
//
// 			"committee/user_ids": "A",
//
// 			"committee/forward_to_committee_ids": "B",
//
// 			"committee/forwarding_user_id": "B",
//
// 			"committee/manager_ids": "B",
//
// 			"committee/receive_forwardings_from_committee_ids": "B",
//
//
// 		// group
//
// 			"group/admin_group_for_meeting_id": "A",
//
// 			"group/default_group_for_meeting_id": "A",
//
// 			"group/external_id": "A",
//
// 			"group/id": "A",
//
// 			"group/mediafile_access_group_ids": "A",
//
// 			"group/mediafile_inherited_access_group_ids": "A",
//
// 			"group/meeting_id": "A",
//
// 			"group/meeting_user_ids": "A",
//
// 			"group/name": "A",
//
// 			"group/permissions": "A",
//
// 			"group/poll_ids": "A",
//
// 			"group/read_chat_group_ids": "A",
//
// 			"group/read_comment_section_ids": "A",
//
// 			"group/used_as_assignment_poll_default_id": "A",
//
// 			"group/used_as_motion_poll_default_id": "A",
//
// 			"group/used_as_poll_default_id": "A",
//
// 			"group/used_as_topic_poll_default_id": "A",
//
// 			"group/weight": "A",
//
// 			"group/write_chat_group_ids": "A",
//
// 			"group/write_comment_section_ids": "A",
//
//
// 		// import_preview
//
// 			"import_preview/created": "A",
//
// 			"import_preview/id": "A",
//
// 			"import_preview/name": "A",
//
// 			"import_preview/result": "A",
//
// 			"import_preview/state": "A",
//
//
// 		// list_of_speakers
//
// 			"list_of_speakers/closed": "A",
//
// 			"list_of_speakers/content_object_id": "A",
//
// 			"list_of_speakers/id": "A",
//
// 			"list_of_speakers/meeting_id": "A",
//
// 			"list_of_speakers/projection_ids": "A",
//
// 			"list_of_speakers/sequential_number": "A",
//
// 			"list_of_speakers/speaker_ids": "A",
//
// 			"list_of_speakers/structure_level_list_of_speakers_ids": "A",
//
//
// 		// mediafile
//
// 			"mediafile/access_group_ids": "A",
//
// 			"mediafile/attachment_ids": "A",
//
// 			"mediafile/child_ids": "A",
//
// 			"mediafile/create_timestamp": "A",
//
// 			"mediafile/filename": "A",
//
// 			"mediafile/filesize": "A",
//
// 			"mediafile/id": "A",
//
// 			"mediafile/inherited_access_group_ids": "A",
//
// 			"mediafile/is_directory": "A",
//
// 			"mediafile/is_public": "A",
//
// 			"mediafile/list_of_speakers_id": "A",
//
// 			"mediafile/mimetype": "A",
//
// 			"mediafile/owner_id": "A",
//
// 			"mediafile/parent_id": "A",
//
// 			"mediafile/pdf_information": "A",
//
// 			"mediafile/projection_ids": "A",
//
// 			"mediafile/title": "A",
//
// 			"mediafile/token": "A",
//
// 			"mediafile/used_as_font_bold_in_meeting_id": "A",
//
// 			"mediafile/used_as_font_bold_italic_in_meeting_id": "A",
//
// 			"mediafile/used_as_font_chyron_speaker_name_in_meeting_id": "A",
//
// 			"mediafile/used_as_font_italic_in_meeting_id": "A",
//
// 			"mediafile/used_as_font_monospace_in_meeting_id": "A",
//
// 			"mediafile/used_as_font_projector_h1_in_meeting_id": "A",
//
// 			"mediafile/used_as_font_projector_h2_in_meeting_id": "A",
//
// 			"mediafile/used_as_font_regular_in_meeting_id": "A",
//
// 			"mediafile/used_as_logo_pdf_ballot_paper_in_meeting_id": "A",
//
// 			"mediafile/used_as_logo_pdf_footer_l_in_meeting_id": "A",
//
// 			"mediafile/used_as_logo_pdf_footer_r_in_meeting_id": "A",
//
// 			"mediafile/used_as_logo_pdf_header_l_in_meeting_id": "A",
//
// 			"mediafile/used_as_logo_pdf_header_r_in_meeting_id": "A",
//
// 			"mediafile/used_as_logo_projector_header_in_meeting_id": "A",
//
// 			"mediafile/used_as_logo_projector_main_in_meeting_id": "A",
//
// 			"mediafile/used_as_logo_web_header_in_meeting_id": "A",
//
//
// 		// meeting
//
// 			"meeting/enable_anonymous": "A",
//
// 			"meeting/external_id": "A",
//
// 			"meeting/forwarded_motion_ids": "A",
//
// 			"meeting/id": "A",
//
// 			"meeting/is_active_in_organization_id": "A",
//
// 			"meeting/is_archived_in_organization_id": "A",
//
// 			"meeting/language": "A",
//
// 			"meeting/motion_ids": "A",
//
// 			"meeting/name": "A",
//
// 			"meeting/admin_group_id": "B",
//
// 			"meeting/agenda_enable_numbering": "B",
//
// 			"meeting/agenda_item_creation": "B",
//
// 			"meeting/agenda_item_ids": "B",
//
// 			"meeting/agenda_new_items_default_visibility": "B",
//
// 			"meeting/agenda_number_prefix": "B",
//
// 			"meeting/agenda_numeral_system": "B",
//
// 			"meeting/agenda_show_internal_items_on_projector": "B",
//
// 			"meeting/agenda_show_subtitles": "B",
//
// 			"meeting/all_projection_ids": "B",
//
// 			"meeting/assignment_candidate_ids": "B",
//
// 			"meeting/assignment_ids": "B",
//
// 			"meeting/assignment_poll_add_candidates_to_list_of_speakers": "B",
//
// 			"meeting/assignment_poll_ballot_paper_number": "B",
//
// 			"meeting/assignment_poll_ballot_paper_selection": "B",
//
// 			"meeting/assignment_poll_default_backend": "B",
//
// 			"meeting/assignment_poll_default_group_ids": "B",
//
// 			"meeting/assignment_poll_default_method": "B",
//
// 			"meeting/assignment_poll_default_onehundred_percent_base": "B",
//
// 			"meeting/assignment_poll_default_type": "B",
//
// 			"meeting/assignment_poll_enable_max_votes_per_option": "B",
//
// 			"meeting/assignment_poll_sort_poll_result_by_votes": "B",
//
// 			"meeting/assignments_export_preamble": "B",
//
// 			"meeting/assignments_export_title": "B",
//
// 			"meeting/chat_group_ids": "B",
//
// 			"meeting/chat_message_ids": "B",
//
// 			"meeting/committee_id": "B",
//
// 			"meeting/custom_translations": "B",
//
// 			"meeting/default_group_id": "B",
//
// 			"meeting/default_meeting_for_committee_id": "B",
//
// 			"meeting/default_projector_agenda_item_list_ids": "B",
//
// 			"meeting/default_projector_amendment_ids": "B",
//
// 			"meeting/default_projector_assignment_ids": "B",
//
// 			"meeting/default_projector_assignment_poll_ids": "B",
//
// 			"meeting/default_projector_countdown_ids": "B",
//
// 			"meeting/default_projector_current_list_of_speakers_ids": "B",
//
// 			"meeting/default_projector_list_of_speakers_ids": "B",
//
// 			"meeting/default_projector_mediafile_ids": "B",
//
// 			"meeting/default_projector_message_ids": "B",
//
// 			"meeting/default_projector_motion_block_ids": "B",
//
// 			"meeting/default_projector_motion_ids": "B",
//
// 			"meeting/default_projector_motion_poll_ids": "B",
//
// 			"meeting/default_projector_poll_ids": "B",
//
// 			"meeting/default_projector_topic_ids": "B",
//
// 			"meeting/description": "B",
//
// 			"meeting/end_time": "B",
//
// 			"meeting/export_csv_encoding": "B",
//
// 			"meeting/export_csv_separator": "B",
//
// 			"meeting/export_pdf_fontsize": "B",
//
// 			"meeting/export_pdf_line_height": "B",
//
// 			"meeting/export_pdf_page_margin_bottom": "B",
//
// 			"meeting/export_pdf_page_margin_left": "B",
//
// 			"meeting/export_pdf_page_margin_right": "B",
//
// 			"meeting/export_pdf_page_margin_top": "B",
//
// 			"meeting/export_pdf_pagenumber_alignment": "B",
//
// 			"meeting/export_pdf_pagesize": "B",
//
// 			"meeting/font_bold_id": "B",
//
// 			"meeting/font_bold_italic_id": "B",
//
// 			"meeting/font_chyron_speaker_name_id": "B",
//
// 			"meeting/font_italic_id": "B",
//
// 			"meeting/font_monospace_id": "B",
//
// 			"meeting/font_projector_h1_id": "B",
//
// 			"meeting/font_projector_h2_id": "B",
//
// 			"meeting/font_regular_id": "B",
//
// 			"meeting/group_ids": "B",
//
// 			"meeting/imported_at": "B",
//
// 			"meeting/jitsi_domain": "B",
//
// 			"meeting/jitsi_room_name": "B",
//
// 			"meeting/jitsi_room_password": "B",
//
// 			"meeting/list_of_speakers_allow_multiple_speakers": "B",
//
// 			"meeting/list_of_speakers_amount_last_on_projector": "B",
//
// 			"meeting/list_of_speakers_amount_next_on_projector": "B",
//
// 			"meeting/list_of_speakers_can_create_point_of_order_for_others": "B",
//
// 			"meeting/list_of_speakers_can_set_contribution_self": "B",
//
// 			"meeting/list_of_speakers_closing_disables_point_of_order": "B",
//
// 			"meeting/list_of_speakers_countdown_id": "B",
//
// 			"meeting/list_of_speakers_couple_countdown": "B",
//
// 			"meeting/list_of_speakers_default_structure_level_time": "B",
//
// 			"meeting/list_of_speakers_enable_interposed_question": "B",
//
// 			"meeting/list_of_speakers_enable_point_of_order_categories": "B",
//
// 			"meeting/list_of_speakers_enable_point_of_order_speakers": "B",
//
// 			"meeting/list_of_speakers_enable_pro_contra_speech": "B",
//
// 			"meeting/list_of_speakers_ids": "B",
//
// 			"meeting/list_of_speakers_initially_closed": "B",
//
// 			"meeting/list_of_speakers_intervention_time": "B",
//
// 			"meeting/list_of_speakers_present_users_only": "B",
//
// 			"meeting/list_of_speakers_show_amount_of_speakers_on_slide": "B",
//
// 			"meeting/list_of_speakers_show_first_contribution": "B",
//
// 			"meeting/list_of_speakers_speaker_note_for_everyone": "B",
//
// 			"meeting/location": "B",
//
// 			"meeting/logo_pdf_ballot_paper_id": "B",
//
// 			"meeting/logo_pdf_footer_l_id": "B",
//
// 			"meeting/logo_pdf_footer_r_id": "B",
//
// 			"meeting/logo_pdf_header_l_id": "B",
//
// 			"meeting/logo_pdf_header_r_id": "B",
//
// 			"meeting/logo_projector_header_id": "B",
//
// 			"meeting/logo_projector_main_id": "B",
//
// 			"meeting/logo_web_header_id": "B",
//
// 			"meeting/mediafile_ids": "B",
//
// 			"meeting/meeting_user_ids": "B",
//
// 			"meeting/motion_block_ids": "B",
//
// 			"meeting/motion_category_ids": "B",
//
// 			"meeting/motion_change_recommendation_ids": "B",
//
// 			"meeting/motion_comment_ids": "B",
//
// 			"meeting/motion_comment_section_ids": "B",
//
// 			"meeting/motion_editor_ids": "B",
//
// 			"meeting/motion_poll_ballot_paper_number": "B",
//
// 			"meeting/motion_poll_ballot_paper_selection": "B",
//
// 			"meeting/motion_poll_default_backend": "B",
//
// 			"meeting/motion_poll_default_group_ids": "B",
//
// 			"meeting/motion_poll_default_onehundred_percent_base": "B",
//
// 			"meeting/motion_poll_default_type": "B",
//
// 			"meeting/motion_state_ids": "B",
//
// 			"meeting/motion_statute_paragraph_ids": "B",
//
// 			"meeting/motion_submitter_ids": "B",
//
// 			"meeting/motion_workflow_ids": "B",
//
// 			"meeting/motion_working_group_speaker_ids": "B",
//
// 			"meeting/motions_amendments_enabled": "B",
//
// 			"meeting/motions_amendments_in_main_list": "B",
//
// 			"meeting/motions_amendments_multiple_paragraphs": "B",
//
// 			"meeting/motions_amendments_of_amendments": "B",
//
// 			"meeting/motions_amendments_prefix": "B",
//
// 			"meeting/motions_amendments_text_mode": "B",
//
// 			"meeting/motions_block_slide_columns": "B",
//
// 			"meeting/motions_default_amendment_workflow_id": "B",
//
// 			"meeting/motions_default_line_numbering": "B",
//
// 			"meeting/motions_default_sorting": "B",
//
// 			"meeting/motions_default_statute_amendment_workflow_id": "B",
//
// 			"meeting/motions_default_workflow_id": "B",
//
// 			"meeting/motions_enable_editor": "B",
//
// 			"meeting/motions_enable_reason_on_projector": "B",
//
// 			"meeting/motions_enable_recommendation_on_projector": "B",
//
// 			"meeting/motions_enable_sidebox_on_projector": "B",
//
// 			"meeting/motions_enable_text_on_projector": "B",
//
// 			"meeting/motions_enable_working_group_speaker": "B",
//
// 			"meeting/motions_export_follow_recommendation": "B",
//
// 			"meeting/motions_export_preamble": "B",
//
// 			"meeting/motions_export_submitter_recommendation": "B",
//
// 			"meeting/motions_export_title": "B",
//
// 			"meeting/motions_line_length": "B",
//
// 			"meeting/motions_number_min_digits": "B",
//
// 			"meeting/motions_number_type": "B",
//
// 			"meeting/motions_number_with_blank": "B",
//
// 			"meeting/motions_preamble": "B",
//
// 			"meeting/motions_reason_required": "B",
//
// 			"meeting/motions_recommendation_text_mode": "B",
//
// 			"meeting/motions_recommendations_by": "B",
//
// 			"meeting/motions_show_referring_motions": "B",
//
// 			"meeting/motions_show_sequential_number": "B",
//
// 			"meeting/motions_statute_recommendations_by": "B",
//
// 			"meeting/motions_statutes_enabled": "B",
//
// 			"meeting/motions_supporters_min_amount": "B",
//
// 			"meeting/option_ids": "B",
//
// 			"meeting/organization_tag_ids": "B",
//
// 			"meeting/personal_note_ids": "B",
//
// 			"meeting/point_of_order_category_ids": "B",
//
// 			"meeting/poll_ballot_paper_number": "B",
//
// 			"meeting/poll_ballot_paper_selection": "B",
//
// 			"meeting/poll_candidate_ids": "B",
//
// 			"meeting/poll_candidate_list_ids": "B",
//
// 			"meeting/poll_countdown_id": "B",
//
// 			"meeting/poll_couple_countdown": "B",
//
// 			"meeting/poll_default_backend": "B",
//
// 			"meeting/poll_default_group_ids": "B",
//
// 			"meeting/poll_default_method": "B",
//
// 			"meeting/poll_default_onehundred_percent_base": "B",
//
// 			"meeting/poll_default_type": "B",
//
// 			"meeting/poll_ids": "B",
//
// 			"meeting/poll_sort_poll_result_by_votes": "B",
//
// 			"meeting/present_user_ids": "B",
//
// 			"meeting/projection_ids": "B",
//
// 			"meeting/projector_countdown_default_time": "B",
//
// 			"meeting/projector_countdown_ids": "B",
//
// 			"meeting/projector_countdown_warning_time": "B",
//
// 			"meeting/projector_ids": "B",
//
// 			"meeting/projector_message_ids": "B",
//
// 			"meeting/reference_projector_id": "B",
//
// 			"meeting/speaker_ids": "B",
//
// 			"meeting/start_time": "B",
//
// 			"meeting/structure_level_ids": "B",
//
// 			"meeting/structure_level_list_of_speakers_ids": "B",
//
// 			"meeting/tag_ids": "B",
//
// 			"meeting/template_for_organization_id": "B",
//
// 			"meeting/topic_ids": "B",
//
// 			"meeting/topic_poll_default_group_ids": "B",
//
// 			"meeting/user_ids": "B",
//
// 			"meeting/users_allow_self_set_present": "B",
//
// 			"meeting/users_email_body": "B",
//
// 			"meeting/users_email_replyto": "B",
//
// 			"meeting/users_email_sender": "B",
//
// 			"meeting/users_email_subject": "B",
//
// 			"meeting/users_enable_presence_view": "B",
//
// 			"meeting/users_enable_vote_delegations": "B",
//
// 			"meeting/users_enable_vote_weight": "B",
//
// 			"meeting/users_pdf_welcometext": "B",
//
// 			"meeting/users_pdf_welcometitle": "B",
//
// 			"meeting/users_pdf_wlan_encryption": "B",
//
// 			"meeting/users_pdf_wlan_password": "B",
//
// 			"meeting/users_pdf_wlan_ssid": "B",
//
// 			"meeting/vote_ids": "B",
//
// 			"meeting/applause_enable": "C",
//
// 			"meeting/applause_max_amount": "C",
//
// 			"meeting/applause_min_amount": "C",
//
// 			"meeting/applause_particle_image_url": "C",
//
// 			"meeting/applause_show_level": "C",
//
// 			"meeting/applause_timeout": "C",
//
// 			"meeting/applause_type": "C",
//
// 			"meeting/conference_auto_connect": "C",
//
// 			"meeting/conference_auto_connect_next_speakers": "C",
//
// 			"meeting/conference_enable_helpdesk": "C",
//
// 			"meeting/conference_los_restriction": "C",
//
// 			"meeting/conference_open_microphone": "C",
//
// 			"meeting/conference_open_video": "C",
//
// 			"meeting/conference_show": "C",
//
// 			"meeting/conference_stream_poster_url": "C",
//
// 			"meeting/conference_stream_url": "C",
//
// 			"meeting/welcome_text": "C",
//
// 			"meeting/welcome_title": "C",
//
//
// 		// meeting_user
//
// 			"meeting_user/about_me": "A",
//
// 			"meeting_user/assignment_candidate_ids": "A",
//
// 			"meeting_user/chat_message_ids": "A",
//
// 			"meeting_user/group_ids": "A",
//
// 			"meeting_user/id": "A",
//
// 			"meeting_user/meeting_id": "A",
//
// 			"meeting_user/motion_editor_ids": "A",
//
// 			"meeting_user/motion_submitter_ids": "A",
//
// 			"meeting_user/motion_working_group_speaker_ids": "A",
//
// 			"meeting_user/number": "A",
//
// 			"meeting_user/speaker_ids": "A",
//
// 			"meeting_user/structure_level_ids": "A",
//
// 			"meeting_user/supported_motion_ids": "A",
//
// 			"meeting_user/user_id": "A",
//
// 			"meeting_user/vote_delegated_to_id": "A",
//
// 			"meeting_user/vote_delegations_from_ids": "A",
//
// 			"meeting_user/vote_weight": "A",
//
// 			"meeting_user/personal_note_ids": "B",
//
// 			"meeting_user/comment": "D",
//
//
// 		// motion
//
// 			"motion/all_derived_motion_ids": "A",
//
// 			"motion/all_origin_ids": "A",
//
// 			"motion/derived_motion_ids": "A",
//
// 			"motion/forwarded": "A",
//
// 			"motion/id": "A",
//
// 			"motion/meeting_id": "A",
//
// 			"motion/origin_id": "A",
//
// 			"motion/origin_meeting_id": "A",
//
// 			"motion/editor_ids": "B",
//
// 			"motion/working_group_speaker_ids": "B",
//
// 			"motion/agenda_item_id": "C",
//
// 			"motion/amendment_ids": "C",
//
// 			"motion/amendment_paragraphs": "C",
//
// 			"motion/attachment_ids": "C",
//
// 			"motion/block_id": "C",
//
// 			"motion/category_id": "C",
//
// 			"motion/category_weight": "C",
//
// 			"motion/change_recommendation_ids": "C",
//
// 			"motion/comment_ids": "C",
//
// 			"motion/created": "C",
//
// 			"motion/identical_motion_ids": "C",
//
// 			"motion/last_modified": "C",
//
// 			"motion/lead_motion_id": "C",
//
// 			"motion/list_of_speakers_id": "C",
//
// 			"motion/modified_final_version": "C",
//
// 			"motion/number": "C",
//
// 			"motion/option_ids": "C",
//
// 			"motion/personal_note_ids": "C",
//
// 			"motion/poll_ids": "C",
//
// 			"motion/projection_ids": "C",
//
// 			"motion/reason": "C",
//
// 			"motion/recommendation_extension": "C",
//
// 			"motion/recommendation_extension_reference_ids": "C",
//
// 			"motion/referenced_in_motion_recommendation_extension_ids": "C",
//
// 			"motion/referenced_in_motion_state_extension_ids": "C",
//
// 			"motion/sequential_number": "C",
//
// 			"motion/sort_child_ids": "C",
//
// 			"motion/sort_parent_id": "C",
//
// 			"motion/sort_weight": "C",
//
// 			"motion/start_line_number": "C",
//
// 			"motion/state_extension": "C",
//
// 			"motion/state_extension_reference_ids": "C",
//
// 			"motion/state_id": "C",
//
// 			"motion/statute_paragraph_id": "C",
//
// 			"motion/submitter_ids": "C",
//
// 			"motion/supporter_meeting_user_ids": "C",
//
// 			"motion/tag_ids": "C",
//
// 			"motion/text": "C",
//
// 			"motion/title": "C",
//
// 			"motion/workflow_timestamp": "C",
//
// 			"motion/number_value": "D",
//
// 			"motion/text_hash": "D",
//
// 			"motion/recommendation_id": "E",
//
//
// 		// motion_block
//
// 			"motion_block/agenda_item_id": "A",
//
// 			"motion_block/id": "A",
//
// 			"motion_block/internal": "A",
//
// 			"motion_block/list_of_speakers_id": "A",
//
// 			"motion_block/meeting_id": "A",
//
// 			"motion_block/motion_ids": "A",
//
// 			"motion_block/projection_ids": "A",
//
// 			"motion_block/sequential_number": "A",
//
// 			"motion_block/title": "A",
//
//
// 		// motion_category
//
// 			"motion_category/child_ids": "A",
//
// 			"motion_category/id": "A",
//
// 			"motion_category/level": "A",
//
// 			"motion_category/meeting_id": "A",
//
// 			"motion_category/motion_ids": "A",
//
// 			"motion_category/name": "A",
//
// 			"motion_category/parent_id": "A",
//
// 			"motion_category/prefix": "A",
//
// 			"motion_category/sequential_number": "A",
//
// 			"motion_category/weight": "A",
//
//
// 		// motion_change_recommendation
//
// 			"motion_change_recommendation/creation_time": "A",
//
// 			"motion_change_recommendation/id": "A",
//
// 			"motion_change_recommendation/internal": "A",
//
// 			"motion_change_recommendation/line_from": "A",
//
// 			"motion_change_recommendation/line_to": "A",
//
// 			"motion_change_recommendation/meeting_id": "A",
//
// 			"motion_change_recommendation/motion_id": "A",
//
// 			"motion_change_recommendation/other_description": "A",
//
// 			"motion_change_recommendation/rejected": "A",
//
// 			"motion_change_recommendation/text": "A",
//
// 			"motion_change_recommendation/type": "A",
//
//
// 		// motion_comment
//
// 			"motion_comment/comment": "A",
//
// 			"motion_comment/id": "A",
//
// 			"motion_comment/meeting_id": "A",
//
// 			"motion_comment/motion_id": "A",
//
// 			"motion_comment/section_id": "A",
//
//
// 		// motion_comment_section
//
// 			"motion_comment_section/comment_ids": "A",
//
// 			"motion_comment_section/id": "A",
//
// 			"motion_comment_section/meeting_id": "A",
//
// 			"motion_comment_section/name": "A",
//
// 			"motion_comment_section/read_group_ids": "A",
//
// 			"motion_comment_section/sequential_number": "A",
//
// 			"motion_comment_section/submitter_can_write": "A",
//
// 			"motion_comment_section/weight": "A",
//
// 			"motion_comment_section/write_group_ids": "A",
//
//
// 		// motion_editor
//
// 			"motion_editor/id": "A",
//
// 			"motion_editor/meeting_id": "A",
//
// 			"motion_editor/meeting_user_id": "A",
//
// 			"motion_editor/motion_id": "A",
//
// 			"motion_editor/weight": "A",
//
//
// 		// motion_state
//
// 			"motion_state/allow_create_poll": "A",
//
// 			"motion_state/allow_motion_forwarding": "A",
//
// 			"motion_state/allow_submitter_edit": "A",
//
// 			"motion_state/allow_support": "A",
//
// 			"motion_state/css_class": "A",
//
// 			"motion_state/first_state_of_workflow_id": "A",
//
// 			"motion_state/id": "A",
//
// 			"motion_state/is_internal": "A",
//
// 			"motion_state/meeting_id": "A",
//
// 			"motion_state/merge_amendment_into_final": "A",
//
// 			"motion_state/motion_ids": "A",
//
// 			"motion_state/motion_recommendation_ids": "A",
//
// 			"motion_state/name": "A",
//
// 			"motion_state/next_state_ids": "A",
//
// 			"motion_state/previous_state_ids": "A",
//
// 			"motion_state/recommendation_label": "A",
//
// 			"motion_state/restrictions": "A",
//
// 			"motion_state/set_number": "A",
//
// 			"motion_state/set_workflow_timestamp": "A",
//
// 			"motion_state/show_recommendation_extension_field": "A",
//
// 			"motion_state/show_state_extension_field": "A",
//
// 			"motion_state/submitter_withdraw_back_ids": "A",
//
// 			"motion_state/submitter_withdraw_state_id": "A",
//
// 			"motion_state/weight": "A",
//
// 			"motion_state/workflow_id": "A",
//
//
// 		// motion_statute_paragraph
//
// 			"motion_statute_paragraph/id": "A",
//
// 			"motion_statute_paragraph/meeting_id": "A",
//
// 			"motion_statute_paragraph/motion_ids": "A",
//
// 			"motion_statute_paragraph/sequential_number": "A",
//
// 			"motion_statute_paragraph/text": "A",
//
// 			"motion_statute_paragraph/title": "A",
//
// 			"motion_statute_paragraph/weight": "A",
//
//
// 		// motion_submitter
//
// 			"motion_submitter/id": "A",
//
// 			"motion_submitter/meeting_id": "A",
//
// 			"motion_submitter/meeting_user_id": "A",
//
// 			"motion_submitter/motion_id": "A",
//
// 			"motion_submitter/weight": "A",
//
//
// 		// motion_workflow
//
// 			"motion_workflow/default_amendment_workflow_meeting_id": "A",
//
// 			"motion_workflow/default_statute_amendment_workflow_meeting_id": "A",
//
// 			"motion_workflow/default_workflow_meeting_id": "A",
//
// 			"motion_workflow/first_state_id": "A",
//
// 			"motion_workflow/id": "A",
//
// 			"motion_workflow/meeting_id": "A",
//
// 			"motion_workflow/name": "A",
//
// 			"motion_workflow/sequential_number": "A",
//
// 			"motion_workflow/state_ids": "A",
//
//
// 		// motion_working_group_speaker
//
// 			"motion_working_group_speaker/id": "A",
//
// 			"motion_working_group_speaker/meeting_id": "A",
//
// 			"motion_working_group_speaker/meeting_user_id": "A",
//
// 			"motion_working_group_speaker/motion_id": "A",
//
// 			"motion_working_group_speaker/weight": "A",
//
//
// 		// option
//
// 			"option/content_object_id": "A",
//
// 			"option/id": "A",
//
// 			"option/meeting_id": "A",
//
// 			"option/poll_id": "A",
//
// 			"option/text": "A",
//
// 			"option/used_as_global_option_in_poll_id": "A",
//
// 			"option/vote_ids": "A",
//
// 			"option/weight": "A",
//
// 			"option/abstain": "B",
//
// 			"option/no": "B",
//
// 			"option/yes": "B",
//
//
// 		// organization
//
// 			"organization/default_language": "A",
//
// 			"organization/description": "A",
//
// 			"organization/genders": "A",
//
// 			"organization/id": "A",
//
// 			"organization/legal_notice": "A",
//
// 			"organization/login_text": "A",
//
// 			"organization/mediafile_ids": "A",
//
// 			"organization/name": "A",
//
// 			"organization/privacy_policy": "A",
//
// 			"organization/saml_attr_mapping": "A",
//
// 			"organization/saml_enabled": "A",
//
// 			"organization/saml_login_button_text": "A",
//
// 			"organization/saml_metadata_idp": "A",
//
// 			"organization/saml_metadata_sp": "A",
//
// 			"organization/saml_private_key": "A",
//
// 			"organization/template_meeting_ids": "A",
//
// 			"organization/theme_id": "A",
//
// 			"organization/theme_ids": "A",
//
// 			"organization/url": "A",
//
// 			"organization/users_email_body": "A",
//
// 			"organization/users_email_replyto": "A",
//
// 			"organization/users_email_sender": "A",
//
// 			"organization/users_email_subject": "A",
//
// 			"organization/vote_decrypt_public_main_key": "A",
//
// 			"organization/active_meeting_ids": "B",
//
// 			"organization/archived_meeting_ids": "B",
//
// 			"organization/committee_ids": "B",
//
// 			"organization/enable_chat": "B",
//
// 			"organization/enable_electronic_voting": "B",
//
// 			"organization/limit_of_meetings": "B",
//
// 			"organization/limit_of_users": "B",
//
// 			"organization/organization_tag_ids": "B",
//
// 			"organization/reset_password_verbose_errors": "B",
//
// 			"organization/user_ids": "C",
//
//
// 		// organization_tag
//
// 			"organization_tag/color": "A",
//
// 			"organization_tag/id": "A",
//
// 			"organization_tag/name": "A",
//
// 			"organization_tag/organization_id": "A",
//
// 			"organization_tag/tagged_ids": "A",
//
//
// 		// personal_note
//
// 			"personal_note/content_object_id": "A",
//
// 			"personal_note/id": "A",
//
// 			"personal_note/meeting_id": "A",
//
// 			"personal_note/meeting_user_id": "A",
//
// 			"personal_note/note": "A",
//
// 			"personal_note/star": "A",
//
//
// 		// point_of_order_category
//
// 			"point_of_order_category/id": "A",
//
// 			"point_of_order_category/meeting_id": "A",
//
// 			"point_of_order_category/rank": "A",
//
// 			"point_of_order_category/speaker_ids": "A",
//
// 			"point_of_order_category/text": "A",
//
//
// 		// poll
//
// 			"poll/backend": "A",
//
// 			"poll/content_object_id": "A",
//
// 			"poll/crypt_key": "A",
//
// 			"poll/crypt_signature": "A",
//
// 			"poll/description": "A",
//
// 			"poll/entitled_group_ids": "A",
//
// 			"poll/entitled_users_at_stop": "A",
//
// 			"poll/global_abstain": "A",
//
// 			"poll/global_no": "A",
//
// 			"poll/global_option_id": "A",
//
// 			"poll/global_yes": "A",
//
// 			"poll/id": "A",
//
// 			"poll/is_pseudoanonymized": "A",
//
// 			"poll/max_votes_amount": "A",
//
// 			"poll/max_votes_per_option": "A",
//
// 			"poll/meeting_id": "A",
//
// 			"poll/min_votes_amount": "A",
//
// 			"poll/onehundred_percent_base": "A",
//
// 			"poll/option_ids": "A",
//
// 			"poll/pollmethod": "A",
//
// 			"poll/projection_ids": "A",
//
// 			"poll/sequential_number": "A",
//
// 			"poll/state": "A",
//
// 			"poll/title": "A",
//
// 			"poll/type": "A",
//
// 			"poll/voted_ids": "A",
//
// 			"poll/votes_raw": "B",
//
// 			"poll/votes_signature": "B",
//
// 			"poll/votesinvalid": "B",
//
// 			"poll/votesvalid": "B",
//
// 			"poll/vote_count": "C",
//
// 			"poll/votescast": "D",
//
//
// 		// poll_candidate
//
// 			"poll_candidate/id": "A",
//
// 			"poll_candidate/meeting_id": "A",
//
// 			"poll_candidate/poll_candidate_list_id": "A",
//
// 			"poll_candidate/user_id": "A",
//
// 			"poll_candidate/weight": "A",
//
//
// 		// poll_candidate_list
//
// 			"poll_candidate_list/id": "A",
//
// 			"poll_candidate_list/meeting_id": "A",
//
// 			"poll_candidate_list/option_id": "A",
//
// 			"poll_candidate_list/poll_candidate_ids": "A",
//
//
// 		// projection
//
// 			"projection/content": "A",
//
// 			"projection/content_object_id": "A",
//
// 			"projection/current_projector_id": "A",
//
// 			"projection/history_projector_id": "A",
//
// 			"projection/id": "A",
//
// 			"projection/meeting_id": "A",
//
// 			"projection/options": "A",
//
// 			"projection/preview_projector_id": "A",
//
// 			"projection/stable": "A",
//
// 			"projection/type": "A",
//
// 			"projection/weight": "A",
//
//
// 		// projector
//
// 			"projector/aspect_ratio_denominator": "A",
//
// 			"projector/aspect_ratio_numerator": "A",
//
// 			"projector/background_color": "A",
//
// 			"projector/chyron_background_color": "A",
//
// 			"projector/chyron_font_color": "A",
//
// 			"projector/color": "A",
//
// 			"projector/current_projection_ids": "A",
//
// 			"projector/header_background_color": "A",
//
// 			"projector/header_font_color": "A",
//
// 			"projector/header_h1_color": "A",
//
// 			"projector/history_projection_ids": "A",
//
// 			"projector/id": "A",
//
// 			"projector/is_internal": "A",
//
// 			"projector/meeting_id": "A",
//
// 			"projector/name": "A",
//
// 			"projector/preview_projection_ids": "A",
//
// 			"projector/scale": "A",
//
// 			"projector/scroll": "A",
//
// 			"projector/sequential_number": "A",
//
// 			"projector/show_clock": "A",
//
// 			"projector/show_header_footer": "A",
//
// 			"projector/show_logo": "A",
//
// 			"projector/show_title": "A",
//
// 			"projector/used_as_default_projector_for_agenda_item_list_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_amendment_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_assignment_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_assignment_poll_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_countdown_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_current_list_of_speakers_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_list_of_speakers_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_mediafile_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_message_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_motion_block_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_motion_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_motion_poll_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_poll_in_meeting_id": "A",
//
// 			"projector/used_as_default_projector_for_topic_in_meeting_id": "A",
//
// 			"projector/used_as_reference_projector_meeting_id": "A",
//
// 			"projector/width": "A",
//
//
// 		// projector_countdown
//
// 			"projector_countdown/countdown_time": "A",
//
// 			"projector_countdown/default_time": "A",
//
// 			"projector_countdown/description": "A",
//
// 			"projector_countdown/id": "A",
//
// 			"projector_countdown/meeting_id": "A",
//
// 			"projector_countdown/projection_ids": "A",
//
// 			"projector_countdown/running": "A",
//
// 			"projector_countdown/title": "A",
//
// 			"projector_countdown/used_as_list_of_speakers_countdown_meeting_id": "A",
//
// 			"projector_countdown/used_as_poll_countdown_meeting_id": "A",
//
//
// 		// projector_message
//
// 			"projector_message/id": "A",
//
// 			"projector_message/meeting_id": "A",
//
// 			"projector_message/message": "A",
//
// 			"projector_message/projection_ids": "A",
//
//
// 		// speaker
//
// 			"speaker/begin_time": "A",
//
// 			"speaker/end_time": "A",
//
// 			"speaker/id": "A",
//
// 			"speaker/list_of_speakers_id": "A",
//
// 			"speaker/meeting_id": "A",
//
// 			"speaker/meeting_user_id": "A",
//
// 			"speaker/note": "A",
//
// 			"speaker/pause_time": "A",
//
// 			"speaker/point_of_order": "A",
//
// 			"speaker/point_of_order_category_id": "A",
//
// 			"speaker/speech_state": "A",
//
// 			"speaker/structure_level_list_of_speakers_id": "A",
//
// 			"speaker/total_pause": "A",
//
// 			"speaker/unpause_time": "A",
//
// 			"speaker/weight": "A",
//
//
// 		// structure_level
//
// 			"structure_level/color": "A",
//
// 			"structure_level/default_time": "A",
//
// 			"structure_level/id": "A",
//
// 			"structure_level/meeting_id": "A",
//
// 			"structure_level/meeting_user_ids": "A",
//
// 			"structure_level/name": "A",
//
// 			"structure_level/structure_level_list_of_speakers_ids": "A",
//
//
// 		// structure_level_list_of_speakers
//
// 			"structure_level_list_of_speakers/additional_time": "A",
//
// 			"structure_level_list_of_speakers/current_start_time": "A",
//
// 			"structure_level_list_of_speakers/id": "A",
//
// 			"structure_level_list_of_speakers/initial_time": "A",
//
// 			"structure_level_list_of_speakers/list_of_speakers_id": "A",
//
// 			"structure_level_list_of_speakers/meeting_id": "A",
//
// 			"structure_level_list_of_speakers/remaining_time": "A",
//
// 			"structure_level_list_of_speakers/speaker_ids": "A",
//
// 			"structure_level_list_of_speakers/structure_level_id": "A",
//
//
// 		// tag
//
// 			"tag/id": "A",
//
// 			"tag/meeting_id": "A",
//
// 			"tag/name": "A",
//
// 			"tag/tagged_ids": "A",
//
//
// 		// theme
//
// 			"theme/abstain": "A",
//
// 			"theme/accent_100": "A",
//
// 			"theme/accent_200": "A",
//
// 			"theme/accent_300": "A",
//
// 			"theme/accent_400": "A",
//
// 			"theme/accent_50": "A",
//
// 			"theme/accent_500": "A",
//
// 			"theme/accent_600": "A",
//
// 			"theme/accent_700": "A",
//
// 			"theme/accent_800": "A",
//
// 			"theme/accent_900": "A",
//
// 			"theme/accent_a100": "A",
//
// 			"theme/accent_a200": "A",
//
// 			"theme/accent_a400": "A",
//
// 			"theme/accent_a700": "A",
//
// 			"theme/headbar": "A",
//
// 			"theme/id": "A",
//
// 			"theme/name": "A",
//
// 			"theme/no": "A",
//
// 			"theme/organization_id": "A",
//
// 			"theme/primary_100": "A",
//
// 			"theme/primary_200": "A",
//
// 			"theme/primary_300": "A",
//
// 			"theme/primary_400": "A",
//
// 			"theme/primary_50": "A",
//
// 			"theme/primary_500": "A",
//
// 			"theme/primary_600": "A",
//
// 			"theme/primary_700": "A",
//
// 			"theme/primary_800": "A",
//
// 			"theme/primary_900": "A",
//
// 			"theme/primary_a100": "A",
//
// 			"theme/primary_a200": "A",
//
// 			"theme/primary_a400": "A",
//
// 			"theme/primary_a700": "A",
//
// 			"theme/theme_for_organization_id": "A",
//
// 			"theme/warn_100": "A",
//
// 			"theme/warn_200": "A",
//
// 			"theme/warn_300": "A",
//
// 			"theme/warn_400": "A",
//
// 			"theme/warn_50": "A",
//
// 			"theme/warn_500": "A",
//
// 			"theme/warn_600": "A",
//
// 			"theme/warn_700": "A",
//
// 			"theme/warn_800": "A",
//
// 			"theme/warn_900": "A",
//
// 			"theme/warn_a100": "A",
//
// 			"theme/warn_a200": "A",
//
// 			"theme/warn_a400": "A",
//
// 			"theme/warn_a700": "A",
//
// 			"theme/yes": "A",
//
//
// 		// topic
//
// 			"topic/agenda_item_id": "A",
//
// 			"topic/attachment_ids": "A",
//
// 			"topic/id": "A",
//
// 			"topic/list_of_speakers_id": "A",
//
// 			"topic/meeting_id": "A",
//
// 			"topic/poll_ids": "A",
//
// 			"topic/projection_ids": "A",
//
// 			"topic/sequential_number": "A",
//
// 			"topic/text": "A",
//
// 			"topic/title": "A",
//
//
// 		// user
//
// 			"user/default_vote_weight": "A",
//
// 			"user/delegated_vote_ids": "A",
//
// 			"user/first_name": "A",
//
// 			"user/gender": "A",
//
// 			"user/id": "A",
//
// 			"user/is_demo_user": "A",
//
// 			"user/is_physical_person": "A",
//
// 			"user/is_present_in_meeting_ids": "A",
//
// 			"user/last_login": "A",
//
// 			"user/last_name": "A",
//
// 			"user/meeting_user_ids": "A",
//
// 			"user/option_ids": "A",
//
// 			"user/poll_candidate_ids": "A",
//
// 			"user/poll_voted_ids": "A",
//
// 			"user/pronoun": "A",
//
// 			"user/saml_id": "A",
//
// 			"user/title": "A",
//
// 			"user/username": "A",
//
// 			"user/vote_ids": "A",
//
// 			"user/email": "B",
//
// 			"user/can_change_own_password": "D",
//
// 			"user/is_active": "D",
//
// 			"user/last_email_sent": "D",
//
// 			"user/committee_ids": "E",
//
// 			"user/committee_management_ids": "E",
//
// 			"user/forwarding_committee_ids": "E",
//
// 			"user/meeting_ids": "E",
//
// 			"user/organization_management_level": "E",
//
// 			"user/organization_id": "F",
//
// 			"user/password": "G",
//
// 			"user/default_password": "H",
//
//
// 		// vote
//
// 			"vote/delegated_user_id": "A",
//
// 			"vote/id": "A",
//
// 			"vote/meeting_id": "A",
//
// 			"vote/option_id": "A",
//
// 			"vote/user_id": "A",
//
// 			"vote/value": "A",
//
// 			"vote/weight": "A",
//
// 			"vote/user_token": "B",
//
//
// }

// TODO: Do I still need this?
// // collectionFields is an index from a collection name to all of its fieldnames.
// var collectionFields = map[string][]string{
//
// 		"action_worker": { "created", "id", "name", "result", "state", "timestamp",  },
//
// 		"agenda_item": { "child_ids", "closed", "content_object_id", "id", "is_hidden", "is_internal", "item_number", "level", "meeting_id", "parent_id", "projection_ids", "tag_ids", "type", "weight", "duration", "comment", "moderator_notes",  },
//
// 		"assignment": { "agenda_item_id", "attachment_ids", "candidate_ids", "default_poll_description", "description", "id", "list_of_speakers_id", "meeting_id", "number_poll_candidates", "open_posts", "phase", "poll_ids", "projection_ids", "sequential_number", "tag_ids", "title",  },
//
// 		"assignment_candidate": { "assignment_id", "id", "meeting_id", "meeting_user_id", "weight",  },
//
// 		"chat_group": { "chat_message_ids", "id", "meeting_id", "name", "read_group_ids", "weight", "write_group_ids",  },
//
// 		"chat_message": { "chat_group_id", "content", "created", "id", "meeting_id", "meeting_user_id",  },
//
// 		"committee": { "default_meeting_id", "description", "external_id", "id", "meeting_ids", "name", "organization_id", "organization_tag_ids", "user_ids", "forward_to_committee_ids", "forwarding_user_id", "manager_ids", "receive_forwardings_from_committee_ids",  },
//
// 		"group": { "admin_group_for_meeting_id", "default_group_for_meeting_id", "external_id", "id", "mediafile_access_group_ids", "mediafile_inherited_access_group_ids", "meeting_id", "meeting_user_ids", "name", "permissions", "poll_ids", "read_chat_group_ids", "read_comment_section_ids", "used_as_assignment_poll_default_id", "used_as_motion_poll_default_id", "used_as_poll_default_id", "used_as_topic_poll_default_id", "weight", "write_chat_group_ids", "write_comment_section_ids",  },
//
// 		"import_preview": { "created", "id", "name", "result", "state",  },
//
// 		"list_of_speakers": { "closed", "content_object_id", "id", "meeting_id", "projection_ids", "sequential_number", "speaker_ids", "structure_level_list_of_speakers_ids",  },
//
// 		"mediafile": { "access_group_ids", "attachment_ids", "child_ids", "create_timestamp", "filename", "filesize", "id", "inherited_access_group_ids", "is_directory", "is_public", "list_of_speakers_id", "mimetype", "owner_id", "parent_id", "pdf_information", "projection_ids", "title", "token", "used_as_font_bold_in_meeting_id", "used_as_font_bold_italic_in_meeting_id", "used_as_font_chyron_speaker_name_in_meeting_id", "used_as_font_italic_in_meeting_id", "used_as_font_monospace_in_meeting_id", "used_as_font_projector_h1_in_meeting_id", "used_as_font_projector_h2_in_meeting_id", "used_as_font_regular_in_meeting_id", "used_as_logo_pdf_ballot_paper_in_meeting_id", "used_as_logo_pdf_footer_l_in_meeting_id", "used_as_logo_pdf_footer_r_in_meeting_id", "used_as_logo_pdf_header_l_in_meeting_id", "used_as_logo_pdf_header_r_in_meeting_id", "used_as_logo_projector_header_in_meeting_id", "used_as_logo_projector_main_in_meeting_id", "used_as_logo_web_header_in_meeting_id",  },
//
// 		"meeting": { "enable_anonymous", "external_id", "forwarded_motion_ids", "id", "is_active_in_organization_id", "is_archived_in_organization_id", "language", "motion_ids", "name", "admin_group_id", "agenda_enable_numbering", "agenda_item_creation", "agenda_item_ids", "agenda_new_items_default_visibility", "agenda_number_prefix", "agenda_numeral_system", "agenda_show_internal_items_on_projector", "agenda_show_subtitles", "all_projection_ids", "assignment_candidate_ids", "assignment_ids", "assignment_poll_add_candidates_to_list_of_speakers", "assignment_poll_ballot_paper_number", "assignment_poll_ballot_paper_selection", "assignment_poll_default_backend", "assignment_poll_default_group_ids", "assignment_poll_default_method", "assignment_poll_default_onehundred_percent_base", "assignment_poll_default_type", "assignment_poll_enable_max_votes_per_option", "assignment_poll_sort_poll_result_by_votes", "assignments_export_preamble", "assignments_export_title", "chat_group_ids", "chat_message_ids", "committee_id", "custom_translations", "default_group_id", "default_meeting_for_committee_id", "default_projector_agenda_item_list_ids", "default_projector_amendment_ids", "default_projector_assignment_ids", "default_projector_assignment_poll_ids", "default_projector_countdown_ids", "default_projector_current_list_of_speakers_ids", "default_projector_list_of_speakers_ids", "default_projector_mediafile_ids", "default_projector_message_ids", "default_projector_motion_block_ids", "default_projector_motion_ids", "default_projector_motion_poll_ids", "default_projector_poll_ids", "default_projector_topic_ids", "description", "end_time", "export_csv_encoding", "export_csv_separator", "export_pdf_fontsize", "export_pdf_line_height", "export_pdf_page_margin_bottom", "export_pdf_page_margin_left", "export_pdf_page_margin_right", "export_pdf_page_margin_top", "export_pdf_pagenumber_alignment", "export_pdf_pagesize", "font_bold_id", "font_bold_italic_id", "font_chyron_speaker_name_id", "font_italic_id", "font_monospace_id", "font_projector_h1_id", "font_projector_h2_id", "font_regular_id", "group_ids", "imported_at", "jitsi_domain", "jitsi_room_name", "jitsi_room_password", "list_of_speakers_allow_multiple_speakers", "list_of_speakers_amount_last_on_projector", "list_of_speakers_amount_next_on_projector", "list_of_speakers_can_create_point_of_order_for_others", "list_of_speakers_can_set_contribution_self", "list_of_speakers_closing_disables_point_of_order", "list_of_speakers_countdown_id", "list_of_speakers_couple_countdown", "list_of_speakers_default_structure_level_time", "list_of_speakers_enable_interposed_question", "list_of_speakers_enable_point_of_order_categories", "list_of_speakers_enable_point_of_order_speakers", "list_of_speakers_enable_pro_contra_speech", "list_of_speakers_ids", "list_of_speakers_initially_closed", "list_of_speakers_intervention_time", "list_of_speakers_present_users_only", "list_of_speakers_show_amount_of_speakers_on_slide", "list_of_speakers_show_first_contribution", "list_of_speakers_speaker_note_for_everyone", "location", "logo_pdf_ballot_paper_id", "logo_pdf_footer_l_id", "logo_pdf_footer_r_id", "logo_pdf_header_l_id", "logo_pdf_header_r_id", "logo_projector_header_id", "logo_projector_main_id", "logo_web_header_id", "mediafile_ids", "meeting_user_ids", "motion_block_ids", "motion_category_ids", "motion_change_recommendation_ids", "motion_comment_ids", "motion_comment_section_ids", "motion_editor_ids", "motion_poll_ballot_paper_number", "motion_poll_ballot_paper_selection", "motion_poll_default_backend", "motion_poll_default_group_ids", "motion_poll_default_onehundred_percent_base", "motion_poll_default_type", "motion_state_ids", "motion_statute_paragraph_ids", "motion_submitter_ids", "motion_workflow_ids", "motion_working_group_speaker_ids", "motions_amendments_enabled", "motions_amendments_in_main_list", "motions_amendments_multiple_paragraphs", "motions_amendments_of_amendments", "motions_amendments_prefix", "motions_amendments_text_mode", "motions_block_slide_columns", "motions_default_amendment_workflow_id", "motions_default_line_numbering", "motions_default_sorting", "motions_default_statute_amendment_workflow_id", "motions_default_workflow_id", "motions_enable_editor", "motions_enable_reason_on_projector", "motions_enable_recommendation_on_projector", "motions_enable_sidebox_on_projector", "motions_enable_text_on_projector", "motions_enable_working_group_speaker", "motions_export_follow_recommendation", "motions_export_preamble", "motions_export_submitter_recommendation", "motions_export_title", "motions_line_length", "motions_number_min_digits", "motions_number_type", "motions_number_with_blank", "motions_preamble", "motions_reason_required", "motions_recommendation_text_mode", "motions_recommendations_by", "motions_show_referring_motions", "motions_show_sequential_number", "motions_statute_recommendations_by", "motions_statutes_enabled", "motions_supporters_min_amount", "option_ids", "organization_tag_ids", "personal_note_ids", "point_of_order_category_ids", "poll_ballot_paper_number", "poll_ballot_paper_selection", "poll_candidate_ids", "poll_candidate_list_ids", "poll_countdown_id", "poll_couple_countdown", "poll_default_backend", "poll_default_group_ids", "poll_default_method", "poll_default_onehundred_percent_base", "poll_default_type", "poll_ids", "poll_sort_poll_result_by_votes", "present_user_ids", "projection_ids", "projector_countdown_default_time", "projector_countdown_ids", "projector_countdown_warning_time", "projector_ids", "projector_message_ids", "reference_projector_id", "speaker_ids", "start_time", "structure_level_ids", "structure_level_list_of_speakers_ids", "tag_ids", "template_for_organization_id", "topic_ids", "topic_poll_default_group_ids", "user_ids", "users_allow_self_set_present", "users_email_body", "users_email_replyto", "users_email_sender", "users_email_subject", "users_enable_presence_view", "users_enable_vote_delegations", "users_enable_vote_weight", "users_pdf_welcometext", "users_pdf_welcometitle", "users_pdf_wlan_encryption", "users_pdf_wlan_password", "users_pdf_wlan_ssid", "vote_ids", "applause_enable", "applause_max_amount", "applause_min_amount", "applause_particle_image_url", "applause_show_level", "applause_timeout", "applause_type", "conference_auto_connect", "conference_auto_connect_next_speakers", "conference_enable_helpdesk", "conference_los_restriction", "conference_open_microphone", "conference_open_video", "conference_show", "conference_stream_poster_url", "conference_stream_url", "welcome_text", "welcome_title",  },
//
// 		"meeting_user": { "about_me", "assignment_candidate_ids", "chat_message_ids", "group_ids", "id", "meeting_id", "motion_editor_ids", "motion_submitter_ids", "motion_working_group_speaker_ids", "number", "speaker_ids", "structure_level_ids", "supported_motion_ids", "user_id", "vote_delegated_to_id", "vote_delegations_from_ids", "vote_weight", "personal_note_ids", "comment",  },
//
// 		"motion": { "all_derived_motion_ids", "all_origin_ids", "derived_motion_ids", "forwarded", "id", "meeting_id", "origin_id", "origin_meeting_id", "editor_ids", "working_group_speaker_ids", "agenda_item_id", "amendment_ids", "amendment_paragraphs", "attachment_ids", "block_id", "category_id", "category_weight", "change_recommendation_ids", "comment_ids", "created", "identical_motion_ids", "last_modified", "lead_motion_id", "list_of_speakers_id", "modified_final_version", "number", "option_ids", "personal_note_ids", "poll_ids", "projection_ids", "reason", "recommendation_extension", "recommendation_extension_reference_ids", "referenced_in_motion_recommendation_extension_ids", "referenced_in_motion_state_extension_ids", "sequential_number", "sort_child_ids", "sort_parent_id", "sort_weight", "start_line_number", "state_extension", "state_extension_reference_ids", "state_id", "statute_paragraph_id", "submitter_ids", "supporter_meeting_user_ids", "tag_ids", "text", "title", "workflow_timestamp", "number_value", "text_hash", "recommendation_id",  },
//
// 		"motion_block": { "agenda_item_id", "id", "internal", "list_of_speakers_id", "meeting_id", "motion_ids", "projection_ids", "sequential_number", "title",  },
//
// 		"motion_category": { "child_ids", "id", "level", "meeting_id", "motion_ids", "name", "parent_id", "prefix", "sequential_number", "weight",  },
//
// 		"motion_change_recommendation": { "creation_time", "id", "internal", "line_from", "line_to", "meeting_id", "motion_id", "other_description", "rejected", "text", "type",  },
//
// 		"motion_comment": { "comment", "id", "meeting_id", "motion_id", "section_id",  },
//
// 		"motion_comment_section": { "comment_ids", "id", "meeting_id", "name", "read_group_ids", "sequential_number", "submitter_can_write", "weight", "write_group_ids",  },
//
// 		"motion_editor": { "id", "meeting_id", "meeting_user_id", "motion_id", "weight",  },
//
// 		"motion_state": { "allow_create_poll", "allow_motion_forwarding", "allow_submitter_edit", "allow_support", "css_class", "first_state_of_workflow_id", "id", "is_internal", "meeting_id", "merge_amendment_into_final", "motion_ids", "motion_recommendation_ids", "name", "next_state_ids", "previous_state_ids", "recommendation_label", "restrictions", "set_number", "set_workflow_timestamp", "show_recommendation_extension_field", "show_state_extension_field", "submitter_withdraw_back_ids", "submitter_withdraw_state_id", "weight", "workflow_id",  },
//
// 		"motion_statute_paragraph": { "id", "meeting_id", "motion_ids", "sequential_number", "text", "title", "weight",  },
//
// 		"motion_submitter": { "id", "meeting_id", "meeting_user_id", "motion_id", "weight",  },
//
// 		"motion_workflow": { "default_amendment_workflow_meeting_id", "default_statute_amendment_workflow_meeting_id", "default_workflow_meeting_id", "first_state_id", "id", "meeting_id", "name", "sequential_number", "state_ids",  },
//
// 		"motion_working_group_speaker": { "id", "meeting_id", "meeting_user_id", "motion_id", "weight",  },
//
// 		"option": { "content_object_id", "id", "meeting_id", "poll_id", "text", "used_as_global_option_in_poll_id", "vote_ids", "weight", "abstain", "no", "yes",  },
//
// 		"organization": { "default_language", "description", "genders", "id", "legal_notice", "login_text", "mediafile_ids", "name", "privacy_policy", "saml_attr_mapping", "saml_enabled", "saml_login_button_text", "saml_metadata_idp", "saml_metadata_sp", "saml_private_key", "template_meeting_ids", "theme_id", "theme_ids", "url", "users_email_body", "users_email_replyto", "users_email_sender", "users_email_subject", "vote_decrypt_public_main_key", "active_meeting_ids", "archived_meeting_ids", "committee_ids", "enable_chat", "enable_electronic_voting", "limit_of_meetings", "limit_of_users", "organization_tag_ids", "reset_password_verbose_errors", "user_ids",  },
//
// 		"organization_tag": { "color", "id", "name", "organization_id", "tagged_ids",  },
//
// 		"personal_note": { "content_object_id", "id", "meeting_id", "meeting_user_id", "note", "star",  },
//
// 		"point_of_order_category": { "id", "meeting_id", "rank", "speaker_ids", "text",  },
//
// 		"poll": { "backend", "content_object_id", "crypt_key", "crypt_signature", "description", "entitled_group_ids", "entitled_users_at_stop", "global_abstain", "global_no", "global_option_id", "global_yes", "id", "is_pseudoanonymized", "max_votes_amount", "max_votes_per_option", "meeting_id", "min_votes_amount", "onehundred_percent_base", "option_ids", "pollmethod", "projection_ids", "sequential_number", "state", "title", "type", "voted_ids", "votes_raw", "votes_signature", "votesinvalid", "votesvalid", "vote_count", "votescast",  },
//
// 		"poll_candidate": { "id", "meeting_id", "poll_candidate_list_id", "user_id", "weight",  },
//
// 		"poll_candidate_list": { "id", "meeting_id", "option_id", "poll_candidate_ids",  },
//
// 		"projection": { "content", "content_object_id", "current_projector_id", "history_projector_id", "id", "meeting_id", "options", "preview_projector_id", "stable", "type", "weight",  },
//
// 		"projector": { "aspect_ratio_denominator", "aspect_ratio_numerator", "background_color", "chyron_background_color", "chyron_font_color", "color", "current_projection_ids", "header_background_color", "header_font_color", "header_h1_color", "history_projection_ids", "id", "is_internal", "meeting_id", "name", "preview_projection_ids", "scale", "scroll", "sequential_number", "show_clock", "show_header_footer", "show_logo", "show_title", "used_as_default_projector_for_agenda_item_list_in_meeting_id", "used_as_default_projector_for_amendment_in_meeting_id", "used_as_default_projector_for_assignment_in_meeting_id", "used_as_default_projector_for_assignment_poll_in_meeting_id", "used_as_default_projector_for_countdown_in_meeting_id", "used_as_default_projector_for_current_list_of_speakers_in_meeting_id", "used_as_default_projector_for_list_of_speakers_in_meeting_id", "used_as_default_projector_for_mediafile_in_meeting_id", "used_as_default_projector_for_message_in_meeting_id", "used_as_default_projector_for_motion_block_in_meeting_id", "used_as_default_projector_for_motion_in_meeting_id", "used_as_default_projector_for_motion_poll_in_meeting_id", "used_as_default_projector_for_poll_in_meeting_id", "used_as_default_projector_for_topic_in_meeting_id", "used_as_reference_projector_meeting_id", "width",  },
//
// 		"projector_countdown": { "countdown_time", "default_time", "description", "id", "meeting_id", "projection_ids", "running", "title", "used_as_list_of_speakers_countdown_meeting_id", "used_as_poll_countdown_meeting_id",  },
//
// 		"projector_message": { "id", "meeting_id", "message", "projection_ids",  },
//
// 		"speaker": { "begin_time", "end_time", "id", "list_of_speakers_id", "meeting_id", "meeting_user_id", "note", "pause_time", "point_of_order", "point_of_order_category_id", "speech_state", "structure_level_list_of_speakers_id", "total_pause", "unpause_time", "weight",  },
//
// 		"structure_level": { "color", "default_time", "id", "meeting_id", "meeting_user_ids", "name", "structure_level_list_of_speakers_ids",  },
//
// 		"structure_level_list_of_speakers": { "additional_time", "current_start_time", "id", "initial_time", "list_of_speakers_id", "meeting_id", "remaining_time", "speaker_ids", "structure_level_id",  },
//
// 		"tag": { "id", "meeting_id", "name", "tagged_ids",  },
//
// 		"theme": { "abstain", "accent_100", "accent_200", "accent_300", "accent_400", "accent_50", "accent_500", "accent_600", "accent_700", "accent_800", "accent_900", "accent_a100", "accent_a200", "accent_a400", "accent_a700", "headbar", "id", "name", "no", "organization_id", "primary_100", "primary_200", "primary_300", "primary_400", "primary_50", "primary_500", "primary_600", "primary_700", "primary_800", "primary_900", "primary_a100", "primary_a200", "primary_a400", "primary_a700", "theme_for_organization_id", "warn_100", "warn_200", "warn_300", "warn_400", "warn_50", "warn_500", "warn_600", "warn_700", "warn_800", "warn_900", "warn_a100", "warn_a200", "warn_a400", "warn_a700", "yes",  },
//
// 		"topic": { "agenda_item_id", "attachment_ids", "id", "list_of_speakers_id", "meeting_id", "poll_ids", "projection_ids", "sequential_number", "text", "title",  },
//
// 		"user": { "default_vote_weight", "delegated_vote_ids", "first_name", "gender", "id", "is_demo_user", "is_physical_person", "is_present_in_meeting_ids", "last_login", "last_name", "meeting_user_ids", "option_ids", "poll_candidate_ids", "poll_voted_ids", "pronoun", "saml_id", "title", "username", "vote_ids", "email", "can_change_own_password", "is_active", "last_email_sent", "committee_ids", "committee_management_ids", "forwarding_committee_ids", "meeting_ids", "organization_management_level", "organization_id", "password", "default_password",  },
//
// 		"vote": { "delegated_user_id", "id", "meeting_id", "option_id", "user_id", "value", "weight", "user_token",  },
//
// }
