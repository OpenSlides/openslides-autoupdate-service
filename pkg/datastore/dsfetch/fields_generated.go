// Code generated from models.yml DO NOT EDIT.
package dsfetch

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// ValueBool is a lazy value from the datastore.
type ValueBool struct {
	collection string
	id         int
	field      string
	value      bool
	required   bool

	executed bool
	isNull   bool

	lazies []*bool

	fetch *Fetch
}

// Value returns the value.
func (v *ValueBool) Value(ctx context.Context) (bool, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return false, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      float32
	required   bool

	executed bool
	isNull   bool

	lazies []*float32

	fetch *Fetch
}

// Value returns the value.
func (v *ValueFloat) Value(ctx context.Context) (float32, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return 0, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      []int
	required   bool

	executed bool
	isNull   bool

	lazies []*[]int

	fetch *Fetch
}

// Value returns the value.
func (v *ValueIDSlice) Value(ctx context.Context) ([]int, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return nil, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      int
	required   bool

	executed bool
	isNull   bool

	lazies []*int

	fetch *Fetch
}

// Value returns the value.
func (v *ValueInt) Value(ctx context.Context) (int, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return 0, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      []int
	required   bool

	executed bool
	isNull   bool

	lazies []*[]int

	fetch *Fetch
}

// Value returns the value.
func (v *ValueIntSlice) Value(ctx context.Context) ([]int, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return nil, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      json.RawMessage
	required   bool

	executed bool
	isNull   bool

	lazies []*json.RawMessage

	fetch *Fetch
}

// Value returns the value.
func (v *ValueJSON) Value(ctx context.Context) (json.RawMessage, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return nil, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      int
	required   bool

	executed bool
	isNull   bool

	lazies []*int

	fetch *Fetch
}

// Value returns the value.
func (v *ValueMaybeInt) Value(ctx context.Context) (int, bool, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return 0, false, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      string
	required   bool

	executed bool
	isNull   bool

	lazies []*string

	fetch *Fetch
}

// Value returns the value.
func (v *ValueMaybeString) Value(ctx context.Context) (string, bool, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return "", false, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      string
	required   bool

	executed bool
	isNull   bool

	lazies []*string

	fetch *Fetch
}

// Value returns the value.
func (v *ValueString) Value(ctx context.Context) (string, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return "", v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	collection string
	id         int
	field      string
	value      []string
	required   bool

	executed bool
	isNull   bool

	lazies []*[]string

	fetch *Fetch
}

// Value returns the value.
func (v *ValueStringSlice) Value(ctx context.Context) ([]string, error) {
	if err := v.fetch.err; err != nil {
		v.fetch.err = nil
		return nil, v.fetch.err
	}

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
		if v.required {
			return fmt.Errorf("database is corrupted. Required field %s/%d/%s is null", v.collection, v.id, v.field)
		}
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
	key := dskey.Key{Collection: "action_worker", ID: actionWorkerID, Field: "created"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "created", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_ID(actionWorkerID int) *ValueInt {
	key := dskey.Key{Collection: "action_worker", ID: actionWorkerID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_Name(actionWorkerID int) *ValueString {
	key := dskey.Key{Collection: "action_worker", ID: actionWorkerID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_Result(actionWorkerID int) *ValueJSON {
	key := dskey.Key{Collection: "action_worker", ID: actionWorkerID, Field: "result"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "result"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_State(actionWorkerID int) *ValueString {
	key := dskey.Key{Collection: "action_worker", ID: actionWorkerID, Field: "state"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "state", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_Timestamp(actionWorkerID int) *ValueInt {
	key := dskey.Key{Collection: "action_worker", ID: actionWorkerID, Field: "timestamp"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "timestamp", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ChildIDs(agendaItemID int) *ValueIntSlice {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "child_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "agendaItem", id: agendaItemID, field: "child_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Closed(agendaItemID int) *ValueBool {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "closed"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "agendaItem", id: agendaItemID, field: "closed"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Comment(agendaItemID int) *ValueString {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "comment"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "agendaItem", id: agendaItemID, field: "comment"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ContentObjectID(agendaItemID int) *ValueString {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "content_object_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "agendaItem", id: agendaItemID, field: "content_object_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Duration(agendaItemID int) *ValueInt {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "duration"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "duration"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ID(agendaItemID int) *ValueInt {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_IsHidden(agendaItemID int) *ValueBool {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "is_hidden"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "agendaItem", id: agendaItemID, field: "is_hidden"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_IsInternal(agendaItemID int) *ValueBool {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "is_internal"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "agendaItem", id: agendaItemID, field: "is_internal"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ItemNumber(agendaItemID int) *ValueString {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "item_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "agendaItem", id: agendaItemID, field: "item_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Level(agendaItemID int) *ValueInt {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "level"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_MeetingID(agendaItemID int) *ValueInt {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ParentID(agendaItemID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "parent_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "parent_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ProjectionIDs(agendaItemID int) *ValueIntSlice {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "agendaItem", id: agendaItemID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_TagIDs(agendaItemID int) *ValueIntSlice {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "tag_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "agendaItem", id: agendaItemID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Type(agendaItemID int) *ValueString {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "agendaItem", id: agendaItemID, field: "type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Weight(agendaItemID int) *ValueInt {
	key := dskey.Key{Collection: "agenda_item", ID: agendaItemID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_AssignmentID(assignmentCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "assignment_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "assignment_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_ID(assignmentCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_MeetingID(assignmentCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_MeetingUserID(assignmentCandidateID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "meeting_user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "meeting_user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_Weight(assignmentCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "assignment_candidate", ID: assignmentCandidateID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_AgendaItemID(assignmentID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "agenda_item_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "assignment", id: assignmentID, field: "agenda_item_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_AttachmentIDs(assignmentID int) *ValueIntSlice {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "attachment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "attachment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_CandidateIDs(assignmentID int) *ValueIntSlice {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "candidate_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_DefaultPollDescription(assignmentID int) *ValueString {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "default_poll_description"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "assignment", id: assignmentID, field: "default_poll_description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_Description(assignmentID int) *ValueString {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "description"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "assignment", id: assignmentID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_ID(assignmentID int) *ValueInt {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_ListOfSpeakersID(assignmentID int) *ValueInt {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "list_of_speakers_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_MeetingID(assignmentID int) *ValueInt {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_NumberPollCandidates(assignmentID int) *ValueBool {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "number_poll_candidates"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "assignment", id: assignmentID, field: "number_poll_candidates"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_OpenPosts(assignmentID int) *ValueInt {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "open_posts"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "open_posts"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_Phase(assignmentID int) *ValueString {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "phase"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "assignment", id: assignmentID, field: "phase"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_PollIDs(assignmentID int) *ValueIntSlice {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "poll_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_ProjectionIDs(assignmentID int) *ValueIntSlice {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_SequentialNumber(assignmentID int) *ValueInt {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_TagIDs(assignmentID int) *ValueIntSlice {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "tag_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_Title(assignmentID int) *ValueString {
	key := dskey.Key{Collection: "assignment", ID: assignmentID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "assignment", id: assignmentID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_ChatMessageIDs(chatGroupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "chat_group", ID: chatGroupID, Field: "chat_message_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "chatGroup", id: chatGroupID, field: "chat_message_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_ID(chatGroupID int) *ValueInt {
	key := dskey.Key{Collection: "chat_group", ID: chatGroupID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatGroup", id: chatGroupID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_MeetingID(chatGroupID int) *ValueInt {
	key := dskey.Key{Collection: "chat_group", ID: chatGroupID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatGroup", id: chatGroupID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_Name(chatGroupID int) *ValueString {
	key := dskey.Key{Collection: "chat_group", ID: chatGroupID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "chatGroup", id: chatGroupID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_ReadGroupIDs(chatGroupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "chat_group", ID: chatGroupID, Field: "read_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "chatGroup", id: chatGroupID, field: "read_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_Weight(chatGroupID int) *ValueInt {
	key := dskey.Key{Collection: "chat_group", ID: chatGroupID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatGroup", id: chatGroupID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_WriteGroupIDs(chatGroupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "chat_group", ID: chatGroupID, Field: "write_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "chatGroup", id: chatGroupID, field: "write_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_ChatGroupID(chatMessageID int) *ValueInt {
	key := dskey.Key{Collection: "chat_message", ID: chatMessageID, Field: "chat_group_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "chat_group_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_Content(chatMessageID int) *ValueString {
	key := dskey.Key{Collection: "chat_message", ID: chatMessageID, Field: "content"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "chatMessage", id: chatMessageID, field: "content", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_Created(chatMessageID int) *ValueInt {
	key := dskey.Key{Collection: "chat_message", ID: chatMessageID, Field: "created"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "created", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_ID(chatMessageID int) *ValueInt {
	key := dskey.Key{Collection: "chat_message", ID: chatMessageID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_MeetingID(chatMessageID int) *ValueInt {
	key := dskey.Key{Collection: "chat_message", ID: chatMessageID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_MeetingUserID(chatMessageID int) *ValueInt {
	key := dskey.Key{Collection: "chat_message", ID: chatMessageID, Field: "meeting_user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "meeting_user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_DefaultMeetingID(committeeID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "default_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "committee", id: committeeID, field: "default_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_Description(committeeID int) *ValueString {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "description"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "committee", id: committeeID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ForwardToCommitteeIDs(committeeID int) *ValueIntSlice {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "forward_to_committee_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "forward_to_committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ForwardingUserID(committeeID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "forwarding_user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "committee", id: committeeID, field: "forwarding_user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ID(committeeID int) *ValueInt {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "committee", id: committeeID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ManagerIDs(committeeID int) *ValueIntSlice {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "manager_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "manager_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_MeetingIDs(committeeID int) *ValueIntSlice {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "meeting_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_Name(committeeID int) *ValueString {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "committee", id: committeeID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_OrganizationID(committeeID int) *ValueInt {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "organization_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "committee", id: committeeID, field: "organization_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_OrganizationTagIDs(committeeID int) *ValueIntSlice {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "organization_tag_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "organization_tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ReceiveForwardingsFromCommitteeIDs(committeeID int) *ValueIntSlice {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "receive_forwardings_from_committee_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "receive_forwardings_from_committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_UserIDs(committeeID int) *ValueIntSlice {
	key := dskey.Key{Collection: "committee", ID: committeeID, Field: "user_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_AdminGroupForMeetingID(groupID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "admin_group_for_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "admin_group_for_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_DefaultGroupForMeetingID(groupID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "default_group_for_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "default_group_for_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_ID(groupID int) *ValueInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "group", id: groupID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_MediafileAccessGroupIDs(groupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "mediafile_access_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "mediafile_access_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_MediafileInheritedAccessGroupIDs(groupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "mediafile_inherited_access_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "mediafile_inherited_access_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_MeetingID(groupID int) *ValueInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "group", id: groupID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_MeetingUserIDs(groupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "meeting_user_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "meeting_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_Name(groupID int) *ValueString {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "group", id: groupID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_Permissions(groupID int) *ValueStringSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "permissions"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "group", id: groupID, field: "permissions"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_PollIDs(groupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "poll_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_ReadChatGroupIDs(groupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "read_chat_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "read_chat_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_ReadCommentSectionIDs(groupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "read_comment_section_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "read_comment_section_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_UsedAsAssignmentPollDefaultID(groupID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "used_as_assignment_poll_default_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "used_as_assignment_poll_default_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_UsedAsMotionPollDefaultID(groupID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "used_as_motion_poll_default_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "used_as_motion_poll_default_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_UsedAsPollDefaultID(groupID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "used_as_poll_default_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "used_as_poll_default_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_UsedAsTopicPollDefaultID(groupID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "used_as_topic_poll_default_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "used_as_topic_poll_default_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_Weight(groupID int) *ValueInt {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "group", id: groupID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_WriteChatGroupIDs(groupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "write_chat_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "write_chat_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_WriteCommentSectionIDs(groupID int) *ValueIntSlice {
	key := dskey.Key{Collection: "group", ID: groupID, Field: "write_comment_section_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "write_comment_section_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_Closed(listOfSpeakersID int) *ValueBool {
	key := dskey.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "closed"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "closed"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ContentObjectID(listOfSpeakersID int) *ValueString {
	key := dskey.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "content_object_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "content_object_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ID(listOfSpeakersID int) *ValueInt {
	key := dskey.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_MeetingID(listOfSpeakersID int) *ValueInt {
	key := dskey.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ProjectionIDs(listOfSpeakersID int) *ValueIntSlice {
	key := dskey.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_SequentialNumber(listOfSpeakersID int) *ValueInt {
	key := dskey.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_SpeakerIDs(listOfSpeakersID int) *ValueIntSlice {
	key := dskey.Key{Collection: "list_of_speakers", ID: listOfSpeakersID, Field: "speaker_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "speaker_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_AccessGroupIDs(mediafileID int) *ValueIntSlice {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "access_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "access_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_AttachmentIDs(mediafileID int) *ValueStringSlice {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "attachment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "attachment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ChildIDs(mediafileID int) *ValueIntSlice {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "child_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "child_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_CreateTimestamp(mediafileID int) *ValueInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "create_timestamp"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "mediafile", id: mediafileID, field: "create_timestamp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Filename(mediafileID int) *ValueString {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "filename"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "filename"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Filesize(mediafileID int) *ValueInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "filesize"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "mediafile", id: mediafileID, field: "filesize"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ID(mediafileID int) *ValueInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "mediafile", id: mediafileID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_InheritedAccessGroupIDs(mediafileID int) *ValueIntSlice {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "inherited_access_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "inherited_access_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_IsDirectory(mediafileID int) *ValueBool {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "is_directory"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "mediafile", id: mediafileID, field: "is_directory"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_IsPublic(mediafileID int) *ValueBool {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "is_public"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "mediafile", id: mediafileID, field: "is_public", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ListOfSpeakersID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "list_of_speakers_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "list_of_speakers_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Mimetype(mediafileID int) *ValueString {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "mimetype"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "mimetype"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_OwnerID(mediafileID int) *ValueString {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "owner_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "owner_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ParentID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "parent_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "parent_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_PdfInformation(mediafileID int) *ValueJSON {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "pdf_information"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "mediafile", id: mediafileID, field: "pdf_information"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ProjectionIDs(mediafileID int) *ValueIntSlice {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Title(mediafileID int) *ValueString {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Token(mediafileID int) *ValueString {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "token"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "token"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontBoldInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_bold_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_bold_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontBoldItalicInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_bold_italic_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_bold_italic_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontChyronSpeakerNameInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_chyron_speaker_name_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_chyron_speaker_name_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontItalicInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_italic_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_italic_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontMonospaceInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_monospace_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_monospace_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontProjectorH1InMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_projector_h1_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_projector_h1_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontProjectorH2InMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_projector_h2_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_projector_h2_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontRegularInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_font_regular_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_regular_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfBallotPaperInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_pdf_ballot_paper_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_ballot_paper_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfFooterLInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_pdf_footer_l_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_footer_l_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfFooterRInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_pdf_footer_r_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_footer_r_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfHeaderLInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_pdf_header_l_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_header_l_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfHeaderRInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_pdf_header_r_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_header_r_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoProjectorHeaderInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_projector_header_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_projector_header_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoProjectorMainInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_projector_main_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_projector_main_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoWebHeaderInMeetingID(mediafileID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "mediafile", ID: mediafileID, Field: "used_as_logo_web_header_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_web_header_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_AboutMe(meetingUserID int) *ValueString {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "about_me"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "about_me"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_AssignmentCandidateIDs(meetingUserID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "assignment_candidate_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "assignment_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_ChatMessageIDs(meetingUserID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "chat_message_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "chat_message_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_Comment(meetingUserID int) *ValueString {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "comment"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "comment"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_GroupIDs(meetingUserID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_ID(meetingUserID int) *ValueInt {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meetingUser", id: meetingUserID, field: "id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_MeetingID(meetingUserID int) *ValueInt {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meetingUser", id: meetingUserID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_MotionSubmitterIDs(meetingUserID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "motion_submitter_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "motion_submitter_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_Number(meetingUserID int) *ValueString {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_PersonalNoteIDs(meetingUserID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "personal_note_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "personal_note_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_SpeakerIDs(meetingUserID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "speaker_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "speaker_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_StructureLevel(meetingUserID int) *ValueString {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "structure_level"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "structure_level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_SupportedMotionIDs(meetingUserID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "supported_motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "supported_motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_UserID(meetingUserID int) *ValueInt {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meetingUser", id: meetingUserID, field: "user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_VoteDelegatedToID(meetingUserID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "vote_delegated_to_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meetingUser", id: meetingUserID, field: "vote_delegated_to_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_VoteDelegationsFromIDs(meetingUserID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "vote_delegations_from_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "vote_delegations_from_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_VoteWeight(meetingUserID int) *ValueString {
	key := dskey.Key{Collection: "meeting_user", ID: meetingUserID, Field: "vote_weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "vote_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AdminGroupID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "admin_group_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "admin_group_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaEnableNumbering(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "agenda_enable_numbering"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "agenda_enable_numbering"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaItemCreation(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "agenda_item_creation"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "agenda_item_creation"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaItemIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "agenda_item_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "agenda_item_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaNewItemsDefaultVisibility(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "agenda_new_items_default_visibility"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "agenda_new_items_default_visibility"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaNumberPrefix(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "agenda_number_prefix"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "agenda_number_prefix"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaNumeralSystem(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "agenda_numeral_system"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "agenda_numeral_system"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaShowInternalItemsOnProjector(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "agenda_show_internal_items_on_projector"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "agenda_show_internal_items_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaShowSubtitles(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "agenda_show_subtitles"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "agenda_show_subtitles"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AllProjectionIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "all_projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "all_projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseEnable(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "applause_enable"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "applause_enable"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseMaxAmount(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "applause_max_amount"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "applause_max_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseMinAmount(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "applause_min_amount"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "applause_min_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseParticleImageUrl(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "applause_particle_image_url"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "applause_particle_image_url"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseShowLevel(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "applause_show_level"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "applause_show_level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseTimeout(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "applause_timeout"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "applause_timeout"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseType(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "applause_type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "applause_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentCandidateIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_candidate_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "assignment_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "assignment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollAddCandidatesToListOfSpeakers(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_add_candidates_to_list_of_speakers"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_add_candidates_to_list_of_speakers"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollBallotPaperNumber(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_ballot_paper_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_ballot_paper_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollBallotPaperSelection(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_ballot_paper_selection"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_ballot_paper_selection"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultBackend(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_backend"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_backend"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultMethod(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_method"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_method"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_onehundred_percent_base"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_onehundred_percent_base"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultType(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_default_type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollEnableMaxVotesPerOption(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_enable_max_votes_per_option"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_enable_max_votes_per_option"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollSortPollResultByVotes(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignment_poll_sort_poll_result_by_votes"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_sort_poll_result_by_votes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentsExportPreamble(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignments_export_preamble"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignments_export_preamble"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentsExportTitle(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "assignments_export_title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignments_export_title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ChatGroupIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "chat_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "chat_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ChatMessageIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "chat_message_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "chat_message_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_CommitteeID(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "committee_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "committee_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceAutoConnect(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_auto_connect"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_auto_connect"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceAutoConnectNextSpeakers(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_auto_connect_next_speakers"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "conference_auto_connect_next_speakers"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceEnableHelpdesk(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_enable_helpdesk"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_enable_helpdesk"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceLosRestriction(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_los_restriction"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_los_restriction"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceOpenMicrophone(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_open_microphone"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_open_microphone"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceOpenVideo(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_open_video"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_open_video"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceShow(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_show"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_show"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceStreamPosterUrl(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_stream_poster_url"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "conference_stream_poster_url"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceStreamUrl(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "conference_stream_url"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "conference_stream_url"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_CustomTranslations(meetingID int) *ValueJSON {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "custom_translations"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "meeting", id: meetingID, field: "custom_translations"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultGroupID(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_group_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "default_group_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultMeetingForCommitteeID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_meeting_for_committee_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "default_meeting_for_committee_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorAgendaItemIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_agenda_item_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_agenda_item_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorAmendmentIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_amendment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_amendment_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorAssignmentIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_assignment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_assignment_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorAssignmentPollIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_assignment_poll_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_assignment_poll_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_countdown_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_countdown_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorCurrentListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_current_list_of_speakers_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_current_list_of_speakers_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_list_of_speakers_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_list_of_speakers_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMediafileIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_mediafile_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_mediafile_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMessageIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_message_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_message_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMotionBlockIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_motion_block_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_motion_block_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMotionIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_motion_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMotionPollIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_motion_poll_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_motion_poll_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorPollIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_poll_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_poll_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorTopicIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "default_projector_topic_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_topic_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_Description(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "description"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_EnableAnonymous(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "enable_anonymous"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "enable_anonymous"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_EndTime(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "end_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "end_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportCsvEncoding(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_csv_encoding"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "export_csv_encoding"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportCsvSeparator(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_csv_separator"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "export_csv_separator"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfFontsize(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_fontsize"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_fontsize"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfLineHeight(meetingID int) *ValueFloat {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_line_height"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueFloat)
	}
	v := &ValueFloat{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_line_height"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginBottom(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_page_margin_bottom"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_page_margin_bottom"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginLeft(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_page_margin_left"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_page_margin_left"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginRight(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_page_margin_right"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_page_margin_right"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginTop(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_page_margin_top"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_page_margin_top"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPagenumberAlignment(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_pagenumber_alignment"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_pagenumber_alignment"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPagesize(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "export_pdf_pagesize"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_pagesize"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontBoldID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "font_bold_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_bold_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontBoldItalicID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "font_bold_italic_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_bold_italic_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontChyronSpeakerNameID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "font_chyron_speaker_name_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_chyron_speaker_name_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontItalicID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "font_italic_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_italic_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontMonospaceID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "font_monospace_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_monospace_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontProjectorH1ID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "font_projector_h1_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_projector_h1_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontProjectorH2ID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "font_projector_h2_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_projector_h2_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontRegularID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "font_regular_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_regular_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ForwardedMotionIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "forwarded_motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "forwarded_motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_GroupIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ID(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ImportedAt(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "imported_at"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "imported_at"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_IsActiveInOrganizationID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "is_active_in_organization_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "is_active_in_organization_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_IsArchivedInOrganizationID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "is_archived_in_organization_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "is_archived_in_organization_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_JitsiDomain(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "jitsi_domain"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "jitsi_domain"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_JitsiRoomName(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "jitsi_room_name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "jitsi_room_name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_JitsiRoomPassword(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "jitsi_room_password"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "jitsi_room_password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_Language(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "language"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "language"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersAmountLastOnProjector(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_amount_last_on_projector"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_amount_last_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersAmountNextOnProjector(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_amount_next_on_projector"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_amount_next_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCanSetContributionSelf(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_can_set_contribution_self"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_can_set_contribution_self"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCountdownID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_countdown_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_countdown_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCoupleCountdown(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_couple_countdown"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_couple_countdown"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_enable_point_of_order_speakers"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_enable_point_of_order_speakers"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersEnableProContraSpeech(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_enable_pro_contra_speech"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_enable_pro_contra_speech"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersInitiallyClosed(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_initially_closed"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_initially_closed"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersPresentUsersOnly(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_present_users_only"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_present_users_only"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_show_amount_of_speakers_on_slide"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_show_amount_of_speakers_on_slide"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersShowFirstContribution(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_show_first_contribution"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_show_first_contribution"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersSpeakerNoteForEveryone(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "list_of_speakers_speaker_note_for_everyone"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_speaker_note_for_everyone"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_Location(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "location"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "location"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfBallotPaperID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "logo_pdf_ballot_paper_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_ballot_paper_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfFooterLID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "logo_pdf_footer_l_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_footer_l_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfFooterRID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "logo_pdf_footer_r_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_footer_r_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfHeaderLID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "logo_pdf_header_l_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_header_l_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfHeaderRID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "logo_pdf_header_r_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_header_r_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoProjectorHeaderID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "logo_projector_header_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_projector_header_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoProjectorMainID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "logo_projector_main_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_projector_main_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoWebHeaderID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "logo_web_header_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_web_header_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MediafileIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "mediafile_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "mediafile_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MeetingUserIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "meeting_user_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "meeting_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionBlockIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_block_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_block_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionCategoryIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_category_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_category_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionChangeRecommendationIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_change_recommendation_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_change_recommendation_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionCommentIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_comment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_comment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionCommentSectionIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_comment_section_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_comment_section_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollBallotPaperNumber(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_ballot_paper_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_ballot_paper_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollBallotPaperSelection(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_ballot_paper_selection"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_ballot_paper_selection"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultBackend(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_default_backend"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_default_backend"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_default_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_default_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_default_onehundred_percent_base"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_default_onehundred_percent_base"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultType(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_poll_default_type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_default_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionStateIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_state_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_state_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionStatuteParagraphIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_statute_paragraph_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_statute_paragraph_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionSubmitterIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_submitter_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_submitter_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionWorkflowIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motion_workflow_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_workflow_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsEnabled(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_enabled"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_enabled"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsInMainList(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_in_main_list"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_in_main_list"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsMultipleParagraphs(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_multiple_paragraphs"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_multiple_paragraphs"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsOfAmendments(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_of_amendments"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_of_amendments"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsPrefix(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_prefix"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_prefix"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsTextMode(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_amendments_text_mode"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_text_mode"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsBlockSlideColumns(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_block_slide_columns"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_block_slide_columns"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultAmendmentWorkflowID(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_amendment_workflow_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_amendment_workflow_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultLineNumbering(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_line_numbering"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_line_numbering"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultSorting(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_sorting"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_sorting"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultStatuteAmendmentWorkflowID(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_statute_amendment_workflow_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_statute_amendment_workflow_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultWorkflowID(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_default_workflow_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_workflow_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableReasonOnProjector(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_enable_reason_on_projector"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_enable_reason_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableRecommendationOnProjector(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_enable_recommendation_on_projector"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_enable_recommendation_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableSideboxOnProjector(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_enable_sidebox_on_projector"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_enable_sidebox_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableTextOnProjector(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_enable_text_on_projector"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_enable_text_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportFollowRecommendation(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_export_follow_recommendation"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_export_follow_recommendation"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportPreamble(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_export_preamble"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_export_preamble"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportSubmitterRecommendation(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_export_submitter_recommendation"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_export_submitter_recommendation"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportTitle(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_export_title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_export_title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsLineLength(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_line_length"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_line_length"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberMinDigits(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_number_min_digits"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_number_min_digits"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberType(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_number_type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_number_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberWithBlank(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_number_with_blank"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_number_with_blank"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsPreamble(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_preamble"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_preamble"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsReasonRequired(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_reason_required"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_reason_required"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsRecommendationTextMode(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_recommendation_text_mode"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_recommendation_text_mode"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsRecommendationsBy(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_recommendations_by"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_recommendations_by"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsShowReferringMotions(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_show_referring_motions"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_show_referring_motions"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsShowSequentialNumber(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_show_sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_show_sequential_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsStatuteRecommendationsBy(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_statute_recommendations_by"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_statute_recommendations_by"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsStatutesEnabled(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_statutes_enabled"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_statutes_enabled"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsSupportersMinAmount(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "motions_supporters_min_amount"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_supporters_min_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_Name(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_OptionIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "option_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "option_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_OrganizationTagIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "organization_tag_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "organization_tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PersonalNoteIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "personal_note_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "personal_note_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollBallotPaperNumber(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_ballot_paper_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "poll_ballot_paper_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollBallotPaperSelection(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_ballot_paper_selection"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_ballot_paper_selection"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollCandidateIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_candidate_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "poll_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollCandidateListIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_candidate_list_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "poll_candidate_list_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollCountdownID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_countdown_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "poll_countdown_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollCoupleCountdown(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_couple_countdown"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "poll_couple_countdown"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultBackend(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_backend"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_backend"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultMethod(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_method"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_method"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_onehundred_percent_base"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_onehundred_percent_base"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultType(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_default_type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollSortPollResultByVotes(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "poll_sort_poll_result_by_votes"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "poll_sort_poll_result_by_votes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PresentUserIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "present_user_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "present_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectionIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownDefaultTime(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "projector_countdown_default_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "projector_countdown_default_time", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "projector_countdown_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "projector_countdown_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownWarningTime(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "projector_countdown_warning_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "projector_countdown_warning_time", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "projector_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "projector_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorMessageIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "projector_message_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "projector_message_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ReferenceProjectorID(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "reference_projector_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "reference_projector_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_SpeakerIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "speaker_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "speaker_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_StartTime(meetingID int) *ValueInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "start_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "start_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_TagIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "tag_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_TemplateForOrganizationID(meetingID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "template_for_organization_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "template_for_organization_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_TopicIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "topic_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "topic_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_TopicPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "topic_poll_default_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "topic_poll_default_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UserIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "user_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersAllowSelfSetPresent(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_allow_self_set_present"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "users_allow_self_set_present"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailBody(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_email_body"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_email_body"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailReplyto(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_email_replyto"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_email_replyto"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailSender(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_email_sender"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_email_sender"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailSubject(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_email_subject"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_email_subject"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEnablePresenceView(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_enable_presence_view"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "users_enable_presence_view"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEnableVoteDelegations(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_enable_vote_delegations"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "users_enable_vote_delegations"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEnableVoteWeight(meetingID int) *ValueBool {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_enable_vote_weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "users_enable_vote_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWelcometext(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_welcometext"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_welcometext"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWelcometitle(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_welcometitle"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_welcometitle"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanEncryption(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_wlan_encryption"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_wlan_encryption"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanPassword(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_wlan_password"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_wlan_password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanSsid(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "users_pdf_wlan_ssid"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_wlan_ssid"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_VoteIDs(meetingID int) *ValueIntSlice {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "vote_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "vote_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_WelcomeText(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "welcome_text"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "welcome_text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_WelcomeTitle(meetingID int) *ValueString {
	key := dskey.Key{Collection: "meeting", ID: meetingID, Field: "welcome_title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "welcome_title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_AgendaItemID(motionBlockID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "agenda_item_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "agenda_item_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_ID(motionBlockID int) *ValueInt {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_Internal(motionBlockID int) *ValueBool {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "internal"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionBlock", id: motionBlockID, field: "internal"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_ListOfSpeakersID(motionBlockID int) *ValueInt {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "list_of_speakers_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_MeetingID(motionBlockID int) *ValueInt {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_MotionIDs(motionBlockID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionBlock", id: motionBlockID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_ProjectionIDs(motionBlockID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionBlock", id: motionBlockID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_SequentialNumber(motionBlockID int) *ValueInt {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_Title(motionBlockID int) *ValueString {
	key := dskey.Key{Collection: "motion_block", ID: motionBlockID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionBlock", id: motionBlockID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_ChildIDs(motionCategoryID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "child_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "child_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_ID(motionCategoryID int) *ValueInt {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_Level(motionCategoryID int) *ValueInt {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "level"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_MeetingID(motionCategoryID int) *ValueInt {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_MotionIDs(motionCategoryID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_Name(motionCategoryID int) *ValueString {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_ParentID(motionCategoryID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "parent_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "parent_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_Prefix(motionCategoryID int) *ValueString {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "prefix"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "prefix"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_SequentialNumber(motionCategoryID int) *ValueInt {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_Weight(motionCategoryID int) *ValueInt {
	key := dskey.Key{Collection: "motion_category", ID: motionCategoryID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_CreationTime(motionChangeRecommendationID int) *ValueInt {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "creation_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "creation_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_ID(motionChangeRecommendationID int) *ValueInt {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Internal(motionChangeRecommendationID int) *ValueBool {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "internal"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "internal"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_LineFrom(motionChangeRecommendationID int) *ValueInt {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "line_from"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "line_from"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_LineTo(motionChangeRecommendationID int) *ValueInt {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "line_to"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "line_to"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_MeetingID(motionChangeRecommendationID int) *ValueInt {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_MotionID(motionChangeRecommendationID int) *ValueInt {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "motion_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "motion_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_OtherDescription(motionChangeRecommendationID int) *ValueString {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "other_description"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "other_description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Rejected(motionChangeRecommendationID int) *ValueBool {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "rejected"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "rejected"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Text(motionChangeRecommendationID int) *ValueString {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "text"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Type(motionChangeRecommendationID int) *ValueString {
	key := dskey.Key{Collection: "motion_change_recommendation", ID: motionChangeRecommendationID, Field: "type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_CommentIDs(motionCommentSectionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "comment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "comment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_ID(motionCommentSectionID int) *ValueInt {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_MeetingID(motionCommentSectionID int) *ValueInt {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_Name(motionCommentSectionID int) *ValueString {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_ReadGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "read_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "read_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_SequentialNumber(motionCommentSectionID int) *ValueInt {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_SubmitterCanWrite(motionCommentSectionID int) *ValueBool {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "submitter_can_write"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "submitter_can_write"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_Weight(motionCommentSectionID int) *ValueInt {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_WriteGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_comment_section", ID: motionCommentSectionID, Field: "write_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "write_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_Comment(motionCommentID int) *ValueString {
	key := dskey.Key{Collection: "motion_comment", ID: motionCommentID, Field: "comment"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionComment", id: motionCommentID, field: "comment"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_ID(motionCommentID int) *ValueInt {
	key := dskey.Key{Collection: "motion_comment", ID: motionCommentID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionComment", id: motionCommentID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_MeetingID(motionCommentID int) *ValueInt {
	key := dskey.Key{Collection: "motion_comment", ID: motionCommentID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionComment", id: motionCommentID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_MotionID(motionCommentID int) *ValueInt {
	key := dskey.Key{Collection: "motion_comment", ID: motionCommentID, Field: "motion_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionComment", id: motionCommentID, field: "motion_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_SectionID(motionCommentID int) *ValueInt {
	key := dskey.Key{Collection: "motion_comment", ID: motionCommentID, Field: "section_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionComment", id: motionCommentID, field: "section_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_AllowCreatePoll(motionStateID int) *ValueBool {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "allow_create_poll"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "allow_create_poll"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_AllowMotionForwarding(motionStateID int) *ValueBool {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "allow_motion_forwarding"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "allow_motion_forwarding"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_AllowSubmitterEdit(motionStateID int) *ValueBool {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "allow_submitter_edit"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "allow_submitter_edit"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_AllowSupport(motionStateID int) *ValueBool {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "allow_support"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "allow_support"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_CssClass(motionStateID int) *ValueString {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "css_class"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionState", id: motionStateID, field: "css_class", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_FirstStateOfWorkflowID(motionStateID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "first_state_of_workflow_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionState", id: motionStateID, field: "first_state_of_workflow_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_ID(motionStateID int) *ValueInt {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionState", id: motionStateID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_MeetingID(motionStateID int) *ValueInt {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionState", id: motionStateID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_MergeAmendmentIntoFinal(motionStateID int) *ValueString {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "merge_amendment_into_final"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionState", id: motionStateID, field: "merge_amendment_into_final"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_MotionIDs(motionStateID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_MotionRecommendationIDs(motionStateID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "motion_recommendation_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "motion_recommendation_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_Name(motionStateID int) *ValueString {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionState", id: motionStateID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_NextStateIDs(motionStateID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "next_state_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "next_state_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_PreviousStateIDs(motionStateID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "previous_state_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "previous_state_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_RecommendationLabel(motionStateID int) *ValueString {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "recommendation_label"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionState", id: motionStateID, field: "recommendation_label"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_Restrictions(motionStateID int) *ValueStringSlice {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "restrictions"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "motionState", id: motionStateID, field: "restrictions"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_SetCreatedTimestamp(motionStateID int) *ValueBool {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "set_created_timestamp"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "set_created_timestamp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_SetNumber(motionStateID int) *ValueBool {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "set_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "set_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_ShowRecommendationExtensionField(motionStateID int) *ValueBool {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "show_recommendation_extension_field"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "show_recommendation_extension_field"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_ShowStateExtensionField(motionStateID int) *ValueBool {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "show_state_extension_field"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "show_state_extension_field"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_SubmitterWithdrawBackIDs(motionStateID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "submitter_withdraw_back_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "submitter_withdraw_back_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_SubmitterWithdrawStateID(motionStateID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "submitter_withdraw_state_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionState", id: motionStateID, field: "submitter_withdraw_state_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_Weight(motionStateID int) *ValueInt {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionState", id: motionStateID, field: "weight", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_WorkflowID(motionStateID int) *ValueInt {
	key := dskey.Key{Collection: "motion_state", ID: motionStateID, Field: "workflow_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionState", id: motionStateID, field: "workflow_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_ID(motionStatuteParagraphID int) *ValueInt {
	key := dskey.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_MeetingID(motionStatuteParagraphID int) *ValueInt {
	key := dskey.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_MotionIDs(motionStatuteParagraphID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_SequentialNumber(motionStatuteParagraphID int) *ValueInt {
	key := dskey.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Text(motionStatuteParagraphID int) *ValueString {
	key := dskey.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "text"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Title(motionStatuteParagraphID int) *ValueString {
	key := dskey.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Weight(motionStatuteParagraphID int) *ValueInt {
	key := dskey.Key{Collection: "motion_statute_paragraph", ID: motionStatuteParagraphID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_ID(motionSubmitterID int) *ValueInt {
	key := dskey.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_MeetingID(motionSubmitterID int) *ValueInt {
	key := dskey.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_MeetingUserID(motionSubmitterID int) *ValueInt {
	key := dskey.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "meeting_user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "meeting_user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_MotionID(motionSubmitterID int) *ValueInt {
	key := dskey.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "motion_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "motion_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_Weight(motionSubmitterID int) *ValueInt {
	key := dskey.Key{Collection: "motion_submitter", ID: motionSubmitterID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "default_amendment_workflow_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "default_amendment_workflow_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultStatuteAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "default_statute_amendment_workflow_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "default_statute_amendment_workflow_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "default_workflow_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "default_workflow_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_FirstStateID(motionWorkflowID int) *ValueInt {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "first_state_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "first_state_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_ID(motionWorkflowID int) *ValueInt {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_MeetingID(motionWorkflowID int) *ValueInt {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_Name(motionWorkflowID int) *ValueString {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_SequentialNumber(motionWorkflowID int) *ValueInt {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_StateIDs(motionWorkflowID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion_workflow", ID: motionWorkflowID, Field: "state_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "state_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AgendaItemID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "agenda_item_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "agenda_item_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AllDerivedMotionIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "all_derived_motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "all_derived_motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AllOriginIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "all_origin_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "all_origin_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AmendmentIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "amendment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "amendment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AmendmentParagraphs(motionID int) *ValueJSON {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "amendment_paragraphs"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "motion", id: motionID, field: "amendment_paragraphs"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AttachmentIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "attachment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "attachment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_BlockID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "block_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "block_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_CategoryID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "category_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "category_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_CategoryWeight(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "category_weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "category_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ChangeRecommendationIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "change_recommendation_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "change_recommendation_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_CommentIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "comment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "comment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Created(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "created"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "created"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_DerivedMotionIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "derived_motion_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "derived_motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Forwarded(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "forwarded"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "forwarded"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ID(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_LastModified(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "last_modified"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "last_modified"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_LeadMotionID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "lead_motion_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "lead_motion_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ListOfSpeakersID(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "list_of_speakers_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_MeetingID(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ModifiedFinalVersion(motionID int) *ValueString {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "modified_final_version"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "modified_final_version"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Number(motionID int) *ValueString {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_NumberValue(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "number_value"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "number_value"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_OptionIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "option_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "option_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_OriginID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "origin_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "origin_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_OriginMeetingID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "origin_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "origin_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_PersonalNoteIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "personal_note_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "personal_note_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_PollIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "poll_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ProjectionIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Reason(motionID int) *ValueString {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "reason"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "reason"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_RecommendationExtension(motionID int) *ValueString {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "recommendation_extension"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "recommendation_extension"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_RecommendationExtensionReferenceIDs(motionID int) *ValueStringSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "recommendation_extension_reference_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "motion", id: motionID, field: "recommendation_extension_reference_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_RecommendationID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "recommendation_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "recommendation_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ReferencedInMotionRecommendationExtensionIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "referenced_in_motion_recommendation_extension_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "referenced_in_motion_recommendation_extension_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ReferencedInMotionStateExtensionIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "referenced_in_motion_state_extension_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "referenced_in_motion_state_extension_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SequentialNumber(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SortChildIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "sort_child_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "sort_child_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SortParentID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "sort_parent_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "sort_parent_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SortWeight(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "sort_weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "sort_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StartLineNumber(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "start_line_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "start_line_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StateExtension(motionID int) *ValueString {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "state_extension"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "state_extension"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StateExtensionReferenceIDs(motionID int) *ValueStringSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "state_extension_reference_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "motion", id: motionID, field: "state_extension_reference_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StateID(motionID int) *ValueInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "state_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "state_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StatuteParagraphID(motionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "statute_paragraph_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "statute_paragraph_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SubmitterIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "submitter_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "submitter_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SupporterMeetingUserIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "supporter_meeting_user_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "supporter_meeting_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_TagIDs(motionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "tag_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Text(motionID int) *ValueString {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "text"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Title(motionID int) *ValueString {
	key := dskey.Key{Collection: "motion", ID: motionID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_Abstain(optionID int) *ValueString {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "abstain"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "option", id: optionID, field: "abstain"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_ContentObjectID(optionID int) *ValueMaybeString {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "content_object_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeString)
	}
	v := &ValueMaybeString{fetch: r, collection: "option", id: optionID, field: "content_object_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_ID(optionID int) *ValueInt {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "option", id: optionID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_MeetingID(optionID int) *ValueInt {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "option", id: optionID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_No(optionID int) *ValueString {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "no"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "option", id: optionID, field: "no"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_PollID(optionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "poll_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "option", id: optionID, field: "poll_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_Text(optionID int) *ValueString {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "text"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "option", id: optionID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_UsedAsGlobalOptionInPollID(optionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "used_as_global_option_in_poll_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "option", id: optionID, field: "used_as_global_option_in_poll_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_VoteIDs(optionID int) *ValueIntSlice {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "vote_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "option", id: optionID, field: "vote_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_Weight(optionID int) *ValueInt {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "option", id: optionID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_Yes(optionID int) *ValueString {
	key := dskey.Key{Collection: "option", ID: optionID, Field: "yes"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "option", id: optionID, field: "yes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_Color(organizationTagID int) *ValueString {
	key := dskey.Key{Collection: "organization_tag", ID: organizationTagID, Field: "color"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organizationTag", id: organizationTagID, field: "color", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_ID(organizationTagID int) *ValueInt {
	key := dskey.Key{Collection: "organization_tag", ID: organizationTagID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organizationTag", id: organizationTagID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_Name(organizationTagID int) *ValueString {
	key := dskey.Key{Collection: "organization_tag", ID: organizationTagID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organizationTag", id: organizationTagID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_OrganizationID(organizationTagID int) *ValueInt {
	key := dskey.Key{Collection: "organization_tag", ID: organizationTagID, Field: "organization_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organizationTag", id: organizationTagID, field: "organization_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_TaggedIDs(organizationTagID int) *ValueStringSlice {
	key := dskey.Key{Collection: "organization_tag", ID: organizationTagID, Field: "tagged_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "organizationTag", id: organizationTagID, field: "tagged_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ActiveMeetingIDs(organizationID int) *ValueIntSlice {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "active_meeting_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "active_meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ArchivedMeetingIDs(organizationID int) *ValueIntSlice {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "archived_meeting_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "archived_meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_CommitteeIDs(organizationID int) *ValueIntSlice {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "committee_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_DefaultLanguage(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "default_language"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "default_language", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_Description(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "description"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_EnableChat(organizationID int) *ValueBool {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "enable_chat"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "organization", id: organizationID, field: "enable_chat"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_EnableElectronicVoting(organizationID int) *ValueBool {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "enable_electronic_voting"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "organization", id: organizationID, field: "enable_electronic_voting"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ID(organizationID int) *ValueInt {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organization", id: organizationID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_LegalNotice(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "legal_notice"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "legal_notice"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_LimitOfMeetings(organizationID int) *ValueInt {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "limit_of_meetings"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organization", id: organizationID, field: "limit_of_meetings"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_LimitOfUsers(organizationID int) *ValueInt {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "limit_of_users"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organization", id: organizationID, field: "limit_of_users"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_LoginText(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "login_text"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "login_text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_MediafileIDs(organizationID int) *ValueIntSlice {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "mediafile_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "mediafile_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_Name(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_OrganizationTagIDs(organizationID int) *ValueIntSlice {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "organization_tag_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "organization_tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_PrivacyPolicy(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "privacy_policy"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "privacy_policy"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ResetPasswordVerboseErrors(organizationID int) *ValueBool {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "reset_password_verbose_errors"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "organization", id: organizationID, field: "reset_password_verbose_errors"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlAttrMapping(organizationID int) *ValueJSON {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "saml_attr_mapping"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "organization", id: organizationID, field: "saml_attr_mapping"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlEnabled(organizationID int) *ValueBool {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "saml_enabled"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "organization", id: organizationID, field: "saml_enabled"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlLoginButtonText(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "saml_login_button_text"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "saml_login_button_text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlMetadataIDp(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "saml_metadata_idp"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "saml_metadata_idp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlMetadataSp(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "saml_metadata_sp"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "saml_metadata_sp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlPrivateKey(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "saml_private_key"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "saml_private_key"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_TemplateMeetingIDs(organizationID int) *ValueIntSlice {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "template_meeting_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "template_meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ThemeID(organizationID int) *ValueInt {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "theme_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organization", id: organizationID, field: "theme_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ThemeIDs(organizationID int) *ValueIntSlice {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "theme_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "theme_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_Url(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "url"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "url"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UserIDs(organizationID int) *ValueIntSlice {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "user_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UsersEmailBody(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "users_email_body"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "users_email_body"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UsersEmailReplyto(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "users_email_replyto"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "users_email_replyto"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UsersEmailSender(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "users_email_sender"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "users_email_sender"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UsersEmailSubject(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "users_email_subject"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "users_email_subject"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_VoteDecryptPublicMainKey(organizationID int) *ValueString {
	key := dskey.Key{Collection: "organization", ID: organizationID, Field: "vote_decrypt_public_main_key"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "vote_decrypt_public_main_key"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_ContentObjectID(personalNoteID int) *ValueMaybeString {
	key := dskey.Key{Collection: "personal_note", ID: personalNoteID, Field: "content_object_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeString)
	}
	v := &ValueMaybeString{fetch: r, collection: "personalNote", id: personalNoteID, field: "content_object_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_ID(personalNoteID int) *ValueInt {
	key := dskey.Key{Collection: "personal_note", ID: personalNoteID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "personalNote", id: personalNoteID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_MeetingID(personalNoteID int) *ValueInt {
	key := dskey.Key{Collection: "personal_note", ID: personalNoteID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "personalNote", id: personalNoteID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_MeetingUserID(personalNoteID int) *ValueInt {
	key := dskey.Key{Collection: "personal_note", ID: personalNoteID, Field: "meeting_user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "personalNote", id: personalNoteID, field: "meeting_user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_Note(personalNoteID int) *ValueString {
	key := dskey.Key{Collection: "personal_note", ID: personalNoteID, Field: "note"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "personalNote", id: personalNoteID, field: "note"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_Star(personalNoteID int) *ValueBool {
	key := dskey.Key{Collection: "personal_note", ID: personalNoteID, Field: "star"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "personalNote", id: personalNoteID, field: "star"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidateList_ID(pollCandidateListID int) *ValueInt {
	key := dskey.Key{Collection: "poll_candidate_list", ID: pollCandidateListID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidateList", id: pollCandidateListID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidateList_MeetingID(pollCandidateListID int) *ValueInt {
	key := dskey.Key{Collection: "poll_candidate_list", ID: pollCandidateListID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidateList", id: pollCandidateListID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidateList_OptionID(pollCandidateListID int) *ValueInt {
	key := dskey.Key{Collection: "poll_candidate_list", ID: pollCandidateListID, Field: "option_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidateList", id: pollCandidateListID, field: "option_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidateList_PollCandidateIDs(pollCandidateListID int) *ValueIntSlice {
	key := dskey.Key{Collection: "poll_candidate_list", ID: pollCandidateListID, Field: "poll_candidate_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "pollCandidateList", id: pollCandidateListID, field: "poll_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_ID(pollCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "poll_candidate", ID: pollCandidateID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_MeetingID(pollCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "poll_candidate", ID: pollCandidateID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_PollCandidateListID(pollCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "poll_candidate", ID: pollCandidateID, Field: "poll_candidate_list_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "poll_candidate_list_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_UserID(pollCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "poll_candidate", ID: pollCandidateID, Field: "user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_Weight(pollCandidateID int) *ValueInt {
	key := dskey.Key{Collection: "poll_candidate", ID: pollCandidateID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "weight", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Backend(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "backend"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "backend", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_ContentObjectID(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "content_object_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "content_object_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_CryptKey(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "crypt_key"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "crypt_key"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_CryptSignature(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "crypt_signature"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "crypt_signature"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Description(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "description"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_EntitledGroupIDs(pollID int) *ValueIntSlice {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "entitled_group_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "poll", id: pollID, field: "entitled_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_EntitledUsersAtStop(pollID int) *ValueJSON {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "entitled_users_at_stop"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "poll", id: pollID, field: "entitled_users_at_stop"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_GlobalAbstain(pollID int) *ValueBool {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "global_abstain"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "poll", id: pollID, field: "global_abstain"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_GlobalNo(pollID int) *ValueBool {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "global_no"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "poll", id: pollID, field: "global_no"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_GlobalOptionID(pollID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "global_option_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "poll", id: pollID, field: "global_option_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_GlobalYes(pollID int) *ValueBool {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "global_yes"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "poll", id: pollID, field: "global_yes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_ID(pollID int) *ValueInt {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_IsPseudoanonymized(pollID int) *ValueBool {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "is_pseudoanonymized"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "poll", id: pollID, field: "is_pseudoanonymized"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_MaxVotesAmount(pollID int) *ValueInt {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "max_votes_amount"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "max_votes_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_MaxVotesPerOption(pollID int) *ValueInt {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "max_votes_per_option"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "max_votes_per_option"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_MeetingID(pollID int) *ValueInt {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_MinVotesAmount(pollID int) *ValueInt {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "min_votes_amount"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "min_votes_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_OnehundredPercentBase(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "onehundred_percent_base"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "onehundred_percent_base", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_OptionIDs(pollID int) *ValueIntSlice {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "option_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "poll", id: pollID, field: "option_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Pollmethod(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "pollmethod"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "pollmethod", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_ProjectionIDs(pollID int) *ValueIntSlice {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "poll", id: pollID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_SequentialNumber(pollID int) *ValueInt {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_State(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "state"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "state"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Title(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Type(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "type", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_VoteCount(pollID int) *ValueInt {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "vote_count"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "vote_count"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_VotedIDs(pollID int) *ValueIntSlice {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "voted_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "poll", id: pollID, field: "voted_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_VotesRaw(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "votes_raw"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votes_raw"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_VotesSignature(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "votes_signature"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votes_signature"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Votescast(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "votescast"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votescast"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Votesinvalid(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "votesinvalid"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votesinvalid"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Votesvalid(pollID int) *ValueString {
	key := dskey.Key{Collection: "poll", ID: pollID, Field: "votesvalid"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votesvalid"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Content(projectionID int) *ValueJSON {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "content"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "projection", id: projectionID, field: "content"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_ContentObjectID(projectionID int) *ValueString {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "content_object_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projection", id: projectionID, field: "content_object_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_CurrentProjectorID(projectionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "current_projector_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projection", id: projectionID, field: "current_projector_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_HistoryProjectorID(projectionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "history_projector_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projection", id: projectionID, field: "history_projector_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_ID(projectionID int) *ValueInt {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projection", id: projectionID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_MeetingID(projectionID int) *ValueInt {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projection", id: projectionID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Options(projectionID int) *ValueJSON {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "options"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "projection", id: projectionID, field: "options"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_PreviewProjectorID(projectionID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "preview_projector_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projection", id: projectionID, field: "preview_projector_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Stable(projectionID int) *ValueBool {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "stable"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projection", id: projectionID, field: "stable"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Type(projectionID int) *ValueString {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "type"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projection", id: projectionID, field: "type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Weight(projectionID int) *ValueInt {
	key := dskey.Key{Collection: "projection", ID: projectionID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projection", id: projectionID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_CountdownTime(projectorCountdownID int) *ValueFloat {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "countdown_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueFloat)
	}
	v := &ValueFloat{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "countdown_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_DefaultTime(projectorCountdownID int) *ValueInt {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "default_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "default_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Description(projectorCountdownID int) *ValueString {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "description"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_ID(projectorCountdownID int) *ValueInt {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_MeetingID(projectorCountdownID int) *ValueInt {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_ProjectionIDs(projectorCountdownID int) *ValueIntSlice {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Running(projectorCountdownID int) *ValueBool {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "running"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "running"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Title(projectorCountdownID int) *ValueString {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "used_as_list_of_speakers_countdown_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "used_as_list_of_speakers_countdown_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_UsedAsPollCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector_countdown", ID: projectorCountdownID, Field: "used_as_poll_countdown_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "used_as_poll_countdown_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorMessage_ID(projectorMessageID int) *ValueInt {
	key := dskey.Key{Collection: "projector_message", ID: projectorMessageID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorMessage", id: projectorMessageID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorMessage_MeetingID(projectorMessageID int) *ValueInt {
	key := dskey.Key{Collection: "projector_message", ID: projectorMessageID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorMessage", id: projectorMessageID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorMessage_Message(projectorMessageID int) *ValueString {
	key := dskey.Key{Collection: "projector_message", ID: projectorMessageID, Field: "message"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projectorMessage", id: projectorMessageID, field: "message"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorMessage_ProjectionIDs(projectorMessageID int) *ValueIntSlice {
	key := dskey.Key{Collection: "projector_message", ID: projectorMessageID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projectorMessage", id: projectorMessageID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_AspectRatioDenominator(projectorID int) *ValueInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "aspect_ratio_denominator"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "aspect_ratio_denominator"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_AspectRatioNumerator(projectorID int) *ValueInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "aspect_ratio_numerator"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "aspect_ratio_numerator"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_BackgroundColor(projectorID int) *ValueString {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "background_color"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "background_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ChyronBackgroundColor(projectorID int) *ValueString {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "chyron_background_color"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "chyron_background_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ChyronFontColor(projectorID int) *ValueString {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "chyron_font_color"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "chyron_font_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Color(projectorID int) *ValueString {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "color"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_CurrentProjectionIDs(projectorID int) *ValueIntSlice {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "current_projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projector", id: projectorID, field: "current_projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_HeaderBackgroundColor(projectorID int) *ValueString {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "header_background_color"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "header_background_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_HeaderFontColor(projectorID int) *ValueString {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "header_font_color"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "header_font_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_HeaderH1Color(projectorID int) *ValueString {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "header_h1_color"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "header_h1_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_HistoryProjectionIDs(projectorID int) *ValueIntSlice {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "history_projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projector", id: projectorID, field: "history_projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ID(projectorID int) *ValueInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_IsInternal(projectorID int) *ValueBool {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "is_internal"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "is_internal"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_MeetingID(projectorID int) *ValueInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Name(projectorID int) *ValueString {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_PreviewProjectionIDs(projectorID int) *ValueIntSlice {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "preview_projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projector", id: projectorID, field: "preview_projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Scale(projectorID int) *ValueInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "scale"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "scale"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Scroll(projectorID int) *ValueInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "scroll"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "scroll"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_SequentialNumber(projectorID int) *ValueInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ShowClock(projectorID int) *ValueBool {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "show_clock"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "show_clock"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ShowHeaderFooter(projectorID int) *ValueBool {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "show_header_footer"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "show_header_footer"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ShowLogo(projectorID int) *ValueBool {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "show_logo"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "show_logo"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ShowTitle(projectorID int) *ValueBool {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "show_title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "show_title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAgendaItemInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_agenda_item_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_agenda_item_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAmendmentInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_amendment_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_amendment_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAssignmentInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_assignment_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_assignment_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAssignmentPollInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_assignment_poll_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_assignment_poll_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForCountdownInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_countdown_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_countdown_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_current_list_of_speakers_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_current_list_of_speakers_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForListOfSpeakersInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_list_of_speakers_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_list_of_speakers_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMediafileInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_mediafile_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_mediafile_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMessageInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_message_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_message_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionBlockInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_motion_block_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_motion_block_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_motion_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_motion_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionPollInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_motion_poll_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_motion_poll_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForPollInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_poll_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_poll_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForTopicInMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_default_projector_for_topic_in_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_topic_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsReferenceProjectorMeetingID(projectorID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "used_as_reference_projector_meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_reference_projector_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Width(projectorID int) *ValueInt {
	key := dskey.Key{Collection: "projector", ID: projectorID, Field: "width"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "width"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_BeginTime(speakerID int) *ValueInt {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "begin_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "begin_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_EndTime(speakerID int) *ValueInt {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "end_time"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "end_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_ID(speakerID int) *ValueInt {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_ListOfSpeakersID(speakerID int) *ValueInt {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "list_of_speakers_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_MeetingID(speakerID int) *ValueInt {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_MeetingUserID(speakerID int) *ValueInt {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "meeting_user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "meeting_user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_Note(speakerID int) *ValueString {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "note"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "speaker", id: speakerID, field: "note"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_PointOfOrder(speakerID int) *ValueBool {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "point_of_order"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "speaker", id: speakerID, field: "point_of_order"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_SpeechState(speakerID int) *ValueString {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "speech_state"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "speaker", id: speakerID, field: "speech_state"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_Weight(speakerID int) *ValueInt {
	key := dskey.Key{Collection: "speaker", ID: speakerID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Tag_ID(tagID int) *ValueInt {
	key := dskey.Key{Collection: "tag", ID: tagID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "tag", id: tagID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Tag_MeetingID(tagID int) *ValueInt {
	key := dskey.Key{Collection: "tag", ID: tagID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "tag", id: tagID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Tag_Name(tagID int) *ValueString {
	key := dskey.Key{Collection: "tag", ID: tagID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "tag", id: tagID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Tag_TaggedIDs(tagID int) *ValueStringSlice {
	key := dskey.Key{Collection: "tag", ID: tagID, Field: "tagged_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "tag", id: tagID, field: "tagged_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Abstain(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "abstain"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "abstain"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent100(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_100"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent200(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_200"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent300(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_300"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_300"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent400(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_400"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent50(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_50"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_50"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent500(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_500"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_500", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent600(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_600"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_600"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent700(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_700"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent800(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_800"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_800"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent900(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_900"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_900"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_AccentA100(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_a100"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_a100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_AccentA200(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_a200"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_a200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_AccentA400(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_a400"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_a400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_AccentA700(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "accent_a700"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_a700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Headbar(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "headbar"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "headbar"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_ID(themeID int) *ValueInt {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "theme", id: themeID, field: "id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Name(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_No(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "no"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "no"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_OrganizationID(themeID int) *ValueInt {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "organization_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "theme", id: themeID, field: "organization_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary100(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_100"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary200(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_200"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary300(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_300"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_300"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary400(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_400"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary50(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_50"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_50"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary500(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_500"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_500", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary600(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_600"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_600"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary700(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_700"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary800(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_800"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_800"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary900(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_900"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_900"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_PrimaryA100(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_a100"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_a100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_PrimaryA200(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_a200"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_a200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_PrimaryA400(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_a400"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_a400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_PrimaryA700(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "primary_a700"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_a700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_ThemeForOrganizationID(themeID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "theme_for_organization_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "theme", id: themeID, field: "theme_for_organization_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn100(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_100"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn200(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_200"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn300(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_300"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_300"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn400(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_400"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn50(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_50"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_50"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn500(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_500"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_500", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn600(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_600"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_600"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn700(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_700"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn800(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_800"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_800"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn900(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_900"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_900"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_WarnA100(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_a100"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_a100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_WarnA200(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_a200"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_a200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_WarnA400(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_a400"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_a400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_WarnA700(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "warn_a700"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_a700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Yes(themeID int) *ValueString {
	key := dskey.Key{Collection: "theme", ID: themeID, Field: "yes"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "yes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_AgendaItemID(topicID int) *ValueInt {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "agenda_item_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "agenda_item_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_AttachmentIDs(topicID int) *ValueIntSlice {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "attachment_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "topic", id: topicID, field: "attachment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_ID(topicID int) *ValueInt {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_ListOfSpeakersID(topicID int) *ValueInt {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "list_of_speakers_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_MeetingID(topicID int) *ValueInt {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_PollIDs(topicID int) *ValueIntSlice {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "poll_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "topic", id: topicID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_ProjectionIDs(topicID int) *ValueIntSlice {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "projection_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "topic", id: topicID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_SequentialNumber(topicID int) *ValueInt {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "sequential_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_TagIDs(topicID int) *ValueIntSlice {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "tag_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "topic", id: topicID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_Text(topicID int) *ValueString {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "text"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "topic", id: topicID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_Title(topicID int) *ValueString {
	key := dskey.Key{Collection: "topic", ID: topicID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "topic", id: topicID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_CanChangeOwnPassword(userID int) *ValueBool {
	key := dskey.Key{Collection: "user", ID: userID, Field: "can_change_own_password"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "user", id: userID, field: "can_change_own_password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_CommitteeIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "committee_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_CommitteeManagementIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "committee_management_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "committee_management_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DefaultNumber(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "default_number"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "default_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DefaultPassword(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "default_password"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "default_password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DefaultStructureLevel(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "default_structure_level"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "default_structure_level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DefaultVoteWeight(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "default_vote_weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "default_vote_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DelegatedVoteIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "delegated_vote_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "delegated_vote_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Email(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "email"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "email"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_FirstName(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "first_name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "first_name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_ForwardingCommitteeIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "forwarding_committee_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "forwarding_committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Gender(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "gender"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "gender"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_ID(userID int) *ValueInt {
	key := dskey.Key{Collection: "user", ID: userID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "user", id: userID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_IsActive(userID int) *ValueBool {
	key := dskey.Key{Collection: "user", ID: userID, Field: "is_active"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "user", id: userID, field: "is_active"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_IsDemoUser(userID int) *ValueBool {
	key := dskey.Key{Collection: "user", ID: userID, Field: "is_demo_user"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "user", id: userID, field: "is_demo_user"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_IsPhysicalPerson(userID int) *ValueBool {
	key := dskey.Key{Collection: "user", ID: userID, Field: "is_physical_person"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "user", id: userID, field: "is_physical_person"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_IsPresentInMeetingIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "is_present_in_meeting_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "is_present_in_meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_LastEmailSent(userID int) *ValueInt {
	key := dskey.Key{Collection: "user", ID: userID, Field: "last_email_sent"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "user", id: userID, field: "last_email_sent"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_LastLogin(userID int) *ValueInt {
	key := dskey.Key{Collection: "user", ID: userID, Field: "last_login"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "user", id: userID, field: "last_login"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_LastName(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "last_name"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "last_name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_MeetingIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "meeting_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_MeetingUserIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "meeting_user_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "meeting_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_OptionIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "option_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "option_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_OrganizationID(userID int) *ValueInt {
	key := dskey.Key{Collection: "user", ID: userID, Field: "organization_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "user", id: userID, field: "organization_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_OrganizationManagementLevel(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "organization_management_level"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "organization_management_level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Password(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "password"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_PollCandidateIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "poll_candidate_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "poll_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_PollVotedIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "poll_voted_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "poll_voted_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Pronoun(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "pronoun"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "pronoun"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_SamlID(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "saml_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "saml_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Title(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "title"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Username(userID int) *ValueString {
	key := dskey.Key{Collection: "user", ID: userID, Field: "username"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "username", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_VoteIDs(userID int) *ValueIntSlice {
	key := dskey.Key{Collection: "user", ID: userID, Field: "vote_ids"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "vote_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_DelegatedUserID(voteID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "vote", ID: voteID, Field: "delegated_user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "vote", id: voteID, field: "delegated_user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_ID(voteID int) *ValueInt {
	key := dskey.Key{Collection: "vote", ID: voteID, Field: "id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "vote", id: voteID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_MeetingID(voteID int) *ValueInt {
	key := dskey.Key{Collection: "vote", ID: voteID, Field: "meeting_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "vote", id: voteID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_OptionID(voteID int) *ValueInt {
	key := dskey.Key{Collection: "vote", ID: voteID, Field: "option_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "vote", id: voteID, field: "option_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_UserID(voteID int) *ValueMaybeInt {
	key := dskey.Key{Collection: "vote", ID: voteID, Field: "user_id"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "vote", id: voteID, field: "user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_UserToken(voteID int) *ValueString {
	key := dskey.Key{Collection: "vote", ID: voteID, Field: "user_token"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "vote", id: voteID, field: "user_token", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_Value(voteID int) *ValueString {
	key := dskey.Key{Collection: "vote", ID: voteID, Field: "value"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "vote", id: voteID, field: "value"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_Weight(voteID int) *ValueString {
	key := dskey.Key{Collection: "vote", ID: voteID, Field: "weight"}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "vote", id: voteID, field: "weight"}
	r.requested[key] = v
	return v
}
