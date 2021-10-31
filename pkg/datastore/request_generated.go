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
	if v.request.err != nil {
		return false, v.request.err
	}

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
	if v.request.err != nil {
		return 0, v.request.err
	}

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
	if v.request.err != nil {
		return nil, v.request.err
	}

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
	if v.request.err != nil {
		return 0, v.request.err
	}

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
	if v.request.err != nil {
		return nil, v.request.err
	}

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
	if v.request.err != nil {
		return nil, v.request.err
	}

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
	if v.request.err != nil {
		return 0, false, v.request.err
	}

	if v.executed {
		return v.value, !v.isNull, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return 0, false, fmt.Errorf("executing request: %w", err)
	}

	return v.value, !v.isNull, nil
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
	if v.request.err != nil {
		return "", false, v.request.err
	}

	if v.executed {
		return v.value, !v.isNull, nil
	}

	if err := v.request.Execute(ctx); err != nil {
		return "", false, fmt.Errorf("executing request: %w", err)
	}

	return v.value, !v.isNull, nil
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
	if v.request.err != nil {
		return "", v.request.err
	}

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
	if v.request.err != nil {
		return nil, v.request.err
	}

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
	r.requested[fmt.Sprintf("agenda_item/%d/child_ids", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_Closed(agendaItemID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/closed", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_Comment(agendaItemID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/comment", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_ContentObjectID(agendaItemID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/content_object_id", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_Duration(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/duration", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_ID(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/id", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_IsHidden(agendaItemID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/is_hidden", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_IsInternal(agendaItemID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/is_internal", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_ItemNumber(agendaItemID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/item_number", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_Level(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/level", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_MeetingID(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/meeting_id", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_ParentID(agendaItemID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/parent_id", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_ProjectionIDs(agendaItemID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/projection_ids", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_TagIDs(agendaItemID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/tag_ids", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_Type(agendaItemID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/type", agendaItemID)] = v
	return v
}

func (r *Request) AgendaItem_Weight(agendaItemID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("agenda_item/%d/weight", agendaItemID)] = v
	return v
}

func (r *Request) AssignmentCandidate_AssignmentID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment_candidate/%d/assignment_id", assignmentCandidateID)] = v
	return v
}

func (r *Request) AssignmentCandidate_ID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment_candidate/%d/id", assignmentCandidateID)] = v
	return v
}

func (r *Request) AssignmentCandidate_MeetingID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment_candidate/%d/meeting_id", assignmentCandidateID)] = v
	return v
}

func (r *Request) AssignmentCandidate_UserID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment_candidate/%d/user_id", assignmentCandidateID)] = v
	return v
}

func (r *Request) AssignmentCandidate_Weight(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment_candidate/%d/weight", assignmentCandidateID)] = v
	return v
}

func (r *Request) Assignment_AgendaItemID(assignmentID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("assignment/%d/agenda_item_id", assignmentID)] = v
	return v
}

func (r *Request) Assignment_AttachmentIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("assignment/%d/attachment_ids", assignmentID)] = v
	return v
}

func (r *Request) Assignment_CandidateIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("assignment/%d/candidate_ids", assignmentID)] = v
	return v
}

func (r *Request) Assignment_DefaultPollDescription(assignmentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("assignment/%d/default_poll_description", assignmentID)] = v
	return v
}

func (r *Request) Assignment_Description(assignmentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("assignment/%d/description", assignmentID)] = v
	return v
}

func (r *Request) Assignment_ID(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment/%d/id", assignmentID)] = v
	return v
}

func (r *Request) Assignment_ListOfSpeakersID(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment/%d/list_of_speakers_id", assignmentID)] = v
	return v
}

func (r *Request) Assignment_MeetingID(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment/%d/meeting_id", assignmentID)] = v
	return v
}

func (r *Request) Assignment_NumberPollCandidates(assignmentID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("assignment/%d/number_poll_candidates", assignmentID)] = v
	return v
}

func (r *Request) Assignment_OpenPosts(assignmentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("assignment/%d/open_posts", assignmentID)] = v
	return v
}

func (r *Request) Assignment_Phase(assignmentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("assignment/%d/phase", assignmentID)] = v
	return v
}

func (r *Request) Assignment_PollIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("assignment/%d/poll_ids", assignmentID)] = v
	return v
}

func (r *Request) Assignment_ProjectionIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("assignment/%d/projection_ids", assignmentID)] = v
	return v
}

func (r *Request) Assignment_TagIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("assignment/%d/tag_ids", assignmentID)] = v
	return v
}

func (r *Request) Assignment_Title(assignmentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("assignment/%d/title", assignmentID)] = v
	return v
}

func (r *Request) ChatGroup_ID(chatGroupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("chat_group/%d/id", chatGroupID)] = v
	return v
}

func (r *Request) ChatGroup_MeetingID(chatGroupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("chat_group/%d/meeting_id", chatGroupID)] = v
	return v
}

func (r *Request) ChatGroup_Name(chatGroupID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("chat_group/%d/name", chatGroupID)] = v
	return v
}

func (r *Request) ChatGroup_ReadGroupIDs(chatGroupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("chat_group/%d/read_group_ids", chatGroupID)] = v
	return v
}

func (r *Request) ChatGroup_Weight(chatGroupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("chat_group/%d/weight", chatGroupID)] = v
	return v
}

func (r *Request) ChatGroup_WriteGroupIDs(chatGroupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("chat_group/%d/write_group_ids", chatGroupID)] = v
	return v
}

func (r *Request) Committee_DefaultMeetingID(committeeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("committee/%d/default_meeting_id", committeeID)] = v
	return v
}

func (r *Request) Committee_Description(committeeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("committee/%d/description", committeeID)] = v
	return v
}

func (r *Request) Committee_ForwardToCommitteeIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("committee/%d/forward_to_committee_ids", committeeID)] = v
	return v
}

func (r *Request) Committee_ID(committeeID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("committee/%d/id", committeeID)] = v
	return v
}

func (r *Request) Committee_MeetingIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("committee/%d/meeting_ids", committeeID)] = v
	return v
}

func (r *Request) Committee_Name(committeeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("committee/%d/name", committeeID)] = v
	return v
}

func (r *Request) Committee_OrganizationID(committeeID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("committee/%d/organization_id", committeeID)] = v
	return v
}

func (r *Request) Committee_OrganizationTagIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("committee/%d/organization_tag_ids", committeeID)] = v
	return v
}

func (r *Request) Committee_ReceiveForwardingsFromCommitteeIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("committee/%d/receive_forwardings_from_committee_ids", committeeID)] = v
	return v
}

func (r *Request) Committee_TemplateMeetingID(committeeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("committee/%d/template_meeting_id", committeeID)] = v
	return v
}

func (r *Request) Committee_UserIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("committee/%d/user_ids", committeeID)] = v
	return v
}

func (r *Request) Group_AdminGroupForMeetingID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("group/%d/admin_group_for_meeting_id", groupID)] = v
	return v
}

func (r *Request) Group_DefaultGroupForMeetingID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("group/%d/default_group_for_meeting_id", groupID)] = v
	return v
}

func (r *Request) Group_ID(groupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("group/%d/id", groupID)] = v
	return v
}

func (r *Request) Group_MediafileAccessGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/mediafile_access_group_ids", groupID)] = v
	return v
}

func (r *Request) Group_MediafileInheritedAccessGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/mediafile_inherited_access_group_ids", groupID)] = v
	return v
}

func (r *Request) Group_MeetingID(groupID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("group/%d/meeting_id", groupID)] = v
	return v
}

func (r *Request) Group_Name(groupID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("group/%d/name", groupID)] = v
	return v
}

func (r *Request) Group_Permissions(groupID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/permissions", groupID)] = v
	return v
}

func (r *Request) Group_PollIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/poll_ids", groupID)] = v
	return v
}

func (r *Request) Group_ReadChatGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/read_chat_group_ids", groupID)] = v
	return v
}

func (r *Request) Group_ReadCommentSectionIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/read_comment_section_ids", groupID)] = v
	return v
}

func (r *Request) Group_UsedAsAssignmentPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("group/%d/used_as_assignment_poll_default_id", groupID)] = v
	return v
}

func (r *Request) Group_UsedAsMotionPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("group/%d/used_as_motion_poll_default_id", groupID)] = v
	return v
}

func (r *Request) Group_UsedAsPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("group/%d/used_as_poll_default_id", groupID)] = v
	return v
}

func (r *Request) Group_UserIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/user_ids", groupID)] = v
	return v
}

func (r *Request) Group_WriteChatGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/write_chat_group_ids", groupID)] = v
	return v
}

func (r *Request) Group_WriteCommentSectionIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("group/%d/write_comment_section_ids", groupID)] = v
	return v
}

func (r *Request) ListOfSpeakers_Closed(listOfSpeakersID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("list_of_speakers/%d/closed", listOfSpeakersID)] = v
	return v
}

func (r *Request) ListOfSpeakers_ContentObjectID(listOfSpeakersID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("list_of_speakers/%d/content_object_id", listOfSpeakersID)] = v
	return v
}

func (r *Request) ListOfSpeakers_ID(listOfSpeakersID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("list_of_speakers/%d/id", listOfSpeakersID)] = v
	return v
}

func (r *Request) ListOfSpeakers_MeetingID(listOfSpeakersID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("list_of_speakers/%d/meeting_id", listOfSpeakersID)] = v
	return v
}

func (r *Request) ListOfSpeakers_ProjectionIDs(listOfSpeakersID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("list_of_speakers/%d/projection_ids", listOfSpeakersID)] = v
	return v
}

func (r *Request) ListOfSpeakers_SpeakerIDs(listOfSpeakersID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("list_of_speakers/%d/speaker_ids", listOfSpeakersID)] = v
	return v
}

func (r *Request) Mediafile_AccessGroupIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/access_group_ids", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_AttachmentIDs(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/attachment_ids", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_ChildIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/child_ids", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_CreateTimestamp(mediafileID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/create_timestamp", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_Filename(mediafileID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/filename", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_Filesize(mediafileID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/filesize", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_ID(mediafileID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/id", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_InheritedAccessGroupIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/inherited_access_group_ids", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_IsDirectory(mediafileID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/is_directory", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_IsPublic(mediafileID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/is_public", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_ListOfSpeakersID(mediafileID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/list_of_speakers_id", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_MeetingID(mediafileID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/meeting_id", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_Mimetype(mediafileID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/mimetype", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_ParentID(mediafileID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/parent_id", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_PdfInformation(mediafileID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/pdf_information", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_ProjectionIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/projection_ids", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_Title(mediafileID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/title", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_UsedAsFontInMeetingIDTmpl(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/used_as_font_$_in_meeting_id", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_UsedAsFontInMeetingID(mediafileID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/used_as_font_$%s_in_meeting_id", mediafileID, replacement)] = v
	return v
}

func (r *Request) Mediafile_UsedAsLogoInMeetingIDTmpl(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/used_as_logo_$_in_meeting_id", mediafileID)] = v
	return v
}

func (r *Request) Mediafile_UsedAsLogoInMeetingID(mediafileID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("mediafile/%d/used_as_logo_$%s_in_meeting_id", mediafileID, replacement)] = v
	return v
}

func (r *Request) Meeting_AdminGroupID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/admin_group_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_AgendaEnableNumbering(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/agenda_enable_numbering", meetingID)] = v
	return v
}

func (r *Request) Meeting_AgendaItemCreation(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/agenda_item_creation", meetingID)] = v
	return v
}

func (r *Request) Meeting_AgendaItemIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/agenda_item_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_AgendaNewItemsDefaultVisibility(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/agenda_new_items_default_visibility", meetingID)] = v
	return v
}

func (r *Request) Meeting_AgendaNumberPrefix(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/agenda_number_prefix", meetingID)] = v
	return v
}

func (r *Request) Meeting_AgendaNumeralSystem(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/agenda_numeral_system", meetingID)] = v
	return v
}

func (r *Request) Meeting_AgendaShowInternalItemsOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/agenda_show_internal_items_on_projector", meetingID)] = v
	return v
}

func (r *Request) Meeting_AgendaShowSubtitles(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/agenda_show_subtitles", meetingID)] = v
	return v
}

func (r *Request) Meeting_AllProjectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/all_projection_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_ApplauseEnable(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/applause_enable", meetingID)] = v
	return v
}

func (r *Request) Meeting_ApplauseMaxAmount(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/applause_max_amount", meetingID)] = v
	return v
}

func (r *Request) Meeting_ApplauseMinAmount(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/applause_min_amount", meetingID)] = v
	return v
}

func (r *Request) Meeting_ApplauseParticleImageUrl(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/applause_particle_image_url", meetingID)] = v
	return v
}

func (r *Request) Meeting_ApplauseShowLevel(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/applause_show_level", meetingID)] = v
	return v
}

func (r *Request) Meeting_ApplauseTimeout(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/applause_timeout", meetingID)] = v
	return v
}

func (r *Request) Meeting_ApplauseType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/applause_type", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentCandidateIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_candidate_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentPollAddCandidatesToListOfSpeakers(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_poll_add_candidates_to_list_of_speakers", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentPollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_poll_ballot_paper_number", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentPollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_poll_ballot_paper_selection", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_poll_default_100_percent_base", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_poll_default_group_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefaultMethod(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_poll_default_method", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentPollDefaultType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_poll_default_type", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentPollSortPollResultByVotes(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignment_poll_sort_poll_result_by_votes", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentsExportPreamble(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignments_export_preamble", meetingID)] = v
	return v
}

func (r *Request) Meeting_AssignmentsExportTitle(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/assignments_export_title", meetingID)] = v
	return v
}

func (r *Request) Meeting_ChatGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/chat_group_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_CommitteeID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/committee_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceAutoConnect(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_auto_connect", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceAutoConnectNextSpeakers(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_auto_connect_next_speakers", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceEnableHelpdesk(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_enable_helpdesk", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceLosRestriction(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_los_restriction", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceOpenMicrophone(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_open_microphone", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceOpenVideo(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_open_video", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceShow(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_show", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceStreamPosterUrl(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_stream_poster_url", meetingID)] = v
	return v
}

func (r *Request) Meeting_ConferenceStreamUrl(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/conference_stream_url", meetingID)] = v
	return v
}

func (r *Request) Meeting_CustomTranslations(meetingID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[fmt.Sprintf("meeting/%d/custom_translations", meetingID)] = v
	return v
}

func (r *Request) Meeting_DefaultGroupID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/default_group_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_DefaultMeetingForCommitteeID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/default_meeting_for_committee_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_DefaultProjectorIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/default_projector_$_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_DefaultProjectorID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/default_projector_$%s_id", meetingID, replacement)] = v
	return v
}

func (r *Request) Meeting_Description(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/description", meetingID)] = v
	return v
}

func (r *Request) Meeting_EnableAnonymous(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/enable_anonymous", meetingID)] = v
	return v
}

func (r *Request) Meeting_EnableChat(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/enable_chat", meetingID)] = v
	return v
}

func (r *Request) Meeting_EndTime(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/end_time", meetingID)] = v
	return v
}

func (r *Request) Meeting_ExportCsvEncoding(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/export_csv_encoding", meetingID)] = v
	return v
}

func (r *Request) Meeting_ExportCsvSeparator(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/export_csv_separator", meetingID)] = v
	return v
}

func (r *Request) Meeting_ExportPdfFontsize(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/export_pdf_fontsize", meetingID)] = v
	return v
}

func (r *Request) Meeting_ExportPdfPagenumberAlignment(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/export_pdf_pagenumber_alignment", meetingID)] = v
	return v
}

func (r *Request) Meeting_ExportPdfPagesize(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/export_pdf_pagesize", meetingID)] = v
	return v
}

func (r *Request) Meeting_FontIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/font_$_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_FontID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/font_$%s_id", meetingID, replacement)] = v
	return v
}

func (r *Request) Meeting_GroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/group_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_ID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/id", meetingID)] = v
	return v
}

func (r *Request) Meeting_ImportedAt(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/imported_at", meetingID)] = v
	return v
}

func (r *Request) Meeting_IsActiveInOrganizationID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/is_active_in_organization_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_JitsiDomain(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/jitsi_domain", meetingID)] = v
	return v
}

func (r *Request) Meeting_JitsiRoomName(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/jitsi_room_name", meetingID)] = v
	return v
}

func (r *Request) Meeting_JitsiRoomPassword(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/jitsi_room_password", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersAmountLastOnProjector(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_amount_last_on_projector", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersAmountNextOnProjector(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_amount_next_on_projector", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersCanSetContributionSelf(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_can_set_contribution_self", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersCountdownID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_countdown_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersCoupleCountdown(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_couple_countdown", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_enable_point_of_order_speakers", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersEnableProContraSpeech(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_enable_pro_contra_speech", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersInitiallyClosed(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_initially_closed", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersPresentUsersOnly(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_present_users_only", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_show_amount_of_speakers_on_slide", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersShowFirstContribution(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_show_first_contribution", meetingID)] = v
	return v
}

func (r *Request) Meeting_ListOfSpeakersSpeakerNoteForEveryone(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/list_of_speakers_speaker_note_for_everyone", meetingID)] = v
	return v
}

func (r *Request) Meeting_Location(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/location", meetingID)] = v
	return v
}

func (r *Request) Meeting_LogoIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/logo_$_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_LogoID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/logo_$%s_id", meetingID, replacement)] = v
	return v
}

func (r *Request) Meeting_MediafileIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/mediafile_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionBlockIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_block_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionCategoryIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_category_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionChangeRecommendationIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_change_recommendation_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionCommentIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_comment_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionCommentSectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_comment_section_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionPollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_poll_ballot_paper_number", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionPollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_poll_ballot_paper_selection", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionPollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_poll_default_100_percent_base", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_poll_default_group_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionPollDefaultType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_poll_default_type", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionStateIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_state_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionStatuteParagraphIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_statute_paragraph_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionSubmitterIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_submitter_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionWorkflowIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motion_workflow_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsEnabled(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_amendments_enabled", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsInMainList(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_amendments_in_main_list", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsMultipleParagraphs(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_amendments_multiple_paragraphs", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsOfAmendments(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_amendments_of_amendments", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsPrefix(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_amendments_prefix", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsAmendmentsTextMode(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_amendments_text_mode", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultAmendmentWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_default_amendment_workflow_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultLineNumbering(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_default_line_numbering", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultSorting(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_default_sorting", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultStatuteAmendmentWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_default_statute_amendment_workflow_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsDefaultWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_default_workflow_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsEnableReasonOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_enable_reason_on_projector", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsEnableRecommendationOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_enable_recommendation_on_projector", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsEnableSideboxOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_enable_sidebox_on_projector", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsEnableTextOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_enable_text_on_projector", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsExportFollowRecommendation(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_export_follow_recommendation", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsExportPreamble(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_export_preamble", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsExportSubmitterRecommendation(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_export_submitter_recommendation", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsExportTitle(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_export_title", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsLineLength(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_line_length", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsNumberMinDigits(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_number_min_digits", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsNumberType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_number_type", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsNumberWithBlank(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_number_with_blank", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsPreamble(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_preamble", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsReasonRequired(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_reason_required", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsRecommendationTextMode(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_recommendation_text_mode", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsRecommendationsBy(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_recommendations_by", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsShowReferringMotions(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_show_referring_motions", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsShowSequentialNumber(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_show_sequential_number", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsStatuteRecommendationsBy(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_statute_recommendations_by", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsStatutesEnabled(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_statutes_enabled", meetingID)] = v
	return v
}

func (r *Request) Meeting_MotionsSupportersMinAmount(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/motions_supporters_min_amount", meetingID)] = v
	return v
}

func (r *Request) Meeting_Name(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/name", meetingID)] = v
	return v
}

func (r *Request) Meeting_OptionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/option_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_OrganizationTagIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/organization_tag_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_PersonalNoteIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/personal_note_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_ballot_paper_number", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_ballot_paper_selection", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollCountdownID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_countdown_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollCoupleCountdown(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_couple_countdown", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_default_100_percent_base", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_default_group_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollDefaultMethod(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_default_method", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollDefaultType(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_default_type", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_PollSortPollResultByVotes(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/poll_sort_poll_result_by_votes", meetingID)] = v
	return v
}

func (r *Request) Meeting_PresentUserIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/present_user_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_ProjectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/projection_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_ProjectorCountdownDefaultTime(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/projector_countdown_default_time", meetingID)] = v
	return v
}

func (r *Request) Meeting_ProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/projector_countdown_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_ProjectorCountdownWarningTime(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/projector_countdown_warning_time", meetingID)] = v
	return v
}

func (r *Request) Meeting_ProjectorIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/projector_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_ProjectorMessageIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/projector_message_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_ReferenceProjectorID(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/reference_projector_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_SpeakerIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/speaker_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_StartTime(meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/start_time", meetingID)] = v
	return v
}

func (r *Request) Meeting_TagIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/tag_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_TemplateForCommitteeID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("meeting/%d/template_for_committee_id", meetingID)] = v
	return v
}

func (r *Request) Meeting_TopicIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/topic_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_UrlName(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/url_name", meetingID)] = v
	return v
}

func (r *Request) Meeting_UserIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/user_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersAllowSelfSetPresent(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_allow_self_set_present", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersEmailBody(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_email_body", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersEmailReplyto(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_email_replyto", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersEmailSender(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_email_sender", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersEmailSubject(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_email_subject", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersEnablePresenceView(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_enable_presence_view", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersEnableVoteWeight(meetingID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_enable_vote_weight", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersPdfUrl(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_pdf_url", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersPdfWelcometext(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_pdf_welcometext", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersPdfWelcometitle(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_pdf_welcometitle", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersPdfWlanEncryption(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_pdf_wlan_encryption", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersPdfWlanPassword(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_pdf_wlan_password", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersPdfWlanSsid(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_pdf_wlan_ssid", meetingID)] = v
	return v
}

func (r *Request) Meeting_UsersSortBy(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/users_sort_by", meetingID)] = v
	return v
}

func (r *Request) Meeting_VoteIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("meeting/%d/vote_ids", meetingID)] = v
	return v
}

func (r *Request) Meeting_WelcomeText(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/welcome_text", meetingID)] = v
	return v
}

func (r *Request) Meeting_WelcomeTitle(meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("meeting/%d/welcome_title", meetingID)] = v
	return v
}

func (r *Request) MotionBlock_AgendaItemID(motionBlockID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion_block/%d/agenda_item_id", motionBlockID)] = v
	return v
}

func (r *Request) MotionBlock_ID(motionBlockID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_block/%d/id", motionBlockID)] = v
	return v
}

func (r *Request) MotionBlock_Internal(motionBlockID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_block/%d/internal", motionBlockID)] = v
	return v
}

func (r *Request) MotionBlock_ListOfSpeakersID(motionBlockID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_block/%d/list_of_speakers_id", motionBlockID)] = v
	return v
}

func (r *Request) MotionBlock_MeetingID(motionBlockID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_block/%d/meeting_id", motionBlockID)] = v
	return v
}

func (r *Request) MotionBlock_MotionIDs(motionBlockID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_block/%d/motion_ids", motionBlockID)] = v
	return v
}

func (r *Request) MotionBlock_ProjectionIDs(motionBlockID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_block/%d/projection_ids", motionBlockID)] = v
	return v
}

func (r *Request) MotionBlock_Title(motionBlockID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_block/%d/title", motionBlockID)] = v
	return v
}

func (r *Request) MotionCategory_ChildIDs(motionCategoryID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/child_ids", motionCategoryID)] = v
	return v
}

func (r *Request) MotionCategory_ID(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/id", motionCategoryID)] = v
	return v
}

func (r *Request) MotionCategory_Level(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/level", motionCategoryID)] = v
	return v
}

func (r *Request) MotionCategory_MeetingID(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/meeting_id", motionCategoryID)] = v
	return v
}

func (r *Request) MotionCategory_MotionIDs(motionCategoryID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/motion_ids", motionCategoryID)] = v
	return v
}

func (r *Request) MotionCategory_Name(motionCategoryID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/name", motionCategoryID)] = v
	return v
}

func (r *Request) MotionCategory_ParentID(motionCategoryID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/parent_id", motionCategoryID)] = v
	return v
}

func (r *Request) MotionCategory_Prefix(motionCategoryID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/prefix", motionCategoryID)] = v
	return v
}

func (r *Request) MotionCategory_Weight(motionCategoryID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_category/%d/weight", motionCategoryID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_CreationTime(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/creation_time", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_ID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/id", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_Internal(motionChangeRecommendationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/internal", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_LineFrom(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/line_from", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_LineTo(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/line_to", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_MeetingID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/meeting_id", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_MotionID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/motion_id", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_OtherDescription(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/other_description", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_Rejected(motionChangeRecommendationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/rejected", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_Text(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/text", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionChangeRecommendation_Type(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_change_recommendation/%d/type", motionChangeRecommendationID)] = v
	return v
}

func (r *Request) MotionCommentSection_CommentIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_comment_section/%d/comment_ids", motionCommentSectionID)] = v
	return v
}

func (r *Request) MotionCommentSection_ID(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_comment_section/%d/id", motionCommentSectionID)] = v
	return v
}

func (r *Request) MotionCommentSection_MeetingID(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_comment_section/%d/meeting_id", motionCommentSectionID)] = v
	return v
}

func (r *Request) MotionCommentSection_Name(motionCommentSectionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_comment_section/%d/name", motionCommentSectionID)] = v
	return v
}

func (r *Request) MotionCommentSection_ReadGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_comment_section/%d/read_group_ids", motionCommentSectionID)] = v
	return v
}

func (r *Request) MotionCommentSection_Weight(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_comment_section/%d/weight", motionCommentSectionID)] = v
	return v
}

func (r *Request) MotionCommentSection_WriteGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_comment_section/%d/write_group_ids", motionCommentSectionID)] = v
	return v
}

func (r *Request) MotionComment_Comment(motionCommentID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_comment/%d/comment", motionCommentID)] = v
	return v
}

func (r *Request) MotionComment_ID(motionCommentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_comment/%d/id", motionCommentID)] = v
	return v
}

func (r *Request) MotionComment_MeetingID(motionCommentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_comment/%d/meeting_id", motionCommentID)] = v
	return v
}

func (r *Request) MotionComment_MotionID(motionCommentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_comment/%d/motion_id", motionCommentID)] = v
	return v
}

func (r *Request) MotionComment_SectionID(motionCommentID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_comment/%d/section_id", motionCommentID)] = v
	return v
}

func (r *Request) MotionState_AllowCreatePoll(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/allow_create_poll", motionStateID)] = v
	return v
}

func (r *Request) MotionState_AllowSubmitterEdit(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/allow_submitter_edit", motionStateID)] = v
	return v
}

func (r *Request) MotionState_AllowSupport(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/allow_support", motionStateID)] = v
	return v
}

func (r *Request) MotionState_CssClass(motionStateID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/css_class", motionStateID)] = v
	return v
}

func (r *Request) MotionState_DontSetIDentifier(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/dont_set_identifier", motionStateID)] = v
	return v
}

func (r *Request) MotionState_FirstStateOfWorkflowID(motionStateID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/first_state_of_workflow_id", motionStateID)] = v
	return v
}

func (r *Request) MotionState_ID(motionStateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/id", motionStateID)] = v
	return v
}

func (r *Request) MotionState_MeetingID(motionStateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/meeting_id", motionStateID)] = v
	return v
}

func (r *Request) MotionState_MergeAmendmentIntoFinal(motionStateID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/merge_amendment_into_final", motionStateID)] = v
	return v
}

func (r *Request) MotionState_MotionIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/motion_ids", motionStateID)] = v
	return v
}

func (r *Request) MotionState_MotionRecommendationIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/motion_recommendation_ids", motionStateID)] = v
	return v
}

func (r *Request) MotionState_Name(motionStateID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/name", motionStateID)] = v
	return v
}

func (r *Request) MotionState_NextStateIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/next_state_ids", motionStateID)] = v
	return v
}

func (r *Request) MotionState_PreviousStateIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/previous_state_ids", motionStateID)] = v
	return v
}

func (r *Request) MotionState_RecommendationLabel(motionStateID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/recommendation_label", motionStateID)] = v
	return v
}

func (r *Request) MotionState_Restrictions(motionStateID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/restrictions", motionStateID)] = v
	return v
}

func (r *Request) MotionState_SetNumber(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/set_number", motionStateID)] = v
	return v
}

func (r *Request) MotionState_ShowRecommendationExtensionField(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/show_recommendation_extension_field", motionStateID)] = v
	return v
}

func (r *Request) MotionState_ShowStateExtensionField(motionStateID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/show_state_extension_field", motionStateID)] = v
	return v
}

func (r *Request) MotionState_Weight(motionStateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/weight", motionStateID)] = v
	return v
}

func (r *Request) MotionState_WorkflowID(motionStateID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_state/%d/workflow_id", motionStateID)] = v
	return v
}

func (r *Request) MotionStatuteParagraph_ID(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_statute_paragraph/%d/id", motionStatuteParagraphID)] = v
	return v
}

func (r *Request) MotionStatuteParagraph_MeetingID(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_statute_paragraph/%d/meeting_id", motionStatuteParagraphID)] = v
	return v
}

func (r *Request) MotionStatuteParagraph_MotionIDs(motionStatuteParagraphID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_statute_paragraph/%d/motion_ids", motionStatuteParagraphID)] = v
	return v
}

func (r *Request) MotionStatuteParagraph_Text(motionStatuteParagraphID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_statute_paragraph/%d/text", motionStatuteParagraphID)] = v
	return v
}

func (r *Request) MotionStatuteParagraph_Title(motionStatuteParagraphID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_statute_paragraph/%d/title", motionStatuteParagraphID)] = v
	return v
}

func (r *Request) MotionStatuteParagraph_Weight(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_statute_paragraph/%d/weight", motionStatuteParagraphID)] = v
	return v
}

func (r *Request) MotionSubmitter_ID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_submitter/%d/id", motionSubmitterID)] = v
	return v
}

func (r *Request) MotionSubmitter_MeetingID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_submitter/%d/meeting_id", motionSubmitterID)] = v
	return v
}

func (r *Request) MotionSubmitter_MotionID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_submitter/%d/motion_id", motionSubmitterID)] = v
	return v
}

func (r *Request) MotionSubmitter_UserID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_submitter/%d/user_id", motionSubmitterID)] = v
	return v
}

func (r *Request) MotionSubmitter_Weight(motionSubmitterID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_submitter/%d/weight", motionSubmitterID)] = v
	return v
}

func (r *Request) MotionWorkflow_DefaultAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion_workflow/%d/default_amendment_workflow_meeting_id", motionWorkflowID)] = v
	return v
}

func (r *Request) MotionWorkflow_DefaultStatuteAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion_workflow/%d/default_statute_amendment_workflow_meeting_id", motionWorkflowID)] = v
	return v
}

func (r *Request) MotionWorkflow_DefaultWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion_workflow/%d/default_workflow_meeting_id", motionWorkflowID)] = v
	return v
}

func (r *Request) MotionWorkflow_FirstStateID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_workflow/%d/first_state_id", motionWorkflowID)] = v
	return v
}

func (r *Request) MotionWorkflow_ID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_workflow/%d/id", motionWorkflowID)] = v
	return v
}

func (r *Request) MotionWorkflow_MeetingID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion_workflow/%d/meeting_id", motionWorkflowID)] = v
	return v
}

func (r *Request) MotionWorkflow_Name(motionWorkflowID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion_workflow/%d/name", motionWorkflowID)] = v
	return v
}

func (r *Request) MotionWorkflow_StateIDs(motionWorkflowID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion_workflow/%d/state_ids", motionWorkflowID)] = v
	return v
}

func (r *Request) Motion_AgendaItemID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/agenda_item_id", motionID)] = v
	return v
}

func (r *Request) Motion_AllDerivedMotionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/all_derived_motion_ids", motionID)] = v
	return v
}

func (r *Request) Motion_AllOriginIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/all_origin_ids", motionID)] = v
	return v
}

func (r *Request) Motion_AmendmentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/amendment_ids", motionID)] = v
	return v
}

func (r *Request) Motion_AmendmentParagraphTmpl(motionID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/amendment_paragraph_$", motionID)] = v
	return v
}

func (r *Request) Motion_AmendmentParagraph(motionID int, replacement string) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion/%d/amendment_paragraph_$%s", motionID, replacement)] = v
	return v
}

func (r *Request) Motion_AttachmentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/attachment_ids", motionID)] = v
	return v
}

func (r *Request) Motion_BlockID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/block_id", motionID)] = v
	return v
}

func (r *Request) Motion_CategoryID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/category_id", motionID)] = v
	return v
}

func (r *Request) Motion_CategoryWeight(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/category_weight", motionID)] = v
	return v
}

func (r *Request) Motion_ChangeRecommendationIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/change_recommendation_ids", motionID)] = v
	return v
}

func (r *Request) Motion_CommentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/comment_ids", motionID)] = v
	return v
}

func (r *Request) Motion_Created(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/created", motionID)] = v
	return v
}

func (r *Request) Motion_DerivedMotionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/derived_motion_ids", motionID)] = v
	return v
}

func (r *Request) Motion_ID(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/id", motionID)] = v
	return v
}

func (r *Request) Motion_LastModified(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/last_modified", motionID)] = v
	return v
}

func (r *Request) Motion_LeadMotionID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/lead_motion_id", motionID)] = v
	return v
}

func (r *Request) Motion_ListOfSpeakersID(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/list_of_speakers_id", motionID)] = v
	return v
}

func (r *Request) Motion_MeetingID(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/meeting_id", motionID)] = v
	return v
}

func (r *Request) Motion_ModifiedFinalVersion(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion/%d/modified_final_version", motionID)] = v
	return v
}

func (r *Request) Motion_Number(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion/%d/number", motionID)] = v
	return v
}

func (r *Request) Motion_NumberValue(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/number_value", motionID)] = v
	return v
}

func (r *Request) Motion_OptionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/option_ids", motionID)] = v
	return v
}

func (r *Request) Motion_OriginID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/origin_id", motionID)] = v
	return v
}

func (r *Request) Motion_PersonalNoteIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/personal_note_ids", motionID)] = v
	return v
}

func (r *Request) Motion_PollIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/poll_ids", motionID)] = v
	return v
}

func (r *Request) Motion_ProjectionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/projection_ids", motionID)] = v
	return v
}

func (r *Request) Motion_Reason(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion/%d/reason", motionID)] = v
	return v
}

func (r *Request) Motion_RecommendationExtension(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion/%d/recommendation_extension", motionID)] = v
	return v
}

func (r *Request) Motion_RecommendationExtensionReferenceIDs(motionID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/recommendation_extension_reference_ids", motionID)] = v
	return v
}

func (r *Request) Motion_RecommendationID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/recommendation_id", motionID)] = v
	return v
}

func (r *Request) Motion_ReferencedInMotionRecommendationExtensionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/referenced_in_motion_recommendation_extension_ids", motionID)] = v
	return v
}

func (r *Request) Motion_SequentialNumber(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/sequential_number", motionID)] = v
	return v
}

func (r *Request) Motion_SortChildIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/sort_child_ids", motionID)] = v
	return v
}

func (r *Request) Motion_SortParentID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/sort_parent_id", motionID)] = v
	return v
}

func (r *Request) Motion_SortWeight(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/sort_weight", motionID)] = v
	return v
}

func (r *Request) Motion_StateExtension(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion/%d/state_extension", motionID)] = v
	return v
}

func (r *Request) Motion_StateID(motionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/state_id", motionID)] = v
	return v
}

func (r *Request) Motion_StatuteParagraphID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("motion/%d/statute_paragraph_id", motionID)] = v
	return v
}

func (r *Request) Motion_SubmitterIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/submitter_ids", motionID)] = v
	return v
}

func (r *Request) Motion_SupporterIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/supporter_ids", motionID)] = v
	return v
}

func (r *Request) Motion_TagIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("motion/%d/tag_ids", motionID)] = v
	return v
}

func (r *Request) Motion_Text(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion/%d/text", motionID)] = v
	return v
}

func (r *Request) Motion_Title(motionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("motion/%d/title", motionID)] = v
	return v
}

func (r *Request) Option_Abstain(optionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("option/%d/abstain", optionID)] = v
	return v
}

func (r *Request) Option_ContentObjectID(optionID int) *ValueMaybeString {
	v := &ValueMaybeString{request: r}
	r.requested[fmt.Sprintf("option/%d/content_object_id", optionID)] = v
	return v
}

func (r *Request) Option_ID(optionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("option/%d/id", optionID)] = v
	return v
}

func (r *Request) Option_MeetingID(optionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("option/%d/meeting_id", optionID)] = v
	return v
}

func (r *Request) Option_No(optionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("option/%d/no", optionID)] = v
	return v
}

func (r *Request) Option_PollID(optionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("option/%d/poll_id", optionID)] = v
	return v
}

func (r *Request) Option_Text(optionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("option/%d/text", optionID)] = v
	return v
}

func (r *Request) Option_UsedAsGlobalOptionInPollID(optionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("option/%d/used_as_global_option_in_poll_id", optionID)] = v
	return v
}

func (r *Request) Option_VoteIDs(optionID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("option/%d/vote_ids", optionID)] = v
	return v
}

func (r *Request) Option_Weight(optionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("option/%d/weight", optionID)] = v
	return v
}

func (r *Request) Option_Yes(optionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("option/%d/yes", optionID)] = v
	return v
}

func (r *Request) OrganizationTag_Color(organizationTagID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("organization_tag/%d/color", organizationTagID)] = v
	return v
}

func (r *Request) OrganizationTag_ID(organizationTagID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("organization_tag/%d/id", organizationTagID)] = v
	return v
}

func (r *Request) OrganizationTag_Name(organizationTagID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("organization_tag/%d/name", organizationTagID)] = v
	return v
}

func (r *Request) OrganizationTag_OrganizationID(organizationTagID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("organization_tag/%d/organization_id", organizationTagID)] = v
	return v
}

func (r *Request) OrganizationTag_TaggedIDs(organizationTagID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("organization_tag/%d/tagged_ids", organizationTagID)] = v
	return v
}

func (r *Request) Organization_ActiveMeetingIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("organization/%d/active_meeting_ids", organizationID)] = v
	return v
}

func (r *Request) Organization_CommitteeIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("organization/%d/committee_ids", organizationID)] = v
	return v
}

func (r *Request) Organization_Description(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("organization/%d/description", organizationID)] = v
	return v
}

func (r *Request) Organization_EnableElectronicVoting(organizationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("organization/%d/enable_electronic_voting", organizationID)] = v
	return v
}

func (r *Request) Organization_ID(organizationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("organization/%d/id", organizationID)] = v
	return v
}

func (r *Request) Organization_LegalNotice(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("organization/%d/legal_notice", organizationID)] = v
	return v
}

func (r *Request) Organization_LimitOfMeetings(organizationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("organization/%d/limit_of_meetings", organizationID)] = v
	return v
}

func (r *Request) Organization_LimitOfUsers(organizationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("organization/%d/limit_of_users", organizationID)] = v
	return v
}

func (r *Request) Organization_LoginText(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("organization/%d/login_text", organizationID)] = v
	return v
}

func (r *Request) Organization_Name(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("organization/%d/name", organizationID)] = v
	return v
}

func (r *Request) Organization_OrganizationTagIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("organization/%d/organization_tag_ids", organizationID)] = v
	return v
}

func (r *Request) Organization_PrivacyPolicy(organizationID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("organization/%d/privacy_policy", organizationID)] = v
	return v
}

func (r *Request) Organization_ResetPasswordVerboseErrors(organizationID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("organization/%d/reset_password_verbose_errors", organizationID)] = v
	return v
}

func (r *Request) Organization_ResourceIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("organization/%d/resource_ids", organizationID)] = v
	return v
}

func (r *Request) Organization_ThemeID(organizationID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("organization/%d/theme_id", organizationID)] = v
	return v
}

func (r *Request) Organization_ThemeIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("organization/%d/theme_ids", organizationID)] = v
	return v
}

func (r *Request) PersonalNote_ContentObjectID(personalNoteID int) *ValueMaybeString {
	v := &ValueMaybeString{request: r}
	r.requested[fmt.Sprintf("personal_note/%d/content_object_id", personalNoteID)] = v
	return v
}

func (r *Request) PersonalNote_ID(personalNoteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("personal_note/%d/id", personalNoteID)] = v
	return v
}

func (r *Request) PersonalNote_MeetingID(personalNoteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("personal_note/%d/meeting_id", personalNoteID)] = v
	return v
}

func (r *Request) PersonalNote_Note(personalNoteID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("personal_note/%d/note", personalNoteID)] = v
	return v
}

func (r *Request) PersonalNote_Star(personalNoteID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("personal_note/%d/star", personalNoteID)] = v
	return v
}

func (r *Request) PersonalNote_UserID(personalNoteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("personal_note/%d/user_id", personalNoteID)] = v
	return v
}

func (r *Request) Poll_Backend(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/backend", pollID)] = v
	return v
}

func (r *Request) Poll_ContentObjectID(pollID int) *ValueMaybeString {
	v := &ValueMaybeString{request: r}
	r.requested[fmt.Sprintf("poll/%d/content_object_id", pollID)] = v
	return v
}

func (r *Request) Poll_Description(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/description", pollID)] = v
	return v
}

func (r *Request) Poll_EntitledGroupIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("poll/%d/entitled_group_ids", pollID)] = v
	return v
}

func (r *Request) Poll_EntitledUsersAtStop(pollID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[fmt.Sprintf("poll/%d/entitled_users_at_stop", pollID)] = v
	return v
}

func (r *Request) Poll_GlobalAbstain(pollID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("poll/%d/global_abstain", pollID)] = v
	return v
}

func (r *Request) Poll_GlobalNo(pollID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("poll/%d/global_no", pollID)] = v
	return v
}

func (r *Request) Poll_GlobalOptionID(pollID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("poll/%d/global_option_id", pollID)] = v
	return v
}

func (r *Request) Poll_GlobalYes(pollID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("poll/%d/global_yes", pollID)] = v
	return v
}

func (r *Request) Poll_ID(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("poll/%d/id", pollID)] = v
	return v
}

func (r *Request) Poll_IsPseudoanonymized(pollID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("poll/%d/is_pseudoanonymized", pollID)] = v
	return v
}

func (r *Request) Poll_MaxVotesAmount(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("poll/%d/max_votes_amount", pollID)] = v
	return v
}

func (r *Request) Poll_MeetingID(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("poll/%d/meeting_id", pollID)] = v
	return v
}

func (r *Request) Poll_MinVotesAmount(pollID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("poll/%d/min_votes_amount", pollID)] = v
	return v
}

func (r *Request) Poll_OnehundredPercentBase(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/onehundred_percent_base", pollID)] = v
	return v
}

func (r *Request) Poll_OptionIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("poll/%d/option_ids", pollID)] = v
	return v
}

func (r *Request) Poll_Pollmethod(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/pollmethod", pollID)] = v
	return v
}

func (r *Request) Poll_ProjectionIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("poll/%d/projection_ids", pollID)] = v
	return v
}

func (r *Request) Poll_State(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/state", pollID)] = v
	return v
}

func (r *Request) Poll_Title(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/title", pollID)] = v
	return v
}

func (r *Request) Poll_Type(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/type", pollID)] = v
	return v
}

func (r *Request) Poll_VotedIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("poll/%d/voted_ids", pollID)] = v
	return v
}

func (r *Request) Poll_Votescast(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/votescast", pollID)] = v
	return v
}

func (r *Request) Poll_Votesinvalid(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/votesinvalid", pollID)] = v
	return v
}

func (r *Request) Poll_Votesvalid(pollID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("poll/%d/votesvalid", pollID)] = v
	return v
}

func (r *Request) Projection_Content(projectionID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[fmt.Sprintf("projection/%d/content", projectionID)] = v
	return v
}

func (r *Request) Projection_ContentObjectID(projectionID int) *ValueMaybeString {
	v := &ValueMaybeString{request: r}
	r.requested[fmt.Sprintf("projection/%d/content_object_id", projectionID)] = v
	return v
}

func (r *Request) Projection_CurrentProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("projection/%d/current_projector_id", projectionID)] = v
	return v
}

func (r *Request) Projection_HistoryProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("projection/%d/history_projector_id", projectionID)] = v
	return v
}

func (r *Request) Projection_ID(projectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projection/%d/id", projectionID)] = v
	return v
}

func (r *Request) Projection_MeetingID(projectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projection/%d/meeting_id", projectionID)] = v
	return v
}

func (r *Request) Projection_Options(projectionID int) *ValueJSON {
	v := &ValueJSON{request: r}
	r.requested[fmt.Sprintf("projection/%d/options", projectionID)] = v
	return v
}

func (r *Request) Projection_PreviewProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("projection/%d/preview_projector_id", projectionID)] = v
	return v
}

func (r *Request) Projection_Stable(projectionID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("projection/%d/stable", projectionID)] = v
	return v
}

func (r *Request) Projection_Type(projectionID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projection/%d/type", projectionID)] = v
	return v
}

func (r *Request) Projection_Weight(projectionID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projection/%d/weight", projectionID)] = v
	return v
}

func (r *Request) ProjectorCountdown_CountdownTime(projectorCountdownID int) *ValueFloat {
	v := &ValueFloat{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/countdown_time", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_DefaultTime(projectorCountdownID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/default_time", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_Description(projectorCountdownID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/description", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_ID(projectorCountdownID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/id", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_MeetingID(projectorCountdownID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/meeting_id", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_ProjectionIDs(projectorCountdownID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/projection_ids", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_Running(projectorCountdownID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/running", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_Title(projectorCountdownID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/title", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_UsedAsListOfSpeakerCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/used_as_list_of_speaker_countdown_meeting_id", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorCountdown_UsedAsPollCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("projector_countdown/%d/used_as_poll_countdown_meeting_id", projectorCountdownID)] = v
	return v
}

func (r *Request) ProjectorMessage_ID(projectorMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector_message/%d/id", projectorMessageID)] = v
	return v
}

func (r *Request) ProjectorMessage_MeetingID(projectorMessageID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector_message/%d/meeting_id", projectorMessageID)] = v
	return v
}

func (r *Request) ProjectorMessage_Message(projectorMessageID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector_message/%d/message", projectorMessageID)] = v
	return v
}

func (r *Request) ProjectorMessage_ProjectionIDs(projectorMessageID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("projector_message/%d/projection_ids", projectorMessageID)] = v
	return v
}

func (r *Request) Projector_AspectRatioDenominator(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/aspect_ratio_denominator", projectorID)] = v
	return v
}

func (r *Request) Projector_AspectRatioNumerator(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/aspect_ratio_numerator", projectorID)] = v
	return v
}

func (r *Request) Projector_BackgroundColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector/%d/background_color", projectorID)] = v
	return v
}

func (r *Request) Projector_ChyronBackgroundColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector/%d/chyron_background_color", projectorID)] = v
	return v
}

func (r *Request) Projector_ChyronFontColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector/%d/chyron_font_color", projectorID)] = v
	return v
}

func (r *Request) Projector_Color(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector/%d/color", projectorID)] = v
	return v
}

func (r *Request) Projector_CurrentProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("projector/%d/current_projection_ids", projectorID)] = v
	return v
}

func (r *Request) Projector_HeaderBackgroundColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector/%d/header_background_color", projectorID)] = v
	return v
}

func (r *Request) Projector_HeaderFontColor(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector/%d/header_font_color", projectorID)] = v
	return v
}

func (r *Request) Projector_HeaderH1Color(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector/%d/header_h1_color", projectorID)] = v
	return v
}

func (r *Request) Projector_HistoryProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("projector/%d/history_projection_ids", projectorID)] = v
	return v
}

func (r *Request) Projector_ID(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/id", projectorID)] = v
	return v
}

func (r *Request) Projector_MeetingID(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/meeting_id", projectorID)] = v
	return v
}

func (r *Request) Projector_Name(projectorID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("projector/%d/name", projectorID)] = v
	return v
}

func (r *Request) Projector_PreviewProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("projector/%d/preview_projection_ids", projectorID)] = v
	return v
}

func (r *Request) Projector_Scale(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/scale", projectorID)] = v
	return v
}

func (r *Request) Projector_Scroll(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/scroll", projectorID)] = v
	return v
}

func (r *Request) Projector_ShowClock(projectorID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("projector/%d/show_clock", projectorID)] = v
	return v
}

func (r *Request) Projector_ShowHeaderFooter(projectorID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("projector/%d/show_header_footer", projectorID)] = v
	return v
}

func (r *Request) Projector_ShowLogo(projectorID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("projector/%d/show_logo", projectorID)] = v
	return v
}

func (r *Request) Projector_ShowTitle(projectorID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("projector/%d/show_title", projectorID)] = v
	return v
}

func (r *Request) Projector_UsedAsDefaultInMeetingIDTmpl(projectorID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("projector/%d/used_as_default_$_in_meeting_id", projectorID)] = v
	return v
}

func (r *Request) Projector_UsedAsDefaultInMeetingID(projectorID int, replacement string) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/used_as_default_$%s_in_meeting_id", projectorID, replacement)] = v
	return v
}

func (r *Request) Projector_UsedAsReferenceProjectorMeetingID(projectorID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/used_as_reference_projector_meeting_id", projectorID)] = v
	return v
}

func (r *Request) Projector_Width(projectorID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("projector/%d/width", projectorID)] = v
	return v
}

func (r *Request) Resource_Filesize(resourceID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("resource/%d/filesize", resourceID)] = v
	return v
}

func (r *Request) Resource_ID(resourceID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("resource/%d/id", resourceID)] = v
	return v
}

func (r *Request) Resource_Mimetype(resourceID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("resource/%d/mimetype", resourceID)] = v
	return v
}

func (r *Request) Resource_OrganizationID(resourceID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("resource/%d/organization_id", resourceID)] = v
	return v
}

func (r *Request) Resource_Token(resourceID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("resource/%d/token", resourceID)] = v
	return v
}

func (r *Request) Speaker_BeginTime(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("speaker/%d/begin_time", speakerID)] = v
	return v
}

func (r *Request) Speaker_EndTime(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("speaker/%d/end_time", speakerID)] = v
	return v
}

func (r *Request) Speaker_ID(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("speaker/%d/id", speakerID)] = v
	return v
}

func (r *Request) Speaker_ListOfSpeakersID(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("speaker/%d/list_of_speakers_id", speakerID)] = v
	return v
}

func (r *Request) Speaker_MeetingID(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("speaker/%d/meeting_id", speakerID)] = v
	return v
}

func (r *Request) Speaker_Note(speakerID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("speaker/%d/note", speakerID)] = v
	return v
}

func (r *Request) Speaker_PointOfOrder(speakerID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("speaker/%d/point_of_order", speakerID)] = v
	return v
}

func (r *Request) Speaker_SpeechState(speakerID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("speaker/%d/speech_state", speakerID)] = v
	return v
}

func (r *Request) Speaker_UserID(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("speaker/%d/user_id", speakerID)] = v
	return v
}

func (r *Request) Speaker_Weight(speakerID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("speaker/%d/weight", speakerID)] = v
	return v
}

func (r *Request) Tag_ID(tagID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("tag/%d/id", tagID)] = v
	return v
}

func (r *Request) Tag_MeetingID(tagID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("tag/%d/meeting_id", tagID)] = v
	return v
}

func (r *Request) Tag_Name(tagID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("tag/%d/name", tagID)] = v
	return v
}

func (r *Request) Tag_TaggedIDs(tagID int) *ValueStringSlice {
	v := &ValueStringSlice{request: r}
	r.requested[fmt.Sprintf("tag/%d/tagged_ids", tagID)] = v
	return v
}

func (r *Request) Theme_Accent100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_100", themeID)] = v
	return v
}

func (r *Request) Theme_Accent200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_200", themeID)] = v
	return v
}

func (r *Request) Theme_Accent300(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_300", themeID)] = v
	return v
}

func (r *Request) Theme_Accent400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_400", themeID)] = v
	return v
}

func (r *Request) Theme_Accent50(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_50", themeID)] = v
	return v
}

func (r *Request) Theme_Accent500(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_500", themeID)] = v
	return v
}

func (r *Request) Theme_Accent600(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_600", themeID)] = v
	return v
}

func (r *Request) Theme_Accent700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_700", themeID)] = v
	return v
}

func (r *Request) Theme_Accent800(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_800", themeID)] = v
	return v
}

func (r *Request) Theme_Accent900(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_900", themeID)] = v
	return v
}

func (r *Request) Theme_AccentA100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_a100", themeID)] = v
	return v
}

func (r *Request) Theme_AccentA200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_a200", themeID)] = v
	return v
}

func (r *Request) Theme_AccentA400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_a400", themeID)] = v
	return v
}

func (r *Request) Theme_AccentA700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/accent_a700", themeID)] = v
	return v
}

func (r *Request) Theme_ID(themeID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("theme/%d/id", themeID)] = v
	return v
}

func (r *Request) Theme_Name(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/name", themeID)] = v
	return v
}

func (r *Request) Theme_OrganizationID(themeID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("theme/%d/organization_id", themeID)] = v
	return v
}

func (r *Request) Theme_Primary100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_100", themeID)] = v
	return v
}

func (r *Request) Theme_Primary200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_200", themeID)] = v
	return v
}

func (r *Request) Theme_Primary300(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_300", themeID)] = v
	return v
}

func (r *Request) Theme_Primary400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_400", themeID)] = v
	return v
}

func (r *Request) Theme_Primary50(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_50", themeID)] = v
	return v
}

func (r *Request) Theme_Primary500(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_500", themeID)] = v
	return v
}

func (r *Request) Theme_Primary600(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_600", themeID)] = v
	return v
}

func (r *Request) Theme_Primary700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_700", themeID)] = v
	return v
}

func (r *Request) Theme_Primary800(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_800", themeID)] = v
	return v
}

func (r *Request) Theme_Primary900(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_900", themeID)] = v
	return v
}

func (r *Request) Theme_PrimaryA100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_a100", themeID)] = v
	return v
}

func (r *Request) Theme_PrimaryA200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_a200", themeID)] = v
	return v
}

func (r *Request) Theme_PrimaryA400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_a400", themeID)] = v
	return v
}

func (r *Request) Theme_PrimaryA700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/primary_a700", themeID)] = v
	return v
}

func (r *Request) Theme_ThemeForOrganizationID(themeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("theme/%d/theme_for_organization_id", themeID)] = v
	return v
}

func (r *Request) Theme_Warn100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_100", themeID)] = v
	return v
}

func (r *Request) Theme_Warn200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_200", themeID)] = v
	return v
}

func (r *Request) Theme_Warn300(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_300", themeID)] = v
	return v
}

func (r *Request) Theme_Warn400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_400", themeID)] = v
	return v
}

func (r *Request) Theme_Warn50(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_50", themeID)] = v
	return v
}

func (r *Request) Theme_Warn500(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_500", themeID)] = v
	return v
}

func (r *Request) Theme_Warn600(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_600", themeID)] = v
	return v
}

func (r *Request) Theme_Warn700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_700", themeID)] = v
	return v
}

func (r *Request) Theme_Warn800(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_800", themeID)] = v
	return v
}

func (r *Request) Theme_Warn900(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_900", themeID)] = v
	return v
}

func (r *Request) Theme_WarnA100(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_a100", themeID)] = v
	return v
}

func (r *Request) Theme_WarnA200(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_a200", themeID)] = v
	return v
}

func (r *Request) Theme_WarnA400(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_a400", themeID)] = v
	return v
}

func (r *Request) Theme_WarnA700(themeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("theme/%d/warn_a700", themeID)] = v
	return v
}

func (r *Request) Topic_AgendaItemID(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("topic/%d/agenda_item_id", topicID)] = v
	return v
}

func (r *Request) Topic_AttachmentIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("topic/%d/attachment_ids", topicID)] = v
	return v
}

func (r *Request) Topic_ID(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("topic/%d/id", topicID)] = v
	return v
}

func (r *Request) Topic_ListOfSpeakersID(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("topic/%d/list_of_speakers_id", topicID)] = v
	return v
}

func (r *Request) Topic_MeetingID(topicID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("topic/%d/meeting_id", topicID)] = v
	return v
}

func (r *Request) Topic_OptionIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("topic/%d/option_ids", topicID)] = v
	return v
}

func (r *Request) Topic_ProjectionIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("topic/%d/projection_ids", topicID)] = v
	return v
}

func (r *Request) Topic_TagIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("topic/%d/tag_ids", topicID)] = v
	return v
}

func (r *Request) Topic_Text(topicID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("topic/%d/text", topicID)] = v
	return v
}

func (r *Request) Topic_Title(topicID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("topic/%d/title", topicID)] = v
	return v
}

func (r *Request) User_AboutMeTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/about_me_$", userID)] = v
	return v
}

func (r *Request) User_AboutMe(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/about_me_$%d", userID, meetingID)] = v
	return v
}

func (r *Request) User_AssignmentCandidateIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/assignment_candidate_$_ids", userID)] = v
	return v
}

func (r *Request) User_AssignmentCandidateIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/assignment_candidate_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_CanChangeOwnPassword(userID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("user/%d/can_change_own_password", userID)] = v
	return v
}

func (r *Request) User_CommentTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/comment_$", userID)] = v
	return v
}

func (r *Request) User_Comment(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/comment_$%d", userID, meetingID)] = v
	return v
}

func (r *Request) User_CommitteeIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/committee_ids", userID)] = v
	return v
}

func (r *Request) User_CommitteeManagementLevelTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/committee_$_management_level", userID)] = v
	return v
}

func (r *Request) User_CommitteeManagementLevel(userID int, committeeID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/committee_$%d_management_level", userID, committeeID)] = v
	return v
}

func (r *Request) User_DefaultNumber(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/default_number", userID)] = v
	return v
}

func (r *Request) User_DefaultPassword(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/default_password", userID)] = v
	return v
}

func (r *Request) User_DefaultStructureLevel(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/default_structure_level", userID)] = v
	return v
}

func (r *Request) User_DefaultVoteWeight(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/default_vote_weight", userID)] = v
	return v
}

func (r *Request) User_Email(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/email", userID)] = v
	return v
}

func (r *Request) User_FirstName(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/first_name", userID)] = v
	return v
}

func (r *Request) User_Gender(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/gender", userID)] = v
	return v
}

func (r *Request) User_GroupIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/group_$_ids", userID)] = v
	return v
}

func (r *Request) User_GroupIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/group_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_ID(userID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("user/%d/id", userID)] = v
	return v
}

func (r *Request) User_IsActive(userID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("user/%d/is_active", userID)] = v
	return v
}

func (r *Request) User_IsDemoUser(userID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("user/%d/is_demo_user", userID)] = v
	return v
}

func (r *Request) User_IsPhysicalPerson(userID int) *ValueBool {
	v := &ValueBool{request: r}
	r.requested[fmt.Sprintf("user/%d/is_physical_person", userID)] = v
	return v
}

func (r *Request) User_IsPresentInMeetingIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/is_present_in_meeting_ids", userID)] = v
	return v
}

func (r *Request) User_LastEmailSend(userID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("user/%d/last_email_send", userID)] = v
	return v
}

func (r *Request) User_LastName(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/last_name", userID)] = v
	return v
}

func (r *Request) User_MeetingIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/meeting_ids", userID)] = v
	return v
}

func (r *Request) User_NumberTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/number_$", userID)] = v
	return v
}

func (r *Request) User_Number(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/number_$%d", userID, meetingID)] = v
	return v
}

func (r *Request) User_OptionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/option_$_ids", userID)] = v
	return v
}

func (r *Request) User_OptionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/option_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_OrganizationManagementLevel(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/organization_management_level", userID)] = v
	return v
}

func (r *Request) User_Password(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/password", userID)] = v
	return v
}

func (r *Request) User_PersonalNoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/personal_note_$_ids", userID)] = v
	return v
}

func (r *Request) User_PersonalNoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/personal_note_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_PollVotedIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/poll_voted_$_ids", userID)] = v
	return v
}

func (r *Request) User_PollVotedIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/poll_voted_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_ProjectionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/projection_$_ids", userID)] = v
	return v
}

func (r *Request) User_ProjectionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/projection_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_SpeakerIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/speaker_$_ids", userID)] = v
	return v
}

func (r *Request) User_SpeakerIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/speaker_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_StructureLevelTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/structure_level_$", userID)] = v
	return v
}

func (r *Request) User_StructureLevel(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/structure_level_$%d", userID, meetingID)] = v
	return v
}

func (r *Request) User_SubmittedMotionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/submitted_motion_$_ids", userID)] = v
	return v
}

func (r *Request) User_SubmittedMotionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/submitted_motion_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_SupportedMotionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/supported_motion_$_ids", userID)] = v
	return v
}

func (r *Request) User_SupportedMotionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/supported_motion_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_Title(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/title", userID)] = v
	return v
}

func (r *Request) User_Username(userID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/username", userID)] = v
	return v
}

func (r *Request) User_VoteDelegatedToIDTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_delegated_$_to_id", userID)] = v
	return v
}

func (r *Request) User_VoteDelegatedToID(userID int, meetingID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_delegated_$%d_to_id", userID, meetingID)] = v
	return v
}

func (r *Request) User_VoteDelegatedVoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_delegated_vote_$_ids", userID)] = v
	return v
}

func (r *Request) User_VoteDelegatedVoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_delegated_vote_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_VoteDelegationsFromIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_delegations_$_from_ids", userID)] = v
	return v
}

func (r *Request) User_VoteDelegationsFromIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_delegations_$%d_from_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_VoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_$_ids", userID)] = v
	return v
}

func (r *Request) User_VoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_$%d_ids", userID, meetingID)] = v
	return v
}

func (r *Request) User_VoteWeightTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_weight_$", userID)] = v
	return v
}

func (r *Request) User_VoteWeight(userID int, meetingID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("user/%d/vote_weight_$%d", userID, meetingID)] = v
	return v
}

func (r *Request) Vote_DelegatedUserID(voteID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("vote/%d/delegated_user_id", voteID)] = v
	return v
}

func (r *Request) Vote_ID(voteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("vote/%d/id", voteID)] = v
	return v
}

func (r *Request) Vote_MeetingID(voteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("vote/%d/meeting_id", voteID)] = v
	return v
}

func (r *Request) Vote_OptionID(voteID int) *ValueInt {
	v := &ValueInt{request: r}
	r.requested[fmt.Sprintf("vote/%d/option_id", voteID)] = v
	return v
}

func (r *Request) Vote_UserID(voteID int) *ValueMaybeInt {
	v := &ValueMaybeInt{request: r}
	r.requested[fmt.Sprintf("vote/%d/user_id", voteID)] = v
	return v
}

func (r *Request) Vote_UserToken(voteID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("vote/%d/user_token", voteID)] = v
	return v
}

func (r *Request) Vote_Value(voteID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("vote/%d/value", voteID)] = v
	return v
}

func (r *Request) Vote_Weight(voteID int) *ValueString {
	v := &ValueString{request: r}
	r.requested[fmt.Sprintf("vote/%d/weight", voteID)] = v
	return v
}
