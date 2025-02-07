// Code generated from models.yml DO NOT EDIT.
package dsfetch

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/fastjson"
)

type loader[T any] interface {
	lazy(ds *Fetch, id int)
	preload(ds *Fetch, id int)
	*T
}

// ValueCollection is a generic struct, where the loader interface is
// implemented by the pointer of C.
type ValueCollection[C any, T loader[C]] struct {
	id    int
	fetch *Fetch
}

func (v *ValueCollection[T, P]) Value(ctx context.Context) (T, error) {
	var collection T
	v.Lazy(&collection)

	if err := v.fetch.Execute(ctx); err != nil {
		var zero T
		return zero, err
	}
	return collection, nil
}

func (v *ValueCollection[T, P]) Lazy(collection P) {
	collection.lazy(v.fetch, v.id)
}

func (v *ValueCollection[T, P]) Preload() {
	var collection P
	collection.preload(v.fetch, v.id)
}

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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueBool) convert(p []byte) (bool, error) {
	var zero bool
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	var value bool
	if err := json.Unmarshal(p, &value); err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueBool) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueFloat) convert(p []byte) (float32, error) {
	var zero float32
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	var value float32
	if err := json.Unmarshal(p, &value); err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueFloat) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueInt) convert(p []byte) (int, error) {
	var zero int
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	value, err := fastjson.DecodeInt(p)
	if err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueInt) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueIntSlice) convert(p []byte) ([]int, error) {
	var zero []int
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	value, err := fastjson.DecodeIntList(p)
	if err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueIntSlice) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueJSON) convert(p []byte) (json.RawMessage, error) {
	var zero json.RawMessage
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	var value json.RawMessage
	if err := json.Unmarshal(p, &value); err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueJSON) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueMaybeInt) convert(p []byte) (Maybe[int], error) {
	var zero Maybe[int]
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	var value Maybe[int]
	if err := json.Unmarshal(p, &value); err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueMaybeInt) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueMaybeString) convert(p []byte) (Maybe[string], error) {
	var zero Maybe[string]
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	var value Maybe[string]
	if err := json.Unmarshal(p, &value); err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueMaybeString) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueString) convert(p []byte) (string, error) {
	var zero string
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	var value string
	if err := json.Unmarshal(p, &value); err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueString) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	value, err := v.convert(rawValue)
	if err != nil {
		return zero, fmt.Errorf("converting raw value: %w", err)
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

// convert converts the json value to the type.
func (v *ValueStringSlice) convert(p []byte) ([]string, error) {
	var zero []string
	if p == nil {
		if v.required {
			return zero, fmt.Errorf("database is corrupted. Required field %s is null", v.key)
		}
		return zero, nil
	}
	var value []string
	if err := json.Unmarshal(p, &value); err != nil {
		return zero, fmt.Errorf("decoding value %q: %w", p, err)
	}
	return value, nil
}

// setLazy sets the lazy values defiend with Lazy or Preload.
func (v *ValueStringSlice) setLazy(p []byte) error {
	value, err := v.convert(p)
	if err != nil {
		return fmt.Errorf("converting value: %w", err)
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) ActionWorker_UserID(actionWorkerID int) *ValueInt {
	key, err := dskey.FromParts("action_worker", actionWorkerID, "user_id")
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) Assignment_AttachmentMeetingMediafileIDs(assignmentID int) *ValueIntSlice {
	key, err := dskey.FromParts("assignment", assignmentID, "attachment_meeting_mediafile_ids")
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatMessage_MeetingID(chatMessageID int) *ValueInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ChatMessage_MeetingUserID(chatMessageID int) *ValueMaybeInt {
	key, err := dskey.FromParts("chat_message", chatMessageID, "meeting_user_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) Gender_ID(genderID int) *ValueInt {
	key, err := dskey.FromParts("gender", genderID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Gender_Name(genderID int) *ValueString {
	key, err := dskey.FromParts("gender", genderID, "name")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key, required: true}
}

func (r *Fetch) Gender_OrganizationID(genderID int) *ValueInt {
	key, err := dskey.FromParts("gender", genderID, "organization_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Gender_UserIDs(genderID int) *ValueIntSlice {
	key, err := dskey.FromParts("gender", genderID, "user_ids")
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

func (r *Fetch) Group_AnonymousGroupForMeetingID(groupID int) *ValueMaybeInt {
	key, err := dskey.FromParts("group", groupID, "anonymous_group_for_meeting_id")
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

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Group_MeetingID(groupID int) *ValueInt {
	key, err := dskey.FromParts("group", groupID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Group_MeetingMediafileAccessGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "meeting_mediafile_access_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Group_MeetingMediafileInheritedAccessGroupIDs(groupID int) *ValueIntSlice {
	key, err := dskey.FromParts("group", groupID, "meeting_mediafile_inherited_access_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ListOfSpeakers_MeetingID(listOfSpeakersID int) *ValueInt {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) ListOfSpeakers_ModeratorNotes(listOfSpeakersID int) *ValueString {
	key, err := dskey.FromParts("list_of_speakers", listOfSpeakersID, "moderator_notes")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
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

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) Mediafile_IsDirectory(mediafileID int) *ValueBool {
	key, err := dskey.FromParts("mediafile", mediafileID, "is_directory")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
}

func (r *Fetch) Mediafile_MeetingMediafileIDs(mediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("mediafile", mediafileID, "meeting_mediafile_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
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

func (r *Fetch) Mediafile_PublishedToMeetingsInOrganizationID(mediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("mediafile", mediafileID, "published_to_meetings_in_organization_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
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

func (r *Fetch) MeetingMediafile_AccessGroupIDs(meetingMediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "access_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_AttachmentIDs(meetingMediafileID int) *ValueStringSlice {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "attachment_ids")
	if err != nil {
		return &ValueStringSlice{err: err}
	}

	return &ValueStringSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_ID(meetingMediafileID int) *ValueInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MeetingMediafile_InheritedAccessGroupIDs(meetingMediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "inherited_access_group_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_IsPublic(meetingMediafileID int) *ValueBool {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "is_public")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key, required: true}
}

func (r *Fetch) MeetingMediafile_ListOfSpeakersID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "list_of_speakers_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_MediafileID(meetingMediafileID int) *ValueInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "mediafile_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MeetingMediafile_MeetingID(meetingMediafileID int) *ValueInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "meeting_id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
}

func (r *Fetch) MeetingMediafile_ProjectionIDs(meetingMediafileID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "projection_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsFontBoldInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_font_bold_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsFontBoldItalicInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_font_bold_italic_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsFontChyronSpeakerNameInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_font_chyron_speaker_name_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsFontItalicInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_font_italic_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsFontMonospaceInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_font_monospace_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsFontProjectorH1InMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_font_projector_h1_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsFontProjectorH2InMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_font_projector_h2_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsFontRegularInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_font_regular_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsLogoPdfBallotPaperInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_logo_pdf_ballot_paper_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsLogoPdfFooterLInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_logo_pdf_footer_l_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsLogoPdfFooterRInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_logo_pdf_footer_r_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsLogoPdfHeaderLInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_logo_pdf_header_l_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsLogoPdfHeaderRInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_logo_pdf_header_r_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsLogoProjectorHeaderInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_logo_projector_header_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsLogoProjectorMainInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_logo_projector_main_in_meeting_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) MeetingMediafile_UsedAsLogoWebHeaderInMeetingID(meetingMediafileID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting_mediafile", meetingMediafileID, "used_as_logo_web_header_in_meeting_id")
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

func (r *Fetch) MeetingUser_LockedOut(meetingUserID int) *ValueBool {
	key, err := dskey.FromParts("meeting_user", meetingUserID, "locked_out")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
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

func (r *Fetch) Meeting_AnonymousGroupID(meetingID int) *ValueMaybeInt {
	key, err := dskey.FromParts("meeting", meetingID, "anonymous_group_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) Meeting_LockedFromInside(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "locked_from_inside")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
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

func (r *Fetch) Meeting_MeetingMediafileIDs(meetingID int) *ValueIntSlice {
	key, err := dskey.FromParts("meeting", meetingID, "meeting_mediafile_ids")
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

func (r *Fetch) Meeting_MotionPollDefaultMethod(meetingID int) *ValueString {
	key, err := dskey.FromParts("meeting", meetingID, "motion_poll_default_method")
	if err != nil {
		return &ValueString{err: err}
	}

	return &ValueString{fetch: r, key: key}
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

func (r *Fetch) Meeting_MotionsCreateEnableAdditionalSubmitterText(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_create_enable_additional_submitter_text")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
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

func (r *Fetch) Meeting_MotionsHideMetadataBackground(meetingID int) *ValueBool {
	key, err := dskey.FromParts("meeting", meetingID, "motions_hide_metadata_background")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) MotionSubmitter_ID(motionSubmitterID int) *ValueInt {
	key, err := dskey.FromParts("motion_submitter", motionSubmitterID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) Motion_AttachmentMeetingMediafileIDs(motionID int) *ValueIntSlice {
	key, err := dskey.FromParts("motion", motionID, "attachment_meeting_mediafile_ids")
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) Organization_EnableAnonymous(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "enable_anonymous")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
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

func (r *Fetch) Organization_GenderIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "gender_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_ID(organizationID int) *ValueInt {
	key, err := dskey.FromParts("organization", organizationID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) Organization_PublishedMediafileIDs(organizationID int) *ValueIntSlice {
	key, err := dskey.FromParts("organization", organizationID, "published_mediafile_ids")
	if err != nil {
		return &ValueIntSlice{err: err}
	}

	return &ValueIntSlice{fetch: r, key: key}
}

func (r *Fetch) Organization_RequireDuplicateFrom(organizationID int) *ValueBool {
	key, err := dskey.FromParts("organization", organizationID, "require_duplicate_from")
	if err != nil {
		return &ValueBool{err: err}
	}

	return &ValueBool{fetch: r, key: key}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) Topic_AttachmentMeetingMediafileIDs(topicID int) *ValueIntSlice {
	key, err := dskey.FromParts("topic", topicID, "attachment_meeting_mediafile_ids")
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

	return &ValueInt{fetch: r, key: key, required: true}
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

func (r *Fetch) User_GenderID(userID int) *ValueMaybeInt {
	key, err := dskey.FromParts("user", userID, "gender_id")
	if err != nil {
		return &ValueMaybeInt{err: err}
	}

	return &ValueMaybeInt{fetch: r, key: key}
}

func (r *Fetch) User_ID(userID int) *ValueInt {
	key, err := dskey.FromParts("user", userID, "id")
	if err != nil {
		return &ValueInt{err: err}
	}

	return &ValueInt{fetch: r, key: key, required: true}
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

	return &ValueInt{fetch: r, key: key, required: true}
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

// ActionWorker has all fields from action_worker.
type ActionWorker struct {
	Created   int
	ID        int
	Name      string
	Result    json.RawMessage
	State     string
	Timestamp int
	UserID    int
	fetch     *Fetch
}

func (c *ActionWorker) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ActionWorker_Created(id).Lazy(&c.Created)
	ds.ActionWorker_ID(id).Lazy(&c.ID)
	ds.ActionWorker_Name(id).Lazy(&c.Name)
	ds.ActionWorker_Result(id).Lazy(&c.Result)
	ds.ActionWorker_State(id).Lazy(&c.State)
	ds.ActionWorker_Timestamp(id).Lazy(&c.Timestamp)
	ds.ActionWorker_UserID(id).Lazy(&c.UserID)
}

func (c *ActionWorker) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.ActionWorker_Created(id).Preload()
	ds.ActionWorker_ID(id).Preload()
	ds.ActionWorker_Name(id).Preload()
	ds.ActionWorker_Result(id).Preload()
	ds.ActionWorker_State(id).Preload()
	ds.ActionWorker_Timestamp(id).Preload()
	ds.ActionWorker_UserID(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) ActionWorker(id int) *ValueCollection[ActionWorker, *ActionWorker] {
	return &ValueCollection[ActionWorker, *ActionWorker]{
		id:    id,
		fetch: r,
	}
}

// AgendaItem has all fields from agenda_item.
type AgendaItem struct {
	ChildIDs        []int
	Closed          bool
	Comment         string
	ContentObjectID string
	Duration        int
	ID              int
	IsHidden        bool
	IsInternal      bool
	ItemNumber      string
	Level           int
	MeetingID       int
	ParentID        Maybe[int]
	ProjectionIDs   []int
	TagIDs          []int
	Type            string
	Weight          int
	fetch           *Fetch
}

func (c *AgendaItem) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.AgendaItem_ChildIDs(id).Lazy(&c.ChildIDs)
	ds.AgendaItem_Closed(id).Lazy(&c.Closed)
	ds.AgendaItem_Comment(id).Lazy(&c.Comment)
	ds.AgendaItem_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.AgendaItem_Duration(id).Lazy(&c.Duration)
	ds.AgendaItem_ID(id).Lazy(&c.ID)
	ds.AgendaItem_IsHidden(id).Lazy(&c.IsHidden)
	ds.AgendaItem_IsInternal(id).Lazy(&c.IsInternal)
	ds.AgendaItem_ItemNumber(id).Lazy(&c.ItemNumber)
	ds.AgendaItem_Level(id).Lazy(&c.Level)
	ds.AgendaItem_MeetingID(id).Lazy(&c.MeetingID)
	ds.AgendaItem_ParentID(id).Lazy(&c.ParentID)
	ds.AgendaItem_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.AgendaItem_TagIDs(id).Lazy(&c.TagIDs)
	ds.AgendaItem_Type(id).Lazy(&c.Type)
	ds.AgendaItem_Weight(id).Lazy(&c.Weight)
}

func (c *AgendaItem) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.AgendaItem_ChildIDs(id).Preload()
	ds.AgendaItem_Closed(id).Preload()
	ds.AgendaItem_Comment(id).Preload()
	ds.AgendaItem_ContentObjectID(id).Preload()
	ds.AgendaItem_Duration(id).Preload()
	ds.AgendaItem_ID(id).Preload()
	ds.AgendaItem_IsHidden(id).Preload()
	ds.AgendaItem_IsInternal(id).Preload()
	ds.AgendaItem_ItemNumber(id).Preload()
	ds.AgendaItem_Level(id).Preload()
	ds.AgendaItem_MeetingID(id).Preload()
	ds.AgendaItem_ParentID(id).Preload()
	ds.AgendaItem_ProjectionIDs(id).Preload()
	ds.AgendaItem_TagIDs(id).Preload()
	ds.AgendaItem_Type(id).Preload()
	ds.AgendaItem_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) AgendaItem(id int) *ValueCollection[AgendaItem, *AgendaItem] {
	return &ValueCollection[AgendaItem, *AgendaItem]{
		id:    id,
		fetch: r,
	}
}

// Assignment has all fields from assignment.
type Assignment struct {
	AgendaItemID                  Maybe[int]
	AttachmentMeetingMediafileIDs []int
	CandidateIDs                  []int
	DefaultPollDescription        string
	Description                   string
	ID                            int
	ListOfSpeakersID              int
	MeetingID                     int
	NumberPollCandidates          bool
	OpenPosts                     int
	Phase                         string
	PollIDs                       []int
	ProjectionIDs                 []int
	SequentialNumber              int
	TagIDs                        []int
	Title                         string
	fetch                         *Fetch
}

func (c *Assignment) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Assignment_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.Assignment_AttachmentMeetingMediafileIDs(id).Lazy(&c.AttachmentMeetingMediafileIDs)
	ds.Assignment_CandidateIDs(id).Lazy(&c.CandidateIDs)
	ds.Assignment_DefaultPollDescription(id).Lazy(&c.DefaultPollDescription)
	ds.Assignment_Description(id).Lazy(&c.Description)
	ds.Assignment_ID(id).Lazy(&c.ID)
	ds.Assignment_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.Assignment_MeetingID(id).Lazy(&c.MeetingID)
	ds.Assignment_NumberPollCandidates(id).Lazy(&c.NumberPollCandidates)
	ds.Assignment_OpenPosts(id).Lazy(&c.OpenPosts)
	ds.Assignment_Phase(id).Lazy(&c.Phase)
	ds.Assignment_PollIDs(id).Lazy(&c.PollIDs)
	ds.Assignment_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Assignment_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Assignment_TagIDs(id).Lazy(&c.TagIDs)
	ds.Assignment_Title(id).Lazy(&c.Title)
}

func (c *Assignment) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Assignment_AgendaItemID(id).Preload()
	ds.Assignment_AttachmentMeetingMediafileIDs(id).Preload()
	ds.Assignment_CandidateIDs(id).Preload()
	ds.Assignment_DefaultPollDescription(id).Preload()
	ds.Assignment_Description(id).Preload()
	ds.Assignment_ID(id).Preload()
	ds.Assignment_ListOfSpeakersID(id).Preload()
	ds.Assignment_MeetingID(id).Preload()
	ds.Assignment_NumberPollCandidates(id).Preload()
	ds.Assignment_OpenPosts(id).Preload()
	ds.Assignment_Phase(id).Preload()
	ds.Assignment_PollIDs(id).Preload()
	ds.Assignment_ProjectionIDs(id).Preload()
	ds.Assignment_SequentialNumber(id).Preload()
	ds.Assignment_TagIDs(id).Preload()
	ds.Assignment_Title(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Assignment(id int) *ValueCollection[Assignment, *Assignment] {
	return &ValueCollection[Assignment, *Assignment]{
		id:    id,
		fetch: r,
	}
}

// AssignmentCandidate has all fields from assignment_candidate.
type AssignmentCandidate struct {
	AssignmentID  int
	ID            int
	MeetingID     int
	MeetingUserID Maybe[int]
	Weight        int
	fetch         *Fetch
}

func (c *AssignmentCandidate) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.AssignmentCandidate_AssignmentID(id).Lazy(&c.AssignmentID)
	ds.AssignmentCandidate_ID(id).Lazy(&c.ID)
	ds.AssignmentCandidate_MeetingID(id).Lazy(&c.MeetingID)
	ds.AssignmentCandidate_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.AssignmentCandidate_Weight(id).Lazy(&c.Weight)
}

func (c *AssignmentCandidate) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.AssignmentCandidate_AssignmentID(id).Preload()
	ds.AssignmentCandidate_ID(id).Preload()
	ds.AssignmentCandidate_MeetingID(id).Preload()
	ds.AssignmentCandidate_MeetingUserID(id).Preload()
	ds.AssignmentCandidate_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) AssignmentCandidate(id int) *ValueCollection[AssignmentCandidate, *AssignmentCandidate] {
	return &ValueCollection[AssignmentCandidate, *AssignmentCandidate]{
		id:    id,
		fetch: r,
	}
}

// ChatGroup has all fields from chat_group.
type ChatGroup struct {
	ChatMessageIDs []int
	ID             int
	MeetingID      int
	Name           string
	ReadGroupIDs   []int
	Weight         int
	WriteGroupIDs  []int
	fetch          *Fetch
}

func (c *ChatGroup) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ChatGroup_ChatMessageIDs(id).Lazy(&c.ChatMessageIDs)
	ds.ChatGroup_ID(id).Lazy(&c.ID)
	ds.ChatGroup_MeetingID(id).Lazy(&c.MeetingID)
	ds.ChatGroup_Name(id).Lazy(&c.Name)
	ds.ChatGroup_ReadGroupIDs(id).Lazy(&c.ReadGroupIDs)
	ds.ChatGroup_Weight(id).Lazy(&c.Weight)
	ds.ChatGroup_WriteGroupIDs(id).Lazy(&c.WriteGroupIDs)
}

func (c *ChatGroup) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.ChatGroup_ChatMessageIDs(id).Preload()
	ds.ChatGroup_ID(id).Preload()
	ds.ChatGroup_MeetingID(id).Preload()
	ds.ChatGroup_Name(id).Preload()
	ds.ChatGroup_ReadGroupIDs(id).Preload()
	ds.ChatGroup_Weight(id).Preload()
	ds.ChatGroup_WriteGroupIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) ChatGroup(id int) *ValueCollection[ChatGroup, *ChatGroup] {
	return &ValueCollection[ChatGroup, *ChatGroup]{
		id:    id,
		fetch: r,
	}
}

// ChatMessage has all fields from chat_message.
type ChatMessage struct {
	ChatGroupID   int
	Content       string
	Created       int
	ID            int
	MeetingID     int
	MeetingUserID Maybe[int]
	fetch         *Fetch
}

func (c *ChatMessage) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ChatMessage_ChatGroupID(id).Lazy(&c.ChatGroupID)
	ds.ChatMessage_Content(id).Lazy(&c.Content)
	ds.ChatMessage_Created(id).Lazy(&c.Created)
	ds.ChatMessage_ID(id).Lazy(&c.ID)
	ds.ChatMessage_MeetingID(id).Lazy(&c.MeetingID)
	ds.ChatMessage_MeetingUserID(id).Lazy(&c.MeetingUserID)
}

func (c *ChatMessage) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.ChatMessage_ChatGroupID(id).Preload()
	ds.ChatMessage_Content(id).Preload()
	ds.ChatMessage_Created(id).Preload()
	ds.ChatMessage_ID(id).Preload()
	ds.ChatMessage_MeetingID(id).Preload()
	ds.ChatMessage_MeetingUserID(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) ChatMessage(id int) *ValueCollection[ChatMessage, *ChatMessage] {
	return &ValueCollection[ChatMessage, *ChatMessage]{
		id:    id,
		fetch: r,
	}
}

// Committee has all fields from committee.
type Committee struct {
	DefaultMeetingID                   Maybe[int]
	Description                        string
	ExternalID                         string
	ForwardToCommitteeIDs              []int
	ForwardingUserID                   Maybe[int]
	ID                                 int
	ManagerIDs                         []int
	MeetingIDs                         []int
	Name                               string
	OrganizationID                     int
	OrganizationTagIDs                 []int
	ReceiveForwardingsFromCommitteeIDs []int
	UserIDs                            []int
	fetch                              *Fetch
}

func (c *Committee) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Committee_DefaultMeetingID(id).Lazy(&c.DefaultMeetingID)
	ds.Committee_Description(id).Lazy(&c.Description)
	ds.Committee_ExternalID(id).Lazy(&c.ExternalID)
	ds.Committee_ForwardToCommitteeIDs(id).Lazy(&c.ForwardToCommitteeIDs)
	ds.Committee_ForwardingUserID(id).Lazy(&c.ForwardingUserID)
	ds.Committee_ID(id).Lazy(&c.ID)
	ds.Committee_ManagerIDs(id).Lazy(&c.ManagerIDs)
	ds.Committee_MeetingIDs(id).Lazy(&c.MeetingIDs)
	ds.Committee_Name(id).Lazy(&c.Name)
	ds.Committee_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.Committee_OrganizationTagIDs(id).Lazy(&c.OrganizationTagIDs)
	ds.Committee_ReceiveForwardingsFromCommitteeIDs(id).Lazy(&c.ReceiveForwardingsFromCommitteeIDs)
	ds.Committee_UserIDs(id).Lazy(&c.UserIDs)
}

func (c *Committee) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Committee_DefaultMeetingID(id).Preload()
	ds.Committee_Description(id).Preload()
	ds.Committee_ExternalID(id).Preload()
	ds.Committee_ForwardToCommitteeIDs(id).Preload()
	ds.Committee_ForwardingUserID(id).Preload()
	ds.Committee_ID(id).Preload()
	ds.Committee_ManagerIDs(id).Preload()
	ds.Committee_MeetingIDs(id).Preload()
	ds.Committee_Name(id).Preload()
	ds.Committee_OrganizationID(id).Preload()
	ds.Committee_OrganizationTagIDs(id).Preload()
	ds.Committee_ReceiveForwardingsFromCommitteeIDs(id).Preload()
	ds.Committee_UserIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Committee(id int) *ValueCollection[Committee, *Committee] {
	return &ValueCollection[Committee, *Committee]{
		id:    id,
		fetch: r,
	}
}

// Gender has all fields from gender.
type Gender struct {
	ID             int
	Name           string
	OrganizationID int
	UserIDs        []int
	fetch          *Fetch
}

func (c *Gender) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Gender_ID(id).Lazy(&c.ID)
	ds.Gender_Name(id).Lazy(&c.Name)
	ds.Gender_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.Gender_UserIDs(id).Lazy(&c.UserIDs)
}

func (c *Gender) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Gender_ID(id).Preload()
	ds.Gender_Name(id).Preload()
	ds.Gender_OrganizationID(id).Preload()
	ds.Gender_UserIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Gender(id int) *ValueCollection[Gender, *Gender] {
	return &ValueCollection[Gender, *Gender]{
		id:    id,
		fetch: r,
	}
}

// Group has all fields from group.
type Group struct {
	AdminGroupForMeetingID                  Maybe[int]
	AnonymousGroupForMeetingID              Maybe[int]
	DefaultGroupForMeetingID                Maybe[int]
	ExternalID                              string
	ID                                      int
	MeetingID                               int
	MeetingMediafileAccessGroupIDs          []int
	MeetingMediafileInheritedAccessGroupIDs []int
	MeetingUserIDs                          []int
	Name                                    string
	Permissions                             []string
	PollIDs                                 []int
	ReadChatGroupIDs                        []int
	ReadCommentSectionIDs                   []int
	UsedAsAssignmentPollDefaultID           Maybe[int]
	UsedAsMotionPollDefaultID               Maybe[int]
	UsedAsPollDefaultID                     Maybe[int]
	UsedAsTopicPollDefaultID                Maybe[int]
	Weight                                  int
	WriteChatGroupIDs                       []int
	WriteCommentSectionIDs                  []int
	fetch                                   *Fetch
}

func (c *Group) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Group_AdminGroupForMeetingID(id).Lazy(&c.AdminGroupForMeetingID)
	ds.Group_AnonymousGroupForMeetingID(id).Lazy(&c.AnonymousGroupForMeetingID)
	ds.Group_DefaultGroupForMeetingID(id).Lazy(&c.DefaultGroupForMeetingID)
	ds.Group_ExternalID(id).Lazy(&c.ExternalID)
	ds.Group_ID(id).Lazy(&c.ID)
	ds.Group_MeetingID(id).Lazy(&c.MeetingID)
	ds.Group_MeetingMediafileAccessGroupIDs(id).Lazy(&c.MeetingMediafileAccessGroupIDs)
	ds.Group_MeetingMediafileInheritedAccessGroupIDs(id).Lazy(&c.MeetingMediafileInheritedAccessGroupIDs)
	ds.Group_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.Group_Name(id).Lazy(&c.Name)
	ds.Group_Permissions(id).Lazy(&c.Permissions)
	ds.Group_PollIDs(id).Lazy(&c.PollIDs)
	ds.Group_ReadChatGroupIDs(id).Lazy(&c.ReadChatGroupIDs)
	ds.Group_ReadCommentSectionIDs(id).Lazy(&c.ReadCommentSectionIDs)
	ds.Group_UsedAsAssignmentPollDefaultID(id).Lazy(&c.UsedAsAssignmentPollDefaultID)
	ds.Group_UsedAsMotionPollDefaultID(id).Lazy(&c.UsedAsMotionPollDefaultID)
	ds.Group_UsedAsPollDefaultID(id).Lazy(&c.UsedAsPollDefaultID)
	ds.Group_UsedAsTopicPollDefaultID(id).Lazy(&c.UsedAsTopicPollDefaultID)
	ds.Group_Weight(id).Lazy(&c.Weight)
	ds.Group_WriteChatGroupIDs(id).Lazy(&c.WriteChatGroupIDs)
	ds.Group_WriteCommentSectionIDs(id).Lazy(&c.WriteCommentSectionIDs)
}

func (c *Group) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Group_AdminGroupForMeetingID(id).Preload()
	ds.Group_AnonymousGroupForMeetingID(id).Preload()
	ds.Group_DefaultGroupForMeetingID(id).Preload()
	ds.Group_ExternalID(id).Preload()
	ds.Group_ID(id).Preload()
	ds.Group_MeetingID(id).Preload()
	ds.Group_MeetingMediafileAccessGroupIDs(id).Preload()
	ds.Group_MeetingMediafileInheritedAccessGroupIDs(id).Preload()
	ds.Group_MeetingUserIDs(id).Preload()
	ds.Group_Name(id).Preload()
	ds.Group_Permissions(id).Preload()
	ds.Group_PollIDs(id).Preload()
	ds.Group_ReadChatGroupIDs(id).Preload()
	ds.Group_ReadCommentSectionIDs(id).Preload()
	ds.Group_UsedAsAssignmentPollDefaultID(id).Preload()
	ds.Group_UsedAsMotionPollDefaultID(id).Preload()
	ds.Group_UsedAsPollDefaultID(id).Preload()
	ds.Group_UsedAsTopicPollDefaultID(id).Preload()
	ds.Group_Weight(id).Preload()
	ds.Group_WriteChatGroupIDs(id).Preload()
	ds.Group_WriteCommentSectionIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Group(id int) *ValueCollection[Group, *Group] {
	return &ValueCollection[Group, *Group]{
		id:    id,
		fetch: r,
	}
}

// ImportPreview has all fields from import_preview.
type ImportPreview struct {
	Created int
	ID      int
	Name    string
	Result  json.RawMessage
	State   string
	fetch   *Fetch
}

func (c *ImportPreview) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ImportPreview_Created(id).Lazy(&c.Created)
	ds.ImportPreview_ID(id).Lazy(&c.ID)
	ds.ImportPreview_Name(id).Lazy(&c.Name)
	ds.ImportPreview_Result(id).Lazy(&c.Result)
	ds.ImportPreview_State(id).Lazy(&c.State)
}

func (c *ImportPreview) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.ImportPreview_Created(id).Preload()
	ds.ImportPreview_ID(id).Preload()
	ds.ImportPreview_Name(id).Preload()
	ds.ImportPreview_Result(id).Preload()
	ds.ImportPreview_State(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) ImportPreview(id int) *ValueCollection[ImportPreview, *ImportPreview] {
	return &ValueCollection[ImportPreview, *ImportPreview]{
		id:    id,
		fetch: r,
	}
}

// ListOfSpeakers has all fields from list_of_speakers.
type ListOfSpeakers struct {
	Closed                          bool
	ContentObjectID                 string
	ID                              int
	MeetingID                       int
	ModeratorNotes                  string
	ProjectionIDs                   []int
	SequentialNumber                int
	SpeakerIDs                      []int
	StructureLevelListOfSpeakersIDs []int
	fetch                           *Fetch
}

func (c *ListOfSpeakers) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ListOfSpeakers_Closed(id).Lazy(&c.Closed)
	ds.ListOfSpeakers_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.ListOfSpeakers_ID(id).Lazy(&c.ID)
	ds.ListOfSpeakers_MeetingID(id).Lazy(&c.MeetingID)
	ds.ListOfSpeakers_ModeratorNotes(id).Lazy(&c.ModeratorNotes)
	ds.ListOfSpeakers_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.ListOfSpeakers_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.ListOfSpeakers_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.ListOfSpeakers_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
}

func (c *ListOfSpeakers) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.ListOfSpeakers_Closed(id).Preload()
	ds.ListOfSpeakers_ContentObjectID(id).Preload()
	ds.ListOfSpeakers_ID(id).Preload()
	ds.ListOfSpeakers_MeetingID(id).Preload()
	ds.ListOfSpeakers_ModeratorNotes(id).Preload()
	ds.ListOfSpeakers_ProjectionIDs(id).Preload()
	ds.ListOfSpeakers_SequentialNumber(id).Preload()
	ds.ListOfSpeakers_SpeakerIDs(id).Preload()
	ds.ListOfSpeakers_StructureLevelListOfSpeakersIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) ListOfSpeakers(id int) *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	return &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
		id:    id,
		fetch: r,
	}
}

// Mediafile has all fields from mediafile.
type Mediafile struct {
	ChildIDs                            []int
	CreateTimestamp                     int
	Filename                            string
	Filesize                            int
	ID                                  int
	IsDirectory                         bool
	MeetingMediafileIDs                 []int
	Mimetype                            string
	OwnerID                             string
	ParentID                            Maybe[int]
	PdfInformation                      json.RawMessage
	PublishedToMeetingsInOrganizationID Maybe[int]
	Title                               string
	Token                               string
	fetch                               *Fetch
}

func (c *Mediafile) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Mediafile_ChildIDs(id).Lazy(&c.ChildIDs)
	ds.Mediafile_CreateTimestamp(id).Lazy(&c.CreateTimestamp)
	ds.Mediafile_Filename(id).Lazy(&c.Filename)
	ds.Mediafile_Filesize(id).Lazy(&c.Filesize)
	ds.Mediafile_ID(id).Lazy(&c.ID)
	ds.Mediafile_IsDirectory(id).Lazy(&c.IsDirectory)
	ds.Mediafile_MeetingMediafileIDs(id).Lazy(&c.MeetingMediafileIDs)
	ds.Mediafile_Mimetype(id).Lazy(&c.Mimetype)
	ds.Mediafile_OwnerID(id).Lazy(&c.OwnerID)
	ds.Mediafile_ParentID(id).Lazy(&c.ParentID)
	ds.Mediafile_PdfInformation(id).Lazy(&c.PdfInformation)
	ds.Mediafile_PublishedToMeetingsInOrganizationID(id).Lazy(&c.PublishedToMeetingsInOrganizationID)
	ds.Mediafile_Title(id).Lazy(&c.Title)
	ds.Mediafile_Token(id).Lazy(&c.Token)
}

func (c *Mediafile) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Mediafile_ChildIDs(id).Preload()
	ds.Mediafile_CreateTimestamp(id).Preload()
	ds.Mediafile_Filename(id).Preload()
	ds.Mediafile_Filesize(id).Preload()
	ds.Mediafile_ID(id).Preload()
	ds.Mediafile_IsDirectory(id).Preload()
	ds.Mediafile_MeetingMediafileIDs(id).Preload()
	ds.Mediafile_Mimetype(id).Preload()
	ds.Mediafile_OwnerID(id).Preload()
	ds.Mediafile_ParentID(id).Preload()
	ds.Mediafile_PdfInformation(id).Preload()
	ds.Mediafile_PublishedToMeetingsInOrganizationID(id).Preload()
	ds.Mediafile_Title(id).Preload()
	ds.Mediafile_Token(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Mediafile(id int) *ValueCollection[Mediafile, *Mediafile] {
	return &ValueCollection[Mediafile, *Mediafile]{
		id:    id,
		fetch: r,
	}
}

// Meeting has all fields from meeting.
type Meeting struct {
	AdminGroupID                                 Maybe[int]
	AgendaEnableNumbering                        bool
	AgendaItemCreation                           string
	AgendaItemIDs                                []int
	AgendaNewItemsDefaultVisibility              string
	AgendaNumberPrefix                           string
	AgendaNumeralSystem                          string
	AgendaShowInternalItemsOnProjector           bool
	AgendaShowSubtitles                          bool
	AgendaShowTopicNavigationOnDetailView        bool
	AllProjectionIDs                             []int
	AnonymousGroupID                             Maybe[int]
	ApplauseEnable                               bool
	ApplauseMaxAmount                            int
	ApplauseMinAmount                            int
	ApplauseParticleImageUrl                     string
	ApplauseShowLevel                            bool
	ApplauseTimeout                              int
	ApplauseType                                 string
	AssignmentCandidateIDs                       []int
	AssignmentIDs                                []int
	AssignmentPollAddCandidatesToListOfSpeakers  bool
	AssignmentPollBallotPaperNumber              int
	AssignmentPollBallotPaperSelection           string
	AssignmentPollDefaultBackend                 string
	AssignmentPollDefaultGroupIDs                []int
	AssignmentPollDefaultMethod                  string
	AssignmentPollDefaultOnehundredPercentBase   string
	AssignmentPollDefaultType                    string
	AssignmentPollEnableMaxVotesPerOption        bool
	AssignmentPollSortPollResultByVotes          bool
	AssignmentsExportPreamble                    string
	AssignmentsExportTitle                       string
	ChatGroupIDs                                 []int
	ChatMessageIDs                               []int
	CommitteeID                                  int
	ConferenceAutoConnect                        bool
	ConferenceAutoConnectNextSpeakers            int
	ConferenceEnableHelpdesk                     bool
	ConferenceLosRestriction                     bool
	ConferenceOpenMicrophone                     bool
	ConferenceOpenVideo                          bool
	ConferenceShow                               bool
	ConferenceStreamPosterUrl                    string
	ConferenceStreamUrl                          string
	CustomTranslations                           json.RawMessage
	DefaultGroupID                               int
	DefaultMeetingForCommitteeID                 Maybe[int]
	DefaultProjectorAgendaItemListIDs            []int
	DefaultProjectorAmendmentIDs                 []int
	DefaultProjectorAssignmentIDs                []int
	DefaultProjectorAssignmentPollIDs            []int
	DefaultProjectorCountdownIDs                 []int
	DefaultProjectorCurrentListOfSpeakersIDs     []int
	DefaultProjectorListOfSpeakersIDs            []int
	DefaultProjectorMediafileIDs                 []int
	DefaultProjectorMessageIDs                   []int
	DefaultProjectorMotionBlockIDs               []int
	DefaultProjectorMotionIDs                    []int
	DefaultProjectorMotionPollIDs                []int
	DefaultProjectorPollIDs                      []int
	DefaultProjectorTopicIDs                     []int
	Description                                  string
	EnableAnonymous                              bool
	EndTime                                      int
	ExportCsvEncoding                            string
	ExportCsvSeparator                           string
	ExportPdfFontsize                            int
	ExportPdfLineHeight                          float32
	ExportPdfPageMarginBottom                    int
	ExportPdfPageMarginLeft                      int
	ExportPdfPageMarginRight                     int
	ExportPdfPageMarginTop                       int
	ExportPdfPagenumberAlignment                 string
	ExportPdfPagesize                            string
	ExternalID                                   string
	FontBoldID                                   Maybe[int]
	FontBoldItalicID                             Maybe[int]
	FontChyronSpeakerNameID                      Maybe[int]
	FontItalicID                                 Maybe[int]
	FontMonospaceID                              Maybe[int]
	FontProjectorH1ID                            Maybe[int]
	FontProjectorH2ID                            Maybe[int]
	FontRegularID                                Maybe[int]
	ForwardedMotionIDs                           []int
	GroupIDs                                     []int
	ID                                           int
	ImportedAt                                   int
	IsActiveInOrganizationID                     Maybe[int]
	IsArchivedInOrganizationID                   Maybe[int]
	JitsiDomain                                  string
	JitsiRoomName                                string
	JitsiRoomPassword                            string
	Language                                     string
	ListOfSpeakersAllowMultipleSpeakers          bool
	ListOfSpeakersAmountLastOnProjector          int
	ListOfSpeakersAmountNextOnProjector          int
	ListOfSpeakersCanCreatePointOfOrderForOthers bool
	ListOfSpeakersCanSetContributionSelf         bool
	ListOfSpeakersClosingDisablesPointOfOrder    bool
	ListOfSpeakersCountdownID                    Maybe[int]
	ListOfSpeakersCoupleCountdown                bool
	ListOfSpeakersDefaultStructureLevelTime      int
	ListOfSpeakersEnableInterposedQuestion       bool
	ListOfSpeakersEnablePointOfOrderCategories   bool
	ListOfSpeakersEnablePointOfOrderSpeakers     bool
	ListOfSpeakersEnableProContraSpeech          bool
	ListOfSpeakersHideContributionCount          bool
	ListOfSpeakersIDs                            []int
	ListOfSpeakersInitiallyClosed                bool
	ListOfSpeakersInterventionTime               int
	ListOfSpeakersPresentUsersOnly               bool
	ListOfSpeakersShowAmountOfSpeakersOnSlide    bool
	ListOfSpeakersShowFirstContribution          bool
	ListOfSpeakersSpeakerNoteForEveryone         bool
	Location                                     string
	LockedFromInside                             bool
	LogoPdfBallotPaperID                         Maybe[int]
	LogoPdfFooterLID                             Maybe[int]
	LogoPdfFooterRID                             Maybe[int]
	LogoPdfHeaderLID                             Maybe[int]
	LogoPdfHeaderRID                             Maybe[int]
	LogoProjectorHeaderID                        Maybe[int]
	LogoProjectorMainID                          Maybe[int]
	LogoWebHeaderID                              Maybe[int]
	MediafileIDs                                 []int
	MeetingMediafileIDs                          []int
	MeetingUserIDs                               []int
	MotionBlockIDs                               []int
	MotionCategoryIDs                            []int
	MotionChangeRecommendationIDs                []int
	MotionCommentIDs                             []int
	MotionCommentSectionIDs                      []int
	MotionEditorIDs                              []int
	MotionIDs                                    []int
	MotionPollBallotPaperNumber                  int
	MotionPollBallotPaperSelection               string
	MotionPollDefaultBackend                     string
	MotionPollDefaultGroupIDs                    []int
	MotionPollDefaultMethod                      string
	MotionPollDefaultOnehundredPercentBase       string
	MotionPollDefaultType                        string
	MotionStateIDs                               []int
	MotionSubmitterIDs                           []int
	MotionWorkflowIDs                            []int
	MotionWorkingGroupSpeakerIDs                 []int
	MotionsAmendmentsEnabled                     bool
	MotionsAmendmentsInMainList                  bool
	MotionsAmendmentsMultipleParagraphs          bool
	MotionsAmendmentsOfAmendments                bool
	MotionsAmendmentsPrefix                      string
	MotionsAmendmentsTextMode                    string
	MotionsBlockSlideColumns                     int
	MotionsCreateEnableAdditionalSubmitterText   bool
	MotionsDefaultAmendmentWorkflowID            int
	MotionsDefaultLineNumbering                  string
	MotionsDefaultSorting                        string
	MotionsDefaultWorkflowID                     int
	MotionsEnableEditor                          bool
	MotionsEnableReasonOnProjector               bool
	MotionsEnableRecommendationOnProjector       bool
	MotionsEnableSideboxOnProjector              bool
	MotionsEnableTextOnProjector                 bool
	MotionsEnableWorkingGroupSpeaker             bool
	MotionsExportFollowRecommendation            bool
	MotionsExportPreamble                        string
	MotionsExportSubmitterRecommendation         bool
	MotionsExportTitle                           string
	MotionsHideMetadataBackground                bool
	MotionsLineLength                            int
	MotionsNumberMinDigits                       int
	MotionsNumberType                            string
	MotionsNumberWithBlank                       bool
	MotionsPreamble                              string
	MotionsReasonRequired                        bool
	MotionsRecommendationTextMode                string
	MotionsRecommendationsBy                     string
	MotionsShowReferringMotions                  bool
	MotionsShowSequentialNumber                  bool
	MotionsSupportersMinAmount                   int
	Name                                         string
	OptionIDs                                    []int
	OrganizationTagIDs                           []int
	PersonalNoteIDs                              []int
	PointOfOrderCategoryIDs                      []int
	PollBallotPaperNumber                        int
	PollBallotPaperSelection                     string
	PollCandidateIDs                             []int
	PollCandidateListIDs                         []int
	PollCountdownID                              Maybe[int]
	PollCoupleCountdown                          bool
	PollDefaultBackend                           string
	PollDefaultGroupIDs                          []int
	PollDefaultMethod                            string
	PollDefaultOnehundredPercentBase             string
	PollDefaultType                              string
	PollIDs                                      []int
	PollSortPollResultByVotes                    bool
	PresentUserIDs                               []int
	ProjectionIDs                                []int
	ProjectorCountdownDefaultTime                int
	ProjectorCountdownIDs                        []int
	ProjectorCountdownWarningTime                int
	ProjectorIDs                                 []int
	ProjectorMessageIDs                          []int
	ReferenceProjectorID                         int
	SpeakerIDs                                   []int
	StartTime                                    int
	StructureLevelIDs                            []int
	StructureLevelListOfSpeakersIDs              []int
	TagIDs                                       []int
	TemplateForOrganizationID                    Maybe[int]
	TopicIDs                                     []int
	TopicPollDefaultGroupIDs                     []int
	UserIDs                                      []int
	UsersAllowSelfSetPresent                     bool
	UsersEmailBody                               string
	UsersEmailReplyto                            string
	UsersEmailSender                             string
	UsersEmailSubject                            string
	UsersEnablePresenceView                      bool
	UsersEnableVoteDelegations                   bool
	UsersEnableVoteWeight                        bool
	UsersForbidDelegatorAsSubmitter              bool
	UsersForbidDelegatorAsSupporter              bool
	UsersForbidDelegatorInListOfSpeakers         bool
	UsersForbidDelegatorToVote                   bool
	UsersPdfWelcometext                          string
	UsersPdfWelcometitle                         string
	UsersPdfWlanEncryption                       string
	UsersPdfWlanPassword                         string
	UsersPdfWlanSsid                             string
	VoteIDs                                      []int
	WelcomeText                                  string
	WelcomeTitle                                 string
	fetch                                        *Fetch
}

func (c *Meeting) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Meeting_AdminGroupID(id).Lazy(&c.AdminGroupID)
	ds.Meeting_AgendaEnableNumbering(id).Lazy(&c.AgendaEnableNumbering)
	ds.Meeting_AgendaItemCreation(id).Lazy(&c.AgendaItemCreation)
	ds.Meeting_AgendaItemIDs(id).Lazy(&c.AgendaItemIDs)
	ds.Meeting_AgendaNewItemsDefaultVisibility(id).Lazy(&c.AgendaNewItemsDefaultVisibility)
	ds.Meeting_AgendaNumberPrefix(id).Lazy(&c.AgendaNumberPrefix)
	ds.Meeting_AgendaNumeralSystem(id).Lazy(&c.AgendaNumeralSystem)
	ds.Meeting_AgendaShowInternalItemsOnProjector(id).Lazy(&c.AgendaShowInternalItemsOnProjector)
	ds.Meeting_AgendaShowSubtitles(id).Lazy(&c.AgendaShowSubtitles)
	ds.Meeting_AgendaShowTopicNavigationOnDetailView(id).Lazy(&c.AgendaShowTopicNavigationOnDetailView)
	ds.Meeting_AllProjectionIDs(id).Lazy(&c.AllProjectionIDs)
	ds.Meeting_AnonymousGroupID(id).Lazy(&c.AnonymousGroupID)
	ds.Meeting_ApplauseEnable(id).Lazy(&c.ApplauseEnable)
	ds.Meeting_ApplauseMaxAmount(id).Lazy(&c.ApplauseMaxAmount)
	ds.Meeting_ApplauseMinAmount(id).Lazy(&c.ApplauseMinAmount)
	ds.Meeting_ApplauseParticleImageUrl(id).Lazy(&c.ApplauseParticleImageUrl)
	ds.Meeting_ApplauseShowLevel(id).Lazy(&c.ApplauseShowLevel)
	ds.Meeting_ApplauseTimeout(id).Lazy(&c.ApplauseTimeout)
	ds.Meeting_ApplauseType(id).Lazy(&c.ApplauseType)
	ds.Meeting_AssignmentCandidateIDs(id).Lazy(&c.AssignmentCandidateIDs)
	ds.Meeting_AssignmentIDs(id).Lazy(&c.AssignmentIDs)
	ds.Meeting_AssignmentPollAddCandidatesToListOfSpeakers(id).Lazy(&c.AssignmentPollAddCandidatesToListOfSpeakers)
	ds.Meeting_AssignmentPollBallotPaperNumber(id).Lazy(&c.AssignmentPollBallotPaperNumber)
	ds.Meeting_AssignmentPollBallotPaperSelection(id).Lazy(&c.AssignmentPollBallotPaperSelection)
	ds.Meeting_AssignmentPollDefaultBackend(id).Lazy(&c.AssignmentPollDefaultBackend)
	ds.Meeting_AssignmentPollDefaultGroupIDs(id).Lazy(&c.AssignmentPollDefaultGroupIDs)
	ds.Meeting_AssignmentPollDefaultMethod(id).Lazy(&c.AssignmentPollDefaultMethod)
	ds.Meeting_AssignmentPollDefaultOnehundredPercentBase(id).Lazy(&c.AssignmentPollDefaultOnehundredPercentBase)
	ds.Meeting_AssignmentPollDefaultType(id).Lazy(&c.AssignmentPollDefaultType)
	ds.Meeting_AssignmentPollEnableMaxVotesPerOption(id).Lazy(&c.AssignmentPollEnableMaxVotesPerOption)
	ds.Meeting_AssignmentPollSortPollResultByVotes(id).Lazy(&c.AssignmentPollSortPollResultByVotes)
	ds.Meeting_AssignmentsExportPreamble(id).Lazy(&c.AssignmentsExportPreamble)
	ds.Meeting_AssignmentsExportTitle(id).Lazy(&c.AssignmentsExportTitle)
	ds.Meeting_ChatGroupIDs(id).Lazy(&c.ChatGroupIDs)
	ds.Meeting_ChatMessageIDs(id).Lazy(&c.ChatMessageIDs)
	ds.Meeting_CommitteeID(id).Lazy(&c.CommitteeID)
	ds.Meeting_ConferenceAutoConnect(id).Lazy(&c.ConferenceAutoConnect)
	ds.Meeting_ConferenceAutoConnectNextSpeakers(id).Lazy(&c.ConferenceAutoConnectNextSpeakers)
	ds.Meeting_ConferenceEnableHelpdesk(id).Lazy(&c.ConferenceEnableHelpdesk)
	ds.Meeting_ConferenceLosRestriction(id).Lazy(&c.ConferenceLosRestriction)
	ds.Meeting_ConferenceOpenMicrophone(id).Lazy(&c.ConferenceOpenMicrophone)
	ds.Meeting_ConferenceOpenVideo(id).Lazy(&c.ConferenceOpenVideo)
	ds.Meeting_ConferenceShow(id).Lazy(&c.ConferenceShow)
	ds.Meeting_ConferenceStreamPosterUrl(id).Lazy(&c.ConferenceStreamPosterUrl)
	ds.Meeting_ConferenceStreamUrl(id).Lazy(&c.ConferenceStreamUrl)
	ds.Meeting_CustomTranslations(id).Lazy(&c.CustomTranslations)
	ds.Meeting_DefaultGroupID(id).Lazy(&c.DefaultGroupID)
	ds.Meeting_DefaultMeetingForCommitteeID(id).Lazy(&c.DefaultMeetingForCommitteeID)
	ds.Meeting_DefaultProjectorAgendaItemListIDs(id).Lazy(&c.DefaultProjectorAgendaItemListIDs)
	ds.Meeting_DefaultProjectorAmendmentIDs(id).Lazy(&c.DefaultProjectorAmendmentIDs)
	ds.Meeting_DefaultProjectorAssignmentIDs(id).Lazy(&c.DefaultProjectorAssignmentIDs)
	ds.Meeting_DefaultProjectorAssignmentPollIDs(id).Lazy(&c.DefaultProjectorAssignmentPollIDs)
	ds.Meeting_DefaultProjectorCountdownIDs(id).Lazy(&c.DefaultProjectorCountdownIDs)
	ds.Meeting_DefaultProjectorCurrentListOfSpeakersIDs(id).Lazy(&c.DefaultProjectorCurrentListOfSpeakersIDs)
	ds.Meeting_DefaultProjectorListOfSpeakersIDs(id).Lazy(&c.DefaultProjectorListOfSpeakersIDs)
	ds.Meeting_DefaultProjectorMediafileIDs(id).Lazy(&c.DefaultProjectorMediafileIDs)
	ds.Meeting_DefaultProjectorMessageIDs(id).Lazy(&c.DefaultProjectorMessageIDs)
	ds.Meeting_DefaultProjectorMotionBlockIDs(id).Lazy(&c.DefaultProjectorMotionBlockIDs)
	ds.Meeting_DefaultProjectorMotionIDs(id).Lazy(&c.DefaultProjectorMotionIDs)
	ds.Meeting_DefaultProjectorMotionPollIDs(id).Lazy(&c.DefaultProjectorMotionPollIDs)
	ds.Meeting_DefaultProjectorPollIDs(id).Lazy(&c.DefaultProjectorPollIDs)
	ds.Meeting_DefaultProjectorTopicIDs(id).Lazy(&c.DefaultProjectorTopicIDs)
	ds.Meeting_Description(id).Lazy(&c.Description)
	ds.Meeting_EnableAnonymous(id).Lazy(&c.EnableAnonymous)
	ds.Meeting_EndTime(id).Lazy(&c.EndTime)
	ds.Meeting_ExportCsvEncoding(id).Lazy(&c.ExportCsvEncoding)
	ds.Meeting_ExportCsvSeparator(id).Lazy(&c.ExportCsvSeparator)
	ds.Meeting_ExportPdfFontsize(id).Lazy(&c.ExportPdfFontsize)
	ds.Meeting_ExportPdfLineHeight(id).Lazy(&c.ExportPdfLineHeight)
	ds.Meeting_ExportPdfPageMarginBottom(id).Lazy(&c.ExportPdfPageMarginBottom)
	ds.Meeting_ExportPdfPageMarginLeft(id).Lazy(&c.ExportPdfPageMarginLeft)
	ds.Meeting_ExportPdfPageMarginRight(id).Lazy(&c.ExportPdfPageMarginRight)
	ds.Meeting_ExportPdfPageMarginTop(id).Lazy(&c.ExportPdfPageMarginTop)
	ds.Meeting_ExportPdfPagenumberAlignment(id).Lazy(&c.ExportPdfPagenumberAlignment)
	ds.Meeting_ExportPdfPagesize(id).Lazy(&c.ExportPdfPagesize)
	ds.Meeting_ExternalID(id).Lazy(&c.ExternalID)
	ds.Meeting_FontBoldID(id).Lazy(&c.FontBoldID)
	ds.Meeting_FontBoldItalicID(id).Lazy(&c.FontBoldItalicID)
	ds.Meeting_FontChyronSpeakerNameID(id).Lazy(&c.FontChyronSpeakerNameID)
	ds.Meeting_FontItalicID(id).Lazy(&c.FontItalicID)
	ds.Meeting_FontMonospaceID(id).Lazy(&c.FontMonospaceID)
	ds.Meeting_FontProjectorH1ID(id).Lazy(&c.FontProjectorH1ID)
	ds.Meeting_FontProjectorH2ID(id).Lazy(&c.FontProjectorH2ID)
	ds.Meeting_FontRegularID(id).Lazy(&c.FontRegularID)
	ds.Meeting_ForwardedMotionIDs(id).Lazy(&c.ForwardedMotionIDs)
	ds.Meeting_GroupIDs(id).Lazy(&c.GroupIDs)
	ds.Meeting_ID(id).Lazy(&c.ID)
	ds.Meeting_ImportedAt(id).Lazy(&c.ImportedAt)
	ds.Meeting_IsActiveInOrganizationID(id).Lazy(&c.IsActiveInOrganizationID)
	ds.Meeting_IsArchivedInOrganizationID(id).Lazy(&c.IsArchivedInOrganizationID)
	ds.Meeting_JitsiDomain(id).Lazy(&c.JitsiDomain)
	ds.Meeting_JitsiRoomName(id).Lazy(&c.JitsiRoomName)
	ds.Meeting_JitsiRoomPassword(id).Lazy(&c.JitsiRoomPassword)
	ds.Meeting_Language(id).Lazy(&c.Language)
	ds.Meeting_ListOfSpeakersAllowMultipleSpeakers(id).Lazy(&c.ListOfSpeakersAllowMultipleSpeakers)
	ds.Meeting_ListOfSpeakersAmountLastOnProjector(id).Lazy(&c.ListOfSpeakersAmountLastOnProjector)
	ds.Meeting_ListOfSpeakersAmountNextOnProjector(id).Lazy(&c.ListOfSpeakersAmountNextOnProjector)
	ds.Meeting_ListOfSpeakersCanCreatePointOfOrderForOthers(id).Lazy(&c.ListOfSpeakersCanCreatePointOfOrderForOthers)
	ds.Meeting_ListOfSpeakersCanSetContributionSelf(id).Lazy(&c.ListOfSpeakersCanSetContributionSelf)
	ds.Meeting_ListOfSpeakersClosingDisablesPointOfOrder(id).Lazy(&c.ListOfSpeakersClosingDisablesPointOfOrder)
	ds.Meeting_ListOfSpeakersCountdownID(id).Lazy(&c.ListOfSpeakersCountdownID)
	ds.Meeting_ListOfSpeakersCoupleCountdown(id).Lazy(&c.ListOfSpeakersCoupleCountdown)
	ds.Meeting_ListOfSpeakersDefaultStructureLevelTime(id).Lazy(&c.ListOfSpeakersDefaultStructureLevelTime)
	ds.Meeting_ListOfSpeakersEnableInterposedQuestion(id).Lazy(&c.ListOfSpeakersEnableInterposedQuestion)
	ds.Meeting_ListOfSpeakersEnablePointOfOrderCategories(id).Lazy(&c.ListOfSpeakersEnablePointOfOrderCategories)
	ds.Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(id).Lazy(&c.ListOfSpeakersEnablePointOfOrderSpeakers)
	ds.Meeting_ListOfSpeakersEnableProContraSpeech(id).Lazy(&c.ListOfSpeakersEnableProContraSpeech)
	ds.Meeting_ListOfSpeakersHideContributionCount(id).Lazy(&c.ListOfSpeakersHideContributionCount)
	ds.Meeting_ListOfSpeakersIDs(id).Lazy(&c.ListOfSpeakersIDs)
	ds.Meeting_ListOfSpeakersInitiallyClosed(id).Lazy(&c.ListOfSpeakersInitiallyClosed)
	ds.Meeting_ListOfSpeakersInterventionTime(id).Lazy(&c.ListOfSpeakersInterventionTime)
	ds.Meeting_ListOfSpeakersPresentUsersOnly(id).Lazy(&c.ListOfSpeakersPresentUsersOnly)
	ds.Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(id).Lazy(&c.ListOfSpeakersShowAmountOfSpeakersOnSlide)
	ds.Meeting_ListOfSpeakersShowFirstContribution(id).Lazy(&c.ListOfSpeakersShowFirstContribution)
	ds.Meeting_ListOfSpeakersSpeakerNoteForEveryone(id).Lazy(&c.ListOfSpeakersSpeakerNoteForEveryone)
	ds.Meeting_Location(id).Lazy(&c.Location)
	ds.Meeting_LockedFromInside(id).Lazy(&c.LockedFromInside)
	ds.Meeting_LogoPdfBallotPaperID(id).Lazy(&c.LogoPdfBallotPaperID)
	ds.Meeting_LogoPdfFooterLID(id).Lazy(&c.LogoPdfFooterLID)
	ds.Meeting_LogoPdfFooterRID(id).Lazy(&c.LogoPdfFooterRID)
	ds.Meeting_LogoPdfHeaderLID(id).Lazy(&c.LogoPdfHeaderLID)
	ds.Meeting_LogoPdfHeaderRID(id).Lazy(&c.LogoPdfHeaderRID)
	ds.Meeting_LogoProjectorHeaderID(id).Lazy(&c.LogoProjectorHeaderID)
	ds.Meeting_LogoProjectorMainID(id).Lazy(&c.LogoProjectorMainID)
	ds.Meeting_LogoWebHeaderID(id).Lazy(&c.LogoWebHeaderID)
	ds.Meeting_MediafileIDs(id).Lazy(&c.MediafileIDs)
	ds.Meeting_MeetingMediafileIDs(id).Lazy(&c.MeetingMediafileIDs)
	ds.Meeting_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.Meeting_MotionBlockIDs(id).Lazy(&c.MotionBlockIDs)
	ds.Meeting_MotionCategoryIDs(id).Lazy(&c.MotionCategoryIDs)
	ds.Meeting_MotionChangeRecommendationIDs(id).Lazy(&c.MotionChangeRecommendationIDs)
	ds.Meeting_MotionCommentIDs(id).Lazy(&c.MotionCommentIDs)
	ds.Meeting_MotionCommentSectionIDs(id).Lazy(&c.MotionCommentSectionIDs)
	ds.Meeting_MotionEditorIDs(id).Lazy(&c.MotionEditorIDs)
	ds.Meeting_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.Meeting_MotionPollBallotPaperNumber(id).Lazy(&c.MotionPollBallotPaperNumber)
	ds.Meeting_MotionPollBallotPaperSelection(id).Lazy(&c.MotionPollBallotPaperSelection)
	ds.Meeting_MotionPollDefaultBackend(id).Lazy(&c.MotionPollDefaultBackend)
	ds.Meeting_MotionPollDefaultGroupIDs(id).Lazy(&c.MotionPollDefaultGroupIDs)
	ds.Meeting_MotionPollDefaultMethod(id).Lazy(&c.MotionPollDefaultMethod)
	ds.Meeting_MotionPollDefaultOnehundredPercentBase(id).Lazy(&c.MotionPollDefaultOnehundredPercentBase)
	ds.Meeting_MotionPollDefaultType(id).Lazy(&c.MotionPollDefaultType)
	ds.Meeting_MotionStateIDs(id).Lazy(&c.MotionStateIDs)
	ds.Meeting_MotionSubmitterIDs(id).Lazy(&c.MotionSubmitterIDs)
	ds.Meeting_MotionWorkflowIDs(id).Lazy(&c.MotionWorkflowIDs)
	ds.Meeting_MotionWorkingGroupSpeakerIDs(id).Lazy(&c.MotionWorkingGroupSpeakerIDs)
	ds.Meeting_MotionsAmendmentsEnabled(id).Lazy(&c.MotionsAmendmentsEnabled)
	ds.Meeting_MotionsAmendmentsInMainList(id).Lazy(&c.MotionsAmendmentsInMainList)
	ds.Meeting_MotionsAmendmentsMultipleParagraphs(id).Lazy(&c.MotionsAmendmentsMultipleParagraphs)
	ds.Meeting_MotionsAmendmentsOfAmendments(id).Lazy(&c.MotionsAmendmentsOfAmendments)
	ds.Meeting_MotionsAmendmentsPrefix(id).Lazy(&c.MotionsAmendmentsPrefix)
	ds.Meeting_MotionsAmendmentsTextMode(id).Lazy(&c.MotionsAmendmentsTextMode)
	ds.Meeting_MotionsBlockSlideColumns(id).Lazy(&c.MotionsBlockSlideColumns)
	ds.Meeting_MotionsCreateEnableAdditionalSubmitterText(id).Lazy(&c.MotionsCreateEnableAdditionalSubmitterText)
	ds.Meeting_MotionsDefaultAmendmentWorkflowID(id).Lazy(&c.MotionsDefaultAmendmentWorkflowID)
	ds.Meeting_MotionsDefaultLineNumbering(id).Lazy(&c.MotionsDefaultLineNumbering)
	ds.Meeting_MotionsDefaultSorting(id).Lazy(&c.MotionsDefaultSorting)
	ds.Meeting_MotionsDefaultWorkflowID(id).Lazy(&c.MotionsDefaultWorkflowID)
	ds.Meeting_MotionsEnableEditor(id).Lazy(&c.MotionsEnableEditor)
	ds.Meeting_MotionsEnableReasonOnProjector(id).Lazy(&c.MotionsEnableReasonOnProjector)
	ds.Meeting_MotionsEnableRecommendationOnProjector(id).Lazy(&c.MotionsEnableRecommendationOnProjector)
	ds.Meeting_MotionsEnableSideboxOnProjector(id).Lazy(&c.MotionsEnableSideboxOnProjector)
	ds.Meeting_MotionsEnableTextOnProjector(id).Lazy(&c.MotionsEnableTextOnProjector)
	ds.Meeting_MotionsEnableWorkingGroupSpeaker(id).Lazy(&c.MotionsEnableWorkingGroupSpeaker)
	ds.Meeting_MotionsExportFollowRecommendation(id).Lazy(&c.MotionsExportFollowRecommendation)
	ds.Meeting_MotionsExportPreamble(id).Lazy(&c.MotionsExportPreamble)
	ds.Meeting_MotionsExportSubmitterRecommendation(id).Lazy(&c.MotionsExportSubmitterRecommendation)
	ds.Meeting_MotionsExportTitle(id).Lazy(&c.MotionsExportTitle)
	ds.Meeting_MotionsHideMetadataBackground(id).Lazy(&c.MotionsHideMetadataBackground)
	ds.Meeting_MotionsLineLength(id).Lazy(&c.MotionsLineLength)
	ds.Meeting_MotionsNumberMinDigits(id).Lazy(&c.MotionsNumberMinDigits)
	ds.Meeting_MotionsNumberType(id).Lazy(&c.MotionsNumberType)
	ds.Meeting_MotionsNumberWithBlank(id).Lazy(&c.MotionsNumberWithBlank)
	ds.Meeting_MotionsPreamble(id).Lazy(&c.MotionsPreamble)
	ds.Meeting_MotionsReasonRequired(id).Lazy(&c.MotionsReasonRequired)
	ds.Meeting_MotionsRecommendationTextMode(id).Lazy(&c.MotionsRecommendationTextMode)
	ds.Meeting_MotionsRecommendationsBy(id).Lazy(&c.MotionsRecommendationsBy)
	ds.Meeting_MotionsShowReferringMotions(id).Lazy(&c.MotionsShowReferringMotions)
	ds.Meeting_MotionsShowSequentialNumber(id).Lazy(&c.MotionsShowSequentialNumber)
	ds.Meeting_MotionsSupportersMinAmount(id).Lazy(&c.MotionsSupportersMinAmount)
	ds.Meeting_Name(id).Lazy(&c.Name)
	ds.Meeting_OptionIDs(id).Lazy(&c.OptionIDs)
	ds.Meeting_OrganizationTagIDs(id).Lazy(&c.OrganizationTagIDs)
	ds.Meeting_PersonalNoteIDs(id).Lazy(&c.PersonalNoteIDs)
	ds.Meeting_PointOfOrderCategoryIDs(id).Lazy(&c.PointOfOrderCategoryIDs)
	ds.Meeting_PollBallotPaperNumber(id).Lazy(&c.PollBallotPaperNumber)
	ds.Meeting_PollBallotPaperSelection(id).Lazy(&c.PollBallotPaperSelection)
	ds.Meeting_PollCandidateIDs(id).Lazy(&c.PollCandidateIDs)
	ds.Meeting_PollCandidateListIDs(id).Lazy(&c.PollCandidateListIDs)
	ds.Meeting_PollCountdownID(id).Lazy(&c.PollCountdownID)
	ds.Meeting_PollCoupleCountdown(id).Lazy(&c.PollCoupleCountdown)
	ds.Meeting_PollDefaultBackend(id).Lazy(&c.PollDefaultBackend)
	ds.Meeting_PollDefaultGroupIDs(id).Lazy(&c.PollDefaultGroupIDs)
	ds.Meeting_PollDefaultMethod(id).Lazy(&c.PollDefaultMethod)
	ds.Meeting_PollDefaultOnehundredPercentBase(id).Lazy(&c.PollDefaultOnehundredPercentBase)
	ds.Meeting_PollDefaultType(id).Lazy(&c.PollDefaultType)
	ds.Meeting_PollIDs(id).Lazy(&c.PollIDs)
	ds.Meeting_PollSortPollResultByVotes(id).Lazy(&c.PollSortPollResultByVotes)
	ds.Meeting_PresentUserIDs(id).Lazy(&c.PresentUserIDs)
	ds.Meeting_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Meeting_ProjectorCountdownDefaultTime(id).Lazy(&c.ProjectorCountdownDefaultTime)
	ds.Meeting_ProjectorCountdownIDs(id).Lazy(&c.ProjectorCountdownIDs)
	ds.Meeting_ProjectorCountdownWarningTime(id).Lazy(&c.ProjectorCountdownWarningTime)
	ds.Meeting_ProjectorIDs(id).Lazy(&c.ProjectorIDs)
	ds.Meeting_ProjectorMessageIDs(id).Lazy(&c.ProjectorMessageIDs)
	ds.Meeting_ReferenceProjectorID(id).Lazy(&c.ReferenceProjectorID)
	ds.Meeting_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.Meeting_StartTime(id).Lazy(&c.StartTime)
	ds.Meeting_StructureLevelIDs(id).Lazy(&c.StructureLevelIDs)
	ds.Meeting_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
	ds.Meeting_TagIDs(id).Lazy(&c.TagIDs)
	ds.Meeting_TemplateForOrganizationID(id).Lazy(&c.TemplateForOrganizationID)
	ds.Meeting_TopicIDs(id).Lazy(&c.TopicIDs)
	ds.Meeting_TopicPollDefaultGroupIDs(id).Lazy(&c.TopicPollDefaultGroupIDs)
	ds.Meeting_UserIDs(id).Lazy(&c.UserIDs)
	ds.Meeting_UsersAllowSelfSetPresent(id).Lazy(&c.UsersAllowSelfSetPresent)
	ds.Meeting_UsersEmailBody(id).Lazy(&c.UsersEmailBody)
	ds.Meeting_UsersEmailReplyto(id).Lazy(&c.UsersEmailReplyto)
	ds.Meeting_UsersEmailSender(id).Lazy(&c.UsersEmailSender)
	ds.Meeting_UsersEmailSubject(id).Lazy(&c.UsersEmailSubject)
	ds.Meeting_UsersEnablePresenceView(id).Lazy(&c.UsersEnablePresenceView)
	ds.Meeting_UsersEnableVoteDelegations(id).Lazy(&c.UsersEnableVoteDelegations)
	ds.Meeting_UsersEnableVoteWeight(id).Lazy(&c.UsersEnableVoteWeight)
	ds.Meeting_UsersForbidDelegatorAsSubmitter(id).Lazy(&c.UsersForbidDelegatorAsSubmitter)
	ds.Meeting_UsersForbidDelegatorAsSupporter(id).Lazy(&c.UsersForbidDelegatorAsSupporter)
	ds.Meeting_UsersForbidDelegatorInListOfSpeakers(id).Lazy(&c.UsersForbidDelegatorInListOfSpeakers)
	ds.Meeting_UsersForbidDelegatorToVote(id).Lazy(&c.UsersForbidDelegatorToVote)
	ds.Meeting_UsersPdfWelcometext(id).Lazy(&c.UsersPdfWelcometext)
	ds.Meeting_UsersPdfWelcometitle(id).Lazy(&c.UsersPdfWelcometitle)
	ds.Meeting_UsersPdfWlanEncryption(id).Lazy(&c.UsersPdfWlanEncryption)
	ds.Meeting_UsersPdfWlanPassword(id).Lazy(&c.UsersPdfWlanPassword)
	ds.Meeting_UsersPdfWlanSsid(id).Lazy(&c.UsersPdfWlanSsid)
	ds.Meeting_VoteIDs(id).Lazy(&c.VoteIDs)
	ds.Meeting_WelcomeText(id).Lazy(&c.WelcomeText)
	ds.Meeting_WelcomeTitle(id).Lazy(&c.WelcomeTitle)
}

func (c *Meeting) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Meeting_AdminGroupID(id).Preload()
	ds.Meeting_AgendaEnableNumbering(id).Preload()
	ds.Meeting_AgendaItemCreation(id).Preload()
	ds.Meeting_AgendaItemIDs(id).Preload()
	ds.Meeting_AgendaNewItemsDefaultVisibility(id).Preload()
	ds.Meeting_AgendaNumberPrefix(id).Preload()
	ds.Meeting_AgendaNumeralSystem(id).Preload()
	ds.Meeting_AgendaShowInternalItemsOnProjector(id).Preload()
	ds.Meeting_AgendaShowSubtitles(id).Preload()
	ds.Meeting_AgendaShowTopicNavigationOnDetailView(id).Preload()
	ds.Meeting_AllProjectionIDs(id).Preload()
	ds.Meeting_AnonymousGroupID(id).Preload()
	ds.Meeting_ApplauseEnable(id).Preload()
	ds.Meeting_ApplauseMaxAmount(id).Preload()
	ds.Meeting_ApplauseMinAmount(id).Preload()
	ds.Meeting_ApplauseParticleImageUrl(id).Preload()
	ds.Meeting_ApplauseShowLevel(id).Preload()
	ds.Meeting_ApplauseTimeout(id).Preload()
	ds.Meeting_ApplauseType(id).Preload()
	ds.Meeting_AssignmentCandidateIDs(id).Preload()
	ds.Meeting_AssignmentIDs(id).Preload()
	ds.Meeting_AssignmentPollAddCandidatesToListOfSpeakers(id).Preload()
	ds.Meeting_AssignmentPollBallotPaperNumber(id).Preload()
	ds.Meeting_AssignmentPollBallotPaperSelection(id).Preload()
	ds.Meeting_AssignmentPollDefaultBackend(id).Preload()
	ds.Meeting_AssignmentPollDefaultGroupIDs(id).Preload()
	ds.Meeting_AssignmentPollDefaultMethod(id).Preload()
	ds.Meeting_AssignmentPollDefaultOnehundredPercentBase(id).Preload()
	ds.Meeting_AssignmentPollDefaultType(id).Preload()
	ds.Meeting_AssignmentPollEnableMaxVotesPerOption(id).Preload()
	ds.Meeting_AssignmentPollSortPollResultByVotes(id).Preload()
	ds.Meeting_AssignmentsExportPreamble(id).Preload()
	ds.Meeting_AssignmentsExportTitle(id).Preload()
	ds.Meeting_ChatGroupIDs(id).Preload()
	ds.Meeting_ChatMessageIDs(id).Preload()
	ds.Meeting_CommitteeID(id).Preload()
	ds.Meeting_ConferenceAutoConnect(id).Preload()
	ds.Meeting_ConferenceAutoConnectNextSpeakers(id).Preload()
	ds.Meeting_ConferenceEnableHelpdesk(id).Preload()
	ds.Meeting_ConferenceLosRestriction(id).Preload()
	ds.Meeting_ConferenceOpenMicrophone(id).Preload()
	ds.Meeting_ConferenceOpenVideo(id).Preload()
	ds.Meeting_ConferenceShow(id).Preload()
	ds.Meeting_ConferenceStreamPosterUrl(id).Preload()
	ds.Meeting_ConferenceStreamUrl(id).Preload()
	ds.Meeting_CustomTranslations(id).Preload()
	ds.Meeting_DefaultGroupID(id).Preload()
	ds.Meeting_DefaultMeetingForCommitteeID(id).Preload()
	ds.Meeting_DefaultProjectorAgendaItemListIDs(id).Preload()
	ds.Meeting_DefaultProjectorAmendmentIDs(id).Preload()
	ds.Meeting_DefaultProjectorAssignmentIDs(id).Preload()
	ds.Meeting_DefaultProjectorAssignmentPollIDs(id).Preload()
	ds.Meeting_DefaultProjectorCountdownIDs(id).Preload()
	ds.Meeting_DefaultProjectorCurrentListOfSpeakersIDs(id).Preload()
	ds.Meeting_DefaultProjectorListOfSpeakersIDs(id).Preload()
	ds.Meeting_DefaultProjectorMediafileIDs(id).Preload()
	ds.Meeting_DefaultProjectorMessageIDs(id).Preload()
	ds.Meeting_DefaultProjectorMotionBlockIDs(id).Preload()
	ds.Meeting_DefaultProjectorMotionIDs(id).Preload()
	ds.Meeting_DefaultProjectorMotionPollIDs(id).Preload()
	ds.Meeting_DefaultProjectorPollIDs(id).Preload()
	ds.Meeting_DefaultProjectorTopicIDs(id).Preload()
	ds.Meeting_Description(id).Preload()
	ds.Meeting_EnableAnonymous(id).Preload()
	ds.Meeting_EndTime(id).Preload()
	ds.Meeting_ExportCsvEncoding(id).Preload()
	ds.Meeting_ExportCsvSeparator(id).Preload()
	ds.Meeting_ExportPdfFontsize(id).Preload()
	ds.Meeting_ExportPdfLineHeight(id).Preload()
	ds.Meeting_ExportPdfPageMarginBottom(id).Preload()
	ds.Meeting_ExportPdfPageMarginLeft(id).Preload()
	ds.Meeting_ExportPdfPageMarginRight(id).Preload()
	ds.Meeting_ExportPdfPageMarginTop(id).Preload()
	ds.Meeting_ExportPdfPagenumberAlignment(id).Preload()
	ds.Meeting_ExportPdfPagesize(id).Preload()
	ds.Meeting_ExternalID(id).Preload()
	ds.Meeting_FontBoldID(id).Preload()
	ds.Meeting_FontBoldItalicID(id).Preload()
	ds.Meeting_FontChyronSpeakerNameID(id).Preload()
	ds.Meeting_FontItalicID(id).Preload()
	ds.Meeting_FontMonospaceID(id).Preload()
	ds.Meeting_FontProjectorH1ID(id).Preload()
	ds.Meeting_FontProjectorH2ID(id).Preload()
	ds.Meeting_FontRegularID(id).Preload()
	ds.Meeting_ForwardedMotionIDs(id).Preload()
	ds.Meeting_GroupIDs(id).Preload()
	ds.Meeting_ID(id).Preload()
	ds.Meeting_ImportedAt(id).Preload()
	ds.Meeting_IsActiveInOrganizationID(id).Preload()
	ds.Meeting_IsArchivedInOrganizationID(id).Preload()
	ds.Meeting_JitsiDomain(id).Preload()
	ds.Meeting_JitsiRoomName(id).Preload()
	ds.Meeting_JitsiRoomPassword(id).Preload()
	ds.Meeting_Language(id).Preload()
	ds.Meeting_ListOfSpeakersAllowMultipleSpeakers(id).Preload()
	ds.Meeting_ListOfSpeakersAmountLastOnProjector(id).Preload()
	ds.Meeting_ListOfSpeakersAmountNextOnProjector(id).Preload()
	ds.Meeting_ListOfSpeakersCanCreatePointOfOrderForOthers(id).Preload()
	ds.Meeting_ListOfSpeakersCanSetContributionSelf(id).Preload()
	ds.Meeting_ListOfSpeakersClosingDisablesPointOfOrder(id).Preload()
	ds.Meeting_ListOfSpeakersCountdownID(id).Preload()
	ds.Meeting_ListOfSpeakersCoupleCountdown(id).Preload()
	ds.Meeting_ListOfSpeakersDefaultStructureLevelTime(id).Preload()
	ds.Meeting_ListOfSpeakersEnableInterposedQuestion(id).Preload()
	ds.Meeting_ListOfSpeakersEnablePointOfOrderCategories(id).Preload()
	ds.Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(id).Preload()
	ds.Meeting_ListOfSpeakersEnableProContraSpeech(id).Preload()
	ds.Meeting_ListOfSpeakersHideContributionCount(id).Preload()
	ds.Meeting_ListOfSpeakersIDs(id).Preload()
	ds.Meeting_ListOfSpeakersInitiallyClosed(id).Preload()
	ds.Meeting_ListOfSpeakersInterventionTime(id).Preload()
	ds.Meeting_ListOfSpeakersPresentUsersOnly(id).Preload()
	ds.Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(id).Preload()
	ds.Meeting_ListOfSpeakersShowFirstContribution(id).Preload()
	ds.Meeting_ListOfSpeakersSpeakerNoteForEveryone(id).Preload()
	ds.Meeting_Location(id).Preload()
	ds.Meeting_LockedFromInside(id).Preload()
	ds.Meeting_LogoPdfBallotPaperID(id).Preload()
	ds.Meeting_LogoPdfFooterLID(id).Preload()
	ds.Meeting_LogoPdfFooterRID(id).Preload()
	ds.Meeting_LogoPdfHeaderLID(id).Preload()
	ds.Meeting_LogoPdfHeaderRID(id).Preload()
	ds.Meeting_LogoProjectorHeaderID(id).Preload()
	ds.Meeting_LogoProjectorMainID(id).Preload()
	ds.Meeting_LogoWebHeaderID(id).Preload()
	ds.Meeting_MediafileIDs(id).Preload()
	ds.Meeting_MeetingMediafileIDs(id).Preload()
	ds.Meeting_MeetingUserIDs(id).Preload()
	ds.Meeting_MotionBlockIDs(id).Preload()
	ds.Meeting_MotionCategoryIDs(id).Preload()
	ds.Meeting_MotionChangeRecommendationIDs(id).Preload()
	ds.Meeting_MotionCommentIDs(id).Preload()
	ds.Meeting_MotionCommentSectionIDs(id).Preload()
	ds.Meeting_MotionEditorIDs(id).Preload()
	ds.Meeting_MotionIDs(id).Preload()
	ds.Meeting_MotionPollBallotPaperNumber(id).Preload()
	ds.Meeting_MotionPollBallotPaperSelection(id).Preload()
	ds.Meeting_MotionPollDefaultBackend(id).Preload()
	ds.Meeting_MotionPollDefaultGroupIDs(id).Preload()
	ds.Meeting_MotionPollDefaultMethod(id).Preload()
	ds.Meeting_MotionPollDefaultOnehundredPercentBase(id).Preload()
	ds.Meeting_MotionPollDefaultType(id).Preload()
	ds.Meeting_MotionStateIDs(id).Preload()
	ds.Meeting_MotionSubmitterIDs(id).Preload()
	ds.Meeting_MotionWorkflowIDs(id).Preload()
	ds.Meeting_MotionWorkingGroupSpeakerIDs(id).Preload()
	ds.Meeting_MotionsAmendmentsEnabled(id).Preload()
	ds.Meeting_MotionsAmendmentsInMainList(id).Preload()
	ds.Meeting_MotionsAmendmentsMultipleParagraphs(id).Preload()
	ds.Meeting_MotionsAmendmentsOfAmendments(id).Preload()
	ds.Meeting_MotionsAmendmentsPrefix(id).Preload()
	ds.Meeting_MotionsAmendmentsTextMode(id).Preload()
	ds.Meeting_MotionsBlockSlideColumns(id).Preload()
	ds.Meeting_MotionsCreateEnableAdditionalSubmitterText(id).Preload()
	ds.Meeting_MotionsDefaultAmendmentWorkflowID(id).Preload()
	ds.Meeting_MotionsDefaultLineNumbering(id).Preload()
	ds.Meeting_MotionsDefaultSorting(id).Preload()
	ds.Meeting_MotionsDefaultWorkflowID(id).Preload()
	ds.Meeting_MotionsEnableEditor(id).Preload()
	ds.Meeting_MotionsEnableReasonOnProjector(id).Preload()
	ds.Meeting_MotionsEnableRecommendationOnProjector(id).Preload()
	ds.Meeting_MotionsEnableSideboxOnProjector(id).Preload()
	ds.Meeting_MotionsEnableTextOnProjector(id).Preload()
	ds.Meeting_MotionsEnableWorkingGroupSpeaker(id).Preload()
	ds.Meeting_MotionsExportFollowRecommendation(id).Preload()
	ds.Meeting_MotionsExportPreamble(id).Preload()
	ds.Meeting_MotionsExportSubmitterRecommendation(id).Preload()
	ds.Meeting_MotionsExportTitle(id).Preload()
	ds.Meeting_MotionsHideMetadataBackground(id).Preload()
	ds.Meeting_MotionsLineLength(id).Preload()
	ds.Meeting_MotionsNumberMinDigits(id).Preload()
	ds.Meeting_MotionsNumberType(id).Preload()
	ds.Meeting_MotionsNumberWithBlank(id).Preload()
	ds.Meeting_MotionsPreamble(id).Preload()
	ds.Meeting_MotionsReasonRequired(id).Preload()
	ds.Meeting_MotionsRecommendationTextMode(id).Preload()
	ds.Meeting_MotionsRecommendationsBy(id).Preload()
	ds.Meeting_MotionsShowReferringMotions(id).Preload()
	ds.Meeting_MotionsShowSequentialNumber(id).Preload()
	ds.Meeting_MotionsSupportersMinAmount(id).Preload()
	ds.Meeting_Name(id).Preload()
	ds.Meeting_OptionIDs(id).Preload()
	ds.Meeting_OrganizationTagIDs(id).Preload()
	ds.Meeting_PersonalNoteIDs(id).Preload()
	ds.Meeting_PointOfOrderCategoryIDs(id).Preload()
	ds.Meeting_PollBallotPaperNumber(id).Preload()
	ds.Meeting_PollBallotPaperSelection(id).Preload()
	ds.Meeting_PollCandidateIDs(id).Preload()
	ds.Meeting_PollCandidateListIDs(id).Preload()
	ds.Meeting_PollCountdownID(id).Preload()
	ds.Meeting_PollCoupleCountdown(id).Preload()
	ds.Meeting_PollDefaultBackend(id).Preload()
	ds.Meeting_PollDefaultGroupIDs(id).Preload()
	ds.Meeting_PollDefaultMethod(id).Preload()
	ds.Meeting_PollDefaultOnehundredPercentBase(id).Preload()
	ds.Meeting_PollDefaultType(id).Preload()
	ds.Meeting_PollIDs(id).Preload()
	ds.Meeting_PollSortPollResultByVotes(id).Preload()
	ds.Meeting_PresentUserIDs(id).Preload()
	ds.Meeting_ProjectionIDs(id).Preload()
	ds.Meeting_ProjectorCountdownDefaultTime(id).Preload()
	ds.Meeting_ProjectorCountdownIDs(id).Preload()
	ds.Meeting_ProjectorCountdownWarningTime(id).Preload()
	ds.Meeting_ProjectorIDs(id).Preload()
	ds.Meeting_ProjectorMessageIDs(id).Preload()
	ds.Meeting_ReferenceProjectorID(id).Preload()
	ds.Meeting_SpeakerIDs(id).Preload()
	ds.Meeting_StartTime(id).Preload()
	ds.Meeting_StructureLevelIDs(id).Preload()
	ds.Meeting_StructureLevelListOfSpeakersIDs(id).Preload()
	ds.Meeting_TagIDs(id).Preload()
	ds.Meeting_TemplateForOrganizationID(id).Preload()
	ds.Meeting_TopicIDs(id).Preload()
	ds.Meeting_TopicPollDefaultGroupIDs(id).Preload()
	ds.Meeting_UserIDs(id).Preload()
	ds.Meeting_UsersAllowSelfSetPresent(id).Preload()
	ds.Meeting_UsersEmailBody(id).Preload()
	ds.Meeting_UsersEmailReplyto(id).Preload()
	ds.Meeting_UsersEmailSender(id).Preload()
	ds.Meeting_UsersEmailSubject(id).Preload()
	ds.Meeting_UsersEnablePresenceView(id).Preload()
	ds.Meeting_UsersEnableVoteDelegations(id).Preload()
	ds.Meeting_UsersEnableVoteWeight(id).Preload()
	ds.Meeting_UsersForbidDelegatorAsSubmitter(id).Preload()
	ds.Meeting_UsersForbidDelegatorAsSupporter(id).Preload()
	ds.Meeting_UsersForbidDelegatorInListOfSpeakers(id).Preload()
	ds.Meeting_UsersForbidDelegatorToVote(id).Preload()
	ds.Meeting_UsersPdfWelcometext(id).Preload()
	ds.Meeting_UsersPdfWelcometitle(id).Preload()
	ds.Meeting_UsersPdfWlanEncryption(id).Preload()
	ds.Meeting_UsersPdfWlanPassword(id).Preload()
	ds.Meeting_UsersPdfWlanSsid(id).Preload()
	ds.Meeting_VoteIDs(id).Preload()
	ds.Meeting_WelcomeText(id).Preload()
	ds.Meeting_WelcomeTitle(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Meeting(id int) *ValueCollection[Meeting, *Meeting] {
	return &ValueCollection[Meeting, *Meeting]{
		id:    id,
		fetch: r,
	}
}

// MeetingMediafile has all fields from meeting_mediafile.
type MeetingMediafile struct {
	AccessGroupIDs                         []int
	AttachmentIDs                          []string
	ID                                     int
	InheritedAccessGroupIDs                []int
	IsPublic                               bool
	ListOfSpeakersID                       Maybe[int]
	MediafileID                            int
	MeetingID                              int
	ProjectionIDs                          []int
	UsedAsFontBoldInMeetingID              Maybe[int]
	UsedAsFontBoldItalicInMeetingID        Maybe[int]
	UsedAsFontChyronSpeakerNameInMeetingID Maybe[int]
	UsedAsFontItalicInMeetingID            Maybe[int]
	UsedAsFontMonospaceInMeetingID         Maybe[int]
	UsedAsFontProjectorH1InMeetingID       Maybe[int]
	UsedAsFontProjectorH2InMeetingID       Maybe[int]
	UsedAsFontRegularInMeetingID           Maybe[int]
	UsedAsLogoPdfBallotPaperInMeetingID    Maybe[int]
	UsedAsLogoPdfFooterLInMeetingID        Maybe[int]
	UsedAsLogoPdfFooterRInMeetingID        Maybe[int]
	UsedAsLogoPdfHeaderLInMeetingID        Maybe[int]
	UsedAsLogoPdfHeaderRInMeetingID        Maybe[int]
	UsedAsLogoProjectorHeaderInMeetingID   Maybe[int]
	UsedAsLogoProjectorMainInMeetingID     Maybe[int]
	UsedAsLogoWebHeaderInMeetingID         Maybe[int]
	fetch                                  *Fetch
}

func (c *MeetingMediafile) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MeetingMediafile_AccessGroupIDs(id).Lazy(&c.AccessGroupIDs)
	ds.MeetingMediafile_AttachmentIDs(id).Lazy(&c.AttachmentIDs)
	ds.MeetingMediafile_ID(id).Lazy(&c.ID)
	ds.MeetingMediafile_InheritedAccessGroupIDs(id).Lazy(&c.InheritedAccessGroupIDs)
	ds.MeetingMediafile_IsPublic(id).Lazy(&c.IsPublic)
	ds.MeetingMediafile_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.MeetingMediafile_MediafileID(id).Lazy(&c.MediafileID)
	ds.MeetingMediafile_MeetingID(id).Lazy(&c.MeetingID)
	ds.MeetingMediafile_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.MeetingMediafile_UsedAsFontBoldInMeetingID(id).Lazy(&c.UsedAsFontBoldInMeetingID)
	ds.MeetingMediafile_UsedAsFontBoldItalicInMeetingID(id).Lazy(&c.UsedAsFontBoldItalicInMeetingID)
	ds.MeetingMediafile_UsedAsFontChyronSpeakerNameInMeetingID(id).Lazy(&c.UsedAsFontChyronSpeakerNameInMeetingID)
	ds.MeetingMediafile_UsedAsFontItalicInMeetingID(id).Lazy(&c.UsedAsFontItalicInMeetingID)
	ds.MeetingMediafile_UsedAsFontMonospaceInMeetingID(id).Lazy(&c.UsedAsFontMonospaceInMeetingID)
	ds.MeetingMediafile_UsedAsFontProjectorH1InMeetingID(id).Lazy(&c.UsedAsFontProjectorH1InMeetingID)
	ds.MeetingMediafile_UsedAsFontProjectorH2InMeetingID(id).Lazy(&c.UsedAsFontProjectorH2InMeetingID)
	ds.MeetingMediafile_UsedAsFontRegularInMeetingID(id).Lazy(&c.UsedAsFontRegularInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfBallotPaperInMeetingID(id).Lazy(&c.UsedAsLogoPdfBallotPaperInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfFooterLInMeetingID(id).Lazy(&c.UsedAsLogoPdfFooterLInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfFooterRInMeetingID(id).Lazy(&c.UsedAsLogoPdfFooterRInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfHeaderLInMeetingID(id).Lazy(&c.UsedAsLogoPdfHeaderLInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfHeaderRInMeetingID(id).Lazy(&c.UsedAsLogoPdfHeaderRInMeetingID)
	ds.MeetingMediafile_UsedAsLogoProjectorHeaderInMeetingID(id).Lazy(&c.UsedAsLogoProjectorHeaderInMeetingID)
	ds.MeetingMediafile_UsedAsLogoProjectorMainInMeetingID(id).Lazy(&c.UsedAsLogoProjectorMainInMeetingID)
	ds.MeetingMediafile_UsedAsLogoWebHeaderInMeetingID(id).Lazy(&c.UsedAsLogoWebHeaderInMeetingID)
}

func (c *MeetingMediafile) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MeetingMediafile_AccessGroupIDs(id).Preload()
	ds.MeetingMediafile_AttachmentIDs(id).Preload()
	ds.MeetingMediafile_ID(id).Preload()
	ds.MeetingMediafile_InheritedAccessGroupIDs(id).Preload()
	ds.MeetingMediafile_IsPublic(id).Preload()
	ds.MeetingMediafile_ListOfSpeakersID(id).Preload()
	ds.MeetingMediafile_MediafileID(id).Preload()
	ds.MeetingMediafile_MeetingID(id).Preload()
	ds.MeetingMediafile_ProjectionIDs(id).Preload()
	ds.MeetingMediafile_UsedAsFontBoldInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontBoldItalicInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontChyronSpeakerNameInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontItalicInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontMonospaceInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontProjectorH1InMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontProjectorH2InMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontRegularInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfBallotPaperInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfFooterLInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfFooterRInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfHeaderLInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfHeaderRInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoProjectorHeaderInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoProjectorMainInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoWebHeaderInMeetingID(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MeetingMediafile(id int) *ValueCollection[MeetingMediafile, *MeetingMediafile] {
	return &ValueCollection[MeetingMediafile, *MeetingMediafile]{
		id:    id,
		fetch: r,
	}
}

// MeetingUser has all fields from meeting_user.
type MeetingUser struct {
	AboutMe                      string
	AssignmentCandidateIDs       []int
	ChatMessageIDs               []int
	Comment                      string
	GroupIDs                     []int
	ID                           int
	LockedOut                    bool
	MeetingID                    int
	MotionEditorIDs              []int
	MotionSubmitterIDs           []int
	MotionWorkingGroupSpeakerIDs []int
	Number                       string
	PersonalNoteIDs              []int
	SpeakerIDs                   []int
	StructureLevelIDs            []int
	SupportedMotionIDs           []int
	UserID                       int
	VoteDelegatedToID            Maybe[int]
	VoteDelegationsFromIDs       []int
	VoteWeight                   string
	fetch                        *Fetch
}

func (c *MeetingUser) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MeetingUser_AboutMe(id).Lazy(&c.AboutMe)
	ds.MeetingUser_AssignmentCandidateIDs(id).Lazy(&c.AssignmentCandidateIDs)
	ds.MeetingUser_ChatMessageIDs(id).Lazy(&c.ChatMessageIDs)
	ds.MeetingUser_Comment(id).Lazy(&c.Comment)
	ds.MeetingUser_GroupIDs(id).Lazy(&c.GroupIDs)
	ds.MeetingUser_ID(id).Lazy(&c.ID)
	ds.MeetingUser_LockedOut(id).Lazy(&c.LockedOut)
	ds.MeetingUser_MeetingID(id).Lazy(&c.MeetingID)
	ds.MeetingUser_MotionEditorIDs(id).Lazy(&c.MotionEditorIDs)
	ds.MeetingUser_MotionSubmitterIDs(id).Lazy(&c.MotionSubmitterIDs)
	ds.MeetingUser_MotionWorkingGroupSpeakerIDs(id).Lazy(&c.MotionWorkingGroupSpeakerIDs)
	ds.MeetingUser_Number(id).Lazy(&c.Number)
	ds.MeetingUser_PersonalNoteIDs(id).Lazy(&c.PersonalNoteIDs)
	ds.MeetingUser_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.MeetingUser_StructureLevelIDs(id).Lazy(&c.StructureLevelIDs)
	ds.MeetingUser_SupportedMotionIDs(id).Lazy(&c.SupportedMotionIDs)
	ds.MeetingUser_UserID(id).Lazy(&c.UserID)
	ds.MeetingUser_VoteDelegatedToID(id).Lazy(&c.VoteDelegatedToID)
	ds.MeetingUser_VoteDelegationsFromIDs(id).Lazy(&c.VoteDelegationsFromIDs)
	ds.MeetingUser_VoteWeight(id).Lazy(&c.VoteWeight)
}

func (c *MeetingUser) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MeetingUser_AboutMe(id).Preload()
	ds.MeetingUser_AssignmentCandidateIDs(id).Preload()
	ds.MeetingUser_ChatMessageIDs(id).Preload()
	ds.MeetingUser_Comment(id).Preload()
	ds.MeetingUser_GroupIDs(id).Preload()
	ds.MeetingUser_ID(id).Preload()
	ds.MeetingUser_LockedOut(id).Preload()
	ds.MeetingUser_MeetingID(id).Preload()
	ds.MeetingUser_MotionEditorIDs(id).Preload()
	ds.MeetingUser_MotionSubmitterIDs(id).Preload()
	ds.MeetingUser_MotionWorkingGroupSpeakerIDs(id).Preload()
	ds.MeetingUser_Number(id).Preload()
	ds.MeetingUser_PersonalNoteIDs(id).Preload()
	ds.MeetingUser_SpeakerIDs(id).Preload()
	ds.MeetingUser_StructureLevelIDs(id).Preload()
	ds.MeetingUser_SupportedMotionIDs(id).Preload()
	ds.MeetingUser_UserID(id).Preload()
	ds.MeetingUser_VoteDelegatedToID(id).Preload()
	ds.MeetingUser_VoteDelegationsFromIDs(id).Preload()
	ds.MeetingUser_VoteWeight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MeetingUser(id int) *ValueCollection[MeetingUser, *MeetingUser] {
	return &ValueCollection[MeetingUser, *MeetingUser]{
		id:    id,
		fetch: r,
	}
}

// Motion has all fields from motion.
type Motion struct {
	AdditionalSubmitter                          string
	AgendaItemID                                 Maybe[int]
	AllDerivedMotionIDs                          []int
	AllOriginIDs                                 []int
	AmendmentIDs                                 []int
	AmendmentParagraphs                          json.RawMessage
	AttachmentMeetingMediafileIDs                []int
	BlockID                                      Maybe[int]
	CategoryID                                   Maybe[int]
	CategoryWeight                               int
	ChangeRecommendationIDs                      []int
	CommentIDs                                   []int
	Created                                      int
	DerivedMotionIDs                             []int
	EditorIDs                                    []int
	Forwarded                                    int
	ID                                           int
	IDenticalMotionIDs                           []int
	LastModified                                 int
	LeadMotionID                                 Maybe[int]
	ListOfSpeakersID                             int
	MeetingID                                    int
	ModifiedFinalVersion                         string
	Number                                       string
	NumberValue                                  int
	OptionIDs                                    []int
	OriginID                                     Maybe[int]
	OriginMeetingID                              Maybe[int]
	PersonalNoteIDs                              []int
	PollIDs                                      []int
	ProjectionIDs                                []int
	Reason                                       string
	RecommendationExtension                      string
	RecommendationExtensionReferenceIDs          []string
	RecommendationID                             Maybe[int]
	ReferencedInMotionRecommendationExtensionIDs []int
	ReferencedInMotionStateExtensionIDs          []int
	SequentialNumber                             int
	SortChildIDs                                 []int
	SortParentID                                 Maybe[int]
	SortWeight                                   int
	StartLineNumber                              int
	StateExtension                               string
	StateExtensionReferenceIDs                   []string
	StateID                                      int
	SubmitterIDs                                 []int
	SupporterMeetingUserIDs                      []int
	TagIDs                                       []int
	Text                                         string
	TextHash                                     string
	Title                                        string
	WorkflowTimestamp                            int
	WorkingGroupSpeakerIDs                       []int
	fetch                                        *Fetch
}

func (c *Motion) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Motion_AdditionalSubmitter(id).Lazy(&c.AdditionalSubmitter)
	ds.Motion_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.Motion_AllDerivedMotionIDs(id).Lazy(&c.AllDerivedMotionIDs)
	ds.Motion_AllOriginIDs(id).Lazy(&c.AllOriginIDs)
	ds.Motion_AmendmentIDs(id).Lazy(&c.AmendmentIDs)
	ds.Motion_AmendmentParagraphs(id).Lazy(&c.AmendmentParagraphs)
	ds.Motion_AttachmentMeetingMediafileIDs(id).Lazy(&c.AttachmentMeetingMediafileIDs)
	ds.Motion_BlockID(id).Lazy(&c.BlockID)
	ds.Motion_CategoryID(id).Lazy(&c.CategoryID)
	ds.Motion_CategoryWeight(id).Lazy(&c.CategoryWeight)
	ds.Motion_ChangeRecommendationIDs(id).Lazy(&c.ChangeRecommendationIDs)
	ds.Motion_CommentIDs(id).Lazy(&c.CommentIDs)
	ds.Motion_Created(id).Lazy(&c.Created)
	ds.Motion_DerivedMotionIDs(id).Lazy(&c.DerivedMotionIDs)
	ds.Motion_EditorIDs(id).Lazy(&c.EditorIDs)
	ds.Motion_Forwarded(id).Lazy(&c.Forwarded)
	ds.Motion_ID(id).Lazy(&c.ID)
	ds.Motion_IDenticalMotionIDs(id).Lazy(&c.IDenticalMotionIDs)
	ds.Motion_LastModified(id).Lazy(&c.LastModified)
	ds.Motion_LeadMotionID(id).Lazy(&c.LeadMotionID)
	ds.Motion_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.Motion_MeetingID(id).Lazy(&c.MeetingID)
	ds.Motion_ModifiedFinalVersion(id).Lazy(&c.ModifiedFinalVersion)
	ds.Motion_Number(id).Lazy(&c.Number)
	ds.Motion_NumberValue(id).Lazy(&c.NumberValue)
	ds.Motion_OptionIDs(id).Lazy(&c.OptionIDs)
	ds.Motion_OriginID(id).Lazy(&c.OriginID)
	ds.Motion_OriginMeetingID(id).Lazy(&c.OriginMeetingID)
	ds.Motion_PersonalNoteIDs(id).Lazy(&c.PersonalNoteIDs)
	ds.Motion_PollIDs(id).Lazy(&c.PollIDs)
	ds.Motion_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Motion_Reason(id).Lazy(&c.Reason)
	ds.Motion_RecommendationExtension(id).Lazy(&c.RecommendationExtension)
	ds.Motion_RecommendationExtensionReferenceIDs(id).Lazy(&c.RecommendationExtensionReferenceIDs)
	ds.Motion_RecommendationID(id).Lazy(&c.RecommendationID)
	ds.Motion_ReferencedInMotionRecommendationExtensionIDs(id).Lazy(&c.ReferencedInMotionRecommendationExtensionIDs)
	ds.Motion_ReferencedInMotionStateExtensionIDs(id).Lazy(&c.ReferencedInMotionStateExtensionIDs)
	ds.Motion_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Motion_SortChildIDs(id).Lazy(&c.SortChildIDs)
	ds.Motion_SortParentID(id).Lazy(&c.SortParentID)
	ds.Motion_SortWeight(id).Lazy(&c.SortWeight)
	ds.Motion_StartLineNumber(id).Lazy(&c.StartLineNumber)
	ds.Motion_StateExtension(id).Lazy(&c.StateExtension)
	ds.Motion_StateExtensionReferenceIDs(id).Lazy(&c.StateExtensionReferenceIDs)
	ds.Motion_StateID(id).Lazy(&c.StateID)
	ds.Motion_SubmitterIDs(id).Lazy(&c.SubmitterIDs)
	ds.Motion_SupporterMeetingUserIDs(id).Lazy(&c.SupporterMeetingUserIDs)
	ds.Motion_TagIDs(id).Lazy(&c.TagIDs)
	ds.Motion_Text(id).Lazy(&c.Text)
	ds.Motion_TextHash(id).Lazy(&c.TextHash)
	ds.Motion_Title(id).Lazy(&c.Title)
	ds.Motion_WorkflowTimestamp(id).Lazy(&c.WorkflowTimestamp)
	ds.Motion_WorkingGroupSpeakerIDs(id).Lazy(&c.WorkingGroupSpeakerIDs)
}

func (c *Motion) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Motion_AdditionalSubmitter(id).Preload()
	ds.Motion_AgendaItemID(id).Preload()
	ds.Motion_AllDerivedMotionIDs(id).Preload()
	ds.Motion_AllOriginIDs(id).Preload()
	ds.Motion_AmendmentIDs(id).Preload()
	ds.Motion_AmendmentParagraphs(id).Preload()
	ds.Motion_AttachmentMeetingMediafileIDs(id).Preload()
	ds.Motion_BlockID(id).Preload()
	ds.Motion_CategoryID(id).Preload()
	ds.Motion_CategoryWeight(id).Preload()
	ds.Motion_ChangeRecommendationIDs(id).Preload()
	ds.Motion_CommentIDs(id).Preload()
	ds.Motion_Created(id).Preload()
	ds.Motion_DerivedMotionIDs(id).Preload()
	ds.Motion_EditorIDs(id).Preload()
	ds.Motion_Forwarded(id).Preload()
	ds.Motion_ID(id).Preload()
	ds.Motion_IDenticalMotionIDs(id).Preload()
	ds.Motion_LastModified(id).Preload()
	ds.Motion_LeadMotionID(id).Preload()
	ds.Motion_ListOfSpeakersID(id).Preload()
	ds.Motion_MeetingID(id).Preload()
	ds.Motion_ModifiedFinalVersion(id).Preload()
	ds.Motion_Number(id).Preload()
	ds.Motion_NumberValue(id).Preload()
	ds.Motion_OptionIDs(id).Preload()
	ds.Motion_OriginID(id).Preload()
	ds.Motion_OriginMeetingID(id).Preload()
	ds.Motion_PersonalNoteIDs(id).Preload()
	ds.Motion_PollIDs(id).Preload()
	ds.Motion_ProjectionIDs(id).Preload()
	ds.Motion_Reason(id).Preload()
	ds.Motion_RecommendationExtension(id).Preload()
	ds.Motion_RecommendationExtensionReferenceIDs(id).Preload()
	ds.Motion_RecommendationID(id).Preload()
	ds.Motion_ReferencedInMotionRecommendationExtensionIDs(id).Preload()
	ds.Motion_ReferencedInMotionStateExtensionIDs(id).Preload()
	ds.Motion_SequentialNumber(id).Preload()
	ds.Motion_SortChildIDs(id).Preload()
	ds.Motion_SortParentID(id).Preload()
	ds.Motion_SortWeight(id).Preload()
	ds.Motion_StartLineNumber(id).Preload()
	ds.Motion_StateExtension(id).Preload()
	ds.Motion_StateExtensionReferenceIDs(id).Preload()
	ds.Motion_StateID(id).Preload()
	ds.Motion_SubmitterIDs(id).Preload()
	ds.Motion_SupporterMeetingUserIDs(id).Preload()
	ds.Motion_TagIDs(id).Preload()
	ds.Motion_Text(id).Preload()
	ds.Motion_TextHash(id).Preload()
	ds.Motion_Title(id).Preload()
	ds.Motion_WorkflowTimestamp(id).Preload()
	ds.Motion_WorkingGroupSpeakerIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Motion(id int) *ValueCollection[Motion, *Motion] {
	return &ValueCollection[Motion, *Motion]{
		id:    id,
		fetch: r,
	}
}

// MotionBlock has all fields from motion_block.
type MotionBlock struct {
	AgendaItemID     Maybe[int]
	ID               int
	Internal         bool
	ListOfSpeakersID int
	MeetingID        int
	MotionIDs        []int
	ProjectionIDs    []int
	SequentialNumber int
	Title            string
	fetch            *Fetch
}

func (c *MotionBlock) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionBlock_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.MotionBlock_ID(id).Lazy(&c.ID)
	ds.MotionBlock_Internal(id).Lazy(&c.Internal)
	ds.MotionBlock_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.MotionBlock_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionBlock_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.MotionBlock_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.MotionBlock_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionBlock_Title(id).Lazy(&c.Title)
}

func (c *MotionBlock) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionBlock_AgendaItemID(id).Preload()
	ds.MotionBlock_ID(id).Preload()
	ds.MotionBlock_Internal(id).Preload()
	ds.MotionBlock_ListOfSpeakersID(id).Preload()
	ds.MotionBlock_MeetingID(id).Preload()
	ds.MotionBlock_MotionIDs(id).Preload()
	ds.MotionBlock_ProjectionIDs(id).Preload()
	ds.MotionBlock_SequentialNumber(id).Preload()
	ds.MotionBlock_Title(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionBlock(id int) *ValueCollection[MotionBlock, *MotionBlock] {
	return &ValueCollection[MotionBlock, *MotionBlock]{
		id:    id,
		fetch: r,
	}
}

// MotionCategory has all fields from motion_category.
type MotionCategory struct {
	ChildIDs         []int
	ID               int
	Level            int
	MeetingID        int
	MotionIDs        []int
	Name             string
	ParentID         Maybe[int]
	Prefix           string
	SequentialNumber int
	Weight           int
	fetch            *Fetch
}

func (c *MotionCategory) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionCategory_ChildIDs(id).Lazy(&c.ChildIDs)
	ds.MotionCategory_ID(id).Lazy(&c.ID)
	ds.MotionCategory_Level(id).Lazy(&c.Level)
	ds.MotionCategory_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionCategory_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.MotionCategory_Name(id).Lazy(&c.Name)
	ds.MotionCategory_ParentID(id).Lazy(&c.ParentID)
	ds.MotionCategory_Prefix(id).Lazy(&c.Prefix)
	ds.MotionCategory_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionCategory_Weight(id).Lazy(&c.Weight)
}

func (c *MotionCategory) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionCategory_ChildIDs(id).Preload()
	ds.MotionCategory_ID(id).Preload()
	ds.MotionCategory_Level(id).Preload()
	ds.MotionCategory_MeetingID(id).Preload()
	ds.MotionCategory_MotionIDs(id).Preload()
	ds.MotionCategory_Name(id).Preload()
	ds.MotionCategory_ParentID(id).Preload()
	ds.MotionCategory_Prefix(id).Preload()
	ds.MotionCategory_SequentialNumber(id).Preload()
	ds.MotionCategory_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionCategory(id int) *ValueCollection[MotionCategory, *MotionCategory] {
	return &ValueCollection[MotionCategory, *MotionCategory]{
		id:    id,
		fetch: r,
	}
}

// MotionChangeRecommendation has all fields from motion_change_recommendation.
type MotionChangeRecommendation struct {
	CreationTime     int
	ID               int
	Internal         bool
	LineFrom         int
	LineTo           int
	MeetingID        int
	MotionID         int
	OtherDescription string
	Rejected         bool
	Text             string
	Type             string
	fetch            *Fetch
}

func (c *MotionChangeRecommendation) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionChangeRecommendation_CreationTime(id).Lazy(&c.CreationTime)
	ds.MotionChangeRecommendation_ID(id).Lazy(&c.ID)
	ds.MotionChangeRecommendation_Internal(id).Lazy(&c.Internal)
	ds.MotionChangeRecommendation_LineFrom(id).Lazy(&c.LineFrom)
	ds.MotionChangeRecommendation_LineTo(id).Lazy(&c.LineTo)
	ds.MotionChangeRecommendation_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionChangeRecommendation_MotionID(id).Lazy(&c.MotionID)
	ds.MotionChangeRecommendation_OtherDescription(id).Lazy(&c.OtherDescription)
	ds.MotionChangeRecommendation_Rejected(id).Lazy(&c.Rejected)
	ds.MotionChangeRecommendation_Text(id).Lazy(&c.Text)
	ds.MotionChangeRecommendation_Type(id).Lazy(&c.Type)
}

func (c *MotionChangeRecommendation) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionChangeRecommendation_CreationTime(id).Preload()
	ds.MotionChangeRecommendation_ID(id).Preload()
	ds.MotionChangeRecommendation_Internal(id).Preload()
	ds.MotionChangeRecommendation_LineFrom(id).Preload()
	ds.MotionChangeRecommendation_LineTo(id).Preload()
	ds.MotionChangeRecommendation_MeetingID(id).Preload()
	ds.MotionChangeRecommendation_MotionID(id).Preload()
	ds.MotionChangeRecommendation_OtherDescription(id).Preload()
	ds.MotionChangeRecommendation_Rejected(id).Preload()
	ds.MotionChangeRecommendation_Text(id).Preload()
	ds.MotionChangeRecommendation_Type(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionChangeRecommendation(id int) *ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation] {
	return &ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation]{
		id:    id,
		fetch: r,
	}
}

// MotionComment has all fields from motion_comment.
type MotionComment struct {
	Comment   string
	ID        int
	MeetingID int
	MotionID  int
	SectionID int
	fetch     *Fetch
}

func (c *MotionComment) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionComment_Comment(id).Lazy(&c.Comment)
	ds.MotionComment_ID(id).Lazy(&c.ID)
	ds.MotionComment_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionComment_MotionID(id).Lazy(&c.MotionID)
	ds.MotionComment_SectionID(id).Lazy(&c.SectionID)
}

func (c *MotionComment) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionComment_Comment(id).Preload()
	ds.MotionComment_ID(id).Preload()
	ds.MotionComment_MeetingID(id).Preload()
	ds.MotionComment_MotionID(id).Preload()
	ds.MotionComment_SectionID(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionComment(id int) *ValueCollection[MotionComment, *MotionComment] {
	return &ValueCollection[MotionComment, *MotionComment]{
		id:    id,
		fetch: r,
	}
}

// MotionCommentSection has all fields from motion_comment_section.
type MotionCommentSection struct {
	CommentIDs        []int
	ID                int
	MeetingID         int
	Name              string
	ReadGroupIDs      []int
	SequentialNumber  int
	SubmitterCanWrite bool
	Weight            int
	WriteGroupIDs     []int
	fetch             *Fetch
}

func (c *MotionCommentSection) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionCommentSection_CommentIDs(id).Lazy(&c.CommentIDs)
	ds.MotionCommentSection_ID(id).Lazy(&c.ID)
	ds.MotionCommentSection_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionCommentSection_Name(id).Lazy(&c.Name)
	ds.MotionCommentSection_ReadGroupIDs(id).Lazy(&c.ReadGroupIDs)
	ds.MotionCommentSection_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionCommentSection_SubmitterCanWrite(id).Lazy(&c.SubmitterCanWrite)
	ds.MotionCommentSection_Weight(id).Lazy(&c.Weight)
	ds.MotionCommentSection_WriteGroupIDs(id).Lazy(&c.WriteGroupIDs)
}

func (c *MotionCommentSection) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionCommentSection_CommentIDs(id).Preload()
	ds.MotionCommentSection_ID(id).Preload()
	ds.MotionCommentSection_MeetingID(id).Preload()
	ds.MotionCommentSection_Name(id).Preload()
	ds.MotionCommentSection_ReadGroupIDs(id).Preload()
	ds.MotionCommentSection_SequentialNumber(id).Preload()
	ds.MotionCommentSection_SubmitterCanWrite(id).Preload()
	ds.MotionCommentSection_Weight(id).Preload()
	ds.MotionCommentSection_WriteGroupIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionCommentSection(id int) *ValueCollection[MotionCommentSection, *MotionCommentSection] {
	return &ValueCollection[MotionCommentSection, *MotionCommentSection]{
		id:    id,
		fetch: r,
	}
}

// MotionEditor has all fields from motion_editor.
type MotionEditor struct {
	ID            int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        int
	fetch         *Fetch
}

func (c *MotionEditor) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionEditor_ID(id).Lazy(&c.ID)
	ds.MotionEditor_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionEditor_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionEditor_MotionID(id).Lazy(&c.MotionID)
	ds.MotionEditor_Weight(id).Lazy(&c.Weight)
}

func (c *MotionEditor) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionEditor_ID(id).Preload()
	ds.MotionEditor_MeetingID(id).Preload()
	ds.MotionEditor_MeetingUserID(id).Preload()
	ds.MotionEditor_MotionID(id).Preload()
	ds.MotionEditor_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionEditor(id int) *ValueCollection[MotionEditor, *MotionEditor] {
	return &ValueCollection[MotionEditor, *MotionEditor]{
		id:    id,
		fetch: r,
	}
}

// MotionState has all fields from motion_state.
type MotionState struct {
	AllowCreatePoll                  bool
	AllowMotionForwarding            bool
	AllowSubmitterEdit               bool
	AllowSupport                     bool
	CssClass                         string
	FirstStateOfWorkflowID           Maybe[int]
	ID                               int
	IsInternal                       bool
	MeetingID                        int
	MergeAmendmentIntoFinal          string
	MotionIDs                        []int
	MotionRecommendationIDs          []int
	Name                             string
	NextStateIDs                     []int
	PreviousStateIDs                 []int
	RecommendationLabel              string
	Restrictions                     []string
	SetNumber                        bool
	SetWorkflowTimestamp             bool
	ShowRecommendationExtensionField bool
	ShowStateExtensionField          bool
	SubmitterWithdrawBackIDs         []int
	SubmitterWithdrawStateID         Maybe[int]
	Weight                           int
	WorkflowID                       int
	fetch                            *Fetch
}

func (c *MotionState) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionState_AllowCreatePoll(id).Lazy(&c.AllowCreatePoll)
	ds.MotionState_AllowMotionForwarding(id).Lazy(&c.AllowMotionForwarding)
	ds.MotionState_AllowSubmitterEdit(id).Lazy(&c.AllowSubmitterEdit)
	ds.MotionState_AllowSupport(id).Lazy(&c.AllowSupport)
	ds.MotionState_CssClass(id).Lazy(&c.CssClass)
	ds.MotionState_FirstStateOfWorkflowID(id).Lazy(&c.FirstStateOfWorkflowID)
	ds.MotionState_ID(id).Lazy(&c.ID)
	ds.MotionState_IsInternal(id).Lazy(&c.IsInternal)
	ds.MotionState_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionState_MergeAmendmentIntoFinal(id).Lazy(&c.MergeAmendmentIntoFinal)
	ds.MotionState_MotionIDs(id).Lazy(&c.MotionIDs)
	ds.MotionState_MotionRecommendationIDs(id).Lazy(&c.MotionRecommendationIDs)
	ds.MotionState_Name(id).Lazy(&c.Name)
	ds.MotionState_NextStateIDs(id).Lazy(&c.NextStateIDs)
	ds.MotionState_PreviousStateIDs(id).Lazy(&c.PreviousStateIDs)
	ds.MotionState_RecommendationLabel(id).Lazy(&c.RecommendationLabel)
	ds.MotionState_Restrictions(id).Lazy(&c.Restrictions)
	ds.MotionState_SetNumber(id).Lazy(&c.SetNumber)
	ds.MotionState_SetWorkflowTimestamp(id).Lazy(&c.SetWorkflowTimestamp)
	ds.MotionState_ShowRecommendationExtensionField(id).Lazy(&c.ShowRecommendationExtensionField)
	ds.MotionState_ShowStateExtensionField(id).Lazy(&c.ShowStateExtensionField)
	ds.MotionState_SubmitterWithdrawBackIDs(id).Lazy(&c.SubmitterWithdrawBackIDs)
	ds.MotionState_SubmitterWithdrawStateID(id).Lazy(&c.SubmitterWithdrawStateID)
	ds.MotionState_Weight(id).Lazy(&c.Weight)
	ds.MotionState_WorkflowID(id).Lazy(&c.WorkflowID)
}

func (c *MotionState) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionState_AllowCreatePoll(id).Preload()
	ds.MotionState_AllowMotionForwarding(id).Preload()
	ds.MotionState_AllowSubmitterEdit(id).Preload()
	ds.MotionState_AllowSupport(id).Preload()
	ds.MotionState_CssClass(id).Preload()
	ds.MotionState_FirstStateOfWorkflowID(id).Preload()
	ds.MotionState_ID(id).Preload()
	ds.MotionState_IsInternal(id).Preload()
	ds.MotionState_MeetingID(id).Preload()
	ds.MotionState_MergeAmendmentIntoFinal(id).Preload()
	ds.MotionState_MotionIDs(id).Preload()
	ds.MotionState_MotionRecommendationIDs(id).Preload()
	ds.MotionState_Name(id).Preload()
	ds.MotionState_NextStateIDs(id).Preload()
	ds.MotionState_PreviousStateIDs(id).Preload()
	ds.MotionState_RecommendationLabel(id).Preload()
	ds.MotionState_Restrictions(id).Preload()
	ds.MotionState_SetNumber(id).Preload()
	ds.MotionState_SetWorkflowTimestamp(id).Preload()
	ds.MotionState_ShowRecommendationExtensionField(id).Preload()
	ds.MotionState_ShowStateExtensionField(id).Preload()
	ds.MotionState_SubmitterWithdrawBackIDs(id).Preload()
	ds.MotionState_SubmitterWithdrawStateID(id).Preload()
	ds.MotionState_Weight(id).Preload()
	ds.MotionState_WorkflowID(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionState(id int) *ValueCollection[MotionState, *MotionState] {
	return &ValueCollection[MotionState, *MotionState]{
		id:    id,
		fetch: r,
	}
}

// MotionSubmitter has all fields from motion_submitter.
type MotionSubmitter struct {
	ID            int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        int
	fetch         *Fetch
}

func (c *MotionSubmitter) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionSubmitter_ID(id).Lazy(&c.ID)
	ds.MotionSubmitter_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionSubmitter_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionSubmitter_MotionID(id).Lazy(&c.MotionID)
	ds.MotionSubmitter_Weight(id).Lazy(&c.Weight)
}

func (c *MotionSubmitter) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionSubmitter_ID(id).Preload()
	ds.MotionSubmitter_MeetingID(id).Preload()
	ds.MotionSubmitter_MeetingUserID(id).Preload()
	ds.MotionSubmitter_MotionID(id).Preload()
	ds.MotionSubmitter_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionSubmitter(id int) *ValueCollection[MotionSubmitter, *MotionSubmitter] {
	return &ValueCollection[MotionSubmitter, *MotionSubmitter]{
		id:    id,
		fetch: r,
	}
}

// MotionWorkflow has all fields from motion_workflow.
type MotionWorkflow struct {
	DefaultAmendmentWorkflowMeetingID Maybe[int]
	DefaultWorkflowMeetingID          Maybe[int]
	FirstStateID                      int
	ID                                int
	MeetingID                         int
	Name                              string
	SequentialNumber                  int
	StateIDs                          []int
	fetch                             *Fetch
}

func (c *MotionWorkflow) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionWorkflow_DefaultAmendmentWorkflowMeetingID(id).Lazy(&c.DefaultAmendmentWorkflowMeetingID)
	ds.MotionWorkflow_DefaultWorkflowMeetingID(id).Lazy(&c.DefaultWorkflowMeetingID)
	ds.MotionWorkflow_FirstStateID(id).Lazy(&c.FirstStateID)
	ds.MotionWorkflow_ID(id).Lazy(&c.ID)
	ds.MotionWorkflow_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionWorkflow_Name(id).Lazy(&c.Name)
	ds.MotionWorkflow_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.MotionWorkflow_StateIDs(id).Lazy(&c.StateIDs)
}

func (c *MotionWorkflow) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionWorkflow_DefaultAmendmentWorkflowMeetingID(id).Preload()
	ds.MotionWorkflow_DefaultWorkflowMeetingID(id).Preload()
	ds.MotionWorkflow_FirstStateID(id).Preload()
	ds.MotionWorkflow_ID(id).Preload()
	ds.MotionWorkflow_MeetingID(id).Preload()
	ds.MotionWorkflow_Name(id).Preload()
	ds.MotionWorkflow_SequentialNumber(id).Preload()
	ds.MotionWorkflow_StateIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionWorkflow(id int) *ValueCollection[MotionWorkflow, *MotionWorkflow] {
	return &ValueCollection[MotionWorkflow, *MotionWorkflow]{
		id:    id,
		fetch: r,
	}
}

// MotionWorkingGroupSpeaker has all fields from motion_working_group_speaker.
type MotionWorkingGroupSpeaker struct {
	ID            int
	MeetingID     int
	MeetingUserID int
	MotionID      int
	Weight        int
	fetch         *Fetch
}

func (c *MotionWorkingGroupSpeaker) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionWorkingGroupSpeaker_ID(id).Lazy(&c.ID)
	ds.MotionWorkingGroupSpeaker_MeetingID(id).Lazy(&c.MeetingID)
	ds.MotionWorkingGroupSpeaker_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.MotionWorkingGroupSpeaker_MotionID(id).Lazy(&c.MotionID)
	ds.MotionWorkingGroupSpeaker_Weight(id).Lazy(&c.Weight)
}

func (c *MotionWorkingGroupSpeaker) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.MotionWorkingGroupSpeaker_ID(id).Preload()
	ds.MotionWorkingGroupSpeaker_MeetingID(id).Preload()
	ds.MotionWorkingGroupSpeaker_MeetingUserID(id).Preload()
	ds.MotionWorkingGroupSpeaker_MotionID(id).Preload()
	ds.MotionWorkingGroupSpeaker_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) MotionWorkingGroupSpeaker(id int) *ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker] {
	return &ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{
		id:    id,
		fetch: r,
	}
}

// Option has all fields from option.
type Option struct {
	Abstain                    string
	ContentObjectID            Maybe[string]
	ID                         int
	MeetingID                  int
	No                         string
	PollID                     Maybe[int]
	Text                       string
	UsedAsGlobalOptionInPollID Maybe[int]
	VoteIDs                    []int
	Weight                     int
	Yes                        string
	fetch                      *Fetch
}

func (c *Option) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Option_Abstain(id).Lazy(&c.Abstain)
	ds.Option_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.Option_ID(id).Lazy(&c.ID)
	ds.Option_MeetingID(id).Lazy(&c.MeetingID)
	ds.Option_No(id).Lazy(&c.No)
	ds.Option_PollID(id).Lazy(&c.PollID)
	ds.Option_Text(id).Lazy(&c.Text)
	ds.Option_UsedAsGlobalOptionInPollID(id).Lazy(&c.UsedAsGlobalOptionInPollID)
	ds.Option_VoteIDs(id).Lazy(&c.VoteIDs)
	ds.Option_Weight(id).Lazy(&c.Weight)
	ds.Option_Yes(id).Lazy(&c.Yes)
}

func (c *Option) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Option_Abstain(id).Preload()
	ds.Option_ContentObjectID(id).Preload()
	ds.Option_ID(id).Preload()
	ds.Option_MeetingID(id).Preload()
	ds.Option_No(id).Preload()
	ds.Option_PollID(id).Preload()
	ds.Option_Text(id).Preload()
	ds.Option_UsedAsGlobalOptionInPollID(id).Preload()
	ds.Option_VoteIDs(id).Preload()
	ds.Option_Weight(id).Preload()
	ds.Option_Yes(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Option(id int) *ValueCollection[Option, *Option] {
	return &ValueCollection[Option, *Option]{
		id:    id,
		fetch: r,
	}
}

// Organization has all fields from organization.
type Organization struct {
	ActiveMeetingIDs           []int
	ArchivedMeetingIDs         []int
	CommitteeIDs               []int
	DefaultLanguage            string
	Description                string
	EnableAnonymous            bool
	EnableChat                 bool
	EnableElectronicVoting     bool
	GenderIDs                  []int
	ID                         int
	LegalNotice                string
	LimitOfMeetings            int
	LimitOfUsers               int
	LoginText                  string
	MediafileIDs               []int
	Name                       string
	OrganizationTagIDs         []int
	PrivacyPolicy              string
	PublishedMediafileIDs      []int
	RequireDuplicateFrom       bool
	ResetPasswordVerboseErrors bool
	SamlAttrMapping            json.RawMessage
	SamlEnabled                bool
	SamlLoginButtonText        string
	SamlMetadataIDp            string
	SamlMetadataSp             string
	SamlPrivateKey             string
	TemplateMeetingIDs         []int
	ThemeID                    int
	ThemeIDs                   []int
	Url                        string
	UserIDs                    []int
	UsersEmailBody             string
	UsersEmailReplyto          string
	UsersEmailSender           string
	UsersEmailSubject          string
	VoteDecryptPublicMainKey   string
	fetch                      *Fetch
}

func (c *Organization) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Organization_ActiveMeetingIDs(id).Lazy(&c.ActiveMeetingIDs)
	ds.Organization_ArchivedMeetingIDs(id).Lazy(&c.ArchivedMeetingIDs)
	ds.Organization_CommitteeIDs(id).Lazy(&c.CommitteeIDs)
	ds.Organization_DefaultLanguage(id).Lazy(&c.DefaultLanguage)
	ds.Organization_Description(id).Lazy(&c.Description)
	ds.Organization_EnableAnonymous(id).Lazy(&c.EnableAnonymous)
	ds.Organization_EnableChat(id).Lazy(&c.EnableChat)
	ds.Organization_EnableElectronicVoting(id).Lazy(&c.EnableElectronicVoting)
	ds.Organization_GenderIDs(id).Lazy(&c.GenderIDs)
	ds.Organization_ID(id).Lazy(&c.ID)
	ds.Organization_LegalNotice(id).Lazy(&c.LegalNotice)
	ds.Organization_LimitOfMeetings(id).Lazy(&c.LimitOfMeetings)
	ds.Organization_LimitOfUsers(id).Lazy(&c.LimitOfUsers)
	ds.Organization_LoginText(id).Lazy(&c.LoginText)
	ds.Organization_MediafileIDs(id).Lazy(&c.MediafileIDs)
	ds.Organization_Name(id).Lazy(&c.Name)
	ds.Organization_OrganizationTagIDs(id).Lazy(&c.OrganizationTagIDs)
	ds.Organization_PrivacyPolicy(id).Lazy(&c.PrivacyPolicy)
	ds.Organization_PublishedMediafileIDs(id).Lazy(&c.PublishedMediafileIDs)
	ds.Organization_RequireDuplicateFrom(id).Lazy(&c.RequireDuplicateFrom)
	ds.Organization_ResetPasswordVerboseErrors(id).Lazy(&c.ResetPasswordVerboseErrors)
	ds.Organization_SamlAttrMapping(id).Lazy(&c.SamlAttrMapping)
	ds.Organization_SamlEnabled(id).Lazy(&c.SamlEnabled)
	ds.Organization_SamlLoginButtonText(id).Lazy(&c.SamlLoginButtonText)
	ds.Organization_SamlMetadataIDp(id).Lazy(&c.SamlMetadataIDp)
	ds.Organization_SamlMetadataSp(id).Lazy(&c.SamlMetadataSp)
	ds.Organization_SamlPrivateKey(id).Lazy(&c.SamlPrivateKey)
	ds.Organization_TemplateMeetingIDs(id).Lazy(&c.TemplateMeetingIDs)
	ds.Organization_ThemeID(id).Lazy(&c.ThemeID)
	ds.Organization_ThemeIDs(id).Lazy(&c.ThemeIDs)
	ds.Organization_Url(id).Lazy(&c.Url)
	ds.Organization_UserIDs(id).Lazy(&c.UserIDs)
	ds.Organization_UsersEmailBody(id).Lazy(&c.UsersEmailBody)
	ds.Organization_UsersEmailReplyto(id).Lazy(&c.UsersEmailReplyto)
	ds.Organization_UsersEmailSender(id).Lazy(&c.UsersEmailSender)
	ds.Organization_UsersEmailSubject(id).Lazy(&c.UsersEmailSubject)
	ds.Organization_VoteDecryptPublicMainKey(id).Lazy(&c.VoteDecryptPublicMainKey)
}

func (c *Organization) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Organization_ActiveMeetingIDs(id).Preload()
	ds.Organization_ArchivedMeetingIDs(id).Preload()
	ds.Organization_CommitteeIDs(id).Preload()
	ds.Organization_DefaultLanguage(id).Preload()
	ds.Organization_Description(id).Preload()
	ds.Organization_EnableAnonymous(id).Preload()
	ds.Organization_EnableChat(id).Preload()
	ds.Organization_EnableElectronicVoting(id).Preload()
	ds.Organization_GenderIDs(id).Preload()
	ds.Organization_ID(id).Preload()
	ds.Organization_LegalNotice(id).Preload()
	ds.Organization_LimitOfMeetings(id).Preload()
	ds.Organization_LimitOfUsers(id).Preload()
	ds.Organization_LoginText(id).Preload()
	ds.Organization_MediafileIDs(id).Preload()
	ds.Organization_Name(id).Preload()
	ds.Organization_OrganizationTagIDs(id).Preload()
	ds.Organization_PrivacyPolicy(id).Preload()
	ds.Organization_PublishedMediafileIDs(id).Preload()
	ds.Organization_RequireDuplicateFrom(id).Preload()
	ds.Organization_ResetPasswordVerboseErrors(id).Preload()
	ds.Organization_SamlAttrMapping(id).Preload()
	ds.Organization_SamlEnabled(id).Preload()
	ds.Organization_SamlLoginButtonText(id).Preload()
	ds.Organization_SamlMetadataIDp(id).Preload()
	ds.Organization_SamlMetadataSp(id).Preload()
	ds.Organization_SamlPrivateKey(id).Preload()
	ds.Organization_TemplateMeetingIDs(id).Preload()
	ds.Organization_ThemeID(id).Preload()
	ds.Organization_ThemeIDs(id).Preload()
	ds.Organization_Url(id).Preload()
	ds.Organization_UserIDs(id).Preload()
	ds.Organization_UsersEmailBody(id).Preload()
	ds.Organization_UsersEmailReplyto(id).Preload()
	ds.Organization_UsersEmailSender(id).Preload()
	ds.Organization_UsersEmailSubject(id).Preload()
	ds.Organization_VoteDecryptPublicMainKey(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Organization(id int) *ValueCollection[Organization, *Organization] {
	return &ValueCollection[Organization, *Organization]{
		id:    id,
		fetch: r,
	}
}

// OrganizationTag has all fields from organization_tag.
type OrganizationTag struct {
	Color          string
	ID             int
	Name           string
	OrganizationID int
	TaggedIDs      []string
	fetch          *Fetch
}

func (c *OrganizationTag) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.OrganizationTag_Color(id).Lazy(&c.Color)
	ds.OrganizationTag_ID(id).Lazy(&c.ID)
	ds.OrganizationTag_Name(id).Lazy(&c.Name)
	ds.OrganizationTag_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.OrganizationTag_TaggedIDs(id).Lazy(&c.TaggedIDs)
}

func (c *OrganizationTag) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.OrganizationTag_Color(id).Preload()
	ds.OrganizationTag_ID(id).Preload()
	ds.OrganizationTag_Name(id).Preload()
	ds.OrganizationTag_OrganizationID(id).Preload()
	ds.OrganizationTag_TaggedIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) OrganizationTag(id int) *ValueCollection[OrganizationTag, *OrganizationTag] {
	return &ValueCollection[OrganizationTag, *OrganizationTag]{
		id:    id,
		fetch: r,
	}
}

// PersonalNote has all fields from personal_note.
type PersonalNote struct {
	ContentObjectID Maybe[string]
	ID              int
	MeetingID       int
	MeetingUserID   int
	Note            string
	Star            bool
	fetch           *Fetch
}

func (c *PersonalNote) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.PersonalNote_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.PersonalNote_ID(id).Lazy(&c.ID)
	ds.PersonalNote_MeetingID(id).Lazy(&c.MeetingID)
	ds.PersonalNote_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.PersonalNote_Note(id).Lazy(&c.Note)
	ds.PersonalNote_Star(id).Lazy(&c.Star)
}

func (c *PersonalNote) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.PersonalNote_ContentObjectID(id).Preload()
	ds.PersonalNote_ID(id).Preload()
	ds.PersonalNote_MeetingID(id).Preload()
	ds.PersonalNote_MeetingUserID(id).Preload()
	ds.PersonalNote_Note(id).Preload()
	ds.PersonalNote_Star(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) PersonalNote(id int) *ValueCollection[PersonalNote, *PersonalNote] {
	return &ValueCollection[PersonalNote, *PersonalNote]{
		id:    id,
		fetch: r,
	}
}

// PointOfOrderCategory has all fields from point_of_order_category.
type PointOfOrderCategory struct {
	ID         int
	MeetingID  int
	Rank       int
	SpeakerIDs []int
	Text       string
	fetch      *Fetch
}

func (c *PointOfOrderCategory) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.PointOfOrderCategory_ID(id).Lazy(&c.ID)
	ds.PointOfOrderCategory_MeetingID(id).Lazy(&c.MeetingID)
	ds.PointOfOrderCategory_Rank(id).Lazy(&c.Rank)
	ds.PointOfOrderCategory_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.PointOfOrderCategory_Text(id).Lazy(&c.Text)
}

func (c *PointOfOrderCategory) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.PointOfOrderCategory_ID(id).Preload()
	ds.PointOfOrderCategory_MeetingID(id).Preload()
	ds.PointOfOrderCategory_Rank(id).Preload()
	ds.PointOfOrderCategory_SpeakerIDs(id).Preload()
	ds.PointOfOrderCategory_Text(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) PointOfOrderCategory(id int) *ValueCollection[PointOfOrderCategory, *PointOfOrderCategory] {
	return &ValueCollection[PointOfOrderCategory, *PointOfOrderCategory]{
		id:    id,
		fetch: r,
	}
}

// Poll has all fields from poll.
type Poll struct {
	Backend               string
	ContentObjectID       string
	CryptKey              string
	CryptSignature        string
	Description           string
	EntitledGroupIDs      []int
	EntitledUsersAtStop   json.RawMessage
	GlobalAbstain         bool
	GlobalNo              bool
	GlobalOptionID        Maybe[int]
	GlobalYes             bool
	ID                    int
	IsPseudoanonymized    bool
	MaxVotesAmount        int
	MaxVotesPerOption     int
	MeetingID             int
	MinVotesAmount        int
	OnehundredPercentBase string
	OptionIDs             []int
	Pollmethod            string
	ProjectionIDs         []int
	SequentialNumber      int
	State                 string
	Title                 string
	Type                  string
	VoteCount             int
	VotedIDs              []int
	VotesRaw              string
	VotesSignature        string
	Votescast             string
	Votesinvalid          string
	Votesvalid            string
	fetch                 *Fetch
}

func (c *Poll) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Poll_Backend(id).Lazy(&c.Backend)
	ds.Poll_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.Poll_CryptKey(id).Lazy(&c.CryptKey)
	ds.Poll_CryptSignature(id).Lazy(&c.CryptSignature)
	ds.Poll_Description(id).Lazy(&c.Description)
	ds.Poll_EntitledGroupIDs(id).Lazy(&c.EntitledGroupIDs)
	ds.Poll_EntitledUsersAtStop(id).Lazy(&c.EntitledUsersAtStop)
	ds.Poll_GlobalAbstain(id).Lazy(&c.GlobalAbstain)
	ds.Poll_GlobalNo(id).Lazy(&c.GlobalNo)
	ds.Poll_GlobalOptionID(id).Lazy(&c.GlobalOptionID)
	ds.Poll_GlobalYes(id).Lazy(&c.GlobalYes)
	ds.Poll_ID(id).Lazy(&c.ID)
	ds.Poll_IsPseudoanonymized(id).Lazy(&c.IsPseudoanonymized)
	ds.Poll_MaxVotesAmount(id).Lazy(&c.MaxVotesAmount)
	ds.Poll_MaxVotesPerOption(id).Lazy(&c.MaxVotesPerOption)
	ds.Poll_MeetingID(id).Lazy(&c.MeetingID)
	ds.Poll_MinVotesAmount(id).Lazy(&c.MinVotesAmount)
	ds.Poll_OnehundredPercentBase(id).Lazy(&c.OnehundredPercentBase)
	ds.Poll_OptionIDs(id).Lazy(&c.OptionIDs)
	ds.Poll_Pollmethod(id).Lazy(&c.Pollmethod)
	ds.Poll_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Poll_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Poll_State(id).Lazy(&c.State)
	ds.Poll_Title(id).Lazy(&c.Title)
	ds.Poll_Type(id).Lazy(&c.Type)
	ds.Poll_VoteCount(id).Lazy(&c.VoteCount)
	ds.Poll_VotedIDs(id).Lazy(&c.VotedIDs)
	ds.Poll_VotesRaw(id).Lazy(&c.VotesRaw)
	ds.Poll_VotesSignature(id).Lazy(&c.VotesSignature)
	ds.Poll_Votescast(id).Lazy(&c.Votescast)
	ds.Poll_Votesinvalid(id).Lazy(&c.Votesinvalid)
	ds.Poll_Votesvalid(id).Lazy(&c.Votesvalid)
}

func (c *Poll) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Poll_Backend(id).Preload()
	ds.Poll_ContentObjectID(id).Preload()
	ds.Poll_CryptKey(id).Preload()
	ds.Poll_CryptSignature(id).Preload()
	ds.Poll_Description(id).Preload()
	ds.Poll_EntitledGroupIDs(id).Preload()
	ds.Poll_EntitledUsersAtStop(id).Preload()
	ds.Poll_GlobalAbstain(id).Preload()
	ds.Poll_GlobalNo(id).Preload()
	ds.Poll_GlobalOptionID(id).Preload()
	ds.Poll_GlobalYes(id).Preload()
	ds.Poll_ID(id).Preload()
	ds.Poll_IsPseudoanonymized(id).Preload()
	ds.Poll_MaxVotesAmount(id).Preload()
	ds.Poll_MaxVotesPerOption(id).Preload()
	ds.Poll_MeetingID(id).Preload()
	ds.Poll_MinVotesAmount(id).Preload()
	ds.Poll_OnehundredPercentBase(id).Preload()
	ds.Poll_OptionIDs(id).Preload()
	ds.Poll_Pollmethod(id).Preload()
	ds.Poll_ProjectionIDs(id).Preload()
	ds.Poll_SequentialNumber(id).Preload()
	ds.Poll_State(id).Preload()
	ds.Poll_Title(id).Preload()
	ds.Poll_Type(id).Preload()
	ds.Poll_VoteCount(id).Preload()
	ds.Poll_VotedIDs(id).Preload()
	ds.Poll_VotesRaw(id).Preload()
	ds.Poll_VotesSignature(id).Preload()
	ds.Poll_Votescast(id).Preload()
	ds.Poll_Votesinvalid(id).Preload()
	ds.Poll_Votesvalid(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Poll(id int) *ValueCollection[Poll, *Poll] {
	return &ValueCollection[Poll, *Poll]{
		id:    id,
		fetch: r,
	}
}

// PollCandidate has all fields from poll_candidate.
type PollCandidate struct {
	ID                  int
	MeetingID           int
	PollCandidateListID int
	UserID              Maybe[int]
	Weight              int
	fetch               *Fetch
}

func (c *PollCandidate) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.PollCandidate_ID(id).Lazy(&c.ID)
	ds.PollCandidate_MeetingID(id).Lazy(&c.MeetingID)
	ds.PollCandidate_PollCandidateListID(id).Lazy(&c.PollCandidateListID)
	ds.PollCandidate_UserID(id).Lazy(&c.UserID)
	ds.PollCandidate_Weight(id).Lazy(&c.Weight)
}

func (c *PollCandidate) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.PollCandidate_ID(id).Preload()
	ds.PollCandidate_MeetingID(id).Preload()
	ds.PollCandidate_PollCandidateListID(id).Preload()
	ds.PollCandidate_UserID(id).Preload()
	ds.PollCandidate_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) PollCandidate(id int) *ValueCollection[PollCandidate, *PollCandidate] {
	return &ValueCollection[PollCandidate, *PollCandidate]{
		id:    id,
		fetch: r,
	}
}

// PollCandidateList has all fields from poll_candidate_list.
type PollCandidateList struct {
	ID               int
	MeetingID        int
	OptionID         int
	PollCandidateIDs []int
	fetch            *Fetch
}

func (c *PollCandidateList) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.PollCandidateList_ID(id).Lazy(&c.ID)
	ds.PollCandidateList_MeetingID(id).Lazy(&c.MeetingID)
	ds.PollCandidateList_OptionID(id).Lazy(&c.OptionID)
	ds.PollCandidateList_PollCandidateIDs(id).Lazy(&c.PollCandidateIDs)
}

func (c *PollCandidateList) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.PollCandidateList_ID(id).Preload()
	ds.PollCandidateList_MeetingID(id).Preload()
	ds.PollCandidateList_OptionID(id).Preload()
	ds.PollCandidateList_PollCandidateIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) PollCandidateList(id int) *ValueCollection[PollCandidateList, *PollCandidateList] {
	return &ValueCollection[PollCandidateList, *PollCandidateList]{
		id:    id,
		fetch: r,
	}
}

// Projection has all fields from projection.
type Projection struct {
	Content            json.RawMessage
	ContentObjectID    string
	CurrentProjectorID Maybe[int]
	HistoryProjectorID Maybe[int]
	ID                 int
	MeetingID          int
	Options            json.RawMessage
	PreviewProjectorID Maybe[int]
	Stable             bool
	Type               string
	Weight             int
	fetch              *Fetch
}

func (c *Projection) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Projection_Content(id).Lazy(&c.Content)
	ds.Projection_ContentObjectID(id).Lazy(&c.ContentObjectID)
	ds.Projection_CurrentProjectorID(id).Lazy(&c.CurrentProjectorID)
	ds.Projection_HistoryProjectorID(id).Lazy(&c.HistoryProjectorID)
	ds.Projection_ID(id).Lazy(&c.ID)
	ds.Projection_MeetingID(id).Lazy(&c.MeetingID)
	ds.Projection_Options(id).Lazy(&c.Options)
	ds.Projection_PreviewProjectorID(id).Lazy(&c.PreviewProjectorID)
	ds.Projection_Stable(id).Lazy(&c.Stable)
	ds.Projection_Type(id).Lazy(&c.Type)
	ds.Projection_Weight(id).Lazy(&c.Weight)
}

func (c *Projection) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Projection_Content(id).Preload()
	ds.Projection_ContentObjectID(id).Preload()
	ds.Projection_CurrentProjectorID(id).Preload()
	ds.Projection_HistoryProjectorID(id).Preload()
	ds.Projection_ID(id).Preload()
	ds.Projection_MeetingID(id).Preload()
	ds.Projection_Options(id).Preload()
	ds.Projection_PreviewProjectorID(id).Preload()
	ds.Projection_Stable(id).Preload()
	ds.Projection_Type(id).Preload()
	ds.Projection_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Projection(id int) *ValueCollection[Projection, *Projection] {
	return &ValueCollection[Projection, *Projection]{
		id:    id,
		fetch: r,
	}
}

// Projector has all fields from projector.
type Projector struct {
	AspectRatioDenominator                                    int
	AspectRatioNumerator                                      int
	BackgroundColor                                           string
	ChyronBackgroundColor                                     string
	ChyronBackgroundColor2                                    string
	ChyronFontColor                                           string
	ChyronFontColor2                                          string
	Color                                                     string
	CurrentProjectionIDs                                      []int
	HeaderBackgroundColor                                     string
	HeaderFontColor                                           string
	HeaderH1Color                                             string
	HistoryProjectionIDs                                      []int
	ID                                                        int
	IsInternal                                                bool
	MeetingID                                                 int
	Name                                                      string
	PreviewProjectionIDs                                      []int
	Scale                                                     int
	Scroll                                                    int
	SequentialNumber                                          int
	ShowClock                                                 bool
	ShowHeaderFooter                                          bool
	ShowLogo                                                  bool
	ShowTitle                                                 bool
	UsedAsDefaultProjectorForAgendaItemListInMeetingID        Maybe[int]
	UsedAsDefaultProjectorForAmendmentInMeetingID             Maybe[int]
	UsedAsDefaultProjectorForAssignmentInMeetingID            Maybe[int]
	UsedAsDefaultProjectorForAssignmentPollInMeetingID        Maybe[int]
	UsedAsDefaultProjectorForCountdownInMeetingID             Maybe[int]
	UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID Maybe[int]
	UsedAsDefaultProjectorForListOfSpeakersInMeetingID        Maybe[int]
	UsedAsDefaultProjectorForMediafileInMeetingID             Maybe[int]
	UsedAsDefaultProjectorForMessageInMeetingID               Maybe[int]
	UsedAsDefaultProjectorForMotionBlockInMeetingID           Maybe[int]
	UsedAsDefaultProjectorForMotionInMeetingID                Maybe[int]
	UsedAsDefaultProjectorForMotionPollInMeetingID            Maybe[int]
	UsedAsDefaultProjectorForPollInMeetingID                  Maybe[int]
	UsedAsDefaultProjectorForTopicInMeetingID                 Maybe[int]
	UsedAsReferenceProjectorMeetingID                         Maybe[int]
	Width                                                     int
	fetch                                                     *Fetch
}

func (c *Projector) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Projector_AspectRatioDenominator(id).Lazy(&c.AspectRatioDenominator)
	ds.Projector_AspectRatioNumerator(id).Lazy(&c.AspectRatioNumerator)
	ds.Projector_BackgroundColor(id).Lazy(&c.BackgroundColor)
	ds.Projector_ChyronBackgroundColor(id).Lazy(&c.ChyronBackgroundColor)
	ds.Projector_ChyronBackgroundColor2(id).Lazy(&c.ChyronBackgroundColor2)
	ds.Projector_ChyronFontColor(id).Lazy(&c.ChyronFontColor)
	ds.Projector_ChyronFontColor2(id).Lazy(&c.ChyronFontColor2)
	ds.Projector_Color(id).Lazy(&c.Color)
	ds.Projector_CurrentProjectionIDs(id).Lazy(&c.CurrentProjectionIDs)
	ds.Projector_HeaderBackgroundColor(id).Lazy(&c.HeaderBackgroundColor)
	ds.Projector_HeaderFontColor(id).Lazy(&c.HeaderFontColor)
	ds.Projector_HeaderH1Color(id).Lazy(&c.HeaderH1Color)
	ds.Projector_HistoryProjectionIDs(id).Lazy(&c.HistoryProjectionIDs)
	ds.Projector_ID(id).Lazy(&c.ID)
	ds.Projector_IsInternal(id).Lazy(&c.IsInternal)
	ds.Projector_MeetingID(id).Lazy(&c.MeetingID)
	ds.Projector_Name(id).Lazy(&c.Name)
	ds.Projector_PreviewProjectionIDs(id).Lazy(&c.PreviewProjectionIDs)
	ds.Projector_Scale(id).Lazy(&c.Scale)
	ds.Projector_Scroll(id).Lazy(&c.Scroll)
	ds.Projector_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Projector_ShowClock(id).Lazy(&c.ShowClock)
	ds.Projector_ShowHeaderFooter(id).Lazy(&c.ShowHeaderFooter)
	ds.Projector_ShowLogo(id).Lazy(&c.ShowLogo)
	ds.Projector_ShowTitle(id).Lazy(&c.ShowTitle)
	ds.Projector_UsedAsDefaultProjectorForAgendaItemListInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForAgendaItemListInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForAmendmentInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForAmendmentInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForAssignmentInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForAssignmentInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForAssignmentPollInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForAssignmentPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForCountdownInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForCountdownInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForListOfSpeakersInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForListOfSpeakersInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMediafileInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMediafileInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMessageInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMessageInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMotionBlockInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMotionBlockInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMotionInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMotionInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMotionPollInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForMotionPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForPollInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForTopicInMeetingID(id).Lazy(&c.UsedAsDefaultProjectorForTopicInMeetingID)
	ds.Projector_UsedAsReferenceProjectorMeetingID(id).Lazy(&c.UsedAsReferenceProjectorMeetingID)
	ds.Projector_Width(id).Lazy(&c.Width)
}

func (c *Projector) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Projector_AspectRatioDenominator(id).Preload()
	ds.Projector_AspectRatioNumerator(id).Preload()
	ds.Projector_BackgroundColor(id).Preload()
	ds.Projector_ChyronBackgroundColor(id).Preload()
	ds.Projector_ChyronBackgroundColor2(id).Preload()
	ds.Projector_ChyronFontColor(id).Preload()
	ds.Projector_ChyronFontColor2(id).Preload()
	ds.Projector_Color(id).Preload()
	ds.Projector_CurrentProjectionIDs(id).Preload()
	ds.Projector_HeaderBackgroundColor(id).Preload()
	ds.Projector_HeaderFontColor(id).Preload()
	ds.Projector_HeaderH1Color(id).Preload()
	ds.Projector_HistoryProjectionIDs(id).Preload()
	ds.Projector_ID(id).Preload()
	ds.Projector_IsInternal(id).Preload()
	ds.Projector_MeetingID(id).Preload()
	ds.Projector_Name(id).Preload()
	ds.Projector_PreviewProjectionIDs(id).Preload()
	ds.Projector_Scale(id).Preload()
	ds.Projector_Scroll(id).Preload()
	ds.Projector_SequentialNumber(id).Preload()
	ds.Projector_ShowClock(id).Preload()
	ds.Projector_ShowHeaderFooter(id).Preload()
	ds.Projector_ShowLogo(id).Preload()
	ds.Projector_ShowTitle(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForAgendaItemListInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForAmendmentInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForAssignmentInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForAssignmentPollInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForCountdownInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForListOfSpeakersInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMediafileInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMessageInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMotionBlockInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMotionInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMotionPollInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForPollInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForTopicInMeetingID(id).Preload()
	ds.Projector_UsedAsReferenceProjectorMeetingID(id).Preload()
	ds.Projector_Width(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Projector(id int) *ValueCollection[Projector, *Projector] {
	return &ValueCollection[Projector, *Projector]{
		id:    id,
		fetch: r,
	}
}

// ProjectorCountdown has all fields from projector_countdown.
type ProjectorCountdown struct {
	CountdownTime                          float32
	DefaultTime                            int
	Description                            string
	ID                                     int
	MeetingID                              int
	ProjectionIDs                          []int
	Running                                bool
	Title                                  string
	UsedAsListOfSpeakersCountdownMeetingID Maybe[int]
	UsedAsPollCountdownMeetingID           Maybe[int]
	fetch                                  *Fetch
}

func (c *ProjectorCountdown) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ProjectorCountdown_CountdownTime(id).Lazy(&c.CountdownTime)
	ds.ProjectorCountdown_DefaultTime(id).Lazy(&c.DefaultTime)
	ds.ProjectorCountdown_Description(id).Lazy(&c.Description)
	ds.ProjectorCountdown_ID(id).Lazy(&c.ID)
	ds.ProjectorCountdown_MeetingID(id).Lazy(&c.MeetingID)
	ds.ProjectorCountdown_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.ProjectorCountdown_Running(id).Lazy(&c.Running)
	ds.ProjectorCountdown_Title(id).Lazy(&c.Title)
	ds.ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(id).Lazy(&c.UsedAsListOfSpeakersCountdownMeetingID)
	ds.ProjectorCountdown_UsedAsPollCountdownMeetingID(id).Lazy(&c.UsedAsPollCountdownMeetingID)
}

func (c *ProjectorCountdown) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.ProjectorCountdown_CountdownTime(id).Preload()
	ds.ProjectorCountdown_DefaultTime(id).Preload()
	ds.ProjectorCountdown_Description(id).Preload()
	ds.ProjectorCountdown_ID(id).Preload()
	ds.ProjectorCountdown_MeetingID(id).Preload()
	ds.ProjectorCountdown_ProjectionIDs(id).Preload()
	ds.ProjectorCountdown_Running(id).Preload()
	ds.ProjectorCountdown_Title(id).Preload()
	ds.ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(id).Preload()
	ds.ProjectorCountdown_UsedAsPollCountdownMeetingID(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) ProjectorCountdown(id int) *ValueCollection[ProjectorCountdown, *ProjectorCountdown] {
	return &ValueCollection[ProjectorCountdown, *ProjectorCountdown]{
		id:    id,
		fetch: r,
	}
}

// ProjectorMessage has all fields from projector_message.
type ProjectorMessage struct {
	ID            int
	MeetingID     int
	Message       string
	ProjectionIDs []int
	fetch         *Fetch
}

func (c *ProjectorMessage) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.ProjectorMessage_ID(id).Lazy(&c.ID)
	ds.ProjectorMessage_MeetingID(id).Lazy(&c.MeetingID)
	ds.ProjectorMessage_Message(id).Lazy(&c.Message)
	ds.ProjectorMessage_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
}

func (c *ProjectorMessage) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.ProjectorMessage_ID(id).Preload()
	ds.ProjectorMessage_MeetingID(id).Preload()
	ds.ProjectorMessage_Message(id).Preload()
	ds.ProjectorMessage_ProjectionIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) ProjectorMessage(id int) *ValueCollection[ProjectorMessage, *ProjectorMessage] {
	return &ValueCollection[ProjectorMessage, *ProjectorMessage]{
		id:    id,
		fetch: r,
	}
}

// Speaker has all fields from speaker.
type Speaker struct {
	BeginTime                      int
	EndTime                        int
	ID                             int
	ListOfSpeakersID               int
	MeetingID                      int
	MeetingUserID                  Maybe[int]
	Note                           string
	PauseTime                      int
	PointOfOrder                   bool
	PointOfOrderCategoryID         Maybe[int]
	SpeechState                    string
	StructureLevelListOfSpeakersID Maybe[int]
	TotalPause                     int
	UnpauseTime                    int
	Weight                         int
	fetch                          *Fetch
}

func (c *Speaker) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Speaker_BeginTime(id).Lazy(&c.BeginTime)
	ds.Speaker_EndTime(id).Lazy(&c.EndTime)
	ds.Speaker_ID(id).Lazy(&c.ID)
	ds.Speaker_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.Speaker_MeetingID(id).Lazy(&c.MeetingID)
	ds.Speaker_MeetingUserID(id).Lazy(&c.MeetingUserID)
	ds.Speaker_Note(id).Lazy(&c.Note)
	ds.Speaker_PauseTime(id).Lazy(&c.PauseTime)
	ds.Speaker_PointOfOrder(id).Lazy(&c.PointOfOrder)
	ds.Speaker_PointOfOrderCategoryID(id).Lazy(&c.PointOfOrderCategoryID)
	ds.Speaker_SpeechState(id).Lazy(&c.SpeechState)
	ds.Speaker_StructureLevelListOfSpeakersID(id).Lazy(&c.StructureLevelListOfSpeakersID)
	ds.Speaker_TotalPause(id).Lazy(&c.TotalPause)
	ds.Speaker_UnpauseTime(id).Lazy(&c.UnpauseTime)
	ds.Speaker_Weight(id).Lazy(&c.Weight)
}

func (c *Speaker) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Speaker_BeginTime(id).Preload()
	ds.Speaker_EndTime(id).Preload()
	ds.Speaker_ID(id).Preload()
	ds.Speaker_ListOfSpeakersID(id).Preload()
	ds.Speaker_MeetingID(id).Preload()
	ds.Speaker_MeetingUserID(id).Preload()
	ds.Speaker_Note(id).Preload()
	ds.Speaker_PauseTime(id).Preload()
	ds.Speaker_PointOfOrder(id).Preload()
	ds.Speaker_PointOfOrderCategoryID(id).Preload()
	ds.Speaker_SpeechState(id).Preload()
	ds.Speaker_StructureLevelListOfSpeakersID(id).Preload()
	ds.Speaker_TotalPause(id).Preload()
	ds.Speaker_UnpauseTime(id).Preload()
	ds.Speaker_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Speaker(id int) *ValueCollection[Speaker, *Speaker] {
	return &ValueCollection[Speaker, *Speaker]{
		id:    id,
		fetch: r,
	}
}

// StructureLevel has all fields from structure_level.
type StructureLevel struct {
	Color                           string
	DefaultTime                     int
	ID                              int
	MeetingID                       int
	MeetingUserIDs                  []int
	Name                            string
	StructureLevelListOfSpeakersIDs []int
	fetch                           *Fetch
}

func (c *StructureLevel) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.StructureLevel_Color(id).Lazy(&c.Color)
	ds.StructureLevel_DefaultTime(id).Lazy(&c.DefaultTime)
	ds.StructureLevel_ID(id).Lazy(&c.ID)
	ds.StructureLevel_MeetingID(id).Lazy(&c.MeetingID)
	ds.StructureLevel_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.StructureLevel_Name(id).Lazy(&c.Name)
	ds.StructureLevel_StructureLevelListOfSpeakersIDs(id).Lazy(&c.StructureLevelListOfSpeakersIDs)
}

func (c *StructureLevel) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.StructureLevel_Color(id).Preload()
	ds.StructureLevel_DefaultTime(id).Preload()
	ds.StructureLevel_ID(id).Preload()
	ds.StructureLevel_MeetingID(id).Preload()
	ds.StructureLevel_MeetingUserIDs(id).Preload()
	ds.StructureLevel_Name(id).Preload()
	ds.StructureLevel_StructureLevelListOfSpeakersIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) StructureLevel(id int) *ValueCollection[StructureLevel, *StructureLevel] {
	return &ValueCollection[StructureLevel, *StructureLevel]{
		id:    id,
		fetch: r,
	}
}

// StructureLevelListOfSpeakers has all fields from structure_level_list_of_speakers.
type StructureLevelListOfSpeakers struct {
	AdditionalTime   float32
	CurrentStartTime int
	ID               int
	InitialTime      int
	ListOfSpeakersID int
	MeetingID        int
	RemainingTime    float32
	SpeakerIDs       []int
	StructureLevelID int
	fetch            *Fetch
}

func (c *StructureLevelListOfSpeakers) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.StructureLevelListOfSpeakers_AdditionalTime(id).Lazy(&c.AdditionalTime)
	ds.StructureLevelListOfSpeakers_CurrentStartTime(id).Lazy(&c.CurrentStartTime)
	ds.StructureLevelListOfSpeakers_ID(id).Lazy(&c.ID)
	ds.StructureLevelListOfSpeakers_InitialTime(id).Lazy(&c.InitialTime)
	ds.StructureLevelListOfSpeakers_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.StructureLevelListOfSpeakers_MeetingID(id).Lazy(&c.MeetingID)
	ds.StructureLevelListOfSpeakers_RemainingTime(id).Lazy(&c.RemainingTime)
	ds.StructureLevelListOfSpeakers_SpeakerIDs(id).Lazy(&c.SpeakerIDs)
	ds.StructureLevelListOfSpeakers_StructureLevelID(id).Lazy(&c.StructureLevelID)
}

func (c *StructureLevelListOfSpeakers) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.StructureLevelListOfSpeakers_AdditionalTime(id).Preload()
	ds.StructureLevelListOfSpeakers_CurrentStartTime(id).Preload()
	ds.StructureLevelListOfSpeakers_ID(id).Preload()
	ds.StructureLevelListOfSpeakers_InitialTime(id).Preload()
	ds.StructureLevelListOfSpeakers_ListOfSpeakersID(id).Preload()
	ds.StructureLevelListOfSpeakers_MeetingID(id).Preload()
	ds.StructureLevelListOfSpeakers_RemainingTime(id).Preload()
	ds.StructureLevelListOfSpeakers_SpeakerIDs(id).Preload()
	ds.StructureLevelListOfSpeakers_StructureLevelID(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) StructureLevelListOfSpeakers(id int) *ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers] {
	return &ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{
		id:    id,
		fetch: r,
	}
}

// Tag has all fields from tag.
type Tag struct {
	ID        int
	MeetingID int
	Name      string
	TaggedIDs []string
	fetch     *Fetch
}

func (c *Tag) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Tag_ID(id).Lazy(&c.ID)
	ds.Tag_MeetingID(id).Lazy(&c.MeetingID)
	ds.Tag_Name(id).Lazy(&c.Name)
	ds.Tag_TaggedIDs(id).Lazy(&c.TaggedIDs)
}

func (c *Tag) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Tag_ID(id).Preload()
	ds.Tag_MeetingID(id).Preload()
	ds.Tag_Name(id).Preload()
	ds.Tag_TaggedIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Tag(id int) *ValueCollection[Tag, *Tag] {
	return &ValueCollection[Tag, *Tag]{
		id:    id,
		fetch: r,
	}
}

// Theme has all fields from theme.
type Theme struct {
	Abstain                string
	Accent100              string
	Accent200              string
	Accent300              string
	Accent400              string
	Accent50               string
	Accent500              string
	Accent600              string
	Accent700              string
	Accent800              string
	Accent900              string
	AccentA100             string
	AccentA200             string
	AccentA400             string
	AccentA700             string
	Headbar                string
	ID                     int
	Name                   string
	No                     string
	OrganizationID         int
	Primary100             string
	Primary200             string
	Primary300             string
	Primary400             string
	Primary50              string
	Primary500             string
	Primary600             string
	Primary700             string
	Primary800             string
	Primary900             string
	PrimaryA100            string
	PrimaryA200            string
	PrimaryA400            string
	PrimaryA700            string
	ThemeForOrganizationID Maybe[int]
	Warn100                string
	Warn200                string
	Warn300                string
	Warn400                string
	Warn50                 string
	Warn500                string
	Warn600                string
	Warn700                string
	Warn800                string
	Warn900                string
	WarnA100               string
	WarnA200               string
	WarnA400               string
	WarnA700               string
	Yes                    string
	fetch                  *Fetch
}

func (c *Theme) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Theme_Abstain(id).Lazy(&c.Abstain)
	ds.Theme_Accent100(id).Lazy(&c.Accent100)
	ds.Theme_Accent200(id).Lazy(&c.Accent200)
	ds.Theme_Accent300(id).Lazy(&c.Accent300)
	ds.Theme_Accent400(id).Lazy(&c.Accent400)
	ds.Theme_Accent50(id).Lazy(&c.Accent50)
	ds.Theme_Accent500(id).Lazy(&c.Accent500)
	ds.Theme_Accent600(id).Lazy(&c.Accent600)
	ds.Theme_Accent700(id).Lazy(&c.Accent700)
	ds.Theme_Accent800(id).Lazy(&c.Accent800)
	ds.Theme_Accent900(id).Lazy(&c.Accent900)
	ds.Theme_AccentA100(id).Lazy(&c.AccentA100)
	ds.Theme_AccentA200(id).Lazy(&c.AccentA200)
	ds.Theme_AccentA400(id).Lazy(&c.AccentA400)
	ds.Theme_AccentA700(id).Lazy(&c.AccentA700)
	ds.Theme_Headbar(id).Lazy(&c.Headbar)
	ds.Theme_ID(id).Lazy(&c.ID)
	ds.Theme_Name(id).Lazy(&c.Name)
	ds.Theme_No(id).Lazy(&c.No)
	ds.Theme_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.Theme_Primary100(id).Lazy(&c.Primary100)
	ds.Theme_Primary200(id).Lazy(&c.Primary200)
	ds.Theme_Primary300(id).Lazy(&c.Primary300)
	ds.Theme_Primary400(id).Lazy(&c.Primary400)
	ds.Theme_Primary50(id).Lazy(&c.Primary50)
	ds.Theme_Primary500(id).Lazy(&c.Primary500)
	ds.Theme_Primary600(id).Lazy(&c.Primary600)
	ds.Theme_Primary700(id).Lazy(&c.Primary700)
	ds.Theme_Primary800(id).Lazy(&c.Primary800)
	ds.Theme_Primary900(id).Lazy(&c.Primary900)
	ds.Theme_PrimaryA100(id).Lazy(&c.PrimaryA100)
	ds.Theme_PrimaryA200(id).Lazy(&c.PrimaryA200)
	ds.Theme_PrimaryA400(id).Lazy(&c.PrimaryA400)
	ds.Theme_PrimaryA700(id).Lazy(&c.PrimaryA700)
	ds.Theme_ThemeForOrganizationID(id).Lazy(&c.ThemeForOrganizationID)
	ds.Theme_Warn100(id).Lazy(&c.Warn100)
	ds.Theme_Warn200(id).Lazy(&c.Warn200)
	ds.Theme_Warn300(id).Lazy(&c.Warn300)
	ds.Theme_Warn400(id).Lazy(&c.Warn400)
	ds.Theme_Warn50(id).Lazy(&c.Warn50)
	ds.Theme_Warn500(id).Lazy(&c.Warn500)
	ds.Theme_Warn600(id).Lazy(&c.Warn600)
	ds.Theme_Warn700(id).Lazy(&c.Warn700)
	ds.Theme_Warn800(id).Lazy(&c.Warn800)
	ds.Theme_Warn900(id).Lazy(&c.Warn900)
	ds.Theme_WarnA100(id).Lazy(&c.WarnA100)
	ds.Theme_WarnA200(id).Lazy(&c.WarnA200)
	ds.Theme_WarnA400(id).Lazy(&c.WarnA400)
	ds.Theme_WarnA700(id).Lazy(&c.WarnA700)
	ds.Theme_Yes(id).Lazy(&c.Yes)
}

func (c *Theme) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Theme_Abstain(id).Preload()
	ds.Theme_Accent100(id).Preload()
	ds.Theme_Accent200(id).Preload()
	ds.Theme_Accent300(id).Preload()
	ds.Theme_Accent400(id).Preload()
	ds.Theme_Accent50(id).Preload()
	ds.Theme_Accent500(id).Preload()
	ds.Theme_Accent600(id).Preload()
	ds.Theme_Accent700(id).Preload()
	ds.Theme_Accent800(id).Preload()
	ds.Theme_Accent900(id).Preload()
	ds.Theme_AccentA100(id).Preload()
	ds.Theme_AccentA200(id).Preload()
	ds.Theme_AccentA400(id).Preload()
	ds.Theme_AccentA700(id).Preload()
	ds.Theme_Headbar(id).Preload()
	ds.Theme_ID(id).Preload()
	ds.Theme_Name(id).Preload()
	ds.Theme_No(id).Preload()
	ds.Theme_OrganizationID(id).Preload()
	ds.Theme_Primary100(id).Preload()
	ds.Theme_Primary200(id).Preload()
	ds.Theme_Primary300(id).Preload()
	ds.Theme_Primary400(id).Preload()
	ds.Theme_Primary50(id).Preload()
	ds.Theme_Primary500(id).Preload()
	ds.Theme_Primary600(id).Preload()
	ds.Theme_Primary700(id).Preload()
	ds.Theme_Primary800(id).Preload()
	ds.Theme_Primary900(id).Preload()
	ds.Theme_PrimaryA100(id).Preload()
	ds.Theme_PrimaryA200(id).Preload()
	ds.Theme_PrimaryA400(id).Preload()
	ds.Theme_PrimaryA700(id).Preload()
	ds.Theme_ThemeForOrganizationID(id).Preload()
	ds.Theme_Warn100(id).Preload()
	ds.Theme_Warn200(id).Preload()
	ds.Theme_Warn300(id).Preload()
	ds.Theme_Warn400(id).Preload()
	ds.Theme_Warn50(id).Preload()
	ds.Theme_Warn500(id).Preload()
	ds.Theme_Warn600(id).Preload()
	ds.Theme_Warn700(id).Preload()
	ds.Theme_Warn800(id).Preload()
	ds.Theme_Warn900(id).Preload()
	ds.Theme_WarnA100(id).Preload()
	ds.Theme_WarnA200(id).Preload()
	ds.Theme_WarnA400(id).Preload()
	ds.Theme_WarnA700(id).Preload()
	ds.Theme_Yes(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Theme(id int) *ValueCollection[Theme, *Theme] {
	return &ValueCollection[Theme, *Theme]{
		id:    id,
		fetch: r,
	}
}

// Topic has all fields from topic.
type Topic struct {
	AgendaItemID                  int
	AttachmentMeetingMediafileIDs []int
	ID                            int
	ListOfSpeakersID              int
	MeetingID                     int
	PollIDs                       []int
	ProjectionIDs                 []int
	SequentialNumber              int
	Text                          string
	Title                         string
	fetch                         *Fetch
}

func (c *Topic) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Topic_AgendaItemID(id).Lazy(&c.AgendaItemID)
	ds.Topic_AttachmentMeetingMediafileIDs(id).Lazy(&c.AttachmentMeetingMediafileIDs)
	ds.Topic_ID(id).Lazy(&c.ID)
	ds.Topic_ListOfSpeakersID(id).Lazy(&c.ListOfSpeakersID)
	ds.Topic_MeetingID(id).Lazy(&c.MeetingID)
	ds.Topic_PollIDs(id).Lazy(&c.PollIDs)
	ds.Topic_ProjectionIDs(id).Lazy(&c.ProjectionIDs)
	ds.Topic_SequentialNumber(id).Lazy(&c.SequentialNumber)
	ds.Topic_Text(id).Lazy(&c.Text)
	ds.Topic_Title(id).Lazy(&c.Title)
}

func (c *Topic) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Topic_AgendaItemID(id).Preload()
	ds.Topic_AttachmentMeetingMediafileIDs(id).Preload()
	ds.Topic_ID(id).Preload()
	ds.Topic_ListOfSpeakersID(id).Preload()
	ds.Topic_MeetingID(id).Preload()
	ds.Topic_PollIDs(id).Preload()
	ds.Topic_ProjectionIDs(id).Preload()
	ds.Topic_SequentialNumber(id).Preload()
	ds.Topic_Text(id).Preload()
	ds.Topic_Title(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Topic(id int) *ValueCollection[Topic, *Topic] {
	return &ValueCollection[Topic, *Topic]{
		id:    id,
		fetch: r,
	}
}

// User has all fields from user.
type User struct {
	CanChangeOwnPassword        bool
	CommitteeIDs                []int
	CommitteeManagementIDs      []int
	DefaultPassword             string
	DefaultVoteWeight           string
	DelegatedVoteIDs            []int
	Email                       string
	FirstName                   string
	ForwardingCommitteeIDs      []int
	GenderID                    Maybe[int]
	ID                          int
	IsActive                    bool
	IsDemoUser                  bool
	IsPhysicalPerson            bool
	IsPresentInMeetingIDs       []int
	LastEmailSent               int
	LastLogin                   int
	LastName                    string
	MeetingIDs                  []int
	MeetingUserIDs              []int
	MemberNumber                string
	OptionIDs                   []int
	OrganizationID              int
	OrganizationManagementLevel string
	Password                    string
	PollCandidateIDs            []int
	PollVotedIDs                []int
	Pronoun                     string
	SamlID                      string
	Title                       string
	Username                    string
	VoteIDs                     []int
	fetch                       *Fetch
}

func (c *User) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.User_CanChangeOwnPassword(id).Lazy(&c.CanChangeOwnPassword)
	ds.User_CommitteeIDs(id).Lazy(&c.CommitteeIDs)
	ds.User_CommitteeManagementIDs(id).Lazy(&c.CommitteeManagementIDs)
	ds.User_DefaultPassword(id).Lazy(&c.DefaultPassword)
	ds.User_DefaultVoteWeight(id).Lazy(&c.DefaultVoteWeight)
	ds.User_DelegatedVoteIDs(id).Lazy(&c.DelegatedVoteIDs)
	ds.User_Email(id).Lazy(&c.Email)
	ds.User_FirstName(id).Lazy(&c.FirstName)
	ds.User_ForwardingCommitteeIDs(id).Lazy(&c.ForwardingCommitteeIDs)
	ds.User_GenderID(id).Lazy(&c.GenderID)
	ds.User_ID(id).Lazy(&c.ID)
	ds.User_IsActive(id).Lazy(&c.IsActive)
	ds.User_IsDemoUser(id).Lazy(&c.IsDemoUser)
	ds.User_IsPhysicalPerson(id).Lazy(&c.IsPhysicalPerson)
	ds.User_IsPresentInMeetingIDs(id).Lazy(&c.IsPresentInMeetingIDs)
	ds.User_LastEmailSent(id).Lazy(&c.LastEmailSent)
	ds.User_LastLogin(id).Lazy(&c.LastLogin)
	ds.User_LastName(id).Lazy(&c.LastName)
	ds.User_MeetingIDs(id).Lazy(&c.MeetingIDs)
	ds.User_MeetingUserIDs(id).Lazy(&c.MeetingUserIDs)
	ds.User_MemberNumber(id).Lazy(&c.MemberNumber)
	ds.User_OptionIDs(id).Lazy(&c.OptionIDs)
	ds.User_OrganizationID(id).Lazy(&c.OrganizationID)
	ds.User_OrganizationManagementLevel(id).Lazy(&c.OrganizationManagementLevel)
	ds.User_Password(id).Lazy(&c.Password)
	ds.User_PollCandidateIDs(id).Lazy(&c.PollCandidateIDs)
	ds.User_PollVotedIDs(id).Lazy(&c.PollVotedIDs)
	ds.User_Pronoun(id).Lazy(&c.Pronoun)
	ds.User_SamlID(id).Lazy(&c.SamlID)
	ds.User_Title(id).Lazy(&c.Title)
	ds.User_Username(id).Lazy(&c.Username)
	ds.User_VoteIDs(id).Lazy(&c.VoteIDs)
}

func (c *User) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.User_CanChangeOwnPassword(id).Preload()
	ds.User_CommitteeIDs(id).Preload()
	ds.User_CommitteeManagementIDs(id).Preload()
	ds.User_DefaultPassword(id).Preload()
	ds.User_DefaultVoteWeight(id).Preload()
	ds.User_DelegatedVoteIDs(id).Preload()
	ds.User_Email(id).Preload()
	ds.User_FirstName(id).Preload()
	ds.User_ForwardingCommitteeIDs(id).Preload()
	ds.User_GenderID(id).Preload()
	ds.User_ID(id).Preload()
	ds.User_IsActive(id).Preload()
	ds.User_IsDemoUser(id).Preload()
	ds.User_IsPhysicalPerson(id).Preload()
	ds.User_IsPresentInMeetingIDs(id).Preload()
	ds.User_LastEmailSent(id).Preload()
	ds.User_LastLogin(id).Preload()
	ds.User_LastName(id).Preload()
	ds.User_MeetingIDs(id).Preload()
	ds.User_MeetingUserIDs(id).Preload()
	ds.User_MemberNumber(id).Preload()
	ds.User_OptionIDs(id).Preload()
	ds.User_OrganizationID(id).Preload()
	ds.User_OrganizationManagementLevel(id).Preload()
	ds.User_Password(id).Preload()
	ds.User_PollCandidateIDs(id).Preload()
	ds.User_PollVotedIDs(id).Preload()
	ds.User_Pronoun(id).Preload()
	ds.User_SamlID(id).Preload()
	ds.User_Title(id).Preload()
	ds.User_Username(id).Preload()
	ds.User_VoteIDs(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) User(id int) *ValueCollection[User, *User] {
	return &ValueCollection[User, *User]{
		id:    id,
		fetch: r,
	}
}

// Vote has all fields from vote.
type Vote struct {
	DelegatedUserID Maybe[int]
	ID              int
	MeetingID       int
	OptionID        int
	UserID          Maybe[int]
	UserToken       string
	Value           string
	Weight          string
	fetch           *Fetch
}

func (c *Vote) lazy(ds *Fetch, id int) {
	c.fetch = ds
	ds.Vote_DelegatedUserID(id).Lazy(&c.DelegatedUserID)
	ds.Vote_ID(id).Lazy(&c.ID)
	ds.Vote_MeetingID(id).Lazy(&c.MeetingID)
	ds.Vote_OptionID(id).Lazy(&c.OptionID)
	ds.Vote_UserID(id).Lazy(&c.UserID)
	ds.Vote_UserToken(id).Lazy(&c.UserToken)
	ds.Vote_Value(id).Lazy(&c.Value)
	ds.Vote_Weight(id).Lazy(&c.Weight)
}

func (c *Vote) preload(ds *Fetch, id int) {
	c.fetch = ds
	ds.Vote_DelegatedUserID(id).Preload()
	ds.Vote_ID(id).Preload()
	ds.Vote_MeetingID(id).Preload()
	ds.Vote_OptionID(id).Preload()
	ds.Vote_UserID(id).Preload()
	ds.Vote_UserToken(id).Preload()
	ds.Vote_Value(id).Preload()
	ds.Vote_Weight(id).Preload()
}

// TODO: Generate functions for all relations

func (r *Fetch) Vote(id int) *ValueCollection[Vote, *Vote] {
	return &ValueCollection[Vote, *Vote]{
		id:    id,
		fetch: r,
	}
}
