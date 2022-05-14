// Code generated from models.yml DO NOT EDIT.
package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// ValueBool is a lazy value from the datastore.
type ValueBool struct {
	value    bool
	isNull   bool
	executed bool

	lazies []*bool

	request *Request
}

// Value returns the value.
func (v *ValueBool) Value(ctx context.Context) (bool, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return false, fmt.Errorf("executing request: %w", err)
	}

	return v.value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueBool) Lazy(value *bool) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueBool) ErrorLater(ctx context.Context) bool {
	if v.request.err != nil {
		return false
	}

	if v.executed {
		return v.value
	}

	if err := v.request.Execute(ctx); err != nil {
		return false
	}

	return v.value
}

// execute will be called from request.
func (v *ValueBool) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueFloat is a lazy value from the datastore.
type ValueFloat struct {
	value    float32
	isNull   bool
	executed bool

	lazies []*float32

	request *Request
}

// Value returns the value.
func (v *ValueFloat) Value(ctx context.Context) (float32, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0, fmt.Errorf("executing request: %w", err)
	}

	return v.value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueFloat) Lazy(value *float32) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueFloat) ErrorLater(ctx context.Context) float32 {
	if v.request.err != nil {
		return 0
	}

	if v.executed {
		return v.value
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0
	}

	return v.value
}

// execute will be called from request.
func (v *ValueFloat) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueIDSlice is a lazy value from the datastore.
type ValueIDSlice struct {
	value    []int
	isNull   bool
	executed bool

	lazies []*[]int

	request *Request
}

// Value returns the value.
func (v *ValueIDSlice) Value(ctx context.Context) ([]int, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	return v.value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueIDSlice) Lazy(value *[]int) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueIDSlice) ErrorLater(ctx context.Context) []int {
	if v.request.err != nil {
		return nil
	}

	if v.executed {
		return v.value
	}

	if err := v.request.Execute(ctx); err != nil {
		return nil
	}

	return v.value
}

// execute will be called from request.
func (v *ValueIDSlice) execute(p []byte) error {
	var values []string
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &values); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for _, e := range values {
		i, err := strconv.Atoi(e)
		if err != nil {
			return fmt.Errorf("converting value %q: %w", e, err)
		}
		v.value = append(v.value, i)
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueInt is a lazy value from the datastore.
type ValueInt struct {
	value    int
	isNull   bool
	executed bool

	lazies []*int

	request *Request
}

// Value returns the value.
func (v *ValueInt) Value(ctx context.Context) (int, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0, fmt.Errorf("executing request: %w", err)
	}

	return v.value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueInt) Lazy(value *int) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueInt) ErrorLater(ctx context.Context) int {
	if v.request.err != nil {
		return 0
	}

	if v.executed {
		return v.value
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0
	}

	return v.value
}

// execute will be called from request.
func (v *ValueInt) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueIntSlice is a lazy value from the datastore.
type ValueIntSlice struct {
	value    []int
	isNull   bool
	executed bool

	lazies []*[]int

	request *Request
}

// Value returns the value.
func (v *ValueIntSlice) Value(ctx context.Context) ([]int, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	return v.value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueIntSlice) Lazy(value *[]int) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueIntSlice) ErrorLater(ctx context.Context) []int {
	if v.request.err != nil {
		return nil
	}

	if v.executed {
		return v.value
	}

	if err := v.request.Execute(ctx); err != nil {
		return nil
	}

	return v.value
}

// execute will be called from request.
func (v *ValueIntSlice) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueJSON is a lazy value from the datastore.
type ValueJSON struct {
	value    json.RawMessage
	isNull   bool
	executed bool

	lazies []*json.RawMessage

	request *Request
}

// Value returns the value.
func (v *ValueJSON) Value(ctx context.Context) (json.RawMessage, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	return v.value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueJSON) Lazy(value *json.RawMessage) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueJSON) ErrorLater(ctx context.Context) json.RawMessage {
	if v.request.err != nil {
		return nil
	}

	if v.executed {
		return v.value
	}

	if err := v.request.Execute(ctx); err != nil {
		return nil
	}

	return v.value
}

// execute will be called from request.
func (v *ValueJSON) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueMaybeInt is a lazy value from the datastore.
type ValueMaybeInt struct {
	value    int
	isNull   bool
	executed bool

	lazies []*int

	request *Request
}

// Value returns the value.
func (v *ValueMaybeInt) Value(ctx context.Context) (int, bool, error) {
	if v.executed {
		return v.value, !v.isNull, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0, false, fmt.Errorf("executing request: %w", err)
	}

	return v.value, !v.isNull, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueMaybeInt) Lazy(value *int) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueMaybeInt) ErrorLater(ctx context.Context) (int, bool) {
	if v.request.err != nil {
		return 0, false
	}

	if v.executed {
		return v.value, !v.isNull
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0, false
	}

	return v.value, !v.isNull
}

// execute will be called from request.
func (v *ValueMaybeInt) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueMaybeString is a lazy value from the datastore.
type ValueMaybeString struct {
	value    string
	isNull   bool
	executed bool

	lazies []*string

	request *Request
}

// Value returns the value.
func (v *ValueMaybeString) Value(ctx context.Context) (string, bool, error) {
	if v.executed {
		return v.value, !v.isNull, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return "", false, fmt.Errorf("executing request: %w", err)
	}

	return v.value, !v.isNull, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueMaybeString) Lazy(value *string) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueMaybeString) ErrorLater(ctx context.Context) (string, bool) {
	if v.request.err != nil {
		return "", false
	}

	if v.executed {
		return v.value, !v.isNull
	}

	if err := v.request.Execute(ctx); err != nil {
		return "", false
	}

	return v.value, !v.isNull
}

// execute will be called from request.
func (v *ValueMaybeString) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueString is a lazy value from the datastore.
type ValueString struct {
	value    string
	isNull   bool
	executed bool

	lazies []*string

	request *Request
}

// Value returns the value.
func (v *ValueString) Value(ctx context.Context) (string, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return "", fmt.Errorf("executing request: %w", err)
	}

	return v.value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueString) Lazy(value *string) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueString) ErrorLater(ctx context.Context) string {
	if v.request.err != nil {
		return ""
	}

	if v.executed {
		return v.value
	}

	if err := v.request.Execute(ctx); err != nil {
		return ""
	}

	return v.value
}

// execute will be called from request.
func (v *ValueString) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

// ValueStringSlice is a lazy value from the datastore.
type ValueStringSlice struct {
	value    []string
	isNull   bool
	executed bool

	lazies []*[]string

	request *Request
}

// Value returns the value.
func (v *ValueStringSlice) Value(ctx context.Context) ([]string, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	return v.value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueStringSlice) Lazy(value *[]string) {
	v.lazies = append(v.lazies, value)
}

// ErrorLater is like Value but does not return an error.
//
// If an error happs, it is saved internaly. Make sure to call request.Err() later to
// access it.
func (v *ValueStringSlice) ErrorLater(ctx context.Context) []string {
	if v.request.err != nil {
		return nil
	}

	if v.executed {
		return v.value
	}

	if err := v.request.Execute(ctx); err != nil {
		return nil
	}

	return v.value
}

// execute will be called from request.
func (v *ValueStringSlice) execute(p []byte) error {
	if p == nil {
		v.isNull = true
	} else {
		if err := json.Unmarshal(p, &v.value); err != nil {
			return fmt.Errorf("decoding value %q: %v", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

func (r *Request) AgendaItem_ChildIDs(agendaItemID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "child_ids"}] = v
	return v
}

func (r *Request) AgendaItem_Closed(agendaItemID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "closed"}] = v
	return v
}

func (r *Request) AgendaItem_Comment(agendaItemID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "comment"}] = v
	return v
}

func (r *Request) AgendaItem_ContentObjectID(agendaItemID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "content_object_id"}] = v
	return v
}

func (r *Request) AgendaItem_Duration(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "duration"}] = v
	return v
}

func (r *Request) AgendaItem_ID(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "id"}] = v
	return v
}

func (r *Request) AgendaItem_IsHidden(agendaItemID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "is_hidden"}] = v
	return v
}

func (r *Request) AgendaItem_IsInternal(agendaItemID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "is_internal"}] = v
	return v
}

func (r *Request) AgendaItem_ItemNumber(agendaItemID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "item_number"}] = v
	return v
}

func (r *Request) AgendaItem_Level(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "level"}] = v
	return v
}

func (r *Request) AgendaItem_MeetingID(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "meeting_id"}] = v
	return v
}

func (r *Request) AgendaItem_ParentID(agendaItemID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "parent_id"}] = v
	return v
}

func (r *Request) AgendaItem_ProjectionIDs(agendaItemID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "projection_ids"}] = v
	return v
}

func (r *Request) AgendaItem_TagIDs(agendaItemID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "tag_ids"}] = v
	return v
}

func (r *Request) AgendaItem_Type(agendaItemID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "type"}] = v
	return v
}

func (r *Request) AgendaItem_Weight(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"agenda_item", agendaItemID, "weight"}] = v
	return v
}

func (r *Request) AssignmentCandidate_AssignmentID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment_candidate", assignmentCandidateID, "assignment_id"}] = v
	return v
}

func (r *Request) AssignmentCandidate_ID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment_candidate", assignmentCandidateID, "id"}] = v
	return v
}

func (r *Request) AssignmentCandidate_MeetingID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment_candidate", assignmentCandidateID, "meeting_id"}] = v
	return v
}

func (r *Request) AssignmentCandidate_UserID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment_candidate", assignmentCandidateID, "user_id"}] = v
	return v
}

func (r *Request) AssignmentCandidate_Weight(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment_candidate", assignmentCandidateID, "weight"}] = v
	return v
}

func (r *Request) Assignment_AgendaItemID(assignmentID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"assignment", assignmentID, "agenda_item_id"}] = v
	return v
}

func (r *Request) Assignment_AttachmentIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"assignment", assignmentID, "attachment_ids"}] = v
	return v
}

func (r *Request) Assignment_CandidateIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"assignment", assignmentID, "candidate_ids"}] = v
	return v
}

func (r *Request) Assignment_DefaultPollDescription(assignmentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"assignment", assignmentID, "default_poll_description"}] = v
	return v
}

func (r *Request) Assignment_Description(assignmentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"assignment", assignmentID, "description"}] = v
	return v
}

func (r *Request) Assignment_ID(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment", assignmentID, "id"}] = v
	return v
}

func (r *Request) Assignment_ListOfSpeakersID(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment", assignmentID, "list_of_speakers_id"}] = v
	return v
}

func (r *Request) Assignment_MeetingID(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment", assignmentID, "meeting_id"}] = v
	return v
}

func (r *Request) Assignment_NumberPollCandidates(assignmentID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"assignment", assignmentID, "number_poll_candidates"}] = v
	return v
}

func (r *Request) Assignment_OpenPosts(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment", assignmentID, "open_posts"}] = v
	return v
}

func (r *Request) Assignment_Phase(assignmentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"assignment", assignmentID, "phase"}] = v
	return v
}

func (r *Request) Assignment_PollIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"assignment", assignmentID, "poll_ids"}] = v
	return v
}

func (r *Request) Assignment_ProjectionIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"assignment", assignmentID, "projection_ids"}] = v
	return v
}

func (r *Request) Assignment_SequentialNumber(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"assignment", assignmentID, "sequential_number"}] = v
	return v
}

func (r *Request) Assignment_TagIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"assignment", assignmentID, "tag_ids"}] = v
	return v
}

func (r *Request) Assignment_Title(assignmentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"assignment", assignmentID, "title"}] = v
	return v
}

func (r *Request) ChatGroup_ChatMessageIDs(chatGroupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"chat_group", chatGroupID, "chat_message_ids"}] = v
	return v
}

func (r *Request) ChatGroup_ID(chatGroupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"chat_group", chatGroupID, "id"}] = v
	return v
}

func (r *Request) ChatGroup_MeetingID(chatGroupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"chat_group", chatGroupID, "meeting_id"}] = v
	return v
}

func (r *Request) ChatGroup_Name(chatGroupID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"chat_group", chatGroupID, "name"}] = v
	return v
}

func (r *Request) ChatGroup_ReadGroupIDs(chatGroupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"chat_group", chatGroupID, "read_group_ids"}] = v
	return v
}

func (r *Request) ChatGroup_Weight(chatGroupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"chat_group", chatGroupID, "weight"}] = v
	return v
}

func (r *Request) ChatGroup_WriteGroupIDs(chatGroupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"chat_group", chatGroupID, "write_group_ids"}] = v
	return v
}

func (r *Request) ChatMessage_ChatGroupID(chatMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"chat_message", chatMessageID, "chat_group_id"}] = v
	return v
}

func (r *Request) ChatMessage_Content(chatMessageID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"chat_message", chatMessageID, "content"}] = v
	return v
}

func (r *Request) ChatMessage_Created(chatMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"chat_message", chatMessageID, "created"}] = v
	return v
}

func (r *Request) ChatMessage_ID(chatMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"chat_message", chatMessageID, "id"}] = v
	return v
}

func (r *Request) ChatMessage_MeetingID(chatMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"chat_message", chatMessageID, "meeting_id"}] = v
	return v
}

func (r *Request) ChatMessage_UserID(chatMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"chat_message", chatMessageID, "user_id"}] = v
	return v
}

func (r *Request) Committee_DefaultMeetingID(committeeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"committee", committeeID, "default_meeting_id"}] = v
	return v
}

func (r *Request) Committee_Description(committeeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"committee", committeeID, "description"}] = v
	return v
}

func (r *Request) Committee_ForwardToCommitteeIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"committee", committeeID, "forward_to_committee_ids"}] = v
	return v
}

func (r *Request) Committee_ForwardingUserID(committeeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"committee", committeeID, "forwarding_user_id"}] = v
	return v
}

func (r *Request) Committee_ID(committeeID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"committee", committeeID, "id"}] = v
	return v
}

func (r *Request) Committee_MeetingIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"committee", committeeID, "meeting_ids"}] = v
	return v
}

func (r *Request) Committee_Name(committeeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"committee", committeeID, "name"}] = v
	return v
}

func (r *Request) Committee_OrganizationID(committeeID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"committee", committeeID, "organization_id"}] = v
	return v
}

func (r *Request) Committee_OrganizationTagIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"committee", committeeID, "organization_tag_ids"}] = v
	return v
}

func (r *Request) Committee_ReceiveForwardingsFromCommitteeIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"committee", committeeID, "receive_forwardings_from_committee_ids"}] = v
	return v
}

func (r *Request) Committee_UserIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"committee", committeeID, "user_ids"}] = v
	return v
}

func (r *Request) Committee_UserManagementLevelTmpl(committeeID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"committee", committeeID, "user_$_management_level"}] = v
	return v
}

func (r *Request) Committee_UserManagementLevel(committeeID int, replacement string) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"committee", committeeID, fmt.Sprintf("user_$%s_management_level", replacement)}] = v
	return v
}

func (r *Request) Group_AdminGroupForMeetingID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"group", groupID, "admin_group_for_meeting_id"}] = v
	return v
}

func (r *Request) Group_DefaultGroupForMeetingID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"group", groupID, "default_group_for_meeting_id"}] = v
	return v
}

func (r *Request) Group_ID(groupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"group", groupID, "id"}] = v
	return v
}

func (r *Request) Group_MediafileAccessGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"group", groupID, "mediafile_access_group_ids"}] = v
	return v
}

func (r *Request) Group_MediafileInheritedAccessGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"group", groupID, "mediafile_inherited_access_group_ids"}] = v
	return v
}

func (r *Request) Group_MeetingID(groupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"group", groupID, "meeting_id"}] = v
	return v
}

func (r *Request) Group_Name(groupID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"group", groupID, "name"}] = v
	return v
}

func (r *Request) Group_Permissions(groupID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"group", groupID, "permissions"}] = v
	return v
}

func (r *Request) Group_PollIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"group", groupID, "poll_ids"}] = v
	return v
}

func (r *Request) Group_ReadChatGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"group", groupID, "read_chat_group_ids"}] = v
	return v
}

func (r *Request) Group_ReadCommentSectionIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"group", groupID, "read_comment_section_ids"}] = v
	return v
}

func (r *Request) Group_UsedAsAssignmentPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"group", groupID, "used_as_assignment_poll_default_id"}] = v
	return v
}

func (r *Request) Group_UsedAsMotionPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"group", groupID, "used_as_motion_poll_default_id"}] = v
	return v
}

func (r *Request) Group_UsedAsPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"group", groupID, "used_as_poll_default_id"}] = v
	return v
}

func (r *Request) Group_UserIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"group", groupID, "user_ids"}] = v
	return v
}

func (r *Request) Group_Weight(groupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"group", groupID, "weight"}] = v
	return v
}

func (r *Request) Group_WriteChatGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"group", groupID, "write_chat_group_ids"}] = v
	return v
}

func (r *Request) Group_WriteCommentSectionIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"group", groupID, "write_comment_section_ids"}] = v
	return v
}

func (r *Request) ListOfSpeakers_Closed(listOfSpeakersID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"list_of_speakers", listOfSpeakersID, "closed"}] = v
	return v
}

func (r *Request) ListOfSpeakers_ContentObjectID(listOfSpeakersID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"list_of_speakers", listOfSpeakersID, "content_object_id"}] = v
	return v
}

func (r *Request) ListOfSpeakers_ID(listOfSpeakersID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"list_of_speakers", listOfSpeakersID, "id"}] = v
	return v
}

func (r *Request) ListOfSpeakers_MeetingID(listOfSpeakersID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"list_of_speakers", listOfSpeakersID, "meeting_id"}] = v
	return v
}

func (r *Request) ListOfSpeakers_ProjectionIDs(listOfSpeakersID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"list_of_speakers", listOfSpeakersID, "projection_ids"}] = v
	return v
}

func (r *Request) ListOfSpeakers_SequentialNumber(listOfSpeakersID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"list_of_speakers", listOfSpeakersID, "sequential_number"}] = v
	return v
}

func (r *Request) ListOfSpeakers_SpeakerIDs(listOfSpeakersID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"list_of_speakers", listOfSpeakersID, "speaker_ids"}] = v
	return v
}

func (r *Request) Mediafile_AccessGroupIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"mediafile", mediafileID, "access_group_ids"}] = v
	return v
}

func (r *Request) Mediafile_AttachmentIDs(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"mediafile", mediafileID, "attachment_ids"}] = v
	return v
}

func (r *Request) Mediafile_ChildIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"mediafile", mediafileID, "child_ids"}] = v
	return v
}

func (r *Request) Mediafile_CreateTimestamp(mediafileID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"mediafile", mediafileID, "create_timestamp"}] = v
	return v
}

func (r *Request) Mediafile_Filename(mediafileID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"mediafile", mediafileID, "filename"}] = v
	return v
}

func (r *Request) Mediafile_Filesize(mediafileID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"mediafile", mediafileID, "filesize"}] = v
	return v
}

func (r *Request) Mediafile_ID(mediafileID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"mediafile", mediafileID, "id"}] = v
	return v
}

func (r *Request) Mediafile_InheritedAccessGroupIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"mediafile", mediafileID, "inherited_access_group_ids"}] = v
	return v
}

func (r *Request) Mediafile_IsDirectory(mediafileID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"mediafile", mediafileID, "is_directory"}] = v
	return v
}

func (r *Request) Mediafile_IsPublic(mediafileID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"mediafile", mediafileID, "is_public"}] = v
	return v
}

func (r *Request) Mediafile_ListOfSpeakersID(mediafileID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"mediafile", mediafileID, "list_of_speakers_id"}] = v
	return v
}

func (r *Request) Mediafile_Mimetype(mediafileID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"mediafile", mediafileID, "mimetype"}] = v
	return v
}

func (r *Request) Mediafile_OwnerID(mediafileID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"mediafile", mediafileID, "owner_id"}] = v
	return v
}

func (r *Request) Mediafile_ParentID(mediafileID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"mediafile", mediafileID, "parent_id"}] = v
	return v
}

func (r *Request) Mediafile_PdfInformation(mediafileID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[Key{"mediafile", mediafileID, "pdf_information"}] = v
	return v
}

func (r *Request) Mediafile_ProjectionIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"mediafile", mediafileID, "projection_ids"}] = v
	return v
}

func (r *Request) Mediafile_Title(mediafileID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"mediafile", mediafileID, "title"}] = v
	return v
}

func (r *Request) Mediafile_Token(mediafileID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"mediafile", mediafileID, "token"}] = v
	return v
}

func (r *Request) Mediafile_UsedAsFontInMeetingIDTmpl(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"mediafile", mediafileID, "used_as_font_$_in_meeting_id"}] = v
	return v
}

func (r *Request) Mediafile_UsedAsFontInMeetingID(mediafileID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"mediafile", mediafileID, fmt.Sprintf("used_as_font_$%s_in_meeting_id", replacement)}] = v
	return v
}

func (r *Request) Mediafile_UsedAsLogoInMeetingIDTmpl(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"mediafile", mediafileID, "used_as_logo_$_in_meeting_id"}] = v
	return v
}

func (r *Request) Mediafile_UsedAsLogoInMeetingID(mediafileID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"mediafile", mediafileID, fmt.Sprintf("used_as_logo_$%s_in_meeting_id", replacement)}] = v
	return v
}

func (r *Request) Meeting_AdminGroupID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"meeting", meetingID, "admin_group_id"}] = v
	return v
}

func (r *Request) Meeting_AgendaEnableNumbering(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "agenda_enable_numbering"}] = v
	return v
}

func (r *Request) Meeting_AgendaItemCreation(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "agenda_item_creation"}] = v
	return v
}

func (r *Request) Meeting_AgendaItemIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "agenda_item_ids"}] = v
	return v
}

func (r *Request) Meeting_AgendaNewItemsDefaultVisibility(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "agenda_new_items_default_visibility"}] = v
	return v
}

func (r *Request) Meeting_AgendaNumberPrefix(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "agenda_number_prefix"}] = v
	return v
}

func (r *Request) Meeting_AgendaNumeralSystem(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "agenda_numeral_system"}] = v
	return v
}

func (r *Request) Meeting_AgendaShowInternalItemsOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "agenda_show_internal_items_on_projector"}] = v
	return v
}

func (r *Request) Meeting_AgendaShowSubtitles(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "agenda_show_subtitles"}] = v
	return v
}

func (r *Request) Meeting_AllProjectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "all_projection_ids"}] = v
	return v
}

func (r *Request) Meeting_ApplauseEnable(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "applause_enable"}] = v
	return v
}

func (r *Request) Meeting_ApplauseMaxAmount(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "applause_max_amount"}] = v
	return v
}

func (r *Request) Meeting_ApplauseMinAmount(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "applause_min_amount"}] = v
	return v
}

func (r *Request) Meeting_ApplauseParticleImageUrl(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "applause_particle_image_url"}] = v
	return v
}

func (r *Request) Meeting_ApplauseShowLevel(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "applause_show_level"}] = v
	return v
}

func (r *Request) Meeting_ApplauseTimeout(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "applause_timeout"}] = v
	return v
}

func (r *Request) Meeting_ApplauseType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "applause_type"}] = v
	return v
}

func (r *Request) Meeting_AssignmentCandidateIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_candidate_ids"}] = v
	return v
}

func (r *Request) Meeting_AssignmentIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_ids"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollAddCandidatesToListOfSpeakers(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_add_candidates_to_list_of_speakers"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_ballot_paper_number"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_ballot_paper_selection"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_default_100_percent_base"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefaultBackend(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_default_backend"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_default_group_ids"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefaultMethod(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_default_method"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefaultType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_default_type"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollEnableMaxVotesPerOption(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_enable_max_votes_per_option"}] = v
	return v
}

func (r *Request) Meeting_AssignmentPollSortPollResultByVotes(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "assignment_poll_sort_poll_result_by_votes"}] = v
	return v
}

func (r *Request) Meeting_AssignmentsExportPreamble(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "assignments_export_preamble"}] = v
	return v
}

func (r *Request) Meeting_AssignmentsExportTitle(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "assignments_export_title"}] = v
	return v
}

func (r *Request) Meeting_ChatGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "chat_group_ids"}] = v
	return v
}

func (r *Request) Meeting_ChatMessageIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "chat_message_ids"}] = v
	return v
}

func (r *Request) Meeting_CommitteeID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "committee_id"}] = v
	return v
}

func (r *Request) Meeting_ConferenceAutoConnect(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "conference_auto_connect"}] = v
	return v
}

func (r *Request) Meeting_ConferenceAutoConnectNextSpeakers(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "conference_auto_connect_next_speakers"}] = v
	return v
}

func (r *Request) Meeting_ConferenceEnableHelpdesk(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "conference_enable_helpdesk"}] = v
	return v
}

func (r *Request) Meeting_ConferenceLosRestriction(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "conference_los_restriction"}] = v
	return v
}

func (r *Request) Meeting_ConferenceOpenMicrophone(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "conference_open_microphone"}] = v
	return v
}

func (r *Request) Meeting_ConferenceOpenVideo(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "conference_open_video"}] = v
	return v
}

func (r *Request) Meeting_ConferenceShow(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "conference_show"}] = v
	return v
}

func (r *Request) Meeting_ConferenceStreamPosterUrl(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "conference_stream_poster_url"}] = v
	return v
}

func (r *Request) Meeting_ConferenceStreamUrl(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "conference_stream_url"}] = v
	return v
}

func (r *Request) Meeting_CustomTranslations(meetingID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[Key{"meeting", meetingID, "custom_translations"}] = v
	return v
}

func (r *Request) Meeting_DefaultGroupID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "default_group_id"}] = v
	return v
}

func (r *Request) Meeting_DefaultMeetingForCommitteeID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"meeting", meetingID, "default_meeting_for_committee_id"}] = v
	return v
}

func (r *Request) Meeting_DefaultProjectorIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"meeting", meetingID, "default_projector_$_id"}] = v
	return v
}

func (r *Request) Meeting_DefaultProjectorID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, fmt.Sprintf("default_projector_$%s_id", replacement)}] = v
	return v
}

func (r *Request) Meeting_Description(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "description"}] = v
	return v
}

func (r *Request) Meeting_EnableAnonymous(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "enable_anonymous"}] = v
	return v
}

func (r *Request) Meeting_EndTime(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "end_time"}] = v
	return v
}

func (r *Request) Meeting_ExportCsvEncoding(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "export_csv_encoding"}] = v
	return v
}

func (r *Request) Meeting_ExportCsvSeparator(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "export_csv_separator"}] = v
	return v
}

func (r *Request) Meeting_ExportPdfFontsize(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "export_pdf_fontsize"}] = v
	return v
}

func (r *Request) Meeting_ExportPdfLineHeight(meetingID int) *ValueFloat {
	v := &ValueFloat{request: r}
	r.requested[Key{"meeting", meetingID, "export_pdf_line_height"}] = v
	return v
}

func (r *Request) Meeting_ExportPdfPageMarginBottom(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "export_pdf_page_margin_bottom"}] = v
	return v
}

func (r *Request) Meeting_ExportPdfPageMarginLeft(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "export_pdf_page_margin_left"}] = v
	return v
}

func (r *Request) Meeting_ExportPdfPageMarginRight(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "export_pdf_page_margin_right"}] = v
	return v
}

func (r *Request) Meeting_ExportPdfPageMarginTop(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "export_pdf_page_margin_top"}] = v
	return v
}

func (r *Request) Meeting_ExportPdfPagenumberAlignment(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "export_pdf_pagenumber_alignment"}] = v
	return v
}

func (r *Request) Meeting_ExportPdfPagesize(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "export_pdf_pagesize"}] = v
	return v
}

func (r *Request) Meeting_FontIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"meeting", meetingID, "font_$_id"}] = v
	return v
}

func (r *Request) Meeting_FontID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, fmt.Sprintf("font_$%s_id", replacement)}] = v
	return v
}

func (r *Request) Meeting_GroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "group_ids"}] = v
	return v
}

func (r *Request) Meeting_ID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "id"}] = v
	return v
}

func (r *Request) Meeting_ImportedAt(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "imported_at"}] = v
	return v
}

func (r *Request) Meeting_IsActiveInOrganizationID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"meeting", meetingID, "is_active_in_organization_id"}] = v
	return v
}

func (r *Request) Meeting_IsArchivedInOrganizationID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"meeting", meetingID, "is_archived_in_organization_id"}] = v
	return v
}

func (r *Request) Meeting_JitsiDomain(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "jitsi_domain"}] = v
	return v
}

func (r *Request) Meeting_JitsiRoomName(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "jitsi_room_name"}] = v
	return v
}

func (r *Request) Meeting_JitsiRoomPassword(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "jitsi_room_password"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersAmountLastOnProjector(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_amount_last_on_projector"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersAmountNextOnProjector(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_amount_next_on_projector"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersCanSetContributionSelf(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_can_set_contribution_self"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersCountdownID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_countdown_id"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersCoupleCountdown(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_couple_countdown"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_enable_point_of_order_speakers"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersEnableProContraSpeech(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_enable_pro_contra_speech"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_ids"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersInitiallyClosed(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_initially_closed"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersPresentUsersOnly(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_present_users_only"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_show_amount_of_speakers_on_slide"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersShowFirstContribution(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_show_first_contribution"}] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersSpeakerNoteForEveryone(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "list_of_speakers_speaker_note_for_everyone"}] = v
	return v
}

func (r *Request) Meeting_Location(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "location"}] = v
	return v
}

func (r *Request) Meeting_LogoIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"meeting", meetingID, "logo_$_id"}] = v
	return v
}

func (r *Request) Meeting_LogoID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, fmt.Sprintf("logo_$%s_id", replacement)}] = v
	return v
}

func (r *Request) Meeting_MediafileIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "mediafile_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionBlockIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_block_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionCategoryIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_category_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionChangeRecommendationIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_change_recommendation_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionCommentIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_comment_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionCommentSectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_comment_section_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionPollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "motion_poll_ballot_paper_number"}] = v
	return v
}

func (r *Request) Meeting_MotionPollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motion_poll_ballot_paper_selection"}] = v
	return v
}

func (r *Request) Meeting_MotionPollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motion_poll_default_100_percent_base"}] = v
	return v
}

func (r *Request) Meeting_MotionPollDefaultBackend(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motion_poll_default_backend"}] = v
	return v
}

func (r *Request) Meeting_MotionPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_poll_default_group_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionPollDefaultType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motion_poll_default_type"}] = v
	return v
}

func (r *Request) Meeting_MotionStateIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_state_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionStatuteParagraphIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_statute_paragraph_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionSubmitterIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_submitter_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionWorkflowIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "motion_workflow_ids"}] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsEnabled(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_amendments_enabled"}] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsInMainList(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_amendments_in_main_list"}] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsMultipleParagraphs(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_amendments_multiple_paragraphs"}] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsOfAmendments(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_amendments_of_amendments"}] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsPrefix(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_amendments_prefix"}] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsTextMode(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_amendments_text_mode"}] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultAmendmentWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "motions_default_amendment_workflow_id"}] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultLineNumbering(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_default_line_numbering"}] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultSorting(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_default_sorting"}] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultStatuteAmendmentWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "motions_default_statute_amendment_workflow_id"}] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "motions_default_workflow_id"}] = v
	return v
}

func (r *Request) Meeting_MotionsEnableReasonOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_enable_reason_on_projector"}] = v
	return v
}

func (r *Request) Meeting_MotionsEnableRecommendationOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_enable_recommendation_on_projector"}] = v
	return v
}

func (r *Request) Meeting_MotionsEnableSideboxOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_enable_sidebox_on_projector"}] = v
	return v
}

func (r *Request) Meeting_MotionsEnableTextOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_enable_text_on_projector"}] = v
	return v
}

func (r *Request) Meeting_MotionsExportFollowRecommendation(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_export_follow_recommendation"}] = v
	return v
}

func (r *Request) Meeting_MotionsExportPreamble(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_export_preamble"}] = v
	return v
}

func (r *Request) Meeting_MotionsExportSubmitterRecommendation(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_export_submitter_recommendation"}] = v
	return v
}

func (r *Request) Meeting_MotionsExportTitle(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_export_title"}] = v
	return v
}

func (r *Request) Meeting_MotionsLineLength(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "motions_line_length"}] = v
	return v
}

func (r *Request) Meeting_MotionsNumberMinDigits(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "motions_number_min_digits"}] = v
	return v
}

func (r *Request) Meeting_MotionsNumberType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_number_type"}] = v
	return v
}

func (r *Request) Meeting_MotionsNumberWithBlank(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_number_with_blank"}] = v
	return v
}

func (r *Request) Meeting_MotionsPreamble(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_preamble"}] = v
	return v
}

func (r *Request) Meeting_MotionsReasonRequired(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_reason_required"}] = v
	return v
}

func (r *Request) Meeting_MotionsRecommendationTextMode(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_recommendation_text_mode"}] = v
	return v
}

func (r *Request) Meeting_MotionsRecommendationsBy(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_recommendations_by"}] = v
	return v
}

func (r *Request) Meeting_MotionsShowReferringMotions(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_show_referring_motions"}] = v
	return v
}

func (r *Request) Meeting_MotionsShowSequentialNumber(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_show_sequential_number"}] = v
	return v
}

func (r *Request) Meeting_MotionsStatuteRecommendationsBy(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "motions_statute_recommendations_by"}] = v
	return v
}

func (r *Request) Meeting_MotionsStatutesEnabled(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "motions_statutes_enabled"}] = v
	return v
}

func (r *Request) Meeting_MotionsSupportersMinAmount(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "motions_supporters_min_amount"}] = v
	return v
}

func (r *Request) Meeting_Name(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "name"}] = v
	return v
}

func (r *Request) Meeting_OptionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "option_ids"}] = v
	return v
}

func (r *Request) Meeting_OrganizationTagIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "organization_tag_ids"}] = v
	return v
}

func (r *Request) Meeting_PersonalNoteIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "personal_note_ids"}] = v
	return v
}

func (r *Request) Meeting_PollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "poll_ballot_paper_number"}] = v
	return v
}

func (r *Request) Meeting_PollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "poll_ballot_paper_selection"}] = v
	return v
}

func (r *Request) Meeting_PollCountdownID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"meeting", meetingID, "poll_countdown_id"}] = v
	return v
}

func (r *Request) Meeting_PollCoupleCountdown(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "poll_couple_countdown"}] = v
	return v
}

func (r *Request) Meeting_PollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "poll_default_100_percent_base"}] = v
	return v
}

func (r *Request) Meeting_PollDefaultBackend(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "poll_default_backend"}] = v
	return v
}

func (r *Request) Meeting_PollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "poll_default_group_ids"}] = v
	return v
}

func (r *Request) Meeting_PollDefaultMethod(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "poll_default_method"}] = v
	return v
}

func (r *Request) Meeting_PollDefaultType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "poll_default_type"}] = v
	return v
}

func (r *Request) Meeting_PollIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "poll_ids"}] = v
	return v
}

func (r *Request) Meeting_PollSortPollResultByVotes(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "poll_sort_poll_result_by_votes"}] = v
	return v
}

func (r *Request) Meeting_PresentUserIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "present_user_ids"}] = v
	return v
}

func (r *Request) Meeting_ProjectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "projection_ids"}] = v
	return v
}

func (r *Request) Meeting_ProjectorCountdownDefaultTime(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "projector_countdown_default_time"}] = v
	return v
}

func (r *Request) Meeting_ProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "projector_countdown_ids"}] = v
	return v
}

func (r *Request) Meeting_ProjectorCountdownWarningTime(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "projector_countdown_warning_time"}] = v
	return v
}

func (r *Request) Meeting_ProjectorIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "projector_ids"}] = v
	return v
}

func (r *Request) Meeting_ProjectorMessageIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "projector_message_ids"}] = v
	return v
}

func (r *Request) Meeting_ReferenceProjectorID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "reference_projector_id"}] = v
	return v
}

func (r *Request) Meeting_SpeakerIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "speaker_ids"}] = v
	return v
}

func (r *Request) Meeting_StartTime(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"meeting", meetingID, "start_time"}] = v
	return v
}

func (r *Request) Meeting_TagIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "tag_ids"}] = v
	return v
}

func (r *Request) Meeting_TemplateForOrganizationID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"meeting", meetingID, "template_for_organization_id"}] = v
	return v
}

func (r *Request) Meeting_TopicIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "topic_ids"}] = v
	return v
}

func (r *Request) Meeting_UserIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "user_ids"}] = v
	return v
}

func (r *Request) Meeting_UsersAllowSelfSetPresent(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "users_allow_self_set_present"}] = v
	return v
}

func (r *Request) Meeting_UsersEmailBody(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_email_body"}] = v
	return v
}

func (r *Request) Meeting_UsersEmailReplyto(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_email_replyto"}] = v
	return v
}

func (r *Request) Meeting_UsersEmailSender(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_email_sender"}] = v
	return v
}

func (r *Request) Meeting_UsersEmailSubject(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_email_subject"}] = v
	return v
}

func (r *Request) Meeting_UsersEnablePresenceView(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "users_enable_presence_view"}] = v
	return v
}

func (r *Request) Meeting_UsersEnableVoteWeight(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"meeting", meetingID, "users_enable_vote_weight"}] = v
	return v
}

func (r *Request) Meeting_UsersPdfWelcometext(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_pdf_welcometext"}] = v
	return v
}

func (r *Request) Meeting_UsersPdfWelcometitle(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_pdf_welcometitle"}] = v
	return v
}

func (r *Request) Meeting_UsersPdfWlanEncryption(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_pdf_wlan_encryption"}] = v
	return v
}

func (r *Request) Meeting_UsersPdfWlanPassword(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_pdf_wlan_password"}] = v
	return v
}

func (r *Request) Meeting_UsersPdfWlanSsid(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_pdf_wlan_ssid"}] = v
	return v
}

func (r *Request) Meeting_UsersSortBy(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "users_sort_by"}] = v
	return v
}

func (r *Request) Meeting_VoteIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"meeting", meetingID, "vote_ids"}] = v
	return v
}

func (r *Request) Meeting_WelcomeText(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "welcome_text"}] = v
	return v
}

func (r *Request) Meeting_WelcomeTitle(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"meeting", meetingID, "welcome_title"}] = v
	return v
}

func (r *Request) MotionBlock_AgendaItemID(motionBlockID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion_block", motionBlockID, "agenda_item_id"}] = v
	return v
}

func (r *Request) MotionBlock_ID(motionBlockID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_block", motionBlockID, "id"}] = v
	return v
}

func (r *Request) MotionBlock_Internal(motionBlockID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_block", motionBlockID, "internal"}] = v
	return v
}

func (r *Request) MotionBlock_ListOfSpeakersID(motionBlockID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_block", motionBlockID, "list_of_speakers_id"}] = v
	return v
}

func (r *Request) MotionBlock_MeetingID(motionBlockID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_block", motionBlockID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionBlock_MotionIDs(motionBlockID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_block", motionBlockID, "motion_ids"}] = v
	return v
}

func (r *Request) MotionBlock_ProjectionIDs(motionBlockID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_block", motionBlockID, "projection_ids"}] = v
	return v
}

func (r *Request) MotionBlock_SequentialNumber(motionBlockID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_block", motionBlockID, "sequential_number"}] = v
	return v
}

func (r *Request) MotionBlock_Title(motionBlockID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_block", motionBlockID, "title"}] = v
	return v
}

func (r *Request) MotionCategory_ChildIDs(motionCategoryID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "child_ids"}] = v
	return v
}

func (r *Request) MotionCategory_ID(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "id"}] = v
	return v
}

func (r *Request) MotionCategory_Level(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "level"}] = v
	return v
}

func (r *Request) MotionCategory_MeetingID(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionCategory_MotionIDs(motionCategoryID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "motion_ids"}] = v
	return v
}

func (r *Request) MotionCategory_Name(motionCategoryID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "name"}] = v
	return v
}

func (r *Request) MotionCategory_ParentID(motionCategoryID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "parent_id"}] = v
	return v
}

func (r *Request) MotionCategory_Prefix(motionCategoryID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "prefix"}] = v
	return v
}

func (r *Request) MotionCategory_SequentialNumber(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "sequential_number"}] = v
	return v
}

func (r *Request) MotionCategory_Weight(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_category", motionCategoryID, "weight"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_CreationTime(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "creation_time"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_ID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "id"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_Internal(motionChangeRecommendationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "internal"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_LineFrom(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "line_from"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_LineTo(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "line_to"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_MeetingID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_MotionID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "motion_id"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_OtherDescription(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "other_description"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_Rejected(motionChangeRecommendationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "rejected"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_Text(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "text"}] = v
	return v
}

func (r *Request) MotionChangeRecommendation_Type(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_change_recommendation", motionChangeRecommendationID, "type"}] = v
	return v
}

func (r *Request) MotionCommentSection_CommentIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_comment_section", motionCommentSectionID, "comment_ids"}] = v
	return v
}

func (r *Request) MotionCommentSection_ID(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_comment_section", motionCommentSectionID, "id"}] = v
	return v
}

func (r *Request) MotionCommentSection_MeetingID(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_comment_section", motionCommentSectionID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionCommentSection_Name(motionCommentSectionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_comment_section", motionCommentSectionID, "name"}] = v
	return v
}

func (r *Request) MotionCommentSection_ReadGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_comment_section", motionCommentSectionID, "read_group_ids"}] = v
	return v
}

func (r *Request) MotionCommentSection_SequentialNumber(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_comment_section", motionCommentSectionID, "sequential_number"}] = v
	return v
}

func (r *Request) MotionCommentSection_Weight(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_comment_section", motionCommentSectionID, "weight"}] = v
	return v
}

func (r *Request) MotionCommentSection_WriteGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_comment_section", motionCommentSectionID, "write_group_ids"}] = v
	return v
}

func (r *Request) MotionComment_Comment(motionCommentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_comment", motionCommentID, "comment"}] = v
	return v
}

func (r *Request) MotionComment_ID(motionCommentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_comment", motionCommentID, "id"}] = v
	return v
}

func (r *Request) MotionComment_MeetingID(motionCommentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_comment", motionCommentID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionComment_MotionID(motionCommentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_comment", motionCommentID, "motion_id"}] = v
	return v
}

func (r *Request) MotionComment_SectionID(motionCommentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_comment", motionCommentID, "section_id"}] = v
	return v
}

func (r *Request) MotionState_AllowCreatePoll(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_state", motionStateID, "allow_create_poll"}] = v
	return v
}

func (r *Request) MotionState_AllowMotionForwarding(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_state", motionStateID, "allow_motion_forwarding"}] = v
	return v
}

func (r *Request) MotionState_AllowSubmitterEdit(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_state", motionStateID, "allow_submitter_edit"}] = v
	return v
}

func (r *Request) MotionState_AllowSupport(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_state", motionStateID, "allow_support"}] = v
	return v
}

func (r *Request) MotionState_CssClass(motionStateID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_state", motionStateID, "css_class"}] = v
	return v
}

func (r *Request) MotionState_FirstStateOfWorkflowID(motionStateID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion_state", motionStateID, "first_state_of_workflow_id"}] = v
	return v
}

func (r *Request) MotionState_ID(motionStateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_state", motionStateID, "id"}] = v
	return v
}

func (r *Request) MotionState_MeetingID(motionStateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_state", motionStateID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionState_MergeAmendmentIntoFinal(motionStateID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_state", motionStateID, "merge_amendment_into_final"}] = v
	return v
}

func (r *Request) MotionState_MotionIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_state", motionStateID, "motion_ids"}] = v
	return v
}

func (r *Request) MotionState_MotionRecommendationIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_state", motionStateID, "motion_recommendation_ids"}] = v
	return v
}

func (r *Request) MotionState_Name(motionStateID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_state", motionStateID, "name"}] = v
	return v
}

func (r *Request) MotionState_NextStateIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_state", motionStateID, "next_state_ids"}] = v
	return v
}

func (r *Request) MotionState_PreviousStateIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_state", motionStateID, "previous_state_ids"}] = v
	return v
}

func (r *Request) MotionState_RecommendationLabel(motionStateID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_state", motionStateID, "recommendation_label"}] = v
	return v
}

func (r *Request) MotionState_Restrictions(motionStateID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"motion_state", motionStateID, "restrictions"}] = v
	return v
}

func (r *Request) MotionState_SetCreatedTimestamp(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_state", motionStateID, "set_created_timestamp"}] = v
	return v
}

func (r *Request) MotionState_SetNumber(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_state", motionStateID, "set_number"}] = v
	return v
}

func (r *Request) MotionState_ShowRecommendationExtensionField(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_state", motionStateID, "show_recommendation_extension_field"}] = v
	return v
}

func (r *Request) MotionState_ShowStateExtensionField(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"motion_state", motionStateID, "show_state_extension_field"}] = v
	return v
}

func (r *Request) MotionState_Weight(motionStateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_state", motionStateID, "weight"}] = v
	return v
}

func (r *Request) MotionState_WorkflowID(motionStateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_state", motionStateID, "workflow_id"}] = v
	return v
}

func (r *Request) MotionStatuteParagraph_ID(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_statute_paragraph", motionStatuteParagraphID, "id"}] = v
	return v
}

func (r *Request) MotionStatuteParagraph_MeetingID(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_statute_paragraph", motionStatuteParagraphID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionStatuteParagraph_MotionIDs(motionStatuteParagraphID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_statute_paragraph", motionStatuteParagraphID, "motion_ids"}] = v
	return v
}

func (r *Request) MotionStatuteParagraph_SequentialNumber(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_statute_paragraph", motionStatuteParagraphID, "sequential_number"}] = v
	return v
}

func (r *Request) MotionStatuteParagraph_Text(motionStatuteParagraphID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_statute_paragraph", motionStatuteParagraphID, "text"}] = v
	return v
}

func (r *Request) MotionStatuteParagraph_Title(motionStatuteParagraphID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_statute_paragraph", motionStatuteParagraphID, "title"}] = v
	return v
}

func (r *Request) MotionStatuteParagraph_Weight(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_statute_paragraph", motionStatuteParagraphID, "weight"}] = v
	return v
}

func (r *Request) MotionSubmitter_ID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_submitter", motionSubmitterID, "id"}] = v
	return v
}

func (r *Request) MotionSubmitter_MeetingID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_submitter", motionSubmitterID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionSubmitter_MotionID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_submitter", motionSubmitterID, "motion_id"}] = v
	return v
}

func (r *Request) MotionSubmitter_UserID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_submitter", motionSubmitterID, "user_id"}] = v
	return v
}

func (r *Request) MotionSubmitter_Weight(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_submitter", motionSubmitterID, "weight"}] = v
	return v
}

func (r *Request) MotionWorkflow_DefaultAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "default_amendment_workflow_meeting_id"}] = v
	return v
}

func (r *Request) MotionWorkflow_DefaultStatuteAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "default_statute_amendment_workflow_meeting_id"}] = v
	return v
}

func (r *Request) MotionWorkflow_DefaultWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "default_workflow_meeting_id"}] = v
	return v
}

func (r *Request) MotionWorkflow_FirstStateID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "first_state_id"}] = v
	return v
}

func (r *Request) MotionWorkflow_ID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "id"}] = v
	return v
}

func (r *Request) MotionWorkflow_MeetingID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "meeting_id"}] = v
	return v
}

func (r *Request) MotionWorkflow_Name(motionWorkflowID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "name"}] = v
	return v
}

func (r *Request) MotionWorkflow_SequentialNumber(motionWorkflowID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "sequential_number"}] = v
	return v
}

func (r *Request) MotionWorkflow_StateIDs(motionWorkflowID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion_workflow", motionWorkflowID, "state_ids"}] = v
	return v
}

func (r *Request) Motion_AgendaItemID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion", motionID, "agenda_item_id"}] = v
	return v
}

func (r *Request) Motion_AllDerivedMotionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "all_derived_motion_ids"}] = v
	return v
}

func (r *Request) Motion_AllOriginIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "all_origin_ids"}] = v
	return v
}

func (r *Request) Motion_AmendmentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "amendment_ids"}] = v
	return v
}

func (r *Request) Motion_AmendmentParagraphTmpl(motionID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"motion", motionID, "amendment_paragraph_$"}] = v
	return v
}

func (r *Request) Motion_AmendmentParagraph(motionID int, replacement string) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion", motionID, fmt.Sprintf("amendment_paragraph_$%s", replacement)}] = v
	return v
}

func (r *Request) Motion_AttachmentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "attachment_ids"}] = v
	return v
}

func (r *Request) Motion_BlockID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion", motionID, "block_id"}] = v
	return v
}

func (r *Request) Motion_CategoryID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion", motionID, "category_id"}] = v
	return v
}

func (r *Request) Motion_CategoryWeight(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "category_weight"}] = v
	return v
}

func (r *Request) Motion_ChangeRecommendationIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "change_recommendation_ids"}] = v
	return v
}

func (r *Request) Motion_CommentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "comment_ids"}] = v
	return v
}

func (r *Request) Motion_Created(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "created"}] = v
	return v
}

func (r *Request) Motion_DerivedMotionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "derived_motion_ids"}] = v
	return v
}

func (r *Request) Motion_Forwarded(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "forwarded"}] = v
	return v
}

func (r *Request) Motion_ID(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "id"}] = v
	return v
}

func (r *Request) Motion_LastModified(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "last_modified"}] = v
	return v
}

func (r *Request) Motion_LeadMotionID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion", motionID, "lead_motion_id"}] = v
	return v
}

func (r *Request) Motion_ListOfSpeakersID(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "list_of_speakers_id"}] = v
	return v
}

func (r *Request) Motion_MeetingID(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "meeting_id"}] = v
	return v
}

func (r *Request) Motion_ModifiedFinalVersion(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion", motionID, "modified_final_version"}] = v
	return v
}

func (r *Request) Motion_Number(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion", motionID, "number"}] = v
	return v
}

func (r *Request) Motion_NumberValue(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "number_value"}] = v
	return v
}

func (r *Request) Motion_OptionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "option_ids"}] = v
	return v
}

func (r *Request) Motion_OriginID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion", motionID, "origin_id"}] = v
	return v
}

func (r *Request) Motion_PersonalNoteIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "personal_note_ids"}] = v
	return v
}

func (r *Request) Motion_PollIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "poll_ids"}] = v
	return v
}

func (r *Request) Motion_ProjectionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "projection_ids"}] = v
	return v
}

func (r *Request) Motion_Reason(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion", motionID, "reason"}] = v
	return v
}

func (r *Request) Motion_RecommendationExtension(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion", motionID, "recommendation_extension"}] = v
	return v
}

func (r *Request) Motion_RecommendationExtensionReferenceIDs(motionID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"motion", motionID, "recommendation_extension_reference_ids"}] = v
	return v
}

func (r *Request) Motion_RecommendationID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion", motionID, "recommendation_id"}] = v
	return v
}

func (r *Request) Motion_ReferencedInMotionRecommendationExtensionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "referenced_in_motion_recommendation_extension_ids"}] = v
	return v
}

func (r *Request) Motion_SequentialNumber(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "sequential_number"}] = v
	return v
}

func (r *Request) Motion_SortChildIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "sort_child_ids"}] = v
	return v
}

func (r *Request) Motion_SortParentID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion", motionID, "sort_parent_id"}] = v
	return v
}

func (r *Request) Motion_SortWeight(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "sort_weight"}] = v
	return v
}

func (r *Request) Motion_StartLineNumber(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "start_line_number"}] = v
	return v
}

func (r *Request) Motion_StateExtension(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion", motionID, "state_extension"}] = v
	return v
}

func (r *Request) Motion_StateID(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"motion", motionID, "state_id"}] = v
	return v
}

func (r *Request) Motion_StatuteParagraphID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"motion", motionID, "statute_paragraph_id"}] = v
	return v
}

func (r *Request) Motion_SubmitterIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "submitter_ids"}] = v
	return v
}

func (r *Request) Motion_SupporterIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "supporter_ids"}] = v
	return v
}

func (r *Request) Motion_TagIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"motion", motionID, "tag_ids"}] = v
	return v
}

func (r *Request) Motion_Text(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion", motionID, "text"}] = v
	return v
}

func (r *Request) Motion_Title(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"motion", motionID, "title"}] = v
	return v
}

func (r *Request) Option_Abstain(optionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"option", optionID, "abstain"}] = v
	return v
}

func (r *Request) Option_ContentObjectID(optionID int) *ValueMaybeString {
	v := &ValueMaybeString{request: r}
	r.requested[Key{"option", optionID, "content_object_id"}] = v
	return v
}

func (r *Request) Option_ID(optionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"option", optionID, "id"}] = v
	return v
}

func (r *Request) Option_MeetingID(optionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"option", optionID, "meeting_id"}] = v
	return v
}

func (r *Request) Option_No(optionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"option", optionID, "no"}] = v
	return v
}

func (r *Request) Option_PollID(optionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"option", optionID, "poll_id"}] = v
	return v
}

func (r *Request) Option_Text(optionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"option", optionID, "text"}] = v
	return v
}

func (r *Request) Option_UsedAsGlobalOptionInPollID(optionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"option", optionID, "used_as_global_option_in_poll_id"}] = v
	return v
}

func (r *Request) Option_VoteIDs(optionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"option", optionID, "vote_ids"}] = v
	return v
}

func (r *Request) Option_Weight(optionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"option", optionID, "weight"}] = v
	return v
}

func (r *Request) Option_Yes(optionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"option", optionID, "yes"}] = v
	return v
}

func (r *Request) OrganizationTag_Color(organizationTagID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization_tag", organizationTagID, "color"}] = v
	return v
}

func (r *Request) OrganizationTag_ID(organizationTagID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"organization_tag", organizationTagID, "id"}] = v
	return v
}

func (r *Request) OrganizationTag_Name(organizationTagID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization_tag", organizationTagID, "name"}] = v
	return v
}

func (r *Request) OrganizationTag_OrganizationID(organizationTagID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"organization_tag", organizationTagID, "organization_id"}] = v
	return v
}

func (r *Request) OrganizationTag_TaggedIDs(organizationTagID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"organization_tag", organizationTagID, "tagged_ids"}] = v
	return v
}

func (r *Request) Organization_ActiveMeetingIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"organization", organizationID, "active_meeting_ids"}] = v
	return v
}

func (r *Request) Organization_ArchivedMeetingIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"organization", organizationID, "archived_meeting_ids"}] = v
	return v
}

func (r *Request) Organization_CommitteeIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"organization", organizationID, "committee_ids"}] = v
	return v
}

func (r *Request) Organization_Description(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "description"}] = v
	return v
}

func (r *Request) Organization_EnableChat(organizationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"organization", organizationID, "enable_chat"}] = v
	return v
}

func (r *Request) Organization_EnableElectronicVoting(organizationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"organization", organizationID, "enable_electronic_voting"}] = v
	return v
}

func (r *Request) Organization_ID(organizationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"organization", organizationID, "id"}] = v
	return v
}

func (r *Request) Organization_LegalNotice(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "legal_notice"}] = v
	return v
}

func (r *Request) Organization_LimitOfMeetings(organizationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"organization", organizationID, "limit_of_meetings"}] = v
	return v
}

func (r *Request) Organization_LimitOfUsers(organizationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"organization", organizationID, "limit_of_users"}] = v
	return v
}

func (r *Request) Organization_LoginText(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "login_text"}] = v
	return v
}

func (r *Request) Organization_MediafileIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"organization", organizationID, "mediafile_ids"}] = v
	return v
}

func (r *Request) Organization_Name(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "name"}] = v
	return v
}

func (r *Request) Organization_OrganizationTagIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"organization", organizationID, "organization_tag_ids"}] = v
	return v
}

func (r *Request) Organization_PrivacyPolicy(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "privacy_policy"}] = v
	return v
}

func (r *Request) Organization_ResetPasswordVerboseErrors(organizationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"organization", organizationID, "reset_password_verbose_errors"}] = v
	return v
}

func (r *Request) Organization_TemplateMeetingIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"organization", organizationID, "template_meeting_ids"}] = v
	return v
}

func (r *Request) Organization_ThemeID(organizationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"organization", organizationID, "theme_id"}] = v
	return v
}

func (r *Request) Organization_ThemeIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"organization", organizationID, "theme_ids"}] = v
	return v
}

func (r *Request) Organization_Url(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "url"}] = v
	return v
}

func (r *Request) Organization_UsersEmailBody(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "users_email_body"}] = v
	return v
}

func (r *Request) Organization_UsersEmailReplyto(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "users_email_replyto"}] = v
	return v
}

func (r *Request) Organization_UsersEmailSender(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "users_email_sender"}] = v
	return v
}

func (r *Request) Organization_UsersEmailSubject(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"organization", organizationID, "users_email_subject"}] = v
	return v
}

func (r *Request) PersonalNote_ContentObjectID(personalNoteID int) *ValueMaybeString {
	v := &ValueMaybeString{request: r}
	r.requested[Key{"personal_note", personalNoteID, "content_object_id"}] = v
	return v
}

func (r *Request) PersonalNote_ID(personalNoteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"personal_note", personalNoteID, "id"}] = v
	return v
}

func (r *Request) PersonalNote_MeetingID(personalNoteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"personal_note", personalNoteID, "meeting_id"}] = v
	return v
}

func (r *Request) PersonalNote_Note(personalNoteID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"personal_note", personalNoteID, "note"}] = v
	return v
}

func (r *Request) PersonalNote_Star(personalNoteID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"personal_note", personalNoteID, "star"}] = v
	return v
}

func (r *Request) PersonalNote_UserID(personalNoteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"personal_note", personalNoteID, "user_id"}] = v
	return v
}

func (r *Request) Poll_Backend(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "backend"}] = v
	return v
}

func (r *Request) Poll_ContentObjectID(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "content_object_id"}] = v
	return v
}

func (r *Request) Poll_Description(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "description"}] = v
	return v
}

func (r *Request) Poll_EntitledGroupIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"poll", pollID, "entitled_group_ids"}] = v
	return v
}

func (r *Request) Poll_EntitledUsersAtStop(pollID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[Key{"poll", pollID, "entitled_users_at_stop"}] = v
	return v
}

func (r *Request) Poll_GlobalAbstain(pollID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"poll", pollID, "global_abstain"}] = v
	return v
}

func (r *Request) Poll_GlobalNo(pollID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"poll", pollID, "global_no"}] = v
	return v
}

func (r *Request) Poll_GlobalOptionID(pollID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"poll", pollID, "global_option_id"}] = v
	return v
}

func (r *Request) Poll_GlobalYes(pollID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"poll", pollID, "global_yes"}] = v
	return v
}

func (r *Request) Poll_ID(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"poll", pollID, "id"}] = v
	return v
}

func (r *Request) Poll_IsPseudoanonymized(pollID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"poll", pollID, "is_pseudoanonymized"}] = v
	return v
}

func (r *Request) Poll_MaxVotesAmount(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"poll", pollID, "max_votes_amount"}] = v
	return v
}

func (r *Request) Poll_MaxVotesPerOption(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"poll", pollID, "max_votes_per_option"}] = v
	return v
}

func (r *Request) Poll_MeetingID(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"poll", pollID, "meeting_id"}] = v
	return v
}

func (r *Request) Poll_MinVotesAmount(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"poll", pollID, "min_votes_amount"}] = v
	return v
}

func (r *Request) Poll_OnehundredPercentBase(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "onehundred_percent_base"}] = v
	return v
}

func (r *Request) Poll_OptionIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"poll", pollID, "option_ids"}] = v
	return v
}

func (r *Request) Poll_Pollmethod(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "pollmethod"}] = v
	return v
}

func (r *Request) Poll_ProjectionIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"poll", pollID, "projection_ids"}] = v
	return v
}

func (r *Request) Poll_SequentialNumber(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"poll", pollID, "sequential_number"}] = v
	return v
}

func (r *Request) Poll_State(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "state"}] = v
	return v
}

func (r *Request) Poll_Title(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "title"}] = v
	return v
}

func (r *Request) Poll_Type(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "type"}] = v
	return v
}

func (r *Request) Poll_VoteCount(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"poll", pollID, "vote_count"}] = v
	return v
}

func (r *Request) Poll_VotedIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"poll", pollID, "voted_ids"}] = v
	return v
}

func (r *Request) Poll_Votescast(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "votescast"}] = v
	return v
}

func (r *Request) Poll_Votesinvalid(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "votesinvalid"}] = v
	return v
}

func (r *Request) Poll_Votesvalid(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"poll", pollID, "votesvalid"}] = v
	return v
}

func (r *Request) Projection_Content(projectionID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[Key{"projection", projectionID, "content"}] = v
	return v
}

func (r *Request) Projection_ContentObjectID(projectionID int) *ValueMaybeString {
	v := &ValueMaybeString{request: r}
	r.requested[Key{"projection", projectionID, "content_object_id"}] = v
	return v
}

func (r *Request) Projection_CurrentProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"projection", projectionID, "current_projector_id"}] = v
	return v
}

func (r *Request) Projection_HistoryProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"projection", projectionID, "history_projector_id"}] = v
	return v
}

func (r *Request) Projection_ID(projectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projection", projectionID, "id"}] = v
	return v
}

func (r *Request) Projection_MeetingID(projectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projection", projectionID, "meeting_id"}] = v
	return v
}

func (r *Request) Projection_Options(projectionID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[Key{"projection", projectionID, "options"}] = v
	return v
}

func (r *Request) Projection_PreviewProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"projection", projectionID, "preview_projector_id"}] = v
	return v
}

func (r *Request) Projection_Stable(projectionID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"projection", projectionID, "stable"}] = v
	return v
}

func (r *Request) Projection_Type(projectionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projection", projectionID, "type"}] = v
	return v
}

func (r *Request) Projection_Weight(projectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projection", projectionID, "weight"}] = v
	return v
}

func (r *Request) ProjectorCountdown_CountdownTime(projectorCountdownID int) *ValueFloat {
	v := &ValueFloat{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "countdown_time"}] = v
	return v
}

func (r *Request) ProjectorCountdown_DefaultTime(projectorCountdownID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "default_time"}] = v
	return v
}

func (r *Request) ProjectorCountdown_Description(projectorCountdownID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "description"}] = v
	return v
}

func (r *Request) ProjectorCountdown_ID(projectorCountdownID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "id"}] = v
	return v
}

func (r *Request) ProjectorCountdown_MeetingID(projectorCountdownID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "meeting_id"}] = v
	return v
}

func (r *Request) ProjectorCountdown_ProjectionIDs(projectorCountdownID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "projection_ids"}] = v
	return v
}

func (r *Request) ProjectorCountdown_Running(projectorCountdownID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "running"}] = v
	return v
}

func (r *Request) ProjectorCountdown_Title(projectorCountdownID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "title"}] = v
	return v
}

func (r *Request) ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "used_as_list_of_speakers_countdown_meeting_id"}] = v
	return v
}

func (r *Request) ProjectorCountdown_UsedAsPollCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"projector_countdown", projectorCountdownID, "used_as_poll_countdown_meeting_id"}] = v
	return v
}

func (r *Request) ProjectorMessage_ID(projectorMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector_message", projectorMessageID, "id"}] = v
	return v
}

func (r *Request) ProjectorMessage_MeetingID(projectorMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector_message", projectorMessageID, "meeting_id"}] = v
	return v
}

func (r *Request) ProjectorMessage_Message(projectorMessageID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector_message", projectorMessageID, "message"}] = v
	return v
}

func (r *Request) ProjectorMessage_ProjectionIDs(projectorMessageID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"projector_message", projectorMessageID, "projection_ids"}] = v
	return v
}

func (r *Request) Projector_AspectRatioDenominator(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, "aspect_ratio_denominator"}] = v
	return v
}

func (r *Request) Projector_AspectRatioNumerator(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, "aspect_ratio_numerator"}] = v
	return v
}

func (r *Request) Projector_BackgroundColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector", projectorID, "background_color"}] = v
	return v
}

func (r *Request) Projector_ChyronBackgroundColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector", projectorID, "chyron_background_color"}] = v
	return v
}

func (r *Request) Projector_ChyronFontColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector", projectorID, "chyron_font_color"}] = v
	return v
}

func (r *Request) Projector_Color(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector", projectorID, "color"}] = v
	return v
}

func (r *Request) Projector_CurrentProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"projector", projectorID, "current_projection_ids"}] = v
	return v
}

func (r *Request) Projector_HeaderBackgroundColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector", projectorID, "header_background_color"}] = v
	return v
}

func (r *Request) Projector_HeaderFontColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector", projectorID, "header_font_color"}] = v
	return v
}

func (r *Request) Projector_HeaderH1Color(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector", projectorID, "header_h1_color"}] = v
	return v
}

func (r *Request) Projector_HistoryProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"projector", projectorID, "history_projection_ids"}] = v
	return v
}

func (r *Request) Projector_ID(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, "id"}] = v
	return v
}

func (r *Request) Projector_MeetingID(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, "meeting_id"}] = v
	return v
}

func (r *Request) Projector_Name(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"projector", projectorID, "name"}] = v
	return v
}

func (r *Request) Projector_PreviewProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"projector", projectorID, "preview_projection_ids"}] = v
	return v
}

func (r *Request) Projector_Scale(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, "scale"}] = v
	return v
}

func (r *Request) Projector_Scroll(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, "scroll"}] = v
	return v
}

func (r *Request) Projector_SequentialNumber(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, "sequential_number"}] = v
	return v
}

func (r *Request) Projector_ShowClock(projectorID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"projector", projectorID, "show_clock"}] = v
	return v
}

func (r *Request) Projector_ShowHeaderFooter(projectorID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"projector", projectorID, "show_header_footer"}] = v
	return v
}

func (r *Request) Projector_ShowLogo(projectorID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"projector", projectorID, "show_logo"}] = v
	return v
}

func (r *Request) Projector_ShowTitle(projectorID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"projector", projectorID, "show_title"}] = v
	return v
}

func (r *Request) Projector_UsedAsDefaultInMeetingIDTmpl(projectorID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"projector", projectorID, "used_as_default_$_in_meeting_id"}] = v
	return v
}

func (r *Request) Projector_UsedAsDefaultInMeetingID(projectorID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, fmt.Sprintf("used_as_default_$%s_in_meeting_id", replacement)}] = v
	return v
}

func (r *Request) Projector_UsedAsReferenceProjectorMeetingID(projectorID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"projector", projectorID, "used_as_reference_projector_meeting_id"}] = v
	return v
}

func (r *Request) Projector_Width(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"projector", projectorID, "width"}] = v
	return v
}

func (r *Request) Speaker_BeginTime(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"speaker", speakerID, "begin_time"}] = v
	return v
}

func (r *Request) Speaker_EndTime(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"speaker", speakerID, "end_time"}] = v
	return v
}

func (r *Request) Speaker_ID(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"speaker", speakerID, "id"}] = v
	return v
}

func (r *Request) Speaker_ListOfSpeakersID(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"speaker", speakerID, "list_of_speakers_id"}] = v
	return v
}

func (r *Request) Speaker_MeetingID(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"speaker", speakerID, "meeting_id"}] = v
	return v
}

func (r *Request) Speaker_Note(speakerID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"speaker", speakerID, "note"}] = v
	return v
}

func (r *Request) Speaker_PointOfOrder(speakerID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"speaker", speakerID, "point_of_order"}] = v
	return v
}

func (r *Request) Speaker_SpeechState(speakerID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"speaker", speakerID, "speech_state"}] = v
	return v
}

func (r *Request) Speaker_UserID(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"speaker", speakerID, "user_id"}] = v
	return v
}

func (r *Request) Speaker_Weight(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"speaker", speakerID, "weight"}] = v
	return v
}

func (r *Request) Tag_ID(tagID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"tag", tagID, "id"}] = v
	return v
}

func (r *Request) Tag_MeetingID(tagID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"tag", tagID, "meeting_id"}] = v
	return v
}

func (r *Request) Tag_Name(tagID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"tag", tagID, "name"}] = v
	return v
}

func (r *Request) Tag_TaggedIDs(tagID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"tag", tagID, "tagged_ids"}] = v
	return v
}

func (r *Request) Theme_Accent100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_100"}] = v
	return v
}

func (r *Request) Theme_Accent200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_200"}] = v
	return v
}

func (r *Request) Theme_Accent300(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_300"}] = v
	return v
}

func (r *Request) Theme_Accent400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_400"}] = v
	return v
}

func (r *Request) Theme_Accent50(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_50"}] = v
	return v
}

func (r *Request) Theme_Accent500(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_500"}] = v
	return v
}

func (r *Request) Theme_Accent600(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_600"}] = v
	return v
}

func (r *Request) Theme_Accent700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_700"}] = v
	return v
}

func (r *Request) Theme_Accent800(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_800"}] = v
	return v
}

func (r *Request) Theme_Accent900(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_900"}] = v
	return v
}

func (r *Request) Theme_AccentA100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_a100"}] = v
	return v
}

func (r *Request) Theme_AccentA200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_a200"}] = v
	return v
}

func (r *Request) Theme_AccentA400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_a400"}] = v
	return v
}

func (r *Request) Theme_AccentA700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "accent_a700"}] = v
	return v
}

func (r *Request) Theme_ID(themeID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"theme", themeID, "id"}] = v
	return v
}

func (r *Request) Theme_Name(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "name"}] = v
	return v
}

func (r *Request) Theme_OrganizationID(themeID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"theme", themeID, "organization_id"}] = v
	return v
}

func (r *Request) Theme_Primary100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_100"}] = v
	return v
}

func (r *Request) Theme_Primary200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_200"}] = v
	return v
}

func (r *Request) Theme_Primary300(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_300"}] = v
	return v
}

func (r *Request) Theme_Primary400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_400"}] = v
	return v
}

func (r *Request) Theme_Primary50(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_50"}] = v
	return v
}

func (r *Request) Theme_Primary500(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_500"}] = v
	return v
}

func (r *Request) Theme_Primary600(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_600"}] = v
	return v
}

func (r *Request) Theme_Primary700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_700"}] = v
	return v
}

func (r *Request) Theme_Primary800(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_800"}] = v
	return v
}

func (r *Request) Theme_Primary900(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_900"}] = v
	return v
}

func (r *Request) Theme_PrimaryA100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_a100"}] = v
	return v
}

func (r *Request) Theme_PrimaryA200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_a200"}] = v
	return v
}

func (r *Request) Theme_PrimaryA400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_a400"}] = v
	return v
}

func (r *Request) Theme_PrimaryA700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "primary_a700"}] = v
	return v
}

func (r *Request) Theme_ThemeForOrganizationID(themeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"theme", themeID, "theme_for_organization_id"}] = v
	return v
}

func (r *Request) Theme_Warn100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_100"}] = v
	return v
}

func (r *Request) Theme_Warn200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_200"}] = v
	return v
}

func (r *Request) Theme_Warn300(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_300"}] = v
	return v
}

func (r *Request) Theme_Warn400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_400"}] = v
	return v
}

func (r *Request) Theme_Warn50(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_50"}] = v
	return v
}

func (r *Request) Theme_Warn500(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_500"}] = v
	return v
}

func (r *Request) Theme_Warn600(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_600"}] = v
	return v
}

func (r *Request) Theme_Warn700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_700"}] = v
	return v
}

func (r *Request) Theme_Warn800(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_800"}] = v
	return v
}

func (r *Request) Theme_Warn900(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_900"}] = v
	return v
}

func (r *Request) Theme_WarnA100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_a100"}] = v
	return v
}

func (r *Request) Theme_WarnA200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_a200"}] = v
	return v
}

func (r *Request) Theme_WarnA400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_a400"}] = v
	return v
}

func (r *Request) Theme_WarnA700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"theme", themeID, "warn_a700"}] = v
	return v
}

func (r *Request) Topic_AgendaItemID(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"topic", topicID, "agenda_item_id"}] = v
	return v
}

func (r *Request) Topic_AttachmentIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"topic", topicID, "attachment_ids"}] = v
	return v
}

func (r *Request) Topic_ID(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"topic", topicID, "id"}] = v
	return v
}

func (r *Request) Topic_ListOfSpeakersID(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"topic", topicID, "list_of_speakers_id"}] = v
	return v
}

func (r *Request) Topic_MeetingID(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"topic", topicID, "meeting_id"}] = v
	return v
}

func (r *Request) Topic_PollIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"topic", topicID, "poll_ids"}] = v
	return v
}

func (r *Request) Topic_ProjectionIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"topic", topicID, "projection_ids"}] = v
	return v
}

func (r *Request) Topic_SequentialNumber(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"topic", topicID, "sequential_number"}] = v
	return v
}

func (r *Request) Topic_TagIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"topic", topicID, "tag_ids"}] = v
	return v
}

func (r *Request) Topic_Text(topicID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"topic", topicID, "text"}] = v
	return v
}

func (r *Request) Topic_Title(topicID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"topic", topicID, "title"}] = v
	return v
}

func (r *Request) User_AboutMeTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "about_me_$"}] = v
	return v
}

func (r *Request) User_AboutMe(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("about_me_$%d", meetingID)}] = v
	return v
}

func (r *Request) User_AssignmentCandidateIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "assignment_candidate_$_ids"}] = v
	return v
}

func (r *Request) User_AssignmentCandidateIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("assignment_candidate_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_CanChangeOwnPassword(userID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"user", userID, "can_change_own_password"}] = v
	return v
}

func (r *Request) User_ChatMessageIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "chat_message_$_ids"}] = v
	return v
}

func (r *Request) User_ChatMessageIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("chat_message_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_CommentTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "comment_$"}] = v
	return v
}

func (r *Request) User_Comment(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("comment_$%d", meetingID)}] = v
	return v
}

func (r *Request) User_CommitteeIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, "committee_ids"}] = v
	return v
}

func (r *Request) User_CommitteeManagementLevelTmpl(userID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[Key{"user", userID, "committee_$_management_level"}] = v
	return v
}

func (r *Request) User_CommitteeManagementLevel(userID int, replacement string) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("committee_$%s_management_level", replacement)}] = v
	return v
}

func (r *Request) User_DefaultNumber(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "default_number"}] = v
	return v
}

func (r *Request) User_DefaultPassword(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "default_password"}] = v
	return v
}

func (r *Request) User_DefaultStructureLevel(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "default_structure_level"}] = v
	return v
}

func (r *Request) User_DefaultVoteWeight(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "default_vote_weight"}] = v
	return v
}

func (r *Request) User_Email(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "email"}] = v
	return v
}

func (r *Request) User_FirstName(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "first_name"}] = v
	return v
}

func (r *Request) User_ForwardingCommitteeIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, "forwarding_committee_ids"}] = v
	return v
}

func (r *Request) User_Gender(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "gender"}] = v
	return v
}

func (r *Request) User_GroupIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "group_$_ids"}] = v
	return v
}

func (r *Request) User_GroupIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("group_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_ID(userID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"user", userID, "id"}] = v
	return v
}

func (r *Request) User_IsActive(userID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"user", userID, "is_active"}] = v
	return v
}

func (r *Request) User_IsDemoUser(userID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"user", userID, "is_demo_user"}] = v
	return v
}

func (r *Request) User_IsPhysicalPerson(userID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[Key{"user", userID, "is_physical_person"}] = v
	return v
}

func (r *Request) User_IsPresentInMeetingIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, "is_present_in_meeting_ids"}] = v
	return v
}

func (r *Request) User_LastEmailSend(userID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"user", userID, "last_email_send"}] = v
	return v
}

func (r *Request) User_LastName(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "last_name"}] = v
	return v
}

func (r *Request) User_MeetingIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, "meeting_ids"}] = v
	return v
}

func (r *Request) User_NumberTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "number_$"}] = v
	return v
}

func (r *Request) User_Number(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("number_$%d", meetingID)}] = v
	return v
}

func (r *Request) User_OptionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "option_$_ids"}] = v
	return v
}

func (r *Request) User_OptionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("option_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_OrganizationManagementLevel(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "organization_management_level"}] = v
	return v
}

func (r *Request) User_Password(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "password"}] = v
	return v
}

func (r *Request) User_PersonalNoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "personal_note_$_ids"}] = v
	return v
}

func (r *Request) User_PersonalNoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("personal_note_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_PollVotedIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "poll_voted_$_ids"}] = v
	return v
}

func (r *Request) User_PollVotedIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("poll_voted_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_ProjectionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "projection_$_ids"}] = v
	return v
}

func (r *Request) User_ProjectionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("projection_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_Pronoun(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "pronoun"}] = v
	return v
}

func (r *Request) User_SpeakerIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "speaker_$_ids"}] = v
	return v
}

func (r *Request) User_SpeakerIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("speaker_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_StructureLevelTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "structure_level_$"}] = v
	return v
}

func (r *Request) User_StructureLevel(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("structure_level_$%d", meetingID)}] = v
	return v
}

func (r *Request) User_SubmittedMotionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "submitted_motion_$_ids"}] = v
	return v
}

func (r *Request) User_SubmittedMotionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("submitted_motion_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_SupportedMotionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "supported_motion_$_ids"}] = v
	return v
}

func (r *Request) User_SupportedMotionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("supported_motion_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_Title(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "title"}] = v
	return v
}

func (r *Request) User_Username(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, "username"}] = v
	return v
}

func (r *Request) User_VoteDelegatedToIDTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "vote_delegated_$_to_id"}] = v
	return v
}

func (r *Request) User_VoteDelegatedToID(userID int, meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("vote_delegated_$%d_to_id", meetingID)}] = v
	return v
}

func (r *Request) User_VoteDelegatedVoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "vote_delegated_vote_$_ids"}] = v
	return v
}

func (r *Request) User_VoteDelegatedVoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("vote_delegated_vote_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_VoteDelegationsFromIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "vote_delegations_$_from_ids"}] = v
	return v
}

func (r *Request) User_VoteDelegationsFromIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("vote_delegations_$%d_from_ids", meetingID)}] = v
	return v
}

func (r *Request) User_VoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "vote_$_ids"}] = v
	return v
}

func (r *Request) User_VoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("vote_$%d_ids", meetingID)}] = v
	return v
}

func (r *Request) User_VoteWeightTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[Key{"user", userID, "vote_weight_$"}] = v
	return v
}

func (r *Request) User_VoteWeight(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"user", userID, fmt.Sprintf("vote_weight_$%d", meetingID)}] = v
	return v
}

func (r *Request) Vote_DelegatedUserID(voteID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"vote", voteID, "delegated_user_id"}] = v
	return v
}

func (r *Request) Vote_ID(voteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"vote", voteID, "id"}] = v
	return v
}

func (r *Request) Vote_MeetingID(voteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"vote", voteID, "meeting_id"}] = v
	return v
}

func (r *Request) Vote_OptionID(voteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[Key{"vote", voteID, "option_id"}] = v
	return v
}

func (r *Request) Vote_UserID(voteID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[Key{"vote", voteID, "user_id"}] = v
	return v
}

func (r *Request) Vote_UserToken(voteID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"vote", voteID, "user_token"}] = v
	return v
}

func (r *Request) Vote_Value(voteID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"vote", voteID, "value"}] = v
	return v
}

func (r *Request) Vote_Weight(voteID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[Key{"vote", voteID, "weight"}] = v
	return v
}
