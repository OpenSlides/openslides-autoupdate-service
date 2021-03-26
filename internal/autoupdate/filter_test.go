package autoupdate

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterFirstCall(t *testing.T) {
	var f filter
	data := map[string]json.RawMessage{
		"k1": []byte("v1"),
		"k2": nil,
	}

	f.filter(data)

	assert.Equal(
		t,
		map[string]json.RawMessage{
			"k1": []byte("v1"),
		},
		data,
	)
}

func TestFilterChange(t *testing.T) {
	for _, tt := range []struct {
		name    string
		origian map[string]json.RawMessage
		new     map[string]json.RawMessage
		expect  map[string]json.RawMessage
	}{
		{
			"Data does not change",
			map[string]json.RawMessage{
				"k1": []byte("v1"),
			},
			map[string]json.RawMessage{
				"k1": []byte("v1"),
			},
			map[string]json.RawMessage{},
		},
		{
			"nil does not change",
			map[string]json.RawMessage{
				"k1": nil,
			},
			map[string]json.RawMessage{
				"k1": nil,
			},
			map[string]json.RawMessage{},
		},
		{
			"data does change to nil",
			map[string]json.RawMessage{
				"k1": []byte("v1"),
			},
			map[string]json.RawMessage{
				"k1": nil,
			},
			map[string]json.RawMessage{
				"k1": nil,
			},
		},
		{
			"nil does change to data",
			map[string]json.RawMessage{
				"k1": nil,
			},
			map[string]json.RawMessage{
				"k1": []byte("v1"),
			},
			map[string]json.RawMessage{
				"k1": []byte("v1"),
			},
		},
		{
			"new key",
			map[string]json.RawMessage{
				"k1": []byte("v1"),
			},
			map[string]json.RawMessage{
				"k2": []byte("v2"),
			},
			map[string]json.RawMessage{
				"k2": []byte("v2"),
			},
		},
		{
			"new key with old key",
			map[string]json.RawMessage{
				"k1": []byte("v1"),
			},
			map[string]json.RawMessage{
				"k1": []byte("v1"),
				"k2": []byte("v2"),
			},
			map[string]json.RawMessage{
				"k2": []byte("v2"),
			},
		},
		{
			"new key nil",
			map[string]json.RawMessage{
				"k1": []byte("v1"),
			},
			map[string]json.RawMessage{
				"k2": nil,
			},
			map[string]json.RawMessage{},
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
