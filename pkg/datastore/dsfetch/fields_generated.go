// Code generated from models.yml DO NOT EDIT.
package dsfetch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/fastjson"
)

// ValueBool is a value from the datastore.
type ValueBool struct {
	err error

	key      dskey.Key
	required bool

	lazies []*bool

	fetch *Fetch
}

// Value returns the value.
func (v *ValueBool) Value(ctx context.Context) (bool, error) {
	var zero bool
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value bool
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueBool) Lazy(value *bool) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueBool) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueBool) execute(p []byte) error {
	var value bool
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		if err := json.Unmarshal(p, &value); err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

// ValueFloat is a value from the datastore.
type ValueFloat struct {
	err error

	key      dskey.Key
	required bool

	lazies []*float32

	fetch *Fetch
}

// Value returns the value.
func (v *ValueFloat) Value(ctx context.Context) (float32, error) {
	var zero float32
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value float32
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueFloat) Lazy(value *float32) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueFloat) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueFloat) execute(p []byte) error {
	var value float32
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		if err := json.Unmarshal(p, &value); err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

// ValueInt is a value from the datastore.
type ValueInt struct {
	err error

	key      dskey.Key
	required bool

	lazies []*int

	fetch *Fetch
}

// Value returns the value.
func (v *ValueInt) Value(ctx context.Context) (int, error) {
	var zero int
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value int
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueInt) Lazy(value *int) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueInt) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueInt) execute(p []byte) error {
	var value int
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		r, err := fastjson.DecodeInt(p)
		if err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
		value = r
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

// ValueIntSlice is a value from the datastore.
type ValueIntSlice struct {
	err error

	key      dskey.Key
	required bool

	lazies []*[]int

	fetch *Fetch
}

// Value returns the value.
func (v *ValueIntSlice) Value(ctx context.Context) ([]int, error) {
	var zero []int
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value []int
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueIntSlice) Lazy(value *[]int) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueIntSlice) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueIntSlice) execute(p []byte) error {
	var value []int
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		r, err := fastjson.DecodeIntList(p)
		if err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
		value = r
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

// ValueJSON is a value from the datastore.
type ValueJSON struct {
	err error

	key      dskey.Key
	required bool

	lazies []*json.RawMessage

	fetch *Fetch
}

// Value returns the value.
func (v *ValueJSON) Value(ctx context.Context) (json.RawMessage, error) {
	var zero json.RawMessage
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value json.RawMessage
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueJSON) Lazy(value *json.RawMessage) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueJSON) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueJSON) execute(p []byte) error {
	var value json.RawMessage
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		if err := json.Unmarshal(p, &value); err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

// ValueMaybeInt is a value from the datastore.
type ValueMaybeInt struct {
	err error

	key      dskey.Key
	required bool

	lazies []*Maybe[int]

	fetch *Fetch
}

// Value returns the value.
func (v *ValueMaybeInt) Value(ctx context.Context) (Maybe[int], error) {
	var zero Maybe[int]
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value Maybe[int]
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueMaybeInt) Lazy(value *Maybe[int]) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueMaybeInt) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueMaybeInt) execute(p []byte) error {
	var value Maybe[int]
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		if err := json.Unmarshal(p, &value); err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

// ValueMaybeString is a value from the datastore.
type ValueMaybeString struct {
	err error

	key      dskey.Key
	required bool

	lazies []*Maybe[string]

	fetch *Fetch
}

// Value returns the value.
func (v *ValueMaybeString) Value(ctx context.Context) (Maybe[string], error) {
	var zero Maybe[string]
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value Maybe[string]
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueMaybeString) Lazy(value *Maybe[string]) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueMaybeString) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueMaybeString) execute(p []byte) error {
	var value Maybe[string]
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		if err := json.Unmarshal(p, &value); err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

// ValueString is a value from the datastore.
type ValueString struct {
	err error

	key      dskey.Key
	required bool

	lazies []*string

	fetch *Fetch
}

// Value returns the value.
func (v *ValueString) Value(ctx context.Context) (string, error) {
	var zero string
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value string
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueString) Lazy(value *string) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueString) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueString) execute(p []byte) error {
	var value string
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		if err := json.Unmarshal(p, &value); err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

// ValueStringSlice is a value from the datastore.
type ValueStringSlice struct {
	err error

	key      dskey.Key
	required bool

	lazies []*[]string

	fetch *Fetch
}

// Value returns the value.
func (v *ValueStringSlice) Value(ctx context.Context) ([]string, error) {
	var zero []string
	if err := v.err; err != nil {
		return zero, v.err
	}

	rawValue, err := v.fetch.getOneKey(ctx, v.key)
	if err != nil {
		return zero, err
	}

	var value []string
	v.Lazy(&value)

	if err := v.execute(rawValue); err != nil {
		return zero, err
	}

	return value, nil
}

// Lazy sets a value as soon as it es executed.
//
// Make sure to call request.Execute() before using the value.
func (v *ValueStringSlice) Lazy(value *[]string) {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
	v.lazies = append(v.lazies, value)
}

// Preload fetches the value but does nothing with it.
//
// This makes sure, that the value is in the cache.
//
// Make sure to call fetch.Execute().
func (v *ValueStringSlice) Preload() {
	v.fetch.requested[v.key] = append(v.fetch.requested[v.key], v)
}

// execute will be called from request.
func (v *ValueStringSlice) execute(p []byte) error {
	var value []string
	if p == nil {
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
	} else {
		if err := json.Unmarshal(p, &value); err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
	}

	for i := 0; i < len(v.lazies); i++ {
		*v.lazies[i] = value
	}

	return nil
}

func (r *Fetch) ActionWorker_Created(actionWorkerID int) *ValueInt {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "created")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ActionWorker_ID(actionWorkerID int) *ValueInt {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ActionWorker_Name(actionWorkerID int) *ValueString {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ActionWorker_Result(actionWorkerID int) *ValueJSON {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "result")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) ActionWorker_State(actionWorkerID int) *ValueString {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "state")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ActionWorker_Timestamp(actionWorkerID int) *ValueInt {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "timestamp")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) AgendaItem_ChildIDs(agendaItemID int) *ValueIntSlice {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "child_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_Closed(agendaItemID int) *ValueBool {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "closed")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_Comment(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "comment")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_ContentObjectID(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "content_object_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) AgendaItem_Duration(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "duration")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_ID(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_IsHidden(agendaItemID int) *ValueBool {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "is_hidden")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_IsInternal(agendaItemID int) *ValueBool {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "is_internal")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_ItemNumber(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "item_number")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_Level(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "level")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_MeetingID(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) AgendaItem_ModeratorNotes(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "moderator_notes")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_ParentID(agendaItemID int) *ValueMaybeInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "parent_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_ProjectionIDs(agendaItemID int) *ValueIntSlice {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_TagIDs(agendaItemID int) *ValueIntSlice {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "tag_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_Type(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) AgendaItem_Weight(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) AssignmentCandidate_AssignmentID(assignmentCandidateID int) *ValueInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "assignment_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) AssignmentCandidate_ID(assignmentCandidateID int) *ValueInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) AssignmentCandidate_MeetingID(assignmentCandidateID int) *ValueInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) AssignmentCandidate_MeetingUserID(assignmentCandidateID int) *ValueMaybeInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "meeting_user_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) AssignmentCandidate_Weight(assignmentCandidateID int) *ValueInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Assignment_AgendaItemID(assignmentID int) *ValueMaybeInt {
	key, err := dskey.FromParts("assignment", assignmentID, "agenda_item_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Assignment_AttachmentIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "attachment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Assignment_CandidateIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "candidate_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Assignment_DefaultPollDescription(assignmentID int) *ValueString {
	key, err := dskey.FromParts("assignment", assignmentID, "default_poll_description")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Assignment_Description(assignmentID int) *ValueString {
	key, err := dskey.FromParts("assignment", assignmentID, "description")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Assignment_ID(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Assignment_ListOfSpeakersID(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "list_of_speakers_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Assignment_MeetingID(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Assignment_NumberPollCandidates(assignmentID int) *ValueBool {
	key, err := dskey.FromParts("assignment", assignmentID, "number_poll_candidates")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Assignment_OpenPosts(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "open_posts")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Assignment_Phase(assignmentID int) *ValueString {
	key, err := dskey.FromParts("assignment", assignmentID, "phase")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Assignment_PollIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "poll_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Assignment_ProjectionIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Assignment_SequentialNumber(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Assignment_TagIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "tag_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Assignment_Title(assignmentID int) *ValueString {
	key, err := dskey.FromParts("assignment", assignmentID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatGroup_ChatMessageIDs(chatGroupID int) *ValueIntSlice {
	key, err := dskey.FromParts("chat_group", chatGroupID, "chat_message_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) ChatGroup_ID(chatGroupID int) *ValueInt {
	key, err := dskey.FromParts("chat_group", chatGroupID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ChatGroup_MeetingID(chatGroupID int) *ValueInt {
	key, err := dskey.FromParts("chat_group", chatGroupID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatGroup_Name(chatGroupID int) *ValueString {
	key, err := dskey.FromParts("chat_group", chatGroupID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatGroup_ReadGroupIDs(chatGroupID int) *ValueIntSlice {
	key, err := dskey.FromParts("chat_group", chatGroupID, "read_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) ChatGroup_Weight(chatGroupID int) *ValueInt {
	key, err := dskey.FromParts("chat_group", chatGroupID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ChatGroup_WriteGroupIDs(chatGroupID int) *ValueIntSlice {
	key, err := dskey.FromParts("chat_group", chatGroupID, "write_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) ChatMessage_ChatGroupID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "chat_group_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatMessage_Content(chatMessageID int) *ValueString {
	key, err := dskey.FromParts("chat_message", chatMessageID, "content")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatMessage_Created(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "created")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatMessage_ID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ChatMessage_MeetingID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatMessage_MeetingUserID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "meeting_user_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Committee_DefaultMeetingID(committeeID int) *ValueMaybeInt {
	key, err := dskey.FromParts("committee", committeeID, "default_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Committee_Description(committeeID int) *ValueString {
	key, err := dskey.FromParts("committee", committeeID, "description")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Committee_ExternalID(committeeID int) *ValueString {
	key, err := dskey.FromParts("committee", committeeID, "external_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Committee_ForwardToCommitteeIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "forward_to_committee_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Committee_ForwardingUserID(committeeID int) *ValueMaybeInt {
	key, err := dskey.FromParts("committee", committeeID, "forwarding_user_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Committee_ID(committeeID int) *ValueInt {
	key, err := dskey.FromParts("committee", committeeID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Committee_ManagerIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "manager_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Committee_MeetingIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "meeting_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Committee_Name(committeeID int) *ValueString {
	key, err := dskey.FromParts("committee", committeeID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Committee_OrganizationID(committeeID int) *ValueInt {
	key, err := dskey.FromParts("committee", committeeID, "organization_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Committee_OrganizationTagIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "organization_tag_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Committee_ReceiveForwardingsFromCommitteeIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "receive_forwardings_from_committee_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Committee_UserIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_AdminGroupForMeetingID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "admin_group_for_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Group_DefaultGroupForMeetingID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "default_group_for_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Group_ExternalID(groupID int) *ValueString {
	key, err := dskey.FromParts("group", groupID, "external_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Group_ID(groupID int) *ValueInt {
	key, err := dskey.FromParts("group", groupID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Group_MediafileAccessGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "mediafile_access_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_MediafileInheritedAccessGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "mediafile_inherited_access_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_MeetingID(groupID int) *ValueInt {
	key, err := dskey.FromParts("group", groupID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Group_MeetingUserIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "meeting_user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_Name(groupID int) *ValueString {
	key, err := dskey.FromParts("group", groupID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Group_Permissions(groupID int) *ValueStringSlice {
	key, err := dskey.FromParts("group", groupID, "permissions")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) Group_PollIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "poll_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_ReadChatGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "read_chat_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_ReadCommentSectionIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "read_comment_section_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_UsedAsAssignmentPollDefaultID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "used_as_assignment_poll_default_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Group_UsedAsMotionPollDefaultID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "used_as_motion_poll_default_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Group_UsedAsPollDefaultID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "used_as_poll_default_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Group_UsedAsTopicPollDefaultID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "used_as_topic_poll_default_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Group_Weight(groupID int) *ValueInt {
	key, err := dskey.FromParts("group", groupID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Group_WriteChatGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "write_chat_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_WriteCommentSectionIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "write_comment_section_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) ImportPreview_Created(importPreviewID int) *ValueInt {
	key, err := dskey.FromParts("import_preview", importPreviewID, "created")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ImportPreview_ID(importPreviewID int) *ValueInt {
	key, err := dskey.FromParts("import_preview", importPreviewID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ImportPreview_Name(importPreviewID int) *ValueString {
	key, err := dskey.FromParts("import_preview", importPreviewID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ImportPreview_Result(importPreviewID int) *ValueJSON {
	key, err := dskey.FromParts("import_preview", importPreviewID, "result")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) ImportPreview_State(importPreviewID int) *ValueString {
	key, err := dskey.FromParts("import_preview", importPreviewID, "state")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ListOfSpeakers_Closed(listOfSpeakersID int) *ValueBool {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "closed")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) ListOfSpeakers_ContentObjectID(listOfSpeakersID int) *ValueString {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "content_object_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ListOfSpeakers_ID(listOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ListOfSpeakers_MeetingID(listOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ListOfSpeakers_ProjectionIDs(listOfSpeakersID int) *ValueIntSlice {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) ListOfSpeakers_SequentialNumber(listOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ListOfSpeakers_SpeakerIDs(listOfSpeakersID int) *ValueIntSlice {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "speaker_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) ListOfSpeakers_StructureLevelListOfSpeakersIDs(listOfSpeakersID int) *ValueIntSlice {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "structure_level_list_of_speakers_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Mediafile_AccessGroupIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "access_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Mediafile_AttachmentIDs(mediafileID int) *ValueStringSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "attachment_ids")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) Mediafile_ChildIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "child_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Mediafile_CreateTimestamp(mediafileID int) *ValueInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "create_timestamp")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_Filename(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "filename")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Mediafile_Filesize(mediafileID int) *ValueInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "filesize")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_ID(mediafileID int) *ValueInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_InheritedAccessGroupIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "inherited_access_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Mediafile_IsDirectory(mediafileID int) *ValueBool {
	key, err := dskey.FromParts("mediafile", mediafileID, "is_directory")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Mediafile_IsPublic(mediafileID int) *ValueBool {
	key, err := dskey.FromParts("mediafile", mediafileID, "is_public")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key, required: true}
}

func (r *Fetch) Mediafile_ListOfSpeakersID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "list_of_speakers_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_Mimetype(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "mimetype")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Mediafile_OwnerID(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "owner_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Mediafile_ParentID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "parent_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_PdfInformation(mediafileID int) *ValueJSON {
	key, err := dskey.FromParts("mediafile", mediafileID, "pdf_information")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) Mediafile_ProjectionIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Mediafile_Title(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Mediafile_Token(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "token")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsFontBoldInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_bold_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsFontBoldItalicInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_bold_italic_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsFontChyronSpeakerNameInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_chyron_speaker_name_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsFontItalicInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_italic_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsFontMonospaceInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_monospace_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsFontProjectorH1InMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_projector_h1_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsFontProjectorH2InMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_projector_h2_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsFontRegularInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_regular_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsLogoPdfBallotPaperInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_ballot_paper_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsLogoPdfFooterLInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_footer_l_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsLogoPdfFooterRInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_footer_r_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsLogoPdfHeaderLInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_header_l_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsLogoPdfHeaderRInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_header_r_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsLogoProjectorHeaderInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_projector_header_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsLogoProjectorMainInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_projector_main_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Mediafile_UsedAsLogoWebHeaderInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_web_header_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_AboutMe(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "about_me")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_AssignmentCandidateIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "assignment_candidate_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_ChatMessageIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "chat_message_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_Comment(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "comment")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_GroupIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_ID(meetingUserID int) *ValueInt {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MeetingUser_MeetingID(meetingUserID int) *ValueInt {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MeetingUser_MotionEditorIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "motion_editor_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_MotionSubmitterIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "motion_submitter_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_MotionWorkingGroupSpeakerIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "motion_working_group_speaker_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_Number(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "number")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_PersonalNoteIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "personal_note_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_SpeakerIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "speaker_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_StructureLevelIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "structure_level_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_SupportedMotionIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "supported_motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_UserID(meetingUserID int) *ValueInt {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "user_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MeetingUser_VoteDelegatedToID(meetingUserID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "vote_delegated_to_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_VoteDelegationsFromIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "vote_delegations_from_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingUser_VoteWeight(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "vote_weight")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AdminGroupID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "admin_group_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaEnableNumbering(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_enable_numbering")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaItemCreation(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_item_creation")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaItemIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_item_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaNewItemsDefaultVisibility(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_new_items_default_visibility")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaNumberPrefix(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_number_prefix")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaNumeralSystem(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_numeral_system")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaShowInternalItemsOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_show_internal_items_on_projector")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaShowSubtitles(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_show_subtitles")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_AgendaShowTopicNavigationOnDetailView(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_show_topic_navigation_on_detail_view")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_AllProjectionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "all_projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ApplauseEnable(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "applause_enable")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ApplauseMaxAmount(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "applause_max_amount")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ApplauseMinAmount(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "applause_min_amount")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ApplauseParticleImageUrl(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "applause_particle_image_url")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_ApplauseShowLevel(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "applause_show_level")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ApplauseTimeout(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "applause_timeout")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ApplauseType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "applause_type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentCandidateIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_candidate_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollAddCandidatesToListOfSpeakers(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_add_candidates_to_list_of_speakers")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollBallotPaperNumber(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_ballot_paper_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollBallotPaperSelection(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_ballot_paper_selection")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollDefaultBackend(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_backend")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollDefaultMethod(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_method")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_onehundred_percent_base")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollDefaultType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollEnableMaxVotesPerOption(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_enable_max_votes_per_option")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentPollSortPollResultByVotes(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_sort_poll_result_by_votes")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentsExportPreamble(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignments_export_preamble")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_AssignmentsExportTitle(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignments_export_title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_ChatGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "chat_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ChatMessageIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "chat_message_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_CommitteeID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "committee_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_ConferenceAutoConnect(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_auto_connect")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ConferenceAutoConnectNextSpeakers(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "conference_auto_connect_next_speakers")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ConferenceEnableHelpdesk(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_enable_helpdesk")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ConferenceLosRestriction(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_los_restriction")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ConferenceOpenMicrophone(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_open_microphone")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ConferenceOpenVideo(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_open_video")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ConferenceShow(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_show")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ConferenceStreamPosterUrl(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "conference_stream_poster_url")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_ConferenceStreamUrl(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "conference_stream_url")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_CustomTranslations(meetingID int) *ValueJSON {
	key, err := dskey.FromParts("meeting", meetingID, "custom_translations")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) Meeting_DefaultGroupID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "default_group_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultMeetingForCommitteeID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "default_meeting_for_committee_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_DefaultProjectorAgendaItemListIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_agenda_item_list_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorAmendmentIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_amendment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorAssignmentIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_assignment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorAssignmentPollIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_assignment_poll_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_countdown_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorCurrentListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_current_list_of_speakers_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_list_of_speakers_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorMediafileIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_mediafile_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorMessageIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_message_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorMotionBlockIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_motion_block_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorMotionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorMotionPollIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_motion_poll_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorPollIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_poll_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_DefaultProjectorTopicIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_topic_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_Description(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "description")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_EnableAnonymous(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "enable_anonymous")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_EndTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "end_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportCsvEncoding(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "export_csv_encoding")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportCsvSeparator(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "export_csv_separator")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportPdfFontsize(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_fontsize")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportPdfLineHeight(meetingID int) *ValueFloat {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_line_height")
	if err != nil {
		return &ValueFloat{err: err}
	}

	return &ValueFloat{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportPdfPageMarginBottom(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_page_margin_bottom")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportPdfPageMarginLeft(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_page_margin_left")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportPdfPageMarginRight(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_page_margin_right")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportPdfPageMarginTop(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_page_margin_top")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportPdfPagenumberAlignment(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_pagenumber_alignment")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExportPdfPagesize(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_pagesize")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_ExternalID(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "external_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_FontBoldID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_bold_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_FontBoldItalicID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_bold_italic_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_FontChyronSpeakerNameID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_chyron_speaker_name_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_FontItalicID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_italic_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_FontMonospaceID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_monospace_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_FontProjectorH1ID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_projector_h1_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_FontProjectorH2ID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_projector_h2_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_FontRegularID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_regular_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ForwardedMotionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "forwarded_motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_GroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ImportedAt(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "imported_at")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_IsActiveInOrganizationID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "is_active_in_organization_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_IsArchivedInOrganizationID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "is_archived_in_organization_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_JitsiDomain(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "jitsi_domain")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_JitsiRoomName(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "jitsi_room_name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_JitsiRoomPassword(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "jitsi_room_password")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_Language(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "language")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_ListOfSpeakersAllowMultipleSpeakers(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_allow_multiple_speakers")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersAmountLastOnProjector(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_amount_last_on_projector")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersAmountNextOnProjector(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_amount_next_on_projector")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersCanCreatePointOfOrderForOthers(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_can_create_point_of_order_for_others")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersCanSetContributionSelf(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_can_set_contribution_self")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersClosingDisablesPointOfOrder(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_closing_disables_point_of_order")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersCountdownID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_countdown_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersCoupleCountdown(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_couple_countdown")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersDefaultStructureLevelTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_default_structure_level_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersEnableInterposedQuestion(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_enable_interposed_question")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersEnablePointOfOrderCategories(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_enable_point_of_order_categories")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_enable_point_of_order_speakers")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersEnableProContraSpeech(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_enable_pro_contra_speech")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersHideContributionCount(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_hide_contribution_count")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersInitiallyClosed(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_initially_closed")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersInterventionTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_intervention_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersPresentUsersOnly(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_present_users_only")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_show_amount_of_speakers_on_slide")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersShowFirstContribution(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_show_first_contribution")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_ListOfSpeakersSpeakerNoteForEveryone(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_speaker_note_for_everyone")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_Location(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "location")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_LogoPdfBallotPaperID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_ballot_paper_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_LogoPdfFooterLID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_footer_l_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_LogoPdfFooterRID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_footer_r_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_LogoPdfHeaderLID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_header_l_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_LogoPdfHeaderRID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_header_r_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_LogoProjectorHeaderID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_projector_header_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_LogoProjectorMainID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_projector_main_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_LogoWebHeaderID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_web_header_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_MediafileIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "mediafile_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MeetingUserIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "meeting_user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionBlockIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_block_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionCategoryIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_category_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionChangeRecommendationIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_change_recommendation_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionCommentIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_comment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionCommentSectionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_comment_section_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionEditorIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_editor_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionPollBallotPaperNumber(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_ballot_paper_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionPollBallotPaperSelection(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_ballot_paper_selection")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionPollDefaultBackend(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_backend")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionPollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_onehundred_percent_base")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionPollDefaultType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionStateIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_state_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionStatuteParagraphIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_statute_paragraph_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionSubmitterIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_submitter_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionWorkflowIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_workflow_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionWorkingGroupSpeakerIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_working_group_speaker_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsAmendmentsEnabled(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_enabled")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsAmendmentsInMainList(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_in_main_list")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsAmendmentsMultipleParagraphs(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_multiple_paragraphs")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsAmendmentsOfAmendments(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_of_amendments")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsAmendmentsPrefix(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_prefix")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsAmendmentsTextMode(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_text_mode")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsBlockSlideColumns(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_block_slide_columns")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsDefaultAmendmentWorkflowID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_amendment_workflow_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_MotionsDefaultLineNumbering(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_line_numbering")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsDefaultSorting(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_sorting")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsDefaultStatuteAmendmentWorkflowID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_statute_amendment_workflow_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_MotionsDefaultWorkflowID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_workflow_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_MotionsEnableEditor(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_editor")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsEnableReasonOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_reason_on_projector")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsEnableRecommendationOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_recommendation_on_projector")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsEnableSideboxOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_sidebox_on_projector")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsEnableTextOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_text_on_projector")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsEnableWorkingGroupSpeaker(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_working_group_speaker")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsExportFollowRecommendation(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_export_follow_recommendation")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsExportPreamble(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_export_preamble")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsExportSubmitterRecommendation(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_export_submitter_recommendation")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsExportTitle(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_export_title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsLineLength(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_line_length")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsNumberMinDigits(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_number_min_digits")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsNumberType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_number_type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsNumberWithBlank(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_number_with_blank")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsPreamble(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_preamble")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsReasonRequired(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_reason_required")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsRecommendationTextMode(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_recommendation_text_mode")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsRecommendationsBy(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_recommendations_by")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsShowReferringMotions(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_show_referring_motions")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsShowSequentialNumber(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_show_sequential_number")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsStatuteRecommendationsBy(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_statute_recommendations_by")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsStatutesEnabled(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_statutes_enabled")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_MotionsSupportersMinAmount(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_supporters_min_amount")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_Name(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_OptionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "option_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_OrganizationTagIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "organization_tag_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_PersonalNoteIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "personal_note_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_PointOfOrderCategoryIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "point_of_order_category_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollBallotPaperNumber(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "poll_ballot_paper_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollBallotPaperSelection(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_ballot_paper_selection")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollCandidateIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "poll_candidate_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollCandidateListIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "poll_candidate_list_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollCountdownID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "poll_countdown_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollCoupleCountdown(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "poll_couple_countdown")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollDefaultBackend(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_backend")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollDefaultMethod(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_method")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_onehundred_percent_base")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollDefaultType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "poll_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_PollSortPollResultByVotes(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "poll_sort_poll_result_by_votes")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_PresentUserIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "present_user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ProjectionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ProjectorCountdownDefaultTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "projector_countdown_default_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_ProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "projector_countdown_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ProjectorCountdownWarningTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "projector_countdown_warning_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_ProjectorIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "projector_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ProjectorMessageIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "projector_message_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_ReferenceProjectorID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "reference_projector_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Meeting_SpeakerIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "speaker_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_StartTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "start_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_StructureLevelIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "structure_level_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_StructureLevelListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "structure_level_list_of_speakers_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_TagIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "tag_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_TemplateForOrganizationID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "template_for_organization_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Meeting_TopicIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "topic_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_TopicPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "topic_poll_default_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_UserIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersAllowSelfSetPresent(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_allow_self_set_present")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersEmailBody(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_email_body")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersEmailReplyto(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_email_replyto")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersEmailSender(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_email_sender")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersEmailSubject(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_email_subject")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersEnablePresenceView(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_enable_presence_view")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersEnableVoteDelegations(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_enable_vote_delegations")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersEnableVoteWeight(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_enable_vote_weight")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersForbidDelegatorAsSubmitter(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_forbid_delegator_as_submitter")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersForbidDelegatorAsSupporter(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_forbid_delegator_as_supporter")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersForbidDelegatorInListOfSpeakers(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_forbid_delegator_in_list_of_speakers")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersForbidDelegatorToVote(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_forbid_delegator_to_vote")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersPdfWelcometext(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_welcometext")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersPdfWelcometitle(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_welcometitle")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersPdfWlanEncryption(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_wlan_encryption")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersPdfWlanPassword(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_wlan_password")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_UsersPdfWlanSsid(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_wlan_ssid")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_VoteIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "vote_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Meeting_WelcomeText(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "welcome_text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Meeting_WelcomeTitle(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "welcome_title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionBlock_AgendaItemID(motionBlockID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "agenda_item_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MotionBlock_ID(motionBlockID int) *ValueInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionBlock_Internal(motionBlockID int) *ValueBool {
	key, err := dskey.FromParts("motion_block", motionBlockID, "internal")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionBlock_ListOfSpeakersID(motionBlockID int) *ValueInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "list_of_speakers_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionBlock_MeetingID(motionBlockID int) *ValueInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionBlock_MotionIDs(motionBlockID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_block", motionBlockID, "motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionBlock_ProjectionIDs(motionBlockID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_block", motionBlockID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionBlock_SequentialNumber(motionBlockID int) *ValueInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionBlock_Title(motionBlockID int) *ValueString {
	key, err := dskey.FromParts("motion_block", motionBlockID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionCategory_ChildIDs(motionCategoryID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "child_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionCategory_ID(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionCategory_Level(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "level")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionCategory_MeetingID(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionCategory_MotionIDs(motionCategoryID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionCategory_Name(motionCategoryID int) *ValueString {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionCategory_ParentID(motionCategoryID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "parent_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MotionCategory_Prefix(motionCategoryID int) *ValueString {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "prefix")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionCategory_SequentialNumber(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionCategory_Weight(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_CreationTime(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "creation_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_ID(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_Internal(motionChangeRecommendationID int) *ValueBool {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "internal")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_LineFrom(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "line_from")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_LineTo(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "line_to")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_MeetingID(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionChangeRecommendation_MotionID(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "motion_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionChangeRecommendation_OtherDescription(motionChangeRecommendationID int) *ValueString {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "other_description")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_Rejected(motionChangeRecommendationID int) *ValueBool {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "rejected")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_Text(motionChangeRecommendationID int) *ValueString {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionChangeRecommendation_Type(motionChangeRecommendationID int) *ValueString {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionCommentSection_CommentIDs(motionCommentSectionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "comment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionCommentSection_ID(motionCommentSectionID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionCommentSection_MeetingID(motionCommentSectionID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionCommentSection_Name(motionCommentSectionID int) *ValueString {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionCommentSection_ReadGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "read_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionCommentSection_SequentialNumber(motionCommentSectionID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionCommentSection_SubmitterCanWrite(motionCommentSectionID int) *ValueBool {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "submitter_can_write")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionCommentSection_Weight(motionCommentSectionID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionCommentSection_WriteGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "write_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionComment_Comment(motionCommentID int) *ValueString {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "comment")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionComment_ID(motionCommentID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionComment_MeetingID(motionCommentID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionComment_MotionID(motionCommentID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "motion_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionComment_SectionID(motionCommentID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "section_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionEditor_ID(motionEditorID int) *ValueInt {
	key, err := dskey.FromParts("motion_editor", motionEditorID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionEditor_MeetingID(motionEditorID int) *ValueInt {
	key, err := dskey.FromParts("motion_editor", motionEditorID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionEditor_MeetingUserID(motionEditorID int) *ValueInt {
	key, err := dskey.FromParts("motion_editor", motionEditorID, "meeting_user_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionEditor_MotionID(motionEditorID int) *ValueInt {
	key, err := dskey.FromParts("motion_editor", motionEditorID, "motion_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionEditor_Weight(motionEditorID int) *ValueInt {
	key, err := dskey.FromParts("motion_editor", motionEditorID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionState_AllowCreatePoll(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "allow_create_poll")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_AllowMotionForwarding(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "allow_motion_forwarding")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_AllowSubmitterEdit(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "allow_submitter_edit")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_AllowSupport(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "allow_support")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_CssClass(motionStateID int) *ValueString {
	key, err := dskey.FromParts("motion_state", motionStateID, "css_class")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionState_FirstStateOfWorkflowID(motionStateID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "first_state_of_workflow_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MotionState_ID(motionStateID int) *ValueInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionState_IsInternal(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "is_internal")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_MeetingID(motionStateID int) *ValueInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionState_MergeAmendmentIntoFinal(motionStateID int) *ValueString {
	key, err := dskey.FromParts("motion_state", motionStateID, "merge_amendment_into_final")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionState_MotionIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionState_MotionRecommendationIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "motion_recommendation_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionState_Name(motionStateID int) *ValueString {
	key, err := dskey.FromParts("motion_state", motionStateID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionState_NextStateIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "next_state_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionState_PreviousStateIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "previous_state_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionState_RecommendationLabel(motionStateID int) *ValueString {
	key, err := dskey.FromParts("motion_state", motionStateID, "recommendation_label")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionState_Restrictions(motionStateID int) *ValueStringSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "restrictions")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) MotionState_SetNumber(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "set_number")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_SetWorkflowTimestamp(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "set_workflow_timestamp")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_ShowRecommendationExtensionField(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "show_recommendation_extension_field")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_ShowStateExtensionField(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "show_state_extension_field")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) MotionState_SubmitterWithdrawBackIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "submitter_withdraw_back_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionState_SubmitterWithdrawStateID(motionStateID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "submitter_withdraw_state_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MotionState_Weight(motionStateID int) *ValueInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionState_WorkflowID(motionStateID int) *ValueInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "workflow_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionStatuteParagraph_ID(motionStatuteParagraphID int) *ValueInt {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionStatuteParagraph_MeetingID(motionStatuteParagraphID int) *ValueInt {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionStatuteParagraph_MotionIDs(motionStatuteParagraphID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionStatuteParagraph_SequentialNumber(motionStatuteParagraphID int) *ValueInt {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionStatuteParagraph_Text(motionStatuteParagraphID int) *ValueString {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) MotionStatuteParagraph_Title(motionStatuteParagraphID int) *ValueString {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionStatuteParagraph_Weight(motionStatuteParagraphID int) *ValueInt {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionSubmitter_ID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionSubmitter_MeetingID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionSubmitter_MeetingUserID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "meeting_user_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionSubmitter_MotionID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "motion_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionSubmitter_Weight(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionWorkflow_DefaultAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "default_amendment_workflow_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MotionWorkflow_DefaultStatuteAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "default_statute_amendment_workflow_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MotionWorkflow_DefaultWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "default_workflow_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MotionWorkflow_FirstStateID(motionWorkflowID int) *ValueInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "first_state_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionWorkflow_ID(motionWorkflowID int) *ValueInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionWorkflow_MeetingID(motionWorkflowID int) *ValueInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionWorkflow_Name(motionWorkflowID int) *ValueString {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionWorkflow_SequentialNumber(motionWorkflowID int) *ValueInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionWorkflow_StateIDs(motionWorkflowID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "state_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MotionWorkingGroupSpeaker_ID(motionWorkingGroupSpeakerID int) *ValueInt {
	key, err := dskey.FromParts("motion_working_group_speaker", motionWorkingGroupSpeakerID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) MotionWorkingGroupSpeaker_MeetingID(motionWorkingGroupSpeakerID int) *ValueInt {
	key, err := dskey.FromParts("motion_working_group_speaker", motionWorkingGroupSpeakerID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionWorkingGroupSpeaker_MeetingUserID(motionWorkingGroupSpeakerID int) *ValueInt {
	key, err := dskey.FromParts("motion_working_group_speaker", motionWorkingGroupSpeakerID, "meeting_user_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionWorkingGroupSpeaker_MotionID(motionWorkingGroupSpeakerID int) *ValueInt {
	key, err := dskey.FromParts("motion_working_group_speaker", motionWorkingGroupSpeakerID, "motion_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MotionWorkingGroupSpeaker_Weight(motionWorkingGroupSpeakerID int) *ValueInt {
	key, err := dskey.FromParts("motion_working_group_speaker", motionWorkingGroupSpeakerID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_AdditionalSubmitter(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "additional_submitter")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Motion_AgendaItemID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "agenda_item_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_AllDerivedMotionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "all_derived_motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_AllOriginIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "all_origin_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_AmendmentIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "amendment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_AmendmentParagraphs(motionID int) *ValueJSON {
	key, err := dskey.FromParts("motion", motionID, "amendment_paragraphs")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) Motion_AttachmentIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "attachment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_BlockID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "block_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_CategoryID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "category_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_CategoryWeight(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "category_weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_ChangeRecommendationIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "change_recommendation_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_CommentIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "comment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_Created(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "created")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_DerivedMotionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "derived_motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_EditorIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "editor_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_Forwarded(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "forwarded")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_ID(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_IDenticalMotionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "identical_motion_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_LastModified(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "last_modified")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_LeadMotionID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "lead_motion_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_ListOfSpeakersID(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "list_of_speakers_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Motion_MeetingID(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Motion_ModifiedFinalVersion(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "modified_final_version")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Motion_Number(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "number")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Motion_NumberValue(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "number_value")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_OptionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "option_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_OriginID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "origin_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_OriginMeetingID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "origin_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_PersonalNoteIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "personal_note_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_PollIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "poll_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_ProjectionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_Reason(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "reason")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Motion_RecommendationExtension(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "recommendation_extension")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Motion_RecommendationExtensionReferenceIDs(motionID int) *ValueStringSlice {
	key, err := dskey.FromParts("motion", motionID, "recommendation_extension_reference_ids")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_RecommendationID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "recommendation_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_ReferencedInMotionRecommendationExtensionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "referenced_in_motion_recommendation_extension_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_ReferencedInMotionStateExtensionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "referenced_in_motion_state_extension_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_SequentialNumber(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Motion_SortChildIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "sort_child_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_SortParentID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "sort_parent_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_SortWeight(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "sort_weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_StartLineNumber(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "start_line_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_StateExtension(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "state_extension")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Motion_StateExtensionReferenceIDs(motionID int) *ValueStringSlice {
	key, err := dskey.FromParts("motion", motionID, "state_extension_reference_ids")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_StateID(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "state_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Motion_StatuteParagraphID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "statute_paragraph_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Motion_SubmitterIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "submitter_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_SupporterMeetingUserIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "supporter_meeting_user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_TagIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "tag_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Motion_Text(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Motion_TextHash(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "text_hash")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Motion_Title(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Motion_WorkflowTimestamp(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "workflow_timestamp")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Motion_WorkingGroupSpeakerIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "working_group_speaker_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Option_Abstain(optionID int) *ValueString {
	key, err := dskey.FromParts("option", optionID, "abstain")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Option_ContentObjectID(optionID int) *ValueMaybeString {
	key, err := dskey.FromParts("option", optionID, "content_object_id")
	if err != nil {
		return &ValueMaybeString{err: err}
	}

	return &ValueMaybeString{fetch: r, key: key}
}

func (r *Fetch) Option_ID(optionID int) *ValueInt {
	key, err := dskey.FromParts("option", optionID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Option_MeetingID(optionID int) *ValueInt {
	key, err := dskey.FromParts("option", optionID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Option_No(optionID int) *ValueString {
	key, err := dskey.FromParts("option", optionID, "no")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Option_PollID(optionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("option", optionID, "poll_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Option_Text(optionID int) *ValueString {
	key, err := dskey.FromParts("option", optionID, "text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Option_UsedAsGlobalOptionInPollID(optionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("option", optionID, "used_as_global_option_in_poll_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Option_VoteIDs(optionID int) *ValueIntSlice {
	key, err := dskey.FromParts("option", optionID, "vote_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Option_Weight(optionID int) *ValueInt {
	key, err := dskey.FromParts("option", optionID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Option_Yes(optionID int) *ValueString {
	key, err := dskey.FromParts("option", optionID, "yes")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) OrganizationTag_Color(organizationTagID int) *ValueString {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) OrganizationTag_ID(organizationTagID int) *ValueInt {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) OrganizationTag_Name(organizationTagID int) *ValueString {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) OrganizationTag_OrganizationID(organizationTagID int) *ValueInt {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "organization_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) OrganizationTag_TaggedIDs(organizationTagID int) *ValueStringSlice {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "tagged_ids")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_ActiveMeetingIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "active_meeting_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_ArchivedMeetingIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "archived_meeting_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_CommitteeIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "committee_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_DefaultLanguage(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "default_language")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Organization_Description(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "description")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_EnableChat(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "enable_chat")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Organization_EnableElectronicVoting(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "enable_electronic_voting")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Organization_Genders(organizationID int) *ValueStringSlice {
	key, err := dskey.FromParts("organization", organizationID, "genders")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_ID(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Organization_LegalNotice(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "legal_notice")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_LimitOfMeetings(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "limit_of_meetings")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Organization_LimitOfUsers(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "limit_of_users")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Organization_LoginText(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "login_text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_MediafileIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "mediafile_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_Name(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_OrganizationTagIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "organization_tag_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_PrivacyPolicy(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "privacy_policy")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_ResetPasswordVerboseErrors(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "reset_password_verbose_errors")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Organization_SamlAttrMapping(organizationID int) *ValueJSON {
	key, err := dskey.FromParts("organization", organizationID, "saml_attr_mapping")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) Organization_SamlEnabled(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "saml_enabled")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Organization_SamlLoginButtonText(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "saml_login_button_text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_SamlMetadataIDp(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "saml_metadata_idp")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_SamlMetadataSp(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "saml_metadata_sp")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_SamlPrivateKey(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "saml_private_key")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_TemplateMeetingIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "template_meeting_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_ThemeID(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "theme_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Organization_ThemeIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "theme_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_Url(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "url")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_UserIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_UsersEmailBody(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "users_email_body")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_UsersEmailReplyto(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "users_email_replyto")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_UsersEmailSender(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "users_email_sender")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_UsersEmailSubject(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "users_email_subject")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Organization_VoteDecryptPublicMainKey(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "vote_decrypt_public_main_key")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) PersonalNote_ContentObjectID(personalNoteID int) *ValueMaybeString {
	key, err := dskey.FromParts("personal_note", personalNoteID, "content_object_id")
	if err != nil {
		return &ValueMaybeString{err: err}
	}

	return &ValueMaybeString{fetch: r, key: key}
}

func (r *Fetch) PersonalNote_ID(personalNoteID int) *ValueInt {
	key, err := dskey.FromParts("personal_note", personalNoteID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) PersonalNote_MeetingID(personalNoteID int) *ValueInt {
	key, err := dskey.FromParts("personal_note", personalNoteID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) PersonalNote_MeetingUserID(personalNoteID int) *ValueInt {
	key, err := dskey.FromParts("personal_note", personalNoteID, "meeting_user_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) PersonalNote_Note(personalNoteID int) *ValueString {
	key, err := dskey.FromParts("personal_note", personalNoteID, "note")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) PersonalNote_Star(personalNoteID int) *ValueBool {
	key, err := dskey.FromParts("personal_note", personalNoteID, "star")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) PointOfOrderCategory_ID(pointOfOrderCategoryID int) *ValueInt {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) PointOfOrderCategory_MeetingID(pointOfOrderCategoryID int) *ValueInt {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) PointOfOrderCategory_Rank(pointOfOrderCategoryID int) *ValueInt {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "rank")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) PointOfOrderCategory_SpeakerIDs(pointOfOrderCategoryID int) *ValueIntSlice {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "speaker_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) PointOfOrderCategory_Text(pointOfOrderCategoryID int) *ValueString {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) PollCandidateList_ID(pollCandidateListID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate_list", pollCandidateListID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) PollCandidateList_MeetingID(pollCandidateListID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate_list", pollCandidateListID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) PollCandidateList_OptionID(pollCandidateListID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate_list", pollCandidateListID, "option_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) PollCandidateList_PollCandidateIDs(pollCandidateListID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll_candidate_list", pollCandidateListID, "poll_candidate_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) PollCandidate_ID(pollCandidateID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) PollCandidate_MeetingID(pollCandidateID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) PollCandidate_PollCandidateListID(pollCandidateID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "poll_candidate_list_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) PollCandidate_UserID(pollCandidateID int) *ValueMaybeInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "user_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) PollCandidate_Weight(pollCandidateID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_Backend(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "backend")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_ContentObjectID(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "content_object_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_CryptKey(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "crypt_key")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Poll_CryptSignature(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "crypt_signature")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Poll_Description(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "description")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Poll_EntitledGroupIDs(pollID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll", pollID, "entitled_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Poll_EntitledUsersAtStop(pollID int) *ValueJSON {
	key, err := dskey.FromParts("poll", pollID, "entitled_users_at_stop")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) Poll_GlobalAbstain(pollID int) *ValueBool {
	key, err := dskey.FromParts("poll", pollID, "global_abstain")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Poll_GlobalNo(pollID int) *ValueBool {
	key, err := dskey.FromParts("poll", pollID, "global_no")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Poll_GlobalOptionID(pollID int) *ValueMaybeInt {
	key, err := dskey.FromParts("poll", pollID, "global_option_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Poll_GlobalYes(pollID int) *ValueBool {
	key, err := dskey.FromParts("poll", pollID, "global_yes")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Poll_ID(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Poll_IsPseudoanonymized(pollID int) *ValueBool {
	key, err := dskey.FromParts("poll", pollID, "is_pseudoanonymized")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Poll_MaxVotesAmount(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "max_votes_amount")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Poll_MaxVotesPerOption(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "max_votes_per_option")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Poll_MeetingID(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_MinVotesAmount(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "min_votes_amount")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Poll_OnehundredPercentBase(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "onehundred_percent_base")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_OptionIDs(pollID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll", pollID, "option_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Poll_Pollmethod(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "pollmethod")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_ProjectionIDs(pollID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll", pollID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Poll_SequentialNumber(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_State(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "state")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Poll_Title(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_Type(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Poll_VoteCount(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "vote_count")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Poll_VotedIDs(pollID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll", pollID, "voted_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Poll_VotesRaw(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votes_raw")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Poll_VotesSignature(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votes_signature")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Poll_Votescast(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votescast")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Poll_Votesinvalid(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votesinvalid")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Poll_Votesvalid(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votesvalid")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projection_Content(projectionID int) *ValueJSON {
	key, err := dskey.FromParts("projection", projectionID, "content")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) Projection_ContentObjectID(projectionID int) *ValueString {
	key, err := dskey.FromParts("projection", projectionID, "content_object_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Projection_CurrentProjectorID(projectionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projection", projectionID, "current_projector_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projection_HistoryProjectorID(projectionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projection", projectionID, "history_projector_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projection_ID(projectionID int) *ValueInt {
	key, err := dskey.FromParts("projection", projectionID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Projection_MeetingID(projectionID int) *ValueInt {
	key, err := dskey.FromParts("projection", projectionID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Projection_Options(projectionID int) *ValueJSON {
	key, err := dskey.FromParts("projection", projectionID, "options")
	if err != nil {
		return &ValueJSON{err: err}
	}

	return &ValueJSON{fetch: r, key: key}
}

func (r *Fetch) Projection_PreviewProjectorID(projectionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projection", projectionID, "preview_projector_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projection_Stable(projectionID int) *ValueBool {
	key, err := dskey.FromParts("projection", projectionID, "stable")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Projection_Type(projectionID int) *ValueString {
	key, err := dskey.FromParts("projection", projectionID, "type")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projection_Weight(projectionID int) *ValueInt {
	key, err := dskey.FromParts("projection", projectionID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ProjectorCountdown_CountdownTime(projectorCountdownID int) *ValueFloat {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "countdown_time")
	if err != nil {
		return &ValueFloat{err: err}
	}

	return &ValueFloat{fetch: r, key: key}
}

func (r *Fetch) ProjectorCountdown_DefaultTime(projectorCountdownID int) *ValueInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "default_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ProjectorCountdown_Description(projectorCountdownID int) *ValueString {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "description")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) ProjectorCountdown_ID(projectorCountdownID int) *ValueInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ProjectorCountdown_MeetingID(projectorCountdownID int) *ValueInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ProjectorCountdown_ProjectionIDs(projectorCountdownID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) ProjectorCountdown_Running(projectorCountdownID int) *ValueBool {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "running")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) ProjectorCountdown_Title(projectorCountdownID int) *ValueString {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "used_as_list_of_speakers_countdown_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) ProjectorCountdown_UsedAsPollCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "used_as_poll_countdown_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) ProjectorMessage_ID(projectorMessageID int) *ValueInt {
	key, err := dskey.FromParts("projector_message", projectorMessageID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) ProjectorMessage_MeetingID(projectorMessageID int) *ValueInt {
	key, err := dskey.FromParts("projector_message", projectorMessageID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ProjectorMessage_Message(projectorMessageID int) *ValueString {
	key, err := dskey.FromParts("projector_message", projectorMessageID, "message")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) ProjectorMessage_ProjectionIDs(projectorMessageID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector_message", projectorMessageID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Projector_AspectRatioDenominator(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "aspect_ratio_denominator")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Projector_AspectRatioNumerator(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "aspect_ratio_numerator")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Projector_BackgroundColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "background_color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_ChyronBackgroundColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "chyron_background_color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_ChyronBackgroundColor2(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "chyron_background_color_2")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_ChyronFontColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "chyron_font_color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_ChyronFontColor2(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "chyron_font_color_2")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_Color(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_CurrentProjectionIDs(projectorID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector", projectorID, "current_projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Projector_HeaderBackgroundColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "header_background_color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_HeaderFontColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "header_font_color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_HeaderH1Color(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "header_h1_color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_HistoryProjectionIDs(projectorID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector", projectorID, "history_projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Projector_ID(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Projector_IsInternal(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "is_internal")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Projector_MeetingID(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Projector_Name(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Projector_PreviewProjectionIDs(projectorID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector", projectorID, "preview_projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Projector_Scale(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "scale")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Projector_Scroll(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "scroll")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Projector_SequentialNumber(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Projector_ShowClock(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "show_clock")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Projector_ShowHeaderFooter(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "show_header_footer")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Projector_ShowLogo(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "show_logo")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Projector_ShowTitle(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "show_title")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAgendaItemListInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_agenda_item_list_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAmendmentInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_amendment_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAssignmentInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_assignment_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAssignmentPollInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_assignment_poll_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForCountdownInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_countdown_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_current_list_of_speakers_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForListOfSpeakersInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_list_of_speakers_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMediafileInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_mediafile_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMessageInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_message_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionBlockInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_motion_block_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_motion_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionPollInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_motion_poll_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForPollInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_poll_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForTopicInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_topic_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_UsedAsReferenceProjectorMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_reference_projector_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Projector_Width(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "width")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_BeginTime(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "begin_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_EndTime(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "end_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_ID(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_ListOfSpeakersID(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "list_of_speakers_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Speaker_MeetingID(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Speaker_MeetingUserID(speakerID int) *ValueMaybeInt {
	key, err := dskey.FromParts("speaker", speakerID, "meeting_user_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_Note(speakerID int) *ValueString {
	key, err := dskey.FromParts("speaker", speakerID, "note")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Speaker_PauseTime(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "pause_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_PointOfOrder(speakerID int) *ValueBool {
	key, err := dskey.FromParts("speaker", speakerID, "point_of_order")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Speaker_PointOfOrderCategoryID(speakerID int) *ValueMaybeInt {
	key, err := dskey.FromParts("speaker", speakerID, "point_of_order_category_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_SpeechState(speakerID int) *ValueString {
	key, err := dskey.FromParts("speaker", speakerID, "speech_state")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Speaker_StructureLevelListOfSpeakersID(speakerID int) *ValueMaybeInt {
	key, err := dskey.FromParts("speaker", speakerID, "structure_level_list_of_speakers_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_TotalPause(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "total_pause")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_UnpauseTime(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "unpause_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Speaker_Weight(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "weight")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) StructureLevelListOfSpeakers_AdditionalTime(structureLevelListOfSpeakersID int) *ValueFloat {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "additional_time")
	if err != nil {
		return &ValueFloat{err: err}
	}

	return &ValueFloat{fetch: r, key: key}
}

func (r *Fetch) StructureLevelListOfSpeakers_CurrentStartTime(structureLevelListOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "current_start_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) StructureLevelListOfSpeakers_ID(structureLevelListOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevelListOfSpeakers_InitialTime(structureLevelListOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "initial_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevelListOfSpeakers_ListOfSpeakersID(structureLevelListOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "list_of_speakers_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevelListOfSpeakers_MeetingID(structureLevelListOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevelListOfSpeakers_RemainingTime(structureLevelListOfSpeakersID int) *ValueFloat {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "remaining_time")
	if err != nil {
		return &ValueFloat{err: err}
	}

	return &ValueFloat{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevelListOfSpeakers_SpeakerIDs(structureLevelListOfSpeakersID int) *ValueIntSlice {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "speaker_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) StructureLevelListOfSpeakers_StructureLevelID(structureLevelListOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("structure_level_list_of_speakers", structureLevelListOfSpeakersID, "structure_level_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevel_Color(structureLevelID int) *ValueString {
	key, err := dskey.FromParts("structure_level", structureLevelID, "color")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) StructureLevel_DefaultTime(structureLevelID int) *ValueInt {
	key, err := dskey.FromParts("structure_level", structureLevelID, "default_time")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) StructureLevel_ID(structureLevelID int) *ValueInt {
	key, err := dskey.FromParts("structure_level", structureLevelID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevel_MeetingID(structureLevelID int) *ValueInt {
	key, err := dskey.FromParts("structure_level", structureLevelID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevel_MeetingUserIDs(structureLevelID int) *ValueIntSlice {
	key, err := dskey.FromParts("structure_level", structureLevelID, "meeting_user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) StructureLevel_Name(structureLevelID int) *ValueString {
	key, err := dskey.FromParts("structure_level", structureLevelID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) StructureLevel_StructureLevelListOfSpeakersIDs(structureLevelID int) *ValueIntSlice {
	key, err := dskey.FromParts("structure_level", structureLevelID, "structure_level_list_of_speakers_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Tag_ID(tagID int) *ValueInt {
	key, err := dskey.FromParts("tag", tagID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Tag_MeetingID(tagID int) *ValueInt {
	key, err := dskey.FromParts("tag", tagID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Tag_Name(tagID int) *ValueString {
	key, err := dskey.FromParts("tag", tagID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Tag_TaggedIDs(tagID int) *ValueStringSlice {
	key, err := dskey.FromParts("tag", tagID, "tagged_ids")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) Theme_Abstain(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "abstain")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_100")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_200")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent300(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_300")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_400")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent50(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_50")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent500(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_500")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Theme_Accent600(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_600")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_700")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent800(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_800")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Accent900(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_900")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_AccentA100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_a100")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_AccentA200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_a200")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_AccentA400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_a400")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_AccentA700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_a700")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Headbar(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "headbar")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_ID(themeID int) *ValueInt {
	key, err := dskey.FromParts("theme", themeID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Theme_Name(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Theme_No(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "no")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_OrganizationID(themeID int) *ValueInt {
	key, err := dskey.FromParts("theme", themeID, "organization_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Theme_Primary100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_100")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Primary200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_200")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Primary300(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_300")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Primary400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_400")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Primary50(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_50")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Primary500(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_500")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Theme_Primary600(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_600")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Primary700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_700")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Primary800(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_800")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Primary900(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_900")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_PrimaryA100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_a100")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_PrimaryA200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_a200")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_PrimaryA400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_a400")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_PrimaryA700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_a700")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_ThemeForOrganizationID(themeID int) *ValueMaybeInt {
	key, err := dskey.FromParts("theme", themeID, "theme_for_organization_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_100")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_200")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn300(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_300")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_400")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn50(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_50")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn500(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_500")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Theme_Warn600(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_600")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_700")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn800(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_800")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Warn900(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_900")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_WarnA100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_a100")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_WarnA200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_a200")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_WarnA400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_a400")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_WarnA700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_a700")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Theme_Yes(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "yes")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Topic_AgendaItemID(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "agenda_item_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Topic_AttachmentIDs(topicID int) *ValueIntSlice {
	key, err := dskey.FromParts("topic", topicID, "attachment_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Topic_ID(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Topic_ListOfSpeakersID(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "list_of_speakers_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Topic_MeetingID(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Topic_PollIDs(topicID int) *ValueIntSlice {
	key, err := dskey.FromParts("topic", topicID, "poll_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Topic_ProjectionIDs(topicID int) *ValueIntSlice {
	key, err := dskey.FromParts("topic", topicID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Topic_SequentialNumber(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "sequential_number")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Topic_Text(topicID int) *ValueString {
	key, err := dskey.FromParts("topic", topicID, "text")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Topic_Title(topicID int) *ValueString {
	key, err := dskey.FromParts("topic", topicID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) User_CanChangeOwnPassword(userID int) *ValueBool {
	key, err := dskey.FromParts("user", userID, "can_change_own_password")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) User_CommitteeIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "committee_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_CommitteeManagementIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "committee_management_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_DefaultPassword(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "default_password")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_DefaultVoteWeight(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "default_vote_weight")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_DelegatedVoteIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "delegated_vote_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_Email(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "email")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_FirstName(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "first_name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_ForwardingCommitteeIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "forwarding_committee_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_Gender(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "gender")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_ID(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) User_IsActive(userID int) *ValueBool {
	key, err := dskey.FromParts("user", userID, "is_active")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) User_IsDemoUser(userID int) *ValueBool {
	key, err := dskey.FromParts("user", userID, "is_demo_user")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) User_IsPhysicalPerson(userID int) *ValueBool {
	key, err := dskey.FromParts("user", userID, "is_physical_person")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) User_IsPresentInMeetingIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "is_present_in_meeting_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_LastEmailSent(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "last_email_sent")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) User_LastLogin(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "last_login")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) User_LastName(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "last_name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_MeetingIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "meeting_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_MeetingUserIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "meeting_user_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_MemberNumber(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "member_number")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_OptionIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "option_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_OrganizationID(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "organization_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) User_OrganizationManagementLevel(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "organization_management_level")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_Password(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "password")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_PollCandidateIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "poll_candidate_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_PollVotedIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "poll_voted_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) User_Pronoun(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "pronoun")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_SamlID(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "saml_id")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_Title(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "title")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) User_Username(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "username")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) User_VoteIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "vote_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Vote_DelegatedUserID(voteID int) *ValueMaybeInt {
	key, err := dskey.FromParts("vote", voteID, "delegated_user_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Vote_ID(voteID int) *ValueInt {
	key, err := dskey.FromParts("vote", voteID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key}
}

func (r *Fetch) Vote_MeetingID(voteID int) *ValueInt {
	key, err := dskey.FromParts("vote", voteID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Vote_OptionID(voteID int) *ValueInt {
	key, err := dskey.FromParts("vote", voteID, "option_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Vote_UserID(voteID int) *ValueMaybeInt {
	key, err := dskey.FromParts("vote", voteID, "user_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) Vote_UserToken(voteID int) *ValueString {
	key, err := dskey.FromParts("vote", voteID, "user_token")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Vote_Value(voteID int) *ValueString {
	key, err := dskey.FromParts("vote", voteID, "value")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}

func (r *Fetch) Vote_Weight(voteID int) *ValueString {
	key, err := dskey.FromParts("vote", voteID, "weight")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
}
