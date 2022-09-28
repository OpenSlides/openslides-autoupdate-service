// Code generated from models.yml DO NOT EDIT.
package dsfetch

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// ValueBool is a lazy value from the datastore.
type ValueBool struct {
	value    bool
	isNull   bool
	executed bool

	lazies []*bool

	fetch *Fetch
}

// Value returns the value.
func (v *ValueBool) Value(ctx context.Context) (bool, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return false
	}

	if v.executed {
		return v.value
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueFloat) Value(ctx context.Context) (float32, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return 0
	}

	if v.executed {
		return v.value
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueIDSlice) Value(ctx context.Context) ([]int, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return nil
	}

	if v.executed {
		return v.value
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueInt) Value(ctx context.Context) (int, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return 0
	}

	if v.executed {
		return v.value
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueIntSlice) Value(ctx context.Context) ([]int, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return nil
	}

	if v.executed {
		return v.value
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueJSON) Value(ctx context.Context) (json.RawMessage, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return nil
	}

	if v.executed {
		return v.value
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueMaybeInt) Value(ctx context.Context) (int, bool, error) {
	if v.executed {
		return v.value, !v.isNull, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return 0, false
	}

	if v.executed {
		return v.value, !v.isNull
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueMaybeString) Value(ctx context.Context) (string, bool, error) {
	if v.executed {
		return v.value, !v.isNull, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return "", false
	}

	if v.executed {
		return v.value, !v.isNull
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueString) Value(ctx context.Context) (string, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return ""
	}

	if v.executed {
		return v.value
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
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

	fetch *Fetch
}

// Value returns the value.
func (v *ValueStringSlice) Value(ctx context.Context) ([]string, error) {
	if v.executed {
		return v.value, nil
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
	if v.fetch.err != nil {
		return nil
	}

	if v.executed {
		return v.value
	}

	if err := v.fetch.Execute(ctx); err != nil {
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
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = v.value
	}

	v.executed = true
	return nil
}

func (r *Fetch) ActionWorker_Created(actionWorkerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "action_worker", ID: actionWorkerID, Field: "created"}] = v
	return v
}

func (r *Fetch) ActionWorker_ID(actionWorkerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "action_worker", ID: actionWorkerID, Field: "id"}] = v
	return v
}

func (r *Fetch) ActionWorker_Name(actionWorkerID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "action_worker", ID: actionWorkerID, Field: "name"}] = v
	return v
}

func (r *Fetch) ActionWorker_Result(actionWorkerID int) *ValueJSON {
	v := &ValueJSON{fetch: r}
	r.requested[datastore.Key{Collection: "action_worker", ID: actionWorkerID, Field: "result"}] = v
	return v
}

func (r *Fetch) ActionWorker_State(actionWorkerID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "action_worker", ID: actionWorkerID, Field: "state"}] = v
	return v
}

func (r *Fetch) ActionWorker_Timestamp(actionWorkerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "action_worker", ID: actionWorkerID, Field: "timestamp"}] = v
	return v
}

func (r *Fetch) AgendaItem_ChildIDs(agendaItemID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "child_ids"}] = v
	return v
}

func (r *Fetch) AgendaItem_Closed(agendaItemID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "closed"}] = v
	return v
}

func (r *Fetch) AgendaItem_Comment(agendaItemID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "comment"}] = v
	return v
}

func (r *Fetch) AgendaItem_ContentObjectID(agendaItemID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "content_object_id"}] = v
	return v
}

func (r *Fetch) AgendaItem_Duration(agendaItemID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "duration"}] = v
	return v
}

func (r *Fetch) AgendaItem_ID(agendaItemID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "id"}] = v
	return v
}

func (r *Fetch) AgendaItem_IsHidden(agendaItemID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "is_hidden"}] = v
	return v
}

func (r *Fetch) AgendaItem_IsInternal(agendaItemID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "is_internal"}] = v
	return v
}

func (r *Fetch) AgendaItem_ItemNumber(agendaItemID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "item_number"}] = v
	return v
}

func (r *Fetch) AgendaItem_Level(agendaItemID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "level"}] = v
	return v
}

func (r *Fetch) AgendaItem_MeetingID(agendaItemID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) AgendaItem_ParentID(agendaItemID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "parent_id"}] = v
	return v
}

func (r *Fetch) AgendaItem_ProjectionIDs(agendaItemID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) AgendaItem_TagIDs(agendaItemID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "tag_ids"}] = v
	return v
}

func (r *Fetch) AgendaItem_Type(agendaItemID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "type"}] = v
	return v
}

func (r *Fetch) AgendaItem_Weight(agendaItemID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "agenda_item", ID: agendaItemID, Field: "weight"}] = v
	return v
}

func (r *Fetch) AssignmentCandidate_AssignmentID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "assignment_id"}] = v
	return v
}

func (r *Fetch) AssignmentCandidate_ID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "id"}] = v
	return v
}

func (r *Fetch) AssignmentCandidate_MeetingID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) AssignmentCandidate_UserID(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "user_id"}] = v
	return v
}

func (r *Fetch) AssignmentCandidate_Weight(assignmentCandidateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "weight"}] = v
	return v
}

func (r *Fetch) Assignment_AgendaItemID(assignmentID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "agenda_item_id"}] = v
	return v
}

func (r *Fetch) Assignment_AttachmentIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "attachment_ids"}] = v
	return v
}

func (r *Fetch) Assignment_CandidateIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "candidate_ids"}] = v
	return v
}

func (r *Fetch) Assignment_DefaultPollDescription(assignmentID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "default_poll_description"}] = v
	return v
}

func (r *Fetch) Assignment_Description(assignmentID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "description"}] = v
	return v
}

func (r *Fetch) Assignment_ID(assignmentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "id"}] = v
	return v
}

func (r *Fetch) Assignment_ListOfSpeakersID(assignmentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "list_of_speakers_id"}] = v
	return v
}

func (r *Fetch) Assignment_MeetingID(assignmentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Assignment_NumberPollCandidates(assignmentID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "number_poll_candidates"}] = v
	return v
}

func (r *Fetch) Assignment_OpenPosts(assignmentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "open_posts"}] = v
	return v
}

func (r *Fetch) Assignment_Phase(assignmentID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "phase"}] = v
	return v
}

func (r *Fetch) Assignment_PollIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "poll_ids"}] = v
	return v
}

func (r *Fetch) Assignment_ProjectionIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) Assignment_SequentialNumber(assignmentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) Assignment_TagIDs(assignmentID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "tag_ids"}] = v
	return v
}

func (r *Fetch) Assignment_Title(assignmentID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "assignment", ID: assignmentID, Field: "title"}] = v
	return v
}

func (r *Fetch) ChatGroup_ChatMessageIDs(chatGroupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "chat_group", ID: chatGroupID, Field: "chat_message_ids"}] = v
	return v
}

func (r *Fetch) ChatGroup_ID(chatGroupID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "chat_group", ID: chatGroupID, Field: "id"}] = v
	return v
}

func (r *Fetch) ChatGroup_MeetingID(chatGroupID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "chat_group", ID: chatGroupID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) ChatGroup_Name(chatGroupID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "chat_group", ID: chatGroupID, Field: "name"}] = v
	return v
}

func (r *Fetch) ChatGroup_ReadGroupIDs(chatGroupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "chat_group", ID: chatGroupID, Field: "read_group_ids"}] = v
	return v
}

func (r *Fetch) ChatGroup_Weight(chatGroupID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "chat_group", ID: chatGroupID, Field: "weight"}] = v
	return v
}

func (r *Fetch) ChatGroup_WriteGroupIDs(chatGroupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "chat_group", ID: chatGroupID, Field: "write_group_ids"}] = v
	return v
}

func (r *Fetch) ChatMessage_ChatGroupID(chatMessageID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "chat_message", ID: chatMessageID, Field: "chat_group_id"}] = v
	return v
}

func (r *Fetch) ChatMessage_Content(chatMessageID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "chat_message", ID: chatMessageID, Field: "content"}] = v
	return v
}

func (r *Fetch) ChatMessage_Created(chatMessageID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "chat_message", ID: chatMessageID, Field: "created"}] = v
	return v
}

func (r *Fetch) ChatMessage_ID(chatMessageID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "chat_message", ID: chatMessageID, Field: "id"}] = v
	return v
}

func (r *Fetch) ChatMessage_MeetingID(chatMessageID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "chat_message", ID: chatMessageID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) ChatMessage_UserID(chatMessageID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "chat_message", ID: chatMessageID, Field: "user_id"}] = v
	return v
}

func (r *Fetch) Committee_DefaultMeetingID(committeeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "default_meeting_id"}] = v
	return v
}

func (r *Fetch) Committee_Description(committeeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "description"}] = v
	return v
}

func (r *Fetch) Committee_ForwardToCommitteeIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "forward_to_committee_ids"}] = v
	return v
}

func (r *Fetch) Committee_ForwardingUserID(committeeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "forwarding_user_id"}] = v
	return v
}

func (r *Fetch) Committee_ID(committeeID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "id"}] = v
	return v
}

func (r *Fetch) Committee_MeetingIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "meeting_ids"}] = v
	return v
}

func (r *Fetch) Committee_Name(committeeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "name"}] = v
	return v
}

func (r *Fetch) Committee_OrganizationID(committeeID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "organization_id"}] = v
	return v
}

func (r *Fetch) Committee_OrganizationTagIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "organization_tag_ids"}] = v
	return v
}

func (r *Fetch) Committee_ReceiveForwardingsFromCommitteeIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "receive_forwardings_from_committee_ids"}] = v
	return v
}

func (r *Fetch) Committee_UserIDs(committeeID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "user_ids"}] = v
	return v
}

func (r *Fetch) Committee_UserManagementLevelTmpl(committeeID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: "user_$_management_level"}] = v
	return v
}

func (r *Fetch) Committee_UserManagementLevel(committeeID int, replacement string) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "committee", ID: committeeID, Field: fmt.Sprintf("user_$%s_management_level", replacement)}] = v
	return v
}

func (r *Fetch) Group_AdminGroupForMeetingID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "admin_group_for_meeting_id"}] = v
	return v
}

func (r *Fetch) Group_DefaultGroupForMeetingID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "default_group_for_meeting_id"}] = v
	return v
}

func (r *Fetch) Group_ID(groupID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "id"}] = v
	return v
}

func (r *Fetch) Group_MediafileAccessGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "mediafile_access_group_ids"}] = v
	return v
}

func (r *Fetch) Group_MediafileInheritedAccessGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "mediafile_inherited_access_group_ids"}] = v
	return v
}

func (r *Fetch) Group_MeetingID(groupID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Group_Name(groupID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "name"}] = v
	return v
}

func (r *Fetch) Group_Permissions(groupID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "permissions"}] = v
	return v
}

func (r *Fetch) Group_PollIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "poll_ids"}] = v
	return v
}

func (r *Fetch) Group_ReadChatGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "read_chat_group_ids"}] = v
	return v
}

func (r *Fetch) Group_ReadCommentSectionIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "read_comment_section_ids"}] = v
	return v
}

func (r *Fetch) Group_UsedAsAssignmentPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "used_as_assignment_poll_default_id"}] = v
	return v
}

func (r *Fetch) Group_UsedAsMotionPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "used_as_motion_poll_default_id"}] = v
	return v
}

func (r *Fetch) Group_UsedAsPollDefaultID(groupID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "used_as_poll_default_id"}] = v
	return v
}

func (r *Fetch) Group_UserIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "user_ids"}] = v
	return v
}

func (r *Fetch) Group_Weight(groupID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "weight"}] = v
	return v
}

func (r *Fetch) Group_WriteChatGroupIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "write_chat_group_ids"}] = v
	return v
}

func (r *Fetch) Group_WriteCommentSectionIDs(groupID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "group", ID: groupID, Field: "write_comment_section_ids"}] = v
	return v
}

func (r *Fetch) ListOfSpeakers_Closed(listOfSpeakersID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "closed"}] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ContentObjectID(listOfSpeakersID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "content_object_id"}] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ID(listOfSpeakersID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "id"}] = v
	return v
}

func (r *Fetch) ListOfSpeakers_MeetingID(listOfSpeakersID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ProjectionIDs(listOfSpeakersID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) ListOfSpeakers_SequentialNumber(listOfSpeakersID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) ListOfSpeakers_SpeakerIDs(listOfSpeakersID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "speaker_ids"}] = v
	return v
}

func (r *Fetch) Mediafile_AccessGroupIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "access_group_ids"}] = v
	return v
}

func (r *Fetch) Mediafile_AttachmentIDs(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "attachment_ids"}] = v
	return v
}

func (r *Fetch) Mediafile_ChildIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "child_ids"}] = v
	return v
}

func (r *Fetch) Mediafile_CreateTimestamp(mediafileID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "create_timestamp"}] = v
	return v
}

func (r *Fetch) Mediafile_Filename(mediafileID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "filename"}] = v
	return v
}

func (r *Fetch) Mediafile_Filesize(mediafileID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "filesize"}] = v
	return v
}

func (r *Fetch) Mediafile_ID(mediafileID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "id"}] = v
	return v
}

func (r *Fetch) Mediafile_InheritedAccessGroupIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "inherited_access_group_ids"}] = v
	return v
}

func (r *Fetch) Mediafile_IsDirectory(mediafileID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "is_directory"}] = v
	return v
}

func (r *Fetch) Mediafile_IsPublic(mediafileID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "is_public"}] = v
	return v
}

func (r *Fetch) Mediafile_ListOfSpeakersID(mediafileID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "list_of_speakers_id"}] = v
	return v
}

func (r *Fetch) Mediafile_Mimetype(mediafileID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "mimetype"}] = v
	return v
}

func (r *Fetch) Mediafile_OwnerID(mediafileID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "owner_id"}] = v
	return v
}

func (r *Fetch) Mediafile_ParentID(mediafileID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "parent_id"}] = v
	return v
}

func (r *Fetch) Mediafile_PdfInformation(mediafileID int) *ValueJSON {
	v := &ValueJSON{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "pdf_information"}] = v
	return v
}

func (r *Fetch) Mediafile_ProjectionIDs(mediafileID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) Mediafile_Title(mediafileID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "title"}] = v
	return v
}

func (r *Fetch) Mediafile_Token(mediafileID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "token"}] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontInMeetingIDTmpl(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_$_in_meeting_id"}] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontInMeetingID(mediafileID int, replacement string) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: fmt.Sprintf("used_as_font_$%s_in_meeting_id", replacement)}] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoInMeetingIDTmpl(mediafileID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_$_in_meeting_id"}] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoInMeetingID(mediafileID int, replacement string) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "mediafile", ID: mediafileID, Field: fmt.Sprintf("used_as_logo_$%s_in_meeting_id", replacement)}] = v
	return v
}

func (r *Fetch) Meeting_AdminGroupID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "admin_group_id"}] = v
	return v
}

func (r *Fetch) Meeting_AgendaEnableNumbering(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "agenda_enable_numbering"}] = v
	return v
}

func (r *Fetch) Meeting_AgendaItemCreation(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "agenda_item_creation"}] = v
	return v
}

func (r *Fetch) Meeting_AgendaItemIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "agenda_item_ids"}] = v
	return v
}

func (r *Fetch) Meeting_AgendaNewItemsDefaultVisibility(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "agenda_new_items_default_visibility"}] = v
	return v
}

func (r *Fetch) Meeting_AgendaNumberPrefix(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "agenda_number_prefix"}] = v
	return v
}

func (r *Fetch) Meeting_AgendaNumeralSystem(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "agenda_numeral_system"}] = v
	return v
}

func (r *Fetch) Meeting_AgendaShowInternalItemsOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "agenda_show_internal_items_on_projector"}] = v
	return v
}

func (r *Fetch) Meeting_AgendaShowSubtitles(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "agenda_show_subtitles"}] = v
	return v
}

func (r *Fetch) Meeting_AllProjectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "all_projection_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ApplauseEnable(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "applause_enable"}] = v
	return v
}

func (r *Fetch) Meeting_ApplauseMaxAmount(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "applause_max_amount"}] = v
	return v
}

func (r *Fetch) Meeting_ApplauseMinAmount(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "applause_min_amount"}] = v
	return v
}

func (r *Fetch) Meeting_ApplauseParticleImageUrl(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "applause_particle_image_url"}] = v
	return v
}

func (r *Fetch) Meeting_ApplauseShowLevel(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "applause_show_level"}] = v
	return v
}

func (r *Fetch) Meeting_ApplauseTimeout(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "applause_timeout"}] = v
	return v
}

func (r *Fetch) Meeting_ApplauseType(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "applause_type"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentCandidateIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_candidate_ids"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_ids"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollAddCandidatesToListOfSpeakers(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_add_candidates_to_list_of_speakers"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_ballot_paper_number"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_ballot_paper_selection"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_100_percent_base"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultBackend(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_backend"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_group_ids"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultMethod(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_method"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultType(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_type"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollEnableMaxVotesPerOption(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_enable_max_votes_per_option"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollSortPollResultByVotes(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_sort_poll_result_by_votes"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentsExportPreamble(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignments_export_preamble"}] = v
	return v
}

func (r *Fetch) Meeting_AssignmentsExportTitle(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "assignments_export_title"}] = v
	return v
}

func (r *Fetch) Meeting_ChatGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "chat_group_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ChatMessageIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "chat_message_ids"}] = v
	return v
}

func (r *Fetch) Meeting_CommitteeID(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "committee_id"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceAutoConnect(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_auto_connect"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceAutoConnectNextSpeakers(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_auto_connect_next_speakers"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceEnableHelpdesk(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_enable_helpdesk"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceLosRestriction(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_los_restriction"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceOpenMicrophone(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_open_microphone"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceOpenVideo(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_open_video"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceShow(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_show"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceStreamPosterUrl(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_stream_poster_url"}] = v
	return v
}

func (r *Fetch) Meeting_ConferenceStreamUrl(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "conference_stream_url"}] = v
	return v
}

func (r *Fetch) Meeting_CustomTranslations(meetingID int) *ValueJSON {
	v := &ValueJSON{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "custom_translations"}] = v
	return v
}

func (r *Fetch) Meeting_DefaultGroupID(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "default_group_id"}] = v
	return v
}

func (r *Fetch) Meeting_DefaultMeetingForCommitteeID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "default_meeting_for_committee_id"}] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_$_id"}] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: fmt.Sprintf("default_projector_$%s_id", replacement)}] = v
	return v
}

func (r *Fetch) Meeting_Description(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "description"}] = v
	return v
}

func (r *Fetch) Meeting_EnableAnonymous(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "enable_anonymous"}] = v
	return v
}

func (r *Fetch) Meeting_EndTime(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "end_time"}] = v
	return v
}

func (r *Fetch) Meeting_ExportCsvEncoding(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_csv_encoding"}] = v
	return v
}

func (r *Fetch) Meeting_ExportCsvSeparator(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_csv_separator"}] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfFontsize(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_fontsize"}] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfLineHeight(meetingID int) *ValueFloat {
	v := &ValueFloat{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_line_height"}] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginBottom(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_page_margin_bottom"}] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginLeft(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_page_margin_left"}] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginRight(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_page_margin_right"}] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginTop(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_page_margin_top"}] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPagenumberAlignment(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_pagenumber_alignment"}] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPagesize(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_pagesize"}] = v
	return v
}

func (r *Fetch) Meeting_FontIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "font_$_id"}] = v
	return v
}

func (r *Fetch) Meeting_FontID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: fmt.Sprintf("font_$%s_id", replacement)}] = v
	return v
}

func (r *Fetch) Meeting_ForwardedMotionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "forwarded_motion_ids"}] = v
	return v
}

func (r *Fetch) Meeting_GroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "group_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ID(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "id"}] = v
	return v
}

func (r *Fetch) Meeting_ImportedAt(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "imported_at"}] = v
	return v
}

func (r *Fetch) Meeting_IsActiveInOrganizationID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "is_active_in_organization_id"}] = v
	return v
}

func (r *Fetch) Meeting_IsArchivedInOrganizationID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "is_archived_in_organization_id"}] = v
	return v
}

func (r *Fetch) Meeting_JitsiDomain(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "jitsi_domain"}] = v
	return v
}

func (r *Fetch) Meeting_JitsiRoomName(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "jitsi_room_name"}] = v
	return v
}

func (r *Fetch) Meeting_JitsiRoomPassword(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "jitsi_room_password"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersAmountLastOnProjector(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_amount_last_on_projector"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersAmountNextOnProjector(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_amount_next_on_projector"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCanSetContributionSelf(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_can_set_contribution_self"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCountdownID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_countdown_id"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCoupleCountdown(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_couple_countdown"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_enable_point_of_order_speakers"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersEnableProContraSpeech(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_enable_pro_contra_speech"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersInitiallyClosed(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_initially_closed"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersPresentUsersOnly(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_present_users_only"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_show_amount_of_speakers_on_slide"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersShowFirstContribution(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_show_first_contribution"}] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersSpeakerNoteForEveryone(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_speaker_note_for_everyone"}] = v
	return v
}

func (r *Fetch) Meeting_Location(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "location"}] = v
	return v
}

func (r *Fetch) Meeting_LogoIDTmpl(meetingID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "logo_$_id"}] = v
	return v
}

func (r *Fetch) Meeting_LogoID(meetingID int, replacement string) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: fmt.Sprintf("logo_$%s_id", replacement)}] = v
	return v
}

func (r *Fetch) Meeting_MediafileIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "mediafile_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionBlockIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_block_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionCategoryIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_category_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionChangeRecommendationIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_change_recommendation_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionCommentIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_comment_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionCommentSectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_comment_section_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionPollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_ballot_paper_number"}] = v
	return v
}

func (r *Fetch) Meeting_MotionPollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_ballot_paper_selection"}] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_default_100_percent_base"}] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultBackend(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_default_backend"}] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_default_group_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultType(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_default_type"}] = v
	return v
}

func (r *Fetch) Meeting_MotionStateIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_state_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionStatuteParagraphIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_statute_paragraph_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionSubmitterIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_submitter_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionWorkflowIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motion_workflow_ids"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsEnabled(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_enabled"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsInMainList(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_in_main_list"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsMultipleParagraphs(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_multiple_paragraphs"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsOfAmendments(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_of_amendments"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsPrefix(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_prefix"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsTextMode(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_text_mode"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultAmendmentWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_amendment_workflow_id"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultLineNumbering(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_line_numbering"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultSorting(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_sorting"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultStatuteAmendmentWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_statute_amendment_workflow_id"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultWorkflowID(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_workflow_id"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableReasonOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_enable_reason_on_projector"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableRecommendationOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_enable_recommendation_on_projector"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableSideboxOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_enable_sidebox_on_projector"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableTextOnProjector(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_enable_text_on_projector"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportFollowRecommendation(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_export_follow_recommendation"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportPreamble(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_export_preamble"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportSubmitterRecommendation(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_export_submitter_recommendation"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportTitle(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_export_title"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsLineLength(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_line_length"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberMinDigits(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_number_min_digits"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberType(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_number_type"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberWithBlank(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_number_with_blank"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsPreamble(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_preamble"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsReasonRequired(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_reason_required"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsRecommendationTextMode(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_recommendation_text_mode"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsRecommendationsBy(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_recommendations_by"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsShowReferringMotions(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_show_referring_motions"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsShowSequentialNumber(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_show_sequential_number"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsStatuteRecommendationsBy(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_statute_recommendations_by"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsStatutesEnabled(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_statutes_enabled"}] = v
	return v
}

func (r *Fetch) Meeting_MotionsSupportersMinAmount(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "motions_supporters_min_amount"}] = v
	return v
}

func (r *Fetch) Meeting_Name(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "name"}] = v
	return v
}

func (r *Fetch) Meeting_OptionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "option_ids"}] = v
	return v
}

func (r *Fetch) Meeting_OrganizationTagIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "organization_tag_ids"}] = v
	return v
}

func (r *Fetch) Meeting_PersonalNoteIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "personal_note_ids"}] = v
	return v
}

func (r *Fetch) Meeting_PollBallotPaperNumber(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_ballot_paper_number"}] = v
	return v
}

func (r *Fetch) Meeting_PollBallotPaperSelection(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_ballot_paper_selection"}] = v
	return v
}

func (r *Fetch) Meeting_PollCountdownID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_countdown_id"}] = v
	return v
}

func (r *Fetch) Meeting_PollCoupleCountdown(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_couple_countdown"}] = v
	return v
}

func (r *Fetch) Meeting_PollDefault100PercentBase(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_100_percent_base"}] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultBackend(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_backend"}] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_group_ids"}] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultMethod(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_method"}] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultType(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_type"}] = v
	return v
}

func (r *Fetch) Meeting_PollIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_ids"}] = v
	return v
}

func (r *Fetch) Meeting_PollSortPollResultByVotes(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "poll_sort_poll_result_by_votes"}] = v
	return v
}

func (r *Fetch) Meeting_PresentUserIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "present_user_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ProjectionIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownDefaultTime(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "projector_countdown_default_time"}] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "projector_countdown_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownWarningTime(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "projector_countdown_warning_time"}] = v
	return v
}

func (r *Fetch) Meeting_ProjectorIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "projector_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ProjectorMessageIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "projector_message_ids"}] = v
	return v
}

func (r *Fetch) Meeting_ReferenceProjectorID(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "reference_projector_id"}] = v
	return v
}

func (r *Fetch) Meeting_SpeakerIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "speaker_ids"}] = v
	return v
}

func (r *Fetch) Meeting_StartTime(meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "start_time"}] = v
	return v
}

func (r *Fetch) Meeting_TagIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "tag_ids"}] = v
	return v
}

func (r *Fetch) Meeting_TemplateForOrganizationID(meetingID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "template_for_organization_id"}] = v
	return v
}

func (r *Fetch) Meeting_TopicIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "topic_ids"}] = v
	return v
}

func (r *Fetch) Meeting_UserIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "user_ids"}] = v
	return v
}

func (r *Fetch) Meeting_UsersAllowSelfSetPresent(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_allow_self_set_present"}] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailBody(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_email_body"}] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailReplyto(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_email_replyto"}] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailSender(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_email_sender"}] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailSubject(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_email_subject"}] = v
	return v
}

func (r *Fetch) Meeting_UsersEnablePresenceView(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_enable_presence_view"}] = v
	return v
}

func (r *Fetch) Meeting_UsersEnableVoteWeight(meetingID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_enable_vote_weight"}] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWelcometext(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_welcometext"}] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWelcometitle(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_welcometitle"}] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanEncryption(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_wlan_encryption"}] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanPassword(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_wlan_password"}] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanSsid(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_wlan_ssid"}] = v
	return v
}

func (r *Fetch) Meeting_VoteIDs(meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "vote_ids"}] = v
	return v
}

func (r *Fetch) Meeting_WelcomeText(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "welcome_text"}] = v
	return v
}

func (r *Fetch) Meeting_WelcomeTitle(meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "meeting", ID: meetingID, Field: "welcome_title"}] = v
	return v
}

func (r *Fetch) MotionBlock_AgendaItemID(motionBlockID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "agenda_item_id"}] = v
	return v
}

func (r *Fetch) MotionBlock_ID(motionBlockID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionBlock_Internal(motionBlockID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "internal"}] = v
	return v
}

func (r *Fetch) MotionBlock_ListOfSpeakersID(motionBlockID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "list_of_speakers_id"}] = v
	return v
}

func (r *Fetch) MotionBlock_MeetingID(motionBlockID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionBlock_MotionIDs(motionBlockID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "motion_ids"}] = v
	return v
}

func (r *Fetch) MotionBlock_ProjectionIDs(motionBlockID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) MotionBlock_SequentialNumber(motionBlockID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) MotionBlock_Title(motionBlockID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_block", ID: motionBlockID, Field: "title"}] = v
	return v
}

func (r *Fetch) MotionCategory_ChildIDs(motionCategoryID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "child_ids"}] = v
	return v
}

func (r *Fetch) MotionCategory_ID(motionCategoryID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionCategory_Level(motionCategoryID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "level"}] = v
	return v
}

func (r *Fetch) MotionCategory_MeetingID(motionCategoryID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionCategory_MotionIDs(motionCategoryID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "motion_ids"}] = v
	return v
}

func (r *Fetch) MotionCategory_Name(motionCategoryID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "name"}] = v
	return v
}

func (r *Fetch) MotionCategory_ParentID(motionCategoryID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "parent_id"}] = v
	return v
}

func (r *Fetch) MotionCategory_Prefix(motionCategoryID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "prefix"}] = v
	return v
}

func (r *Fetch) MotionCategory_SequentialNumber(motionCategoryID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) MotionCategory_Weight(motionCategoryID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_category", ID: motionCategoryID, Field: "weight"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_CreationTime(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "creation_time"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_ID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Internal(motionChangeRecommendationID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "internal"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_LineFrom(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "line_from"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_LineTo(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "line_to"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_MeetingID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_MotionID(motionChangeRecommendationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "motion_id"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_OtherDescription(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "other_description"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Rejected(motionChangeRecommendationID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "rejected"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Text(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "text"}] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Type(motionChangeRecommendationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "type"}] = v
	return v
}

func (r *Fetch) MotionCommentSection_CommentIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "comment_ids"}] = v
	return v
}

func (r *Fetch) MotionCommentSection_ID(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionCommentSection_MeetingID(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionCommentSection_Name(motionCommentSectionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "name"}] = v
	return v
}

func (r *Fetch) MotionCommentSection_ReadGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "read_group_ids"}] = v
	return v
}

func (r *Fetch) MotionCommentSection_SequentialNumber(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) MotionCommentSection_Weight(motionCommentSectionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "weight"}] = v
	return v
}

func (r *Fetch) MotionCommentSection_WriteGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "write_group_ids"}] = v
	return v
}

func (r *Fetch) MotionComment_Comment(motionCommentID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment", ID: motionCommentID, Field: "comment"}] = v
	return v
}

func (r *Fetch) MotionComment_ID(motionCommentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment", ID: motionCommentID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionComment_MeetingID(motionCommentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment", ID: motionCommentID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionComment_MotionID(motionCommentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment", ID: motionCommentID, Field: "motion_id"}] = v
	return v
}

func (r *Fetch) MotionComment_SectionID(motionCommentID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_comment", ID: motionCommentID, Field: "section_id"}] = v
	return v
}

func (r *Fetch) MotionState_AllowCreatePoll(motionStateID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "allow_create_poll"}] = v
	return v
}

func (r *Fetch) MotionState_AllowMotionForwarding(motionStateID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "allow_motion_forwarding"}] = v
	return v
}

func (r *Fetch) MotionState_AllowSubmitterEdit(motionStateID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "allow_submitter_edit"}] = v
	return v
}

func (r *Fetch) MotionState_AllowSupport(motionStateID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "allow_support"}] = v
	return v
}

func (r *Fetch) MotionState_CssClass(motionStateID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "css_class"}] = v
	return v
}

func (r *Fetch) MotionState_FirstStateOfWorkflowID(motionStateID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "first_state_of_workflow_id"}] = v
	return v
}

func (r *Fetch) MotionState_ID(motionStateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionState_MeetingID(motionStateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionState_MergeAmendmentIntoFinal(motionStateID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "merge_amendment_into_final"}] = v
	return v
}

func (r *Fetch) MotionState_MotionIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "motion_ids"}] = v
	return v
}

func (r *Fetch) MotionState_MotionRecommendationIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "motion_recommendation_ids"}] = v
	return v
}

func (r *Fetch) MotionState_Name(motionStateID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "name"}] = v
	return v
}

func (r *Fetch) MotionState_NextStateIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "next_state_ids"}] = v
	return v
}

func (r *Fetch) MotionState_PreviousStateIDs(motionStateID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "previous_state_ids"}] = v
	return v
}

func (r *Fetch) MotionState_RecommendationLabel(motionStateID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "recommendation_label"}] = v
	return v
}

func (r *Fetch) MotionState_Restrictions(motionStateID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "restrictions"}] = v
	return v
}

func (r *Fetch) MotionState_SetCreatedTimestamp(motionStateID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "set_created_timestamp"}] = v
	return v
}

func (r *Fetch) MotionState_SetNumber(motionStateID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "set_number"}] = v
	return v
}

func (r *Fetch) MotionState_ShowRecommendationExtensionField(motionStateID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "show_recommendation_extension_field"}] = v
	return v
}

func (r *Fetch) MotionState_ShowStateExtensionField(motionStateID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "show_state_extension_field"}] = v
	return v
}

func (r *Fetch) MotionState_Weight(motionStateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "weight"}] = v
	return v
}

func (r *Fetch) MotionState_WorkflowID(motionStateID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_state", ID: motionStateID, Field: "workflow_id"}] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_ID(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_MeetingID(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_MotionIDs(motionStatuteParagraphID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "motion_ids"}] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_SequentialNumber(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Text(motionStatuteParagraphID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "text"}] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Title(motionStatuteParagraphID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "title"}] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Weight(motionStatuteParagraphID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "weight"}] = v
	return v
}

func (r *Fetch) MotionSubmitter_ID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionSubmitter_MeetingID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionSubmitter_MotionID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "motion_id"}] = v
	return v
}

func (r *Fetch) MotionSubmitter_UserID(motionSubmitterID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "user_id"}] = v
	return v
}

func (r *Fetch) MotionSubmitter_Weight(motionSubmitterID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "weight"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "default_amendment_workflow_meeting_id"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultStatuteAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "default_statute_amendment_workflow_meeting_id"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "default_workflow_meeting_id"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_FirstStateID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "first_state_id"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_ID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "id"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_MeetingID(motionWorkflowID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_Name(motionWorkflowID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "name"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_SequentialNumber(motionWorkflowID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) MotionWorkflow_StateIDs(motionWorkflowID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "state_ids"}] = v
	return v
}

func (r *Fetch) Motion_AgendaItemID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "agenda_item_id"}] = v
	return v
}

func (r *Fetch) Motion_AllDerivedMotionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "all_derived_motion_ids"}] = v
	return v
}

func (r *Fetch) Motion_AllOriginIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "all_origin_ids"}] = v
	return v
}

func (r *Fetch) Motion_AmendmentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "amendment_ids"}] = v
	return v
}

func (r *Fetch) Motion_AmendmentParagraphTmpl(motionID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "amendment_paragraph_$"}] = v
	return v
}

func (r *Fetch) Motion_AmendmentParagraph(motionID int, replacement string) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: fmt.Sprintf("amendment_paragraph_$%s", replacement)}] = v
	return v
}

func (r *Fetch) Motion_AttachmentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "attachment_ids"}] = v
	return v
}

func (r *Fetch) Motion_BlockID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "block_id"}] = v
	return v
}

func (r *Fetch) Motion_CategoryID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "category_id"}] = v
	return v
}

func (r *Fetch) Motion_CategoryWeight(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "category_weight"}] = v
	return v
}

func (r *Fetch) Motion_ChangeRecommendationIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "change_recommendation_ids"}] = v
	return v
}

func (r *Fetch) Motion_CommentIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "comment_ids"}] = v
	return v
}

func (r *Fetch) Motion_Created(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "created"}] = v
	return v
}

func (r *Fetch) Motion_DerivedMotionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "derived_motion_ids"}] = v
	return v
}

func (r *Fetch) Motion_Forwarded(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "forwarded"}] = v
	return v
}

func (r *Fetch) Motion_ID(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "id"}] = v
	return v
}

func (r *Fetch) Motion_LastModified(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "last_modified"}] = v
	return v
}

func (r *Fetch) Motion_LeadMotionID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "lead_motion_id"}] = v
	return v
}

func (r *Fetch) Motion_ListOfSpeakersID(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "list_of_speakers_id"}] = v
	return v
}

func (r *Fetch) Motion_MeetingID(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Motion_ModifiedFinalVersion(motionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "modified_final_version"}] = v
	return v
}

func (r *Fetch) Motion_Number(motionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "number"}] = v
	return v
}

func (r *Fetch) Motion_NumberValue(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "number_value"}] = v
	return v
}

func (r *Fetch) Motion_OptionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "option_ids"}] = v
	return v
}

func (r *Fetch) Motion_OriginID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "origin_id"}] = v
	return v
}

func (r *Fetch) Motion_OriginMeetingID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "origin_meeting_id"}] = v
	return v
}

func (r *Fetch) Motion_PersonalNoteIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "personal_note_ids"}] = v
	return v
}

func (r *Fetch) Motion_PollIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "poll_ids"}] = v
	return v
}

func (r *Fetch) Motion_ProjectionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) Motion_Reason(motionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "reason"}] = v
	return v
}

func (r *Fetch) Motion_RecommendationExtension(motionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "recommendation_extension"}] = v
	return v
}

func (r *Fetch) Motion_RecommendationExtensionReferenceIDs(motionID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "recommendation_extension_reference_ids"}] = v
	return v
}

func (r *Fetch) Motion_RecommendationID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "recommendation_id"}] = v
	return v
}

func (r *Fetch) Motion_ReferencedInMotionRecommendationExtensionIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "referenced_in_motion_recommendation_extension_ids"}] = v
	return v
}

func (r *Fetch) Motion_SequentialNumber(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) Motion_SortChildIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "sort_child_ids"}] = v
	return v
}

func (r *Fetch) Motion_SortParentID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "sort_parent_id"}] = v
	return v
}

func (r *Fetch) Motion_SortWeight(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "sort_weight"}] = v
	return v
}

func (r *Fetch) Motion_StartLineNumber(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "start_line_number"}] = v
	return v
}

func (r *Fetch) Motion_StateExtension(motionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "state_extension"}] = v
	return v
}

func (r *Fetch) Motion_StateID(motionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "state_id"}] = v
	return v
}

func (r *Fetch) Motion_StatuteParagraphID(motionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "statute_paragraph_id"}] = v
	return v
}

func (r *Fetch) Motion_SubmitterIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "submitter_ids"}] = v
	return v
}

func (r *Fetch) Motion_SupporterIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "supporter_ids"}] = v
	return v
}

func (r *Fetch) Motion_TagIDs(motionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "tag_ids"}] = v
	return v
}

func (r *Fetch) Motion_Text(motionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "text"}] = v
	return v
}

func (r *Fetch) Motion_Title(motionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "motion", ID: motionID, Field: "title"}] = v
	return v
}

func (r *Fetch) Option_Abstain(optionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "abstain"}] = v
	return v
}

func (r *Fetch) Option_ContentObjectID(optionID int) *ValueMaybeString {
	v := &ValueMaybeString{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "content_object_id"}] = v
	return v
}

func (r *Fetch) Option_ID(optionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "id"}] = v
	return v
}

func (r *Fetch) Option_MeetingID(optionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Option_No(optionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "no"}] = v
	return v
}

func (r *Fetch) Option_PollID(optionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "poll_id"}] = v
	return v
}

func (r *Fetch) Option_Text(optionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "text"}] = v
	return v
}

func (r *Fetch) Option_UsedAsGlobalOptionInPollID(optionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "used_as_global_option_in_poll_id"}] = v
	return v
}

func (r *Fetch) Option_VoteIDs(optionID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "vote_ids"}] = v
	return v
}

func (r *Fetch) Option_Weight(optionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "weight"}] = v
	return v
}

func (r *Fetch) Option_Yes(optionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "option", ID: optionID, Field: "yes"}] = v
	return v
}

func (r *Fetch) OrganizationTag_Color(organizationTagID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization_tag", ID: organizationTagID, Field: "color"}] = v
	return v
}

func (r *Fetch) OrganizationTag_ID(organizationTagID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "organization_tag", ID: organizationTagID, Field: "id"}] = v
	return v
}

func (r *Fetch) OrganizationTag_Name(organizationTagID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization_tag", ID: organizationTagID, Field: "name"}] = v
	return v
}

func (r *Fetch) OrganizationTag_OrganizationID(organizationTagID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "organization_tag", ID: organizationTagID, Field: "organization_id"}] = v
	return v
}

func (r *Fetch) OrganizationTag_TaggedIDs(organizationTagID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization_tag", ID: organizationTagID, Field: "tagged_ids"}] = v
	return v
}

func (r *Fetch) Organization_ActiveMeetingIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "active_meeting_ids"}] = v
	return v
}

func (r *Fetch) Organization_ArchivedMeetingIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "archived_meeting_ids"}] = v
	return v
}

func (r *Fetch) Organization_CommitteeIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "committee_ids"}] = v
	return v
}

func (r *Fetch) Organization_Description(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "description"}] = v
	return v
}

func (r *Fetch) Organization_EnableChat(organizationID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "enable_chat"}] = v
	return v
}

func (r *Fetch) Organization_EnableElectronicVoting(organizationID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "enable_electronic_voting"}] = v
	return v
}

func (r *Fetch) Organization_ID(organizationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "id"}] = v
	return v
}

func (r *Fetch) Organization_LegalNotice(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "legal_notice"}] = v
	return v
}

func (r *Fetch) Organization_LimitOfMeetings(organizationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "limit_of_meetings"}] = v
	return v
}

func (r *Fetch) Organization_LimitOfUsers(organizationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "limit_of_users"}] = v
	return v
}

func (r *Fetch) Organization_LoginText(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "login_text"}] = v
	return v
}

func (r *Fetch) Organization_MediafileIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "mediafile_ids"}] = v
	return v
}

func (r *Fetch) Organization_Name(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "name"}] = v
	return v
}

func (r *Fetch) Organization_OrganizationTagIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "organization_tag_ids"}] = v
	return v
}

func (r *Fetch) Organization_PrivacyPolicy(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "privacy_policy"}] = v
	return v
}

func (r *Fetch) Organization_ResetPasswordVerboseErrors(organizationID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "reset_password_verbose_errors"}] = v
	return v
}

func (r *Fetch) Organization_TemplateMeetingIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "template_meeting_ids"}] = v
	return v
}

func (r *Fetch) Organization_ThemeID(organizationID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "theme_id"}] = v
	return v
}

func (r *Fetch) Organization_ThemeIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "theme_ids"}] = v
	return v
}

func (r *Fetch) Organization_Url(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "url"}] = v
	return v
}

func (r *Fetch) Organization_UserIDs(organizationID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "user_ids"}] = v
	return v
}

func (r *Fetch) Organization_UsersEmailBody(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "users_email_body"}] = v
	return v
}

func (r *Fetch) Organization_UsersEmailReplyto(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "users_email_replyto"}] = v
	return v
}

func (r *Fetch) Organization_UsersEmailSender(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "users_email_sender"}] = v
	return v
}

func (r *Fetch) Organization_UsersEmailSubject(organizationID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "organization", ID: organizationID, Field: "users_email_subject"}] = v
	return v
}

func (r *Fetch) PersonalNote_ContentObjectID(personalNoteID int) *ValueMaybeString {
	v := &ValueMaybeString{fetch: r}
	r.requested[datastore.Key{Collection: "personal_note", ID: personalNoteID, Field: "content_object_id"}] = v
	return v
}

func (r *Fetch) PersonalNote_ID(personalNoteID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "personal_note", ID: personalNoteID, Field: "id"}] = v
	return v
}

func (r *Fetch) PersonalNote_MeetingID(personalNoteID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "personal_note", ID: personalNoteID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) PersonalNote_Note(personalNoteID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "personal_note", ID: personalNoteID, Field: "note"}] = v
	return v
}

func (r *Fetch) PersonalNote_Star(personalNoteID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "personal_note", ID: personalNoteID, Field: "star"}] = v
	return v
}

func (r *Fetch) PersonalNote_UserID(personalNoteID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "personal_note", ID: personalNoteID, Field: "user_id"}] = v
	return v
}

func (r *Fetch) Poll_Backend(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "backend"}] = v
	return v
}

func (r *Fetch) Poll_ContentObjectID(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "content_object_id"}] = v
	return v
}

func (r *Fetch) Poll_CryptKey(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "crypt_key"}] = v
	return v
}

func (r *Fetch) Poll_CryptSignature(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "crypt_signature"}] = v
	return v
}

func (r *Fetch) Poll_Description(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "description"}] = v
	return v
}

func (r *Fetch) Poll_EntitledGroupIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "entitled_group_ids"}] = v
	return v
}

func (r *Fetch) Poll_EntitledUsersAtStop(pollID int) *ValueJSON {
	v := &ValueJSON{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "entitled_users_at_stop"}] = v
	return v
}

func (r *Fetch) Poll_GlobalAbstain(pollID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "global_abstain"}] = v
	return v
}

func (r *Fetch) Poll_GlobalNo(pollID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "global_no"}] = v
	return v
}

func (r *Fetch) Poll_GlobalOptionID(pollID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "global_option_id"}] = v
	return v
}

func (r *Fetch) Poll_GlobalYes(pollID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "global_yes"}] = v
	return v
}

func (r *Fetch) Poll_ID(pollID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "id"}] = v
	return v
}

func (r *Fetch) Poll_IsPseudoanonymized(pollID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "is_pseudoanonymized"}] = v
	return v
}

func (r *Fetch) Poll_MaxVotesAmount(pollID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "max_votes_amount"}] = v
	return v
}

func (r *Fetch) Poll_MaxVotesPerOption(pollID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "max_votes_per_option"}] = v
	return v
}

func (r *Fetch) Poll_MeetingID(pollID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Poll_MinVotesAmount(pollID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "min_votes_amount"}] = v
	return v
}

func (r *Fetch) Poll_OnehundredPercentBase(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "onehundred_percent_base"}] = v
	return v
}

func (r *Fetch) Poll_OptionIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "option_ids"}] = v
	return v
}

func (r *Fetch) Poll_Pollmethod(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "pollmethod"}] = v
	return v
}

func (r *Fetch) Poll_ProjectionIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) Poll_SequentialNumber(pollID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) Poll_State(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "state"}] = v
	return v
}

func (r *Fetch) Poll_Title(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "title"}] = v
	return v
}

func (r *Fetch) Poll_Type(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "type"}] = v
	return v
}

func (r *Fetch) Poll_VoteCount(pollID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "vote_count"}] = v
	return v
}

func (r *Fetch) Poll_VotedIDs(pollID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "voted_ids"}] = v
	return v
}

func (r *Fetch) Poll_VotesRaw(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "votes_raw"}] = v
	return v
}

func (r *Fetch) Poll_VotesSignature(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "votes_signature"}] = v
	return v
}

func (r *Fetch) Poll_Votescast(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "votescast"}] = v
	return v
}

func (r *Fetch) Poll_Votesinvalid(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "votesinvalid"}] = v
	return v
}

func (r *Fetch) Poll_Votesvalid(pollID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "poll", ID: pollID, Field: "votesvalid"}] = v
	return v
}

func (r *Fetch) Projection_Content(projectionID int) *ValueJSON {
	v := &ValueJSON{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "content"}] = v
	return v
}

func (r *Fetch) Projection_ContentObjectID(projectionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "content_object_id"}] = v
	return v
}

func (r *Fetch) Projection_CurrentProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "current_projector_id"}] = v
	return v
}

func (r *Fetch) Projection_HistoryProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "history_projector_id"}] = v
	return v
}

func (r *Fetch) Projection_ID(projectionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "id"}] = v
	return v
}

func (r *Fetch) Projection_MeetingID(projectionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Projection_Options(projectionID int) *ValueJSON {
	v := &ValueJSON{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "options"}] = v
	return v
}

func (r *Fetch) Projection_PreviewProjectorID(projectionID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "preview_projector_id"}] = v
	return v
}

func (r *Fetch) Projection_Stable(projectionID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "stable"}] = v
	return v
}

func (r *Fetch) Projection_Type(projectionID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "type"}] = v
	return v
}

func (r *Fetch) Projection_Weight(projectionID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projection", ID: projectionID, Field: "weight"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_CountdownTime(projectorCountdownID int) *ValueFloat {
	v := &ValueFloat{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "countdown_time"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_DefaultTime(projectorCountdownID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "default_time"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Description(projectorCountdownID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "description"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_ID(projectorCountdownID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "id"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_MeetingID(projectorCountdownID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_ProjectionIDs(projectorCountdownID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Running(projectorCountdownID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "running"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Title(projectorCountdownID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "title"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "used_as_list_of_speakers_countdown_meeting_id"}] = v
	return v
}

func (r *Fetch) ProjectorCountdown_UsedAsPollCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "used_as_poll_countdown_meeting_id"}] = v
	return v
}

func (r *Fetch) ProjectorMessage_ID(projectorMessageID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector_message", ID: projectorMessageID, Field: "id"}] = v
	return v
}

func (r *Fetch) ProjectorMessage_MeetingID(projectorMessageID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector_message", ID: projectorMessageID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) ProjectorMessage_Message(projectorMessageID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector_message", ID: projectorMessageID, Field: "message"}] = v
	return v
}

func (r *Fetch) ProjectorMessage_ProjectionIDs(projectorMessageID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "projector_message", ID: projectorMessageID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) Projector_AspectRatioDenominator(projectorID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "aspect_ratio_denominator"}] = v
	return v
}

func (r *Fetch) Projector_AspectRatioNumerator(projectorID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "aspect_ratio_numerator"}] = v
	return v
}

func (r *Fetch) Projector_BackgroundColor(projectorID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "background_color"}] = v
	return v
}

func (r *Fetch) Projector_ChyronBackgroundColor(projectorID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "chyron_background_color"}] = v
	return v
}

func (r *Fetch) Projector_ChyronFontColor(projectorID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "chyron_font_color"}] = v
	return v
}

func (r *Fetch) Projector_Color(projectorID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "color"}] = v
	return v
}

func (r *Fetch) Projector_CurrentProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "current_projection_ids"}] = v
	return v
}

func (r *Fetch) Projector_HeaderBackgroundColor(projectorID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "header_background_color"}] = v
	return v
}

func (r *Fetch) Projector_HeaderFontColor(projectorID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "header_font_color"}] = v
	return v
}

func (r *Fetch) Projector_HeaderH1Color(projectorID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "header_h1_color"}] = v
	return v
}

func (r *Fetch) Projector_HistoryProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "history_projection_ids"}] = v
	return v
}

func (r *Fetch) Projector_ID(projectorID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "id"}] = v
	return v
}

func (r *Fetch) Projector_MeetingID(projectorID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Projector_Name(projectorID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "name"}] = v
	return v
}

func (r *Fetch) Projector_PreviewProjectionIDs(projectorID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "preview_projection_ids"}] = v
	return v
}

func (r *Fetch) Projector_Scale(projectorID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "scale"}] = v
	return v
}

func (r *Fetch) Projector_Scroll(projectorID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "scroll"}] = v
	return v
}

func (r *Fetch) Projector_SequentialNumber(projectorID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) Projector_ShowClock(projectorID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "show_clock"}] = v
	return v
}

func (r *Fetch) Projector_ShowHeaderFooter(projectorID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "show_header_footer"}] = v
	return v
}

func (r *Fetch) Projector_ShowLogo(projectorID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "show_logo"}] = v
	return v
}

func (r *Fetch) Projector_ShowTitle(projectorID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "show_title"}] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultInMeetingIDTmpl(projectorID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_$_in_meeting_id"}] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultInMeetingID(projectorID int, replacement string) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: fmt.Sprintf("used_as_default_$%s_in_meeting_id", replacement)}] = v
	return v
}

func (r *Fetch) Projector_UsedAsReferenceProjectorMeetingID(projectorID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "used_as_reference_projector_meeting_id"}] = v
	return v
}

func (r *Fetch) Projector_Width(projectorID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "projector", ID: projectorID, Field: "width"}] = v
	return v
}

func (r *Fetch) Speaker_BeginTime(speakerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "begin_time"}] = v
	return v
}

func (r *Fetch) Speaker_EndTime(speakerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "end_time"}] = v
	return v
}

func (r *Fetch) Speaker_ID(speakerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "id"}] = v
	return v
}

func (r *Fetch) Speaker_ListOfSpeakersID(speakerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "list_of_speakers_id"}] = v
	return v
}

func (r *Fetch) Speaker_MeetingID(speakerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Speaker_Note(speakerID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "note"}] = v
	return v
}

func (r *Fetch) Speaker_PointOfOrder(speakerID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "point_of_order"}] = v
	return v
}

func (r *Fetch) Speaker_SpeechState(speakerID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "speech_state"}] = v
	return v
}

func (r *Fetch) Speaker_UserID(speakerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "user_id"}] = v
	return v
}

func (r *Fetch) Speaker_Weight(speakerID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "speaker", ID: speakerID, Field: "weight"}] = v
	return v
}

func (r *Fetch) Tag_ID(tagID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "tag", ID: tagID, Field: "id"}] = v
	return v
}

func (r *Fetch) Tag_MeetingID(tagID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "tag", ID: tagID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Tag_Name(tagID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "tag", ID: tagID, Field: "name"}] = v
	return v
}

func (r *Fetch) Tag_TaggedIDs(tagID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "tag", ID: tagID, Field: "tagged_ids"}] = v
	return v
}

func (r *Fetch) Theme_Accent100(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_100"}] = v
	return v
}

func (r *Fetch) Theme_Accent200(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_200"}] = v
	return v
}

func (r *Fetch) Theme_Accent300(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_300"}] = v
	return v
}

func (r *Fetch) Theme_Accent400(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_400"}] = v
	return v
}

func (r *Fetch) Theme_Accent50(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_50"}] = v
	return v
}

func (r *Fetch) Theme_Accent500(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_500"}] = v
	return v
}

func (r *Fetch) Theme_Accent600(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_600"}] = v
	return v
}

func (r *Fetch) Theme_Accent700(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_700"}] = v
	return v
}

func (r *Fetch) Theme_Accent800(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_800"}] = v
	return v
}

func (r *Fetch) Theme_Accent900(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_900"}] = v
	return v
}

func (r *Fetch) Theme_AccentA100(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_a100"}] = v
	return v
}

func (r *Fetch) Theme_AccentA200(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_a200"}] = v
	return v
}

func (r *Fetch) Theme_AccentA400(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_a400"}] = v
	return v
}

func (r *Fetch) Theme_AccentA700(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "accent_a700"}] = v
	return v
}

func (r *Fetch) Theme_ID(themeID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "id"}] = v
	return v
}

func (r *Fetch) Theme_Name(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "name"}] = v
	return v
}

func (r *Fetch) Theme_OrganizationID(themeID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "organization_id"}] = v
	return v
}

func (r *Fetch) Theme_Primary100(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_100"}] = v
	return v
}

func (r *Fetch) Theme_Primary200(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_200"}] = v
	return v
}

func (r *Fetch) Theme_Primary300(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_300"}] = v
	return v
}

func (r *Fetch) Theme_Primary400(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_400"}] = v
	return v
}

func (r *Fetch) Theme_Primary50(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_50"}] = v
	return v
}

func (r *Fetch) Theme_Primary500(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_500"}] = v
	return v
}

func (r *Fetch) Theme_Primary600(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_600"}] = v
	return v
}

func (r *Fetch) Theme_Primary700(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_700"}] = v
	return v
}

func (r *Fetch) Theme_Primary800(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_800"}] = v
	return v
}

func (r *Fetch) Theme_Primary900(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_900"}] = v
	return v
}

func (r *Fetch) Theme_PrimaryA100(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_a100"}] = v
	return v
}

func (r *Fetch) Theme_PrimaryA200(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_a200"}] = v
	return v
}

func (r *Fetch) Theme_PrimaryA400(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_a400"}] = v
	return v
}

func (r *Fetch) Theme_PrimaryA700(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "primary_a700"}] = v
	return v
}

func (r *Fetch) Theme_ThemeForOrganizationID(themeID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "theme_for_organization_id"}] = v
	return v
}

func (r *Fetch) Theme_Warn100(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_100"}] = v
	return v
}

func (r *Fetch) Theme_Warn200(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_200"}] = v
	return v
}

func (r *Fetch) Theme_Warn300(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_300"}] = v
	return v
}

func (r *Fetch) Theme_Warn400(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_400"}] = v
	return v
}

func (r *Fetch) Theme_Warn50(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_50"}] = v
	return v
}

func (r *Fetch) Theme_Warn500(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_500"}] = v
	return v
}

func (r *Fetch) Theme_Warn600(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_600"}] = v
	return v
}

func (r *Fetch) Theme_Warn700(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_700"}] = v
	return v
}

func (r *Fetch) Theme_Warn800(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_800"}] = v
	return v
}

func (r *Fetch) Theme_Warn900(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_900"}] = v
	return v
}

func (r *Fetch) Theme_WarnA100(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_a100"}] = v
	return v
}

func (r *Fetch) Theme_WarnA200(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_a200"}] = v
	return v
}

func (r *Fetch) Theme_WarnA400(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_a400"}] = v
	return v
}

func (r *Fetch) Theme_WarnA700(themeID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "theme", ID: themeID, Field: "warn_a700"}] = v
	return v
}

func (r *Fetch) Topic_AgendaItemID(topicID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "agenda_item_id"}] = v
	return v
}

func (r *Fetch) Topic_AttachmentIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "attachment_ids"}] = v
	return v
}

func (r *Fetch) Topic_ID(topicID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "id"}] = v
	return v
}

func (r *Fetch) Topic_ListOfSpeakersID(topicID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "list_of_speakers_id"}] = v
	return v
}

func (r *Fetch) Topic_MeetingID(topicID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Topic_PollIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "poll_ids"}] = v
	return v
}

func (r *Fetch) Topic_ProjectionIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "projection_ids"}] = v
	return v
}

func (r *Fetch) Topic_SequentialNumber(topicID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "sequential_number"}] = v
	return v
}

func (r *Fetch) Topic_TagIDs(topicID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "tag_ids"}] = v
	return v
}

func (r *Fetch) Topic_Text(topicID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "text"}] = v
	return v
}

func (r *Fetch) Topic_Title(topicID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "topic", ID: topicID, Field: "title"}] = v
	return v
}

func (r *Fetch) User_AboutMeTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "about_me_$"}] = v
	return v
}

func (r *Fetch) User_AboutMe(userID int, meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("about_me_$%d", meetingID)}] = v
	return v
}

func (r *Fetch) User_AssignmentCandidateIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "assignment_candidate_$_ids"}] = v
	return v
}

func (r *Fetch) User_AssignmentCandidateIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("assignment_candidate_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_CanChangeOwnPassword(userID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "can_change_own_password"}] = v
	return v
}

func (r *Fetch) User_ChatMessageIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "chat_message_$_ids"}] = v
	return v
}

func (r *Fetch) User_ChatMessageIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("chat_message_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_CommentTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "comment_$"}] = v
	return v
}

func (r *Fetch) User_Comment(userID int, meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("comment_$%d", meetingID)}] = v
	return v
}

func (r *Fetch) User_CommitteeIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "committee_ids"}] = v
	return v
}

func (r *Fetch) User_CommitteeManagementLevelTmpl(userID int) *ValueStringSlice {
	v := &ValueStringSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "committee_$_management_level"}] = v
	return v
}

func (r *Fetch) User_CommitteeManagementLevel(userID int, replacement string) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("committee_$%s_management_level", replacement)}] = v
	return v
}

func (r *Fetch) User_DefaultNumber(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "default_number"}] = v
	return v
}

func (r *Fetch) User_DefaultPassword(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "default_password"}] = v
	return v
}

func (r *Fetch) User_DefaultStructureLevel(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "default_structure_level"}] = v
	return v
}

func (r *Fetch) User_DefaultVoteWeight(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "default_vote_weight"}] = v
	return v
}

func (r *Fetch) User_Email(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "email"}] = v
	return v
}

func (r *Fetch) User_FirstName(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "first_name"}] = v
	return v
}

func (r *Fetch) User_ForwardingCommitteeIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "forwarding_committee_ids"}] = v
	return v
}

func (r *Fetch) User_Gender(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "gender"}] = v
	return v
}

func (r *Fetch) User_GroupIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "group_$_ids"}] = v
	return v
}

func (r *Fetch) User_GroupIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("group_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_ID(userID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "id"}] = v
	return v
}

func (r *Fetch) User_IsActive(userID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "is_active"}] = v
	return v
}

func (r *Fetch) User_IsDemoUser(userID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "is_demo_user"}] = v
	return v
}

func (r *Fetch) User_IsPhysicalPerson(userID int) *ValueBool {
	v := &ValueBool{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "is_physical_person"}] = v
	return v
}

func (r *Fetch) User_IsPresentInMeetingIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "is_present_in_meeting_ids"}] = v
	return v
}

func (r *Fetch) User_LastEmailSend(userID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "last_email_send"}] = v
	return v
}

func (r *Fetch) User_LastName(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "last_name"}] = v
	return v
}

func (r *Fetch) User_MeetingIDs(userID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "meeting_ids"}] = v
	return v
}

func (r *Fetch) User_NumberTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "number_$"}] = v
	return v
}

func (r *Fetch) User_Number(userID int, meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("number_$%d", meetingID)}] = v
	return v
}

func (r *Fetch) User_OptionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "option_$_ids"}] = v
	return v
}

func (r *Fetch) User_OptionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("option_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_OrganizationID(userID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "organization_id"}] = v
	return v
}

func (r *Fetch) User_OrganizationManagementLevel(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "organization_management_level"}] = v
	return v
}

func (r *Fetch) User_Password(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "password"}] = v
	return v
}

func (r *Fetch) User_PersonalNoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "personal_note_$_ids"}] = v
	return v
}

func (r *Fetch) User_PersonalNoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("personal_note_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_PollVotedIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "poll_voted_$_ids"}] = v
	return v
}

func (r *Fetch) User_PollVotedIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("poll_voted_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_ProjectionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "projection_$_ids"}] = v
	return v
}

func (r *Fetch) User_ProjectionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("projection_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_Pronoun(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "pronoun"}] = v
	return v
}

func (r *Fetch) User_SpeakerIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "speaker_$_ids"}] = v
	return v
}

func (r *Fetch) User_SpeakerIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("speaker_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_StructureLevelTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "structure_level_$"}] = v
	return v
}

func (r *Fetch) User_StructureLevel(userID int, meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("structure_level_$%d", meetingID)}] = v
	return v
}

func (r *Fetch) User_SubmittedMotionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "submitted_motion_$_ids"}] = v
	return v
}

func (r *Fetch) User_SubmittedMotionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("submitted_motion_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_SupportedMotionIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "supported_motion_$_ids"}] = v
	return v
}

func (r *Fetch) User_SupportedMotionIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("supported_motion_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_Title(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "title"}] = v
	return v
}

func (r *Fetch) User_Username(userID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "username"}] = v
	return v
}

func (r *Fetch) User_VoteDelegatedToIDTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "vote_delegated_$_to_id"}] = v
	return v
}

func (r *Fetch) User_VoteDelegatedToID(userID int, meetingID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("vote_delegated_$%d_to_id", meetingID)}] = v
	return v
}

func (r *Fetch) User_VoteDelegatedVoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "vote_delegated_vote_$_ids"}] = v
	return v
}

func (r *Fetch) User_VoteDelegatedVoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("vote_delegated_vote_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_VoteDelegationsFromIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "vote_delegations_$_from_ids"}] = v
	return v
}

func (r *Fetch) User_VoteDelegationsFromIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("vote_delegations_$%d_from_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_VoteIDsTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "vote_$_ids"}] = v
	return v
}

func (r *Fetch) User_VoteIDs(userID int, meetingID int) *ValueIntSlice {
	v := &ValueIntSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("vote_$%d_ids", meetingID)}] = v
	return v
}

func (r *Fetch) User_VoteWeightTmpl(userID int) *ValueIDSlice {
	v := &ValueIDSlice{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: "vote_weight_$"}] = v
	return v
}

func (r *Fetch) User_VoteWeight(userID int, meetingID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "user", ID: userID, Field: fmt.Sprintf("vote_weight_$%d", meetingID)}] = v
	return v
}

func (r *Fetch) Vote_DelegatedUserID(voteID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "vote", ID: voteID, Field: "delegated_user_id"}] = v
	return v
}

func (r *Fetch) Vote_ID(voteID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "vote", ID: voteID, Field: "id"}] = v
	return v
}

func (r *Fetch) Vote_MeetingID(voteID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "vote", ID: voteID, Field: "meeting_id"}] = v
	return v
}

func (r *Fetch) Vote_OptionID(voteID int) *ValueInt {
	v := &ValueInt{fetch: r}
	r.requested[datastore.Key{Collection: "vote", ID: voteID, Field: "option_id"}] = v
	return v
}

func (r *Fetch) Vote_UserID(voteID int) *ValueMaybeInt {
	v := &ValueMaybeInt{fetch: r}
	r.requested[datastore.Key{Collection: "vote", ID: voteID, Field: "user_id"}] = v
	return v
}

func (r *Fetch) Vote_UserToken(voteID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "vote", ID: voteID, Field: "user_token"}] = v
	return v
}

func (r *Fetch) Vote_Value(voteID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "vote", ID: voteID, Field: "value"}] = v
	return v
}

func (r *Fetch) Vote_Weight(voteID int) *ValueString {
	v := &ValueString{fetch: r}
	r.requested[datastore.Key{Collection: "vote", ID: voteID, Field: "weight"}] = v
	return v
}
