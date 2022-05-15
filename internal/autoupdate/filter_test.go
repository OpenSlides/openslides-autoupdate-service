package autoupdate

import (
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

func MustKey(in string) datastore.Key {
	k, err := datastore.KeyFromString(in)
	if err != nil {
		panic(err)
	}
	return k
}

var (
	myKey1 = MustKey("collection/1/field")
	myKey2 = MustKey("collection/2/field")
)

func TestFilterFirstCall(t *testing.T) {
	data := map[datastore.Key][]byte{
		myKey1: []byte("v1"),
		myKey2: nil,
	}

	new(filter).filter(data)

	expect := map[datastore.Key][]byte{
		myKey1: []byte("v1"),
	}

	if !reflect.DeepEqual(data, expect) {
		t.Errorf("got %v, expected %v", data, expect)
	}
}

func TestFilterChange(t *testing.T) {
	for _, tt := range []struct {
		name    string
		origian map[datastore.Key][]byte
		new     map[datastore.Key][]byte
		expect  map[datastore.Key][]byte
	}{
		{
			"Data does not change",
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{},
		},
		{
			"nil does not change",
			map[datastore.Key][]byte{
				myKey1: nil,
			},
			map[datastore.Key][]byte{
				myKey1: nil,
			},
			map[datastore.Key][]byte{},
		},
		{
			"data does change to nil",
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey1: nil,
			},
			map[datastore.Key][]byte{
				myKey1: nil,
			},
		},
		{
			"nil does change to data",
			map[datastore.Key][]byte{
				myKey1: nil,
			},
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
		},
		{
			"new key",
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
				myKey2: []byte("v2"),
			},
			map[datastore.Key][]byte{
				myKey2: []byte("v2"),
			},
		},
		{
			"new key with old key",
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
				myKey2: []byte("v2"),
			},
			map[datastore.Key][]byte{
				myKey2: []byte("v2"),
			},
		},
		{
			"new key nil",
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
				myKey2: nil,
			},
			map[datastore.Key][]byte{},
		},
		{
			"don't ask for second key",
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey2: []byte("v2"),
			},
			map[datastore.Key][]byte{
				myKey2: []byte("v2"),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var f filter
			f.filter(tt.origian)
			f.filter(tt.new)
			if !reflect.DeepEqual(tt.new, tt.expect) {
				t.Errorf("got %v, expected %v", tt.new, tt.expect)
			}
		})
	}
}

func TestFilterChangeTwice(t *testing.T) {
	for _, tt := range []struct {
		name        string
		origian     map[datastore.Key][]byte
		firstChange map[datastore.Key][]byte
		new         map[datastore.Key][]byte
		expect      map[datastore.Key][]byte
	}{
		{
			"Key does not change",
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[datastore.Key][]byte{},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var f filter
			f.filter(tt.origian)
			f.filter(tt.firstChange)
			f.filter(tt.new)
			if !reflect.DeepEqual(tt.new, tt.expect) {
				t.Errorf("got %v, expected %v", tt.new, tt.expect)
			}
		})
	}
}
