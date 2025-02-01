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

// ImportPreview has all fields from import_preview.
type ImportPreview struct {
	Created int
	ID      int
	Name    string
	Result  json.RawMessage
	State   string
}

func (t *ImportPreview) lazy(ds *Fetch, id int) {
	ds.ImportPreview_Created(id).Lazy(&t.Created)
	ds.ImportPreview_ID(id).Lazy(&t.ID)
	ds.ImportPreview_Name(id).Lazy(&t.Name)
	ds.ImportPreview_Result(id).Lazy(&t.Result)
	ds.ImportPreview_State(id).Lazy(&t.State)
}

func (t *ImportPreview) preload(ds *Fetch, id int) {
	ds.ImportPreview_Created(id).Preload()
	ds.ImportPreview_ID(id).Preload()
	ds.ImportPreview_Name(id).Preload()
	ds.ImportPreview_Result(id).Preload()
	ds.ImportPreview_State(id).Preload()
}

func (r *Fetch) ImportPreview(id int) *ValueCollection[ImportPreview, *ImportPreview] {
	return &ValueCollection[ImportPreview, *ImportPreview]{
		id:    id,
		fetch: r,
	}
}

// MotionCategory has all fields from motion_category.
type MotionCategory struct {
	ChildIDs         []int
	Name             string
	ParentID         Maybe[int]
	ID               int
	Level            int
	MeetingID        int
	MotionIDs        []int
	Prefix           string
	SequentialNumber int
	Weight           int
}

func (t *MotionCategory) lazy(ds *Fetch, id int) {
	ds.MotionCategory_ChildIDs(id).Lazy(&t.ChildIDs)
	ds.MotionCategory_Name(id).Lazy(&t.Name)
	ds.MotionCategory_ParentID(id).Lazy(&t.ParentID)
	ds.MotionCategory_ID(id).Lazy(&t.ID)
	ds.MotionCategory_Level(id).Lazy(&t.Level)
	ds.MotionCategory_MeetingID(id).Lazy(&t.MeetingID)
	ds.MotionCategory_MotionIDs(id).Lazy(&t.MotionIDs)
	ds.MotionCategory_Prefix(id).Lazy(&t.Prefix)
	ds.MotionCategory_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.MotionCategory_Weight(id).Lazy(&t.Weight)
}

func (t *MotionCategory) preload(ds *Fetch, id int) {
	ds.MotionCategory_ChildIDs(id).Preload()
	ds.MotionCategory_Name(id).Preload()
	ds.MotionCategory_ParentID(id).Preload()
	ds.MotionCategory_ID(id).Preload()
	ds.MotionCategory_Level(id).Preload()
	ds.MotionCategory_MeetingID(id).Preload()
	ds.MotionCategory_MotionIDs(id).Preload()
	ds.MotionCategory_Prefix(id).Preload()
	ds.MotionCategory_SequentialNumber(id).Preload()
	ds.MotionCategory_Weight(id).Preload()
}

func (r *Fetch) MotionCategory(id int) *ValueCollection[MotionCategory, *MotionCategory] {
	return &ValueCollection[MotionCategory, *MotionCategory]{
		id:    id,
		fetch: r,
	}
}

// PollCandidateList has all fields from poll_candidate_list.
type PollCandidateList struct {
	PollCandidateIDs []int
	ID               int
	MeetingID        int
	OptionID         int
}

func (t *PollCandidateList) lazy(ds *Fetch, id int) {
	ds.PollCandidateList_PollCandidateIDs(id).Lazy(&t.PollCandidateIDs)
	ds.PollCandidateList_ID(id).Lazy(&t.ID)
	ds.PollCandidateList_MeetingID(id).Lazy(&t.MeetingID)
	ds.PollCandidateList_OptionID(id).Lazy(&t.OptionID)
}

func (t *PollCandidateList) preload(ds *Fetch, id int) {
	ds.PollCandidateList_PollCandidateIDs(id).Preload()
	ds.PollCandidateList_ID(id).Preload()
	ds.PollCandidateList_MeetingID(id).Preload()
	ds.PollCandidateList_OptionID(id).Preload()
}

func (r *Fetch) PollCandidateList(id int) *ValueCollection[PollCandidateList, *PollCandidateList] {
	return &ValueCollection[PollCandidateList, *PollCandidateList]{
		id:    id,
		fetch: r,
	}
}

// Projector has all fields from projector.
type Projector struct {
	UsedAsDefaultProjectorForMotionInMeetingID                Maybe[int]
	UsedAsReferenceProjectorMeetingID                         Maybe[int]
	BackgroundColor                                           string
	HeaderH1Color                                             string
	UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID Maybe[int]
	UsedAsDefaultProjectorForListOfSpeakersInMeetingID        Maybe[int]
	AspectRatioDenominator                                    int
	UsedAsDefaultProjectorForMessageInMeetingID               Maybe[int]
	ChyronBackgroundColor                                     string
	CurrentProjectionIDs                                      []int
	UsedAsDefaultProjectorForMotionPollInMeetingID            Maybe[int]
	UsedAsDefaultProjectorForPollInMeetingID                  Maybe[int]
	UsedAsDefaultProjectorForMediafileInMeetingID             Maybe[int]
	ChyronFontColor                                           string
	HeaderFontColor                                           string
	PreviewProjectionIDs                                      []int
	ShowHeaderFooter                                          bool
	Scale                                                     int
	ShowClock                                                 bool
	ShowLogo                                                  bool
	UsedAsDefaultProjectorForAgendaItemListInMeetingID        Maybe[int]
	ChyronFontColor2                                          string
	Color                                                     string
	HeaderBackgroundColor                                     string
	HistoryProjectionIDs                                      []int
	UsedAsDefaultProjectorForAssignmentPollInMeetingID        Maybe[int]
	UsedAsDefaultProjectorForMotionBlockInMeetingID           Maybe[int]
	Width                                                     int
	ChyronBackgroundColor2                                    string
	MeetingID                                                 int
	SequentialNumber                                          int
	ShowTitle                                                 bool
	UsedAsDefaultProjectorForAssignmentInMeetingID            Maybe[int]
	UsedAsDefaultProjectorForTopicInMeetingID                 Maybe[int]
	ID                                                        int
	IsInternal                                                bool
	Name                                                      string
	UsedAsDefaultProjectorForAmendmentInMeetingID             Maybe[int]
	AspectRatioNumerator                                      int
	Scroll                                                    int
	UsedAsDefaultProjectorForCountdownInMeetingID             Maybe[int]
}

func (t *Projector) lazy(ds *Fetch, id int) {
	ds.Projector_UsedAsDefaultProjectorForMotionInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForMotionInMeetingID)
	ds.Projector_UsedAsReferenceProjectorMeetingID(id).Lazy(&t.UsedAsReferenceProjectorMeetingID)
	ds.Projector_BackgroundColor(id).Lazy(&t.BackgroundColor)
	ds.Projector_HeaderH1Color(id).Lazy(&t.HeaderH1Color)
	ds.Projector_UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForListOfSpeakersInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForListOfSpeakersInMeetingID)
	ds.Projector_AspectRatioDenominator(id).Lazy(&t.AspectRatioDenominator)
	ds.Projector_UsedAsDefaultProjectorForMessageInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForMessageInMeetingID)
	ds.Projector_ChyronBackgroundColor(id).Lazy(&t.ChyronBackgroundColor)
	ds.Projector_CurrentProjectionIDs(id).Lazy(&t.CurrentProjectionIDs)
	ds.Projector_UsedAsDefaultProjectorForMotionPollInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForMotionPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForPollInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMediafileInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForMediafileInMeetingID)
	ds.Projector_ChyronFontColor(id).Lazy(&t.ChyronFontColor)
	ds.Projector_HeaderFontColor(id).Lazy(&t.HeaderFontColor)
	ds.Projector_PreviewProjectionIDs(id).Lazy(&t.PreviewProjectionIDs)
	ds.Projector_ShowHeaderFooter(id).Lazy(&t.ShowHeaderFooter)
	ds.Projector_Scale(id).Lazy(&t.Scale)
	ds.Projector_ShowClock(id).Lazy(&t.ShowClock)
	ds.Projector_ShowLogo(id).Lazy(&t.ShowLogo)
	ds.Projector_UsedAsDefaultProjectorForAgendaItemListInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForAgendaItemListInMeetingID)
	ds.Projector_ChyronFontColor2(id).Lazy(&t.ChyronFontColor2)
	ds.Projector_Color(id).Lazy(&t.Color)
	ds.Projector_HeaderBackgroundColor(id).Lazy(&t.HeaderBackgroundColor)
	ds.Projector_HistoryProjectionIDs(id).Lazy(&t.HistoryProjectionIDs)
	ds.Projector_UsedAsDefaultProjectorForAssignmentPollInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForAssignmentPollInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForMotionBlockInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForMotionBlockInMeetingID)
	ds.Projector_Width(id).Lazy(&t.Width)
	ds.Projector_ChyronBackgroundColor2(id).Lazy(&t.ChyronBackgroundColor2)
	ds.Projector_MeetingID(id).Lazy(&t.MeetingID)
	ds.Projector_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.Projector_ShowTitle(id).Lazy(&t.ShowTitle)
	ds.Projector_UsedAsDefaultProjectorForAssignmentInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForAssignmentInMeetingID)
	ds.Projector_UsedAsDefaultProjectorForTopicInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForTopicInMeetingID)
	ds.Projector_ID(id).Lazy(&t.ID)
	ds.Projector_IsInternal(id).Lazy(&t.IsInternal)
	ds.Projector_Name(id).Lazy(&t.Name)
	ds.Projector_UsedAsDefaultProjectorForAmendmentInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForAmendmentInMeetingID)
	ds.Projector_AspectRatioNumerator(id).Lazy(&t.AspectRatioNumerator)
	ds.Projector_Scroll(id).Lazy(&t.Scroll)
	ds.Projector_UsedAsDefaultProjectorForCountdownInMeetingID(id).Lazy(&t.UsedAsDefaultProjectorForCountdownInMeetingID)
}

func (t *Projector) preload(ds *Fetch, id int) {
	ds.Projector_UsedAsDefaultProjectorForMotionInMeetingID(id).Preload()
	ds.Projector_UsedAsReferenceProjectorMeetingID(id).Preload()
	ds.Projector_BackgroundColor(id).Preload()
	ds.Projector_HeaderH1Color(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForCurrentListOfSpeakersInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForListOfSpeakersInMeetingID(id).Preload()
	ds.Projector_AspectRatioDenominator(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMessageInMeetingID(id).Preload()
	ds.Projector_ChyronBackgroundColor(id).Preload()
	ds.Projector_CurrentProjectionIDs(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMotionPollInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForPollInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMediafileInMeetingID(id).Preload()
	ds.Projector_ChyronFontColor(id).Preload()
	ds.Projector_HeaderFontColor(id).Preload()
	ds.Projector_PreviewProjectionIDs(id).Preload()
	ds.Projector_ShowHeaderFooter(id).Preload()
	ds.Projector_Scale(id).Preload()
	ds.Projector_ShowClock(id).Preload()
	ds.Projector_ShowLogo(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForAgendaItemListInMeetingID(id).Preload()
	ds.Projector_ChyronFontColor2(id).Preload()
	ds.Projector_Color(id).Preload()
	ds.Projector_HeaderBackgroundColor(id).Preload()
	ds.Projector_HistoryProjectionIDs(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForAssignmentPollInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForMotionBlockInMeetingID(id).Preload()
	ds.Projector_Width(id).Preload()
	ds.Projector_ChyronBackgroundColor2(id).Preload()
	ds.Projector_MeetingID(id).Preload()
	ds.Projector_SequentialNumber(id).Preload()
	ds.Projector_ShowTitle(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForAssignmentInMeetingID(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForTopicInMeetingID(id).Preload()
	ds.Projector_ID(id).Preload()
	ds.Projector_IsInternal(id).Preload()
	ds.Projector_Name(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForAmendmentInMeetingID(id).Preload()
	ds.Projector_AspectRatioNumerator(id).Preload()
	ds.Projector_Scroll(id).Preload()
	ds.Projector_UsedAsDefaultProjectorForCountdownInMeetingID(id).Preload()
}

func (r *Fetch) Projector(id int) *ValueCollection[Projector, *Projector] {
	return &ValueCollection[Projector, *Projector]{
		id:    id,
		fetch: r,
	}
}

// Group has all fields from group.
type Group struct {
	ExternalID                              string
	MeetingMediafileAccessGroupIDs          []int
	Name                                    string
	Permissions                             []string
	ReadCommentSectionIDs                   []int
	UsedAsPollDefaultID                     Maybe[int]
	WriteCommentSectionIDs                  []int
	DefaultGroupForMeetingID                Maybe[int]
	ID                                      int
	MeetingID                               int
	MeetingUserIDs                          []int
	ReadChatGroupIDs                        []int
	UsedAsMotionPollDefaultID               Maybe[int]
	AdminGroupForMeetingID                  Maybe[int]
	PollIDs                                 []int
	UsedAsAssignmentPollDefaultID           Maybe[int]
	WriteChatGroupIDs                       []int
	AnonymousGroupForMeetingID              Maybe[int]
	MeetingMediafileInheritedAccessGroupIDs []int
	UsedAsTopicPollDefaultID                Maybe[int]
	Weight                                  int
}

func (t *Group) lazy(ds *Fetch, id int) {
	ds.Group_ExternalID(id).Lazy(&t.ExternalID)
	ds.Group_MeetingMediafileAccessGroupIDs(id).Lazy(&t.MeetingMediafileAccessGroupIDs)
	ds.Group_Name(id).Lazy(&t.Name)
	ds.Group_Permissions(id).Lazy(&t.Permissions)
	ds.Group_ReadCommentSectionIDs(id).Lazy(&t.ReadCommentSectionIDs)
	ds.Group_UsedAsPollDefaultID(id).Lazy(&t.UsedAsPollDefaultID)
	ds.Group_WriteCommentSectionIDs(id).Lazy(&t.WriteCommentSectionIDs)
	ds.Group_DefaultGroupForMeetingID(id).Lazy(&t.DefaultGroupForMeetingID)
	ds.Group_ID(id).Lazy(&t.ID)
	ds.Group_MeetingID(id).Lazy(&t.MeetingID)
	ds.Group_MeetingUserIDs(id).Lazy(&t.MeetingUserIDs)
	ds.Group_ReadChatGroupIDs(id).Lazy(&t.ReadChatGroupIDs)
	ds.Group_UsedAsMotionPollDefaultID(id).Lazy(&t.UsedAsMotionPollDefaultID)
	ds.Group_AdminGroupForMeetingID(id).Lazy(&t.AdminGroupForMeetingID)
	ds.Group_PollIDs(id).Lazy(&t.PollIDs)
	ds.Group_UsedAsAssignmentPollDefaultID(id).Lazy(&t.UsedAsAssignmentPollDefaultID)
	ds.Group_WriteChatGroupIDs(id).Lazy(&t.WriteChatGroupIDs)
	ds.Group_AnonymousGroupForMeetingID(id).Lazy(&t.AnonymousGroupForMeetingID)
	ds.Group_MeetingMediafileInheritedAccessGroupIDs(id).Lazy(&t.MeetingMediafileInheritedAccessGroupIDs)
	ds.Group_UsedAsTopicPollDefaultID(id).Lazy(&t.UsedAsTopicPollDefaultID)
	ds.Group_Weight(id).Lazy(&t.Weight)
}

func (t *Group) preload(ds *Fetch, id int) {
	ds.Group_ExternalID(id).Preload()
	ds.Group_MeetingMediafileAccessGroupIDs(id).Preload()
	ds.Group_Name(id).Preload()
	ds.Group_Permissions(id).Preload()
	ds.Group_ReadCommentSectionIDs(id).Preload()
	ds.Group_UsedAsPollDefaultID(id).Preload()
	ds.Group_WriteCommentSectionIDs(id).Preload()
	ds.Group_DefaultGroupForMeetingID(id).Preload()
	ds.Group_ID(id).Preload()
	ds.Group_MeetingID(id).Preload()
	ds.Group_MeetingUserIDs(id).Preload()
	ds.Group_ReadChatGroupIDs(id).Preload()
	ds.Group_UsedAsMotionPollDefaultID(id).Preload()
	ds.Group_AdminGroupForMeetingID(id).Preload()
	ds.Group_PollIDs(id).Preload()
	ds.Group_UsedAsAssignmentPollDefaultID(id).Preload()
	ds.Group_WriteChatGroupIDs(id).Preload()
	ds.Group_AnonymousGroupForMeetingID(id).Preload()
	ds.Group_MeetingMediafileInheritedAccessGroupIDs(id).Preload()
	ds.Group_UsedAsTopicPollDefaultID(id).Preload()
	ds.Group_Weight(id).Preload()
}

func (r *Fetch) Group(id int) *ValueCollection[Group, *Group] {
	return &ValueCollection[Group, *Group]{
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
}

func (t *Gender) lazy(ds *Fetch, id int) {
	ds.Gender_ID(id).Lazy(&t.ID)
	ds.Gender_Name(id).Lazy(&t.Name)
	ds.Gender_OrganizationID(id).Lazy(&t.OrganizationID)
	ds.Gender_UserIDs(id).Lazy(&t.UserIDs)
}

func (t *Gender) preload(ds *Fetch, id int) {
	ds.Gender_ID(id).Preload()
	ds.Gender_Name(id).Preload()
	ds.Gender_OrganizationID(id).Preload()
	ds.Gender_UserIDs(id).Preload()
}

func (r *Fetch) Gender(id int) *ValueCollection[Gender, *Gender] {
	return &ValueCollection[Gender, *Gender]{
		id:    id,
		fetch: r,
	}
}

// MotionCommentSection has all fields from motion_comment_section.
type MotionCommentSection struct {
	MeetingID         int
	Name              string
	ReadGroupIDs      []int
	SequentialNumber  int
	Weight            int
	WriteGroupIDs     []int
	CommentIDs        []int
	ID                int
	SubmitterCanWrite bool
}

func (t *MotionCommentSection) lazy(ds *Fetch, id int) {
	ds.MotionCommentSection_MeetingID(id).Lazy(&t.MeetingID)
	ds.MotionCommentSection_Name(id).Lazy(&t.Name)
	ds.MotionCommentSection_ReadGroupIDs(id).Lazy(&t.ReadGroupIDs)
	ds.MotionCommentSection_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.MotionCommentSection_Weight(id).Lazy(&t.Weight)
	ds.MotionCommentSection_WriteGroupIDs(id).Lazy(&t.WriteGroupIDs)
	ds.MotionCommentSection_CommentIDs(id).Lazy(&t.CommentIDs)
	ds.MotionCommentSection_ID(id).Lazy(&t.ID)
	ds.MotionCommentSection_SubmitterCanWrite(id).Lazy(&t.SubmitterCanWrite)
}

func (t *MotionCommentSection) preload(ds *Fetch, id int) {
	ds.MotionCommentSection_MeetingID(id).Preload()
	ds.MotionCommentSection_Name(id).Preload()
	ds.MotionCommentSection_ReadGroupIDs(id).Preload()
	ds.MotionCommentSection_SequentialNumber(id).Preload()
	ds.MotionCommentSection_Weight(id).Preload()
	ds.MotionCommentSection_WriteGroupIDs(id).Preload()
	ds.MotionCommentSection_CommentIDs(id).Preload()
	ds.MotionCommentSection_ID(id).Preload()
	ds.MotionCommentSection_SubmitterCanWrite(id).Preload()
}

func (r *Fetch) MotionCommentSection(id int) *ValueCollection[MotionCommentSection, *MotionCommentSection] {
	return &ValueCollection[MotionCommentSection, *MotionCommentSection]{
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
}

func (t *MotionSubmitter) lazy(ds *Fetch, id int) {
	ds.MotionSubmitter_ID(id).Lazy(&t.ID)
	ds.MotionSubmitter_MeetingID(id).Lazy(&t.MeetingID)
	ds.MotionSubmitter_MeetingUserID(id).Lazy(&t.MeetingUserID)
	ds.MotionSubmitter_MotionID(id).Lazy(&t.MotionID)
	ds.MotionSubmitter_Weight(id).Lazy(&t.Weight)
}

func (t *MotionSubmitter) preload(ds *Fetch, id int) {
	ds.MotionSubmitter_ID(id).Preload()
	ds.MotionSubmitter_MeetingID(id).Preload()
	ds.MotionSubmitter_MeetingUserID(id).Preload()
	ds.MotionSubmitter_MotionID(id).Preload()
	ds.MotionSubmitter_Weight(id).Preload()
}

func (r *Fetch) MotionSubmitter(id int) *ValueCollection[MotionSubmitter, *MotionSubmitter] {
	return &ValueCollection[MotionSubmitter, *MotionSubmitter]{
		id:    id,
		fetch: r,
	}
}

// Organization has all fields from organization.
type Organization struct {
	EnableAnonymous            bool
	GenderIDs                  []int
	SamlMetadataSp             string
	PrivacyPolicy              string
	RequireDuplicateFrom       bool
	ResetPasswordVerboseErrors bool
	SamlAttrMapping            json.RawMessage
	SamlLoginButtonText        string
	UsersEmailSender           string
	EnableElectronicVoting     bool
	LoginText                  string
	MediafileIDs               []int
	OrganizationTagIDs         []int
	Url                        string
	ArchivedMeetingIDs         []int
	DefaultLanguage            string
	EnableChat                 bool
	SamlPrivateKey             string
	Description                string
	LegalNotice                string
	LimitOfMeetings            int
	ActiveMeetingIDs           []int
	CommitteeIDs               []int
	ThemeID                    int
	UserIDs                    []int
	UsersEmailBody             string
	UsersEmailReplyto          string
	VoteDecryptPublicMainKey   string
	ID                         int
	LimitOfUsers               int
	TemplateMeetingIDs         []int
	UsersEmailSubject          string
	Name                       string
	PublishedMediafileIDs      []int
	SamlEnabled                bool
	SamlMetadataIDp            string
	ThemeIDs                   []int
}

func (t *Organization) lazy(ds *Fetch, id int) {
	ds.Organization_EnableAnonymous(id).Lazy(&t.EnableAnonymous)
	ds.Organization_GenderIDs(id).Lazy(&t.GenderIDs)
	ds.Organization_SamlMetadataSp(id).Lazy(&t.SamlMetadataSp)
	ds.Organization_PrivacyPolicy(id).Lazy(&t.PrivacyPolicy)
	ds.Organization_RequireDuplicateFrom(id).Lazy(&t.RequireDuplicateFrom)
	ds.Organization_ResetPasswordVerboseErrors(id).Lazy(&t.ResetPasswordVerboseErrors)
	ds.Organization_SamlAttrMapping(id).Lazy(&t.SamlAttrMapping)
	ds.Organization_SamlLoginButtonText(id).Lazy(&t.SamlLoginButtonText)
	ds.Organization_UsersEmailSender(id).Lazy(&t.UsersEmailSender)
	ds.Organization_EnableElectronicVoting(id).Lazy(&t.EnableElectronicVoting)
	ds.Organization_LoginText(id).Lazy(&t.LoginText)
	ds.Organization_MediafileIDs(id).Lazy(&t.MediafileIDs)
	ds.Organization_OrganizationTagIDs(id).Lazy(&t.OrganizationTagIDs)
	ds.Organization_Url(id).Lazy(&t.Url)
	ds.Organization_ArchivedMeetingIDs(id).Lazy(&t.ArchivedMeetingIDs)
	ds.Organization_DefaultLanguage(id).Lazy(&t.DefaultLanguage)
	ds.Organization_EnableChat(id).Lazy(&t.EnableChat)
	ds.Organization_SamlPrivateKey(id).Lazy(&t.SamlPrivateKey)
	ds.Organization_Description(id).Lazy(&t.Description)
	ds.Organization_LegalNotice(id).Lazy(&t.LegalNotice)
	ds.Organization_LimitOfMeetings(id).Lazy(&t.LimitOfMeetings)
	ds.Organization_ActiveMeetingIDs(id).Lazy(&t.ActiveMeetingIDs)
	ds.Organization_CommitteeIDs(id).Lazy(&t.CommitteeIDs)
	ds.Organization_ThemeID(id).Lazy(&t.ThemeID)
	ds.Organization_UserIDs(id).Lazy(&t.UserIDs)
	ds.Organization_UsersEmailBody(id).Lazy(&t.UsersEmailBody)
	ds.Organization_UsersEmailReplyto(id).Lazy(&t.UsersEmailReplyto)
	ds.Organization_VoteDecryptPublicMainKey(id).Lazy(&t.VoteDecryptPublicMainKey)
	ds.Organization_ID(id).Lazy(&t.ID)
	ds.Organization_LimitOfUsers(id).Lazy(&t.LimitOfUsers)
	ds.Organization_TemplateMeetingIDs(id).Lazy(&t.TemplateMeetingIDs)
	ds.Organization_UsersEmailSubject(id).Lazy(&t.UsersEmailSubject)
	ds.Organization_Name(id).Lazy(&t.Name)
	ds.Organization_PublishedMediafileIDs(id).Lazy(&t.PublishedMediafileIDs)
	ds.Organization_SamlEnabled(id).Lazy(&t.SamlEnabled)
	ds.Organization_SamlMetadataIDp(id).Lazy(&t.SamlMetadataIDp)
	ds.Organization_ThemeIDs(id).Lazy(&t.ThemeIDs)
}

func (t *Organization) preload(ds *Fetch, id int) {
	ds.Organization_EnableAnonymous(id).Preload()
	ds.Organization_GenderIDs(id).Preload()
	ds.Organization_SamlMetadataSp(id).Preload()
	ds.Organization_PrivacyPolicy(id).Preload()
	ds.Organization_RequireDuplicateFrom(id).Preload()
	ds.Organization_ResetPasswordVerboseErrors(id).Preload()
	ds.Organization_SamlAttrMapping(id).Preload()
	ds.Organization_SamlLoginButtonText(id).Preload()
	ds.Organization_UsersEmailSender(id).Preload()
	ds.Organization_EnableElectronicVoting(id).Preload()
	ds.Organization_LoginText(id).Preload()
	ds.Organization_MediafileIDs(id).Preload()
	ds.Organization_OrganizationTagIDs(id).Preload()
	ds.Organization_Url(id).Preload()
	ds.Organization_ArchivedMeetingIDs(id).Preload()
	ds.Organization_DefaultLanguage(id).Preload()
	ds.Organization_EnableChat(id).Preload()
	ds.Organization_SamlPrivateKey(id).Preload()
	ds.Organization_Description(id).Preload()
	ds.Organization_LegalNotice(id).Preload()
	ds.Organization_LimitOfMeetings(id).Preload()
	ds.Organization_ActiveMeetingIDs(id).Preload()
	ds.Organization_CommitteeIDs(id).Preload()
	ds.Organization_ThemeID(id).Preload()
	ds.Organization_UserIDs(id).Preload()
	ds.Organization_UsersEmailBody(id).Preload()
	ds.Organization_UsersEmailReplyto(id).Preload()
	ds.Organization_VoteDecryptPublicMainKey(id).Preload()
	ds.Organization_ID(id).Preload()
	ds.Organization_LimitOfUsers(id).Preload()
	ds.Organization_TemplateMeetingIDs(id).Preload()
	ds.Organization_UsersEmailSubject(id).Preload()
	ds.Organization_Name(id).Preload()
	ds.Organization_PublishedMediafileIDs(id).Preload()
	ds.Organization_SamlEnabled(id).Preload()
	ds.Organization_SamlMetadataIDp(id).Preload()
	ds.Organization_ThemeIDs(id).Preload()
}

func (r *Fetch) Organization(id int) *ValueCollection[Organization, *Organization] {
	return &ValueCollection[Organization, *Organization]{
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
}

func (t *PersonalNote) lazy(ds *Fetch, id int) {
	ds.PersonalNote_ContentObjectID(id).Lazy(&t.ContentObjectID)
	ds.PersonalNote_ID(id).Lazy(&t.ID)
	ds.PersonalNote_MeetingID(id).Lazy(&t.MeetingID)
	ds.PersonalNote_MeetingUserID(id).Lazy(&t.MeetingUserID)
	ds.PersonalNote_Note(id).Lazy(&t.Note)
	ds.PersonalNote_Star(id).Lazy(&t.Star)
}

func (t *PersonalNote) preload(ds *Fetch, id int) {
	ds.PersonalNote_ContentObjectID(id).Preload()
	ds.PersonalNote_ID(id).Preload()
	ds.PersonalNote_MeetingID(id).Preload()
	ds.PersonalNote_MeetingUserID(id).Preload()
	ds.PersonalNote_Note(id).Preload()
	ds.PersonalNote_Star(id).Preload()
}

func (r *Fetch) PersonalNote(id int) *ValueCollection[PersonalNote, *PersonalNote] {
	return &ValueCollection[PersonalNote, *PersonalNote]{
		id:    id,
		fetch: r,
	}
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
}

func (t *ActionWorker) lazy(ds *Fetch, id int) {
	ds.ActionWorker_Created(id).Lazy(&t.Created)
	ds.ActionWorker_ID(id).Lazy(&t.ID)
	ds.ActionWorker_Name(id).Lazy(&t.Name)
	ds.ActionWorker_Result(id).Lazy(&t.Result)
	ds.ActionWorker_State(id).Lazy(&t.State)
	ds.ActionWorker_Timestamp(id).Lazy(&t.Timestamp)
	ds.ActionWorker_UserID(id).Lazy(&t.UserID)
}

func (t *ActionWorker) preload(ds *Fetch, id int) {
	ds.ActionWorker_Created(id).Preload()
	ds.ActionWorker_ID(id).Preload()
	ds.ActionWorker_Name(id).Preload()
	ds.ActionWorker_Result(id).Preload()
	ds.ActionWorker_State(id).Preload()
	ds.ActionWorker_Timestamp(id).Preload()
	ds.ActionWorker_UserID(id).Preload()
}

func (r *Fetch) ActionWorker(id int) *ValueCollection[ActionWorker, *ActionWorker] {
	return &ValueCollection[ActionWorker, *ActionWorker]{
		id:    id,
		fetch: r,
	}
}

// OrganizationTag has all fields from organization_tag.
type OrganizationTag struct {
	OrganizationID int
	TaggedIDs      []string
	Color          string
	ID             int
	Name           string
}

func (t *OrganizationTag) lazy(ds *Fetch, id int) {
	ds.OrganizationTag_OrganizationID(id).Lazy(&t.OrganizationID)
	ds.OrganizationTag_TaggedIDs(id).Lazy(&t.TaggedIDs)
	ds.OrganizationTag_Color(id).Lazy(&t.Color)
	ds.OrganizationTag_ID(id).Lazy(&t.ID)
	ds.OrganizationTag_Name(id).Lazy(&t.Name)
}

func (t *OrganizationTag) preload(ds *Fetch, id int) {
	ds.OrganizationTag_OrganizationID(id).Preload()
	ds.OrganizationTag_TaggedIDs(id).Preload()
	ds.OrganizationTag_Color(id).Preload()
	ds.OrganizationTag_ID(id).Preload()
	ds.OrganizationTag_Name(id).Preload()
}

func (r *Fetch) OrganizationTag(id int) *ValueCollection[OrganizationTag, *OrganizationTag] {
	return &ValueCollection[OrganizationTag, *OrganizationTag]{
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
}

func (t *ProjectorMessage) lazy(ds *Fetch, id int) {
	ds.ProjectorMessage_ID(id).Lazy(&t.ID)
	ds.ProjectorMessage_MeetingID(id).Lazy(&t.MeetingID)
	ds.ProjectorMessage_Message(id).Lazy(&t.Message)
	ds.ProjectorMessage_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
}

func (t *ProjectorMessage) preload(ds *Fetch, id int) {
	ds.ProjectorMessage_ID(id).Preload()
	ds.ProjectorMessage_MeetingID(id).Preload()
	ds.ProjectorMessage_Message(id).Preload()
	ds.ProjectorMessage_ProjectionIDs(id).Preload()
}

func (r *Fetch) ProjectorMessage(id int) *ValueCollection[ProjectorMessage, *ProjectorMessage] {
	return &ValueCollection[ProjectorMessage, *ProjectorMessage]{
		id:    id,
		fetch: r,
	}
}

// StructureLevelListOfSpeakers has all fields from structure_level_list_of_speakers.
type StructureLevelListOfSpeakers struct {
	ID               int
	RemainingTime    float32
	SpeakerIDs       []int
	ListOfSpeakersID int
	MeetingID        int
	StructureLevelID int
	AdditionalTime   float32
	CurrentStartTime int
	InitialTime      int
}

func (t *StructureLevelListOfSpeakers) lazy(ds *Fetch, id int) {
	ds.StructureLevelListOfSpeakers_ID(id).Lazy(&t.ID)
	ds.StructureLevelListOfSpeakers_RemainingTime(id).Lazy(&t.RemainingTime)
	ds.StructureLevelListOfSpeakers_SpeakerIDs(id).Lazy(&t.SpeakerIDs)
	ds.StructureLevelListOfSpeakers_ListOfSpeakersID(id).Lazy(&t.ListOfSpeakersID)
	ds.StructureLevelListOfSpeakers_MeetingID(id).Lazy(&t.MeetingID)
	ds.StructureLevelListOfSpeakers_StructureLevelID(id).Lazy(&t.StructureLevelID)
	ds.StructureLevelListOfSpeakers_AdditionalTime(id).Lazy(&t.AdditionalTime)
	ds.StructureLevelListOfSpeakers_CurrentStartTime(id).Lazy(&t.CurrentStartTime)
	ds.StructureLevelListOfSpeakers_InitialTime(id).Lazy(&t.InitialTime)
}

func (t *StructureLevelListOfSpeakers) preload(ds *Fetch, id int) {
	ds.StructureLevelListOfSpeakers_ID(id).Preload()
	ds.StructureLevelListOfSpeakers_RemainingTime(id).Preload()
	ds.StructureLevelListOfSpeakers_SpeakerIDs(id).Preload()
	ds.StructureLevelListOfSpeakers_ListOfSpeakersID(id).Preload()
	ds.StructureLevelListOfSpeakers_MeetingID(id).Preload()
	ds.StructureLevelListOfSpeakers_StructureLevelID(id).Preload()
	ds.StructureLevelListOfSpeakers_AdditionalTime(id).Preload()
	ds.StructureLevelListOfSpeakers_CurrentStartTime(id).Preload()
	ds.StructureLevelListOfSpeakers_InitialTime(id).Preload()
}

func (r *Fetch) StructureLevelListOfSpeakers(id int) *ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers] {
	return &ValueCollection[StructureLevelListOfSpeakers, *StructureLevelListOfSpeakers]{
		id:    id,
		fetch: r,
	}
}

// Topic has all fields from topic.
type Topic struct {
	AttachmentMeetingMediafileIDs []int
	ListOfSpeakersID              int
	PollIDs                       []int
	Title                         string
	AgendaItemID                  int
	ID                            int
	MeetingID                     int
	ProjectionIDs                 []int
	SequentialNumber              int
	Text                          string
}

func (t *Topic) lazy(ds *Fetch, id int) {
	ds.Topic_AttachmentMeetingMediafileIDs(id).Lazy(&t.AttachmentMeetingMediafileIDs)
	ds.Topic_ListOfSpeakersID(id).Lazy(&t.ListOfSpeakersID)
	ds.Topic_PollIDs(id).Lazy(&t.PollIDs)
	ds.Topic_Title(id).Lazy(&t.Title)
	ds.Topic_AgendaItemID(id).Lazy(&t.AgendaItemID)
	ds.Topic_ID(id).Lazy(&t.ID)
	ds.Topic_MeetingID(id).Lazy(&t.MeetingID)
	ds.Topic_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.Topic_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.Topic_Text(id).Lazy(&t.Text)
}

func (t *Topic) preload(ds *Fetch, id int) {
	ds.Topic_AttachmentMeetingMediafileIDs(id).Preload()
	ds.Topic_ListOfSpeakersID(id).Preload()
	ds.Topic_PollIDs(id).Preload()
	ds.Topic_Title(id).Preload()
	ds.Topic_AgendaItemID(id).Preload()
	ds.Topic_ID(id).Preload()
	ds.Topic_MeetingID(id).Preload()
	ds.Topic_ProjectionIDs(id).Preload()
	ds.Topic_SequentialNumber(id).Preload()
	ds.Topic_Text(id).Preload()
}

func (r *Fetch) Topic(id int) *ValueCollection[Topic, *Topic] {
	return &ValueCollection[Topic, *Topic]{
		id:    id,
		fetch: r,
	}
}

// Vote has all fields from vote.
type Vote struct {
	MeetingID       int
	OptionID        int
	UserID          Maybe[int]
	UserToken       string
	Value           string
	Weight          string
	DelegatedUserID Maybe[int]
	ID              int
}

func (t *Vote) lazy(ds *Fetch, id int) {
	ds.Vote_MeetingID(id).Lazy(&t.MeetingID)
	ds.Vote_OptionID(id).Lazy(&t.OptionID)
	ds.Vote_UserID(id).Lazy(&t.UserID)
	ds.Vote_UserToken(id).Lazy(&t.UserToken)
	ds.Vote_Value(id).Lazy(&t.Value)
	ds.Vote_Weight(id).Lazy(&t.Weight)
	ds.Vote_DelegatedUserID(id).Lazy(&t.DelegatedUserID)
	ds.Vote_ID(id).Lazy(&t.ID)
}

func (t *Vote) preload(ds *Fetch, id int) {
	ds.Vote_MeetingID(id).Preload()
	ds.Vote_OptionID(id).Preload()
	ds.Vote_UserID(id).Preload()
	ds.Vote_UserToken(id).Preload()
	ds.Vote_Value(id).Preload()
	ds.Vote_Weight(id).Preload()
	ds.Vote_DelegatedUserID(id).Preload()
	ds.Vote_ID(id).Preload()
}

func (r *Fetch) Vote(id int) *ValueCollection[Vote, *Vote] {
	return &ValueCollection[Vote, *Vote]{
		id:    id,
		fetch: r,
	}
}

// Motion has all fields from motion.
type Motion struct {
	ChangeRecommendationIDs                      []int
	RecommendationID                             Maybe[int]
	Text                                         string
	BlockID                                      Maybe[int]
	MeetingID                                    int
	AgendaItemID                                 Maybe[int]
	StateExtension                               string
	SubmitterIDs                                 []int
	TextHash                                     string
	Title                                        string
	WorkflowTimestamp                            int
	SortWeight                                   int
	SequentialNumber                             int
	ProjectionIDs                                []int
	TagIDs                                       []int
	WorkingGroupSpeakerIDs                       []int
	Forwarded                                    int
	ID                                           int
	LeadMotionID                                 Maybe[int]
	Number                                       string
	OriginMeetingID                              Maybe[int]
	CategoryID                                   Maybe[int]
	OriginID                                     Maybe[int]
	Reason                                       string
	ReferencedInMotionStateExtensionIDs          []int
	StateExtensionReferenceIDs                   []string
	AllDerivedMotionIDs                          []int
	AmendmentIDs                                 []int
	Created                                      int
	ModifiedFinalVersion                         string
	OptionIDs                                    []int
	RecommendationExtensionReferenceIDs          []string
	AllOriginIDs                                 []int
	CategoryWeight                               int
	DerivedMotionIDs                             []int
	ListOfSpeakersID                             int
	ReferencedInMotionRecommendationExtensionIDs []int
	SortChildIDs                                 []int
	AdditionalSubmitter                          string
	SortParentID                                 Maybe[int]
	PollIDs                                      []int
	NumberValue                                  int
	StartLineNumber                              int
	CommentIDs                                   []int
	EditorIDs                                    []int
	LastModified                                 int
	RecommendationExtension                      string
	StateID                                      int
	AttachmentMeetingMediafileIDs                []int
	IDenticalMotionIDs                           []int
	PersonalNoteIDs                              []int
	SupporterMeetingUserIDs                      []int
	AmendmentParagraphs                          json.RawMessage
}

func (t *Motion) lazy(ds *Fetch, id int) {
	ds.Motion_ChangeRecommendationIDs(id).Lazy(&t.ChangeRecommendationIDs)
	ds.Motion_RecommendationID(id).Lazy(&t.RecommendationID)
	ds.Motion_Text(id).Lazy(&t.Text)
	ds.Motion_BlockID(id).Lazy(&t.BlockID)
	ds.Motion_MeetingID(id).Lazy(&t.MeetingID)
	ds.Motion_AgendaItemID(id).Lazy(&t.AgendaItemID)
	ds.Motion_StateExtension(id).Lazy(&t.StateExtension)
	ds.Motion_SubmitterIDs(id).Lazy(&t.SubmitterIDs)
	ds.Motion_TextHash(id).Lazy(&t.TextHash)
	ds.Motion_Title(id).Lazy(&t.Title)
	ds.Motion_WorkflowTimestamp(id).Lazy(&t.WorkflowTimestamp)
	ds.Motion_SortWeight(id).Lazy(&t.SortWeight)
	ds.Motion_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.Motion_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.Motion_TagIDs(id).Lazy(&t.TagIDs)
	ds.Motion_WorkingGroupSpeakerIDs(id).Lazy(&t.WorkingGroupSpeakerIDs)
	ds.Motion_Forwarded(id).Lazy(&t.Forwarded)
	ds.Motion_ID(id).Lazy(&t.ID)
	ds.Motion_LeadMotionID(id).Lazy(&t.LeadMotionID)
	ds.Motion_Number(id).Lazy(&t.Number)
	ds.Motion_OriginMeetingID(id).Lazy(&t.OriginMeetingID)
	ds.Motion_CategoryID(id).Lazy(&t.CategoryID)
	ds.Motion_OriginID(id).Lazy(&t.OriginID)
	ds.Motion_Reason(id).Lazy(&t.Reason)
	ds.Motion_ReferencedInMotionStateExtensionIDs(id).Lazy(&t.ReferencedInMotionStateExtensionIDs)
	ds.Motion_StateExtensionReferenceIDs(id).Lazy(&t.StateExtensionReferenceIDs)
	ds.Motion_AllDerivedMotionIDs(id).Lazy(&t.AllDerivedMotionIDs)
	ds.Motion_AmendmentIDs(id).Lazy(&t.AmendmentIDs)
	ds.Motion_Created(id).Lazy(&t.Created)
	ds.Motion_ModifiedFinalVersion(id).Lazy(&t.ModifiedFinalVersion)
	ds.Motion_OptionIDs(id).Lazy(&t.OptionIDs)
	ds.Motion_RecommendationExtensionReferenceIDs(id).Lazy(&t.RecommendationExtensionReferenceIDs)
	ds.Motion_AllOriginIDs(id).Lazy(&t.AllOriginIDs)
	ds.Motion_CategoryWeight(id).Lazy(&t.CategoryWeight)
	ds.Motion_DerivedMotionIDs(id).Lazy(&t.DerivedMotionIDs)
	ds.Motion_ListOfSpeakersID(id).Lazy(&t.ListOfSpeakersID)
	ds.Motion_ReferencedInMotionRecommendationExtensionIDs(id).Lazy(&t.ReferencedInMotionRecommendationExtensionIDs)
	ds.Motion_SortChildIDs(id).Lazy(&t.SortChildIDs)
	ds.Motion_AdditionalSubmitter(id).Lazy(&t.AdditionalSubmitter)
	ds.Motion_SortParentID(id).Lazy(&t.SortParentID)
	ds.Motion_PollIDs(id).Lazy(&t.PollIDs)
	ds.Motion_NumberValue(id).Lazy(&t.NumberValue)
	ds.Motion_StartLineNumber(id).Lazy(&t.StartLineNumber)
	ds.Motion_CommentIDs(id).Lazy(&t.CommentIDs)
	ds.Motion_EditorIDs(id).Lazy(&t.EditorIDs)
	ds.Motion_LastModified(id).Lazy(&t.LastModified)
	ds.Motion_RecommendationExtension(id).Lazy(&t.RecommendationExtension)
	ds.Motion_StateID(id).Lazy(&t.StateID)
	ds.Motion_AttachmentMeetingMediafileIDs(id).Lazy(&t.AttachmentMeetingMediafileIDs)
	ds.Motion_IDenticalMotionIDs(id).Lazy(&t.IDenticalMotionIDs)
	ds.Motion_PersonalNoteIDs(id).Lazy(&t.PersonalNoteIDs)
	ds.Motion_SupporterMeetingUserIDs(id).Lazy(&t.SupporterMeetingUserIDs)
	ds.Motion_AmendmentParagraphs(id).Lazy(&t.AmendmentParagraphs)
}

func (t *Motion) preload(ds *Fetch, id int) {
	ds.Motion_ChangeRecommendationIDs(id).Preload()
	ds.Motion_RecommendationID(id).Preload()
	ds.Motion_Text(id).Preload()
	ds.Motion_BlockID(id).Preload()
	ds.Motion_MeetingID(id).Preload()
	ds.Motion_AgendaItemID(id).Preload()
	ds.Motion_StateExtension(id).Preload()
	ds.Motion_SubmitterIDs(id).Preload()
	ds.Motion_TextHash(id).Preload()
	ds.Motion_Title(id).Preload()
	ds.Motion_WorkflowTimestamp(id).Preload()
	ds.Motion_SortWeight(id).Preload()
	ds.Motion_SequentialNumber(id).Preload()
	ds.Motion_ProjectionIDs(id).Preload()
	ds.Motion_TagIDs(id).Preload()
	ds.Motion_WorkingGroupSpeakerIDs(id).Preload()
	ds.Motion_Forwarded(id).Preload()
	ds.Motion_ID(id).Preload()
	ds.Motion_LeadMotionID(id).Preload()
	ds.Motion_Number(id).Preload()
	ds.Motion_OriginMeetingID(id).Preload()
	ds.Motion_CategoryID(id).Preload()
	ds.Motion_OriginID(id).Preload()
	ds.Motion_Reason(id).Preload()
	ds.Motion_ReferencedInMotionStateExtensionIDs(id).Preload()
	ds.Motion_StateExtensionReferenceIDs(id).Preload()
	ds.Motion_AllDerivedMotionIDs(id).Preload()
	ds.Motion_AmendmentIDs(id).Preload()
	ds.Motion_Created(id).Preload()
	ds.Motion_ModifiedFinalVersion(id).Preload()
	ds.Motion_OptionIDs(id).Preload()
	ds.Motion_RecommendationExtensionReferenceIDs(id).Preload()
	ds.Motion_AllOriginIDs(id).Preload()
	ds.Motion_CategoryWeight(id).Preload()
	ds.Motion_DerivedMotionIDs(id).Preload()
	ds.Motion_ListOfSpeakersID(id).Preload()
	ds.Motion_ReferencedInMotionRecommendationExtensionIDs(id).Preload()
	ds.Motion_SortChildIDs(id).Preload()
	ds.Motion_AdditionalSubmitter(id).Preload()
	ds.Motion_SortParentID(id).Preload()
	ds.Motion_PollIDs(id).Preload()
	ds.Motion_NumberValue(id).Preload()
	ds.Motion_StartLineNumber(id).Preload()
	ds.Motion_CommentIDs(id).Preload()
	ds.Motion_EditorIDs(id).Preload()
	ds.Motion_LastModified(id).Preload()
	ds.Motion_RecommendationExtension(id).Preload()
	ds.Motion_StateID(id).Preload()
	ds.Motion_AttachmentMeetingMediafileIDs(id).Preload()
	ds.Motion_IDenticalMotionIDs(id).Preload()
	ds.Motion_PersonalNoteIDs(id).Preload()
	ds.Motion_SupporterMeetingUserIDs(id).Preload()
	ds.Motion_AmendmentParagraphs(id).Preload()
}

func (r *Fetch) Motion(id int) *ValueCollection[Motion, *Motion] {
	return &ValueCollection[Motion, *Motion]{
		id:    id,
		fetch: r,
	}
}

// ChatGroup has all fields from chat_group.
type ChatGroup struct {
	ReadGroupIDs   []int
	Weight         int
	WriteGroupIDs  []int
	ChatMessageIDs []int
	ID             int
	MeetingID      int
	Name           string
}

func (t *ChatGroup) lazy(ds *Fetch, id int) {
	ds.ChatGroup_ReadGroupIDs(id).Lazy(&t.ReadGroupIDs)
	ds.ChatGroup_Weight(id).Lazy(&t.Weight)
	ds.ChatGroup_WriteGroupIDs(id).Lazy(&t.WriteGroupIDs)
	ds.ChatGroup_ChatMessageIDs(id).Lazy(&t.ChatMessageIDs)
	ds.ChatGroup_ID(id).Lazy(&t.ID)
	ds.ChatGroup_MeetingID(id).Lazy(&t.MeetingID)
	ds.ChatGroup_Name(id).Lazy(&t.Name)
}

func (t *ChatGroup) preload(ds *Fetch, id int) {
	ds.ChatGroup_ReadGroupIDs(id).Preload()
	ds.ChatGroup_Weight(id).Preload()
	ds.ChatGroup_WriteGroupIDs(id).Preload()
	ds.ChatGroup_ChatMessageIDs(id).Preload()
	ds.ChatGroup_ID(id).Preload()
	ds.ChatGroup_MeetingID(id).Preload()
	ds.ChatGroup_Name(id).Preload()
}

func (r *Fetch) ChatGroup(id int) *ValueCollection[ChatGroup, *ChatGroup] {
	return &ValueCollection[ChatGroup, *ChatGroup]{
		id:    id,
		fetch: r,
	}
}

// ChatMessage has all fields from chat_message.
type ChatMessage struct {
	MeetingUserID Maybe[int]
	ChatGroupID   int
	Content       string
	Created       int
	ID            int
	MeetingID     int
}

func (t *ChatMessage) lazy(ds *Fetch, id int) {
	ds.ChatMessage_MeetingUserID(id).Lazy(&t.MeetingUserID)
	ds.ChatMessage_ChatGroupID(id).Lazy(&t.ChatGroupID)
	ds.ChatMessage_Content(id).Lazy(&t.Content)
	ds.ChatMessage_Created(id).Lazy(&t.Created)
	ds.ChatMessage_ID(id).Lazy(&t.ID)
	ds.ChatMessage_MeetingID(id).Lazy(&t.MeetingID)
}

func (t *ChatMessage) preload(ds *Fetch, id int) {
	ds.ChatMessage_MeetingUserID(id).Preload()
	ds.ChatMessage_ChatGroupID(id).Preload()
	ds.ChatMessage_Content(id).Preload()
	ds.ChatMessage_Created(id).Preload()
	ds.ChatMessage_ID(id).Preload()
	ds.ChatMessage_MeetingID(id).Preload()
}

func (r *Fetch) ChatMessage(id int) *ValueCollection[ChatMessage, *ChatMessage] {
	return &ValueCollection[ChatMessage, *ChatMessage]{
		id:    id,
		fetch: r,
	}
}

// Committee has all fields from committee.
type Committee struct {
	ForwardToCommitteeIDs              []int
	ForwardingUserID                   Maybe[int]
	ID                                 int
	ManagerIDs                         []int
	DefaultMeetingID                   Maybe[int]
	ExternalID                         string
	Name                               string
	OrganizationID                     int
	OrganizationTagIDs                 []int
	ReceiveForwardingsFromCommitteeIDs []int
	UserIDs                            []int
	Description                        string
	MeetingIDs                         []int
}

func (t *Committee) lazy(ds *Fetch, id int) {
	ds.Committee_ForwardToCommitteeIDs(id).Lazy(&t.ForwardToCommitteeIDs)
	ds.Committee_ForwardingUserID(id).Lazy(&t.ForwardingUserID)
	ds.Committee_ID(id).Lazy(&t.ID)
	ds.Committee_ManagerIDs(id).Lazy(&t.ManagerIDs)
	ds.Committee_DefaultMeetingID(id).Lazy(&t.DefaultMeetingID)
	ds.Committee_ExternalID(id).Lazy(&t.ExternalID)
	ds.Committee_Name(id).Lazy(&t.Name)
	ds.Committee_OrganizationID(id).Lazy(&t.OrganizationID)
	ds.Committee_OrganizationTagIDs(id).Lazy(&t.OrganizationTagIDs)
	ds.Committee_ReceiveForwardingsFromCommitteeIDs(id).Lazy(&t.ReceiveForwardingsFromCommitteeIDs)
	ds.Committee_UserIDs(id).Lazy(&t.UserIDs)
	ds.Committee_Description(id).Lazy(&t.Description)
	ds.Committee_MeetingIDs(id).Lazy(&t.MeetingIDs)
}

func (t *Committee) preload(ds *Fetch, id int) {
	ds.Committee_ForwardToCommitteeIDs(id).Preload()
	ds.Committee_ForwardingUserID(id).Preload()
	ds.Committee_ID(id).Preload()
	ds.Committee_ManagerIDs(id).Preload()
	ds.Committee_DefaultMeetingID(id).Preload()
	ds.Committee_ExternalID(id).Preload()
	ds.Committee_Name(id).Preload()
	ds.Committee_OrganizationID(id).Preload()
	ds.Committee_OrganizationTagIDs(id).Preload()
	ds.Committee_ReceiveForwardingsFromCommitteeIDs(id).Preload()
	ds.Committee_UserIDs(id).Preload()
	ds.Committee_Description(id).Preload()
	ds.Committee_MeetingIDs(id).Preload()
}

func (r *Fetch) Committee(id int) *ValueCollection[Committee, *Committee] {
	return &ValueCollection[Committee, *Committee]{
		id:    id,
		fetch: r,
	}
}

// MeetingMediafile has all fields from meeting_mediafile.
type MeetingMediafile struct {
	IsPublic                               bool
	UsedAsFontChyronSpeakerNameInMeetingID Maybe[int]
	UsedAsLogoPdfFooterRInMeetingID        Maybe[int]
	UsedAsLogoPdfHeaderLInMeetingID        Maybe[int]
	UsedAsLogoPdfHeaderRInMeetingID        Maybe[int]
	AccessGroupIDs                         []int
	AttachmentIDs                          []string
	MediafileID                            int
	UsedAsFontProjectorH1InMeetingID       Maybe[int]
	UsedAsLogoPdfBallotPaperInMeetingID    Maybe[int]
	UsedAsFontRegularInMeetingID           Maybe[int]
	UsedAsLogoProjectorMainInMeetingID     Maybe[int]
	ListOfSpeakersID                       Maybe[int]
	MeetingID                              int
	ProjectionIDs                          []int
	UsedAsFontBoldItalicInMeetingID        Maybe[int]
	UsedAsFontItalicInMeetingID            Maybe[int]
	UsedAsFontProjectorH2InMeetingID       Maybe[int]
	UsedAsLogoWebHeaderInMeetingID         Maybe[int]
	ID                                     int
	InheritedAccessGroupIDs                []int
	UsedAsFontBoldInMeetingID              Maybe[int]
	UsedAsFontMonospaceInMeetingID         Maybe[int]
	UsedAsLogoPdfFooterLInMeetingID        Maybe[int]
	UsedAsLogoProjectorHeaderInMeetingID   Maybe[int]
}

func (t *MeetingMediafile) lazy(ds *Fetch, id int) {
	ds.MeetingMediafile_IsPublic(id).Lazy(&t.IsPublic)
	ds.MeetingMediafile_UsedAsFontChyronSpeakerNameInMeetingID(id).Lazy(&t.UsedAsFontChyronSpeakerNameInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfFooterRInMeetingID(id).Lazy(&t.UsedAsLogoPdfFooterRInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfHeaderLInMeetingID(id).Lazy(&t.UsedAsLogoPdfHeaderLInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfHeaderRInMeetingID(id).Lazy(&t.UsedAsLogoPdfHeaderRInMeetingID)
	ds.MeetingMediafile_AccessGroupIDs(id).Lazy(&t.AccessGroupIDs)
	ds.MeetingMediafile_AttachmentIDs(id).Lazy(&t.AttachmentIDs)
	ds.MeetingMediafile_MediafileID(id).Lazy(&t.MediafileID)
	ds.MeetingMediafile_UsedAsFontProjectorH1InMeetingID(id).Lazy(&t.UsedAsFontProjectorH1InMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfBallotPaperInMeetingID(id).Lazy(&t.UsedAsLogoPdfBallotPaperInMeetingID)
	ds.MeetingMediafile_UsedAsFontRegularInMeetingID(id).Lazy(&t.UsedAsFontRegularInMeetingID)
	ds.MeetingMediafile_UsedAsLogoProjectorMainInMeetingID(id).Lazy(&t.UsedAsLogoProjectorMainInMeetingID)
	ds.MeetingMediafile_ListOfSpeakersID(id).Lazy(&t.ListOfSpeakersID)
	ds.MeetingMediafile_MeetingID(id).Lazy(&t.MeetingID)
	ds.MeetingMediafile_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.MeetingMediafile_UsedAsFontBoldItalicInMeetingID(id).Lazy(&t.UsedAsFontBoldItalicInMeetingID)
	ds.MeetingMediafile_UsedAsFontItalicInMeetingID(id).Lazy(&t.UsedAsFontItalicInMeetingID)
	ds.MeetingMediafile_UsedAsFontProjectorH2InMeetingID(id).Lazy(&t.UsedAsFontProjectorH2InMeetingID)
	ds.MeetingMediafile_UsedAsLogoWebHeaderInMeetingID(id).Lazy(&t.UsedAsLogoWebHeaderInMeetingID)
	ds.MeetingMediafile_ID(id).Lazy(&t.ID)
	ds.MeetingMediafile_InheritedAccessGroupIDs(id).Lazy(&t.InheritedAccessGroupIDs)
	ds.MeetingMediafile_UsedAsFontBoldInMeetingID(id).Lazy(&t.UsedAsFontBoldInMeetingID)
	ds.MeetingMediafile_UsedAsFontMonospaceInMeetingID(id).Lazy(&t.UsedAsFontMonospaceInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfFooterLInMeetingID(id).Lazy(&t.UsedAsLogoPdfFooterLInMeetingID)
	ds.MeetingMediafile_UsedAsLogoProjectorHeaderInMeetingID(id).Lazy(&t.UsedAsLogoProjectorHeaderInMeetingID)
}

func (t *MeetingMediafile) preload(ds *Fetch, id int) {
	ds.MeetingMediafile_IsPublic(id).Preload()
	ds.MeetingMediafile_UsedAsFontChyronSpeakerNameInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfFooterRInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfHeaderLInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfHeaderRInMeetingID(id).Preload()
	ds.MeetingMediafile_AccessGroupIDs(id).Preload()
	ds.MeetingMediafile_AttachmentIDs(id).Preload()
	ds.MeetingMediafile_MediafileID(id).Preload()
	ds.MeetingMediafile_UsedAsFontProjectorH1InMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfBallotPaperInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontRegularInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoProjectorMainInMeetingID(id).Preload()
	ds.MeetingMediafile_ListOfSpeakersID(id).Preload()
	ds.MeetingMediafile_MeetingID(id).Preload()
	ds.MeetingMediafile_ProjectionIDs(id).Preload()
	ds.MeetingMediafile_UsedAsFontBoldItalicInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontItalicInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontProjectorH2InMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoWebHeaderInMeetingID(id).Preload()
	ds.MeetingMediafile_ID(id).Preload()
	ds.MeetingMediafile_InheritedAccessGroupIDs(id).Preload()
	ds.MeetingMediafile_UsedAsFontBoldInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsFontMonospaceInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoPdfFooterLInMeetingID(id).Preload()
	ds.MeetingMediafile_UsedAsLogoProjectorHeaderInMeetingID(id).Preload()
}

func (r *Fetch) MeetingMediafile(id int) *ValueCollection[MeetingMediafile, *MeetingMediafile] {
	return &ValueCollection[MeetingMediafile, *MeetingMediafile]{
		id:    id,
		fetch: r,
	}
}

// MotionWorkflow has all fields from motion_workflow.
type MotionWorkflow struct {
	Name                              string
	SequentialNumber                  int
	StateIDs                          []int
	DefaultAmendmentWorkflowMeetingID Maybe[int]
	DefaultWorkflowMeetingID          Maybe[int]
	FirstStateID                      int
	ID                                int
	MeetingID                         int
}

func (t *MotionWorkflow) lazy(ds *Fetch, id int) {
	ds.MotionWorkflow_Name(id).Lazy(&t.Name)
	ds.MotionWorkflow_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.MotionWorkflow_StateIDs(id).Lazy(&t.StateIDs)
	ds.MotionWorkflow_DefaultAmendmentWorkflowMeetingID(id).Lazy(&t.DefaultAmendmentWorkflowMeetingID)
	ds.MotionWorkflow_DefaultWorkflowMeetingID(id).Lazy(&t.DefaultWorkflowMeetingID)
	ds.MotionWorkflow_FirstStateID(id).Lazy(&t.FirstStateID)
	ds.MotionWorkflow_ID(id).Lazy(&t.ID)
	ds.MotionWorkflow_MeetingID(id).Lazy(&t.MeetingID)
}

func (t *MotionWorkflow) preload(ds *Fetch, id int) {
	ds.MotionWorkflow_Name(id).Preload()
	ds.MotionWorkflow_SequentialNumber(id).Preload()
	ds.MotionWorkflow_StateIDs(id).Preload()
	ds.MotionWorkflow_DefaultAmendmentWorkflowMeetingID(id).Preload()
	ds.MotionWorkflow_DefaultWorkflowMeetingID(id).Preload()
	ds.MotionWorkflow_FirstStateID(id).Preload()
	ds.MotionWorkflow_ID(id).Preload()
	ds.MotionWorkflow_MeetingID(id).Preload()
}

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
}

func (t *MotionWorkingGroupSpeaker) lazy(ds *Fetch, id int) {
	ds.MotionWorkingGroupSpeaker_ID(id).Lazy(&t.ID)
	ds.MotionWorkingGroupSpeaker_MeetingID(id).Lazy(&t.MeetingID)
	ds.MotionWorkingGroupSpeaker_MeetingUserID(id).Lazy(&t.MeetingUserID)
	ds.MotionWorkingGroupSpeaker_MotionID(id).Lazy(&t.MotionID)
	ds.MotionWorkingGroupSpeaker_Weight(id).Lazy(&t.Weight)
}

func (t *MotionWorkingGroupSpeaker) preload(ds *Fetch, id int) {
	ds.MotionWorkingGroupSpeaker_ID(id).Preload()
	ds.MotionWorkingGroupSpeaker_MeetingID(id).Preload()
	ds.MotionWorkingGroupSpeaker_MeetingUserID(id).Preload()
	ds.MotionWorkingGroupSpeaker_MotionID(id).Preload()
	ds.MotionWorkingGroupSpeaker_Weight(id).Preload()
}

func (r *Fetch) MotionWorkingGroupSpeaker(id int) *ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker] {
	return &ValueCollection[MotionWorkingGroupSpeaker, *MotionWorkingGroupSpeaker]{
		id:    id,
		fetch: r,
	}
}

// Option has all fields from option.
type Option struct {
	ID                         int
	No                         string
	PollID                     Maybe[int]
	Text                       string
	VoteIDs                    []int
	Weight                     int
	Abstain                    string
	ContentObjectID            Maybe[string]
	MeetingID                  int
	UsedAsGlobalOptionInPollID Maybe[int]
	Yes                        string
}

func (t *Option) lazy(ds *Fetch, id int) {
	ds.Option_ID(id).Lazy(&t.ID)
	ds.Option_No(id).Lazy(&t.No)
	ds.Option_PollID(id).Lazy(&t.PollID)
	ds.Option_Text(id).Lazy(&t.Text)
	ds.Option_VoteIDs(id).Lazy(&t.VoteIDs)
	ds.Option_Weight(id).Lazy(&t.Weight)
	ds.Option_Abstain(id).Lazy(&t.Abstain)
	ds.Option_ContentObjectID(id).Lazy(&t.ContentObjectID)
	ds.Option_MeetingID(id).Lazy(&t.MeetingID)
	ds.Option_UsedAsGlobalOptionInPollID(id).Lazy(&t.UsedAsGlobalOptionInPollID)
	ds.Option_Yes(id).Lazy(&t.Yes)
}

func (t *Option) preload(ds *Fetch, id int) {
	ds.Option_ID(id).Preload()
	ds.Option_No(id).Preload()
	ds.Option_PollID(id).Preload()
	ds.Option_Text(id).Preload()
	ds.Option_VoteIDs(id).Preload()
	ds.Option_Weight(id).Preload()
	ds.Option_Abstain(id).Preload()
	ds.Option_ContentObjectID(id).Preload()
	ds.Option_MeetingID(id).Preload()
	ds.Option_UsedAsGlobalOptionInPollID(id).Preload()
	ds.Option_Yes(id).Preload()
}

func (r *Fetch) Option(id int) *ValueCollection[Option, *Option] {
	return &ValueCollection[Option, *Option]{
		id:    id,
		fetch: r,
	}
}

// AgendaItem has all fields from agenda_item.
type AgendaItem struct {
	IsInternal      bool
	TagIDs          []int
	Type            string
	ContentObjectID string
	Duration        int
	ProjectionIDs   []int
	Comment         string
	ItemNumber      string
	MeetingID       int
	Weight          int
	ChildIDs        []int
	Closed          bool
	Level           int
	ParentID        Maybe[int]
	ID              int
	IsHidden        bool
}

func (t *AgendaItem) lazy(ds *Fetch, id int) {
	ds.AgendaItem_IsInternal(id).Lazy(&t.IsInternal)
	ds.AgendaItem_TagIDs(id).Lazy(&t.TagIDs)
	ds.AgendaItem_Type(id).Lazy(&t.Type)
	ds.AgendaItem_ContentObjectID(id).Lazy(&t.ContentObjectID)
	ds.AgendaItem_Duration(id).Lazy(&t.Duration)
	ds.AgendaItem_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.AgendaItem_Comment(id).Lazy(&t.Comment)
	ds.AgendaItem_ItemNumber(id).Lazy(&t.ItemNumber)
	ds.AgendaItem_MeetingID(id).Lazy(&t.MeetingID)
	ds.AgendaItem_Weight(id).Lazy(&t.Weight)
	ds.AgendaItem_ChildIDs(id).Lazy(&t.ChildIDs)
	ds.AgendaItem_Closed(id).Lazy(&t.Closed)
	ds.AgendaItem_Level(id).Lazy(&t.Level)
	ds.AgendaItem_ParentID(id).Lazy(&t.ParentID)
	ds.AgendaItem_ID(id).Lazy(&t.ID)
	ds.AgendaItem_IsHidden(id).Lazy(&t.IsHidden)
}

func (t *AgendaItem) preload(ds *Fetch, id int) {
	ds.AgendaItem_IsInternal(id).Preload()
	ds.AgendaItem_TagIDs(id).Preload()
	ds.AgendaItem_Type(id).Preload()
	ds.AgendaItem_ContentObjectID(id).Preload()
	ds.AgendaItem_Duration(id).Preload()
	ds.AgendaItem_ProjectionIDs(id).Preload()
	ds.AgendaItem_Comment(id).Preload()
	ds.AgendaItem_ItemNumber(id).Preload()
	ds.AgendaItem_MeetingID(id).Preload()
	ds.AgendaItem_Weight(id).Preload()
	ds.AgendaItem_ChildIDs(id).Preload()
	ds.AgendaItem_Closed(id).Preload()
	ds.AgendaItem_Level(id).Preload()
	ds.AgendaItem_ParentID(id).Preload()
	ds.AgendaItem_ID(id).Preload()
	ds.AgendaItem_IsHidden(id).Preload()
}

func (r *Fetch) AgendaItem(id int) *ValueCollection[AgendaItem, *AgendaItem] {
	return &ValueCollection[AgendaItem, *AgendaItem]{
		id:    id,
		fetch: r,
	}
}

// ProjectorCountdown has all fields from projector_countdown.
type ProjectorCountdown struct {
	Title                                  string
	UsedAsPollCountdownMeetingID           Maybe[int]
	CountdownTime                          float32
	DefaultTime                            int
	ID                                     int
	MeetingID                              int
	ProjectionIDs                          []int
	Description                            string
	Running                                bool
	UsedAsListOfSpeakersCountdownMeetingID Maybe[int]
}

func (t *ProjectorCountdown) lazy(ds *Fetch, id int) {
	ds.ProjectorCountdown_Title(id).Lazy(&t.Title)
	ds.ProjectorCountdown_UsedAsPollCountdownMeetingID(id).Lazy(&t.UsedAsPollCountdownMeetingID)
	ds.ProjectorCountdown_CountdownTime(id).Lazy(&t.CountdownTime)
	ds.ProjectorCountdown_DefaultTime(id).Lazy(&t.DefaultTime)
	ds.ProjectorCountdown_ID(id).Lazy(&t.ID)
	ds.ProjectorCountdown_MeetingID(id).Lazy(&t.MeetingID)
	ds.ProjectorCountdown_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.ProjectorCountdown_Description(id).Lazy(&t.Description)
	ds.ProjectorCountdown_Running(id).Lazy(&t.Running)
	ds.ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(id).Lazy(&t.UsedAsListOfSpeakersCountdownMeetingID)
}

func (t *ProjectorCountdown) preload(ds *Fetch, id int) {
	ds.ProjectorCountdown_Title(id).Preload()
	ds.ProjectorCountdown_UsedAsPollCountdownMeetingID(id).Preload()
	ds.ProjectorCountdown_CountdownTime(id).Preload()
	ds.ProjectorCountdown_DefaultTime(id).Preload()
	ds.ProjectorCountdown_ID(id).Preload()
	ds.ProjectorCountdown_MeetingID(id).Preload()
	ds.ProjectorCountdown_ProjectionIDs(id).Preload()
	ds.ProjectorCountdown_Description(id).Preload()
	ds.ProjectorCountdown_Running(id).Preload()
	ds.ProjectorCountdown_UsedAsListOfSpeakersCountdownMeetingID(id).Preload()
}

func (r *Fetch) ProjectorCountdown(id int) *ValueCollection[ProjectorCountdown, *ProjectorCountdown] {
	return &ValueCollection[ProjectorCountdown, *ProjectorCountdown]{
		id:    id,
		fetch: r,
	}
}

// PointOfOrderCategory has all fields from point_of_order_category.
type PointOfOrderCategory struct {
	Text       string
	ID         int
	MeetingID  int
	Rank       int
	SpeakerIDs []int
}

func (t *PointOfOrderCategory) lazy(ds *Fetch, id int) {
	ds.PointOfOrderCategory_Text(id).Lazy(&t.Text)
	ds.PointOfOrderCategory_ID(id).Lazy(&t.ID)
	ds.PointOfOrderCategory_MeetingID(id).Lazy(&t.MeetingID)
	ds.PointOfOrderCategory_Rank(id).Lazy(&t.Rank)
	ds.PointOfOrderCategory_SpeakerIDs(id).Lazy(&t.SpeakerIDs)
}

func (t *PointOfOrderCategory) preload(ds *Fetch, id int) {
	ds.PointOfOrderCategory_Text(id).Preload()
	ds.PointOfOrderCategory_ID(id).Preload()
	ds.PointOfOrderCategory_MeetingID(id).Preload()
	ds.PointOfOrderCategory_Rank(id).Preload()
	ds.PointOfOrderCategory_SpeakerIDs(id).Preload()
}

func (r *Fetch) PointOfOrderCategory(id int) *ValueCollection[PointOfOrderCategory, *PointOfOrderCategory] {
	return &ValueCollection[PointOfOrderCategory, *PointOfOrderCategory]{
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
}

func (t *AssignmentCandidate) lazy(ds *Fetch, id int) {
	ds.AssignmentCandidate_AssignmentID(id).Lazy(&t.AssignmentID)
	ds.AssignmentCandidate_ID(id).Lazy(&t.ID)
	ds.AssignmentCandidate_MeetingID(id).Lazy(&t.MeetingID)
	ds.AssignmentCandidate_MeetingUserID(id).Lazy(&t.MeetingUserID)
	ds.AssignmentCandidate_Weight(id).Lazy(&t.Weight)
}

func (t *AssignmentCandidate) preload(ds *Fetch, id int) {
	ds.AssignmentCandidate_AssignmentID(id).Preload()
	ds.AssignmentCandidate_ID(id).Preload()
	ds.AssignmentCandidate_MeetingID(id).Preload()
	ds.AssignmentCandidate_MeetingUserID(id).Preload()
	ds.AssignmentCandidate_Weight(id).Preload()
}

func (r *Fetch) AssignmentCandidate(id int) *ValueCollection[AssignmentCandidate, *AssignmentCandidate] {
	return &ValueCollection[AssignmentCandidate, *AssignmentCandidate]{
		id:    id,
		fetch: r,
	}
}

// Mediafile has all fields from mediafile.
type Mediafile struct {
	Mimetype                            string
	PublishedToMeetingsInOrganizationID Maybe[int]
	Filename                            string
	ID                                  int
	MeetingMediafileIDs                 []int
	PdfInformation                      json.RawMessage
	Token                               string
	ChildIDs                            []int
	OwnerID                             string
	ParentID                            Maybe[int]
	CreateTimestamp                     int
	Filesize                            int
	IsDirectory                         bool
	Title                               string
}

func (t *Mediafile) lazy(ds *Fetch, id int) {
	ds.Mediafile_Mimetype(id).Lazy(&t.Mimetype)
	ds.Mediafile_PublishedToMeetingsInOrganizationID(id).Lazy(&t.PublishedToMeetingsInOrganizationID)
	ds.Mediafile_Filename(id).Lazy(&t.Filename)
	ds.Mediafile_ID(id).Lazy(&t.ID)
	ds.Mediafile_MeetingMediafileIDs(id).Lazy(&t.MeetingMediafileIDs)
	ds.Mediafile_PdfInformation(id).Lazy(&t.PdfInformation)
	ds.Mediafile_Token(id).Lazy(&t.Token)
	ds.Mediafile_ChildIDs(id).Lazy(&t.ChildIDs)
	ds.Mediafile_OwnerID(id).Lazy(&t.OwnerID)
	ds.Mediafile_ParentID(id).Lazy(&t.ParentID)
	ds.Mediafile_CreateTimestamp(id).Lazy(&t.CreateTimestamp)
	ds.Mediafile_Filesize(id).Lazy(&t.Filesize)
	ds.Mediafile_IsDirectory(id).Lazy(&t.IsDirectory)
	ds.Mediafile_Title(id).Lazy(&t.Title)
}

func (t *Mediafile) preload(ds *Fetch, id int) {
	ds.Mediafile_Mimetype(id).Preload()
	ds.Mediafile_PublishedToMeetingsInOrganizationID(id).Preload()
	ds.Mediafile_Filename(id).Preload()
	ds.Mediafile_ID(id).Preload()
	ds.Mediafile_MeetingMediafileIDs(id).Preload()
	ds.Mediafile_PdfInformation(id).Preload()
	ds.Mediafile_Token(id).Preload()
	ds.Mediafile_ChildIDs(id).Preload()
	ds.Mediafile_OwnerID(id).Preload()
	ds.Mediafile_ParentID(id).Preload()
	ds.Mediafile_CreateTimestamp(id).Preload()
	ds.Mediafile_Filesize(id).Preload()
	ds.Mediafile_IsDirectory(id).Preload()
	ds.Mediafile_Title(id).Preload()
}

func (r *Fetch) Mediafile(id int) *ValueCollection[Mediafile, *Mediafile] {
	return &ValueCollection[Mediafile, *Mediafile]{
		id:    id,
		fetch: r,
	}
}

// MeetingUser has all fields from meeting_user.
type MeetingUser struct {
	GroupIDs                     []int
	ID                           int
	MotionSubmitterIDs           []int
	VoteDelegatedToID            Maybe[int]
	MeetingID                    int
	Number                       string
	SupportedMotionIDs           []int
	VoteWeight                   string
	AboutMe                      string
	Comment                      string
	MotionWorkingGroupSpeakerIDs []int
	StructureLevelIDs            []int
	SpeakerIDs                   []int
	UserID                       int
	VoteDelegationsFromIDs       []int
	AssignmentCandidateIDs       []int
	ChatMessageIDs               []int
	LockedOut                    bool
	MotionEditorIDs              []int
	PersonalNoteIDs              []int
}

func (t *MeetingUser) lazy(ds *Fetch, id int) {
	ds.MeetingUser_GroupIDs(id).Lazy(&t.GroupIDs)
	ds.MeetingUser_ID(id).Lazy(&t.ID)
	ds.MeetingUser_MotionSubmitterIDs(id).Lazy(&t.MotionSubmitterIDs)
	ds.MeetingUser_VoteDelegatedToID(id).Lazy(&t.VoteDelegatedToID)
	ds.MeetingUser_MeetingID(id).Lazy(&t.MeetingID)
	ds.MeetingUser_Number(id).Lazy(&t.Number)
	ds.MeetingUser_SupportedMotionIDs(id).Lazy(&t.SupportedMotionIDs)
	ds.MeetingUser_VoteWeight(id).Lazy(&t.VoteWeight)
	ds.MeetingUser_AboutMe(id).Lazy(&t.AboutMe)
	ds.MeetingUser_Comment(id).Lazy(&t.Comment)
	ds.MeetingUser_MotionWorkingGroupSpeakerIDs(id).Lazy(&t.MotionWorkingGroupSpeakerIDs)
	ds.MeetingUser_StructureLevelIDs(id).Lazy(&t.StructureLevelIDs)
	ds.MeetingUser_SpeakerIDs(id).Lazy(&t.SpeakerIDs)
	ds.MeetingUser_UserID(id).Lazy(&t.UserID)
	ds.MeetingUser_VoteDelegationsFromIDs(id).Lazy(&t.VoteDelegationsFromIDs)
	ds.MeetingUser_AssignmentCandidateIDs(id).Lazy(&t.AssignmentCandidateIDs)
	ds.MeetingUser_ChatMessageIDs(id).Lazy(&t.ChatMessageIDs)
	ds.MeetingUser_LockedOut(id).Lazy(&t.LockedOut)
	ds.MeetingUser_MotionEditorIDs(id).Lazy(&t.MotionEditorIDs)
	ds.MeetingUser_PersonalNoteIDs(id).Lazy(&t.PersonalNoteIDs)
}

func (t *MeetingUser) preload(ds *Fetch, id int) {
	ds.MeetingUser_GroupIDs(id).Preload()
	ds.MeetingUser_ID(id).Preload()
	ds.MeetingUser_MotionSubmitterIDs(id).Preload()
	ds.MeetingUser_VoteDelegatedToID(id).Preload()
	ds.MeetingUser_MeetingID(id).Preload()
	ds.MeetingUser_Number(id).Preload()
	ds.MeetingUser_SupportedMotionIDs(id).Preload()
	ds.MeetingUser_VoteWeight(id).Preload()
	ds.MeetingUser_AboutMe(id).Preload()
	ds.MeetingUser_Comment(id).Preload()
	ds.MeetingUser_MotionWorkingGroupSpeakerIDs(id).Preload()
	ds.MeetingUser_StructureLevelIDs(id).Preload()
	ds.MeetingUser_SpeakerIDs(id).Preload()
	ds.MeetingUser_UserID(id).Preload()
	ds.MeetingUser_VoteDelegationsFromIDs(id).Preload()
	ds.MeetingUser_AssignmentCandidateIDs(id).Preload()
	ds.MeetingUser_ChatMessageIDs(id).Preload()
	ds.MeetingUser_LockedOut(id).Preload()
	ds.MeetingUser_MotionEditorIDs(id).Preload()
	ds.MeetingUser_PersonalNoteIDs(id).Preload()
}

func (r *Fetch) MeetingUser(id int) *ValueCollection[MeetingUser, *MeetingUser] {
	return &ValueCollection[MeetingUser, *MeetingUser]{
		id:    id,
		fetch: r,
	}
}

// MotionBlock has all fields from motion_block.
type MotionBlock struct {
	MotionIDs        []int
	ProjectionIDs    []int
	SequentialNumber int
	Title            string
	MeetingID        int
	ID               int
	Internal         bool
	ListOfSpeakersID int
	AgendaItemID     Maybe[int]
}

func (t *MotionBlock) lazy(ds *Fetch, id int) {
	ds.MotionBlock_MotionIDs(id).Lazy(&t.MotionIDs)
	ds.MotionBlock_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.MotionBlock_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.MotionBlock_Title(id).Lazy(&t.Title)
	ds.MotionBlock_MeetingID(id).Lazy(&t.MeetingID)
	ds.MotionBlock_ID(id).Lazy(&t.ID)
	ds.MotionBlock_Internal(id).Lazy(&t.Internal)
	ds.MotionBlock_ListOfSpeakersID(id).Lazy(&t.ListOfSpeakersID)
	ds.MotionBlock_AgendaItemID(id).Lazy(&t.AgendaItemID)
}

func (t *MotionBlock) preload(ds *Fetch, id int) {
	ds.MotionBlock_MotionIDs(id).Preload()
	ds.MotionBlock_ProjectionIDs(id).Preload()
	ds.MotionBlock_SequentialNumber(id).Preload()
	ds.MotionBlock_Title(id).Preload()
	ds.MotionBlock_MeetingID(id).Preload()
	ds.MotionBlock_ID(id).Preload()
	ds.MotionBlock_Internal(id).Preload()
	ds.MotionBlock_ListOfSpeakersID(id).Preload()
	ds.MotionBlock_AgendaItemID(id).Preload()
}

func (r *Fetch) MotionBlock(id int) *ValueCollection[MotionBlock, *MotionBlock] {
	return &ValueCollection[MotionBlock, *MotionBlock]{
		id:    id,
		fetch: r,
	}
}

// MotionChangeRecommendation has all fields from motion_change_recommendation.
type MotionChangeRecommendation struct {
	CreationTime     int
	MeetingID        int
	OtherDescription string
	Rejected         bool
	Type             string
	ID               int
	Internal         bool
	LineFrom         int
	LineTo           int
	MotionID         int
	Text             string
}

func (t *MotionChangeRecommendation) lazy(ds *Fetch, id int) {
	ds.MotionChangeRecommendation_CreationTime(id).Lazy(&t.CreationTime)
	ds.MotionChangeRecommendation_MeetingID(id).Lazy(&t.MeetingID)
	ds.MotionChangeRecommendation_OtherDescription(id).Lazy(&t.OtherDescription)
	ds.MotionChangeRecommendation_Rejected(id).Lazy(&t.Rejected)
	ds.MotionChangeRecommendation_Type(id).Lazy(&t.Type)
	ds.MotionChangeRecommendation_ID(id).Lazy(&t.ID)
	ds.MotionChangeRecommendation_Internal(id).Lazy(&t.Internal)
	ds.MotionChangeRecommendation_LineFrom(id).Lazy(&t.LineFrom)
	ds.MotionChangeRecommendation_LineTo(id).Lazy(&t.LineTo)
	ds.MotionChangeRecommendation_MotionID(id).Lazy(&t.MotionID)
	ds.MotionChangeRecommendation_Text(id).Lazy(&t.Text)
}

func (t *MotionChangeRecommendation) preload(ds *Fetch, id int) {
	ds.MotionChangeRecommendation_CreationTime(id).Preload()
	ds.MotionChangeRecommendation_MeetingID(id).Preload()
	ds.MotionChangeRecommendation_OtherDescription(id).Preload()
	ds.MotionChangeRecommendation_Rejected(id).Preload()
	ds.MotionChangeRecommendation_Type(id).Preload()
	ds.MotionChangeRecommendation_ID(id).Preload()
	ds.MotionChangeRecommendation_Internal(id).Preload()
	ds.MotionChangeRecommendation_LineFrom(id).Preload()
	ds.MotionChangeRecommendation_LineTo(id).Preload()
	ds.MotionChangeRecommendation_MotionID(id).Preload()
	ds.MotionChangeRecommendation_Text(id).Preload()
}

func (r *Fetch) MotionChangeRecommendation(id int) *ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation] {
	return &ValueCollection[MotionChangeRecommendation, *MotionChangeRecommendation]{
		id:    id,
		fetch: r,
	}
}

// MotionComment has all fields from motion_comment.
type MotionComment struct {
	MotionID  int
	SectionID int
	Comment   string
	ID        int
	MeetingID int
}

func (t *MotionComment) lazy(ds *Fetch, id int) {
	ds.MotionComment_MotionID(id).Lazy(&t.MotionID)
	ds.MotionComment_SectionID(id).Lazy(&t.SectionID)
	ds.MotionComment_Comment(id).Lazy(&t.Comment)
	ds.MotionComment_ID(id).Lazy(&t.ID)
	ds.MotionComment_MeetingID(id).Lazy(&t.MeetingID)
}

func (t *MotionComment) preload(ds *Fetch, id int) {
	ds.MotionComment_MotionID(id).Preload()
	ds.MotionComment_SectionID(id).Preload()
	ds.MotionComment_Comment(id).Preload()
	ds.MotionComment_ID(id).Preload()
	ds.MotionComment_MeetingID(id).Preload()
}

func (r *Fetch) MotionComment(id int) *ValueCollection[MotionComment, *MotionComment] {
	return &ValueCollection[MotionComment, *MotionComment]{
		id:    id,
		fetch: r,
	}
}

// MotionEditor has all fields from motion_editor.
type MotionEditor struct {
	MeetingUserID int
	MotionID      int
	Weight        int
	ID            int
	MeetingID     int
}

func (t *MotionEditor) lazy(ds *Fetch, id int) {
	ds.MotionEditor_MeetingUserID(id).Lazy(&t.MeetingUserID)
	ds.MotionEditor_MotionID(id).Lazy(&t.MotionID)
	ds.MotionEditor_Weight(id).Lazy(&t.Weight)
	ds.MotionEditor_ID(id).Lazy(&t.ID)
	ds.MotionEditor_MeetingID(id).Lazy(&t.MeetingID)
}

func (t *MotionEditor) preload(ds *Fetch, id int) {
	ds.MotionEditor_MeetingUserID(id).Preload()
	ds.MotionEditor_MotionID(id).Preload()
	ds.MotionEditor_Weight(id).Preload()
	ds.MotionEditor_ID(id).Preload()
	ds.MotionEditor_MeetingID(id).Preload()
}

func (r *Fetch) MotionEditor(id int) *ValueCollection[MotionEditor, *MotionEditor] {
	return &ValueCollection[MotionEditor, *MotionEditor]{
		id:    id,
		fetch: r,
	}
}

// Assignment has all fields from assignment.
type Assignment struct {
	AttachmentMeetingMediafileIDs []int
	CandidateIDs                  []int
	Description                   string
	ID                            int
	ListOfSpeakersID              int
	MeetingID                     int
	Phase                         string
	PollIDs                       []int
	ProjectionIDs                 []int
	Title                         string
	DefaultPollDescription        string
	NumberPollCandidates          bool
	OpenPosts                     int
	TagIDs                        []int
	AgendaItemID                  Maybe[int]
	SequentialNumber              int
}

func (t *Assignment) lazy(ds *Fetch, id int) {
	ds.Assignment_AttachmentMeetingMediafileIDs(id).Lazy(&t.AttachmentMeetingMediafileIDs)
	ds.Assignment_CandidateIDs(id).Lazy(&t.CandidateIDs)
	ds.Assignment_Description(id).Lazy(&t.Description)
	ds.Assignment_ID(id).Lazy(&t.ID)
	ds.Assignment_ListOfSpeakersID(id).Lazy(&t.ListOfSpeakersID)
	ds.Assignment_MeetingID(id).Lazy(&t.MeetingID)
	ds.Assignment_Phase(id).Lazy(&t.Phase)
	ds.Assignment_PollIDs(id).Lazy(&t.PollIDs)
	ds.Assignment_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.Assignment_Title(id).Lazy(&t.Title)
	ds.Assignment_DefaultPollDescription(id).Lazy(&t.DefaultPollDescription)
	ds.Assignment_NumberPollCandidates(id).Lazy(&t.NumberPollCandidates)
	ds.Assignment_OpenPosts(id).Lazy(&t.OpenPosts)
	ds.Assignment_TagIDs(id).Lazy(&t.TagIDs)
	ds.Assignment_AgendaItemID(id).Lazy(&t.AgendaItemID)
	ds.Assignment_SequentialNumber(id).Lazy(&t.SequentialNumber)
}

func (t *Assignment) preload(ds *Fetch, id int) {
	ds.Assignment_AttachmentMeetingMediafileIDs(id).Preload()
	ds.Assignment_CandidateIDs(id).Preload()
	ds.Assignment_Description(id).Preload()
	ds.Assignment_ID(id).Preload()
	ds.Assignment_ListOfSpeakersID(id).Preload()
	ds.Assignment_MeetingID(id).Preload()
	ds.Assignment_Phase(id).Preload()
	ds.Assignment_PollIDs(id).Preload()
	ds.Assignment_ProjectionIDs(id).Preload()
	ds.Assignment_Title(id).Preload()
	ds.Assignment_DefaultPollDescription(id).Preload()
	ds.Assignment_NumberPollCandidates(id).Preload()
	ds.Assignment_OpenPosts(id).Preload()
	ds.Assignment_TagIDs(id).Preload()
	ds.Assignment_AgendaItemID(id).Preload()
	ds.Assignment_SequentialNumber(id).Preload()
}

func (r *Fetch) Assignment(id int) *ValueCollection[Assignment, *Assignment] {
	return &ValueCollection[Assignment, *Assignment]{
		id:    id,
		fetch: r,
	}
}

// Meeting has all fields from meeting.
type Meeting struct {
	ReferenceProjectorID                         int
	AssignmentPollDefaultOnehundredPercentBase   string
	DefaultProjectorMotionIDs                    []int
	PollCoupleCountdown                          bool
	ConferenceOpenMicrophone                     bool
	LogoPdfHeaderRID                             Maybe[int]
	UsersEmailBody                               string
	ChatMessageIDs                               []int
	AssignmentPollDefaultGroupIDs                []int
	AssignmentPollDefaultType                    string
	MotionCommentSectionIDs                      []int
	MotionsPreamble                              string
	UsersPdfWlanSsid                             string
	WelcomeText                                  string
	PollBallotPaperNumber                        int
	TopicPollDefaultGroupIDs                     []int
	UsersEmailReplyto                            string
	ApplauseType                                 string
	DefaultProjectorAgendaItemListIDs            []int
	ListOfSpeakersAllowMultipleSpeakers          bool
	MotionsEnableEditor                          bool
	ProjectorCountdownDefaultTime                int
	ConferenceShow                               bool
	ExportPdfPageMarginLeft                      int
	GroupIDs                                     []int
	AgendaNewItemsDefaultVisibility              string
	JitsiRoomPassword                            string
	ListOfSpeakersIDs                            []int
	LogoWebHeaderID                              Maybe[int]
	MotionsRecommendationsBy                     string
	Description                                  string
	IsActiveInOrganizationID                     Maybe[int]
	StructureLevelIDs                            []int
	UsersAllowSelfSetPresent                     bool
	AgendaNumeralSystem                          string
	MotionsExportTitle                           string
	MotionsSupportersMinAmount                   int
	UsersForbidDelegatorAsSubmitter              bool
	UsersPdfWlanPassword                         string
	AnonymousGroupID                             Maybe[int]
	ConferenceStreamPosterUrl                    string
	FontItalicID                                 Maybe[int]
	PollCandidateIDs                             []int
	ProjectorIDs                                 []int
	AssignmentCandidateIDs                       []int
	DefaultProjectorCountdownIDs                 []int
	ListOfSpeakersAmountLastOnProjector          int
	ListOfSpeakersClosingDisablesPointOfOrder    bool
	MotionsAmendmentsOfAmendments                bool
	MotionsEnableTextOnProjector                 bool
	AgendaNumberPrefix                           string
	MotionChangeRecommendationIDs                []int
	UsersEmailSender                             string
	ExportPdfPagesize                            string
	LogoProjectorHeaderID                        Maybe[int]
	DefaultMeetingForCommitteeID                 Maybe[int]
	EndTime                                      int
	ListOfSpeakersInterventionTime               int
	MotionsDefaultWorkflowID                     int
	UsersForbidDelegatorInListOfSpeakers         bool
	AssignmentPollAddCandidatesToListOfSpeakers  bool
	AssignmentPollEnableMaxVotesPerOption        bool
	ListOfSpeakersCoupleCountdown                bool
	AgendaShowTopicNavigationOnDetailView        bool
	MotionPollDefaultBackend                     string
	MotionPollDefaultType                        string
	MotionEditorIDs                              []int
	Name                                         string
	StructureLevelListOfSpeakersIDs              []int
	AssignmentPollDefaultMethod                  string
	FontBoldID                                   Maybe[int]
	ListOfSpeakersEnablePointOfOrderCategories   bool
	LogoPdfFooterRID                             Maybe[int]
	MotionPollBallotPaperSelection               string
	MotionPollDefaultOnehundredPercentBase       string
	UsersPdfWlanEncryption                       string
	ListOfSpeakersAmountNextOnProjector          int
	ListOfSpeakersCanSetContributionSelf         bool
	ListOfSpeakersDefaultStructureLevelTime      int
	MotionsExportFollowRecommendation            bool
	AllProjectionIDs                             []int
	EnableAnonymous                              bool
	LogoPdfBallotPaperID                         Maybe[int]
	PollCountdownID                              Maybe[int]
	UserIDs                                      []int
	MotionCommentIDs                             []int
	UsersEnableVoteWeight                        bool
	DefaultGroupID                               int
	MeetingMediafileIDs                          []int
	UsersEmailSubject                            string
	JitsiRoomName                                string
	MotionWorkflowIDs                            []int
	MotionsEnableReasonOnProjector               bool
	OptionIDs                                    []int
	OrganizationTagIDs                           []int
	DefaultProjectorMessageIDs                   []int
	ExportPdfPageMarginRight                     int
	AssignmentPollSortPollResultByVotes          bool
	ConferenceStreamUrl                          string
	MeetingUserIDs                               []int
	VoteIDs                                      []int
	AgendaItemIDs                                []int
	ApplauseEnable                               bool
	DefaultProjectorMotionBlockIDs               []int
	PersonalNoteIDs                              []int
	PollDefaultMethod                            string
	PollIDs                                      []int
	ApplauseMaxAmount                            int
	ConferenceOpenVideo                          bool
	MediafileIDs                                 []int
	UsersPdfWelcometitle                         string
	ExportPdfLineHeight                          float32
	MotionWorkingGroupSpeakerIDs                 []int
	MotionsEnableSideboxOnProjector              bool
	MotionsExportPreamble                        string
	AgendaEnableNumbering                        bool
	AssignmentPollBallotPaperSelection           string
	FontRegularID                                Maybe[int]
	JitsiDomain                                  string
	MotionsAmendmentsInMainList                  bool
	MotionsNumberWithBlank                       bool
	AssignmentPollBallotPaperNumber              int
	PollDefaultGroupIDs                          []int
	SpeakerIDs                                   []int
	PollBallotPaperSelection                     string
	PresentUserIDs                               []int
	AssignmentPollDefaultBackend                 string
	FontProjectorH1ID                            Maybe[int]
	ListOfSpeakersCanCreatePointOfOrderForOthers bool
	MotionsAmendmentsMultipleParagraphs          bool
	MotionsBlockSlideColumns                     int
	MotionsNumberType                            string
	MotionIDs                                    []int
	ImportedAt                                   int
	MotionsCreateEnableAdditionalSubmitterText   bool
	MotionsShowReferringMotions                  bool
	FontBoldItalicID                             Maybe[int]
	LockedFromInside                             bool
	ApplauseTimeout                              int
	ProjectionIDs                                []int
	IsArchivedInOrganizationID                   Maybe[int]
	Location                                     string
	PollDefaultOnehundredPercentBase             string
	UsersEnableVoteDelegations                   bool
	ListOfSpeakersCountdownID                    Maybe[int]
	MotionsAmendmentsEnabled                     bool
	MotionsRecommendationTextMode                string
	DefaultProjectorListOfSpeakersIDs            []int
	DefaultProjectorTopicIDs                     []int
	ListOfSpeakersSpeakerNoteForEveryone         bool
	MotionBlockIDs                               []int
	LogoPdfFooterLID                             Maybe[int]
	MotionsAmendmentsTextMode                    string
	PointOfOrderCategoryIDs                      []int
	UsersPdfWelcometext                          string
	WelcomeTitle                                 string
	AgendaShowInternalItemsOnProjector           bool
	ExportCsvSeparator                           string
	ListOfSpeakersEnableInterposedQuestion       bool
	LogoPdfHeaderLID                             Maybe[int]
	ProjectorCountdownIDs                        []int
	UsersForbidDelegatorToVote                   bool
	ApplauseShowLevel                            bool
	ListOfSpeakersEnableProContraSpeech          bool
	ListOfSpeakersHideContributionCount          bool
	ListOfSpeakersPresentUsersOnly               bool
	LogoProjectorMainID                          Maybe[int]
	MotionStateIDs                               []int
	ListOfSpeakersInitiallyClosed                bool
	TopicIDs                                     []int
	DefaultProjectorCurrentListOfSpeakersIDs     []int
	DefaultProjectorPollIDs                      []int
	Language                                     string
	ConferenceAutoConnectNextSpeakers            int
	ExportPdfPagenumberAlignment                 string
	FontProjectorH2ID                            Maybe[int]
	MotionSubmitterIDs                           []int
	StartTime                                    int
	UsersEnablePresenceView                      bool
	AssignmentsExportPreamble                    string
	ExportPdfFontsize                            int
	ExportPdfPageMarginTop                       int
	ListOfSpeakersShowAmountOfSpeakersOnSlide    bool
	MotionsAmendmentsPrefix                      string
	MotionsEnableRecommendationOnProjector       bool
	ExportPdfPageMarginBottom                    int
	ListOfSpeakersShowFirstContribution          bool
	MotionsNumberMinDigits                       int
	DefaultProjectorAssignmentIDs                []int
	ID                                           int
	MotionsReasonRequired                        bool
	TemplateForOrganizationID                    Maybe[int]
	MotionPollBallotPaperNumber                  int
	FontChyronSpeakerNameID                      Maybe[int]
	PollCandidateListIDs                         []int
	MotionsEnableWorkingGroupSpeaker             bool
	MotionsHideMetadataBackground                bool
	ApplauseMinAmount                            int
	ApplauseParticleImageUrl                     string
	ChatGroupIDs                                 []int
	ConferenceEnableHelpdesk                     bool
	CustomTranslations                           json.RawMessage
	MotionCategoryIDs                            []int
	ConferenceAutoConnect                        bool
	ExternalID                                   string
	MotionsLineLength                            int
	PollDefaultType                              string
	AgendaShowSubtitles                          bool
	CommitteeID                                  int
	FontMonospaceID                              Maybe[int]
	MotionsDefaultAmendmentWorkflowID            int
	MotionsDefaultLineNumbering                  string
	AssignmentsExportTitle                       string
	DefaultProjectorAssignmentPollIDs            []int
	ListOfSpeakersEnablePointOfOrderSpeakers     bool
	ExportCsvEncoding                            string
	PollSortPollResultByVotes                    bool
	TagIDs                                       []int
	ConferenceLosRestriction                     bool
	DefaultProjectorMotionPollIDs                []int
	AdminGroupID                                 Maybe[int]
	DefaultProjectorMediafileIDs                 []int
	MotionsExportSubmitterRecommendation         bool
	UsersForbidDelegatorAsSupporter              bool
	MotionPollDefaultGroupIDs                    []int
	MotionPollDefaultMethod                      string
	MotionsDefaultSorting                        string
	PollDefaultBackend                           string
	AssignmentIDs                                []int
	ProjectorMessageIDs                          []int
	AgendaItemCreation                           string
	DefaultProjectorAmendmentIDs                 []int
	ForwardedMotionIDs                           []int
	MotionsShowSequentialNumber                  bool
	ProjectorCountdownWarningTime                int
}

func (t *Meeting) lazy(ds *Fetch, id int) {
	ds.Meeting_ReferenceProjectorID(id).Lazy(&t.ReferenceProjectorID)
	ds.Meeting_AssignmentPollDefaultOnehundredPercentBase(id).Lazy(&t.AssignmentPollDefaultOnehundredPercentBase)
	ds.Meeting_DefaultProjectorMotionIDs(id).Lazy(&t.DefaultProjectorMotionIDs)
	ds.Meeting_PollCoupleCountdown(id).Lazy(&t.PollCoupleCountdown)
	ds.Meeting_ConferenceOpenMicrophone(id).Lazy(&t.ConferenceOpenMicrophone)
	ds.Meeting_LogoPdfHeaderRID(id).Lazy(&t.LogoPdfHeaderRID)
	ds.Meeting_UsersEmailBody(id).Lazy(&t.UsersEmailBody)
	ds.Meeting_ChatMessageIDs(id).Lazy(&t.ChatMessageIDs)
	ds.Meeting_AssignmentPollDefaultGroupIDs(id).Lazy(&t.AssignmentPollDefaultGroupIDs)
	ds.Meeting_AssignmentPollDefaultType(id).Lazy(&t.AssignmentPollDefaultType)
	ds.Meeting_MotionCommentSectionIDs(id).Lazy(&t.MotionCommentSectionIDs)
	ds.Meeting_MotionsPreamble(id).Lazy(&t.MotionsPreamble)
	ds.Meeting_UsersPdfWlanSsid(id).Lazy(&t.UsersPdfWlanSsid)
	ds.Meeting_WelcomeText(id).Lazy(&t.WelcomeText)
	ds.Meeting_PollBallotPaperNumber(id).Lazy(&t.PollBallotPaperNumber)
	ds.Meeting_TopicPollDefaultGroupIDs(id).Lazy(&t.TopicPollDefaultGroupIDs)
	ds.Meeting_UsersEmailReplyto(id).Lazy(&t.UsersEmailReplyto)
	ds.Meeting_ApplauseType(id).Lazy(&t.ApplauseType)
	ds.Meeting_DefaultProjectorAgendaItemListIDs(id).Lazy(&t.DefaultProjectorAgendaItemListIDs)
	ds.Meeting_ListOfSpeakersAllowMultipleSpeakers(id).Lazy(&t.ListOfSpeakersAllowMultipleSpeakers)
	ds.Meeting_MotionsEnableEditor(id).Lazy(&t.MotionsEnableEditor)
	ds.Meeting_ProjectorCountdownDefaultTime(id).Lazy(&t.ProjectorCountdownDefaultTime)
	ds.Meeting_ConferenceShow(id).Lazy(&t.ConferenceShow)
	ds.Meeting_ExportPdfPageMarginLeft(id).Lazy(&t.ExportPdfPageMarginLeft)
	ds.Meeting_GroupIDs(id).Lazy(&t.GroupIDs)
	ds.Meeting_AgendaNewItemsDefaultVisibility(id).Lazy(&t.AgendaNewItemsDefaultVisibility)
	ds.Meeting_JitsiRoomPassword(id).Lazy(&t.JitsiRoomPassword)
	ds.Meeting_ListOfSpeakersIDs(id).Lazy(&t.ListOfSpeakersIDs)
	ds.Meeting_LogoWebHeaderID(id).Lazy(&t.LogoWebHeaderID)
	ds.Meeting_MotionsRecommendationsBy(id).Lazy(&t.MotionsRecommendationsBy)
	ds.Meeting_Description(id).Lazy(&t.Description)
	ds.Meeting_IsActiveInOrganizationID(id).Lazy(&t.IsActiveInOrganizationID)
	ds.Meeting_StructureLevelIDs(id).Lazy(&t.StructureLevelIDs)
	ds.Meeting_UsersAllowSelfSetPresent(id).Lazy(&t.UsersAllowSelfSetPresent)
	ds.Meeting_AgendaNumeralSystem(id).Lazy(&t.AgendaNumeralSystem)
	ds.Meeting_MotionsExportTitle(id).Lazy(&t.MotionsExportTitle)
	ds.Meeting_MotionsSupportersMinAmount(id).Lazy(&t.MotionsSupportersMinAmount)
	ds.Meeting_UsersForbidDelegatorAsSubmitter(id).Lazy(&t.UsersForbidDelegatorAsSubmitter)
	ds.Meeting_UsersPdfWlanPassword(id).Lazy(&t.UsersPdfWlanPassword)
	ds.Meeting_AnonymousGroupID(id).Lazy(&t.AnonymousGroupID)
	ds.Meeting_ConferenceStreamPosterUrl(id).Lazy(&t.ConferenceStreamPosterUrl)
	ds.Meeting_FontItalicID(id).Lazy(&t.FontItalicID)
	ds.Meeting_PollCandidateIDs(id).Lazy(&t.PollCandidateIDs)
	ds.Meeting_ProjectorIDs(id).Lazy(&t.ProjectorIDs)
	ds.Meeting_AssignmentCandidateIDs(id).Lazy(&t.AssignmentCandidateIDs)
	ds.Meeting_DefaultProjectorCountdownIDs(id).Lazy(&t.DefaultProjectorCountdownIDs)
	ds.Meeting_ListOfSpeakersAmountLastOnProjector(id).Lazy(&t.ListOfSpeakersAmountLastOnProjector)
	ds.Meeting_ListOfSpeakersClosingDisablesPointOfOrder(id).Lazy(&t.ListOfSpeakersClosingDisablesPointOfOrder)
	ds.Meeting_MotionsAmendmentsOfAmendments(id).Lazy(&t.MotionsAmendmentsOfAmendments)
	ds.Meeting_MotionsEnableTextOnProjector(id).Lazy(&t.MotionsEnableTextOnProjector)
	ds.Meeting_AgendaNumberPrefix(id).Lazy(&t.AgendaNumberPrefix)
	ds.Meeting_MotionChangeRecommendationIDs(id).Lazy(&t.MotionChangeRecommendationIDs)
	ds.Meeting_UsersEmailSender(id).Lazy(&t.UsersEmailSender)
	ds.Meeting_ExportPdfPagesize(id).Lazy(&t.ExportPdfPagesize)
	ds.Meeting_LogoProjectorHeaderID(id).Lazy(&t.LogoProjectorHeaderID)
	ds.Meeting_DefaultMeetingForCommitteeID(id).Lazy(&t.DefaultMeetingForCommitteeID)
	ds.Meeting_EndTime(id).Lazy(&t.EndTime)
	ds.Meeting_ListOfSpeakersInterventionTime(id).Lazy(&t.ListOfSpeakersInterventionTime)
	ds.Meeting_MotionsDefaultWorkflowID(id).Lazy(&t.MotionsDefaultWorkflowID)
	ds.Meeting_UsersForbidDelegatorInListOfSpeakers(id).Lazy(&t.UsersForbidDelegatorInListOfSpeakers)
	ds.Meeting_AssignmentPollAddCandidatesToListOfSpeakers(id).Lazy(&t.AssignmentPollAddCandidatesToListOfSpeakers)
	ds.Meeting_AssignmentPollEnableMaxVotesPerOption(id).Lazy(&t.AssignmentPollEnableMaxVotesPerOption)
	ds.Meeting_ListOfSpeakersCoupleCountdown(id).Lazy(&t.ListOfSpeakersCoupleCountdown)
	ds.Meeting_AgendaShowTopicNavigationOnDetailView(id).Lazy(&t.AgendaShowTopicNavigationOnDetailView)
	ds.Meeting_MotionPollDefaultBackend(id).Lazy(&t.MotionPollDefaultBackend)
	ds.Meeting_MotionPollDefaultType(id).Lazy(&t.MotionPollDefaultType)
	ds.Meeting_MotionEditorIDs(id).Lazy(&t.MotionEditorIDs)
	ds.Meeting_Name(id).Lazy(&t.Name)
	ds.Meeting_StructureLevelListOfSpeakersIDs(id).Lazy(&t.StructureLevelListOfSpeakersIDs)
	ds.Meeting_AssignmentPollDefaultMethod(id).Lazy(&t.AssignmentPollDefaultMethod)
	ds.Meeting_FontBoldID(id).Lazy(&t.FontBoldID)
	ds.Meeting_ListOfSpeakersEnablePointOfOrderCategories(id).Lazy(&t.ListOfSpeakersEnablePointOfOrderCategories)
	ds.Meeting_LogoPdfFooterRID(id).Lazy(&t.LogoPdfFooterRID)
	ds.Meeting_MotionPollBallotPaperSelection(id).Lazy(&t.MotionPollBallotPaperSelection)
	ds.Meeting_MotionPollDefaultOnehundredPercentBase(id).Lazy(&t.MotionPollDefaultOnehundredPercentBase)
	ds.Meeting_UsersPdfWlanEncryption(id).Lazy(&t.UsersPdfWlanEncryption)
	ds.Meeting_ListOfSpeakersAmountNextOnProjector(id).Lazy(&t.ListOfSpeakersAmountNextOnProjector)
	ds.Meeting_ListOfSpeakersCanSetContributionSelf(id).Lazy(&t.ListOfSpeakersCanSetContributionSelf)
	ds.Meeting_ListOfSpeakersDefaultStructureLevelTime(id).Lazy(&t.ListOfSpeakersDefaultStructureLevelTime)
	ds.Meeting_MotionsExportFollowRecommendation(id).Lazy(&t.MotionsExportFollowRecommendation)
	ds.Meeting_AllProjectionIDs(id).Lazy(&t.AllProjectionIDs)
	ds.Meeting_EnableAnonymous(id).Lazy(&t.EnableAnonymous)
	ds.Meeting_LogoPdfBallotPaperID(id).Lazy(&t.LogoPdfBallotPaperID)
	ds.Meeting_PollCountdownID(id).Lazy(&t.PollCountdownID)
	ds.Meeting_UserIDs(id).Lazy(&t.UserIDs)
	ds.Meeting_MotionCommentIDs(id).Lazy(&t.MotionCommentIDs)
	ds.Meeting_UsersEnableVoteWeight(id).Lazy(&t.UsersEnableVoteWeight)
	ds.Meeting_DefaultGroupID(id).Lazy(&t.DefaultGroupID)
	ds.Meeting_MeetingMediafileIDs(id).Lazy(&t.MeetingMediafileIDs)
	ds.Meeting_UsersEmailSubject(id).Lazy(&t.UsersEmailSubject)
	ds.Meeting_JitsiRoomName(id).Lazy(&t.JitsiRoomName)
	ds.Meeting_MotionWorkflowIDs(id).Lazy(&t.MotionWorkflowIDs)
	ds.Meeting_MotionsEnableReasonOnProjector(id).Lazy(&t.MotionsEnableReasonOnProjector)
	ds.Meeting_OptionIDs(id).Lazy(&t.OptionIDs)
	ds.Meeting_OrganizationTagIDs(id).Lazy(&t.OrganizationTagIDs)
	ds.Meeting_DefaultProjectorMessageIDs(id).Lazy(&t.DefaultProjectorMessageIDs)
	ds.Meeting_ExportPdfPageMarginRight(id).Lazy(&t.ExportPdfPageMarginRight)
	ds.Meeting_AssignmentPollSortPollResultByVotes(id).Lazy(&t.AssignmentPollSortPollResultByVotes)
	ds.Meeting_ConferenceStreamUrl(id).Lazy(&t.ConferenceStreamUrl)
	ds.Meeting_MeetingUserIDs(id).Lazy(&t.MeetingUserIDs)
	ds.Meeting_VoteIDs(id).Lazy(&t.VoteIDs)
	ds.Meeting_AgendaItemIDs(id).Lazy(&t.AgendaItemIDs)
	ds.Meeting_ApplauseEnable(id).Lazy(&t.ApplauseEnable)
	ds.Meeting_DefaultProjectorMotionBlockIDs(id).Lazy(&t.DefaultProjectorMotionBlockIDs)
	ds.Meeting_PersonalNoteIDs(id).Lazy(&t.PersonalNoteIDs)
	ds.Meeting_PollDefaultMethod(id).Lazy(&t.PollDefaultMethod)
	ds.Meeting_PollIDs(id).Lazy(&t.PollIDs)
	ds.Meeting_ApplauseMaxAmount(id).Lazy(&t.ApplauseMaxAmount)
	ds.Meeting_ConferenceOpenVideo(id).Lazy(&t.ConferenceOpenVideo)
	ds.Meeting_MediafileIDs(id).Lazy(&t.MediafileIDs)
	ds.Meeting_UsersPdfWelcometitle(id).Lazy(&t.UsersPdfWelcometitle)
	ds.Meeting_ExportPdfLineHeight(id).Lazy(&t.ExportPdfLineHeight)
	ds.Meeting_MotionWorkingGroupSpeakerIDs(id).Lazy(&t.MotionWorkingGroupSpeakerIDs)
	ds.Meeting_MotionsEnableSideboxOnProjector(id).Lazy(&t.MotionsEnableSideboxOnProjector)
	ds.Meeting_MotionsExportPreamble(id).Lazy(&t.MotionsExportPreamble)
	ds.Meeting_AgendaEnableNumbering(id).Lazy(&t.AgendaEnableNumbering)
	ds.Meeting_AssignmentPollBallotPaperSelection(id).Lazy(&t.AssignmentPollBallotPaperSelection)
	ds.Meeting_FontRegularID(id).Lazy(&t.FontRegularID)
	ds.Meeting_JitsiDomain(id).Lazy(&t.JitsiDomain)
	ds.Meeting_MotionsAmendmentsInMainList(id).Lazy(&t.MotionsAmendmentsInMainList)
	ds.Meeting_MotionsNumberWithBlank(id).Lazy(&t.MotionsNumberWithBlank)
	ds.Meeting_AssignmentPollBallotPaperNumber(id).Lazy(&t.AssignmentPollBallotPaperNumber)
	ds.Meeting_PollDefaultGroupIDs(id).Lazy(&t.PollDefaultGroupIDs)
	ds.Meeting_SpeakerIDs(id).Lazy(&t.SpeakerIDs)
	ds.Meeting_PollBallotPaperSelection(id).Lazy(&t.PollBallotPaperSelection)
	ds.Meeting_PresentUserIDs(id).Lazy(&t.PresentUserIDs)
	ds.Meeting_AssignmentPollDefaultBackend(id).Lazy(&t.AssignmentPollDefaultBackend)
	ds.Meeting_FontProjectorH1ID(id).Lazy(&t.FontProjectorH1ID)
	ds.Meeting_ListOfSpeakersCanCreatePointOfOrderForOthers(id).Lazy(&t.ListOfSpeakersCanCreatePointOfOrderForOthers)
	ds.Meeting_MotionsAmendmentsMultipleParagraphs(id).Lazy(&t.MotionsAmendmentsMultipleParagraphs)
	ds.Meeting_MotionsBlockSlideColumns(id).Lazy(&t.MotionsBlockSlideColumns)
	ds.Meeting_MotionsNumberType(id).Lazy(&t.MotionsNumberType)
	ds.Meeting_MotionIDs(id).Lazy(&t.MotionIDs)
	ds.Meeting_ImportedAt(id).Lazy(&t.ImportedAt)
	ds.Meeting_MotionsCreateEnableAdditionalSubmitterText(id).Lazy(&t.MotionsCreateEnableAdditionalSubmitterText)
	ds.Meeting_MotionsShowReferringMotions(id).Lazy(&t.MotionsShowReferringMotions)
	ds.Meeting_FontBoldItalicID(id).Lazy(&t.FontBoldItalicID)
	ds.Meeting_LockedFromInside(id).Lazy(&t.LockedFromInside)
	ds.Meeting_ApplauseTimeout(id).Lazy(&t.ApplauseTimeout)
	ds.Meeting_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.Meeting_IsArchivedInOrganizationID(id).Lazy(&t.IsArchivedInOrganizationID)
	ds.Meeting_Location(id).Lazy(&t.Location)
	ds.Meeting_PollDefaultOnehundredPercentBase(id).Lazy(&t.PollDefaultOnehundredPercentBase)
	ds.Meeting_UsersEnableVoteDelegations(id).Lazy(&t.UsersEnableVoteDelegations)
	ds.Meeting_ListOfSpeakersCountdownID(id).Lazy(&t.ListOfSpeakersCountdownID)
	ds.Meeting_MotionsAmendmentsEnabled(id).Lazy(&t.MotionsAmendmentsEnabled)
	ds.Meeting_MotionsRecommendationTextMode(id).Lazy(&t.MotionsRecommendationTextMode)
	ds.Meeting_DefaultProjectorListOfSpeakersIDs(id).Lazy(&t.DefaultProjectorListOfSpeakersIDs)
	ds.Meeting_DefaultProjectorTopicIDs(id).Lazy(&t.DefaultProjectorTopicIDs)
	ds.Meeting_ListOfSpeakersSpeakerNoteForEveryone(id).Lazy(&t.ListOfSpeakersSpeakerNoteForEveryone)
	ds.Meeting_MotionBlockIDs(id).Lazy(&t.MotionBlockIDs)
	ds.Meeting_LogoPdfFooterLID(id).Lazy(&t.LogoPdfFooterLID)
	ds.Meeting_MotionsAmendmentsTextMode(id).Lazy(&t.MotionsAmendmentsTextMode)
	ds.Meeting_PointOfOrderCategoryIDs(id).Lazy(&t.PointOfOrderCategoryIDs)
	ds.Meeting_UsersPdfWelcometext(id).Lazy(&t.UsersPdfWelcometext)
	ds.Meeting_WelcomeTitle(id).Lazy(&t.WelcomeTitle)
	ds.Meeting_AgendaShowInternalItemsOnProjector(id).Lazy(&t.AgendaShowInternalItemsOnProjector)
	ds.Meeting_ExportCsvSeparator(id).Lazy(&t.ExportCsvSeparator)
	ds.Meeting_ListOfSpeakersEnableInterposedQuestion(id).Lazy(&t.ListOfSpeakersEnableInterposedQuestion)
	ds.Meeting_LogoPdfHeaderLID(id).Lazy(&t.LogoPdfHeaderLID)
	ds.Meeting_ProjectorCountdownIDs(id).Lazy(&t.ProjectorCountdownIDs)
	ds.Meeting_UsersForbidDelegatorToVote(id).Lazy(&t.UsersForbidDelegatorToVote)
	ds.Meeting_ApplauseShowLevel(id).Lazy(&t.ApplauseShowLevel)
	ds.Meeting_ListOfSpeakersEnableProContraSpeech(id).Lazy(&t.ListOfSpeakersEnableProContraSpeech)
	ds.Meeting_ListOfSpeakersHideContributionCount(id).Lazy(&t.ListOfSpeakersHideContributionCount)
	ds.Meeting_ListOfSpeakersPresentUsersOnly(id).Lazy(&t.ListOfSpeakersPresentUsersOnly)
	ds.Meeting_LogoProjectorMainID(id).Lazy(&t.LogoProjectorMainID)
	ds.Meeting_MotionStateIDs(id).Lazy(&t.MotionStateIDs)
	ds.Meeting_ListOfSpeakersInitiallyClosed(id).Lazy(&t.ListOfSpeakersInitiallyClosed)
	ds.Meeting_TopicIDs(id).Lazy(&t.TopicIDs)
	ds.Meeting_DefaultProjectorCurrentListOfSpeakersIDs(id).Lazy(&t.DefaultProjectorCurrentListOfSpeakersIDs)
	ds.Meeting_DefaultProjectorPollIDs(id).Lazy(&t.DefaultProjectorPollIDs)
	ds.Meeting_Language(id).Lazy(&t.Language)
	ds.Meeting_ConferenceAutoConnectNextSpeakers(id).Lazy(&t.ConferenceAutoConnectNextSpeakers)
	ds.Meeting_ExportPdfPagenumberAlignment(id).Lazy(&t.ExportPdfPagenumberAlignment)
	ds.Meeting_FontProjectorH2ID(id).Lazy(&t.FontProjectorH2ID)
	ds.Meeting_MotionSubmitterIDs(id).Lazy(&t.MotionSubmitterIDs)
	ds.Meeting_StartTime(id).Lazy(&t.StartTime)
	ds.Meeting_UsersEnablePresenceView(id).Lazy(&t.UsersEnablePresenceView)
	ds.Meeting_AssignmentsExportPreamble(id).Lazy(&t.AssignmentsExportPreamble)
	ds.Meeting_ExportPdfFontsize(id).Lazy(&t.ExportPdfFontsize)
	ds.Meeting_ExportPdfPageMarginTop(id).Lazy(&t.ExportPdfPageMarginTop)
	ds.Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(id).Lazy(&t.ListOfSpeakersShowAmountOfSpeakersOnSlide)
	ds.Meeting_MotionsAmendmentsPrefix(id).Lazy(&t.MotionsAmendmentsPrefix)
	ds.Meeting_MotionsEnableRecommendationOnProjector(id).Lazy(&t.MotionsEnableRecommendationOnProjector)
	ds.Meeting_ExportPdfPageMarginBottom(id).Lazy(&t.ExportPdfPageMarginBottom)
	ds.Meeting_ListOfSpeakersShowFirstContribution(id).Lazy(&t.ListOfSpeakersShowFirstContribution)
	ds.Meeting_MotionsNumberMinDigits(id).Lazy(&t.MotionsNumberMinDigits)
	ds.Meeting_DefaultProjectorAssignmentIDs(id).Lazy(&t.DefaultProjectorAssignmentIDs)
	ds.Meeting_ID(id).Lazy(&t.ID)
	ds.Meeting_MotionsReasonRequired(id).Lazy(&t.MotionsReasonRequired)
	ds.Meeting_TemplateForOrganizationID(id).Lazy(&t.TemplateForOrganizationID)
	ds.Meeting_MotionPollBallotPaperNumber(id).Lazy(&t.MotionPollBallotPaperNumber)
	ds.Meeting_FontChyronSpeakerNameID(id).Lazy(&t.FontChyronSpeakerNameID)
	ds.Meeting_PollCandidateListIDs(id).Lazy(&t.PollCandidateListIDs)
	ds.Meeting_MotionsEnableWorkingGroupSpeaker(id).Lazy(&t.MotionsEnableWorkingGroupSpeaker)
	ds.Meeting_MotionsHideMetadataBackground(id).Lazy(&t.MotionsHideMetadataBackground)
	ds.Meeting_ApplauseMinAmount(id).Lazy(&t.ApplauseMinAmount)
	ds.Meeting_ApplauseParticleImageUrl(id).Lazy(&t.ApplauseParticleImageUrl)
	ds.Meeting_ChatGroupIDs(id).Lazy(&t.ChatGroupIDs)
	ds.Meeting_ConferenceEnableHelpdesk(id).Lazy(&t.ConferenceEnableHelpdesk)
	ds.Meeting_CustomTranslations(id).Lazy(&t.CustomTranslations)
	ds.Meeting_MotionCategoryIDs(id).Lazy(&t.MotionCategoryIDs)
	ds.Meeting_ConferenceAutoConnect(id).Lazy(&t.ConferenceAutoConnect)
	ds.Meeting_ExternalID(id).Lazy(&t.ExternalID)
	ds.Meeting_MotionsLineLength(id).Lazy(&t.MotionsLineLength)
	ds.Meeting_PollDefaultType(id).Lazy(&t.PollDefaultType)
	ds.Meeting_AgendaShowSubtitles(id).Lazy(&t.AgendaShowSubtitles)
	ds.Meeting_CommitteeID(id).Lazy(&t.CommitteeID)
	ds.Meeting_FontMonospaceID(id).Lazy(&t.FontMonospaceID)
	ds.Meeting_MotionsDefaultAmendmentWorkflowID(id).Lazy(&t.MotionsDefaultAmendmentWorkflowID)
	ds.Meeting_MotionsDefaultLineNumbering(id).Lazy(&t.MotionsDefaultLineNumbering)
	ds.Meeting_AssignmentsExportTitle(id).Lazy(&t.AssignmentsExportTitle)
	ds.Meeting_DefaultProjectorAssignmentPollIDs(id).Lazy(&t.DefaultProjectorAssignmentPollIDs)
	ds.Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(id).Lazy(&t.ListOfSpeakersEnablePointOfOrderSpeakers)
	ds.Meeting_ExportCsvEncoding(id).Lazy(&t.ExportCsvEncoding)
	ds.Meeting_PollSortPollResultByVotes(id).Lazy(&t.PollSortPollResultByVotes)
	ds.Meeting_TagIDs(id).Lazy(&t.TagIDs)
	ds.Meeting_ConferenceLosRestriction(id).Lazy(&t.ConferenceLosRestriction)
	ds.Meeting_DefaultProjectorMotionPollIDs(id).Lazy(&t.DefaultProjectorMotionPollIDs)
	ds.Meeting_AdminGroupID(id).Lazy(&t.AdminGroupID)
	ds.Meeting_DefaultProjectorMediafileIDs(id).Lazy(&t.DefaultProjectorMediafileIDs)
	ds.Meeting_MotionsExportSubmitterRecommendation(id).Lazy(&t.MotionsExportSubmitterRecommendation)
	ds.Meeting_UsersForbidDelegatorAsSupporter(id).Lazy(&t.UsersForbidDelegatorAsSupporter)
	ds.Meeting_MotionPollDefaultGroupIDs(id).Lazy(&t.MotionPollDefaultGroupIDs)
	ds.Meeting_MotionPollDefaultMethod(id).Lazy(&t.MotionPollDefaultMethod)
	ds.Meeting_MotionsDefaultSorting(id).Lazy(&t.MotionsDefaultSorting)
	ds.Meeting_PollDefaultBackend(id).Lazy(&t.PollDefaultBackend)
	ds.Meeting_AssignmentIDs(id).Lazy(&t.AssignmentIDs)
	ds.Meeting_ProjectorMessageIDs(id).Lazy(&t.ProjectorMessageIDs)
	ds.Meeting_AgendaItemCreation(id).Lazy(&t.AgendaItemCreation)
	ds.Meeting_DefaultProjectorAmendmentIDs(id).Lazy(&t.DefaultProjectorAmendmentIDs)
	ds.Meeting_ForwardedMotionIDs(id).Lazy(&t.ForwardedMotionIDs)
	ds.Meeting_MotionsShowSequentialNumber(id).Lazy(&t.MotionsShowSequentialNumber)
	ds.Meeting_ProjectorCountdownWarningTime(id).Lazy(&t.ProjectorCountdownWarningTime)
}

func (t *Meeting) preload(ds *Fetch, id int) {
	ds.Meeting_ReferenceProjectorID(id).Preload()
	ds.Meeting_AssignmentPollDefaultOnehundredPercentBase(id).Preload()
	ds.Meeting_DefaultProjectorMotionIDs(id).Preload()
	ds.Meeting_PollCoupleCountdown(id).Preload()
	ds.Meeting_ConferenceOpenMicrophone(id).Preload()
	ds.Meeting_LogoPdfHeaderRID(id).Preload()
	ds.Meeting_UsersEmailBody(id).Preload()
	ds.Meeting_ChatMessageIDs(id).Preload()
	ds.Meeting_AssignmentPollDefaultGroupIDs(id).Preload()
	ds.Meeting_AssignmentPollDefaultType(id).Preload()
	ds.Meeting_MotionCommentSectionIDs(id).Preload()
	ds.Meeting_MotionsPreamble(id).Preload()
	ds.Meeting_UsersPdfWlanSsid(id).Preload()
	ds.Meeting_WelcomeText(id).Preload()
	ds.Meeting_PollBallotPaperNumber(id).Preload()
	ds.Meeting_TopicPollDefaultGroupIDs(id).Preload()
	ds.Meeting_UsersEmailReplyto(id).Preload()
	ds.Meeting_ApplauseType(id).Preload()
	ds.Meeting_DefaultProjectorAgendaItemListIDs(id).Preload()
	ds.Meeting_ListOfSpeakersAllowMultipleSpeakers(id).Preload()
	ds.Meeting_MotionsEnableEditor(id).Preload()
	ds.Meeting_ProjectorCountdownDefaultTime(id).Preload()
	ds.Meeting_ConferenceShow(id).Preload()
	ds.Meeting_ExportPdfPageMarginLeft(id).Preload()
	ds.Meeting_GroupIDs(id).Preload()
	ds.Meeting_AgendaNewItemsDefaultVisibility(id).Preload()
	ds.Meeting_JitsiRoomPassword(id).Preload()
	ds.Meeting_ListOfSpeakersIDs(id).Preload()
	ds.Meeting_LogoWebHeaderID(id).Preload()
	ds.Meeting_MotionsRecommendationsBy(id).Preload()
	ds.Meeting_Description(id).Preload()
	ds.Meeting_IsActiveInOrganizationID(id).Preload()
	ds.Meeting_StructureLevelIDs(id).Preload()
	ds.Meeting_UsersAllowSelfSetPresent(id).Preload()
	ds.Meeting_AgendaNumeralSystem(id).Preload()
	ds.Meeting_MotionsExportTitle(id).Preload()
	ds.Meeting_MotionsSupportersMinAmount(id).Preload()
	ds.Meeting_UsersForbidDelegatorAsSubmitter(id).Preload()
	ds.Meeting_UsersPdfWlanPassword(id).Preload()
	ds.Meeting_AnonymousGroupID(id).Preload()
	ds.Meeting_ConferenceStreamPosterUrl(id).Preload()
	ds.Meeting_FontItalicID(id).Preload()
	ds.Meeting_PollCandidateIDs(id).Preload()
	ds.Meeting_ProjectorIDs(id).Preload()
	ds.Meeting_AssignmentCandidateIDs(id).Preload()
	ds.Meeting_DefaultProjectorCountdownIDs(id).Preload()
	ds.Meeting_ListOfSpeakersAmountLastOnProjector(id).Preload()
	ds.Meeting_ListOfSpeakersClosingDisablesPointOfOrder(id).Preload()
	ds.Meeting_MotionsAmendmentsOfAmendments(id).Preload()
	ds.Meeting_MotionsEnableTextOnProjector(id).Preload()
	ds.Meeting_AgendaNumberPrefix(id).Preload()
	ds.Meeting_MotionChangeRecommendationIDs(id).Preload()
	ds.Meeting_UsersEmailSender(id).Preload()
	ds.Meeting_ExportPdfPagesize(id).Preload()
	ds.Meeting_LogoProjectorHeaderID(id).Preload()
	ds.Meeting_DefaultMeetingForCommitteeID(id).Preload()
	ds.Meeting_EndTime(id).Preload()
	ds.Meeting_ListOfSpeakersInterventionTime(id).Preload()
	ds.Meeting_MotionsDefaultWorkflowID(id).Preload()
	ds.Meeting_UsersForbidDelegatorInListOfSpeakers(id).Preload()
	ds.Meeting_AssignmentPollAddCandidatesToListOfSpeakers(id).Preload()
	ds.Meeting_AssignmentPollEnableMaxVotesPerOption(id).Preload()
	ds.Meeting_ListOfSpeakersCoupleCountdown(id).Preload()
	ds.Meeting_AgendaShowTopicNavigationOnDetailView(id).Preload()
	ds.Meeting_MotionPollDefaultBackend(id).Preload()
	ds.Meeting_MotionPollDefaultType(id).Preload()
	ds.Meeting_MotionEditorIDs(id).Preload()
	ds.Meeting_Name(id).Preload()
	ds.Meeting_StructureLevelListOfSpeakersIDs(id).Preload()
	ds.Meeting_AssignmentPollDefaultMethod(id).Preload()
	ds.Meeting_FontBoldID(id).Preload()
	ds.Meeting_ListOfSpeakersEnablePointOfOrderCategories(id).Preload()
	ds.Meeting_LogoPdfFooterRID(id).Preload()
	ds.Meeting_MotionPollBallotPaperSelection(id).Preload()
	ds.Meeting_MotionPollDefaultOnehundredPercentBase(id).Preload()
	ds.Meeting_UsersPdfWlanEncryption(id).Preload()
	ds.Meeting_ListOfSpeakersAmountNextOnProjector(id).Preload()
	ds.Meeting_ListOfSpeakersCanSetContributionSelf(id).Preload()
	ds.Meeting_ListOfSpeakersDefaultStructureLevelTime(id).Preload()
	ds.Meeting_MotionsExportFollowRecommendation(id).Preload()
	ds.Meeting_AllProjectionIDs(id).Preload()
	ds.Meeting_EnableAnonymous(id).Preload()
	ds.Meeting_LogoPdfBallotPaperID(id).Preload()
	ds.Meeting_PollCountdownID(id).Preload()
	ds.Meeting_UserIDs(id).Preload()
	ds.Meeting_MotionCommentIDs(id).Preload()
	ds.Meeting_UsersEnableVoteWeight(id).Preload()
	ds.Meeting_DefaultGroupID(id).Preload()
	ds.Meeting_MeetingMediafileIDs(id).Preload()
	ds.Meeting_UsersEmailSubject(id).Preload()
	ds.Meeting_JitsiRoomName(id).Preload()
	ds.Meeting_MotionWorkflowIDs(id).Preload()
	ds.Meeting_MotionsEnableReasonOnProjector(id).Preload()
	ds.Meeting_OptionIDs(id).Preload()
	ds.Meeting_OrganizationTagIDs(id).Preload()
	ds.Meeting_DefaultProjectorMessageIDs(id).Preload()
	ds.Meeting_ExportPdfPageMarginRight(id).Preload()
	ds.Meeting_AssignmentPollSortPollResultByVotes(id).Preload()
	ds.Meeting_ConferenceStreamUrl(id).Preload()
	ds.Meeting_MeetingUserIDs(id).Preload()
	ds.Meeting_VoteIDs(id).Preload()
	ds.Meeting_AgendaItemIDs(id).Preload()
	ds.Meeting_ApplauseEnable(id).Preload()
	ds.Meeting_DefaultProjectorMotionBlockIDs(id).Preload()
	ds.Meeting_PersonalNoteIDs(id).Preload()
	ds.Meeting_PollDefaultMethod(id).Preload()
	ds.Meeting_PollIDs(id).Preload()
	ds.Meeting_ApplauseMaxAmount(id).Preload()
	ds.Meeting_ConferenceOpenVideo(id).Preload()
	ds.Meeting_MediafileIDs(id).Preload()
	ds.Meeting_UsersPdfWelcometitle(id).Preload()
	ds.Meeting_ExportPdfLineHeight(id).Preload()
	ds.Meeting_MotionWorkingGroupSpeakerIDs(id).Preload()
	ds.Meeting_MotionsEnableSideboxOnProjector(id).Preload()
	ds.Meeting_MotionsExportPreamble(id).Preload()
	ds.Meeting_AgendaEnableNumbering(id).Preload()
	ds.Meeting_AssignmentPollBallotPaperSelection(id).Preload()
	ds.Meeting_FontRegularID(id).Preload()
	ds.Meeting_JitsiDomain(id).Preload()
	ds.Meeting_MotionsAmendmentsInMainList(id).Preload()
	ds.Meeting_MotionsNumberWithBlank(id).Preload()
	ds.Meeting_AssignmentPollBallotPaperNumber(id).Preload()
	ds.Meeting_PollDefaultGroupIDs(id).Preload()
	ds.Meeting_SpeakerIDs(id).Preload()
	ds.Meeting_PollBallotPaperSelection(id).Preload()
	ds.Meeting_PresentUserIDs(id).Preload()
	ds.Meeting_AssignmentPollDefaultBackend(id).Preload()
	ds.Meeting_FontProjectorH1ID(id).Preload()
	ds.Meeting_ListOfSpeakersCanCreatePointOfOrderForOthers(id).Preload()
	ds.Meeting_MotionsAmendmentsMultipleParagraphs(id).Preload()
	ds.Meeting_MotionsBlockSlideColumns(id).Preload()
	ds.Meeting_MotionsNumberType(id).Preload()
	ds.Meeting_MotionIDs(id).Preload()
	ds.Meeting_ImportedAt(id).Preload()
	ds.Meeting_MotionsCreateEnableAdditionalSubmitterText(id).Preload()
	ds.Meeting_MotionsShowReferringMotions(id).Preload()
	ds.Meeting_FontBoldItalicID(id).Preload()
	ds.Meeting_LockedFromInside(id).Preload()
	ds.Meeting_ApplauseTimeout(id).Preload()
	ds.Meeting_ProjectionIDs(id).Preload()
	ds.Meeting_IsArchivedInOrganizationID(id).Preload()
	ds.Meeting_Location(id).Preload()
	ds.Meeting_PollDefaultOnehundredPercentBase(id).Preload()
	ds.Meeting_UsersEnableVoteDelegations(id).Preload()
	ds.Meeting_ListOfSpeakersCountdownID(id).Preload()
	ds.Meeting_MotionsAmendmentsEnabled(id).Preload()
	ds.Meeting_MotionsRecommendationTextMode(id).Preload()
	ds.Meeting_DefaultProjectorListOfSpeakersIDs(id).Preload()
	ds.Meeting_DefaultProjectorTopicIDs(id).Preload()
	ds.Meeting_ListOfSpeakersSpeakerNoteForEveryone(id).Preload()
	ds.Meeting_MotionBlockIDs(id).Preload()
	ds.Meeting_LogoPdfFooterLID(id).Preload()
	ds.Meeting_MotionsAmendmentsTextMode(id).Preload()
	ds.Meeting_PointOfOrderCategoryIDs(id).Preload()
	ds.Meeting_UsersPdfWelcometext(id).Preload()
	ds.Meeting_WelcomeTitle(id).Preload()
	ds.Meeting_AgendaShowInternalItemsOnProjector(id).Preload()
	ds.Meeting_ExportCsvSeparator(id).Preload()
	ds.Meeting_ListOfSpeakersEnableInterposedQuestion(id).Preload()
	ds.Meeting_LogoPdfHeaderLID(id).Preload()
	ds.Meeting_ProjectorCountdownIDs(id).Preload()
	ds.Meeting_UsersForbidDelegatorToVote(id).Preload()
	ds.Meeting_ApplauseShowLevel(id).Preload()
	ds.Meeting_ListOfSpeakersEnableProContraSpeech(id).Preload()
	ds.Meeting_ListOfSpeakersHideContributionCount(id).Preload()
	ds.Meeting_ListOfSpeakersPresentUsersOnly(id).Preload()
	ds.Meeting_LogoProjectorMainID(id).Preload()
	ds.Meeting_MotionStateIDs(id).Preload()
	ds.Meeting_ListOfSpeakersInitiallyClosed(id).Preload()
	ds.Meeting_TopicIDs(id).Preload()
	ds.Meeting_DefaultProjectorCurrentListOfSpeakersIDs(id).Preload()
	ds.Meeting_DefaultProjectorPollIDs(id).Preload()
	ds.Meeting_Language(id).Preload()
	ds.Meeting_ConferenceAutoConnectNextSpeakers(id).Preload()
	ds.Meeting_ExportPdfPagenumberAlignment(id).Preload()
	ds.Meeting_FontProjectorH2ID(id).Preload()
	ds.Meeting_MotionSubmitterIDs(id).Preload()
	ds.Meeting_StartTime(id).Preload()
	ds.Meeting_UsersEnablePresenceView(id).Preload()
	ds.Meeting_AssignmentsExportPreamble(id).Preload()
	ds.Meeting_ExportPdfFontsize(id).Preload()
	ds.Meeting_ExportPdfPageMarginTop(id).Preload()
	ds.Meeting_ListOfSpeakersShowAmountOfSpeakersOnSlide(id).Preload()
	ds.Meeting_MotionsAmendmentsPrefix(id).Preload()
	ds.Meeting_MotionsEnableRecommendationOnProjector(id).Preload()
	ds.Meeting_ExportPdfPageMarginBottom(id).Preload()
	ds.Meeting_ListOfSpeakersShowFirstContribution(id).Preload()
	ds.Meeting_MotionsNumberMinDigits(id).Preload()
	ds.Meeting_DefaultProjectorAssignmentIDs(id).Preload()
	ds.Meeting_ID(id).Preload()
	ds.Meeting_MotionsReasonRequired(id).Preload()
	ds.Meeting_TemplateForOrganizationID(id).Preload()
	ds.Meeting_MotionPollBallotPaperNumber(id).Preload()
	ds.Meeting_FontChyronSpeakerNameID(id).Preload()
	ds.Meeting_PollCandidateListIDs(id).Preload()
	ds.Meeting_MotionsEnableWorkingGroupSpeaker(id).Preload()
	ds.Meeting_MotionsHideMetadataBackground(id).Preload()
	ds.Meeting_ApplauseMinAmount(id).Preload()
	ds.Meeting_ApplauseParticleImageUrl(id).Preload()
	ds.Meeting_ChatGroupIDs(id).Preload()
	ds.Meeting_ConferenceEnableHelpdesk(id).Preload()
	ds.Meeting_CustomTranslations(id).Preload()
	ds.Meeting_MotionCategoryIDs(id).Preload()
	ds.Meeting_ConferenceAutoConnect(id).Preload()
	ds.Meeting_ExternalID(id).Preload()
	ds.Meeting_MotionsLineLength(id).Preload()
	ds.Meeting_PollDefaultType(id).Preload()
	ds.Meeting_AgendaShowSubtitles(id).Preload()
	ds.Meeting_CommitteeID(id).Preload()
	ds.Meeting_FontMonospaceID(id).Preload()
	ds.Meeting_MotionsDefaultAmendmentWorkflowID(id).Preload()
	ds.Meeting_MotionsDefaultLineNumbering(id).Preload()
	ds.Meeting_AssignmentsExportTitle(id).Preload()
	ds.Meeting_DefaultProjectorAssignmentPollIDs(id).Preload()
	ds.Meeting_ListOfSpeakersEnablePointOfOrderSpeakers(id).Preload()
	ds.Meeting_ExportCsvEncoding(id).Preload()
	ds.Meeting_PollSortPollResultByVotes(id).Preload()
	ds.Meeting_TagIDs(id).Preload()
	ds.Meeting_ConferenceLosRestriction(id).Preload()
	ds.Meeting_DefaultProjectorMotionPollIDs(id).Preload()
	ds.Meeting_AdminGroupID(id).Preload()
	ds.Meeting_DefaultProjectorMediafileIDs(id).Preload()
	ds.Meeting_MotionsExportSubmitterRecommendation(id).Preload()
	ds.Meeting_UsersForbidDelegatorAsSupporter(id).Preload()
	ds.Meeting_MotionPollDefaultGroupIDs(id).Preload()
	ds.Meeting_MotionPollDefaultMethod(id).Preload()
	ds.Meeting_MotionsDefaultSorting(id).Preload()
	ds.Meeting_PollDefaultBackend(id).Preload()
	ds.Meeting_AssignmentIDs(id).Preload()
	ds.Meeting_ProjectorMessageIDs(id).Preload()
	ds.Meeting_AgendaItemCreation(id).Preload()
	ds.Meeting_DefaultProjectorAmendmentIDs(id).Preload()
	ds.Meeting_ForwardedMotionIDs(id).Preload()
	ds.Meeting_MotionsShowSequentialNumber(id).Preload()
	ds.Meeting_ProjectorCountdownWarningTime(id).Preload()
}

func (r *Fetch) Meeting(id int) *ValueCollection[Meeting, *Meeting] {
	return &ValueCollection[Meeting, *Meeting]{
		id:    id,
		fetch: r,
	}
}

// Theme has all fields from theme.
type Theme struct {
	PrimaryA200            string
	Accent50               string
	Accent700              string
	Primary600             string
	Headbar                string
	ID                     int
	OrganizationID         int
	Primary300             string
	Primary900             string
	Abstain                string
	Accent400              string
	Accent500              string
	Warn100                string
	WarnA100               string
	Yes                    string
	Warn50                 string
	Warn600                string
	WarnA700               string
	Accent100              string
	Primary100             string
	PrimaryA100            string
	Warn300                string
	Warn500                string
	Warn800                string
	Accent200              string
	AccentA700             string
	ThemeForOrganizationID Maybe[int]
	PrimaryA700            string
	Warn400                string
	WarnA200               string
	Name                   string
	Primary50              string
	PrimaryA400            string
	Warn900                string
	AccentA400             string
	No                     string
	Primary200             string
	Primary400             string
	Primary500             string
	Primary800             string
	Warn700                string
	Accent300              string
	Accent600              string
	Accent800              string
	Primary700             string
	Warn200                string
	WarnA400               string
	Accent900              string
	AccentA100             string
	AccentA200             string
}

func (t *Theme) lazy(ds *Fetch, id int) {
	ds.Theme_PrimaryA200(id).Lazy(&t.PrimaryA200)
	ds.Theme_Accent50(id).Lazy(&t.Accent50)
	ds.Theme_Accent700(id).Lazy(&t.Accent700)
	ds.Theme_Primary600(id).Lazy(&t.Primary600)
	ds.Theme_Headbar(id).Lazy(&t.Headbar)
	ds.Theme_ID(id).Lazy(&t.ID)
	ds.Theme_OrganizationID(id).Lazy(&t.OrganizationID)
	ds.Theme_Primary300(id).Lazy(&t.Primary300)
	ds.Theme_Primary900(id).Lazy(&t.Primary900)
	ds.Theme_Abstain(id).Lazy(&t.Abstain)
	ds.Theme_Accent400(id).Lazy(&t.Accent400)
	ds.Theme_Accent500(id).Lazy(&t.Accent500)
	ds.Theme_Warn100(id).Lazy(&t.Warn100)
	ds.Theme_WarnA100(id).Lazy(&t.WarnA100)
	ds.Theme_Yes(id).Lazy(&t.Yes)
	ds.Theme_Warn50(id).Lazy(&t.Warn50)
	ds.Theme_Warn600(id).Lazy(&t.Warn600)
	ds.Theme_WarnA700(id).Lazy(&t.WarnA700)
	ds.Theme_Accent100(id).Lazy(&t.Accent100)
	ds.Theme_Primary100(id).Lazy(&t.Primary100)
	ds.Theme_PrimaryA100(id).Lazy(&t.PrimaryA100)
	ds.Theme_Warn300(id).Lazy(&t.Warn300)
	ds.Theme_Warn500(id).Lazy(&t.Warn500)
	ds.Theme_Warn800(id).Lazy(&t.Warn800)
	ds.Theme_Accent200(id).Lazy(&t.Accent200)
	ds.Theme_AccentA700(id).Lazy(&t.AccentA700)
	ds.Theme_ThemeForOrganizationID(id).Lazy(&t.ThemeForOrganizationID)
	ds.Theme_PrimaryA700(id).Lazy(&t.PrimaryA700)
	ds.Theme_Warn400(id).Lazy(&t.Warn400)
	ds.Theme_WarnA200(id).Lazy(&t.WarnA200)
	ds.Theme_Name(id).Lazy(&t.Name)
	ds.Theme_Primary50(id).Lazy(&t.Primary50)
	ds.Theme_PrimaryA400(id).Lazy(&t.PrimaryA400)
	ds.Theme_Warn900(id).Lazy(&t.Warn900)
	ds.Theme_AccentA400(id).Lazy(&t.AccentA400)
	ds.Theme_No(id).Lazy(&t.No)
	ds.Theme_Primary200(id).Lazy(&t.Primary200)
	ds.Theme_Primary400(id).Lazy(&t.Primary400)
	ds.Theme_Primary500(id).Lazy(&t.Primary500)
	ds.Theme_Primary800(id).Lazy(&t.Primary800)
	ds.Theme_Warn700(id).Lazy(&t.Warn700)
	ds.Theme_Accent300(id).Lazy(&t.Accent300)
	ds.Theme_Accent600(id).Lazy(&t.Accent600)
	ds.Theme_Accent800(id).Lazy(&t.Accent800)
	ds.Theme_Primary700(id).Lazy(&t.Primary700)
	ds.Theme_Warn200(id).Lazy(&t.Warn200)
	ds.Theme_WarnA400(id).Lazy(&t.WarnA400)
	ds.Theme_Accent900(id).Lazy(&t.Accent900)
	ds.Theme_AccentA100(id).Lazy(&t.AccentA100)
	ds.Theme_AccentA200(id).Lazy(&t.AccentA200)
}

func (t *Theme) preload(ds *Fetch, id int) {
	ds.Theme_PrimaryA200(id).Preload()
	ds.Theme_Accent50(id).Preload()
	ds.Theme_Accent700(id).Preload()
	ds.Theme_Primary600(id).Preload()
	ds.Theme_Headbar(id).Preload()
	ds.Theme_ID(id).Preload()
	ds.Theme_OrganizationID(id).Preload()
	ds.Theme_Primary300(id).Preload()
	ds.Theme_Primary900(id).Preload()
	ds.Theme_Abstain(id).Preload()
	ds.Theme_Accent400(id).Preload()
	ds.Theme_Accent500(id).Preload()
	ds.Theme_Warn100(id).Preload()
	ds.Theme_WarnA100(id).Preload()
	ds.Theme_Yes(id).Preload()
	ds.Theme_Warn50(id).Preload()
	ds.Theme_Warn600(id).Preload()
	ds.Theme_WarnA700(id).Preload()
	ds.Theme_Accent100(id).Preload()
	ds.Theme_Primary100(id).Preload()
	ds.Theme_PrimaryA100(id).Preload()
	ds.Theme_Warn300(id).Preload()
	ds.Theme_Warn500(id).Preload()
	ds.Theme_Warn800(id).Preload()
	ds.Theme_Accent200(id).Preload()
	ds.Theme_AccentA700(id).Preload()
	ds.Theme_ThemeForOrganizationID(id).Preload()
	ds.Theme_PrimaryA700(id).Preload()
	ds.Theme_Warn400(id).Preload()
	ds.Theme_WarnA200(id).Preload()
	ds.Theme_Name(id).Preload()
	ds.Theme_Primary50(id).Preload()
	ds.Theme_PrimaryA400(id).Preload()
	ds.Theme_Warn900(id).Preload()
	ds.Theme_AccentA400(id).Preload()
	ds.Theme_No(id).Preload()
	ds.Theme_Primary200(id).Preload()
	ds.Theme_Primary400(id).Preload()
	ds.Theme_Primary500(id).Preload()
	ds.Theme_Primary800(id).Preload()
	ds.Theme_Warn700(id).Preload()
	ds.Theme_Accent300(id).Preload()
	ds.Theme_Accent600(id).Preload()
	ds.Theme_Accent800(id).Preload()
	ds.Theme_Primary700(id).Preload()
	ds.Theme_Warn200(id).Preload()
	ds.Theme_WarnA400(id).Preload()
	ds.Theme_Accent900(id).Preload()
	ds.Theme_AccentA100(id).Preload()
	ds.Theme_AccentA200(id).Preload()
}

func (r *Fetch) Theme(id int) *ValueCollection[Theme, *Theme] {
	return &ValueCollection[Theme, *Theme]{
		id:    id,
		fetch: r,
	}
}

// ListOfSpeakers has all fields from list_of_speakers.
type ListOfSpeakers struct {
	ID                              int
	ContentObjectID                 string
	MeetingID                       int
	ModeratorNotes                  string
	ProjectionIDs                   []int
	SequentialNumber                int
	SpeakerIDs                      []int
	StructureLevelListOfSpeakersIDs []int
	Closed                          bool
}

func (t *ListOfSpeakers) lazy(ds *Fetch, id int) {
	ds.ListOfSpeakers_ID(id).Lazy(&t.ID)
	ds.ListOfSpeakers_ContentObjectID(id).Lazy(&t.ContentObjectID)
	ds.ListOfSpeakers_MeetingID(id).Lazy(&t.MeetingID)
	ds.ListOfSpeakers_ModeratorNotes(id).Lazy(&t.ModeratorNotes)
	ds.ListOfSpeakers_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.ListOfSpeakers_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.ListOfSpeakers_SpeakerIDs(id).Lazy(&t.SpeakerIDs)
	ds.ListOfSpeakers_StructureLevelListOfSpeakersIDs(id).Lazy(&t.StructureLevelListOfSpeakersIDs)
	ds.ListOfSpeakers_Closed(id).Lazy(&t.Closed)
}

func (t *ListOfSpeakers) preload(ds *Fetch, id int) {
	ds.ListOfSpeakers_ID(id).Preload()
	ds.ListOfSpeakers_ContentObjectID(id).Preload()
	ds.ListOfSpeakers_MeetingID(id).Preload()
	ds.ListOfSpeakers_ModeratorNotes(id).Preload()
	ds.ListOfSpeakers_ProjectionIDs(id).Preload()
	ds.ListOfSpeakers_SequentialNumber(id).Preload()
	ds.ListOfSpeakers_SpeakerIDs(id).Preload()
	ds.ListOfSpeakers_StructureLevelListOfSpeakersIDs(id).Preload()
	ds.ListOfSpeakers_Closed(id).Preload()
}

func (r *Fetch) ListOfSpeakers(id int) *ValueCollection[ListOfSpeakers, *ListOfSpeakers] {
	return &ValueCollection[ListOfSpeakers, *ListOfSpeakers]{
		id:    id,
		fetch: r,
	}
}

// Poll has all fields from poll.
type Poll struct {
	MaxVotesAmount        int
	MaxVotesPerOption     int
	GlobalNo              bool
	CryptSignature        string
	EntitledGroupIDs      []int
	ProjectionIDs         []int
	IsPseudoanonymized    bool
	OptionIDs             []int
	MinVotesAmount        int
	Title                 string
	VoteCount             int
	Votesinvalid          string
	Votesvalid            string
	GlobalAbstain         bool
	MeetingID             int
	EntitledUsersAtStop   json.RawMessage
	GlobalYes             bool
	State                 string
	VotesSignature        string
	Votescast             string
	Backend               string
	ContentObjectID       string
	Type                  string
	CryptKey              string
	GlobalOptionID        Maybe[int]
	Pollmethod            string
	SequentialNumber      int
	VotedIDs              []int
	Description           string
	ID                    int
	OnehundredPercentBase string
	VotesRaw              string
}

func (t *Poll) lazy(ds *Fetch, id int) {
	ds.Poll_MaxVotesAmount(id).Lazy(&t.MaxVotesAmount)
	ds.Poll_MaxVotesPerOption(id).Lazy(&t.MaxVotesPerOption)
	ds.Poll_GlobalNo(id).Lazy(&t.GlobalNo)
	ds.Poll_CryptSignature(id).Lazy(&t.CryptSignature)
	ds.Poll_EntitledGroupIDs(id).Lazy(&t.EntitledGroupIDs)
	ds.Poll_ProjectionIDs(id).Lazy(&t.ProjectionIDs)
	ds.Poll_IsPseudoanonymized(id).Lazy(&t.IsPseudoanonymized)
	ds.Poll_OptionIDs(id).Lazy(&t.OptionIDs)
	ds.Poll_MinVotesAmount(id).Lazy(&t.MinVotesAmount)
	ds.Poll_Title(id).Lazy(&t.Title)
	ds.Poll_VoteCount(id).Lazy(&t.VoteCount)
	ds.Poll_Votesinvalid(id).Lazy(&t.Votesinvalid)
	ds.Poll_Votesvalid(id).Lazy(&t.Votesvalid)
	ds.Poll_GlobalAbstain(id).Lazy(&t.GlobalAbstain)
	ds.Poll_MeetingID(id).Lazy(&t.MeetingID)
	ds.Poll_EntitledUsersAtStop(id).Lazy(&t.EntitledUsersAtStop)
	ds.Poll_GlobalYes(id).Lazy(&t.GlobalYes)
	ds.Poll_State(id).Lazy(&t.State)
	ds.Poll_VotesSignature(id).Lazy(&t.VotesSignature)
	ds.Poll_Votescast(id).Lazy(&t.Votescast)
	ds.Poll_Backend(id).Lazy(&t.Backend)
	ds.Poll_ContentObjectID(id).Lazy(&t.ContentObjectID)
	ds.Poll_Type(id).Lazy(&t.Type)
	ds.Poll_CryptKey(id).Lazy(&t.CryptKey)
	ds.Poll_GlobalOptionID(id).Lazy(&t.GlobalOptionID)
	ds.Poll_Pollmethod(id).Lazy(&t.Pollmethod)
	ds.Poll_SequentialNumber(id).Lazy(&t.SequentialNumber)
	ds.Poll_VotedIDs(id).Lazy(&t.VotedIDs)
	ds.Poll_Description(id).Lazy(&t.Description)
	ds.Poll_ID(id).Lazy(&t.ID)
	ds.Poll_OnehundredPercentBase(id).Lazy(&t.OnehundredPercentBase)
	ds.Poll_VotesRaw(id).Lazy(&t.VotesRaw)
}

func (t *Poll) preload(ds *Fetch, id int) {
	ds.Poll_MaxVotesAmount(id).Preload()
	ds.Poll_MaxVotesPerOption(id).Preload()
	ds.Poll_GlobalNo(id).Preload()
	ds.Poll_CryptSignature(id).Preload()
	ds.Poll_EntitledGroupIDs(id).Preload()
	ds.Poll_ProjectionIDs(id).Preload()
	ds.Poll_IsPseudoanonymized(id).Preload()
	ds.Poll_OptionIDs(id).Preload()
	ds.Poll_MinVotesAmount(id).Preload()
	ds.Poll_Title(id).Preload()
	ds.Poll_VoteCount(id).Preload()
	ds.Poll_Votesinvalid(id).Preload()
	ds.Poll_Votesvalid(id).Preload()
	ds.Poll_GlobalAbstain(id).Preload()
	ds.Poll_MeetingID(id).Preload()
	ds.Poll_EntitledUsersAtStop(id).Preload()
	ds.Poll_GlobalYes(id).Preload()
	ds.Poll_State(id).Preload()
	ds.Poll_VotesSignature(id).Preload()
	ds.Poll_Votescast(id).Preload()
	ds.Poll_Backend(id).Preload()
	ds.Poll_ContentObjectID(id).Preload()
	ds.Poll_Type(id).Preload()
	ds.Poll_CryptKey(id).Preload()
	ds.Poll_GlobalOptionID(id).Preload()
	ds.Poll_Pollmethod(id).Preload()
	ds.Poll_SequentialNumber(id).Preload()
	ds.Poll_VotedIDs(id).Preload()
	ds.Poll_Description(id).Preload()
	ds.Poll_ID(id).Preload()
	ds.Poll_OnehundredPercentBase(id).Preload()
	ds.Poll_VotesRaw(id).Preload()
}

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
}

func (t *PollCandidate) lazy(ds *Fetch, id int) {
	ds.PollCandidate_ID(id).Lazy(&t.ID)
	ds.PollCandidate_MeetingID(id).Lazy(&t.MeetingID)
	ds.PollCandidate_PollCandidateListID(id).Lazy(&t.PollCandidateListID)
	ds.PollCandidate_UserID(id).Lazy(&t.UserID)
	ds.PollCandidate_Weight(id).Lazy(&t.Weight)
}

func (t *PollCandidate) preload(ds *Fetch, id int) {
	ds.PollCandidate_ID(id).Preload()
	ds.PollCandidate_MeetingID(id).Preload()
	ds.PollCandidate_PollCandidateListID(id).Preload()
	ds.PollCandidate_UserID(id).Preload()
	ds.PollCandidate_Weight(id).Preload()
}

func (r *Fetch) PollCandidate(id int) *ValueCollection[PollCandidate, *PollCandidate] {
	return &ValueCollection[PollCandidate, *PollCandidate]{
		id:    id,
		fetch: r,
	}
}

// Projection has all fields from projection.
type Projection struct {
	PreviewProjectorID Maybe[int]
	Content            json.RawMessage
	ContentObjectID    string
	HistoryProjectorID Maybe[int]
	Options            json.RawMessage
	Stable             bool
	Type               string
	Weight             int
	CurrentProjectorID Maybe[int]
	ID                 int
	MeetingID          int
}

func (t *Projection) lazy(ds *Fetch, id int) {
	ds.Projection_PreviewProjectorID(id).Lazy(&t.PreviewProjectorID)
	ds.Projection_Content(id).Lazy(&t.Content)
	ds.Projection_ContentObjectID(id).Lazy(&t.ContentObjectID)
	ds.Projection_HistoryProjectorID(id).Lazy(&t.HistoryProjectorID)
	ds.Projection_Options(id).Lazy(&t.Options)
	ds.Projection_Stable(id).Lazy(&t.Stable)
	ds.Projection_Type(id).Lazy(&t.Type)
	ds.Projection_Weight(id).Lazy(&t.Weight)
	ds.Projection_CurrentProjectorID(id).Lazy(&t.CurrentProjectorID)
	ds.Projection_ID(id).Lazy(&t.ID)
	ds.Projection_MeetingID(id).Lazy(&t.MeetingID)
}

func (t *Projection) preload(ds *Fetch, id int) {
	ds.Projection_PreviewProjectorID(id).Preload()
	ds.Projection_Content(id).Preload()
	ds.Projection_ContentObjectID(id).Preload()
	ds.Projection_HistoryProjectorID(id).Preload()
	ds.Projection_Options(id).Preload()
	ds.Projection_Stable(id).Preload()
	ds.Projection_Type(id).Preload()
	ds.Projection_Weight(id).Preload()
	ds.Projection_CurrentProjectorID(id).Preload()
	ds.Projection_ID(id).Preload()
	ds.Projection_MeetingID(id).Preload()
}

func (r *Fetch) Projection(id int) *ValueCollection[Projection, *Projection] {
	return &ValueCollection[Projection, *Projection]{
		id:    id,
		fetch: r,
	}
}

// Speaker has all fields from speaker.
type Speaker struct {
	BeginTime                      int
	PauseTime                      int
	PointOfOrder                   bool
	PointOfOrderCategoryID         Maybe[int]
	SpeechState                    string
	MeetingID                      int
	TotalPause                     int
	ListOfSpeakersID               int
	StructureLevelListOfSpeakersID Maybe[int]
	UnpauseTime                    int
	EndTime                        int
	ID                             int
	MeetingUserID                  Maybe[int]
	Note                           string
	Weight                         int
}

func (t *Speaker) lazy(ds *Fetch, id int) {
	ds.Speaker_BeginTime(id).Lazy(&t.BeginTime)
	ds.Speaker_PauseTime(id).Lazy(&t.PauseTime)
	ds.Speaker_PointOfOrder(id).Lazy(&t.PointOfOrder)
	ds.Speaker_PointOfOrderCategoryID(id).Lazy(&t.PointOfOrderCategoryID)
	ds.Speaker_SpeechState(id).Lazy(&t.SpeechState)
	ds.Speaker_MeetingID(id).Lazy(&t.MeetingID)
	ds.Speaker_TotalPause(id).Lazy(&t.TotalPause)
	ds.Speaker_ListOfSpeakersID(id).Lazy(&t.ListOfSpeakersID)
	ds.Speaker_StructureLevelListOfSpeakersID(id).Lazy(&t.StructureLevelListOfSpeakersID)
	ds.Speaker_UnpauseTime(id).Lazy(&t.UnpauseTime)
	ds.Speaker_EndTime(id).Lazy(&t.EndTime)
	ds.Speaker_ID(id).Lazy(&t.ID)
	ds.Speaker_MeetingUserID(id).Lazy(&t.MeetingUserID)
	ds.Speaker_Note(id).Lazy(&t.Note)
	ds.Speaker_Weight(id).Lazy(&t.Weight)
}

func (t *Speaker) preload(ds *Fetch, id int) {
	ds.Speaker_BeginTime(id).Preload()
	ds.Speaker_PauseTime(id).Preload()
	ds.Speaker_PointOfOrder(id).Preload()
	ds.Speaker_PointOfOrderCategoryID(id).Preload()
	ds.Speaker_SpeechState(id).Preload()
	ds.Speaker_MeetingID(id).Preload()
	ds.Speaker_TotalPause(id).Preload()
	ds.Speaker_ListOfSpeakersID(id).Preload()
	ds.Speaker_StructureLevelListOfSpeakersID(id).Preload()
	ds.Speaker_UnpauseTime(id).Preload()
	ds.Speaker_EndTime(id).Preload()
	ds.Speaker_ID(id).Preload()
	ds.Speaker_MeetingUserID(id).Preload()
	ds.Speaker_Note(id).Preload()
	ds.Speaker_Weight(id).Preload()
}

func (r *Fetch) Speaker(id int) *ValueCollection[Speaker, *Speaker] {
	return &ValueCollection[Speaker, *Speaker]{
		id:    id,
		fetch: r,
	}
}

// StructureLevel has all fields from structure_level.
type StructureLevel struct {
	MeetingUserIDs                  []int
	Name                            string
	StructureLevelListOfSpeakersIDs []int
	Color                           string
	DefaultTime                     int
	ID                              int
	MeetingID                       int
}

func (t *StructureLevel) lazy(ds *Fetch, id int) {
	ds.StructureLevel_MeetingUserIDs(id).Lazy(&t.MeetingUserIDs)
	ds.StructureLevel_Name(id).Lazy(&t.Name)
	ds.StructureLevel_StructureLevelListOfSpeakersIDs(id).Lazy(&t.StructureLevelListOfSpeakersIDs)
	ds.StructureLevel_Color(id).Lazy(&t.Color)
	ds.StructureLevel_DefaultTime(id).Lazy(&t.DefaultTime)
	ds.StructureLevel_ID(id).Lazy(&t.ID)
	ds.StructureLevel_MeetingID(id).Lazy(&t.MeetingID)
}

func (t *StructureLevel) preload(ds *Fetch, id int) {
	ds.StructureLevel_MeetingUserIDs(id).Preload()
	ds.StructureLevel_Name(id).Preload()
	ds.StructureLevel_StructureLevelListOfSpeakersIDs(id).Preload()
	ds.StructureLevel_Color(id).Preload()
	ds.StructureLevel_DefaultTime(id).Preload()
	ds.StructureLevel_ID(id).Preload()
	ds.StructureLevel_MeetingID(id).Preload()
}

func (r *Fetch) StructureLevel(id int) *ValueCollection[StructureLevel, *StructureLevel] {
	return &ValueCollection[StructureLevel, *StructureLevel]{
		id:    id,
		fetch: r,
	}
}

// User has all fields from user.
type User struct {
	MeetingIDs                  []int
	MeetingUserIDs              []int
	PollCandidateIDs            []int
	LastEmailSent               int
	IsPhysicalPerson            bool
	IsPresentInMeetingIDs       []int
	OrganizationID              int
	DelegatedVoteIDs            []int
	PollVotedIDs                []int
	Pronoun                     string
	SamlID                      string
	ID                          int
	DefaultPassword             string
	OptionIDs                   []int
	Title                       string
	VoteIDs                     []int
	CommitteeManagementIDs      []int
	MemberNumber                string
	Password                    string
	Username                    string
	CommitteeIDs                []int
	IsActive                    bool
	CanChangeOwnPassword        bool
	GenderID                    Maybe[int]
	LastName                    string
	Email                       string
	FirstName                   string
	ForwardingCommitteeIDs      []int
	IsDemoUser                  bool
	LastLogin                   int
	OrganizationManagementLevel string
	DefaultVoteWeight           string
}

func (t *User) lazy(ds *Fetch, id int) {
	ds.User_MeetingIDs(id).Lazy(&t.MeetingIDs)
	ds.User_MeetingUserIDs(id).Lazy(&t.MeetingUserIDs)
	ds.User_PollCandidateIDs(id).Lazy(&t.PollCandidateIDs)
	ds.User_LastEmailSent(id).Lazy(&t.LastEmailSent)
	ds.User_IsPhysicalPerson(id).Lazy(&t.IsPhysicalPerson)
	ds.User_IsPresentInMeetingIDs(id).Lazy(&t.IsPresentInMeetingIDs)
	ds.User_OrganizationID(id).Lazy(&t.OrganizationID)
	ds.User_DelegatedVoteIDs(id).Lazy(&t.DelegatedVoteIDs)
	ds.User_PollVotedIDs(id).Lazy(&t.PollVotedIDs)
	ds.User_Pronoun(id).Lazy(&t.Pronoun)
	ds.User_SamlID(id).Lazy(&t.SamlID)
	ds.User_ID(id).Lazy(&t.ID)
	ds.User_DefaultPassword(id).Lazy(&t.DefaultPassword)
	ds.User_OptionIDs(id).Lazy(&t.OptionIDs)
	ds.User_Title(id).Lazy(&t.Title)
	ds.User_VoteIDs(id).Lazy(&t.VoteIDs)
	ds.User_CommitteeManagementIDs(id).Lazy(&t.CommitteeManagementIDs)
	ds.User_MemberNumber(id).Lazy(&t.MemberNumber)
	ds.User_Password(id).Lazy(&t.Password)
	ds.User_Username(id).Lazy(&t.Username)
	ds.User_CommitteeIDs(id).Lazy(&t.CommitteeIDs)
	ds.User_IsActive(id).Lazy(&t.IsActive)
	ds.User_CanChangeOwnPassword(id).Lazy(&t.CanChangeOwnPassword)
	ds.User_GenderID(id).Lazy(&t.GenderID)
	ds.User_LastName(id).Lazy(&t.LastName)
	ds.User_Email(id).Lazy(&t.Email)
	ds.User_FirstName(id).Lazy(&t.FirstName)
	ds.User_ForwardingCommitteeIDs(id).Lazy(&t.ForwardingCommitteeIDs)
	ds.User_IsDemoUser(id).Lazy(&t.IsDemoUser)
	ds.User_LastLogin(id).Lazy(&t.LastLogin)
	ds.User_OrganizationManagementLevel(id).Lazy(&t.OrganizationManagementLevel)
	ds.User_DefaultVoteWeight(id).Lazy(&t.DefaultVoteWeight)
}

func (t *User) preload(ds *Fetch, id int) {
	ds.User_MeetingIDs(id).Preload()
	ds.User_MeetingUserIDs(id).Preload()
	ds.User_PollCandidateIDs(id).Preload()
	ds.User_LastEmailSent(id).Preload()
	ds.User_IsPhysicalPerson(id).Preload()
	ds.User_IsPresentInMeetingIDs(id).Preload()
	ds.User_OrganizationID(id).Preload()
	ds.User_DelegatedVoteIDs(id).Preload()
	ds.User_PollVotedIDs(id).Preload()
	ds.User_Pronoun(id).Preload()
	ds.User_SamlID(id).Preload()
	ds.User_ID(id).Preload()
	ds.User_DefaultPassword(id).Preload()
	ds.User_OptionIDs(id).Preload()
	ds.User_Title(id).Preload()
	ds.User_VoteIDs(id).Preload()
	ds.User_CommitteeManagementIDs(id).Preload()
	ds.User_MemberNumber(id).Preload()
	ds.User_Password(id).Preload()
	ds.User_Username(id).Preload()
	ds.User_CommitteeIDs(id).Preload()
	ds.User_IsActive(id).Preload()
	ds.User_CanChangeOwnPassword(id).Preload()
	ds.User_GenderID(id).Preload()
	ds.User_LastName(id).Preload()
	ds.User_Email(id).Preload()
	ds.User_FirstName(id).Preload()
	ds.User_ForwardingCommitteeIDs(id).Preload()
	ds.User_IsDemoUser(id).Preload()
	ds.User_LastLogin(id).Preload()
	ds.User_OrganizationManagementLevel(id).Preload()
	ds.User_DefaultVoteWeight(id).Preload()
}

func (r *Fetch) User(id int) *ValueCollection[User, *User] {
	return &ValueCollection[User, *User]{
		id:    id,
		fetch: r,
	}
}

// MotionState has all fields from motion_state.
type MotionState struct {
	WorkflowID                       int
	AllowSubmitterEdit               bool
	CssClass                         string
	Name                             string
	NextStateIDs                     []int
	PreviousStateIDs                 []int
	ShowStateExtensionField          bool
	SubmitterWithdrawStateID         Maybe[int]
	Weight                           int
	AllowCreatePoll                  bool
	AllowSupport                     bool
	MeetingID                        int
	MergeAmendmentIntoFinal          string
	SetWorkflowTimestamp             bool
	SetNumber                        bool
	AllowMotionForwarding            bool
	FirstStateOfWorkflowID           Maybe[int]
	ID                               int
	MotionIDs                        []int
	MotionRecommendationIDs          []int
	IsInternal                       bool
	RecommendationLabel              string
	Restrictions                     []string
	ShowRecommendationExtensionField bool
	SubmitterWithdrawBackIDs         []int
}

func (t *MotionState) lazy(ds *Fetch, id int) {
	ds.MotionState_WorkflowID(id).Lazy(&t.WorkflowID)
	ds.MotionState_AllowSubmitterEdit(id).Lazy(&t.AllowSubmitterEdit)
	ds.MotionState_CssClass(id).Lazy(&t.CssClass)
	ds.MotionState_Name(id).Lazy(&t.Name)
	ds.MotionState_NextStateIDs(id).Lazy(&t.NextStateIDs)
	ds.MotionState_PreviousStateIDs(id).Lazy(&t.PreviousStateIDs)
	ds.MotionState_ShowStateExtensionField(id).Lazy(&t.ShowStateExtensionField)
	ds.MotionState_SubmitterWithdrawStateID(id).Lazy(&t.SubmitterWithdrawStateID)
	ds.MotionState_Weight(id).Lazy(&t.Weight)
	ds.MotionState_AllowCreatePoll(id).Lazy(&t.AllowCreatePoll)
	ds.MotionState_AllowSupport(id).Lazy(&t.AllowSupport)
	ds.MotionState_MeetingID(id).Lazy(&t.MeetingID)
	ds.MotionState_MergeAmendmentIntoFinal(id).Lazy(&t.MergeAmendmentIntoFinal)
	ds.MotionState_SetWorkflowTimestamp(id).Lazy(&t.SetWorkflowTimestamp)
	ds.MotionState_SetNumber(id).Lazy(&t.SetNumber)
	ds.MotionState_AllowMotionForwarding(id).Lazy(&t.AllowMotionForwarding)
	ds.MotionState_FirstStateOfWorkflowID(id).Lazy(&t.FirstStateOfWorkflowID)
	ds.MotionState_ID(id).Lazy(&t.ID)
	ds.MotionState_MotionIDs(id).Lazy(&t.MotionIDs)
	ds.MotionState_MotionRecommendationIDs(id).Lazy(&t.MotionRecommendationIDs)
	ds.MotionState_IsInternal(id).Lazy(&t.IsInternal)
	ds.MotionState_RecommendationLabel(id).Lazy(&t.RecommendationLabel)
	ds.MotionState_Restrictions(id).Lazy(&t.Restrictions)
	ds.MotionState_ShowRecommendationExtensionField(id).Lazy(&t.ShowRecommendationExtensionField)
	ds.MotionState_SubmitterWithdrawBackIDs(id).Lazy(&t.SubmitterWithdrawBackIDs)
}

func (t *MotionState) preload(ds *Fetch, id int) {
	ds.MotionState_WorkflowID(id).Preload()
	ds.MotionState_AllowSubmitterEdit(id).Preload()
	ds.MotionState_CssClass(id).Preload()
	ds.MotionState_Name(id).Preload()
	ds.MotionState_NextStateIDs(id).Preload()
	ds.MotionState_PreviousStateIDs(id).Preload()
	ds.MotionState_ShowStateExtensionField(id).Preload()
	ds.MotionState_SubmitterWithdrawStateID(id).Preload()
	ds.MotionState_Weight(id).Preload()
	ds.MotionState_AllowCreatePoll(id).Preload()
	ds.MotionState_AllowSupport(id).Preload()
	ds.MotionState_MeetingID(id).Preload()
	ds.MotionState_MergeAmendmentIntoFinal(id).Preload()
	ds.MotionState_SetWorkflowTimestamp(id).Preload()
	ds.MotionState_SetNumber(id).Preload()
	ds.MotionState_AllowMotionForwarding(id).Preload()
	ds.MotionState_FirstStateOfWorkflowID(id).Preload()
	ds.MotionState_ID(id).Preload()
	ds.MotionState_MotionIDs(id).Preload()
	ds.MotionState_MotionRecommendationIDs(id).Preload()
	ds.MotionState_IsInternal(id).Preload()
	ds.MotionState_RecommendationLabel(id).Preload()
	ds.MotionState_Restrictions(id).Preload()
	ds.MotionState_ShowRecommendationExtensionField(id).Preload()
	ds.MotionState_SubmitterWithdrawBackIDs(id).Preload()
}

func (r *Fetch) MotionState(id int) *ValueCollection[MotionState, *MotionState] {
	return &ValueCollection[MotionState, *MotionState]{
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
}

func (t *Tag) lazy(ds *Fetch, id int) {
	ds.Tag_ID(id).Lazy(&t.ID)
	ds.Tag_MeetingID(id).Lazy(&t.MeetingID)
	ds.Tag_Name(id).Lazy(&t.Name)
	ds.Tag_TaggedIDs(id).Lazy(&t.TaggedIDs)
}

func (t *Tag) preload(ds *Fetch, id int) {
	ds.Tag_ID(id).Preload()
	ds.Tag_MeetingID(id).Preload()
	ds.Tag_Name(id).Preload()
	ds.Tag_TaggedIDs(id).Preload()
}

func (r *Fetch) Tag(id int) *ValueCollection[Tag, *Tag] {
	return &ValueCollection[Tag, *Tag]{
		id:    id,
		fetch: r,
	}
}
