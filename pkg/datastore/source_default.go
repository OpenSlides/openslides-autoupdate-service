package datastore

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// Updater returns keys that have changes. Blocks until there is
// changed data.
//
// Deprivated: Use redis directly.
type Updater interface {
	Update(context.Context) (map[string][]byte, error)
}

// SourceDatastore receives the data from the datastore-reader via http and
// updates via the redis message bus.
type SourceDatastore struct {
	url     string
	client  *http.Client
	updater Updater // TODO: Replace this with the real redis backend.
}

// NewSourceDatastore initializes a SourceDatastore.
func NewSourceDatastore(url string, updater Updater) *SourceDatastore {
	return &SourceDatastore{
		url: url + urlPath,
		client: &http.Client{
			Timeout: httpTimeout,
		},
		updater: updater,
	}
}

// Get fetches the request keys from the datastore-reader.
func (s *SourceDatastore) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	requestData, err := keysToGetManyRequest(keys)
	if err != nil {
		return nil, fmt.Errorf("creating GetManyRequest: %w", err)
	}

	req, err := http.NewRequest("POST", s.url, bytes.NewReader(requestData))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting keys `%v`: %w", keys, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("datastore returned status %s", resp.Status)
		}
		return nil, fmt.Errorf("datastore returned status %s: %s", resp.Status, body)
	}

	responseData, err := getManyResponseToKeyValue(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	// Add keys that where not returned.
	for _, k := range keys {
		if _, ok := responseData[k]; ok {
			continue
		}
		responseData[k] = nil
	}

	return responseData, nil
}

// Update updates the data from the redis message bus.
func (s *SourceDatastore) Update(ctx context.Context) (map[string][]byte, error) {
	return s.updater.Update(ctx)
}
