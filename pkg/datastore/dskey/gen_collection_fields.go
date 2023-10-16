// Code generated with models.yml DO NOT EDIT.
package dskey

var collectionFields = [...]collectionField{
	{"invalid", "key"},
	{"_meta", "update"},
	{"action_worker", "created"},
	{"action_worker", "id"},
	{"action_worker", "name"},
	{"action_worker", "result"},
	{"action_worker", "state"},
	{"action_worker", "timestamp"},
	{"agenda_item", "child_ids"},
	{"agenda_item", "closed"},
	{"agenda_item", "comment"},
	{"agenda_item", "content_object_id"},
	{"agenda_item", "duration"},
	{"agenda_item", "id"},
	{"agenda_item", "is_hidden"},
	{"agenda_item", "is_internal"},
	{"agenda_item", "item_number"},
	{"agenda_item", "level"},
	{"agenda_item", "meeting_id"},
	{"agenda_item", "parent_id"},
	{"agenda_item", "projection_ids"},
	{"agenda_item", "tag_ids"},
	{"agenda_item", "type"},
	{"agenda_item", "weight"},
	{"assignment", "agenda_item_id"},
	{"assignment", "attachment_ids"},
	{"assignment", "candidate_ids"},
	{"assignment", "default_poll_description"},
	{"assignment", "description"},
	{"assignment", "id"},
	{"assignment", "list_of_speakers_id"},
	{"assignment", "meeting_id"},
	{"assignment", "number_poll_candidates"},
	{"assignment", "open_posts"},
	{"assignment", "phase"},
	{"assignment", "poll_ids"},
	{"assignment", "projection_ids"},
	{"assignment", "sequential_number"},
	{"assignment", "tag_ids"},
	{"assignment", "title"},
	{"assignment_candidate", "assignment_id"},
	{"assignment_candidate", "id"},
	{"assignment_candidate", "meeting_id"},
	{"assignment_candidate", "meeting_user_id"},
	{"assignment_candidate", "weight"},
	{"chat_group", "chat_message_ids"},
	{"chat_group", "id"},
	{"chat_group", "meeting_id"},
	{"chat_group", "name"},
	{"chat_group", "read_group_ids"},
	{"chat_group", "weight"},
	{"chat_group", "write_group_ids"},
	{"chat_message", "chat_group_id"},
	{"chat_message", "content"},
	{"chat_message", "created"},
	{"chat_message", "id"},
	{"chat_message", "meeting_id"},
	{"chat_message", "meeting_user_id"},
	{"committee", "default_meeting_id"},
	{"committee", "description"},
	{"committee", "external_id"},
	{"committee", "forward_to_committee_ids"},
	{"committee", "forwarding_user_id"},
	{"committee", "id"},
	{"committee", "manager_ids"},
	{"committee", "meeting_ids"},
	{"committee", "name"},
	{"committee", "organization_id"},
	{"committee", "organization_tag_ids"},
	{"committee", "receive_forwardings_from_committee_ids"},
	{"committee", "user_ids"},
	{"group", "admin_group_for_meeting_id"},
	{"group", "default_group_for_meeting_id"},
	{"group", "external_id"},
	{"group", "id"},
	{"group", "mediafile_access_group_ids"},
	{"group", "mediafile_inherited_access_group_ids"},
	{"group", "meeting_id"},
	{"group", "meeting_user_ids"},
	{"group", "name"},
	{"group", "permissions"},
	{"group", "poll_ids"},
	{"group", "read_chat_group_ids"},
	{"group", "read_comment_section_ids"},
	{"group", "used_as_assignment_poll_default_id"},
	{"group", "used_as_motion_poll_default_id"},
	{"group", "used_as_poll_default_id"},
	{"group", "used_as_topic_poll_default_id"},
	{"group", "weight"},
	{"group", "write_chat_group_ids"},
	{"group", "write_comment_section_ids"},
	{"import_preview", "created"},
	{"import_preview", "id"},
	{"import_preview", "name"},
	{"import_preview", "result"},
	{"import_preview", "state"},
	{"list_of_speakers", "closed"},
	{"list_of_speakers", "content_object_id"},
	{"list_of_speakers", "id"},
	{"list_of_speakers", "meeting_id"},
	{"list_of_speakers", "projection_ids"},
	{"list_of_speakers", "sequential_number"},
	{"list_of_speakers", "speaker_ids"},
	{"mediafile", "access_group_ids"},
	{"mediafile", "attachment_ids"},
	{"mediafile", "child_ids"},
	{"mediafile", "create_timestamp"},
	{"mediafile", "filename"},
	{"mediafile", "filesize"},
	{"mediafile", "id"},
	{"mediafile", "inherited_access_group_ids"},
	{"mediafile", "is_directory"},
	{"mediafile", "is_public"},
	{"mediafile", "list_of_speakers_id"},
	{"mediafile", "mimetype"},
	{"mediafile", "owner_id"},
	{"mediafile", "parent_id"},
	{"mediafile", "pdf_information"},
	{"mediafile", "projection_ids"},
	{"mediafile", "title"},
	{"mediafile", "token"},
	{"mediafile", "used_as_font_bold_in_meeting_id"},
	{"mediafile", "used_as_font_bold_italic_in_meeting_id"},
	{"mediafile", "used_as_font_chyron_speaker_name_in_meeting_id"},
	{"mediafile", "used_as_font_italic_in_meeting_id"},
	{"mediafile", "used_as_font_monospace_in_meeting_id"},
	{"mediafile", "used_as_font_projector_h1_in_meeting_id"},
	{"mediafile", "used_as_font_projector_h2_in_meeting_id"},
	{"mediafile", "used_as_font_regular_in_meeting_id"},
	{"mediafile", "used_as_logo_pdf_ballot_paper_in_meeting_id"},
	{"mediafile", "used_as_logo_pdf_footer_l_in_meeting_id"},
	{"mediafile", "used_as_logo_pdf_footer_r_in_meeting_id"},
	{"mediafile", "used_as_logo_pdf_header_l_in_meeting_id"},
	{"mediafile", "used_as_logo_pdf_header_r_in_meeting_id"},
	{"mediafile", "used_as_logo_projector_header_in_meeting_id"},
	{"mediafile", "used_as_logo_projector_main_in_meeting_id"},
	{"mediafile", "used_as_logo_web_header_in_meeting_id"},
	{"meeting", "admin_group_id"},
	{"meeting", "agenda_enable_numbering"},
	{"meeting", "agenda_item_creation"},
	{"meeting", "agenda_item_ids"},
	{"meeting", "agenda_new_items_default_visibility"},
	{"meeting", "agenda_number_prefix"},
	{"meeting", "agenda_numeral_system"},
	{"meeting", "agenda_show_internal_items_on_projector"},
	{"meeting", "agenda_show_subtitles"},
	{"meeting", "all_projection_ids"},
	{"meeting", "applause_enable"},
	{"meeting", "applause_max_amount"},
	{"meeting", "applause_min_amount"},
	{"meeting", "applause_particle_image_url"},
	{"meeting", "applause_show_level"},
	{"meeting", "applause_timeout"},
	{"meeting", "applause_type"},
	{"meeting", "assignment_candidate_ids"},
	{"meeting", "assignment_ids"},
	{"meeting", "assignment_poll_add_candidates_to_list_of_speakers"},
	{"meeting", "assignment_poll_ballot_paper_number"},
	{"meeting", "assignment_poll_ballot_paper_selection"},
	{"meeting", "assignment_poll_default_backend"},
	{"meeting", "assignment_poll_default_group_ids"},
	{"meeting", "assignment_poll_default_method"},
	{"meeting", "assignment_poll_default_onehundred_percent_base"},
	{"meeting", "assignment_poll_default_type"},
	{"meeting", "assignment_poll_enable_max_votes_per_option"},
	{"meeting", "assignment_poll_sort_poll_result_by_votes"},
	{"meeting", "assignments_export_preamble"},
	{"meeting", "assignments_export_title"},
	{"meeting", "chat_group_ids"},
	{"meeting", "chat_message_ids"},
	{"meeting", "committee_id"},
	{"meeting", "conference_auto_connect"},
	{"meeting", "conference_auto_connect_next_speakers"},
	{"meeting", "conference_enable_helpdesk"},
	{"meeting", "conference_los_restriction"},
	{"meeting", "conference_open_microphone"},
	{"meeting", "conference_open_video"},
	{"meeting", "conference_show"},
	{"meeting", "conference_stream_poster_url"},
	{"meeting", "conference_stream_url"},
	{"meeting", "custom_translations"},
	{"meeting", "default_group_id"},
	{"meeting", "default_meeting_for_committee_id"},
	{"meeting", "default_projector_agenda_item_list_ids"},
	{"meeting", "default_projector_amendment_ids"},
	{"meeting", "default_projector_assignment_ids"},
	{"meeting", "default_projector_assignment_poll_ids"},
	{"meeting", "default_projector_countdown_ids"},
	{"meeting", "default_projector_current_list_of_speakers_ids"},
	{"meeting", "default_projector_list_of_speakers_ids"},
	{"meeting", "default_projector_mediafile_ids"},
	{"meeting", "default_projector_message_ids"},
	{"meeting", "default_projector_motion_block_ids"},
	{"meeting", "default_projector_motion_ids"},
	{"meeting", "default_projector_motion_poll_ids"},
	{"meeting", "default_projector_poll_ids"},
	{"meeting", "default_projector_topic_ids"},
	{"meeting", "description"},
	{"meeting", "enable_anonymous"},
	{"meeting", "end_time"},
	{"meeting", "export_csv_encoding"},
	{"meeting", "export_csv_separator"},
	{"meeting", "export_pdf_fontsize"},
	{"meeting", "export_pdf_line_height"},
	{"meeting", "export_pdf_page_margin_bottom"},
	{"meeting", "export_pdf_page_margin_left"},
	{"meeting", "export_pdf_page_margin_right"},
	{"meeting", "export_pdf_page_margin_top"},
	{"meeting", "export_pdf_pagenumber_alignment"},
	{"meeting", "export_pdf_pagesize"},
	{"meeting", "external_id"},
	{"meeting", "font_bold_id"},
	{"meeting", "font_bold_italic_id"},
	{"meeting", "font_chyron_speaker_name_id"},
	{"meeting", "font_italic_id"},
	{"meeting", "font_monospace_id"},
	{"meeting", "font_projector_h1_id"},
	{"meeting", "font_projector_h2_id"},
	{"meeting", "font_regular_id"},
	{"meeting", "forwarded_motion_ids"},
	{"meeting", "group_ids"},
	{"meeting", "id"},
	{"meeting", "imported_at"},
	{"meeting", "is_active_in_organization_id"},
	{"meeting", "is_archived_in_organization_id"},
	{"meeting", "jitsi_domain"},
	{"meeting", "jitsi_room_name"},
	{"meeting", "jitsi_room_password"},
	{"meeting", "language"},
	{"meeting", "list_of_speakers_amount_last_on_projector"},
	{"meeting", "list_of_speakers_amount_next_on_projector"},
	{"meeting", "list_of_speakers_can_set_contribution_self"},
	{"meeting", "list_of_speakers_closing_disables_point_of_order"},
	{"meeting", "list_of_speakers_countdown_id"},
	{"meeting", "list_of_speakers_couple_countdown"},
	{"meeting", "list_of_speakers_enable_point_of_order_categories"},
	{"meeting", "list_of_speakers_enable_point_of_order_speakers"},
	{"meeting", "list_of_speakers_enable_pro_contra_speech"},
	{"meeting", "list_of_speakers_ids"},
	{"meeting", "list_of_speakers_initially_closed"},
	{"meeting", "list_of_speakers_present_users_only"},
	{"meeting", "list_of_speakers_show_amount_of_speakers_on_slide"},
	{"meeting", "list_of_speakers_show_first_contribution"},
	{"meeting", "list_of_speakers_speaker_note_for_everyone"},
	{"meeting", "location"},
	{"meeting", "logo_pdf_ballot_paper_id"},
	{"meeting", "logo_pdf_footer_l_id"},
	{"meeting", "logo_pdf_footer_r_id"},
	{"meeting", "logo_pdf_header_l_id"},
	{"meeting", "logo_pdf_header_r_id"},
	{"meeting", "logo_projector_header_id"},
	{"meeting", "logo_projector_main_id"},
	{"meeting", "logo_web_header_id"},
	{"meeting", "mediafile_ids"},
	{"meeting", "meeting_user_ids"},
	{"meeting", "motion_block_ids"},
	{"meeting", "motion_category_ids"},
	{"meeting", "motion_change_recommendation_ids"},
	{"meeting", "motion_comment_ids"},
	{"meeting", "motion_comment_section_ids"},
	{"meeting", "motion_ids"},
	{"meeting", "motion_poll_ballot_paper_number"},
	{"meeting", "motion_poll_ballot_paper_selection"},
	{"meeting", "motion_poll_default_backend"},
	{"meeting", "motion_poll_default_group_ids"},
	{"meeting", "motion_poll_default_onehundred_percent_base"},
	{"meeting", "motion_poll_default_type"},
	{"meeting", "motion_state_ids"},
	{"meeting", "motion_statute_paragraph_ids"},
	{"meeting", "motion_submitter_ids"},
	{"meeting", "motion_workflow_ids"},
	{"meeting", "motions_amendments_enabled"},
	{"meeting", "motions_amendments_in_main_list"},
	{"meeting", "motions_amendments_multiple_paragraphs"},
	{"meeting", "motions_amendments_of_amendments"},
	{"meeting", "motions_amendments_prefix"},
	{"meeting", "motions_amendments_text_mode"},
	{"meeting", "motions_block_slide_columns"},
	{"meeting", "motions_default_amendment_workflow_id"},
	{"meeting", "motions_default_line_numbering"},
	{"meeting", "motions_default_sorting"},
	{"meeting", "motions_default_statute_amendment_workflow_id"},
	{"meeting", "motions_default_workflow_id"},
	{"meeting", "motions_enable_reason_on_projector"},
	{"meeting", "motions_enable_recommendation_on_projector"},
	{"meeting", "motions_enable_sidebox_on_projector"},
	{"meeting", "motions_enable_text_on_projector"},
	{"meeting", "motions_export_follow_recommendation"},
	{"meeting", "motions_export_preamble"},
	{"meeting", "motions_export_submitter_recommendation"},
	{"meeting", "motions_export_title"},
	{"meeting", "motions_line_length"},
	{"meeting", "motions_number_min_digits"},
	{"meeting", "motions_number_type"},
	{"meeting", "motions_number_with_blank"},
	{"meeting", "motions_preamble"},
	{"meeting", "motions_reason_required"},
	{"meeting", "motions_recommendation_text_mode"},
	{"meeting", "motions_recommendations_by"},
	{"meeting", "motions_show_referring_motions"},
	{"meeting", "motions_show_sequential_number"},
	{"meeting", "motions_statute_recommendations_by"},
	{"meeting", "motions_statutes_enabled"},
	{"meeting", "motions_supporters_min_amount"},
	{"meeting", "name"},
	{"meeting", "option_ids"},
	{"meeting", "organization_tag_ids"},
	{"meeting", "personal_note_ids"},
	{"meeting", "point_of_order_category_ids"},
	{"meeting", "poll_ballot_paper_number"},
	{"meeting", "poll_ballot_paper_selection"},
	{"meeting", "poll_candidate_ids"},
	{"meeting", "poll_candidate_list_ids"},
	{"meeting", "poll_countdown_id"},
	{"meeting", "poll_couple_countdown"},
	{"meeting", "poll_default_backend"},
	{"meeting", "poll_default_group_ids"},
	{"meeting", "poll_default_method"},
	{"meeting", "poll_default_onehundred_percent_base"},
	{"meeting", "poll_default_type"},
	{"meeting", "poll_ids"},
	{"meeting", "poll_sort_poll_result_by_votes"},
	{"meeting", "present_user_ids"},
	{"meeting", "projection_ids"},
	{"meeting", "projector_countdown_default_time"},
	{"meeting", "projector_countdown_ids"},
	{"meeting", "projector_countdown_warning_time"},
	{"meeting", "projector_ids"},
	{"meeting", "projector_message_ids"},
	{"meeting", "reference_projector_id"},
	{"meeting", "speaker_ids"},
	{"meeting", "start_time"},
	{"meeting", "tag_ids"},
	{"meeting", "template_for_organization_id"},
	{"meeting", "topic_ids"},
	{"meeting", "topic_poll_default_group_ids"},
	{"meeting", "user_ids"},
	{"meeting", "users_allow_self_set_present"},
	{"meeting", "users_email_body"},
	{"meeting", "users_email_replyto"},
	{"meeting", "users_email_sender"},
	{"meeting", "users_email_subject"},
	{"meeting", "users_enable_presence_view"},
	{"meeting", "users_enable_vote_delegations"},
	{"meeting", "users_enable_vote_weight"},
	{"meeting", "users_pdf_welcometext"},
	{"meeting", "users_pdf_welcometitle"},
	{"meeting", "users_pdf_wlan_encryption"},
	{"meeting", "users_pdf_wlan_password"},
	{"meeting", "users_pdf_wlan_ssid"},
	{"meeting", "vote_ids"},
	{"meeting", "welcome_text"},
	{"meeting", "welcome_title"},
	{"meeting_user", "about_me"},
	{"meeting_user", "assignment_candidate_ids"},
	{"meeting_user", "chat_message_ids"},
	{"meeting_user", "comment"},
	{"meeting_user", "group_ids"},
	{"meeting_user", "id"},
	{"meeting_user", "meeting_id"},
	{"meeting_user", "motion_submitter_ids"},
	{"meeting_user", "number"},
	{"meeting_user", "personal_note_ids"},
	{"meeting_user", "speaker_ids"},
	{"meeting_user", "structure_level"},
	{"meeting_user", "supported_motion_ids"},
	{"meeting_user", "user_id"},
	{"meeting_user", "vote_delegated_to_id"},
	{"meeting_user", "vote_delegations_from_ids"},
	{"meeting_user", "vote_weight"},
	{"motion", "agenda_item_id"},
	{"motion", "all_derived_motion_ids"},
	{"motion", "all_origin_ids"},
	{"motion", "amendment_ids"},
	{"motion", "amendment_paragraphs"},
	{"motion", "attachment_ids"},
	{"motion", "block_id"},
	{"motion", "category_id"},
	{"motion", "category_weight"},
	{"motion", "change_recommendation_ids"},
	{"motion", "comment_ids"},
	{"motion", "created"},
	{"motion", "derived_motion_ids"},
	{"motion", "forwarded"},
	{"motion", "id"},
	{"motion", "last_modified"},
	{"motion", "lead_motion_id"},
	{"motion", "list_of_speakers_id"},
	{"motion", "meeting_id"},
	{"motion", "modified_final_version"},
	{"motion", "number"},
	{"motion", "number_value"},
	{"motion", "option_ids"},
	{"motion", "origin_id"},
	{"motion", "origin_meeting_id"},
	{"motion", "personal_note_ids"},
	{"motion", "poll_ids"},
	{"motion", "projection_ids"},
	{"motion", "reason"},
	{"motion", "recommendation_extension"},
	{"motion", "recommendation_extension_reference_ids"},
	{"motion", "recommendation_id"},
	{"motion", "referenced_in_motion_recommendation_extension_ids"},
	{"motion", "referenced_in_motion_state_extension_ids"},
	{"motion", "sequential_number"},
	{"motion", "sort_child_ids"},
	{"motion", "sort_parent_id"},
	{"motion", "sort_weight"},
	{"motion", "start_line_number"},
	{"motion", "state_extension"},
	{"motion", "state_extension_reference_ids"},
	{"motion", "state_id"},
	{"motion", "statute_paragraph_id"},
	{"motion", "submitter_ids"},
	{"motion", "supporter_meeting_user_ids"},
	{"motion", "tag_ids"},
	{"motion", "text"},
	{"motion", "title"},
	{"motion", "workflow_timestamp"},
	{"motion_block", "agenda_item_id"},
	{"motion_block", "id"},
	{"motion_block", "internal"},
	{"motion_block", "list_of_speakers_id"},
	{"motion_block", "meeting_id"},
	{"motion_block", "motion_ids"},
	{"motion_block", "projection_ids"},
	{"motion_block", "sequential_number"},
	{"motion_block", "title"},
	{"motion_category", "child_ids"},
	{"motion_category", "id"},
	{"motion_category", "level"},
	{"motion_category", "meeting_id"},
	{"motion_category", "motion_ids"},
	{"motion_category", "name"},
	{"motion_category", "parent_id"},
	{"motion_category", "prefix"},
	{"motion_category", "sequential_number"},
	{"motion_category", "weight"},
	{"motion_change_recommendation", "creation_time"},
	{"motion_change_recommendation", "id"},
	{"motion_change_recommendation", "internal"},
	{"motion_change_recommendation", "line_from"},
	{"motion_change_recommendation", "line_to"},
	{"motion_change_recommendation", "meeting_id"},
	{"motion_change_recommendation", "motion_id"},
	{"motion_change_recommendation", "other_description"},
	{"motion_change_recommendation", "rejected"},
	{"motion_change_recommendation", "text"},
	{"motion_change_recommendation", "type"},
	{"motion_comment", "comment"},
	{"motion_comment", "id"},
	{"motion_comment", "meeting_id"},
	{"motion_comment", "motion_id"},
	{"motion_comment", "section_id"},
	{"motion_comment_section", "comment_ids"},
	{"motion_comment_section", "id"},
	{"motion_comment_section", "meeting_id"},
	{"motion_comment_section", "name"},
	{"motion_comment_section", "read_group_ids"},
	{"motion_comment_section", "sequential_number"},
	{"motion_comment_section", "submitter_can_write"},
	{"motion_comment_section", "weight"},
	{"motion_comment_section", "write_group_ids"},
	{"motion_state", "allow_create_poll"},
	{"motion_state", "allow_motion_forwarding"},
	{"motion_state", "allow_submitter_edit"},
	{"motion_state", "allow_support"},
	{"motion_state", "css_class"},
	{"motion_state", "first_state_of_workflow_id"},
	{"motion_state", "id"},
	{"motion_state", "meeting_id"},
	{"motion_state", "merge_amendment_into_final"},
	{"motion_state", "motion_ids"},
	{"motion_state", "motion_recommendation_ids"},
	{"motion_state", "name"},
	{"motion_state", "next_state_ids"},
	{"motion_state", "previous_state_ids"},
	{"motion_state", "recommendation_label"},
	{"motion_state", "restrictions"},
	{"motion_state", "set_number"},
	{"motion_state", "set_workflow_timestamp"},
	{"motion_state", "show_recommendation_extension_field"},
	{"motion_state", "show_state_extension_field"},
	{"motion_state", "submitter_withdraw_back_ids"},
	{"motion_state", "submitter_withdraw_state_id"},
	{"motion_state", "weight"},
	{"motion_state", "workflow_id"},
	{"motion_statute_paragraph", "id"},
	{"motion_statute_paragraph", "meeting_id"},
	{"motion_statute_paragraph", "motion_ids"},
	{"motion_statute_paragraph", "sequential_number"},
	{"motion_statute_paragraph", "text"},
	{"motion_statute_paragraph", "title"},
	{"motion_statute_paragraph", "weight"},
	{"motion_submitter", "id"},
	{"motion_submitter", "meeting_id"},
	{"motion_submitter", "meeting_user_id"},
	{"motion_submitter", "motion_id"},
	{"motion_submitter", "weight"},
	{"motion_workflow", "default_amendment_workflow_meeting_id"},
	{"motion_workflow", "default_statute_amendment_workflow_meeting_id"},
	{"motion_workflow", "default_workflow_meeting_id"},
	{"motion_workflow", "first_state_id"},
	{"motion_workflow", "id"},
	{"motion_workflow", "meeting_id"},
	{"motion_workflow", "name"},
	{"motion_workflow", "sequential_number"},
	{"motion_workflow", "state_ids"},
	{"option", "abstain"},
	{"option", "content_object_id"},
	{"option", "id"},
	{"option", "meeting_id"},
	{"option", "no"},
	{"option", "poll_id"},
	{"option", "text"},
	{"option", "used_as_global_option_in_poll_id"},
	{"option", "vote_ids"},
	{"option", "weight"},
	{"option", "yes"},
	{"organization", "active_meeting_ids"},
	{"organization", "archived_meeting_ids"},
	{"organization", "committee_ids"},
	{"organization", "default_language"},
	{"organization", "description"},
	{"organization", "enable_chat"},
	{"organization", "enable_electronic_voting"},
	{"organization", "genders"},
	{"organization", "id"},
	{"organization", "legal_notice"},
	{"organization", "limit_of_meetings"},
	{"organization", "limit_of_users"},
	{"organization", "login_text"},
	{"organization", "mediafile_ids"},
	{"organization", "name"},
	{"organization", "organization_tag_ids"},
	{"organization", "privacy_policy"},
	{"organization", "reset_password_verbose_errors"},
	{"organization", "saml_attr_mapping"},
	{"organization", "saml_enabled"},
	{"organization", "saml_login_button_text"},
	{"organization", "saml_metadata_idp"},
	{"organization", "saml_metadata_sp"},
	{"organization", "saml_private_key"},
	{"organization", "template_meeting_ids"},
	{"organization", "theme_id"},
	{"organization", "theme_ids"},
	{"organization", "url"},
	{"organization", "user_ids"},
	{"organization", "users_email_body"},
	{"organization", "users_email_replyto"},
	{"organization", "users_email_sender"},
	{"organization", "users_email_subject"},
	{"organization", "vote_decrypt_public_main_key"},
	{"organization_tag", "color"},
	{"organization_tag", "id"},
	{"organization_tag", "name"},
	{"organization_tag", "organization_id"},
	{"organization_tag", "tagged_ids"},
	{"personal_note", "content_object_id"},
	{"personal_note", "id"},
	{"personal_note", "meeting_id"},
	{"personal_note", "meeting_user_id"},
	{"personal_note", "note"},
	{"personal_note", "star"},
	{"point_of_order_category", "id"},
	{"point_of_order_category", "meeting_id"},
	{"point_of_order_category", "rank"},
	{"point_of_order_category", "speaker_ids"},
	{"point_of_order_category", "text"},
	{"poll", "backend"},
	{"poll", "content_object_id"},
	{"poll", "crypt_key"},
	{"poll", "crypt_signature"},
	{"poll", "description"},
	{"poll", "entitled_group_ids"},
	{"poll", "entitled_users_at_stop"},
	{"poll", "global_abstain"},
	{"poll", "global_no"},
	{"poll", "global_option_id"},
	{"poll", "global_yes"},
	{"poll", "id"},
	{"poll", "is_pseudoanonymized"},
	{"poll", "max_votes_amount"},
	{"poll", "max_votes_per_option"},
	{"poll", "meeting_id"},
	{"poll", "min_votes_amount"},
	{"poll", "onehundred_percent_base"},
	{"poll", "option_ids"},
	{"poll", "pollmethod"},
	{"poll", "projection_ids"},
	{"poll", "sequential_number"},
	{"poll", "state"},
	{"poll", "title"},
	{"poll", "type"},
	{"poll", "vote_count"},
	{"poll", "voted_ids"},
	{"poll", "votes_raw"},
	{"poll", "votes_signature"},
	{"poll", "votescast"},
	{"poll", "votesinvalid"},
	{"poll", "votesvalid"},
	{"poll_candidate", "id"},
	{"poll_candidate", "meeting_id"},
	{"poll_candidate", "poll_candidate_list_id"},
	{"poll_candidate", "user_id"},
	{"poll_candidate", "weight"},
	{"poll_candidate_list", "id"},
	{"poll_candidate_list", "meeting_id"},
	{"poll_candidate_list", "option_id"},
	{"poll_candidate_list", "poll_candidate_ids"},
	{"projection", "content"},
	{"projection", "content_object_id"},
	{"projection", "current_projector_id"},
	{"projection", "history_projector_id"},
	{"projection", "id"},
	{"projection", "meeting_id"},
	{"projection", "options"},
	{"projection", "preview_projector_id"},
	{"projection", "stable"},
	{"projection", "type"},
	{"projection", "weight"},
	{"projector", "aspect_ratio_denominator"},
	{"projector", "aspect_ratio_numerator"},
	{"projector", "background_color"},
	{"projector", "chyron_background_color"},
	{"projector", "chyron_font_color"},
	{"projector", "color"},
	{"projector", "current_projection_ids"},
	{"projector", "header_background_color"},
	{"projector", "header_font_color"},
	{"projector", "header_h1_color"},
	{"projector", "history_projection_ids"},
	{"projector", "id"},
	{"projector", "is_internal"},
	{"projector", "meeting_id"},
	{"projector", "name"},
	{"projector", "preview_projection_ids"},
	{"projector", "scale"},
	{"projector", "scroll"},
	{"projector", "sequential_number"},
	{"projector", "show_clock"},
	{"projector", "show_header_footer"},
	{"projector", "show_logo"},
	{"projector", "show_title"},
	{"projector", "used_as_default_projector_for_agenda_item_list_in_meeting_id"},
	{"projector", "used_as_default_projector_for_amendment_in_meeting_id"},
	{"projector", "used_as_default_projector_for_assignment_in_meeting_id"},
	{"projector", "used_as_default_projector_for_assignment_poll_in_meeting_id"},
	{"projector", "used_as_default_projector_for_countdown_in_meeting_id"},
	{"projector", "used_as_default_projector_for_current_list_of_speakers_in_meeting_id"},
	{"projector", "used_as_default_projector_for_list_of_speakers_in_meeting_id"},
	{"projector", "used_as_default_projector_for_mediafile_in_meeting_id"},
	{"projector", "used_as_default_projector_for_message_in_meeting_id"},
	{"projector", "used_as_default_projector_for_motion_block_in_meeting_id"},
	{"projector", "used_as_default_projector_for_motion_in_meeting_id"},
	{"projector", "used_as_default_projector_for_motion_poll_in_meeting_id"},
	{"projector", "used_as_default_projector_for_poll_in_meeting_id"},
	{"projector", "used_as_default_projector_for_topic_in_meeting_id"},
	{"projector", "used_as_reference_projector_meeting_id"},
	{"projector", "width"},
	{"projector_countdown", "countdown_time"},
	{"projector_countdown", "default_time"},
	{"projector_countdown", "description"},
	{"projector_countdown", "id"},
	{"projector_countdown", "meeting_id"},
	{"projector_countdown", "projection_ids"},
	{"projector_countdown", "running"},
	{"projector_countdown", "title"},
	{"projector_countdown", "used_as_list_of_speakers_countdown_meeting_id"},
	{"projector_countdown", "used_as_poll_countdown_meeting_id"},
	{"projector_message", "id"},
	{"projector_message", "meeting_id"},
	{"projector_message", "message"},
	{"projector_message", "projection_ids"},
	{"speaker", "begin_time"},
	{"speaker", "end_time"},
	{"speaker", "id"},
	{"speaker", "list_of_speakers_id"},
	{"speaker", "meeting_id"},
	{"speaker", "meeting_user_id"},
	{"speaker", "note"},
	{"speaker", "point_of_order"},
	{"speaker", "point_of_order_category_id"},
	{"speaker", "speech_state"},
	{"speaker", "weight"},
	{"tag", "id"},
	{"tag", "meeting_id"},
	{"tag", "name"},
	{"tag", "tagged_ids"},
	{"theme", "abstain"},
	{"theme", "accent_100"},
	{"theme", "accent_200"},
	{"theme", "accent_300"},
	{"theme", "accent_400"},
	{"theme", "accent_50"},
	{"theme", "accent_500"},
	{"theme", "accent_600"},
	{"theme", "accent_700"},
	{"theme", "accent_800"},
	{"theme", "accent_900"},
	{"theme", "accent_a100"},
	{"theme", "accent_a200"},
	{"theme", "accent_a400"},
	{"theme", "accent_a700"},
	{"theme", "headbar"},
	{"theme", "id"},
	{"theme", "name"},
	{"theme", "no"},
	{"theme", "organization_id"},
	{"theme", "primary_100"},
	{"theme", "primary_200"},
	{"theme", "primary_300"},
	{"theme", "primary_400"},
	{"theme", "primary_50"},
	{"theme", "primary_500"},
	{"theme", "primary_600"},
	{"theme", "primary_700"},
	{"theme", "primary_800"},
	{"theme", "primary_900"},
	{"theme", "primary_a100"},
	{"theme", "primary_a200"},
	{"theme", "primary_a400"},
	{"theme", "primary_a700"},
	{"theme", "theme_for_organization_id"},
	{"theme", "warn_100"},
	{"theme", "warn_200"},
	{"theme", "warn_300"},
	{"theme", "warn_400"},
	{"theme", "warn_50"},
	{"theme", "warn_500"},
	{"theme", "warn_600"},
	{"theme", "warn_700"},
	{"theme", "warn_800"},
	{"theme", "warn_900"},
	{"theme", "warn_a100"},
	{"theme", "warn_a200"},
	{"theme", "warn_a400"},
	{"theme", "warn_a700"},
	{"theme", "yes"},
	{"topic", "agenda_item_id"},
	{"topic", "attachment_ids"},
	{"topic", "id"},
	{"topic", "list_of_speakers_id"},
	{"topic", "meeting_id"},
	{"topic", "poll_ids"},
	{"topic", "projection_ids"},
	{"topic", "sequential_number"},
	{"topic", "text"},
	{"topic", "title"},
	{"user", "can_change_own_password"},
	{"user", "committee_ids"},
	{"user", "committee_management_ids"},
	{"user", "default_number"},
	{"user", "default_password"},
	{"user", "default_structure_level"},
	{"user", "default_vote_weight"},
	{"user", "delegated_vote_ids"},
	{"user", "email"},
	{"user", "first_name"},
	{"user", "forwarding_committee_ids"},
	{"user", "gender"},
	{"user", "id"},
	{"user", "is_active"},
	{"user", "is_demo_user"},
	{"user", "is_physical_person"},
	{"user", "is_present_in_meeting_ids"},
	{"user", "last_email_sent"},
	{"user", "last_login"},
	{"user", "last_name"},
	{"user", "meeting_ids"},
	{"user", "meeting_user_ids"},
	{"user", "option_ids"},
	{"user", "organization_id"},
	{"user", "organization_management_level"},
	{"user", "password"},
	{"user", "poll_candidate_ids"},
	{"user", "poll_voted_ids"},
	{"user", "pronoun"},
	{"user", "saml_id"},
	{"user", "title"},
	{"user", "username"},
	{"user", "vote_ids"},
	{"vote", "delegated_user_id"},
	{"vote", "id"},
	{"vote", "meeting_id"},
	{"vote", "option_id"},
	{"vote", "user_id"},
	{"vote", "user_token"},
	{"vote", "value"},
	{"vote", "weight"},
}

func collectionFieldToID(cf string) int {
	switch cf {
	case "_meta/update":
		return 1
	case "action_worker/created":
		return 2
	case "action_worker/id":
		return 3
	case "action_worker/name":
		return 4
	case "action_worker/result":
		return 5
	case "action_worker/state":
		return 6
	case "action_worker/timestamp":
		return 7
	case "agenda_item/child_ids":
		return 8
	case "agenda_item/closed":
		return 9
	case "agenda_item/comment":
		return 10
	case "agenda_item/content_object_id":
		return 11
	case "agenda_item/duration":
		return 12
	case "agenda_item/id":
		return 13
	case "agenda_item/is_hidden":
		return 14
	case "agenda_item/is_internal":
		return 15
	case "agenda_item/item_number":
		return 16
	case "agenda_item/level":
		return 17
	case "agenda_item/meeting_id":
		return 18
	case "agenda_item/parent_id":
		return 19
	case "agenda_item/projection_ids":
		return 20
	case "agenda_item/tag_ids":
		return 21
	case "agenda_item/type":
		return 22
	case "agenda_item/weight":
		return 23
	case "assignment/agenda_item_id":
		return 24
	case "assignment/attachment_ids":
		return 25
	case "assignment/candidate_ids":
		return 26
	case "assignment/default_poll_description":
		return 27
	case "assignment/description":
		return 28
	case "assignment/id":
		return 29
	case "assignment/list_of_speakers_id":
		return 30
	case "assignment/meeting_id":
		return 31
	case "assignment/number_poll_candidates":
		return 32
	case "assignment/open_posts":
		return 33
	case "assignment/phase":
		return 34
	case "assignment/poll_ids":
		return 35
	case "assignment/projection_ids":
		return 36
	case "assignment/sequential_number":
		return 37
	case "assignment/tag_ids":
		return 38
	case "assignment/title":
		return 39
	case "assignment_candidate/assignment_id":
		return 40
	case "assignment_candidate/id":
		return 41
	case "assignment_candidate/meeting_id":
		return 42
	case "assignment_candidate/meeting_user_id":
		return 43
	case "assignment_candidate/weight":
		return 44
	case "chat_group/chat_message_ids":
		return 45
	case "chat_group/id":
		return 46
	case "chat_group/meeting_id":
		return 47
	case "chat_group/name":
		return 48
	case "chat_group/read_group_ids":
		return 49
	case "chat_group/weight":
		return 50
	case "chat_group/write_group_ids":
		return 51
	case "chat_message/chat_group_id":
		return 52
	case "chat_message/content":
		return 53
	case "chat_message/created":
		return 54
	case "chat_message/id":
		return 55
	case "chat_message/meeting_id":
		return 56
	case "chat_message/meeting_user_id":
		return 57
	case "committee/default_meeting_id":
		return 58
	case "committee/description":
		return 59
	case "committee/external_id":
		return 60
	case "committee/forward_to_committee_ids":
		return 61
	case "committee/forwarding_user_id":
		return 62
	case "committee/id":
		return 63
	case "committee/manager_ids":
		return 64
	case "committee/meeting_ids":
		return 65
	case "committee/name":
		return 66
	case "committee/organization_id":
		return 67
	case "committee/organization_tag_ids":
		return 68
	case "committee/receive_forwardings_from_committee_ids":
		return 69
	case "committee/user_ids":
		return 70
	case "group/admin_group_for_meeting_id":
		return 71
	case "group/default_group_for_meeting_id":
		return 72
	case "group/external_id":
		return 73
	case "group/id":
		return 74
	case "group/mediafile_access_group_ids":
		return 75
	case "group/mediafile_inherited_access_group_ids":
		return 76
	case "group/meeting_id":
		return 77
	case "group/meeting_user_ids":
		return 78
	case "group/name":
		return 79
	case "group/permissions":
		return 80
	case "group/poll_ids":
		return 81
	case "group/read_chat_group_ids":
		return 82
	case "group/read_comment_section_ids":
		return 83
	case "group/used_as_assignment_poll_default_id":
		return 84
	case "group/used_as_motion_poll_default_id":
		return 85
	case "group/used_as_poll_default_id":
		return 86
	case "group/used_as_topic_poll_default_id":
		return 87
	case "group/weight":
		return 88
	case "group/write_chat_group_ids":
		return 89
	case "group/write_comment_section_ids":
		return 90
	case "import_preview/created":
		return 91
	case "import_preview/id":
		return 92
	case "import_preview/name":
		return 93
	case "import_preview/result":
		return 94
	case "import_preview/state":
		return 95
	case "list_of_speakers/closed":
		return 96
	case "list_of_speakers/content_object_id":
		return 97
	case "list_of_speakers/id":
		return 98
	case "list_of_speakers/meeting_id":
		return 99
	case "list_of_speakers/projection_ids":
		return 100
	case "list_of_speakers/sequential_number":
		return 101
	case "list_of_speakers/speaker_ids":
		return 102
	case "mediafile/access_group_ids":
		return 103
	case "mediafile/attachment_ids":
		return 104
	case "mediafile/child_ids":
		return 105
	case "mediafile/create_timestamp":
		return 106
	case "mediafile/filename":
		return 107
	case "mediafile/filesize":
		return 108
	case "mediafile/id":
		return 109
	case "mediafile/inherited_access_group_ids":
		return 110
	case "mediafile/is_directory":
		return 111
	case "mediafile/is_public":
		return 112
	case "mediafile/list_of_speakers_id":
		return 113
	case "mediafile/mimetype":
		return 114
	case "mediafile/owner_id":
		return 115
	case "mediafile/parent_id":
		return 116
	case "mediafile/pdf_information":
		return 117
	case "mediafile/projection_ids":
		return 118
	case "mediafile/title":
		return 119
	case "mediafile/token":
		return 120
	case "mediafile/used_as_font_bold_in_meeting_id":
		return 121
	case "mediafile/used_as_font_bold_italic_in_meeting_id":
		return 122
	case "mediafile/used_as_font_chyron_speaker_name_in_meeting_id":
		return 123
	case "mediafile/used_as_font_italic_in_meeting_id":
		return 124
	case "mediafile/used_as_font_monospace_in_meeting_id":
		return 125
	case "mediafile/used_as_font_projector_h1_in_meeting_id":
		return 126
	case "mediafile/used_as_font_projector_h2_in_meeting_id":
		return 127
	case "mediafile/used_as_font_regular_in_meeting_id":
		return 128
	case "mediafile/used_as_logo_pdf_ballot_paper_in_meeting_id":
		return 129
	case "mediafile/used_as_logo_pdf_footer_l_in_meeting_id":
		return 130
	case "mediafile/used_as_logo_pdf_footer_r_in_meeting_id":
		return 131
	case "mediafile/used_as_logo_pdf_header_l_in_meeting_id":
		return 132
	case "mediafile/used_as_logo_pdf_header_r_in_meeting_id":
		return 133
	case "mediafile/used_as_logo_projector_header_in_meeting_id":
		return 134
	case "mediafile/used_as_logo_projector_main_in_meeting_id":
		return 135
	case "mediafile/used_as_logo_web_header_in_meeting_id":
		return 136
	case "meeting/admin_group_id":
		return 137
	case "meeting/agenda_enable_numbering":
		return 138
	case "meeting/agenda_item_creation":
		return 139
	case "meeting/agenda_item_ids":
		return 140
	case "meeting/agenda_new_items_default_visibility":
		return 141
	case "meeting/agenda_number_prefix":
		return 142
	case "meeting/agenda_numeral_system":
		return 143
	case "meeting/agenda_show_internal_items_on_projector":
		return 144
	case "meeting/agenda_show_subtitles":
		return 145
	case "meeting/all_projection_ids":
		return 146
	case "meeting/applause_enable":
		return 147
	case "meeting/applause_max_amount":
		return 148
	case "meeting/applause_min_amount":
		return 149
	case "meeting/applause_particle_image_url":
		return 150
	case "meeting/applause_show_level":
		return 151
	case "meeting/applause_timeout":
		return 152
	case "meeting/applause_type":
		return 153
	case "meeting/assignment_candidate_ids":
		return 154
	case "meeting/assignment_ids":
		return 155
	case "meeting/assignment_poll_add_candidates_to_list_of_speakers":
		return 156
	case "meeting/assignment_poll_ballot_paper_number":
		return 157
	case "meeting/assignment_poll_ballot_paper_selection":
		return 158
	case "meeting/assignment_poll_default_backend":
		return 159
	case "meeting/assignment_poll_default_group_ids":
		return 160
	case "meeting/assignment_poll_default_method":
		return 161
	case "meeting/assignment_poll_default_onehundred_percent_base":
		return 162
	case "meeting/assignment_poll_default_type":
		return 163
	case "meeting/assignment_poll_enable_max_votes_per_option":
		return 164
	case "meeting/assignment_poll_sort_poll_result_by_votes":
		return 165
	case "meeting/assignments_export_preamble":
		return 166
	case "meeting/assignments_export_title":
		return 167
	case "meeting/chat_group_ids":
		return 168
	case "meeting/chat_message_ids":
		return 169
	case "meeting/committee_id":
		return 170
	case "meeting/conference_auto_connect":
		return 171
	case "meeting/conference_auto_connect_next_speakers":
		return 172
	case "meeting/conference_enable_helpdesk":
		return 173
	case "meeting/conference_los_restriction":
		return 174
	case "meeting/conference_open_microphone":
		return 175
	case "meeting/conference_open_video":
		return 176
	case "meeting/conference_show":
		return 177
	case "meeting/conference_stream_poster_url":
		return 178
	case "meeting/conference_stream_url":
		return 179
	case "meeting/custom_translations":
		return 180
	case "meeting/default_group_id":
		return 181
	case "meeting/default_meeting_for_committee_id":
		return 182
	case "meeting/default_projector_agenda_item_list_ids":
		return 183
	case "meeting/default_projector_amendment_ids":
		return 184
	case "meeting/default_projector_assignment_ids":
		return 185
	case "meeting/default_projector_assignment_poll_ids":
		return 186
	case "meeting/default_projector_countdown_ids":
		return 187
	case "meeting/default_projector_current_list_of_speakers_ids":
		return 188
	case "meeting/default_projector_list_of_speakers_ids":
		return 189
	case "meeting/default_projector_mediafile_ids":
		return 190
	case "meeting/default_projector_message_ids":
		return 191
	case "meeting/default_projector_motion_block_ids":
		return 192
	case "meeting/default_projector_motion_ids":
		return 193
	case "meeting/default_projector_motion_poll_ids":
		return 194
	case "meeting/default_projector_poll_ids":
		return 195
	case "meeting/default_projector_topic_ids":
		return 196
	case "meeting/description":
		return 197
	case "meeting/enable_anonymous":
		return 198
	case "meeting/end_time":
		return 199
	case "meeting/export_csv_encoding":
		return 200
	case "meeting/export_csv_separator":
		return 201
	case "meeting/export_pdf_fontsize":
		return 202
	case "meeting/export_pdf_line_height":
		return 203
	case "meeting/export_pdf_page_margin_bottom":
		return 204
	case "meeting/export_pdf_page_margin_left":
		return 205
	case "meeting/export_pdf_page_margin_right":
		return 206
	case "meeting/export_pdf_page_margin_top":
		return 207
	case "meeting/export_pdf_pagenumber_alignment":
		return 208
	case "meeting/export_pdf_pagesize":
		return 209
	case "meeting/external_id":
		return 210
	case "meeting/font_bold_id":
		return 211
	case "meeting/font_bold_italic_id":
		return 212
	case "meeting/font_chyron_speaker_name_id":
		return 213
	case "meeting/font_italic_id":
		return 214
	case "meeting/font_monospace_id":
		return 215
	case "meeting/font_projector_h1_id":
		return 216
	case "meeting/font_projector_h2_id":
		return 217
	case "meeting/font_regular_id":
		return 218
	case "meeting/forwarded_motion_ids":
		return 219
	case "meeting/group_ids":
		return 220
	case "meeting/id":
		return 221
	case "meeting/imported_at":
		return 222
	case "meeting/is_active_in_organization_id":
		return 223
	case "meeting/is_archived_in_organization_id":
		return 224
	case "meeting/jitsi_domain":
		return 225
	case "meeting/jitsi_room_name":
		return 226
	case "meeting/jitsi_room_password":
		return 227
	case "meeting/language":
		return 228
	case "meeting/list_of_speakers_amount_last_on_projector":
		return 229
	case "meeting/list_of_speakers_amount_next_on_projector":
		return 230
	case "meeting/list_of_speakers_can_set_contribution_self":
		return 231
	case "meeting/list_of_speakers_closing_disables_point_of_order":
		return 232
	case "meeting/list_of_speakers_countdown_id":
		return 233
	case "meeting/list_of_speakers_couple_countdown":
		return 234
	case "meeting/list_of_speakers_enable_point_of_order_categories":
		return 235
	case "meeting/list_of_speakers_enable_point_of_order_speakers":
		return 236
	case "meeting/list_of_speakers_enable_pro_contra_speech":
		return 237
	case "meeting/list_of_speakers_ids":
		return 238
	case "meeting/list_of_speakers_initially_closed":
		return 239
	case "meeting/list_of_speakers_present_users_only":
		return 240
	case "meeting/list_of_speakers_show_amount_of_speakers_on_slide":
		return 241
	case "meeting/list_of_speakers_show_first_contribution":
		return 242
	case "meeting/list_of_speakers_speaker_note_for_everyone":
		return 243
	case "meeting/location":
		return 244
	case "meeting/logo_pdf_ballot_paper_id":
		return 245
	case "meeting/logo_pdf_footer_l_id":
		return 246
	case "meeting/logo_pdf_footer_r_id":
		return 247
	case "meeting/logo_pdf_header_l_id":
		return 248
	case "meeting/logo_pdf_header_r_id":
		return 249
	case "meeting/logo_projector_header_id":
		return 250
	case "meeting/logo_projector_main_id":
		return 251
	case "meeting/logo_web_header_id":
		return 252
	case "meeting/mediafile_ids":
		return 253
	case "meeting/meeting_user_ids":
		return 254
	case "meeting/motion_block_ids":
		return 255
	case "meeting/motion_category_ids":
		return 256
	case "meeting/motion_change_recommendation_ids":
		return 257
	case "meeting/motion_comment_ids":
		return 258
	case "meeting/motion_comment_section_ids":
		return 259
	case "meeting/motion_ids":
		return 260
	case "meeting/motion_poll_ballot_paper_number":
		return 261
	case "meeting/motion_poll_ballot_paper_selection":
		return 262
	case "meeting/motion_poll_default_backend":
		return 263
	case "meeting/motion_poll_default_group_ids":
		return 264
	case "meeting/motion_poll_default_onehundred_percent_base":
		return 265
	case "meeting/motion_poll_default_type":
		return 266
	case "meeting/motion_state_ids":
		return 267
	case "meeting/motion_statute_paragraph_ids":
		return 268
	case "meeting/motion_submitter_ids":
		return 269
	case "meeting/motion_workflow_ids":
		return 270
	case "meeting/motions_amendments_enabled":
		return 271
	case "meeting/motions_amendments_in_main_list":
		return 272
	case "meeting/motions_amendments_multiple_paragraphs":
		return 273
	case "meeting/motions_amendments_of_amendments":
		return 274
	case "meeting/motions_amendments_prefix":
		return 275
	case "meeting/motions_amendments_text_mode":
		return 276
	case "meeting/motions_block_slide_columns":
		return 277
	case "meeting/motions_default_amendment_workflow_id":
		return 278
	case "meeting/motions_default_line_numbering":
		return 279
	case "meeting/motions_default_sorting":
		return 280
	case "meeting/motions_default_statute_amendment_workflow_id":
		return 281
	case "meeting/motions_default_workflow_id":
		return 282
	case "meeting/motions_enable_reason_on_projector":
		return 283
	case "meeting/motions_enable_recommendation_on_projector":
		return 284
	case "meeting/motions_enable_sidebox_on_projector":
		return 285
	case "meeting/motions_enable_text_on_projector":
		return 286
	case "meeting/motions_export_follow_recommendation":
		return 287
	case "meeting/motions_export_preamble":
		return 288
	case "meeting/motions_export_submitter_recommendation":
		return 289
	case "meeting/motions_export_title":
		return 290
	case "meeting/motions_line_length":
		return 291
	case "meeting/motions_number_min_digits":
		return 292
	case "meeting/motions_number_type":
		return 293
	case "meeting/motions_number_with_blank":
		return 294
	case "meeting/motions_preamble":
		return 295
	case "meeting/motions_reason_required":
		return 296
	case "meeting/motions_recommendation_text_mode":
		return 297
	case "meeting/motions_recommendations_by":
		return 298
	case "meeting/motions_show_referring_motions":
		return 299
	case "meeting/motions_show_sequential_number":
		return 300
	case "meeting/motions_statute_recommendations_by":
		return 301
	case "meeting/motions_statutes_enabled":
		return 302
	case "meeting/motions_supporters_min_amount":
		return 303
	case "meeting/name":
		return 304
	case "meeting/option_ids":
		return 305
	case "meeting/organization_tag_ids":
		return 306
	case "meeting/personal_note_ids":
		return 307
	case "meeting/point_of_order_category_ids":
		return 308
	case "meeting/poll_ballot_paper_number":
		return 309
	case "meeting/poll_ballot_paper_selection":
		return 310
	case "meeting/poll_candidate_ids":
		return 311
	case "meeting/poll_candidate_list_ids":
		return 312
	case "meeting/poll_countdown_id":
		return 313
	case "meeting/poll_couple_countdown":
		return 314
	case "meeting/poll_default_backend":
		return 315
	case "meeting/poll_default_group_ids":
		return 316
	case "meeting/poll_default_method":
		return 317
	case "meeting/poll_default_onehundred_percent_base":
		return 318
	case "meeting/poll_default_type":
		return 319
	case "meeting/poll_ids":
		return 320
	case "meeting/poll_sort_poll_result_by_votes":
		return 321
	case "meeting/present_user_ids":
		return 322
	case "meeting/projection_ids":
		return 323
	case "meeting/projector_countdown_default_time":
		return 324
	case "meeting/projector_countdown_ids":
		return 325
	case "meeting/projector_countdown_warning_time":
		return 326
	case "meeting/projector_ids":
		return 327
	case "meeting/projector_message_ids":
		return 328
	case "meeting/reference_projector_id":
		return 329
	case "meeting/speaker_ids":
		return 330
	case "meeting/start_time":
		return 331
	case "meeting/tag_ids":
		return 332
	case "meeting/template_for_organization_id":
		return 333
	case "meeting/topic_ids":
		return 334
	case "meeting/topic_poll_default_group_ids":
		return 335
	case "meeting/user_ids":
		return 336
	case "meeting/users_allow_self_set_present":
		return 337
	case "meeting/users_email_body":
		return 338
	case "meeting/users_email_replyto":
		return 339
	case "meeting/users_email_sender":
		return 340
	case "meeting/users_email_subject":
		return 341
	case "meeting/users_enable_presence_view":
		return 342
	case "meeting/users_enable_vote_delegations":
		return 343
	case "meeting/users_enable_vote_weight":
		return 344
	case "meeting/users_pdf_welcometext":
		return 345
	case "meeting/users_pdf_welcometitle":
		return 346
	case "meeting/users_pdf_wlan_encryption":
		return 347
	case "meeting/users_pdf_wlan_password":
		return 348
	case "meeting/users_pdf_wlan_ssid":
		return 349
	case "meeting/vote_ids":
		return 350
	case "meeting/welcome_text":
		return 351
	case "meeting/welcome_title":
		return 352
	case "meeting_user/about_me":
		return 353
	case "meeting_user/assignment_candidate_ids":
		return 354
	case "meeting_user/chat_message_ids":
		return 355
	case "meeting_user/comment":
		return 356
	case "meeting_user/group_ids":
		return 357
	case "meeting_user/id":
		return 358
	case "meeting_user/meeting_id":
		return 359
	case "meeting_user/motion_submitter_ids":
		return 360
	case "meeting_user/number":
		return 361
	case "meeting_user/personal_note_ids":
		return 362
	case "meeting_user/speaker_ids":
		return 363
	case "meeting_user/structure_level":
		return 364
	case "meeting_user/supported_motion_ids":
		return 365
	case "meeting_user/user_id":
		return 366
	case "meeting_user/vote_delegated_to_id":
		return 367
	case "meeting_user/vote_delegations_from_ids":
		return 368
	case "meeting_user/vote_weight":
		return 369
	case "motion/agenda_item_id":
		return 370
	case "motion/all_derived_motion_ids":
		return 371
	case "motion/all_origin_ids":
		return 372
	case "motion/amendment_ids":
		return 373
	case "motion/amendment_paragraphs":
		return 374
	case "motion/attachment_ids":
		return 375
	case "motion/block_id":
		return 376
	case "motion/category_id":
		return 377
	case "motion/category_weight":
		return 378
	case "motion/change_recommendation_ids":
		return 379
	case "motion/comment_ids":
		return 380
	case "motion/created":
		return 381
	case "motion/derived_motion_ids":
		return 382
	case "motion/forwarded":
		return 383
	case "motion/id":
		return 384
	case "motion/last_modified":
		return 385
	case "motion/lead_motion_id":
		return 386
	case "motion/list_of_speakers_id":
		return 387
	case "motion/meeting_id":
		return 388
	case "motion/modified_final_version":
		return 389
	case "motion/number":
		return 390
	case "motion/number_value":
		return 391
	case "motion/option_ids":
		return 392
	case "motion/origin_id":
		return 393
	case "motion/origin_meeting_id":
		return 394
	case "motion/personal_note_ids":
		return 395
	case "motion/poll_ids":
		return 396
	case "motion/projection_ids":
		return 397
	case "motion/reason":
		return 398
	case "motion/recommendation_extension":
		return 399
	case "motion/recommendation_extension_reference_ids":
		return 400
	case "motion/recommendation_id":
		return 401
	case "motion/referenced_in_motion_recommendation_extension_ids":
		return 402
	case "motion/referenced_in_motion_state_extension_ids":
		return 403
	case "motion/sequential_number":
		return 404
	case "motion/sort_child_ids":
		return 405
	case "motion/sort_parent_id":
		return 406
	case "motion/sort_weight":
		return 407
	case "motion/start_line_number":
		return 408
	case "motion/state_extension":
		return 409
	case "motion/state_extension_reference_ids":
		return 410
	case "motion/state_id":
		return 411
	case "motion/statute_paragraph_id":
		return 412
	case "motion/submitter_ids":
		return 413
	case "motion/supporter_meeting_user_ids":
		return 414
	case "motion/tag_ids":
		return 415
	case "motion/text":
		return 416
	case "motion/title":
		return 417
	case "motion/workflow_timestamp":
		return 418
	case "motion_block/agenda_item_id":
		return 419
	case "motion_block/id":
		return 420
	case "motion_block/internal":
		return 421
	case "motion_block/list_of_speakers_id":
		return 422
	case "motion_block/meeting_id":
		return 423
	case "motion_block/motion_ids":
		return 424
	case "motion_block/projection_ids":
		return 425
	case "motion_block/sequential_number":
		return 426
	case "motion_block/title":
		return 427
	case "motion_category/child_ids":
		return 428
	case "motion_category/id":
		return 429
	case "motion_category/level":
		return 430
	case "motion_category/meeting_id":
		return 431
	case "motion_category/motion_ids":
		return 432
	case "motion_category/name":
		return 433
	case "motion_category/parent_id":
		return 434
	case "motion_category/prefix":
		return 435
	case "motion_category/sequential_number":
		return 436
	case "motion_category/weight":
		return 437
	case "motion_change_recommendation/creation_time":
		return 438
	case "motion_change_recommendation/id":
		return 439
	case "motion_change_recommendation/internal":
		return 440
	case "motion_change_recommendation/line_from":
		return 441
	case "motion_change_recommendation/line_to":
		return 442
	case "motion_change_recommendation/meeting_id":
		return 443
	case "motion_change_recommendation/motion_id":
		return 444
	case "motion_change_recommendation/other_description":
		return 445
	case "motion_change_recommendation/rejected":
		return 446
	case "motion_change_recommendation/text":
		return 447
	case "motion_change_recommendation/type":
		return 448
	case "motion_comment/comment":
		return 449
	case "motion_comment/id":
		return 450
	case "motion_comment/meeting_id":
		return 451
	case "motion_comment/motion_id":
		return 452
	case "motion_comment/section_id":
		return 453
	case "motion_comment_section/comment_ids":
		return 454
	case "motion_comment_section/id":
		return 455
	case "motion_comment_section/meeting_id":
		return 456
	case "motion_comment_section/name":
		return 457
	case "motion_comment_section/read_group_ids":
		return 458
	case "motion_comment_section/sequential_number":
		return 459
	case "motion_comment_section/submitter_can_write":
		return 460
	case "motion_comment_section/weight":
		return 461
	case "motion_comment_section/write_group_ids":
		return 462
	case "motion_state/allow_create_poll":
		return 463
	case "motion_state/allow_motion_forwarding":
		return 464
	case "motion_state/allow_submitter_edit":
		return 465
	case "motion_state/allow_support":
		return 466
	case "motion_state/css_class":
		return 467
	case "motion_state/first_state_of_workflow_id":
		return 468
	case "motion_state/id":
		return 469
	case "motion_state/meeting_id":
		return 470
	case "motion_state/merge_amendment_into_final":
		return 471
	case "motion_state/motion_ids":
		return 472
	case "motion_state/motion_recommendation_ids":
		return 473
	case "motion_state/name":
		return 474
	case "motion_state/next_state_ids":
		return 475
	case "motion_state/previous_state_ids":
		return 476
	case "motion_state/recommendation_label":
		return 477
	case "motion_state/restrictions":
		return 478
	case "motion_state/set_number":
		return 479
	case "motion_state/set_workflow_timestamp":
		return 480
	case "motion_state/show_recommendation_extension_field":
		return 481
	case "motion_state/show_state_extension_field":
		return 482
	case "motion_state/submitter_withdraw_back_ids":
		return 483
	case "motion_state/submitter_withdraw_state_id":
		return 484
	case "motion_state/weight":
		return 485
	case "motion_state/workflow_id":
		return 486
	case "motion_statute_paragraph/id":
		return 487
	case "motion_statute_paragraph/meeting_id":
		return 488
	case "motion_statute_paragraph/motion_ids":
		return 489
	case "motion_statute_paragraph/sequential_number":
		return 490
	case "motion_statute_paragraph/text":
		return 491
	case "motion_statute_paragraph/title":
		return 492
	case "motion_statute_paragraph/weight":
		return 493
	case "motion_submitter/id":
		return 494
	case "motion_submitter/meeting_id":
		return 495
	case "motion_submitter/meeting_user_id":
		return 496
	case "motion_submitter/motion_id":
		return 497
	case "motion_submitter/weight":
		return 498
	case "motion_workflow/default_amendment_workflow_meeting_id":
		return 499
	case "motion_workflow/default_statute_amendment_workflow_meeting_id":
		return 500
	case "motion_workflow/default_workflow_meeting_id":
		return 501
	case "motion_workflow/first_state_id":
		return 502
	case "motion_workflow/id":
		return 503
	case "motion_workflow/meeting_id":
		return 504
	case "motion_workflow/name":
		return 505
	case "motion_workflow/sequential_number":
		return 506
	case "motion_workflow/state_ids":
		return 507
	case "option/abstain":
		return 508
	case "option/content_object_id":
		return 509
	case "option/id":
		return 510
	case "option/meeting_id":
		return 511
	case "option/no":
		return 512
	case "option/poll_id":
		return 513
	case "option/text":
		return 514
	case "option/used_as_global_option_in_poll_id":
		return 515
	case "option/vote_ids":
		return 516
	case "option/weight":
		return 517
	case "option/yes":
		return 518
	case "organization/active_meeting_ids":
		return 519
	case "organization/archived_meeting_ids":
		return 520
	case "organization/committee_ids":
		return 521
	case "organization/default_language":
		return 522
	case "organization/description":
		return 523
	case "organization/enable_chat":
		return 524
	case "organization/enable_electronic_voting":
		return 525
	case "organization/genders":
		return 526
	case "organization/id":
		return 527
	case "organization/legal_notice":
		return 528
	case "organization/limit_of_meetings":
		return 529
	case "organization/limit_of_users":
		return 530
	case "organization/login_text":
		return 531
	case "organization/mediafile_ids":
		return 532
	case "organization/name":
		return 533
	case "organization/organization_tag_ids":
		return 534
	case "organization/privacy_policy":
		return 535
	case "organization/reset_password_verbose_errors":
		return 536
	case "organization/saml_attr_mapping":
		return 537
	case "organization/saml_enabled":
		return 538
	case "organization/saml_login_button_text":
		return 539
	case "organization/saml_metadata_idp":
		return 540
	case "organization/saml_metadata_sp":
		return 541
	case "organization/saml_private_key":
		return 542
	case "organization/template_meeting_ids":
		return 543
	case "organization/theme_id":
		return 544
	case "organization/theme_ids":
		return 545
	case "organization/url":
		return 546
	case "organization/user_ids":
		return 547
	case "organization/users_email_body":
		return 548
	case "organization/users_email_replyto":
		return 549
	case "organization/users_email_sender":
		return 550
	case "organization/users_email_subject":
		return 551
	case "organization/vote_decrypt_public_main_key":
		return 552
	case "organization_tag/color":
		return 553
	case "organization_tag/id":
		return 554
	case "organization_tag/name":
		return 555
	case "organization_tag/organization_id":
		return 556
	case "organization_tag/tagged_ids":
		return 557
	case "personal_note/content_object_id":
		return 558
	case "personal_note/id":
		return 559
	case "personal_note/meeting_id":
		return 560
	case "personal_note/meeting_user_id":
		return 561
	case "personal_note/note":
		return 562
	case "personal_note/star":
		return 563
	case "point_of_order_category/id":
		return 564
	case "point_of_order_category/meeting_id":
		return 565
	case "point_of_order_category/rank":
		return 566
	case "point_of_order_category/speaker_ids":
		return 567
	case "point_of_order_category/text":
		return 568
	case "poll/backend":
		return 569
	case "poll/content_object_id":
		return 570
	case "poll/crypt_key":
		return 571
	case "poll/crypt_signature":
		return 572
	case "poll/description":
		return 573
	case "poll/entitled_group_ids":
		return 574
	case "poll/entitled_users_at_stop":
		return 575
	case "poll/global_abstain":
		return 576
	case "poll/global_no":
		return 577
	case "poll/global_option_id":
		return 578
	case "poll/global_yes":
		return 579
	case "poll/id":
		return 580
	case "poll/is_pseudoanonymized":
		return 581
	case "poll/max_votes_amount":
		return 582
	case "poll/max_votes_per_option":
		return 583
	case "poll/meeting_id":
		return 584
	case "poll/min_votes_amount":
		return 585
	case "poll/onehundred_percent_base":
		return 586
	case "poll/option_ids":
		return 587
	case "poll/pollmethod":
		return 588
	case "poll/projection_ids":
		return 589
	case "poll/sequential_number":
		return 590
	case "poll/state":
		return 591
	case "poll/title":
		return 592
	case "poll/type":
		return 593
	case "poll/vote_count":
		return 594
	case "poll/voted_ids":
		return 595
	case "poll/votes_raw":
		return 596
	case "poll/votes_signature":
		return 597
	case "poll/votescast":
		return 598
	case "poll/votesinvalid":
		return 599
	case "poll/votesvalid":
		return 600
	case "poll_candidate/id":
		return 601
	case "poll_candidate/meeting_id":
		return 602
	case "poll_candidate/poll_candidate_list_id":
		return 603
	case "poll_candidate/user_id":
		return 604
	case "poll_candidate/weight":
		return 605
	case "poll_candidate_list/id":
		return 606
	case "poll_candidate_list/meeting_id":
		return 607
	case "poll_candidate_list/option_id":
		return 608
	case "poll_candidate_list/poll_candidate_ids":
		return 609
	case "projection/content":
		return 610
	case "projection/content_object_id":
		return 611
	case "projection/current_projector_id":
		return 612
	case "projection/history_projector_id":
		return 613
	case "projection/id":
		return 614
	case "projection/meeting_id":
		return 615
	case "projection/options":
		return 616
	case "projection/preview_projector_id":
		return 617
	case "projection/stable":
		return 618
	case "projection/type":
		return 619
	case "projection/weight":
		return 620
	case "projector/aspect_ratio_denominator":
		return 621
	case "projector/aspect_ratio_numerator":
		return 622
	case "projector/background_color":
		return 623
	case "projector/chyron_background_color":
		return 624
	case "projector/chyron_font_color":
		return 625
	case "projector/color":
		return 626
	case "projector/current_projection_ids":
		return 627
	case "projector/header_background_color":
		return 628
	case "projector/header_font_color":
		return 629
	case "projector/header_h1_color":
		return 630
	case "projector/history_projection_ids":
		return 631
	case "projector/id":
		return 632
	case "projector/is_internal":
		return 633
	case "projector/meeting_id":
		return 634
	case "projector/name":
		return 635
	case "projector/preview_projection_ids":
		return 636
	case "projector/scale":
		return 637
	case "projector/scroll":
		return 638
	case "projector/sequential_number":
		return 639
	case "projector/show_clock":
		return 640
	case "projector/show_header_footer":
		return 641
	case "projector/show_logo":
		return 642
	case "projector/show_title":
		return 643
	case "projector/used_as_default_projector_for_agenda_item_list_in_meeting_id":
		return 644
	case "projector/used_as_default_projector_for_amendment_in_meeting_id":
		return 645
	case "projector/used_as_default_projector_for_assignment_in_meeting_id":
		return 646
	case "projector/used_as_default_projector_for_assignment_poll_in_meeting_id":
		return 647
	case "projector/used_as_default_projector_for_countdown_in_meeting_id":
		return 648
	case "projector/used_as_default_projector_for_current_list_of_speakers_in_meeting_id":
		return 649
	case "projector/used_as_default_projector_for_list_of_speakers_in_meeting_id":
		return 650
	case "projector/used_as_default_projector_for_mediafile_in_meeting_id":
		return 651
	case "projector/used_as_default_projector_for_message_in_meeting_id":
		return 652
	case "projector/used_as_default_projector_for_motion_block_in_meeting_id":
		return 653
	case "projector/used_as_default_projector_for_motion_in_meeting_id":
		return 654
	case "projector/used_as_default_projector_for_motion_poll_in_meeting_id":
		return 655
	case "projector/used_as_default_projector_for_poll_in_meeting_id":
		return 656
	case "projector/used_as_default_projector_for_topic_in_meeting_id":
		return 657
	case "projector/used_as_reference_projector_meeting_id":
		return 658
	case "projector/width":
		return 659
	case "projector_countdown/countdown_time":
		return 660
	case "projector_countdown/default_time":
		return 661
	case "projector_countdown/description":
		return 662
	case "projector_countdown/id":
		return 663
	case "projector_countdown/meeting_id":
		return 664
	case "projector_countdown/projection_ids":
		return 665
	case "projector_countdown/running":
		return 666
	case "projector_countdown/title":
		return 667
	case "projector_countdown/used_as_list_of_speakers_countdown_meeting_id":
		return 668
	case "projector_countdown/used_as_poll_countdown_meeting_id":
		return 669
	case "projector_message/id":
		return 670
	case "projector_message/meeting_id":
		return 671
	case "projector_message/message":
		return 672
	case "projector_message/projection_ids":
		return 673
	case "speaker/begin_time":
		return 674
	case "speaker/end_time":
		return 675
	case "speaker/id":
		return 676
	case "speaker/list_of_speakers_id":
		return 677
	case "speaker/meeting_id":
		return 678
	case "speaker/meeting_user_id":
		return 679
	case "speaker/note":
		return 680
	case "speaker/point_of_order":
		return 681
	case "speaker/point_of_order_category_id":
		return 682
	case "speaker/speech_state":
		return 683
	case "speaker/weight":
		return 684
	case "tag/id":
		return 685
	case "tag/meeting_id":
		return 686
	case "tag/name":
		return 687
	case "tag/tagged_ids":
		return 688
	case "theme/abstain":
		return 689
	case "theme/accent_100":
		return 690
	case "theme/accent_200":
		return 691
	case "theme/accent_300":
		return 692
	case "theme/accent_400":
		return 693
	case "theme/accent_50":
		return 694
	case "theme/accent_500":
		return 695
	case "theme/accent_600":
		return 696
	case "theme/accent_700":
		return 697
	case "theme/accent_800":
		return 698
	case "theme/accent_900":
		return 699
	case "theme/accent_a100":
		return 700
	case "theme/accent_a200":
		return 701
	case "theme/accent_a400":
		return 702
	case "theme/accent_a700":
		return 703
	case "theme/headbar":
		return 704
	case "theme/id":
		return 705
	case "theme/name":
		return 706
	case "theme/no":
		return 707
	case "theme/organization_id":
		return 708
	case "theme/primary_100":
		return 709
	case "theme/primary_200":
		return 710
	case "theme/primary_300":
		return 711
	case "theme/primary_400":
		return 712
	case "theme/primary_50":
		return 713
	case "theme/primary_500":
		return 714
	case "theme/primary_600":
		return 715
	case "theme/primary_700":
		return 716
	case "theme/primary_800":
		return 717
	case "theme/primary_900":
		return 718
	case "theme/primary_a100":
		return 719
	case "theme/primary_a200":
		return 720
	case "theme/primary_a400":
		return 721
	case "theme/primary_a700":
		return 722
	case "theme/theme_for_organization_id":
		return 723
	case "theme/warn_100":
		return 724
	case "theme/warn_200":
		return 725
	case "theme/warn_300":
		return 726
	case "theme/warn_400":
		return 727
	case "theme/warn_50":
		return 728
	case "theme/warn_500":
		return 729
	case "theme/warn_600":
		return 730
	case "theme/warn_700":
		return 731
	case "theme/warn_800":
		return 732
	case "theme/warn_900":
		return 733
	case "theme/warn_a100":
		return 734
	case "theme/warn_a200":
		return 735
	case "theme/warn_a400":
		return 736
	case "theme/warn_a700":
		return 737
	case "theme/yes":
		return 738
	case "topic/agenda_item_id":
		return 739
	case "topic/attachment_ids":
		return 740
	case "topic/id":
		return 741
	case "topic/list_of_speakers_id":
		return 742
	case "topic/meeting_id":
		return 743
	case "topic/poll_ids":
		return 744
	case "topic/projection_ids":
		return 745
	case "topic/sequential_number":
		return 746
	case "topic/text":
		return 747
	case "topic/title":
		return 748
	case "user/can_change_own_password":
		return 749
	case "user/committee_ids":
		return 750
	case "user/committee_management_ids":
		return 751
	case "user/default_number":
		return 752
	case "user/default_password":
		return 753
	case "user/default_structure_level":
		return 754
	case "user/default_vote_weight":
		return 755
	case "user/delegated_vote_ids":
		return 756
	case "user/email":
		return 757
	case "user/first_name":
		return 758
	case "user/forwarding_committee_ids":
		return 759
	case "user/gender":
		return 760
	case "user/id":
		return 761
	case "user/is_active":
		return 762
	case "user/is_demo_user":
		return 763
	case "user/is_physical_person":
		return 764
	case "user/is_present_in_meeting_ids":
		return 765
	case "user/last_email_sent":
		return 766
	case "user/last_login":
		return 767
	case "user/last_name":
		return 768
	case "user/meeting_ids":
		return 769
	case "user/meeting_user_ids":
		return 770
	case "user/option_ids":
		return 771
	case "user/organization_id":
		return 772
	case "user/organization_management_level":
		return 773
	case "user/password":
		return 774
	case "user/poll_candidate_ids":
		return 775
	case "user/poll_voted_ids":
		return 776
	case "user/pronoun":
		return 777
	case "user/saml_id":
		return 778
	case "user/title":
		return 779
	case "user/username":
		return 780
	case "user/vote_ids":
		return 781
	case "vote/delegated_user_id":
		return 782
	case "vote/id":
		return 783
	case "vote/meeting_id":
		return 784
	case "vote/option_id":
		return 785
	case "vote/user_id":
		return 786
	case "vote/user_token":
		return 787
	case "vote/value":
		return 788
	case "vote/weight":
		return 789
	default:
		return -1
	}
}

var collectionModeFields = [...]collectionMode{
	{"invalid", "mode"},
	{"action_worker", "A"},
	{"agenda_item", "A"},
	{"agenda_item", "B"},
	{"agenda_item", "C"},
	{"assignment", "A"},
	{"assignment_candidate", "A"},
	{"chat_group", "A"},
	{"chat_message", "A"},
	{"committee", "A"},
	{"committee", "B"},
	{"group", "A"},
	{"import_preview", "A"},
	{"list_of_speakers", "A"},
	{"mediafile", "A"},
	{"meeting", "A"},
	{"meeting", "B"},
	{"meeting", "C"},
	{"meeting_user", "A"},
	{"meeting_user", "B"},
	{"meeting_user", "D"},
	{"motion", "A"},
	{"motion", "C"},
	{"motion", "D"},
	{"motion_block", "A"},
	{"motion_category", "A"},
	{"motion_change_recommendation", "A"},
	{"motion_comment", "A"},
	{"motion_comment_section", "A"},
	{"motion_state", "A"},
	{"motion_statute_paragraph", "A"},
	{"motion_submitter", "A"},
	{"motion_workflow", "A"},
	{"option", "A"},
	{"option", "B"},
	{"organization", "A"},
	{"organization", "B"},
	{"organization", "C"},
	{"organization_tag", "A"},
	{"personal_note", "A"},
	{"point_of_order_category", "A"},
	{"poll", "A"},
	{"poll", "B"},
	{"poll", "C"},
	{"poll", "D"},
	{"poll", "MANAGE"},
	{"poll_candidate", "A"},
	{"poll_candidate_list", "A"},
	{"projection", "A"},
	{"projector", "A"},
	{"projector_countdown", "A"},
	{"projector_message", "A"},
	{"speaker", "A"},
	{"tag", "A"},
	{"theme", "A"},
	{"topic", "A"},
	{"user", "A"},
	{"user", "D"},
	{"user", "E"},
	{"user", "F"},
	{"user", "G"},
	{"user", "H"},
	{"vote", "A"},
	{"vote", "B"},
}

func collectionModeToID(cf string) int {
	switch cf {
	case "action_worker/A":
		return 1
	case "agenda_item/A":
		return 2
	case "agenda_item/B":
		return 3
	case "agenda_item/C":
		return 4
	case "assignment/A":
		return 5
	case "assignment_candidate/A":
		return 6
	case "chat_group/A":
		return 7
	case "chat_message/A":
		return 8
	case "committee/A":
		return 9
	case "committee/B":
		return 10
	case "group/A":
		return 11
	case "import_preview/A":
		return 12
	case "list_of_speakers/A":
		return 13
	case "mediafile/A":
		return 14
	case "meeting/A":
		return 15
	case "meeting/B":
		return 16
	case "meeting/C":
		return 17
	case "meeting_user/A":
		return 18
	case "meeting_user/B":
		return 19
	case "meeting_user/D":
		return 20
	case "motion/A":
		return 21
	case "motion/C":
		return 22
	case "motion/D":
		return 23
	case "motion_block/A":
		return 24
	case "motion_category/A":
		return 25
	case "motion_change_recommendation/A":
		return 26
	case "motion_comment/A":
		return 27
	case "motion_comment_section/A":
		return 28
	case "motion_state/A":
		return 29
	case "motion_statute_paragraph/A":
		return 30
	case "motion_submitter/A":
		return 31
	case "motion_workflow/A":
		return 32
	case "option/A":
		return 33
	case "option/B":
		return 34
	case "organization/A":
		return 35
	case "organization/B":
		return 36
	case "organization/C":
		return 37
	case "organization_tag/A":
		return 38
	case "personal_note/A":
		return 39
	case "point_of_order_category/A":
		return 40
	case "poll/A":
		return 41
	case "poll/B":
		return 42
	case "poll/C":
		return 43
	case "poll/D":
		return 44
	case "poll/MANAGE":
		return 45
	case "poll_candidate/A":
		return 46
	case "poll_candidate_list/A":
		return 47
	case "projection/A":
		return 48
	case "projector/A":
		return 49
	case "projector_countdown/A":
		return 50
	case "projector_message/A":
		return 51
	case "speaker/A":
		return 52
	case "tag/A":
		return 53
	case "theme/A":
		return 54
	case "topic/A":
		return 55
	case "user/A":
		return 56
	case "user/D":
		return 57
	case "user/E":
		return 58
	case "user/F":
		return 59
	case "user/G":
		return 60
	case "user/H":
		return 61
	case "vote/A":
		return 62
	case "vote/B":
		return 63
	default:
		return -1
	}
}

var collectionFieldToMode = [...]int{
	-1,
	-1,
	1,
	1,
	1,
	1,
	1,
	1,
	2,
	2,
	4,
	2,
	3,
	2,
	2,
	2,
	2,
	2,
	2,
	2,
	2,
	2,
	2,
	2,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	5,
	6,
	6,
	6,
	6,
	6,
	7,
	7,
	7,
	7,
	7,
	7,
	7,
	8,
	8,
	8,
	8,
	8,
	8,
	9,
	9,
	9,
	10,
	10,
	9,
	10,
	9,
	9,
	9,
	9,
	10,
	9,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	11,
	12,
	12,
	12,
	12,
	12,
	13,
	13,
	13,
	13,
	13,
	13,
	13,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	15,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	15,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	15,
	16,
	15,
	16,
	15,
	15,
	16,
	16,
	16,
	15,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	15,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	15,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	16,
	17,
	17,
	18,
	18,
	18,
	20,
	18,
	18,
	18,
	18,
	18,
	19,
	18,
	18,
	18,
	18,
	18,
	18,
	18,
	22,
	21,
	21,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	21,
	21,
	21,
	22,
	22,
	22,
	21,
	22,
	22,
	23,
	22,
	21,
	21,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	22,
	24,
	24,
	24,
	24,
	24,
	24,
	24,
	24,
	24,
	25,
	25,
	25,
	25,
	25,
	25,
	25,
	25,
	25,
	25,
	26,
	26,
	26,
	26,
	26,
	26,
	26,
	26,
	26,
	26,
	26,
	27,
	27,
	27,
	27,
	27,
	28,
	28,
	28,
	28,
	28,
	28,
	28,
	28,
	28,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	29,
	30,
	30,
	30,
	30,
	30,
	30,
	30,
	31,
	31,
	31,
	31,
	31,
	32,
	32,
	32,
	32,
	32,
	32,
	32,
	32,
	32,
	34,
	33,
	33,
	33,
	34,
	33,
	33,
	33,
	33,
	33,
	34,
	36,
	36,
	36,
	35,
	35,
	36,
	36,
	35,
	35,
	35,
	36,
	36,
	35,
	35,
	35,
	36,
	35,
	36,
	35,
	35,
	35,
	35,
	35,
	35,
	35,
	35,
	35,
	35,
	37,
	35,
	35,
	35,
	35,
	35,
	38,
	38,
	38,
	38,
	38,
	39,
	39,
	39,
	39,
	39,
	39,
	40,
	40,
	40,
	40,
	40,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	41,
	43,
	41,
	42,
	42,
	44,
	42,
	42,
	46,
	46,
	46,
	46,
	46,
	47,
	47,
	47,
	47,
	48,
	48,
	48,
	48,
	48,
	48,
	48,
	48,
	48,
	48,
	48,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	49,
	50,
	50,
	50,
	50,
	50,
	50,
	50,
	50,
	50,
	50,
	51,
	51,
	51,
	51,
	52,
	52,
	52,
	52,
	52,
	52,
	52,
	52,
	52,
	52,
	52,
	53,
	53,
	53,
	53,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	54,
	55,
	55,
	55,
	55,
	55,
	55,
	55,
	55,
	55,
	55,
	57,
	58,
	58,
	56,
	61,
	56,
	56,
	56,
	58,
	56,
	58,
	56,
	56,
	57,
	56,
	56,
	56,
	57,
	56,
	56,
	58,
	56,
	56,
	59,
	58,
	60,
	56,
	56,
	56,
	56,
	56,
	56,
	56,
	62,
	62,
	62,
	62,
	62,
	63,
	62,
	62,
}
