// Code generated from models.yml DO NOT EDIT.
package dsfetch

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/fastjson"
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
		r, err := fastjson.DecodeInt(p)
		if err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
		v.value = r
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
		r, err := fastjson.DecodeIntList(p)
		if err != nil {
			return fmt.Errorf("decoding value %q: %w", p, err)
		}
		v.value = r
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
	key, err := dskey.FromParts("action_worker", actionWorkerID, "created")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "created", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_ID(actionWorkerID int) *ValueInt {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_Name(actionWorkerID int) *ValueString {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_Result(actionWorkerID int) *ValueJSON {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "result")
	if err != nil {
		r.err = err
		return &ValueJSON{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "result"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_State(actionWorkerID int) *ValueString {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "state")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "state", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ActionWorker_Timestamp(actionWorkerID int) *ValueInt {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "timestamp")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "actionWorker", id: actionWorkerID, field: "timestamp", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ChildIDs(agendaItemID int) *ValueIntSlice {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "child_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "agendaItem", id: agendaItemID, field: "child_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Closed(agendaItemID int) *ValueBool {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "closed")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "agendaItem", id: agendaItemID, field: "closed"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Comment(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "comment")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "agendaItem", id: agendaItemID, field: "comment"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ContentObjectID(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "content_object_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "agendaItem", id: agendaItemID, field: "content_object_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Duration(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "duration")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "duration"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ID(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_IsHidden(agendaItemID int) *ValueBool {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "is_hidden")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "agendaItem", id: agendaItemID, field: "is_hidden"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_IsInternal(agendaItemID int) *ValueBool {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "is_internal")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "agendaItem", id: agendaItemID, field: "is_internal"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ItemNumber(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "item_number")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "agendaItem", id: agendaItemID, field: "item_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Level(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "level")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_MeetingID(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ParentID(agendaItemID int) *ValueMaybeInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "parent_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "parent_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_ProjectionIDs(agendaItemID int) *ValueIntSlice {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "agendaItem", id: agendaItemID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_TagIDs(agendaItemID int) *ValueIntSlice {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "tag_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "agendaItem", id: agendaItemID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Type(agendaItemID int) *ValueString {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "agendaItem", id: agendaItemID, field: "type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AgendaItem_Weight(agendaItemID int) *ValueInt {
	key, err := dskey.FromParts("agenda_item", agendaItemID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "agendaItem", id: agendaItemID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_AssignmentID(assignmentCandidateID int) *ValueInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "assignment_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "assignment_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_ID(assignmentCandidateID int) *ValueInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_MeetingID(assignmentCandidateID int) *ValueInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_MeetingUserID(assignmentCandidateID int) *ValueMaybeInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "meeting_user_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "meeting_user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) AssignmentCandidate_Weight(assignmentCandidateID int) *ValueInt {
	key, err := dskey.FromParts("assignment_candidate", assignmentCandidateID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignmentCandidate", id: assignmentCandidateID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_AgendaItemID(assignmentID int) *ValueMaybeInt {
	key, err := dskey.FromParts("assignment", assignmentID, "agenda_item_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "assignment", id: assignmentID, field: "agenda_item_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_AttachmentIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "attachment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "attachment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_CandidateIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "candidate_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_DefaultPollDescription(assignmentID int) *ValueString {
	key, err := dskey.FromParts("assignment", assignmentID, "default_poll_description")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "assignment", id: assignmentID, field: "default_poll_description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_Description(assignmentID int) *ValueString {
	key, err := dskey.FromParts("assignment", assignmentID, "description")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "assignment", id: assignmentID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_ID(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_ListOfSpeakersID(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "list_of_speakers_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_MeetingID(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_NumberPollCandidates(assignmentID int) *ValueBool {
	key, err := dskey.FromParts("assignment", assignmentID, "number_poll_candidates")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "assignment", id: assignmentID, field: "number_poll_candidates"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_OpenPosts(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "open_posts")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "open_posts"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_Phase(assignmentID int) *ValueString {
	key, err := dskey.FromParts("assignment", assignmentID, "phase")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "assignment", id: assignmentID, field: "phase"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_PollIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "poll_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_ProjectionIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_SequentialNumber(assignmentID int) *ValueInt {
	key, err := dskey.FromParts("assignment", assignmentID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "assignment", id: assignmentID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_TagIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "tag_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "assignment", id: assignmentID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Assignment_Title(assignmentID int) *ValueString {
	key, err := dskey.FromParts("assignment", assignmentID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "assignment", id: assignmentID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_ChatMessageIDs(chatGroupID int) *ValueIntSlice {
	key, err := dskey.FromParts("chat_group", chatGroupID, "chat_message_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "chatGroup", id: chatGroupID, field: "chat_message_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_ID(chatGroupID int) *ValueInt {
	key, err := dskey.FromParts("chat_group", chatGroupID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatGroup", id: chatGroupID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_MeetingID(chatGroupID int) *ValueInt {
	key, err := dskey.FromParts("chat_group", chatGroupID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatGroup", id: chatGroupID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_Name(chatGroupID int) *ValueString {
	key, err := dskey.FromParts("chat_group", chatGroupID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "chatGroup", id: chatGroupID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_ReadGroupIDs(chatGroupID int) *ValueIntSlice {
	key, err := dskey.FromParts("chat_group", chatGroupID, "read_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "chatGroup", id: chatGroupID, field: "read_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_Weight(chatGroupID int) *ValueInt {
	key, err := dskey.FromParts("chat_group", chatGroupID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatGroup", id: chatGroupID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatGroup_WriteGroupIDs(chatGroupID int) *ValueIntSlice {
	key, err := dskey.FromParts("chat_group", chatGroupID, "write_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "chatGroup", id: chatGroupID, field: "write_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_ChatGroupID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "chat_group_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "chat_group_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_Content(chatMessageID int) *ValueString {
	key, err := dskey.FromParts("chat_message", chatMessageID, "content")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "chatMessage", id: chatMessageID, field: "content", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_Created(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "created")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "created", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_ID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_MeetingID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ChatMessage_MeetingUserID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "meeting_user_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "chatMessage", id: chatMessageID, field: "meeting_user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_DefaultMeetingID(committeeID int) *ValueMaybeInt {
	key, err := dskey.FromParts("committee", committeeID, "default_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "committee", id: committeeID, field: "default_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_Description(committeeID int) *ValueString {
	key, err := dskey.FromParts("committee", committeeID, "description")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "committee", id: committeeID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ExternalID(committeeID int) *ValueString {
	key, err := dskey.FromParts("committee", committeeID, "external_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "committee", id: committeeID, field: "external_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ForwardToCommitteeIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "forward_to_committee_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "forward_to_committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ForwardingUserID(committeeID int) *ValueMaybeInt {
	key, err := dskey.FromParts("committee", committeeID, "forwarding_user_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "committee", id: committeeID, field: "forwarding_user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ID(committeeID int) *ValueInt {
	key, err := dskey.FromParts("committee", committeeID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "committee", id: committeeID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ManagerIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "manager_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "manager_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_MeetingIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "meeting_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_Name(committeeID int) *ValueString {
	key, err := dskey.FromParts("committee", committeeID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "committee", id: committeeID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_OrganizationID(committeeID int) *ValueInt {
	key, err := dskey.FromParts("committee", committeeID, "organization_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "committee", id: committeeID, field: "organization_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_OrganizationTagIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "organization_tag_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "organization_tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_ReceiveForwardingsFromCommitteeIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "receive_forwardings_from_committee_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "receive_forwardings_from_committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Committee_UserIDs(committeeID int) *ValueIntSlice {
	key, err := dskey.FromParts("committee", committeeID, "user_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "committee", id: committeeID, field: "user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_AdminGroupForMeetingID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "admin_group_for_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "admin_group_for_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_DefaultGroupForMeetingID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "default_group_for_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "default_group_for_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_ExternalID(groupID int) *ValueString {
	key, err := dskey.FromParts("group", groupID, "external_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "group", id: groupID, field: "external_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_ID(groupID int) *ValueInt {
	key, err := dskey.FromParts("group", groupID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "group", id: groupID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_MediafileAccessGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "mediafile_access_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "mediafile_access_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_MediafileInheritedAccessGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "mediafile_inherited_access_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "mediafile_inherited_access_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_MeetingID(groupID int) *ValueInt {
	key, err := dskey.FromParts("group", groupID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "group", id: groupID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_MeetingUserIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "meeting_user_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "meeting_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_Name(groupID int) *ValueString {
	key, err := dskey.FromParts("group", groupID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "group", id: groupID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_Permissions(groupID int) *ValueStringSlice {
	key, err := dskey.FromParts("group", groupID, "permissions")
	if err != nil {
		r.err = err
		return &ValueStringSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "group", id: groupID, field: "permissions"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_PollIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "poll_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_ReadChatGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "read_chat_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "read_chat_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_ReadCommentSectionIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "read_comment_section_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "read_comment_section_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_UsedAsAssignmentPollDefaultID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "used_as_assignment_poll_default_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "used_as_assignment_poll_default_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_UsedAsMotionPollDefaultID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "used_as_motion_poll_default_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "used_as_motion_poll_default_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_UsedAsPollDefaultID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "used_as_poll_default_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "used_as_poll_default_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_UsedAsTopicPollDefaultID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "used_as_topic_poll_default_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "group", id: groupID, field: "used_as_topic_poll_default_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_Weight(groupID int) *ValueInt {
	key, err := dskey.FromParts("group", groupID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "group", id: groupID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_WriteChatGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "write_chat_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "write_chat_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Group_WriteCommentSectionIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "write_comment_section_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "group", id: groupID, field: "write_comment_section_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_Closed(listOfSpeakersID int) *ValueBool {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "closed")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "closed"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ContentObjectID(listOfSpeakersID int) *ValueString {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "content_object_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "content_object_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ID(listOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_MeetingID(listOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_ProjectionIDs(listOfSpeakersID int) *ValueIntSlice {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_SequentialNumber(listOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ListOfSpeakers_SpeakerIDs(listOfSpeakersID int) *ValueIntSlice {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "speaker_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "listOfSpeakers", id: listOfSpeakersID, field: "speaker_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_AccessGroupIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "access_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "access_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_AttachmentIDs(mediafileID int) *ValueStringSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "attachment_ids")
	if err != nil {
		r.err = err
		return &ValueStringSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "attachment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ChildIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "child_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "child_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_CreateTimestamp(mediafileID int) *ValueInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "create_timestamp")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "mediafile", id: mediafileID, field: "create_timestamp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Filename(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "filename")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "filename"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Filesize(mediafileID int) *ValueInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "filesize")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "mediafile", id: mediafileID, field: "filesize"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ID(mediafileID int) *ValueInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "mediafile", id: mediafileID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_InheritedAccessGroupIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "inherited_access_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "inherited_access_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_IsDirectory(mediafileID int) *ValueBool {
	key, err := dskey.FromParts("mediafile", mediafileID, "is_directory")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "mediafile", id: mediafileID, field: "is_directory"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_IsPublic(mediafileID int) *ValueBool {
	key, err := dskey.FromParts("mediafile", mediafileID, "is_public")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "mediafile", id: mediafileID, field: "is_public", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ListOfSpeakersID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "list_of_speakers_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "list_of_speakers_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Mimetype(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "mimetype")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "mimetype"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_OwnerID(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "owner_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "owner_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ParentID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "parent_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "parent_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_PdfInformation(mediafileID int) *ValueJSON {
	key, err := dskey.FromParts("mediafile", mediafileID, "pdf_information")
	if err != nil {
		r.err = err
		return &ValueJSON{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "mediafile", id: mediafileID, field: "pdf_information"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_ProjectionIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "mediafile", id: mediafileID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Title(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_Token(mediafileID int) *ValueString {
	key, err := dskey.FromParts("mediafile", mediafileID, "token")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "mediafile", id: mediafileID, field: "token"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontBoldInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_bold_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_bold_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontBoldItalicInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_bold_italic_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_bold_italic_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontChyronSpeakerNameInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_chyron_speaker_name_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_chyron_speaker_name_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontItalicInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_italic_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_italic_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontMonospaceInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_monospace_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_monospace_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontProjectorH1InMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_projector_h1_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_projector_h1_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontProjectorH2InMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_projector_h2_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_projector_h2_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsFontRegularInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_font_regular_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_font_regular_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfBallotPaperInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_ballot_paper_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_ballot_paper_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfFooterLInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_footer_l_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_footer_l_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfFooterRInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_footer_r_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_footer_r_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfHeaderLInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_header_l_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_header_l_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoPdfHeaderRInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_pdf_header_r_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_pdf_header_r_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoProjectorHeaderInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_projector_header_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_projector_header_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoProjectorMainInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_projector_main_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_projector_main_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Mediafile_UsedAsLogoWebHeaderInMeetingID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "used_as_logo_web_header_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "mediafile", id: mediafileID, field: "used_as_logo_web_header_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_AboutMe(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "about_me")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "about_me"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_AssignmentCandidateIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "assignment_candidate_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "assignment_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_ChatMessageIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "chat_message_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "chat_message_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_Comment(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "comment")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "comment"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_GroupIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_ID(meetingUserID int) *ValueInt {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meetingUser", id: meetingUserID, field: "id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_MeetingID(meetingUserID int) *ValueInt {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meetingUser", id: meetingUserID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_MotionSubmitterIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "motion_submitter_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "motion_submitter_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_Number(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "number")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_PersonalNoteIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "personal_note_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "personal_note_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_SpeakerIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "speaker_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "speaker_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_StructureLevel(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "structure_level")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "structure_level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_SupportedMotionIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "supported_motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "supported_motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_UserID(meetingUserID int) *ValueInt {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "user_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meetingUser", id: meetingUserID, field: "user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_VoteDelegatedToID(meetingUserID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "vote_delegated_to_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meetingUser", id: meetingUserID, field: "vote_delegated_to_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_VoteDelegationsFromIDs(meetingUserID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "vote_delegations_from_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meetingUser", id: meetingUserID, field: "vote_delegations_from_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MeetingUser_VoteWeight(meetingUserID int) *ValueString {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "vote_weight")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meetingUser", id: meetingUserID, field: "vote_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AdminGroupID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "admin_group_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "admin_group_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaEnableNumbering(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_enable_numbering")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "agenda_enable_numbering"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaItemCreation(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_item_creation")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "agenda_item_creation"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaItemIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_item_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "agenda_item_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaNewItemsDefaultVisibility(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_new_items_default_visibility")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "agenda_new_items_default_visibility"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaNumberPrefix(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_number_prefix")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "agenda_number_prefix"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaNumeralSystem(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_numeral_system")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "agenda_numeral_system"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaShowInternalItemsOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_show_internal_items_on_projector")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "agenda_show_internal_items_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AgendaShowSubtitles(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "agenda_show_subtitles")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "agenda_show_subtitles"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AllProjectionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "all_projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "all_projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseEnable(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "applause_enable")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "applause_enable"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseMaxAmount(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "applause_max_amount")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "applause_max_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseMinAmount(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "applause_min_amount")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "applause_min_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseParticleImageUrl(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "applause_particle_image_url")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "applause_particle_image_url"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseShowLevel(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "applause_show_level")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "applause_show_level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseTimeout(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "applause_timeout")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "applause_timeout"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ApplauseType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "applause_type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "applause_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentCandidateIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_candidate_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "assignment_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "assignment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollAddCandidatesToListOfSpeakers(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_add_candidates_to_list_of_speakers")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_add_candidates_to_list_of_speakers"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollBallotPaperNumber(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_ballot_paper_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_ballot_paper_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollBallotPaperSelection(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_ballot_paper_selection")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_ballot_paper_selection"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultBackend(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_backend")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_backend"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultMethod(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_method")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_method"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_onehundred_percent_base")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_onehundred_percent_base"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollDefaultType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_default_type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_default_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollEnableMaxVotesPerOption(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_enable_max_votes_per_option")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_enable_max_votes_per_option"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentPollSortPollResultByVotes(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "assignment_poll_sort_poll_result_by_votes")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "assignment_poll_sort_poll_result_by_votes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentsExportPreamble(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignments_export_preamble")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignments_export_preamble"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_AssignmentsExportTitle(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "assignments_export_title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "assignments_export_title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ChatGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "chat_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "chat_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ChatMessageIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "chat_message_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "chat_message_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_CommitteeID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "committee_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "committee_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceAutoConnect(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_auto_connect")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_auto_connect"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceAutoConnectNextSpeakers(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "conference_auto_connect_next_speakers")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "conference_auto_connect_next_speakers"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceEnableHelpdesk(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_enable_helpdesk")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_enable_helpdesk"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceLosRestriction(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_los_restriction")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_los_restriction"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceOpenMicrophone(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_open_microphone")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_open_microphone"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceOpenVideo(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_open_video")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_open_video"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceShow(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "conference_show")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "conference_show"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceStreamPosterUrl(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "conference_stream_poster_url")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "conference_stream_poster_url"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ConferenceStreamUrl(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "conference_stream_url")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "conference_stream_url"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_CustomTranslations(meetingID int) *ValueJSON {
	key, err := dskey.FromParts("meeting", meetingID, "custom_translations")
	if err != nil {
		r.err = err
		return &ValueJSON{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "meeting", id: meetingID, field: "custom_translations"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultGroupID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "default_group_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "default_group_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultMeetingForCommitteeID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "default_meeting_for_committee_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "default_meeting_for_committee_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorAgendaItemListIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_agenda_item_list_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_agenda_item_list_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorAmendmentIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_amendment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_amendment_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorAssignmentIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_assignment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_assignment_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorAssignmentPollIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_assignment_poll_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_assignment_poll_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_countdown_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_countdown_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorCurrentListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_current_list_of_speakers_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_current_list_of_speakers_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_list_of_speakers_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_list_of_speakers_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMediafileIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_mediafile_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_mediafile_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMessageIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_message_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_message_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMotionBlockIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_motion_block_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_motion_block_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMotionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_motion_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorMotionPollIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_motion_poll_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_motion_poll_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorPollIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_poll_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_poll_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_DefaultProjectorTopicIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "default_projector_topic_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "default_projector_topic_ids", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_Description(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "description")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_EnableAnonymous(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "enable_anonymous")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "enable_anonymous"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_EndTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "end_time")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "end_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportCsvEncoding(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "export_csv_encoding")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "export_csv_encoding"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportCsvSeparator(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "export_csv_separator")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "export_csv_separator"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfFontsize(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_fontsize")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_fontsize"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfLineHeight(meetingID int) *ValueFloat {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_line_height")
	if err != nil {
		r.err = err
		return &ValueFloat{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueFloat)
	}
	v := &ValueFloat{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_line_height"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginBottom(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_page_margin_bottom")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_page_margin_bottom"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginLeft(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_page_margin_left")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_page_margin_left"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginRight(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_page_margin_right")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_page_margin_right"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPageMarginTop(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_page_margin_top")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_page_margin_top"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPagenumberAlignment(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_pagenumber_alignment")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_pagenumber_alignment"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExportPdfPagesize(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "export_pdf_pagesize")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "export_pdf_pagesize"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ExternalID(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "external_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "external_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontBoldID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_bold_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_bold_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontBoldItalicID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_bold_italic_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_bold_italic_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontChyronSpeakerNameID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_chyron_speaker_name_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_chyron_speaker_name_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontItalicID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_italic_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_italic_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontMonospaceID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_monospace_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_monospace_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontProjectorH1ID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_projector_h1_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_projector_h1_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontProjectorH2ID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_projector_h2_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_projector_h2_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_FontRegularID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "font_regular_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "font_regular_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ForwardedMotionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "forwarded_motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "forwarded_motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_GroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ImportedAt(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "imported_at")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "imported_at"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_IsActiveInOrganizationID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "is_active_in_organization_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "is_active_in_organization_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_IsArchivedInOrganizationID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "is_archived_in_organization_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "is_archived_in_organization_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_JitsiDomain(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "jitsi_domain")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "jitsi_domain"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_JitsiRoomName(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "jitsi_room_name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "jitsi_room_name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_JitsiRoomPassword(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "jitsi_room_password")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "jitsi_room_password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_Language(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "language")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "language"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersAmountLastOnProjector(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_amount_last_on_projector")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_amount_last_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersAmountNextOnProjector(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_amount_next_on_projector")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_amount_next_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCanSetContributionSelf(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_can_set_contribution_self")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_can_set_contribution_self"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersClosingDisablesPointOfOrder(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_closing_disables_point_of_order")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_closing_disables_point_of_order"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCountdownID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_countdown_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_countdown_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersCoupleCountdown(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_couple_countdown")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_couple_countdown"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersEnablePointOfOrderCategories(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_enable_point_of_order_categories")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_enable_point_of_order_categories"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_enable_point_of_order_speakers")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_enable_point_of_order_speakers"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersEnableProContraSpeech(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_enable_pro_contra_speech")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_enable_pro_contra_speech"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersInitiallyClosed(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_initially_closed")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_initially_closed"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersPresentUsersOnly(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_present_users_only")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_present_users_only"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_show_amount_of_speakers_on_slide")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_show_amount_of_speakers_on_slide"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersShowFirstContribution(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_show_first_contribution")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_show_first_contribution"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ListOfSpeakersSpeakerNoteForEveryone(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "list_of_speakers_speaker_note_for_everyone")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "list_of_speakers_speaker_note_for_everyone"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_Location(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "location")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "location"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfBallotPaperID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_ballot_paper_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_ballot_paper_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfFooterLID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_footer_l_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_footer_l_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfFooterRID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_footer_r_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_footer_r_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfHeaderLID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_header_l_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_header_l_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoPdfHeaderRID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_pdf_header_r_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_pdf_header_r_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoProjectorHeaderID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_projector_header_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_projector_header_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoProjectorMainID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_projector_main_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_projector_main_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_LogoWebHeaderID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "logo_web_header_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "logo_web_header_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MediafileIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "mediafile_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "mediafile_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MeetingUserIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "meeting_user_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "meeting_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionBlockIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_block_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_block_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionCategoryIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_category_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_category_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionChangeRecommendationIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_change_recommendation_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_change_recommendation_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionCommentIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_comment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_comment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionCommentSectionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_comment_section_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_comment_section_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollBallotPaperNumber(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_ballot_paper_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_ballot_paper_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollBallotPaperSelection(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_ballot_paper_selection")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_ballot_paper_selection"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultBackend(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_backend")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_default_backend"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_default_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_onehundred_percent_base")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_default_onehundred_percent_base"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionPollDefaultType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motion_poll_default_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionStateIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_state_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_state_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionStatuteParagraphIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_statute_paragraph_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_statute_paragraph_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionSubmitterIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_submitter_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_submitter_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionWorkflowIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "motion_workflow_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "motion_workflow_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsEnabled(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_enabled")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_enabled"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsInMainList(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_in_main_list")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_in_main_list"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsMultipleParagraphs(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_multiple_paragraphs")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_multiple_paragraphs"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsOfAmendments(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_of_amendments")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_of_amendments"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsPrefix(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_prefix")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_prefix"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsAmendmentsTextMode(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_amendments_text_mode")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_amendments_text_mode"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsBlockSlideColumns(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_block_slide_columns")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_block_slide_columns"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultAmendmentWorkflowID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_amendment_workflow_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_amendment_workflow_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultLineNumbering(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_line_numbering")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_line_numbering"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultSorting(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_sorting")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_sorting"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultStatuteAmendmentWorkflowID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_statute_amendment_workflow_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_statute_amendment_workflow_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsDefaultWorkflowID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_default_workflow_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_default_workflow_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableReasonOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_reason_on_projector")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_enable_reason_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableRecommendationOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_recommendation_on_projector")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_enable_recommendation_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableSideboxOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_sidebox_on_projector")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_enable_sidebox_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsEnableTextOnProjector(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_enable_text_on_projector")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_enable_text_on_projector"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportFollowRecommendation(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_export_follow_recommendation")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_export_follow_recommendation"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportPreamble(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_export_preamble")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_export_preamble"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportSubmitterRecommendation(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_export_submitter_recommendation")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_export_submitter_recommendation"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsExportTitle(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_export_title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_export_title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsLineLength(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_line_length")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_line_length"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberMinDigits(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_number_min_digits")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_number_min_digits"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_number_type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_number_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsNumberWithBlank(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_number_with_blank")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_number_with_blank"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsPreamble(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_preamble")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_preamble"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsReasonRequired(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_reason_required")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_reason_required"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsRecommendationTextMode(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_recommendation_text_mode")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_recommendation_text_mode"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsRecommendationsBy(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_recommendations_by")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_recommendations_by"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsShowReferringMotions(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_show_referring_motions")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_show_referring_motions"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsShowSequentialNumber(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_show_sequential_number")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_show_sequential_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsStatuteRecommendationsBy(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motions_statute_recommendations_by")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "motions_statute_recommendations_by"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsStatutesEnabled(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_statutes_enabled")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "motions_statutes_enabled"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_MotionsSupportersMinAmount(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "motions_supporters_min_amount")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "motions_supporters_min_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_Name(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_OptionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "option_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "option_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_OrganizationTagIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "organization_tag_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "organization_tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PersonalNoteIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "personal_note_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "personal_note_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PointOfOrderCategoryIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "point_of_order_category_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "point_of_order_category_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollBallotPaperNumber(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "poll_ballot_paper_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "poll_ballot_paper_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollBallotPaperSelection(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_ballot_paper_selection")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_ballot_paper_selection"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollCandidateIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "poll_candidate_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "poll_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollCandidateListIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "poll_candidate_list_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "poll_candidate_list_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollCountdownID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "poll_countdown_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "poll_countdown_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollCoupleCountdown(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "poll_couple_countdown")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "poll_couple_countdown"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultBackend(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_backend")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_backend"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultMethod(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_method")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_method"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultOnehundredPercentBase(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_onehundred_percent_base")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_onehundred_percent_base"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollDefaultType(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "poll_default_type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "poll_default_type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "poll_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PollSortPollResultByVotes(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "poll_sort_poll_result_by_votes")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "poll_sort_poll_result_by_votes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_PresentUserIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "present_user_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "present_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectionIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownDefaultTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "projector_countdown_default_time")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "projector_countdown_default_time", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "projector_countdown_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "projector_countdown_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorCountdownWarningTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "projector_countdown_warning_time")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "projector_countdown_warning_time", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "projector_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "projector_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ProjectorMessageIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "projector_message_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "projector_message_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_ReferenceProjectorID(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "reference_projector_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "reference_projector_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_SpeakerIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "speaker_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "speaker_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_StartTime(meetingID int) *ValueInt {
	key, err := dskey.FromParts("meeting", meetingID, "start_time")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "meeting", id: meetingID, field: "start_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_TagIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "tag_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_TemplateForOrganizationID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "template_for_organization_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "meeting", id: meetingID, field: "template_for_organization_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_TopicIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "topic_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "topic_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_TopicPollDefaultGroupIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "topic_poll_default_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "topic_poll_default_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UserIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "user_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersAllowSelfSetPresent(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_allow_self_set_present")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "users_allow_self_set_present"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailBody(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_email_body")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_email_body"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailReplyto(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_email_replyto")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_email_replyto"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailSender(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_email_sender")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_email_sender"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEmailSubject(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_email_subject")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_email_subject"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEnablePresenceView(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_enable_presence_view")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "users_enable_presence_view"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEnableVoteDelegations(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_enable_vote_delegations")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "users_enable_vote_delegations"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersEnableVoteWeight(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "users_enable_vote_weight")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "meeting", id: meetingID, field: "users_enable_vote_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWelcometext(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_welcometext")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_welcometext"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWelcometitle(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_welcometitle")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_welcometitle"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanEncryption(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_wlan_encryption")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_wlan_encryption"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanPassword(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_wlan_password")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_wlan_password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_UsersPdfWlanSsid(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "users_pdf_wlan_ssid")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "users_pdf_wlan_ssid"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_VoteIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "vote_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "meeting", id: meetingID, field: "vote_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_WelcomeText(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "welcome_text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "welcome_text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Meeting_WelcomeTitle(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "welcome_title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "meeting", id: meetingID, field: "welcome_title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_AgendaItemID(motionBlockID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "agenda_item_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "agenda_item_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_ID(motionBlockID int) *ValueInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_Internal(motionBlockID int) *ValueBool {
	key, err := dskey.FromParts("motion_block", motionBlockID, "internal")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionBlock", id: motionBlockID, field: "internal"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_ListOfSpeakersID(motionBlockID int) *ValueInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "list_of_speakers_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_MeetingID(motionBlockID int) *ValueInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_MotionIDs(motionBlockID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_block", motionBlockID, "motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionBlock", id: motionBlockID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_ProjectionIDs(motionBlockID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_block", motionBlockID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionBlock", id: motionBlockID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_SequentialNumber(motionBlockID int) *ValueInt {
	key, err := dskey.FromParts("motion_block", motionBlockID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionBlock", id: motionBlockID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionBlock_Title(motionBlockID int) *ValueString {
	key, err := dskey.FromParts("motion_block", motionBlockID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionBlock", id: motionBlockID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_ChildIDs(motionCategoryID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "child_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "child_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_ID(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_Level(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "level")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_MeetingID(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_MotionIDs(motionCategoryID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_Name(motionCategoryID int) *ValueString {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_ParentID(motionCategoryID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "parent_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "parent_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_Prefix(motionCategoryID int) *ValueString {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "prefix")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "prefix"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_SequentialNumber(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCategory_Weight(motionCategoryID int) *ValueInt {
	key, err := dskey.FromParts("motion_category", motionCategoryID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCategory", id: motionCategoryID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_CreationTime(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "creation_time")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "creation_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_ID(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Internal(motionChangeRecommendationID int) *ValueBool {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "internal")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "internal"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_LineFrom(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "line_from")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "line_from"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_LineTo(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "line_to")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "line_to"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_MeetingID(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_MotionID(motionChangeRecommendationID int) *ValueInt {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "motion_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "motion_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_OtherDescription(motionChangeRecommendationID int) *ValueString {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "other_description")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "other_description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Rejected(motionChangeRecommendationID int) *ValueBool {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "rejected")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "rejected"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Text(motionChangeRecommendationID int) *ValueString {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionChangeRecommendation_Type(motionChangeRecommendationID int) *ValueString {
	key, err := dskey.FromParts("motion_change_recommendation", motionChangeRecommendationID, "type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionChangeRecommendation", id: motionChangeRecommendationID, field: "type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_CommentIDs(motionCommentSectionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "comment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "comment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_ID(motionCommentSectionID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_MeetingID(motionCommentSectionID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_Name(motionCommentSectionID int) *ValueString {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_ReadGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "read_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "read_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_SequentialNumber(motionCommentSectionID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_SubmitterCanWrite(motionCommentSectionID int) *ValueBool {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "submitter_can_write")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "submitter_can_write"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_Weight(motionCommentSectionID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionCommentSection_WriteGroupIDs(motionCommentSectionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_comment_section", motionCommentSectionID, "write_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionCommentSection", id: motionCommentSectionID, field: "write_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_Comment(motionCommentID int) *ValueString {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "comment")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionComment", id: motionCommentID, field: "comment"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_ID(motionCommentID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionComment", id: motionCommentID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_MeetingID(motionCommentID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionComment", id: motionCommentID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_MotionID(motionCommentID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "motion_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionComment", id: motionCommentID, field: "motion_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionComment_SectionID(motionCommentID int) *ValueInt {
	key, err := dskey.FromParts("motion_comment", motionCommentID, "section_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionComment", id: motionCommentID, field: "section_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_AllowCreatePoll(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "allow_create_poll")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "allow_create_poll"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_AllowMotionForwarding(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "allow_motion_forwarding")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "allow_motion_forwarding"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_AllowSubmitterEdit(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "allow_submitter_edit")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "allow_submitter_edit"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_AllowSupport(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "allow_support")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "allow_support"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_CssClass(motionStateID int) *ValueString {
	key, err := dskey.FromParts("motion_state", motionStateID, "css_class")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionState", id: motionStateID, field: "css_class", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_FirstStateOfWorkflowID(motionStateID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "first_state_of_workflow_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionState", id: motionStateID, field: "first_state_of_workflow_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_ID(motionStateID int) *ValueInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionState", id: motionStateID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_MeetingID(motionStateID int) *ValueInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionState", id: motionStateID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_MergeAmendmentIntoFinal(motionStateID int) *ValueString {
	key, err := dskey.FromParts("motion_state", motionStateID, "merge_amendment_into_final")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionState", id: motionStateID, field: "merge_amendment_into_final"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_MotionIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_MotionRecommendationIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "motion_recommendation_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "motion_recommendation_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_Name(motionStateID int) *ValueString {
	key, err := dskey.FromParts("motion_state", motionStateID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionState", id: motionStateID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_NextStateIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "next_state_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "next_state_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_PreviousStateIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "previous_state_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "previous_state_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_RecommendationLabel(motionStateID int) *ValueString {
	key, err := dskey.FromParts("motion_state", motionStateID, "recommendation_label")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionState", id: motionStateID, field: "recommendation_label"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_Restrictions(motionStateID int) *ValueStringSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "restrictions")
	if err != nil {
		r.err = err
		return &ValueStringSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "motionState", id: motionStateID, field: "restrictions"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_SetNumber(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "set_number")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "set_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_SetWorkflowTimestamp(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "set_workflow_timestamp")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "set_workflow_timestamp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_ShowRecommendationExtensionField(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "show_recommendation_extension_field")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "show_recommendation_extension_field"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_ShowStateExtensionField(motionStateID int) *ValueBool {
	key, err := dskey.FromParts("motion_state", motionStateID, "show_state_extension_field")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "motionState", id: motionStateID, field: "show_state_extension_field"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_SubmitterWithdrawBackIDs(motionStateID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_state", motionStateID, "submitter_withdraw_back_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionState", id: motionStateID, field: "submitter_withdraw_back_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_SubmitterWithdrawStateID(motionStateID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "submitter_withdraw_state_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionState", id: motionStateID, field: "submitter_withdraw_state_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_Weight(motionStateID int) *ValueInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionState", id: motionStateID, field: "weight", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionState_WorkflowID(motionStateID int) *ValueInt {
	key, err := dskey.FromParts("motion_state", motionStateID, "workflow_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionState", id: motionStateID, field: "workflow_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_ID(motionStatuteParagraphID int) *ValueInt {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_MeetingID(motionStatuteParagraphID int) *ValueInt {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_MotionIDs(motionStatuteParagraphID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_SequentialNumber(motionStatuteParagraphID int) *ValueInt {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Text(motionStatuteParagraphID int) *ValueString {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Title(motionStatuteParagraphID int) *ValueString {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionStatuteParagraph_Weight(motionStatuteParagraphID int) *ValueInt {
	key, err := dskey.FromParts("motion_statute_paragraph", motionStatuteParagraphID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionStatuteParagraph", id: motionStatuteParagraphID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_ID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_MeetingID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_MeetingUserID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "meeting_user_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "meeting_user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_MotionID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "motion_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "motion_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionSubmitter_Weight(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionSubmitter", id: motionSubmitterID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "default_amendment_workflow_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "default_amendment_workflow_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultStatuteAmendmentWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "default_statute_amendment_workflow_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "default_statute_amendment_workflow_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_DefaultWorkflowMeetingID(motionWorkflowID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "default_workflow_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "default_workflow_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_FirstStateID(motionWorkflowID int) *ValueInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "first_state_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "first_state_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_ID(motionWorkflowID int) *ValueInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_MeetingID(motionWorkflowID int) *ValueInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_Name(motionWorkflowID int) *ValueString {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_SequentialNumber(motionWorkflowID int) *ValueInt {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) MotionWorkflow_StateIDs(motionWorkflowID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion_workflow", motionWorkflowID, "state_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motionWorkflow", id: motionWorkflowID, field: "state_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AgendaItemID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "agenda_item_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "agenda_item_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AllDerivedMotionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "all_derived_motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "all_derived_motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AllOriginIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "all_origin_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "all_origin_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AmendmentIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "amendment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "amendment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AmendmentParagraphs(motionID int) *ValueJSON {
	key, err := dskey.FromParts("motion", motionID, "amendment_paragraphs")
	if err != nil {
		r.err = err
		return &ValueJSON{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "motion", id: motionID, field: "amendment_paragraphs"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_AttachmentIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "attachment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "attachment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_BlockID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "block_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "block_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_CategoryID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "category_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "category_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_CategoryWeight(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "category_weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "category_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ChangeRecommendationIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "change_recommendation_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "change_recommendation_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_CommentIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "comment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "comment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Created(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "created")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "created"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_DerivedMotionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "derived_motion_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "derived_motion_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Forwarded(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "forwarded")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "forwarded"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ID(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_LastModified(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "last_modified")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "last_modified"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_LeadMotionID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "lead_motion_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "lead_motion_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ListOfSpeakersID(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "list_of_speakers_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_MeetingID(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ModifiedFinalVersion(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "modified_final_version")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "modified_final_version"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Number(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "number")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_NumberValue(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "number_value")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "number_value"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_OptionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "option_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "option_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_OriginID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "origin_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "origin_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_OriginMeetingID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "origin_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "origin_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_PersonalNoteIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "personal_note_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "personal_note_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_PollIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "poll_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ProjectionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Reason(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "reason")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "reason"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_RecommendationExtension(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "recommendation_extension")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "recommendation_extension"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_RecommendationExtensionReferenceIDs(motionID int) *ValueStringSlice {
	key, err := dskey.FromParts("motion", motionID, "recommendation_extension_reference_ids")
	if err != nil {
		r.err = err
		return &ValueStringSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "motion", id: motionID, field: "recommendation_extension_reference_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_RecommendationID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "recommendation_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "recommendation_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ReferencedInMotionRecommendationExtensionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "referenced_in_motion_recommendation_extension_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "referenced_in_motion_recommendation_extension_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_ReferencedInMotionStateExtensionIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "referenced_in_motion_state_extension_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "referenced_in_motion_state_extension_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SequentialNumber(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SortChildIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "sort_child_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "sort_child_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SortParentID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "sort_parent_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "sort_parent_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SortWeight(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "sort_weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "sort_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StartLineNumber(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "start_line_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "start_line_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StateExtension(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "state_extension")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "state_extension"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StateExtensionReferenceIDs(motionID int) *ValueStringSlice {
	key, err := dskey.FromParts("motion", motionID, "state_extension_reference_ids")
	if err != nil {
		r.err = err
		return &ValueStringSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "motion", id: motionID, field: "state_extension_reference_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StateID(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "state_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "state_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_StatuteParagraphID(motionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("motion", motionID, "statute_paragraph_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "motion", id: motionID, field: "statute_paragraph_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SubmitterIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "submitter_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "submitter_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_SupporterMeetingUserIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "supporter_meeting_user_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "supporter_meeting_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_TagIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "tag_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "motion", id: motionID, field: "tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Text(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_Title(motionID int) *ValueString {
	key, err := dskey.FromParts("motion", motionID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "motion", id: motionID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Motion_WorkflowTimestamp(motionID int) *ValueInt {
	key, err := dskey.FromParts("motion", motionID, "workflow_timestamp")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "motion", id: motionID, field: "workflow_timestamp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_Abstain(optionID int) *ValueString {
	key, err := dskey.FromParts("option", optionID, "abstain")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "option", id: optionID, field: "abstain"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_ContentObjectID(optionID int) *ValueMaybeString {
	key, err := dskey.FromParts("option", optionID, "content_object_id")
	if err != nil {
		r.err = err
		return &ValueMaybeString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeString)
	}
	v := &ValueMaybeString{fetch: r, collection: "option", id: optionID, field: "content_object_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_ID(optionID int) *ValueInt {
	key, err := dskey.FromParts("option", optionID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "option", id: optionID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_MeetingID(optionID int) *ValueInt {
	key, err := dskey.FromParts("option", optionID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "option", id: optionID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_No(optionID int) *ValueString {
	key, err := dskey.FromParts("option", optionID, "no")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "option", id: optionID, field: "no"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_PollID(optionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("option", optionID, "poll_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "option", id: optionID, field: "poll_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_Text(optionID int) *ValueString {
	key, err := dskey.FromParts("option", optionID, "text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "option", id: optionID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_UsedAsGlobalOptionInPollID(optionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("option", optionID, "used_as_global_option_in_poll_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "option", id: optionID, field: "used_as_global_option_in_poll_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_VoteIDs(optionID int) *ValueIntSlice {
	key, err := dskey.FromParts("option", optionID, "vote_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "option", id: optionID, field: "vote_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_Weight(optionID int) *ValueInt {
	key, err := dskey.FromParts("option", optionID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "option", id: optionID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Option_Yes(optionID int) *ValueString {
	key, err := dskey.FromParts("option", optionID, "yes")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "option", id: optionID, field: "yes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_Color(organizationTagID int) *ValueString {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "color")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organizationTag", id: organizationTagID, field: "color", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_ID(organizationTagID int) *ValueInt {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organizationTag", id: organizationTagID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_Name(organizationTagID int) *ValueString {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organizationTag", id: organizationTagID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_OrganizationID(organizationTagID int) *ValueInt {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "organization_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organizationTag", id: organizationTagID, field: "organization_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) OrganizationTag_TaggedIDs(organizationTagID int) *ValueStringSlice {
	key, err := dskey.FromParts("organization_tag", organizationTagID, "tagged_ids")
	if err != nil {
		r.err = err
		return &ValueStringSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "organizationTag", id: organizationTagID, field: "tagged_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ActiveMeetingIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "active_meeting_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "active_meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ArchivedMeetingIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "archived_meeting_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "archived_meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_CommitteeIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "committee_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_DefaultLanguage(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "default_language")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "default_language", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_Description(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "description")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_EnableChat(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "enable_chat")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "organization", id: organizationID, field: "enable_chat"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_EnableElectronicVoting(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "enable_electronic_voting")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "organization", id: organizationID, field: "enable_electronic_voting"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_Genders(organizationID int) *ValueStringSlice {
	key, err := dskey.FromParts("organization", organizationID, "genders")
	if err != nil {
		r.err = err
		return &ValueStringSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "organization", id: organizationID, field: "genders"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ID(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organization", id: organizationID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_LegalNotice(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "legal_notice")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "legal_notice"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_LimitOfMeetings(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "limit_of_meetings")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organization", id: organizationID, field: "limit_of_meetings"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_LimitOfUsers(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "limit_of_users")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organization", id: organizationID, field: "limit_of_users"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_LoginText(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "login_text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "login_text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_MediafileIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "mediafile_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "mediafile_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_Name(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_OrganizationTagIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "organization_tag_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "organization_tag_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_PrivacyPolicy(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "privacy_policy")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "privacy_policy"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ResetPasswordVerboseErrors(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "reset_password_verbose_errors")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "organization", id: organizationID, field: "reset_password_verbose_errors"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlAttrMapping(organizationID int) *ValueJSON {
	key, err := dskey.FromParts("organization", organizationID, "saml_attr_mapping")
	if err != nil {
		r.err = err
		return &ValueJSON{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "organization", id: organizationID, field: "saml_attr_mapping"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlEnabled(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "saml_enabled")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "organization", id: organizationID, field: "saml_enabled"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlLoginButtonText(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "saml_login_button_text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "saml_login_button_text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlMetadataIDp(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "saml_metadata_idp")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "saml_metadata_idp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlMetadataSp(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "saml_metadata_sp")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "saml_metadata_sp"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_SamlPrivateKey(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "saml_private_key")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "saml_private_key"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_TemplateMeetingIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "template_meeting_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "template_meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ThemeID(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "theme_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "organization", id: organizationID, field: "theme_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_ThemeIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "theme_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "theme_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_Url(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "url")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "url"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UserIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "user_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "organization", id: organizationID, field: "user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UsersEmailBody(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "users_email_body")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "users_email_body"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UsersEmailReplyto(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "users_email_replyto")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "users_email_replyto"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UsersEmailSender(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "users_email_sender")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "users_email_sender"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_UsersEmailSubject(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "users_email_subject")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "users_email_subject"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Organization_VoteDecryptPublicMainKey(organizationID int) *ValueString {
	key, err := dskey.FromParts("organization", organizationID, "vote_decrypt_public_main_key")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "organization", id: organizationID, field: "vote_decrypt_public_main_key"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_ContentObjectID(personalNoteID int) *ValueMaybeString {
	key, err := dskey.FromParts("personal_note", personalNoteID, "content_object_id")
	if err != nil {
		r.err = err
		return &ValueMaybeString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeString)
	}
	v := &ValueMaybeString{fetch: r, collection: "personalNote", id: personalNoteID, field: "content_object_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_ID(personalNoteID int) *ValueInt {
	key, err := dskey.FromParts("personal_note", personalNoteID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "personalNote", id: personalNoteID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_MeetingID(personalNoteID int) *ValueInt {
	key, err := dskey.FromParts("personal_note", personalNoteID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "personalNote", id: personalNoteID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_MeetingUserID(personalNoteID int) *ValueInt {
	key, err := dskey.FromParts("personal_note", personalNoteID, "meeting_user_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "personalNote", id: personalNoteID, field: "meeting_user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_Note(personalNoteID int) *ValueString {
	key, err := dskey.FromParts("personal_note", personalNoteID, "note")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "personalNote", id: personalNoteID, field: "note"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PersonalNote_Star(personalNoteID int) *ValueBool {
	key, err := dskey.FromParts("personal_note", personalNoteID, "star")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "personalNote", id: personalNoteID, field: "star"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PointOfOrderCategory_ID(pointOfOrderCategoryID int) *ValueInt {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pointOfOrderCategory", id: pointOfOrderCategoryID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PointOfOrderCategory_MeetingID(pointOfOrderCategoryID int) *ValueInt {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pointOfOrderCategory", id: pointOfOrderCategoryID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PointOfOrderCategory_Rank(pointOfOrderCategoryID int) *ValueInt {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "rank")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pointOfOrderCategory", id: pointOfOrderCategoryID, field: "rank", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PointOfOrderCategory_SpeakerIDs(pointOfOrderCategoryID int) *ValueIntSlice {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "speaker_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "pointOfOrderCategory", id: pointOfOrderCategoryID, field: "speaker_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PointOfOrderCategory_Text(pointOfOrderCategoryID int) *ValueString {
	key, err := dskey.FromParts("point_of_order_category", pointOfOrderCategoryID, "text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "pointOfOrderCategory", id: pointOfOrderCategoryID, field: "text", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidateList_ID(pollCandidateListID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate_list", pollCandidateListID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidateList", id: pollCandidateListID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidateList_MeetingID(pollCandidateListID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate_list", pollCandidateListID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidateList", id: pollCandidateListID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidateList_OptionID(pollCandidateListID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate_list", pollCandidateListID, "option_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidateList", id: pollCandidateListID, field: "option_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidateList_PollCandidateIDs(pollCandidateListID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll_candidate_list", pollCandidateListID, "poll_candidate_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "pollCandidateList", id: pollCandidateListID, field: "poll_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_ID(pollCandidateID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_MeetingID(pollCandidateID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_PollCandidateListID(pollCandidateID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "poll_candidate_list_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "poll_candidate_list_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_UserID(pollCandidateID int) *ValueMaybeInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "user_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) PollCandidate_Weight(pollCandidateID int) *ValueInt {
	key, err := dskey.FromParts("poll_candidate", pollCandidateID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "pollCandidate", id: pollCandidateID, field: "weight", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Backend(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "backend")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "backend", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_ContentObjectID(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "content_object_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "content_object_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_CryptKey(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "crypt_key")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "crypt_key"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_CryptSignature(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "crypt_signature")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "crypt_signature"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Description(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "description")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_EntitledGroupIDs(pollID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll", pollID, "entitled_group_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "poll", id: pollID, field: "entitled_group_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_EntitledUsersAtStop(pollID int) *ValueJSON {
	key, err := dskey.FromParts("poll", pollID, "entitled_users_at_stop")
	if err != nil {
		r.err = err
		return &ValueJSON{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "poll", id: pollID, field: "entitled_users_at_stop"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_GlobalAbstain(pollID int) *ValueBool {
	key, err := dskey.FromParts("poll", pollID, "global_abstain")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "poll", id: pollID, field: "global_abstain"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_GlobalNo(pollID int) *ValueBool {
	key, err := dskey.FromParts("poll", pollID, "global_no")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "poll", id: pollID, field: "global_no"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_GlobalOptionID(pollID int) *ValueMaybeInt {
	key, err := dskey.FromParts("poll", pollID, "global_option_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "poll", id: pollID, field: "global_option_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_GlobalYes(pollID int) *ValueBool {
	key, err := dskey.FromParts("poll", pollID, "global_yes")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "poll", id: pollID, field: "global_yes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_ID(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_IsPseudoanonymized(pollID int) *ValueBool {
	key, err := dskey.FromParts("poll", pollID, "is_pseudoanonymized")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "poll", id: pollID, field: "is_pseudoanonymized"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_MaxVotesAmount(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "max_votes_amount")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "max_votes_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_MaxVotesPerOption(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "max_votes_per_option")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "max_votes_per_option"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_MeetingID(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_MinVotesAmount(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "min_votes_amount")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "min_votes_amount"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_OnehundredPercentBase(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "onehundred_percent_base")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "onehundred_percent_base", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_OptionIDs(pollID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll", pollID, "option_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "poll", id: pollID, field: "option_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Pollmethod(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "pollmethod")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "pollmethod", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_ProjectionIDs(pollID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll", pollID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "poll", id: pollID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_SequentialNumber(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_State(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "state")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "state"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Title(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Type(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "type", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_VoteCount(pollID int) *ValueInt {
	key, err := dskey.FromParts("poll", pollID, "vote_count")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "poll", id: pollID, field: "vote_count"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_VotedIDs(pollID int) *ValueIntSlice {
	key, err := dskey.FromParts("poll", pollID, "voted_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "poll", id: pollID, field: "voted_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_VotesRaw(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votes_raw")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votes_raw"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_VotesSignature(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votes_signature")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votes_signature"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Votescast(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votescast")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votescast"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Votesinvalid(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votesinvalid")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votesinvalid"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Poll_Votesvalid(pollID int) *ValueString {
	key, err := dskey.FromParts("poll", pollID, "votesvalid")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "poll", id: pollID, field: "votesvalid"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Content(projectionID int) *ValueJSON {
	key, err := dskey.FromParts("projection", projectionID, "content")
	if err != nil {
		r.err = err
		return &ValueJSON{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "projection", id: projectionID, field: "content"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_ContentObjectID(projectionID int) *ValueString {
	key, err := dskey.FromParts("projection", projectionID, "content_object_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projection", id: projectionID, field: "content_object_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_CurrentProjectorID(projectionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projection", projectionID, "current_projector_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projection", id: projectionID, field: "current_projector_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_HistoryProjectorID(projectionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projection", projectionID, "history_projector_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projection", id: projectionID, field: "history_projector_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_ID(projectionID int) *ValueInt {
	key, err := dskey.FromParts("projection", projectionID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projection", id: projectionID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_MeetingID(projectionID int) *ValueInt {
	key, err := dskey.FromParts("projection", projectionID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projection", id: projectionID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Options(projectionID int) *ValueJSON {
	key, err := dskey.FromParts("projection", projectionID, "options")
	if err != nil {
		r.err = err
		return &ValueJSON{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueJSON)
	}
	v := &ValueJSON{fetch: r, collection: "projection", id: projectionID, field: "options"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_PreviewProjectorID(projectionID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projection", projectionID, "preview_projector_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projection", id: projectionID, field: "preview_projector_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Stable(projectionID int) *ValueBool {
	key, err := dskey.FromParts("projection", projectionID, "stable")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projection", id: projectionID, field: "stable"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Type(projectionID int) *ValueString {
	key, err := dskey.FromParts("projection", projectionID, "type")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projection", id: projectionID, field: "type"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projection_Weight(projectionID int) *ValueInt {
	key, err := dskey.FromParts("projection", projectionID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projection", id: projectionID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_CountdownTime(projectorCountdownID int) *ValueFloat {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "countdown_time")
	if err != nil {
		r.err = err
		return &ValueFloat{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueFloat)
	}
	v := &ValueFloat{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "countdown_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_DefaultTime(projectorCountdownID int) *ValueInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "default_time")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "default_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Description(projectorCountdownID int) *ValueString {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "description")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "description"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_ID(projectorCountdownID int) *ValueInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_MeetingID(projectorCountdownID int) *ValueInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_ProjectionIDs(projectorCountdownID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Running(projectorCountdownID int) *ValueBool {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "running")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "running"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_Title(projectorCountdownID int) *ValueString {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "used_as_list_of_speakers_countdown_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "used_as_list_of_speakers_countdown_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorCountdown_UsedAsPollCountdownMeetingID(projectorCountdownID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector_countdown", projectorCountdownID, "used_as_poll_countdown_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projectorCountdown", id: projectorCountdownID, field: "used_as_poll_countdown_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorMessage_ID(projectorMessageID int) *ValueInt {
	key, err := dskey.FromParts("projector_message", projectorMessageID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorMessage", id: projectorMessageID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorMessage_MeetingID(projectorMessageID int) *ValueInt {
	key, err := dskey.FromParts("projector_message", projectorMessageID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projectorMessage", id: projectorMessageID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorMessage_Message(projectorMessageID int) *ValueString {
	key, err := dskey.FromParts("projector_message", projectorMessageID, "message")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projectorMessage", id: projectorMessageID, field: "message"}
	r.requested[key] = v
	return v
}

func (r *Fetch) ProjectorMessage_ProjectionIDs(projectorMessageID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector_message", projectorMessageID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projectorMessage", id: projectorMessageID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_AspectRatioDenominator(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "aspect_ratio_denominator")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "aspect_ratio_denominator"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_AspectRatioNumerator(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "aspect_ratio_numerator")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "aspect_ratio_numerator"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_BackgroundColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "background_color")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "background_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ChyronBackgroundColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "chyron_background_color")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "chyron_background_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ChyronFontColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "chyron_font_color")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "chyron_font_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Color(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "color")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_CurrentProjectionIDs(projectorID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector", projectorID, "current_projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projector", id: projectorID, field: "current_projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_HeaderBackgroundColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "header_background_color")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "header_background_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_HeaderFontColor(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "header_font_color")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "header_font_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_HeaderH1Color(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "header_h1_color")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "header_h1_color"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_HistoryProjectionIDs(projectorID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector", projectorID, "history_projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projector", id: projectorID, field: "history_projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ID(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_IsInternal(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "is_internal")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "is_internal"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_MeetingID(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Name(projectorID int) *ValueString {
	key, err := dskey.FromParts("projector", projectorID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "projector", id: projectorID, field: "name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_PreviewProjectionIDs(projectorID int) *ValueIntSlice {
	key, err := dskey.FromParts("projector", projectorID, "preview_projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "projector", id: projectorID, field: "preview_projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Scale(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "scale")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "scale"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Scroll(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "scroll")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "scroll"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_SequentialNumber(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ShowClock(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "show_clock")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "show_clock"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ShowHeaderFooter(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "show_header_footer")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "show_header_footer"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ShowLogo(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "show_logo")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "show_logo"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_ShowTitle(projectorID int) *ValueBool {
	key, err := dskey.FromParts("projector", projectorID, "show_title")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "projector", id: projectorID, field: "show_title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAgendaItemListInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_agenda_item_list_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_agenda_item_list_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAmendmentInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_amendment_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_amendment_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAssignmentInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_assignment_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_assignment_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForAssignmentPollInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_assignment_poll_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_assignment_poll_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForCountdownInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_countdown_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_countdown_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_current_list_of_speakers_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_current_list_of_speakers_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForListOfSpeakersInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_list_of_speakers_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_list_of_speakers_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMediafileInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_mediafile_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_mediafile_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMessageInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_message_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_message_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionBlockInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_motion_block_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_motion_block_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_motion_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_motion_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForMotionPollInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_motion_poll_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_motion_poll_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForPollInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_poll_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_poll_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsDefaultProjectorForTopicInMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_default_projector_for_topic_in_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_default_projector_for_topic_in_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_UsedAsReferenceProjectorMeetingID(projectorID int) *ValueMaybeInt {
	key, err := dskey.FromParts("projector", projectorID, "used_as_reference_projector_meeting_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "projector", id: projectorID, field: "used_as_reference_projector_meeting_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Projector_Width(projectorID int) *ValueInt {
	key, err := dskey.FromParts("projector", projectorID, "width")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "projector", id: projectorID, field: "width"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_BeginTime(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "begin_time")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "begin_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_EndTime(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "end_time")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "end_time"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_ID(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_ListOfSpeakersID(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "list_of_speakers_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_MeetingID(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_MeetingUserID(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "meeting_user_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "meeting_user_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_Note(speakerID int) *ValueString {
	key, err := dskey.FromParts("speaker", speakerID, "note")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "speaker", id: speakerID, field: "note"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_PointOfOrder(speakerID int) *ValueBool {
	key, err := dskey.FromParts("speaker", speakerID, "point_of_order")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "speaker", id: speakerID, field: "point_of_order"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_PointOfOrderCategoryID(speakerID int) *ValueMaybeInt {
	key, err := dskey.FromParts("speaker", speakerID, "point_of_order_category_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "speaker", id: speakerID, field: "point_of_order_category_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_SpeechState(speakerID int) *ValueString {
	key, err := dskey.FromParts("speaker", speakerID, "speech_state")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "speaker", id: speakerID, field: "speech_state"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Speaker_Weight(speakerID int) *ValueInt {
	key, err := dskey.FromParts("speaker", speakerID, "weight")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "speaker", id: speakerID, field: "weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Tag_ID(tagID int) *ValueInt {
	key, err := dskey.FromParts("tag", tagID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "tag", id: tagID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Tag_MeetingID(tagID int) *ValueInt {
	key, err := dskey.FromParts("tag", tagID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "tag", id: tagID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Tag_Name(tagID int) *ValueString {
	key, err := dskey.FromParts("tag", tagID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "tag", id: tagID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Tag_TaggedIDs(tagID int) *ValueStringSlice {
	key, err := dskey.FromParts("tag", tagID, "tagged_ids")
	if err != nil {
		r.err = err
		return &ValueStringSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueStringSlice)
	}
	v := &ValueStringSlice{fetch: r, collection: "tag", id: tagID, field: "tagged_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Abstain(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "abstain")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "abstain"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_100")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_200")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent300(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_300")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_300"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_400")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent50(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_50")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_50"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent500(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_500")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_500", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent600(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_600")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_600"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_700")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent800(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_800")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_800"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Accent900(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_900")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_900"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_AccentA100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_a100")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_a100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_AccentA200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_a200")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_a200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_AccentA400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_a400")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_a400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_AccentA700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "accent_a700")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "accent_a700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Headbar(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "headbar")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "headbar"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_ID(themeID int) *ValueInt {
	key, err := dskey.FromParts("theme", themeID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "theme", id: themeID, field: "id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Name(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "name", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_No(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "no")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "no"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_OrganizationID(themeID int) *ValueInt {
	key, err := dskey.FromParts("theme", themeID, "organization_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "theme", id: themeID, field: "organization_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_100")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_200")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary300(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_300")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_300"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_400")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary50(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_50")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_50"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary500(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_500")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_500", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary600(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_600")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_600"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_700")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary800(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_800")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_800"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Primary900(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_900")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_900"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_PrimaryA100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_a100")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_a100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_PrimaryA200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_a200")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_a200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_PrimaryA400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_a400")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_a400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_PrimaryA700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "primary_a700")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "primary_a700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_ThemeForOrganizationID(themeID int) *ValueMaybeInt {
	key, err := dskey.FromParts("theme", themeID, "theme_for_organization_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "theme", id: themeID, field: "theme_for_organization_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_100")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_200")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn300(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_300")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_300"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_400")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn50(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_50")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_50"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn500(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_500")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_500", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn600(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_600")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_600"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_700")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn800(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_800")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_800"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Warn900(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_900")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_900"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_WarnA100(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_a100")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_a100"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_WarnA200(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_a200")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_a200"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_WarnA400(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_a400")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_a400"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_WarnA700(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "warn_a700")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "warn_a700"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Theme_Yes(themeID int) *ValueString {
	key, err := dskey.FromParts("theme", themeID, "yes")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "theme", id: themeID, field: "yes"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_AgendaItemID(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "agenda_item_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "agenda_item_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_AttachmentIDs(topicID int) *ValueIntSlice {
	key, err := dskey.FromParts("topic", topicID, "attachment_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "topic", id: topicID, field: "attachment_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_ID(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_ListOfSpeakersID(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "list_of_speakers_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "list_of_speakers_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_MeetingID(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_PollIDs(topicID int) *ValueIntSlice {
	key, err := dskey.FromParts("topic", topicID, "poll_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "topic", id: topicID, field: "poll_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_ProjectionIDs(topicID int) *ValueIntSlice {
	key, err := dskey.FromParts("topic", topicID, "projection_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "topic", id: topicID, field: "projection_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_SequentialNumber(topicID int) *ValueInt {
	key, err := dskey.FromParts("topic", topicID, "sequential_number")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "topic", id: topicID, field: "sequential_number", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_Text(topicID int) *ValueString {
	key, err := dskey.FromParts("topic", topicID, "text")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "topic", id: topicID, field: "text"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Topic_Title(topicID int) *ValueString {
	key, err := dskey.FromParts("topic", topicID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "topic", id: topicID, field: "title", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_CanChangeOwnPassword(userID int) *ValueBool {
	key, err := dskey.FromParts("user", userID, "can_change_own_password")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "user", id: userID, field: "can_change_own_password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_CommitteeIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "committee_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_CommitteeManagementIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "committee_management_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "committee_management_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DefaultNumber(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "default_number")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "default_number"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DefaultPassword(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "default_password")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "default_password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DefaultStructureLevel(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "default_structure_level")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "default_structure_level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DefaultVoteWeight(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "default_vote_weight")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "default_vote_weight"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_DelegatedVoteIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "delegated_vote_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "delegated_vote_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Email(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "email")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "email"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_FirstName(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "first_name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "first_name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_ForwardingCommitteeIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "forwarding_committee_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "forwarding_committee_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Gender(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "gender")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "gender"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_ID(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "user", id: userID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_IsActive(userID int) *ValueBool {
	key, err := dskey.FromParts("user", userID, "is_active")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "user", id: userID, field: "is_active"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_IsDemoUser(userID int) *ValueBool {
	key, err := dskey.FromParts("user", userID, "is_demo_user")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "user", id: userID, field: "is_demo_user"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_IsPhysicalPerson(userID int) *ValueBool {
	key, err := dskey.FromParts("user", userID, "is_physical_person")
	if err != nil {
		r.err = err
		return &ValueBool{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueBool)
	}
	v := &ValueBool{fetch: r, collection: "user", id: userID, field: "is_physical_person"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_IsPresentInMeetingIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "is_present_in_meeting_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "is_present_in_meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_LastEmailSent(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "last_email_sent")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "user", id: userID, field: "last_email_sent"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_LastLogin(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "last_login")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "user", id: userID, field: "last_login"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_LastName(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "last_name")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "last_name"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_MeetingIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "meeting_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "meeting_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_MeetingUserIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "meeting_user_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "meeting_user_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_OptionIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "option_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "option_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_OrganizationID(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "organization_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "user", id: userID, field: "organization_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_OrganizationManagementLevel(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "organization_management_level")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "organization_management_level"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Password(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "password")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "password"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_PollCandidateIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "poll_candidate_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "poll_candidate_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_PollVotedIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "poll_voted_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "poll_voted_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Pronoun(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "pronoun")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "pronoun"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_SamlID(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "saml_id")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "saml_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Title(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "title")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "title"}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_Username(userID int) *ValueString {
	key, err := dskey.FromParts("user", userID, "username")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "user", id: userID, field: "username", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) User_VoteIDs(userID int) *ValueIntSlice {
	key, err := dskey.FromParts("user", userID, "vote_ids")
	if err != nil {
		r.err = err
		return &ValueIntSlice{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueIntSlice)
	}
	v := &ValueIntSlice{fetch: r, collection: "user", id: userID, field: "vote_ids"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_DelegatedUserID(voteID int) *ValueMaybeInt {
	key, err := dskey.FromParts("vote", voteID, "delegated_user_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "vote", id: voteID, field: "delegated_user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_ID(voteID int) *ValueInt {
	key, err := dskey.FromParts("vote", voteID, "id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "vote", id: voteID, field: "id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_MeetingID(voteID int) *ValueInt {
	key, err := dskey.FromParts("vote", voteID, "meeting_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "vote", id: voteID, field: "meeting_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_OptionID(voteID int) *ValueInt {
	key, err := dskey.FromParts("vote", voteID, "option_id")
	if err != nil {
		r.err = err
		return &ValueInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueInt)
	}
	v := &ValueInt{fetch: r, collection: "vote", id: voteID, field: "option_id", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_UserID(voteID int) *ValueMaybeInt {
	key, err := dskey.FromParts("vote", voteID, "user_id")
	if err != nil {
		r.err = err
		return &ValueMaybeInt{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueMaybeInt)
	}
	v := &ValueMaybeInt{fetch: r, collection: "vote", id: voteID, field: "user_id"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_UserToken(voteID int) *ValueString {
	key, err := dskey.FromParts("vote", voteID, "user_token")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "vote", id: voteID, field: "user_token", required: true}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_Value(voteID int) *ValueString {
	key, err := dskey.FromParts("vote", voteID, "value")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "vote", id: voteID, field: "value"}
	r.requested[key] = v
	return v
}

func (r *Fetch) Vote_Weight(voteID int) *ValueString {
	key, err := dskey.FromParts("vote", voteID, "weight")
	if err != nil {
		r.err = err
		return &ValueString{fetch: r}
	}
	if v, ok := r.requested[key]; ok {
		return v.(*ValueString)
	}
	v := &ValueString{fetch: r, collection: "vote", id: voteID, field: "weight"}
	r.requested[key] = v
	return v
}
