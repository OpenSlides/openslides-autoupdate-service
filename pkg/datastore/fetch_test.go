package datastore_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchObject(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id":         "1",
		"testmodel/1/text":       `"my text"`,
		"testmodel/1/friend_ids": "[1,2,3]",
	}))

	object := fetch.Object(context.Background(), "testmodel/1", "id", "text", "friend_ids")
	require.NoError(t, fetch.Err(), "Get returned unexpected error")

	assert.Equal(t, json.RawMessage([]byte("1")), object["id"])
	assert.Equal(t, json.RawMessage([]byte(`"my text"`)), object["text"])
	assert.Equal(t, json.RawMessage([]byte("[1,2,3]")), object["friend_ids"])
	assert.ElementsMatch(t, []string{"testmodel/1/id", "testmodel/1/text", "testmodel/1/friend_ids"}, fetch.Keys())
}

func TestFetchObjectFieldDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{}))

	object := fetch.Object(context.Background(), "testmodel/1", "id")
	require.NoError(t, fetch.Err(), "Get returned unexpected error")

	require.Equal(t, 1, len(object))
	require.Nil(t, object["id"])
	assert.ElementsMatch(t, []string{"testmodel/1/id"}, fetch.Keys())
}

func TestFetchValue(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/text": `"my text"`,
	}))

	var value string
	fetch.Fetch(context.Background(), &value, "testmodel/1/text")

	if err := fetch.Err(); err != nil {
		t.Fatalf("Fetch() returned error: %v", err)
	}
	expect := "my text"
	if value != expect {
		t.Errorf("Fetch() fetched value %q, expected %q", value, expect)
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/text")
}

func TestFetchValueDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, nil))

	var value string
	fetch.Fetch(context.Background(), &value, "testmodel/1/text")

	var errDoesNotExist datastore.DoesNotExistError
	if !errors.As(fetch.Err(), &errDoesNotExist) {
		t.Errorf("Fetch returned error %q, expected datastore.DoesNotExist", fetch.Err())
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/text")
}

func TestFetchValueAfterError(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/text": `"my text"`,
	}))

	var doesNotExistValue string
	fetch.Fetch(context.Background(), &doesNotExistValue, "testmodel/1/does_not_exist")
	var value string
	fetch.Fetch(context.Background(), &value, "testmodel/1/text")

	var errDoesNotExist datastore.DoesNotExistError
	if !errors.As(fetch.Err(), &errDoesNotExist) {
		t.Errorf("Fetch returned error %q, expected datastore.DoesNotExist", fetch.Err())
	}

	if value != "" {
		t.Errorf("Fetch set value after an error to %q", value)
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/does_not_exist")
}

func TestFetchIfExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, nil))

	var value string
	fetch.FetchIfExist(context.Background(), &value, "testmodel/1/text")

	if err := fetch.Err(); err != nil {
		t.Errorf("Fetch returned error: %v", err)
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/text")
}

func TestFetchIfExistAfterError(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, nil))

	var value string
	fetch.Fetch(context.Background(), &value, "testmodel/1/text")
	fetch.FetchIfExist(context.Background(), &value, "testmodel/1/text")

	var errDoesNotExist datastore.DoesNotExistError
	if !errors.As(fetch.Err(), &errDoesNotExist) {
		t.Errorf("Fetch returned error %q, expected datastore.DoesNotExist", fetch.Err())
	}
	if value != "" {
		t.Errorf("Fetch set value after an error to %q", value)
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/text")
}

func ExampleInt() {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id": "1",
	}))

	i := datastore.Int(context.Background(), fetch.Fetch, "testmodel/%d/id", 1)

	fmt.Println(i)
	fmt.Println(fetch.Keys())
	fmt.Println(fetch.Err())
	// Output:
	// 1
	// [testmodel/1/id]
	// <nil>
}

func ExampleInt_doesNotExist() {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, nil))

	i := datastore.Int(context.Background(), fetch.Fetch, "testmodel/%d/id", 1)

	fmt.Println(i)
	fmt.Println(fetch.Keys())
	var errDoesNotExist datastore.DoesNotExistError
	errors.As(fetch.Err(), &errDoesNotExist)
	fmt.Println(errDoesNotExist)
	// Output:
	// 0
	// [testmodel/1/id]
	// testmodel/1/id does not exist.
}

func ExampleInt_wrongType() {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id": `"a string"`,
	}))

	i := datastore.Int(context.Background(), fetch.Fetch, "testmodel/%d/id", 1)

	fmt.Println(i)
	fmt.Println(fetch.Keys())
	fmt.Println(fetch.Err() == nil)
	// Output:
	// 0
	// [testmodel/1/id]
	// false
}

func ExampleInt_ifExist() {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, nil))

	i := datastore.Int(context.Background(), fetch.FetchIfExist, "testmodel/%d/id", 1)

	fmt.Println(i)
	fmt.Println(fetch.Keys())
	fmt.Println(fetch.Err())
	// Output:
	// 0
	// [testmodel/1/id]
	// <nil>
}

func ExampleInts() {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/ids": "[1,2,3]",
	}))

	ints := datastore.Ints(context.Background(), fetch.Fetch, "testmodel/%d/ids", 1)

	fmt.Println(ints)
	fmt.Println(fetch.Keys())
	fmt.Println(fetch.Err())
	// Output:
	// [1 2 3]
	// [testmodel/1/ids]
	// <nil>
}

func ExampleString() {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/name": `"hugo"`,
	}))

	str := datastore.String(context.Background(), fetch.Fetch, "testmodel/%d/name", 1)

	fmt.Println(str)
	fmt.Println(fetch.Keys())
	fmt.Println(fetch.Err())
	// Output:
	// hugo
	// [testmodel/1/name]
	// <nil>
}

func keysEqual(t *testing.T, got []string, expect ...string) {
	t.Helper()

	if len(got) != len(expect) {
		t.Errorf("Got %d fields, expected %d", len(got), len(expect))
		return
	}

	for i := range got {
		if got[i] != expect[i] {
			t.Errorf("Field[%d] == %q, expected %q", i, got[i], expect[i])
		}
	}
}
