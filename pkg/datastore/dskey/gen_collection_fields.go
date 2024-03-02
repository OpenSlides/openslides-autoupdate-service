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
	{"agenda_item", "moderator_notes"},
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
	{"list_of_speakers", "structure_level_list_of_speakers_ids"},
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
	{"meeting", "list_of_speakers_allow_multiple_speakers"},
	{"meeting", "list_of_speakers_amount_last_on_projector"},
	{"meeting", "list_of_speakers_amount_next_on_projector"},
	{"meeting", "list_of_speakers_can_create_point_of_order_for_others"},
	{"meeting", "list_of_speakers_can_set_contribution_self"},
	{"meeting", "list_of_speakers_closing_disables_point_of_order"},
	{"meeting", "list_of_speakers_countdown_id"},
	{"meeting", "list_of_speakers_couple_countdown"},
	{"meeting", "list_of_speakers_default_structure_level_time"},
	{"meeting", "list_of_speakers_enable_interposed_question"},
	{"meeting", "list_of_speakers_enable_point_of_order_categories"},
	{"meeting", "list_of_speakers_enable_point_of_order_speakers"},
	{"meeting", "list_of_speakers_enable_pro_contra_speech"},
	{"meeting", "list_of_speakers_ids"},
	{"meeting", "list_of_speakers_initially_closed"},
	{"meeting", "list_of_speakers_intervention_time"},
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
	{"meeting", "motion_editor_ids"},
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
	{"meeting", "motion_working_group_speaker_ids"},
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
	{"meeting", "motions_enable_editor"},
	{"meeting", "motions_enable_reason_on_projector"},
	{"meeting", "motions_enable_recommendation_on_projector"},
	{"meeting", "motions_enable_sidebox_on_projector"},
	{"meeting", "motions_enable_text_on_projector"},
	{"meeting", "motions_enable_working_group_speaker"},
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
	{"meeting", "structure_level_ids"},
	{"meeting", "structure_level_list_of_speakers_ids"},
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
	{"meeting_user", "motion_editor_ids"},
	{"meeting_user", "motion_submitter_ids"},
	{"meeting_user", "motion_working_group_speaker_ids"},
	{"meeting_user", "number"},
	{"meeting_user", "personal_note_ids"},
	{"meeting_user", "speaker_ids"},
	{"meeting_user", "structure_level_ids"},
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
	{"motion", "editor_ids"},
	{"motion", "forwarded"},
	{"motion", "id"},
	{"motion", "identical_motion_ids"},
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
	{"motion", "text_hash"},
	{"motion", "title"},
	{"motion", "workflow_timestamp"},
	{"motion", "working_group_speaker_ids"},
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
	{"motion_editor", "id"},
	{"motion_editor", "meeting_id"},
	{"motion_editor", "meeting_user_id"},
	{"motion_editor", "motion_id"},
	{"motion_editor", "weight"},
	{"motion_state", "allow_create_poll"},
	{"motion_state", "allow_motion_forwarding"},
	{"motion_state", "allow_submitter_edit"},
	{"motion_state", "allow_support"},
	{"motion_state", "css_class"},
	{"motion_state", "first_state_of_workflow_id"},
	{"motion_state", "id"},
	{"motion_state", "is_internal"},
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
	{"motion_working_group_speaker", "id"},
	{"motion_working_group_speaker", "meeting_id"},
	{"motion_working_group_speaker", "meeting_user_id"},
	{"motion_working_group_speaker", "motion_id"},
	{"motion_working_group_speaker", "weight"},
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
	{"speaker", "pause_time"},
	{"speaker", "point_of_order"},
	{"speaker", "point_of_order_category_id"},
	{"speaker", "speech_state"},
	{"speaker", "structure_level_list_of_speakers_id"},
	{"speaker", "total_pause"},
	{"speaker", "unpause_time"},
	{"speaker", "weight"},
	{"structure_level", "color"},
	{"structure_level", "default_time"},
	{"structure_level", "id"},
	{"structure_level", "meeting_id"},
	{"structure_level", "meeting_user_ids"},
	{"structure_level", "name"},
	{"structure_level", "structure_level_list_of_speakers_ids"},
	{"structure_level_list_of_speakers", "additional_time"},
	{"structure_level_list_of_speakers", "current_start_time"},
	{"structure_level_list_of_speakers", "id"},
	{"structure_level_list_of_speakers", "initial_time"},
	{"structure_level_list_of_speakers", "list_of_speakers_id"},
	{"structure_level_list_of_speakers", "meeting_id"},
	{"structure_level_list_of_speakers", "remaining_time"},
	{"structure_level_list_of_speakers", "speaker_ids"},
	{"structure_level_list_of_speakers", "structure_level_id"},
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
	{"user", "default_password"},
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
	case "agenda_item/moderator_notes":
		return 19
	case "agenda_item/parent_id":
		return 20
	case "agenda_item/projection_ids":
		return 21
	case "agenda_item/tag_ids":
		return 22
	case "agenda_item/type":
		return 23
	case "agenda_item/weight":
		return 24
	case "assignment/agenda_item_id":
		return 25
	case "assignment/attachment_ids":
		return 26
	case "assignment/candidate_ids":
		return 27
	case "assignment/default_poll_description":
		return 28
	case "assignment/description":
		return 29
	case "assignment/id":
		return 30
	case "assignment/list_of_speakers_id":
		return 31
	case "assignment/meeting_id":
		return 32
	case "assignment/number_poll_candidates":
		return 33
	case "assignment/open_posts":
		return 34
	case "assignment/phase":
		return 35
	case "assignment/poll_ids":
		return 36
	case "assignment/projection_ids":
		return 37
	case "assignment/sequential_number":
		return 38
	case "assignment/tag_ids":
		return 39
	case "assignment/title":
		return 40
	case "assignment_candidate/assignment_id":
		return 41
	case "assignment_candidate/id":
		return 42
	case "assignment_candidate/meeting_id":
		return 43
	case "assignment_candidate/meeting_user_id":
		return 44
	case "assignment_candidate/weight":
		return 45
	case "chat_group/chat_message_ids":
		return 46
	case "chat_group/id":
		return 47
	case "chat_group/meeting_id":
		return 48
	case "chat_group/name":
		return 49
	case "chat_group/read_group_ids":
		return 50
	case "chat_group/weight":
		return 51
	case "chat_group/write_group_ids":
		return 52
	case "chat_message/chat_group_id":
		return 53
	case "chat_message/content":
		return 54
	case "chat_message/created":
		return 55
	case "chat_message/id":
		return 56
	case "chat_message/meeting_id":
		return 57
	case "chat_message/meeting_user_id":
		return 58
	case "committee/default_meeting_id":
		return 59
	case "committee/description":
		return 60
	case "committee/external_id":
		return 61
	case "committee/forward_to_committee_ids":
		return 62
	case "committee/forwarding_user_id":
		return 63
	case "committee/id":
		return 64
	case "committee/manager_ids":
		return 65
	case "committee/meeting_ids":
		return 66
	case "committee/name":
		return 67
	case "committee/organization_id":
		return 68
	case "committee/organization_tag_ids":
		return 69
	case "committee/receive_forwardings_from_committee_ids":
		return 70
	case "committee/user_ids":
		return 71
	case "group/admin_group_for_meeting_id":
		return 72
	case "group/default_group_for_meeting_id":
		return 73
	case "group/external_id":
		return 74
	case "group/id":
		return 75
	case "group/mediafile_access_group_ids":
		return 76
	case "group/mediafile_inherited_access_group_ids":
		return 77
	case "group/meeting_id":
		return 78
	case "group/meeting_user_ids":
		return 79
	case "group/name":
		return 80
	case "group/permissions":
		return 81
	case "group/poll_ids":
		return 82
	case "group/read_chat_group_ids":
		return 83
	case "group/read_comment_section_ids":
		return 84
	case "group/used_as_assignment_poll_default_id":
		return 85
	case "group/used_as_motion_poll_default_id":
		return 86
	case "group/used_as_poll_default_id":
		return 87
	case "group/used_as_topic_poll_default_id":
		return 88
	case "group/weight":
		return 89
	case "group/write_chat_group_ids":
		return 90
	case "group/write_comment_section_ids":
		return 91
	case "import_preview/created":
		return 92
	case "import_preview/id":
		return 93
	case "import_preview/name":
		return 94
	case "import_preview/result":
		return 95
	case "import_preview/state":
		return 96
	case "list_of_speakers/closed":
		return 97
	case "list_of_speakers/content_object_id":
		return 98
	case "list_of_speakers/id":
		return 99
	case "list_of_speakers/meeting_id":
		return 100
	case "list_of_speakers/projection_ids":
		return 101
	case "list_of_speakers/sequential_number":
		return 102
	case "list_of_speakers/speaker_ids":
		return 103
	case "list_of_speakers/structure_level_list_of_speakers_ids":
		return 104
	case "mediafile/access_group_ids":
		return 105
	case "mediafile/attachment_ids":
		return 106
	case "mediafile/child_ids":
		return 107
	case "mediafile/create_timestamp":
		return 108
	case "mediafile/filename":
		return 109
	case "mediafile/filesize":
		return 110
	case "mediafile/id":
		return 111
	case "mediafile/inherited_access_group_ids":
		return 112
	case "mediafile/is_directory":
		return 113
	case "mediafile/is_public":
		return 114
	case "mediafile/list_of_speakers_id":
		return 115
	case "mediafile/mimetype":
		return 116
	case "mediafile/owner_id":
		return 117
	case "mediafile/parent_id":
		return 118
	case "mediafile/pdf_information":
		return 119
	case "mediafile/projection_ids":
		return 120
	case "mediafile/title":
		return 121
	case "mediafile/token":
		return 122
	case "mediafile/used_as_font_bold_in_meeting_id":
		return 123
	case "mediafile/used_as_font_bold_italic_in_meeting_id":
		return 124
	case "mediafile/used_as_font_chyron_speaker_name_in_meeting_id":
		return 125
	case "mediafile/used_as_font_italic_in_meeting_id":
		return 126
	case "mediafile/used_as_font_monospace_in_meeting_id":
		return 127
	case "mediafile/used_as_font_projector_h1_in_meeting_id":
		return 128
	case "mediafile/used_as_font_projector_h2_in_meeting_id":
		return 129
	case "mediafile/used_as_font_regular_in_meeting_id":
		return 130
	case "mediafile/used_as_logo_pdf_ballot_paper_in_meeting_id":
		return 131
	case "mediafile/used_as_logo_pdf_footer_l_in_meeting_id":
		return 132
	case "mediafile/used_as_logo_pdf_footer_r_in_meeting_id":
		return 133
	case "mediafile/used_as_logo_pdf_header_l_in_meeting_id":
		return 134
	case "mediafile/used_as_logo_pdf_header_r_in_meeting_id":
		return 135
	case "mediafile/used_as_logo_projector_header_in_meeting_id":
		return 136
	case "mediafile/used_as_logo_projector_main_in_meeting_id":
		return 137
	case "mediafile/used_as_logo_web_header_in_meeting_id":
		return 138
	case "meeting/admin_group_id":
		return 139
	case "meeting/agenda_enable_numbering":
		return 140
	case "meeting/agenda_item_creation":
		return 141
	case "meeting/agenda_item_ids":
		return 142
	case "meeting/agenda_new_items_default_visibility":
		return 143
	case "meeting/agenda_number_prefix":
		return 144
	case "meeting/agenda_numeral_system":
		return 145
	case "meeting/agenda_show_internal_items_on_projector":
		return 146
	case "meeting/agenda_show_subtitles":
		return 147
	case "meeting/all_projection_ids":
		return 148
	case "meeting/applause_enable":
		return 149
	case "meeting/applause_max_amount":
		return 150
	case "meeting/applause_min_amount":
		return 151
	case "meeting/applause_particle_image_url":
		return 152
	case "meeting/applause_show_level":
		return 153
	case "meeting/applause_timeout":
		return 154
	case "meeting/applause_type":
		return 155
	case "meeting/assignment_candidate_ids":
		return 156
	case "meeting/assignment_ids":
		return 157
	case "meeting/assignment_poll_add_candidates_to_list_of_speakers":
		return 158
	case "meeting/assignment_poll_ballot_paper_number":
		return 159
	case "meeting/assignment_poll_ballot_paper_selection":
		return 160
	case "meeting/assignment_poll_default_backend":
		return 161
	case "meeting/assignment_poll_default_group_ids":
		return 162
	case "meeting/assignment_poll_default_method":
		return 163
	case "meeting/assignment_poll_default_onehundred_percent_base":
		return 164
	case "meeting/assignment_poll_default_type":
		return 165
	case "meeting/assignment_poll_enable_max_votes_per_option":
		return 166
	case "meeting/assignment_poll_sort_poll_result_by_votes":
		return 167
	case "meeting/assignments_export_preamble":
		return 168
	case "meeting/assignments_export_title":
		return 169
	case "meeting/chat_group_ids":
		return 170
	case "meeting/chat_message_ids":
		return 171
	case "meeting/committee_id":
		return 172
	case "meeting/conference_auto_connect":
		return 173
	case "meeting/conference_auto_connect_next_speakers":
		return 174
	case "meeting/conference_enable_helpdesk":
		return 175
	case "meeting/conference_los_restriction":
		return 176
	case "meeting/conference_open_microphone":
		return 177
	case "meeting/conference_open_video":
		return 178
	case "meeting/conference_show":
		return 179
	case "meeting/conference_stream_poster_url":
		return 180
	case "meeting/conference_stream_url":
		return 181
	case "meeting/custom_translations":
		return 182
	case "meeting/default_group_id":
		return 183
	case "meeting/default_meeting_for_committee_id":
		return 184
	case "meeting/default_projector_agenda_item_list_ids":
		return 185
	case "meeting/default_projector_amendment_ids":
		return 186
	case "meeting/default_projector_assignment_ids":
		return 187
	case "meeting/default_projector_assignment_poll_ids":
		return 188
	case "meeting/default_projector_countdown_ids":
		return 189
	case "meeting/default_projector_current_list_of_speakers_ids":
		return 190
	case "meeting/default_projector_list_of_speakers_ids":
		return 191
	case "meeting/default_projector_mediafile_ids":
		return 192
	case "meeting/default_projector_message_ids":
		return 193
	case "meeting/default_projector_motion_block_ids":
		return 194
	case "meeting/default_projector_motion_ids":
		return 195
	case "meeting/default_projector_motion_poll_ids":
		return 196
	case "meeting/default_projector_poll_ids":
		return 197
	case "meeting/default_projector_topic_ids":
		return 198
	case "meeting/description":
		return 199
	case "meeting/enable_anonymous":
		return 200
	case "meeting/end_time":
		return 201
	case "meeting/export_csv_encoding":
		return 202
	case "meeting/export_csv_separator":
		return 203
	case "meeting/export_pdf_fontsize":
		return 204
	case "meeting/export_pdf_line_height":
		return 205
	case "meeting/export_pdf_page_margin_bottom":
		return 206
	case "meeting/export_pdf_page_margin_left":
		return 207
	case "meeting/export_pdf_page_margin_right":
		return 208
	case "meeting/export_pdf_page_margin_top":
		return 209
	case "meeting/export_pdf_pagenumber_alignment":
		return 210
	case "meeting/export_pdf_pagesize":
		return 211
	case "meeting/external_id":
		return 212
	case "meeting/font_bold_id":
		return 213
	case "meeting/font_bold_italic_id":
		return 214
	case "meeting/font_chyron_speaker_name_id":
		return 215
	case "meeting/font_italic_id":
		return 216
	case "meeting/font_monospace_id":
		return 217
	case "meeting/font_projector_h1_id":
		return 218
	case "meeting/font_projector_h2_id":
		return 219
	case "meeting/font_regular_id":
		return 220
	case "meeting/forwarded_motion_ids":
		return 221
	case "meeting/group_ids":
		return 222
	case "meeting/id":
		return 223
	case "meeting/imported_at":
		return 224
	case "meeting/is_active_in_organization_id":
		return 225
	case "meeting/is_archived_in_organization_id":
		return 226
	case "meeting/jitsi_domain":
		return 227
	case "meeting/jitsi_room_name":
		return 228
	case "meeting/jitsi_room_password":
		return 229
	case "meeting/language":
		return 230
	case "meeting/list_of_speakers_allow_multiple_speakers":
		return 231
	case "meeting/list_of_speakers_amount_last_on_projector":
		return 232
	case "meeting/list_of_speakers_amount_next_on_projector":
		return 233
	case "meeting/list_of_speakers_can_create_point_of_order_for_others":
		return 234
	case "meeting/list_of_speakers_can_set_contribution_self":
		return 235
	case "meeting/list_of_speakers_closing_disables_point_of_order":
		return 236
	case "meeting/list_of_speakers_countdown_id":
		return 237
	case "meeting/list_of_speakers_couple_countdown":
		return 238
	case "meeting/list_of_speakers_default_structure_level_time":
		return 239
	case "meeting/list_of_speakers_enable_interposed_question":
		return 240
	case "meeting/list_of_speakers_enable_point_of_order_categories":
		return 241
	case "meeting/list_of_speakers_enable_point_of_order_speakers":
		return 242
	case "meeting/list_of_speakers_enable_pro_contra_speech":
		return 243
	case "meeting/list_of_speakers_ids":
		return 244
	case "meeting/list_of_speakers_initially_closed":
		return 245
	case "meeting/list_of_speakers_intervention_time":
		return 246
	case "meeting/list_of_speakers_present_users_only":
		return 247
	case "meeting/list_of_speakers_show_amount_of_speakers_on_slide":
		return 248
	case "meeting/list_of_speakers_show_first_contribution":
		return 249
	case "meeting/list_of_speakers_speaker_note_for_everyone":
		return 250
	case "meeting/location":
		return 251
	case "meeting/logo_pdf_ballot_paper_id":
		return 252
	case "meeting/logo_pdf_footer_l_id":
		return 253
	case "meeting/logo_pdf_footer_r_id":
		return 254
	case "meeting/logo_pdf_header_l_id":
		return 255
	case "meeting/logo_pdf_header_r_id":
		return 256
	case "meeting/logo_projector_header_id":
		return 257
	case "meeting/logo_projector_main_id":
		return 258
	case "meeting/logo_web_header_id":
		return 259
	case "meeting/mediafile_ids":
		return 260
	case "meeting/meeting_user_ids":
		return 261
	case "meeting/motion_block_ids":
		return 262
	case "meeting/motion_category_ids":
		return 263
	case "meeting/motion_change_recommendation_ids":
		return 264
	case "meeting/motion_comment_ids":
		return 265
	case "meeting/motion_comment_section_ids":
		return 266
	case "meeting/motion_editor_ids":
		return 267
	case "meeting/motion_ids":
		return 268
	case "meeting/motion_poll_ballot_paper_number":
		return 269
	case "meeting/motion_poll_ballot_paper_selection":
		return 270
	case "meeting/motion_poll_default_backend":
		return 271
	case "meeting/motion_poll_default_group_ids":
		return 272
	case "meeting/motion_poll_default_onehundred_percent_base":
		return 273
	case "meeting/motion_poll_default_type":
		return 274
	case "meeting/motion_state_ids":
		return 275
	case "meeting/motion_statute_paragraph_ids":
		return 276
	case "meeting/motion_submitter_ids":
		return 277
	case "meeting/motion_workflow_ids":
		return 278
	case "meeting/motion_working_group_speaker_ids":
		return 279
	case "meeting/motions_amendments_enabled":
		return 280
	case "meeting/motions_amendments_in_main_list":
		return 281
	case "meeting/motions_amendments_multiple_paragraphs":
		return 282
	case "meeting/motions_amendments_of_amendments":
		return 283
	case "meeting/motions_amendments_prefix":
		return 284
	case "meeting/motions_amendments_text_mode":
		return 285
	case "meeting/motions_block_slide_columns":
		return 286
	case "meeting/motions_default_amendment_workflow_id":
		return 287
	case "meeting/motions_default_line_numbering":
		return 288
	case "meeting/motions_default_sorting":
		return 289
	case "meeting/motions_default_statute_amendment_workflow_id":
		return 290
	case "meeting/motions_default_workflow_id":
		return 291
	case "meeting/motions_enable_editor":
		return 292
	case "meeting/motions_enable_reason_on_projector":
		return 293
	case "meeting/motions_enable_recommendation_on_projector":
		return 294
	case "meeting/motions_enable_sidebox_on_projector":
		return 295
	case "meeting/motions_enable_text_on_projector":
		return 296
	case "meeting/motions_enable_working_group_speaker":
		return 297
	case "meeting/motions_export_follow_recommendation":
		return 298
	case "meeting/motions_export_preamble":
		return 299
	case "meeting/motions_export_submitter_recommendation":
		return 300
	case "meeting/motions_export_title":
		return 301
	case "meeting/motions_line_length":
		return 302
	case "meeting/motions_number_min_digits":
		return 303
	case "meeting/motions_number_type":
		return 304
	case "meeting/motions_number_with_blank":
		return 305
	case "meeting/motions_preamble":
		return 306
	case "meeting/motions_reason_required":
		return 307
	case "meeting/motions_recommendation_text_mode":
		return 308
	case "meeting/motions_recommendations_by":
		return 309
	case "meeting/motions_show_referring_motions":
		return 310
	case "meeting/motions_show_sequential_number":
		return 311
	case "meeting/motions_statute_recommendations_by":
		return 312
	case "meeting/motions_statutes_enabled":
		return 313
	case "meeting/motions_supporters_min_amount":
		return 314
	case "meeting/name":
		return 315
	case "meeting/option_ids":
		return 316
	case "meeting/organization_tag_ids":
		return 317
	case "meeting/personal_note_ids":
		return 318
	case "meeting/point_of_order_category_ids":
		return 319
	case "meeting/poll_ballot_paper_number":
		return 320
	case "meeting/poll_ballot_paper_selection":
		return 321
	case "meeting/poll_candidate_ids":
		return 322
	case "meeting/poll_candidate_list_ids":
		return 323
	case "meeting/poll_countdown_id":
		return 324
	case "meeting/poll_couple_countdown":
		return 325
	case "meeting/poll_default_backend":
		return 326
	case "meeting/poll_default_group_ids":
		return 327
	case "meeting/poll_default_method":
		return 328
	case "meeting/poll_default_onehundred_percent_base":
		return 329
	case "meeting/poll_default_type":
		return 330
	case "meeting/poll_ids":
		return 331
	case "meeting/poll_sort_poll_result_by_votes":
		return 332
	case "meeting/present_user_ids":
		return 333
	case "meeting/projection_ids":
		return 334
	case "meeting/projector_countdown_default_time":
		return 335
	case "meeting/projector_countdown_ids":
		return 336
	case "meeting/projector_countdown_warning_time":
		return 337
	case "meeting/projector_ids":
		return 338
	case "meeting/projector_message_ids":
		return 339
	case "meeting/reference_projector_id":
		return 340
	case "meeting/speaker_ids":
		return 341
	case "meeting/start_time":
		return 342
	case "meeting/structure_level_ids":
		return 343
	case "meeting/structure_level_list_of_speakers_ids":
		return 344
	case "meeting/tag_ids":
		return 345
	case "meeting/template_for_organization_id":
		return 346
	case "meeting/topic_ids":
		return 347
	case "meeting/topic_poll_default_group_ids":
		return 348
	case "meeting/user_ids":
		return 349
	case "meeting/users_allow_self_set_present":
		return 350
	case "meeting/users_email_body":
		return 351
	case "meeting/users_email_replyto":
		return 352
	case "meeting/users_email_sender":
		return 353
	case "meeting/users_email_subject":
		return 354
	case "meeting/users_enable_presence_view":
		return 355
	case "meeting/users_enable_vote_delegations":
		return 356
	case "meeting/users_enable_vote_weight":
		return 357
	case "meeting/users_pdf_welcometext":
		return 358
	case "meeting/users_pdf_welcometitle":
		return 359
	case "meeting/users_pdf_wlan_encryption":
		return 360
	case "meeting/users_pdf_wlan_password":
		return 361
	case "meeting/users_pdf_wlan_ssid":
		return 362
	case "meeting/vote_ids":
		return 363
	case "meeting/welcome_text":
		return 364
	case "meeting/welcome_title":
		return 365
	case "meeting_user/about_me":
		return 366
	case "meeting_user/assignment_candidate_ids":
		return 367
	case "meeting_user/chat_message_ids":
		return 368
	case "meeting_user/comment":
		return 369
	case "meeting_user/group_ids":
		return 370
	case "meeting_user/id":
		return 371
	case "meeting_user/meeting_id":
		return 372
	case "meeting_user/motion_editor_ids":
		return 373
	case "meeting_user/motion_submitter_ids":
		return 374
	case "meeting_user/motion_working_group_speaker_ids":
		return 375
	case "meeting_user/number":
		return 376
	case "meeting_user/personal_note_ids":
		return 377
	case "meeting_user/speaker_ids":
		return 378
	case "meeting_user/structure_level_ids":
		return 379
	case "meeting_user/supported_motion_ids":
		return 380
	case "meeting_user/user_id":
		return 381
	case "meeting_user/vote_delegated_to_id":
		return 382
	case "meeting_user/vote_delegations_from_ids":
		return 383
	case "meeting_user/vote_weight":
		return 384
	case "motion/agenda_item_id":
		return 385
	case "motion/all_derived_motion_ids":
		return 386
	case "motion/all_origin_ids":
		return 387
	case "motion/amendment_ids":
		return 388
	case "motion/amendment_paragraphs":
		return 389
	case "motion/attachment_ids":
		return 390
	case "motion/block_id":
		return 391
	case "motion/category_id":
		return 392
	case "motion/category_weight":
		return 393
	case "motion/change_recommendation_ids":
		return 394
	case "motion/comment_ids":
		return 395
	case "motion/created":
		return 396
	case "motion/derived_motion_ids":
		return 397
	case "motion/editor_ids":
		return 398
	case "motion/forwarded":
		return 399
	case "motion/id":
		return 400
	case "motion/identical_motion_ids":
		return 401
	case "motion/last_modified":
		return 402
	case "motion/lead_motion_id":
		return 403
	case "motion/list_of_speakers_id":
		return 404
	case "motion/meeting_id":
		return 405
	case "motion/modified_final_version":
		return 406
	case "motion/number":
		return 407
	case "motion/number_value":
		return 408
	case "motion/option_ids":
		return 409
	case "motion/origin_id":
		return 410
	case "motion/origin_meeting_id":
		return 411
	case "motion/personal_note_ids":
		return 412
	case "motion/poll_ids":
		return 413
	case "motion/projection_ids":
		return 414
	case "motion/reason":
		return 415
	case "motion/recommendation_extension":
		return 416
	case "motion/recommendation_extension_reference_ids":
		return 417
	case "motion/recommendation_id":
		return 418
	case "motion/referenced_in_motion_recommendation_extension_ids":
		return 419
	case "motion/referenced_in_motion_state_extension_ids":
		return 420
	case "motion/sequential_number":
		return 421
	case "motion/sort_child_ids":
		return 422
	case "motion/sort_parent_id":
		return 423
	case "motion/sort_weight":
		return 424
	case "motion/start_line_number":
		return 425
	case "motion/state_extension":
		return 426
	case "motion/state_extension_reference_ids":
		return 427
	case "motion/state_id":
		return 428
	case "motion/statute_paragraph_id":
		return 429
	case "motion/submitter_ids":
		return 430
	case "motion/supporter_meeting_user_ids":
		return 431
	case "motion/tag_ids":
		return 432
	case "motion/text":
		return 433
	case "motion/text_hash":
		return 434
	case "motion/title":
		return 435
	case "motion/workflow_timestamp":
		return 436
	case "motion/working_group_speaker_ids":
		return 437
	case "motion_block/agenda_item_id":
		return 438
	case "motion_block/id":
		return 439
	case "motion_block/internal":
		return 440
	case "motion_block/list_of_speakers_id":
		return 441
	case "motion_block/meeting_id":
		return 442
	case "motion_block/motion_ids":
		return 443
	case "motion_block/projection_ids":
		return 444
	case "motion_block/sequential_number":
		return 445
	case "motion_block/title":
		return 446
	case "motion_category/child_ids":
		return 447
	case "motion_category/id":
		return 448
	case "motion_category/level":
		return 449
	case "motion_category/meeting_id":
		return 450
	case "motion_category/motion_ids":
		return 451
	case "motion_category/name":
		return 452
	case "motion_category/parent_id":
		return 453
	case "motion_category/prefix":
		return 454
	case "motion_category/sequential_number":
		return 455
	case "motion_category/weight":
		return 456
	case "motion_change_recommendation/creation_time":
		return 457
	case "motion_change_recommendation/id":
		return 458
	case "motion_change_recommendation/internal":
		return 459
	case "motion_change_recommendation/line_from":
		return 460
	case "motion_change_recommendation/line_to":
		return 461
	case "motion_change_recommendation/meeting_id":
		return 462
	case "motion_change_recommendation/motion_id":
		return 463
	case "motion_change_recommendation/other_description":
		return 464
	case "motion_change_recommendation/rejected":
		return 465
	case "motion_change_recommendation/text":
		return 466
	case "motion_change_recommendation/type":
		return 467
	case "motion_comment/comment":
		return 468
	case "motion_comment/id":
		return 469
	case "motion_comment/meeting_id":
		return 470
	case "motion_comment/motion_id":
		return 471
	case "motion_comment/section_id":
		return 472
	case "motion_comment_section/comment_ids":
		return 473
	case "motion_comment_section/id":
		return 474
	case "motion_comment_section/meeting_id":
		return 475
	case "motion_comment_section/name":
		return 476
	case "motion_comment_section/read_group_ids":
		return 477
	case "motion_comment_section/sequential_number":
		return 478
	case "motion_comment_section/submitter_can_write":
		return 479
	case "motion_comment_section/weight":
		return 480
	case "motion_comment_section/write_group_ids":
		return 481
	case "motion_editor/id":
		return 482
	case "motion_editor/meeting_id":
		return 483
	case "motion_editor/meeting_user_id":
		return 484
	case "motion_editor/motion_id":
		return 485
	case "motion_editor/weight":
		return 486
	case "motion_state/allow_create_poll":
		return 487
	case "motion_state/allow_motion_forwarding":
		return 488
	case "motion_state/allow_submitter_edit":
		return 489
	case "motion_state/allow_support":
		return 490
	case "motion_state/css_class":
		return 491
	case "motion_state/first_state_of_workflow_id":
		return 492
	case "motion_state/id":
		return 493
	case "motion_state/is_internal":
		return 494
	case "motion_state/meeting_id":
		return 495
	case "motion_state/merge_amendment_into_final":
		return 496
	case "motion_state/motion_ids":
		return 497
	case "motion_state/motion_recommendation_ids":
		return 498
	case "motion_state/name":
		return 499
	case "motion_state/next_state_ids":
		return 500
	case "motion_state/previous_state_ids":
		return 501
	case "motion_state/recommendation_label":
		return 502
	case "motion_state/restrictions":
		return 503
	case "motion_state/set_number":
		return 504
	case "motion_state/set_workflow_timestamp":
		return 505
	case "motion_state/show_recommendation_extension_field":
		return 506
	case "motion_state/show_state_extension_field":
		return 507
	case "motion_state/submitter_withdraw_back_ids":
		return 508
	case "motion_state/submitter_withdraw_state_id":
		return 509
	case "motion_state/weight":
		return 510
	case "motion_state/workflow_id":
		return 511
	case "motion_statute_paragraph/id":
		return 512
	case "motion_statute_paragraph/meeting_id":
		return 513
	case "motion_statute_paragraph/motion_ids":
		return 514
	case "motion_statute_paragraph/sequential_number":
		return 515
	case "motion_statute_paragraph/text":
		return 516
	case "motion_statute_paragraph/title":
		return 517
	case "motion_statute_paragraph/weight":
		return 518
	case "motion_submitter/id":
		return 519
	case "motion_submitter/meeting_id":
		return 520
	case "motion_submitter/meeting_user_id":
		return 521
	case "motion_submitter/motion_id":
		return 522
	case "motion_submitter/weight":
		return 523
	case "motion_workflow/default_amendment_workflow_meeting_id":
		return 524
	case "motion_workflow/default_statute_amendment_workflow_meeting_id":
		return 525
	case "motion_workflow/default_workflow_meeting_id":
		return 526
	case "motion_workflow/first_state_id":
		return 527
	case "motion_workflow/id":
		return 528
	case "motion_workflow/meeting_id":
		return 529
	case "motion_workflow/name":
		return 530
	case "motion_workflow/sequential_number":
		return 531
	case "motion_workflow/state_ids":
		return 532
	case "motion_working_group_speaker/id":
		return 533
	case "motion_working_group_speaker/meeting_id":
		return 534
	case "motion_working_group_speaker/meeting_user_id":
		return 535
	case "motion_working_group_speaker/motion_id":
		return 536
	case "motion_working_group_speaker/weight":
		return 537
	case "option/abstain":
		return 538
	case "option/content_object_id":
		return 539
	case "option/id":
		return 540
	case "option/meeting_id":
		return 541
	case "option/no":
		return 542
	case "option/poll_id":
		return 543
	case "option/text":
		return 544
	case "option/used_as_global_option_in_poll_id":
		return 545
	case "option/vote_ids":
		return 546
	case "option/weight":
		return 547
	case "option/yes":
		return 548
	case "organization/active_meeting_ids":
		return 549
	case "organization/archived_meeting_ids":
		return 550
	case "organization/committee_ids":
		return 551
	case "organization/default_language":
		return 552
	case "organization/description":
		return 553
	case "organization/enable_chat":
		return 554
	case "organization/enable_electronic_voting":
		return 555
	case "organization/genders":
		return 556
	case "organization/id":
		return 557
	case "organization/legal_notice":
		return 558
	case "organization/limit_of_meetings":
		return 559
	case "organization/limit_of_users":
		return 560
	case "organization/login_text":
		return 561
	case "organization/mediafile_ids":
		return 562
	case "organization/name":
		return 563
	case "organization/organization_tag_ids":
		return 564
	case "organization/privacy_policy":
		return 565
	case "organization/reset_password_verbose_errors":
		return 566
	case "organization/saml_attr_mapping":
		return 567
	case "organization/saml_enabled":
		return 568
	case "organization/saml_login_button_text":
		return 569
	case "organization/saml_metadata_idp":
		return 570
	case "organization/saml_metadata_sp":
		return 571
	case "organization/saml_private_key":
		return 572
	case "organization/template_meeting_ids":
		return 573
	case "organization/theme_id":
		return 574
	case "organization/theme_ids":
		return 575
	case "organization/url":
		return 576
	case "organization/user_ids":
		return 577
	case "organization/users_email_body":
		return 578
	case "organization/users_email_replyto":
		return 579
	case "organization/users_email_sender":
		return 580
	case "organization/users_email_subject":
		return 581
	case "organization/vote_decrypt_public_main_key":
		return 582
	case "organization_tag/color":
		return 583
	case "organization_tag/id":
		return 584
	case "organization_tag/name":
		return 585
	case "organization_tag/organization_id":
		return 586
	case "organization_tag/tagged_ids":
		return 587
	case "personal_note/content_object_id":
		return 588
	case "personal_note/id":
		return 589
	case "personal_note/meeting_id":
		return 590
	case "personal_note/meeting_user_id":
		return 591
	case "personal_note/note":
		return 592
	case "personal_note/star":
		return 593
	case "point_of_order_category/id":
		return 594
	case "point_of_order_category/meeting_id":
		return 595
	case "point_of_order_category/rank":
		return 596
	case "point_of_order_category/speaker_ids":
		return 597
	case "point_of_order_category/text":
		return 598
	case "poll/backend":
		return 599
	case "poll/content_object_id":
		return 600
	case "poll/crypt_key":
		return 601
	case "poll/crypt_signature":
		return 602
	case "poll/description":
		return 603
	case "poll/entitled_group_ids":
		return 604
	case "poll/entitled_users_at_stop":
		return 605
	case "poll/global_abstain":
		return 606
	case "poll/global_no":
		return 607
	case "poll/global_option_id":
		return 608
	case "poll/global_yes":
		return 609
	case "poll/id":
		return 610
	case "poll/is_pseudoanonymized":
		return 611
	case "poll/max_votes_amount":
		return 612
	case "poll/max_votes_per_option":
		return 613
	case "poll/meeting_id":
		return 614
	case "poll/min_votes_amount":
		return 615
	case "poll/onehundred_percent_base":
		return 616
	case "poll/option_ids":
		return 617
	case "poll/pollmethod":
		return 618
	case "poll/projection_ids":
		return 619
	case "poll/sequential_number":
		return 620
	case "poll/state":
		return 621
	case "poll/title":
		return 622
	case "poll/type":
		return 623
	case "poll/vote_count":
		return 624
	case "poll/voted_ids":
		return 625
	case "poll/votes_raw":
		return 626
	case "poll/votes_signature":
		return 627
	case "poll/votescast":
		return 628
	case "poll/votesinvalid":
		return 629
	case "poll/votesvalid":
		return 630
	case "poll_candidate/id":
		return 631
	case "poll_candidate/meeting_id":
		return 632
	case "poll_candidate/poll_candidate_list_id":
		return 633
	case "poll_candidate/user_id":
		return 634
	case "poll_candidate/weight":
		return 635
	case "poll_candidate_list/id":
		return 636
	case "poll_candidate_list/meeting_id":
		return 637
	case "poll_candidate_list/option_id":
		return 638
	case "poll_candidate_list/poll_candidate_ids":
		return 639
	case "projection/content":
		return 640
	case "projection/content_object_id":
		return 641
	case "projection/current_projector_id":
		return 642
	case "projection/history_projector_id":
		return 643
	case "projection/id":
		return 644
	case "projection/meeting_id":
		return 645
	case "projection/options":
		return 646
	case "projection/preview_projector_id":
		return 647
	case "projection/stable":
		return 648
	case "projection/type":
		return 649
	case "projection/weight":
		return 650
	case "projector/aspect_ratio_denominator":
		return 651
	case "projector/aspect_ratio_numerator":
		return 652
	case "projector/background_color":
		return 653
	case "projector/chyron_background_color":
		return 654
	case "projector/chyron_font_color":
		return 655
	case "projector/color":
		return 656
	case "projector/current_projection_ids":
		return 657
	case "projector/header_background_color":
		return 658
	case "projector/header_font_color":
		return 659
	case "projector/header_h1_color":
		return 660
	case "projector/history_projection_ids":
		return 661
	case "projector/id":
		return 662
	case "projector/is_internal":
		return 663
	case "projector/meeting_id":
		return 664
	case "projector/name":
		return 665
	case "projector/preview_projection_ids":
		return 666
	case "projector/scale":
		return 667
	case "projector/scroll":
		return 668
	case "projector/sequential_number":
		return 669
	case "projector/show_clock":
		return 670
	case "projector/show_header_footer":
		return 671
	case "projector/show_logo":
		return 672
	case "projector/show_title":
		return 673
	case "projector/used_as_default_projector_for_agenda_item_list_in_meeting_id":
		return 674
	case "projector/used_as_default_projector_for_amendment_in_meeting_id":
		return 675
	case "projector/used_as_default_projector_for_assignment_in_meeting_id":
		return 676
	case "projector/used_as_default_projector_for_assignment_poll_in_meeting_id":
		return 677
	case "projector/used_as_default_projector_for_countdown_in_meeting_id":
		return 678
	case "projector/used_as_default_projector_for_current_list_of_speakers_in_meeting_id":
		return 679
	case "projector/used_as_default_projector_for_list_of_speakers_in_meeting_id":
		return 680
	case "projector/used_as_default_projector_for_mediafile_in_meeting_id":
		return 681
	case "projector/used_as_default_projector_for_message_in_meeting_id":
		return 682
	case "projector/used_as_default_projector_for_motion_block_in_meeting_id":
		return 683
	case "projector/used_as_default_projector_for_motion_in_meeting_id":
		return 684
	case "projector/used_as_default_projector_for_motion_poll_in_meeting_id":
		return 685
	case "projector/used_as_default_projector_for_poll_in_meeting_id":
		return 686
	case "projector/used_as_default_projector_for_topic_in_meeting_id":
		return 687
	case "projector/used_as_reference_projector_meeting_id":
		return 688
	case "projector/width":
		return 689
	case "projector_countdown/countdown_time":
		return 690
	case "projector_countdown/default_time":
		return 691
	case "projector_countdown/description":
		return 692
	case "projector_countdown/id":
		return 693
	case "projector_countdown/meeting_id":
		return 694
	case "projector_countdown/projection_ids":
		return 695
	case "projector_countdown/running":
		return 696
	case "projector_countdown/title":
		return 697
	case "projector_countdown/used_as_list_of_speakers_countdown_meeting_id":
		return 698
	case "projector_countdown/used_as_poll_countdown_meeting_id":
		return 699
	case "projector_message/id":
		return 700
	case "projector_message/meeting_id":
		return 701
	case "projector_message/message":
		return 702
	case "projector_message/projection_ids":
		return 703
	case "speaker/begin_time":
		return 704
	case "speaker/end_time":
		return 705
	case "speaker/id":
		return 706
	case "speaker/list_of_speakers_id":
		return 707
	case "speaker/meeting_id":
		return 708
	case "speaker/meeting_user_id":
		return 709
	case "speaker/note":
		return 710
	case "speaker/pause_time":
		return 711
	case "speaker/point_of_order":
		return 712
	case "speaker/point_of_order_category_id":
		return 713
	case "speaker/speech_state":
		return 714
	case "speaker/structure_level_list_of_speakers_id":
		return 715
	case "speaker/total_pause":
		return 716
	case "speaker/unpause_time":
		return 717
	case "speaker/weight":
		return 718
	case "structure_level/color":
		return 719
	case "structure_level/default_time":
		return 720
	case "structure_level/id":
		return 721
	case "structure_level/meeting_id":
		return 722
	case "structure_level/meeting_user_ids":
		return 723
	case "structure_level/name":
		return 724
	case "structure_level/structure_level_list_of_speakers_ids":
		return 725
	case "structure_level_list_of_speakers/additional_time":
		return 726
	case "structure_level_list_of_speakers/current_start_time":
		return 727
	case "structure_level_list_of_speakers/id":
		return 728
	case "structure_level_list_of_speakers/initial_time":
		return 729
	case "structure_level_list_of_speakers/list_of_speakers_id":
		return 730
	case "structure_level_list_of_speakers/meeting_id":
		return 731
	case "structure_level_list_of_speakers/remaining_time":
		return 732
	case "structure_level_list_of_speakers/speaker_ids":
		return 733
	case "structure_level_list_of_speakers/structure_level_id":
		return 734
	case "tag/id":
		return 735
	case "tag/meeting_id":
		return 736
	case "tag/name":
		return 737
	case "tag/tagged_ids":
		return 738
	case "theme/abstain":
		return 739
	case "theme/accent_100":
		return 740
	case "theme/accent_200":
		return 741
	case "theme/accent_300":
		return 742
	case "theme/accent_400":
		return 743
	case "theme/accent_50":
		return 744
	case "theme/accent_500":
		return 745
	case "theme/accent_600":
		return 746
	case "theme/accent_700":
		return 747
	case "theme/accent_800":
		return 748
	case "theme/accent_900":
		return 749
	case "theme/accent_a100":
		return 750
	case "theme/accent_a200":
		return 751
	case "theme/accent_a400":
		return 752
	case "theme/accent_a700":
		return 753
	case "theme/headbar":
		return 754
	case "theme/id":
		return 755
	case "theme/name":
		return 756
	case "theme/no":
		return 757
	case "theme/organization_id":
		return 758
	case "theme/primary_100":
		return 759
	case "theme/primary_200":
		return 760
	case "theme/primary_300":
		return 761
	case "theme/primary_400":
		return 762
	case "theme/primary_50":
		return 763
	case "theme/primary_500":
		return 764
	case "theme/primary_600":
		return 765
	case "theme/primary_700":
		return 766
	case "theme/primary_800":
		return 767
	case "theme/primary_900":
		return 768
	case "theme/primary_a100":
		return 769
	case "theme/primary_a200":
		return 770
	case "theme/primary_a400":
		return 771
	case "theme/primary_a700":
		return 772
	case "theme/theme_for_organization_id":
		return 773
	case "theme/warn_100":
		return 774
	case "theme/warn_200":
		return 775
	case "theme/warn_300":
		return 776
	case "theme/warn_400":
		return 777
	case "theme/warn_50":
		return 778
	case "theme/warn_500":
		return 779
	case "theme/warn_600":
		return 780
	case "theme/warn_700":
		return 781
	case "theme/warn_800":
		return 782
	case "theme/warn_900":
		return 783
	case "theme/warn_a100":
		return 784
	case "theme/warn_a200":
		return 785
	case "theme/warn_a400":
		return 786
	case "theme/warn_a700":
		return 787
	case "theme/yes":
		return 788
	case "topic/agenda_item_id":
		return 789
	case "topic/attachment_ids":
		return 790
	case "topic/id":
		return 791
	case "topic/list_of_speakers_id":
		return 792
	case "topic/meeting_id":
		return 793
	case "topic/poll_ids":
		return 794
	case "topic/projection_ids":
		return 795
	case "topic/sequential_number":
		return 796
	case "topic/text":
		return 797
	case "topic/title":
		return 798
	case "user/can_change_own_password":
		return 799
	case "user/committee_ids":
		return 800
	case "user/committee_management_ids":
		return 801
	case "user/default_password":
		return 802
	case "user/default_vote_weight":
		return 803
	case "user/delegated_vote_ids":
		return 804
	case "user/email":
		return 805
	case "user/first_name":
		return 806
	case "user/forwarding_committee_ids":
		return 807
	case "user/gender":
		return 808
	case "user/id":
		return 809
	case "user/is_active":
		return 810
	case "user/is_demo_user":
		return 811
	case "user/is_physical_person":
		return 812
	case "user/is_present_in_meeting_ids":
		return 813
	case "user/last_email_sent":
		return 814
	case "user/last_login":
		return 815
	case "user/last_name":
		return 816
	case "user/meeting_ids":
		return 817
	case "user/meeting_user_ids":
		return 818
	case "user/option_ids":
		return 819
	case "user/organization_id":
		return 820
	case "user/organization_management_level":
		return 821
	case "user/password":
		return 822
	case "user/poll_candidate_ids":
		return 823
	case "user/poll_voted_ids":
		return 824
	case "user/pronoun":
		return 825
	case "user/saml_id":
		return 826
	case "user/title":
		return 827
	case "user/username":
		return 828
	case "user/vote_ids":
		return 829
	case "vote/delegated_user_id":
		return 830
	case "vote/id":
		return 831
	case "vote/meeting_id":
		return 832
	case "vote/option_id":
		return 833
	case "vote/user_id":
		return 834
	case "vote/user_token":
		return 835
	case "vote/value":
		return 836
	case "vote/weight":
		return 837
	default:
		return -1
	}
}

var collectionFieldToMode = [...]int{
	0,
	0,
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
	5,
	2,
	2,
	2,
	2,
	2,
	6,
	6,
	6,
	6,
	6,
	6,
	6,
	6,
	6,
	6,
	6,
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
	8,
	8,
	8,
	8,
	8,
	8,
	8,
	9,
	9,
	9,
	9,
	9,
	9,
	10,
	10,
	10,
	11,
	11,
	10,
	11,
	10,
	10,
	10,
	10,
	11,
	10,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
	12,
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
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	14,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	15,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	18,
	18,
	18,
	18,
	18,
	18,
	18,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	18,
	18,
	18,
	18,
	18,
	18,
	18,
	18,
	18,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
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
	17,
	17,
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
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	16,
	17,
	16,
	17,
	16,
	16,
	17,
	17,
	17,
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
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
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
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
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
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	17,
	18,
	18,
	19,
	19,
	19,
	21,
	19,
	19,
	19,
	19,
	19,
	19,
	19,
	20,
	19,
	19,
	19,
	19,
	19,
	19,
	19,
	24,
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
	22,
	23,
	22,
	22,
	24,
	24,
	24,
	24,
	22,
	24,
	24,
	25,
	24,
	22,
	22,
	24,
	24,
	24,
	24,
	24,
	24,
	26,
	24,
	24,
	24,
	24,
	24,
	24,
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
	24,
	24,
	23,
	27,
	27,
	27,
	27,
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
	31,
	31,
	31,
	31,
	32,
	32,
	32,
	32,
	32,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	33,
	34,
	34,
	34,
	34,
	34,
	34,
	34,
	35,
	35,
	35,
	35,
	35,
	36,
	36,
	36,
	36,
	36,
	36,
	36,
	36,
	36,
	37,
	37,
	37,
	37,
	37,
	39,
	38,
	38,
	38,
	39,
	38,
	38,
	38,
	38,
	38,
	39,
	41,
	41,
	41,
	40,
	40,
	41,
	41,
	40,
	40,
	40,
	41,
	41,
	40,
	40,
	40,
	41,
	40,
	41,
	40,
	40,
	40,
	40,
	40,
	40,
	40,
	40,
	40,
	40,
	42,
	40,
	40,
	40,
	40,
	40,
	43,
	43,
	43,
	43,
	43,
	44,
	44,
	44,
	44,
	44,
	44,
	45,
	45,
	45,
	45,
	45,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	46,
	48,
	46,
	47,
	47,
	49,
	47,
	47,
	51,
	51,
	51,
	51,
	51,
	52,
	52,
	52,
	52,
	53,
	53,
	53,
	53,
	53,
	53,
	53,
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
	56,
	56,
	56,
	56,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	57,
	58,
	58,
	58,
	58,
	58,
	58,
	58,
	59,
	59,
	59,
	59,
	59,
	59,
	59,
	59,
	59,
	60,
	60,
	60,
	60,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	61,
	62,
	62,
	62,
	62,
	62,
	62,
	62,
	62,
	62,
	62,
	65,
	66,
	66,
	69,
	63,
	63,
	64,
	63,
	66,
	63,
	63,
	65,
	63,
	63,
	63,
	65,
	63,
	63,
	66,
	63,
	63,
	67,
	66,
	68,
	63,
	63,
	63,
	63,
	63,
	63,
	63,
	70,
	70,
	70,
	70,
	70,
	71,
	70,
	70,
}

var relationType = [...]Relation{
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationGenericSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationList,
	RelationList,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationList,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationGenericSingle,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationList,
	RelationList,
	RelationList,
	RelationGenericList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationGenericSingle,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationList,
	RelationNone,
	RelationList,
	RelationList,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationList,
	RelationList,
	RelationList,
	RelationSingle,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationList,
	RelationList,
	RelationList,
	RelationList,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationGenericList,
	RelationSingle,
	RelationList,
	RelationList,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationGenericList,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationGenericSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationGenericList,
	RelationGenericSingle,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationGenericSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationGenericSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationGenericList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationSingle,
	RelationList,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationList,
	RelationList,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationNone,
	RelationList,
	RelationSingle,
	RelationNone,
	RelationSingle,
	RelationSingle,
	RelationSingle,
	RelationNone,
	RelationNone,
	RelationNone,
}

var relationTo = [...]int{
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	20,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	142,
	0,
	8,
	641,
	738,
	0,
	0,
	11,
	106,
	41,
	0,
	0,
	0,
	98,
	157,
	0,
	0,
	0,
	600,
	641,
	0,
	738,
	0,
	27,
	0,
	156,
	367,
	0,
	53,
	0,
	170,
	0,
	83,
	0,
	90,
	46,
	0,
	0,
	0,
	171,
	368,
	184,
	0,
	0,
	70,
	807,
	0,
	801,
	172,
	0,
	551,
	587,
	62,
	800,
	139,
	183,
	0,
	0,
	105,
	112,
	222,
	370,
	0,
	0,
	604,
	50,
	477,
	162,
	272,
	327,
	348,
	0,
	52,
	481,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	244,
	641,
	0,
	707,
	730,
	76,
	0,
	118,
	0,
	0,
	0,
	0,
	77,
	0,
	0,
	98,
	0,
	0,
	107,
	0,
	641,
	0,
	0,
	213,
	214,
	215,
	216,
	217,
	218,
	219,
	220,
	252,
	253,
	254,
	255,
	256,
	257,
	258,
	259,
	72,
	0,
	0,
	18,
	0,
	0,
	0,
	0,
	0,
	645,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	43,
	32,
	0,
	0,
	0,
	0,
	85,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	48,
	57,
	66,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	73,
	59,
	674,
	675,
	676,
	677,
	678,
	679,
	680,
	681,
	682,
	683,
	684,
	685,
	686,
	687,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	123,
	124,
	125,
	126,
	127,
	128,
	129,
	130,
	411,
	78,
	0,
	0,
	549,
	550,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	698,
	0,
	0,
	0,
	0,
	0,
	0,
	100,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	131,
	132,
	133,
	134,
	135,
	136,
	137,
	138,
	117,
	372,
	442,
	450,
	462,
	470,
	475,
	483,
	405,
	0,
	0,
	0,
	86,
	0,
	0,
	495,
	513,
	520,
	529,
	534,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	524,
	0,
	0,
	525,
	526,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	541,
	587,
	590,
	595,
	0,
	0,
	632,
	637,
	699,
	0,
	0,
	87,
	0,
	0,
	0,
	614,
	0,
	813,
	641,
	0,
	694,
	0,
	664,
	701,
	688,
	708,
	0,
	722,
	731,
	736,
	573,
	793,
	88,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	832,
	0,
	0,
	0,
	44,
	58,
	0,
	79,
	0,
	261,
	484,
	521,
	535,
	0,
	591,
	709,
	723,
	431,
	818,
	383,
	382,
	0,
	11,
	387,
	386,
	403,
	0,
	106,
	443,
	451,
	0,
	463,
	471,
	0,
	410,
	485,
	0,
	0,
	401,
	0,
	388,
	98,
	268,
	0,
	0,
	0,
	539,
	397,
	221,
	588,
	600,
	641,
	0,
	0,
	0,
	498,
	417,
	427,
	0,
	423,
	422,
	0,
	0,
	0,
	0,
	497,
	514,
	522,
	380,
	738,
	0,
	0,
	0,
	0,
	536,
	11,
	0,
	0,
	98,
	262,
	391,
	641,
	0,
	0,
	453,
	0,
	0,
	263,
	392,
	0,
	447,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	264,
	394,
	0,
	0,
	0,
	0,
	0,
	0,
	265,
	395,
	473,
	472,
	0,
	266,
	0,
	84,
	0,
	0,
	0,
	91,
	0,
	267,
	373,
	398,
	0,
	0,
	0,
	0,
	0,
	0,
	527,
	0,
	0,
	275,
	0,
	428,
	418,
	0,
	501,
	500,
	0,
	0,
	0,
	0,
	0,
	0,
	509,
	508,
	0,
	532,
	0,
	276,
	429,
	0,
	0,
	0,
	0,
	0,
	277,
	374,
	430,
	0,
	287,
	290,
	291,
	492,
	0,
	278,
	0,
	0,
	511,
	0,
	279,
	375,
	437,
	0,
	0,
	0,
	0,
	316,
	0,
	617,
	0,
	608,
	833,
	0,
	0,
	225,
	226,
	68,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	117,
	0,
	586,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	346,
	773,
	758,
	0,
	820,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	564,
	0,
	0,
	0,
	318,
	377,
	0,
	0,
	0,
	319,
	0,
	713,
	0,
	0,
	0,
	0,
	0,
	0,
	82,
	0,
	0,
	0,
	545,
	0,
	0,
	0,
	0,
	0,
	331,
	0,
	0,
	543,
	0,
	641,
	0,
	0,
	0,
	0,
	0,
	824,
	0,
	0,
	0,
	0,
	0,
	0,
	322,
	639,
	823,
	0,
	0,
	323,
	539,
	633,
	0,
	0,
	657,
	661,
	0,
	148,
	0,
	666,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	642,
	0,
	0,
	0,
	643,
	0,
	0,
	338,
	0,
	647,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	185,
	186,
	187,
	188,
	189,
	190,
	191,
	192,
	193,
	194,
	195,
	196,
	197,
	198,
	340,
	0,
	0,
	0,
	0,
	0,
	336,
	641,
	0,
	0,
	237,
	324,
	0,
	339,
	0,
	641,
	0,
	0,
	0,
	103,
	341,
	378,
	0,
	0,
	0,
	597,
	0,
	733,
	0,
	0,
	0,
	0,
	0,
	0,
	343,
	379,
	0,
	734,
	0,
	0,
	0,
	0,
	104,
	344,
	0,
	715,
	725,
	0,
	345,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	575,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	574,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	11,
	106,
	0,
	98,
	347,
	600,
	641,
	0,
	0,
	0,
	0,
	71,
	65,
	0,
	0,
	830,
	0,
	0,
	63,
	0,
	0,
	0,
	0,
	0,
	333,
	0,
	0,
	0,
	0,
	381,
	539,
	577,
	0,
	0,
	634,
	625,
	0,
	0,
	0,
	0,
	834,
	804,
	0,
	363,
	546,
	829,
	0,
	0,
	0,
}

var relationGenericTo = [...]map[string]int{
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"assignment": 25, "motion": 385, "motion_block": 438, "topic": 789},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"assignment": 31, "mediafile": 115, "motion": 404, "motion_block": 441, "topic": 792},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"assignment": 26, "motion": 390, "topic": 790},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"meeting": 260, "organization": 562},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"motion": 419},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"motion": 420},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"motion": 409, "poll_candidate_list": 638, "user": 819},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"committee": 69, "meeting": 317},
	{"motion": 412},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"assignment": 36, "motion": 413, "topic": 794},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"agenda_item": 21, "assignment": 37, "list_of_speakers": 101, "mediafile": 120, "meeting": 334, "motion": 414, "motion_block": 444, "poll": 619, "projector_countdown": 695, "projector_message": 703, "topic": 795},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	{"agenda_item": 22, "assignment": 39, "motion": 432},
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
}

var collectionModeFields = [...]collectionMode{
	{"invalid", "mode"},
	{"action_worker", "A"},
	{"agenda_item", "A"},
	{"agenda_item", "B"},
	{"agenda_item", "C"},
	{"agenda_item", "D"},
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
	{"motion", "B"},
	{"motion", "C"},
	{"motion", "D"},
	{"motion", "E"},
	{"motion_block", "A"},
	{"motion_category", "A"},
	{"motion_change_recommendation", "A"},
	{"motion_comment", "A"},
	{"motion_comment_section", "A"},
	{"motion_editor", "A"},
	{"motion_state", "A"},
	{"motion_statute_paragraph", "A"},
	{"motion_submitter", "A"},
	{"motion_workflow", "A"},
	{"motion_working_group_speaker", "A"},
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
	{"structure_level", "A"},
	{"structure_level_list_of_speakers", "A"},
	{"tag", "A"},
	{"theme", "A"},
	{"topic", "A"},
	{"user", "A"},
	{"user", "B"},
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
	case "agenda_item/D":
		return 5
	case "assignment/A":
		return 6
	case "assignment_candidate/A":
		return 7
	case "chat_group/A":
		return 8
	case "chat_message/A":
		return 9
	case "committee/A":
		return 10
	case "committee/B":
		return 11
	case "group/A":
		return 12
	case "import_preview/A":
		return 13
	case "list_of_speakers/A":
		return 14
	case "mediafile/A":
		return 15
	case "meeting/A":
		return 16
	case "meeting/B":
		return 17
	case "meeting/C":
		return 18
	case "meeting_user/A":
		return 19
	case "meeting_user/B":
		return 20
	case "meeting_user/D":
		return 21
	case "motion/A":
		return 22
	case "motion/B":
		return 23
	case "motion/C":
		return 24
	case "motion/D":
		return 25
	case "motion/E":
		return 26
	case "motion_block/A":
		return 27
	case "motion_category/A":
		return 28
	case "motion_change_recommendation/A":
		return 29
	case "motion_comment/A":
		return 30
	case "motion_comment_section/A":
		return 31
	case "motion_editor/A":
		return 32
	case "motion_state/A":
		return 33
	case "motion_statute_paragraph/A":
		return 34
	case "motion_submitter/A":
		return 35
	case "motion_workflow/A":
		return 36
	case "motion_working_group_speaker/A":
		return 37
	case "option/A":
		return 38
	case "option/B":
		return 39
	case "organization/A":
		return 40
	case "organization/B":
		return 41
	case "organization/C":
		return 42
	case "organization_tag/A":
		return 43
	case "personal_note/A":
		return 44
	case "point_of_order_category/A":
		return 45
	case "poll/A":
		return 46
	case "poll/B":
		return 47
	case "poll/C":
		return 48
	case "poll/D":
		return 49
	case "poll/MANAGE":
		return 50
	case "poll_candidate/A":
		return 51
	case "poll_candidate_list/A":
		return 52
	case "projection/A":
		return 53
	case "projector/A":
		return 54
	case "projector_countdown/A":
		return 55
	case "projector_message/A":
		return 56
	case "speaker/A":
		return 57
	case "structure_level/A":
		return 58
	case "structure_level_list_of_speakers/A":
		return 59
	case "tag/A":
		return 60
	case "theme/A":
		return 61
	case "topic/A":
		return 62
	case "user/A":
		return 63
	case "user/B":
		return 64
	case "user/D":
		return 65
	case "user/E":
		return 66
	case "user/F":
		return 67
	case "user/G":
		return 68
	case "user/H":
		return 69
	case "vote/A":
		return 70
	case "vote/B":
		return 71
	default:
		return -1
	}
}
