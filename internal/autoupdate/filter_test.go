package autoupdate

import (
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

var (
	myKey1 = dskey.MustKey("user/1/username")
	myKey2 = dskey.MustKey("user/2/username")
)

func TestFilterFirstCall(t *testing.T) {
	data := map[dskey.Key][]byte{
		myKey1: []byte("v1"),
		myKey2: nil,
	}

	new(filter).filter(data)

	expect := map[dskey.Key][]byte{
		myKey1: []byte("v1"),
	}

	if !reflect.DeepEqual(data, expect) {
		t.Errorf("got %v, expected %v", data, expect)
	}
}

func TestFilterChange(t *testing.T) {
	for _, tt := range []struct {
		name    string
		origian map[dskey.Key][]byte
		new     map[dskey.Key][]byte
		expect  map[dskey.Key][]byte
	}{
		{
			"Data does not change",
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{},
		},
		{
			"nil does not change",
			map[dskey.Key][]byte{
				myKey1: nil,
			},
			map[dskey.Key][]byte{
				myKey1: nil,
			},
			map[dskey.Key][]byte{},
		},
		{
			"data does change to nil",
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey1: nil,
			},
			map[dskey.Key][]byte{
				myKey1: nil,
			},
		},
		{
			"nil does change to data",
			map[dskey.Key][]byte{
				myKey1: nil,
			},
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
		},
		{
			"new key",
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
				myKey2: []byte("v2"),
			},
			map[dskey.Key][]byte{
				myKey2: []byte("v2"),
			},
		},
		{
			"new key with old key",
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
				myKey2: []byte("v2"),
			},
			map[dskey.Key][]byte{
				myKey2: []byte("v2"),
			},
		},
		{
			"new key nil",
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
				myKey2: nil,
			},
			map[dskey.Key][]byte{},
		},
		{
			"don't ask for second key",
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey2: []byte("v2"),
			},
			map[dskey.Key][]byte{
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
		origian     map[dskey.Key][]byte
		firstChange map[dskey.Key][]byte
		new         map[dskey.Key][]byte
		expect      map[dskey.Key][]byte
	}{
		{
			"Key does not change",
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{
				myKey1: []byte("v1"),
			},
			map[dskey.Key][]byte{},
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
