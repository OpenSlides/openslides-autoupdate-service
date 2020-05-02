package autoupdate_test

import (
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestConnect(t *testing.T) {
	c, _, _, close := getConnection()
	defer close()

	data, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	if value, ok := data["user/1/name"]; !ok || value != `"Hello World"` {
		t.Errorf("c.Next() returned %v, expected map[user/1/name:\"Hello World\"", data)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	c, _, disconnect, close := getConnection()
	defer close()

	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	disconnect()
	data, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if data != nil {
		t.Errorf("Expect no new data, got: %v", data)
	}
}

func TestConnectionReadNewData(t *testing.T) {
	c, datastore, _, close := getConnection()
	defer close()
	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	datastore.Update(map[string]string{"user/1/name": `"new value"`})
	datastore.Send(test.Str("user/1/name"))
	data, err := c.Next()

	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if got := len(data); got != 1 {
		t.Errorf("Expected data to have one key, got: %d", got)
	}
	if value, ok := data["user/1/name"]; !ok || value != `"new value"` {
		t.Errorf("c.Next() returned %v, expected %v", data, map[string]string{"user/1/name": `"new value"`})
	}
}
