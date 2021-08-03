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
		"testmodel/1/number":     "456",
		"testmodel/1/text":       `"my text"`,
		"testmodel/1/friend_ids": "[1,2,3]",
	}))

	object := fetch.Object(context.Background(), "testmodel/1", "number", "text", "friend_ids")
	require.NoError(t, fetch.Err(), "Get returned unexpected error")

	assert.Equal(t, json.RawMessage([]byte("456")), object["number"])
	assert.Equal(t, json.RawMessage([]byte(`"my text"`)), object["text"])
	assert.Equal(t, json.RawMessage([]byte("[1,2,3]")), object["friend_ids"])
	assert.ElementsMatch(t, []string{"testmodel/1/id", "testmodel/1/number", "testmodel/1/text", "testmodel/1/friend_ids"}, fetch.Keys())
}

func TestFetchObjectOnError(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id":         "1",
		"testmodel/1/number":     "456",
		"testmodel/1/text":       `"my text"`,
		"testmodel/1/friend_ids": "[1,2,3]",
	})
	ds.InjectError(errors.New("some error"))
	fetch := datastore.NewFetcher(ds)

	fetch.Object(context.Background(), "testmodel/1", "number", "text", "friend_ids")
	if err := fetch.Err(); err == nil {
		t.Fatalf("Object did not return an error")
	}

	keysEqual(t, fetch.Keys(), "testmodel/1/id", "testmodel/1/number", "testmodel/1/text", "testmodel/1/friend_ids")
}

func TestFetchObjectDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{}))

	fetch.Object(context.Background(), "testmodel/1", "text")

	var errDoesNotExist datastore.DoesNotExistError
	if err := fetch.Err(); !errors.As(err, &errDoesNotExist) {
		t.Errorf("fetch.Object returned error %v, expected a does not exist error", err)
	}
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

	if err := fetch.Err(); err != nil {
		t.Errorf("Fetch returned unexpected error %v", err)
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/text")
	if value != "" {
		t.Errorf("Fetch of unexpected key returned %q, expected am empty string", value)
	}
}

func TestFetchValueAfterError(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/text": `"my text"`,
	})
	fetch := datastore.NewFetcher(ds)

	myerr := errors.New("some error")
	ds.InjectError(myerr)
	var errorValue string
	fetch.Fetch(context.Background(), &errorValue, "testmodel/1/error_value")

	ds.InjectError(nil)
	var value string
	fetch.Fetch(context.Background(), &value, "testmodel/1/text")

	if err := fetch.Err(); !errors.Is(err, myerr) {
		t.Errorf("Fetch returned error %q, expected %q", fetch.Err(), myerr)
	}

	if value != "" {
		t.Errorf("Fetch set value after an error to %q", value)
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/error_value")
}

func TestFetchIfExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id": "1",
	}))

	var value string
	fetch.FetchIfExist(context.Background(), &value, "testmodel/1/text")

	if err := fetch.Err(); err != nil {
		t.Errorf("Fetch returned error: %v", err)
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/id", "testmodel/1/text")
}

func TestFetchIfExistObjectDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	fetch := datastore.NewFetcher(dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/text": `"some test"`,
	}))

	var value string
	fetch.FetchIfExist(context.Background(), &value, "testmodel/1/text")

	var errDoesNotExist datastore.DoesNotExistError
	if err := fetch.Err(); !errors.As(err, &errDoesNotExist) {
		t.Errorf("FetchIfExist returned error: %q, expected DoesNotExistError", err)
	}
	keysEqual(t, fetch.Keys(), "testmodel/1/id", "testmodel/1/text")
}

func TestFetchIfExistAfterError(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id":   "1",
		"testmodel/1/text": `"some test"`,
	})
	fetch := datastore.NewFetcher(ds)
	myerr := errors.New("some error")

	ds.InjectError(myerr)
	var value string
	fetch.Fetch(context.Background(), &value, "testmodel/1/text")

	ds.InjectError(nil)
	fetch.FetchIfExist(context.Background(), &value, "testmodel/1/text")

	if err := fetch.Err(); !errors.Is(err, myerr) {
		t.Errorf("Fetch returned error %q, expected %q", err, myerr)
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

	i := datastore.Int(context.Background(), fetch.Fetch, "testmodel/%d/number", 1)

	fmt.Println(i)
	fmt.Println(fetch.Keys())
	fmt.Println(fetch.Err())
	// Output:
	// 0
	// [testmodel/1/number]
	// <nil>
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

	i := datastore.Int(context.Background(), fetch.FetchIfExist, "testmodel/%d/number", 1)

	fmt.Println(i)
	fmt.Println(fetch.Keys())
	fmt.Println(fetch.Err())
	// Output:
	// 0
	// [testmodel/1/id testmodel/1/number]
	// testmodel/1 does not exist.
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
		t.Errorf("Got %d keys, expected %d\nGot:\n%v\nExpected:\n%v", len(got), len(expect), got, expect)
		return
	}

	for i := range got {
		if got[i] != expect[i] {
			t.Errorf("Key[%d] == %q, expected %q", i, got[i], expect[i])
		}
	}
}
