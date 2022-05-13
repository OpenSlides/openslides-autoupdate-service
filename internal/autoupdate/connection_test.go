package autoupdate_test

import (
	"context"
	"errors"
	"log"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var userNameKey = autoupdate.MustKey("user/1/name")

func TestConnect(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	next, _ := getConnection(shutdownCtx.Done())

	data, err := next(context.Background())
	if err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	if value, ok := data[userNameKey]; !ok || string(value) != `"Hello World"` {
		t.Errorf("next() returned %v, expected map[user/1/name:\"Hello World\"", data)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	next, _ := getConnection(shutdownCtx.Done())
	ctx, disconnect := context.WithCancel(context.Background())

	if _, err := next(ctx); err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	disconnect()
	data, err := next(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("next() returned error %v, expected context.Canceled", err)
	}
	if data != nil {
		t.Errorf("Expect no new data, got: %v", data)
	}
}

func TestConnectionReadNewData(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	next, ds := getConnection(shutdownCtx.Done())
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	if _, err := next(context.Background()); err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	ds.Send(map[datastore.Key][]byte{userNameKey: []byte(`"new value"`)})
	data, err := next(context.Background())

	if err != nil {
		t.Errorf("next() returned an error: %v", err)
	}
	if got := len(data); got != 1 {
		t.Errorf("Expected data to have one key, got: %d", got)
	}
	if value, ok := data[userNameKey]; !ok || string(value) != `"new value"` {
		t.Errorf("next() returned %v, expected %v", data, map[datastore.Key]string{userNameKey: `"new value"`})
	}
}

func TestConnectionEmptyData(t *testing.T) {
	var (
		doesNotExistKey = autoupdate.MustKey("doesnot/1/exist")
		doesExistKey    = userNameKey
	)

	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(shutdownCtx.Done(), map[datastore.Key][]byte{
		doesExistKey: []byte(`"Hello World"`),
	})
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	s := autoupdate.New(ds, test.RestrictAllowed, "")
	kb, _ := keysbuilder.FromKeys(doesExistKey.String(), doesNotExistKey.String())

	t.Run("First response", func(t *testing.T) {
		next := s.Connect(1, kb)

		data, err := next(context.Background())
		require.NoError(t, err)
		assert.Contains(t, data, doesExistKey, "next() should return the existing key")
		assert.NotContains(t, data, doesNotExistKey, "next() should not return a non existing key")
	})

	for _, tt := range []struct {
		name           string
		update         map[datastore.Key][]byte
		expectBlocking bool
		expectExist    bool
		expectNotExist bool
	}{
		{
			name:           "not exist->not exist",
			update:         map[datastore.Key][]byte{doesNotExistKey: nil},
			expectBlocking: true,
		},
		{
			name:           "not exist->exist",
			update:         map[datastore.Key][]byte{doesNotExistKey: []byte("value")},
			expectExist:    false, // existing key gets filtered.
			expectNotExist: true,
		},
		{
			name:           "exist->not exist",
			update:         map[datastore.Key][]byte{doesExistKey: nil},
			expectExist:    true,
			expectNotExist: false,
		},
		{
			name:           "exist->exist",
			update:         map[datastore.Key][]byte{doesExistKey: []byte("new value")},
			expectExist:    true,
			expectNotExist: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			next := s.Connect(1, kb)
			if _, err := next(context.Background()); err != nil {
				t.Errorf("next() returned an error: %v", err)
			}

			ds.Send(tt.update)

			var data map[datastore.Key][]byte
			var err error
			isBlocking := blocking(func() {
				data, err = next(context.Background())
			})

			require.NoError(t, err)
			if tt.expectBlocking {
				assert.True(t, isBlocking, "Expect next() to block")
			} else {
				assert.False(t, isBlocking, "Expect next() not to block.")
			}

			if tt.expectExist {
				assert.Contains(t, data, doesExistKey)
			} else {
				assert.NotContains(t, data, doesExistKey)
			}

			if tt.expectNotExist {
				assert.Contains(t, data, doesNotExistKey)
			} else {
				assert.NotContains(t, data, doesNotExistKey)
			}
		})
	}

	t.Run("exit->not exist-> not exist", func(t *testing.T) {
		next := s.Connect(1, kb)
		if _, err := next(context.Background()); err != nil {
			t.Errorf("next() returned an error: %v", err)
		}

		// First time not exist
		ds.Send(map[datastore.Key][]byte{doesExistKey: nil})

		blocking(func() {
			next(context.Background())
		})

		// Second time not exist
		ds.Send(map[datastore.Key][]byte{doesExistKey: nil})

		var err error
		isBlocking := blocking(func() {
			_, err = next(context.Background())
		})

		require.NoError(t, err)
		assert.True(t, isBlocking, "second request should be blocking")
	})
}

func TestConnectionFilterData(t *testing.T) {
	t.Skipf("TODO")
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(shutdownCtx.Done(), map[datastore.Key][]byte{
		userNameKey: []byte(`"Hello World"`),
	})

	s := autoupdate.New(ds, test.RestrictAllowed, "")
	kb, _ := keysbuilder.FromKeys(userNameKey.String())
	next := s.Connect(1, kb)
	if _, err := next(context.Background()); err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	// send again, value did not change in restricter
	ds.Send(map[datastore.Key][]byte{userNameKey: []byte(`"Hello World"`)})
	data, err := next(context.Background())

	if err != nil {
		t.Errorf("next() returned an error: %v", err)
	}
	if got := len(data); got != 0 {
		t.Errorf("Got %d keys, expected none", got)
	}
	if _, ok := data[userNameKey]; ok {
		t.Errorf("next() returned %v, expected empty map", data)
	}
}

func TestConntectionFilterOnlyOneKey(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(shutdownCtx.Done(), map[datastore.Key][]byte{
		userNameKey: []byte(`"Hello World"`),
	})
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	s := autoupdate.New(ds, test.RestrictAllowed, "")
	kb, _ := keysbuilder.FromKeys(userNameKey.String())
	next := s.Connect(1, kb)
	if _, err := next(context.Background()); err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	ds.Send(map[datastore.Key][]byte{userNameKey: []byte(`"newname"`)})
	data, err := next(context.Background())

	if err != nil {
		t.Errorf("next() returned an error: %v", err)
	}
	if got := len(data); got != 1 {
		t.Errorf("Expected data to have one key, got: %d", got)
	}
	if _, ok := data[userNameKey]; !ok {
		t.Errorf("Returned value does not have key `user/1/name`")
	}
	if got := string(data[userNameKey]); got != `"newname"` {
		t.Errorf("Expect value `newname` got: %s", got)
	}
}

func TestNextNoReturnWhenDataIsRestricted(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(shutdownCtx.Done(), map[datastore.Key][]byte{
		userNameKey: []byte(`"Hello World"`),
	})

	s := autoupdate.New(ds, test.RestrictNotAllowed, "")
	kb, _ := keysbuilder.FromKeys(userNameKey.String())

	next := s.Connect(1, kb)

	t.Run("first call", func(t *testing.T) {
		var data map[datastore.Key][]byte
		var err error
		isBlocked := blocking(func() {
			data, err = next(context.Background())

		})
		require.NoError(t, err, "next() returned an error")
		assert.Empty(t, data, "next() should return data on first call.")
		assert.False(t, isBlocked, "next() should not block on first call.")
	})

	t.Run("next call", func(t *testing.T) {
		var data map[datastore.Key][]byte
		var err error
		isBlocked := blocking(func() {
			data, err = next(context.Background())

		})
		require.NoError(t, err, "next() returned an error")
		assert.Empty(t, data, "next() returned data")
		assert.True(t, isBlocked, "next() did not block")
	})
}

// TestKeyNotRequestedAnymore tests the case, that an object that was indirectly
// requested gets deleted.
//
// This happens, when a object is requested by a keysbuilder not on the first
// level, but on a second level though a relation-list.
//
// In this case the deleted object is removed from the relation-list-field and
// therefore not requested anymore. So the deleted object is not send to the
// client anymore.
//
// The result is, that the client does not get an update, that the object was
// deleted. Only be looking in the relation-list-field the client knows, that it
// should not be interested in the object anymore.
//
// See Issue https://github.com/OpenSlides/openslides-autoupdate-service/issues/321
func TestKeyNotRequestedAnymore(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	datastore := dsmock.NewMockDatastore(shutdownCtx.Done(), dsmock.YAMLData(`---
	organization/1/organization_tag_ids: [1,2]
	organization_tag/1/id: 1
	organization_tag/2/id: 2
	`))
	go datastore.ListenOnUpdates(shutdownCtx, nil)

	s := autoupdate.New(datastore, test.RestrictAllowed, "")
	kb, err := keysbuilder.FromJSON(strings.NewReader(`{
		"collection":"organization",
		"ids":[
		  1
		],
		"fields":{
		  "organization_tag_ids":{
			"type":"relation-list",
			"collection":"organization_tag",
			"fields":{
			  "id":null
			}
		  }
		}
	  }`))

	if err != nil {
		t.Fatalf("Can not build request: %v", err)
	}

	next := s.Connect(1, kb)

	if _, err := next(shutdownCtx); err != nil {
		t.Fatalf("Getting first data: %v", err)
	}

	datastore.Send(dsmock.YAMLData(`
	organization_tag/2/id: null
	organization/1/organization_tag_ids: [1]
	`))

	secondData, err := next(shutdownCtx)
	if err != nil {
		t.Fatalf("Getting second data: %v", err)
	}

	if len(secondData) != 1 {
		t.Errorf("second data contained 2 values, expected only one. Got: %v", secondData)
	}

	if v := string(secondData[autoupdate.MustKey("organization/1/organization_tag_ids")]); v != "[1]" {
		t.Errorf("Got organization/1/organization_tag_ids: %q, expected [1]", v)
	}

	if v, ok := secondData[autoupdate.MustKey("organization_tag/2/id")]; ok {
		t.Errorf("Got value for deleted object organization_tag/2/id: %s", v)
	}
}

// TestKeyRequestedAgain makes sure, that when a key is requested again, it is
// send to the client, even when it has not changed.
//
// See the TestKeyNotRequestedAnymore test and the issue
//
// https://github.com/OpenSlides/openslides-autoupdate-service/issues/382
func TestKeyRequestedAgain(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	datastore := dsmock.NewMockDatastore(shutdownCtx.Done(), dsmock.YAMLData(`---
	organization/1/organization_tag_ids: [1,2]
	organization_tag/1/id: 1
	organization_tag/2/id: 2
	`))
	go datastore.ListenOnUpdates(shutdownCtx, nil)

	s := autoupdate.New(datastore, test.RestrictAllowed, "")
	kb, err := keysbuilder.FromJSON(strings.NewReader(`{
		"collection":"organization",
		"ids":[
		  1
		],
		"fields":{
		  "organization_tag_ids":{
			"type":"relation-list",
			"collection":"organization_tag",
			"fields":{
			  "id":null
			}
		  }
		}
	  }`))

	if err != nil {
		t.Fatalf("Can not build request: %v", err)
	}

	next := s.Connect(1, kb)

	// Receive the initial data
	if _, err := next(shutdownCtx); err != nil {
		t.Fatalf("Getting first data: %v", err)
	}

	datastore.Send(dsmock.YAMLData(`
	organization/1/organization_tag_ids: [1]
	`))

	if _, err := next(shutdownCtx); err != nil {
		t.Fatalf("Getting second data: %v", err)
	}

	datastore.Send(dsmock.YAMLData(`
	organization/1/organization_tag_ids: [1,2]
	`))

	// Receive the third data
	testData, err := next(shutdownCtx)
	if err != nil {
		t.Fatalf("Getting second data: %v", err)
	}

	if len(testData) != 2 {
		t.Errorf("second data contained %d values, expected two. Got: %v", len(testData), testData)
	}

	if v := string(testData[autoupdate.MustKey("organization/1/organization_tag_ids")]); v != "[1,2]" {
		t.Errorf("Got organization/1/organization_tag_ids: %q, expected [1,2]", v)
	}

	if v := string(testData[autoupdate.MustKey("organization_tag/2/id")]); v != "2" {
		t.Errorf("Got organization_tag/2/id: %q, expected 2", v)
	}
}
