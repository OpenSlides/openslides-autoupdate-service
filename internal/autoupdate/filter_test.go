package autoupdate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterFirstCall(t *testing.T) {
	var f filter
	data := map[string][]byte{
		"k1": []byte("v1"),
		"k2": nil,
	}

	f.filter(data)

	assert.Equal(
		t,
		map[string][]byte{
			"k1": []byte("v1"),
		},
		data,
	)
}

func TestFilterChange(t *testing.T) {
	for _, tt := range []struct {
		name    string
		origian map[string][]byte
		new     map[string][]byte
		expect  map[string][]byte
	}{
		{
			"Data does not change",
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{},
		},
		{
			"nil does not change",
			map[string][]byte{
				"k1": nil,
			},
			map[string][]byte{
				"k1": nil,
			},
			map[string][]byte{},
		},
		{
			"data does change to nil",
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k1": nil,
			},
			map[string][]byte{
				"k1": nil,
			},
		},
		{
			"nil does change to data",
			map[string][]byte{
				"k1": nil,
			},
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k1": []byte("v1"),
			},
		},
		{
			"new key",
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k1": []byte("v1"),
				"k2": []byte("v2"),
			},
			map[string][]byte{
				"k2": []byte("v2"),
			},
		},
		{
			"new key with old key",
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k1": []byte("v1"),
				"k2": []byte("v2"),
			},
			map[string][]byte{
				"k2": []byte("v2"),
			},
		},
		{
			"new key nil",
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k1": []byte("v1"),
				"k2": nil,
			},
			map[string][]byte{},
		},
		{
			"don't ask for second key",
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k2": []byte("v2"),
			},
			map[string][]byte{
				"k2": []byte("v2"),
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var f filter
			f.filter(tt.origian)
			f.filter(tt.new)
			assert.Equal(t, tt.expect, tt.new)
		})
	}
}

func TestFilterChangeTwice(t *testing.T) {
	for _, tt := range []struct {
		name        string
		origian     map[string][]byte
		firstChange map[string][]byte
		new         map[string][]byte
		expect      map[string][]byte
	}{
		{
			"Key does not change",
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{
				"k1": []byte("v1"),
			},
			map[string][]byte{},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var f filter
			f.filter(tt.origian)
			f.filter(tt.firstChange)
			f.filter(tt.new)
			assert.Equal(t, tt.expect, tt.new)
		})
	}
}
