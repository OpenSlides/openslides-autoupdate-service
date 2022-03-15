package datastore

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	urlGetMany            = "/internal/datastore/reader/get_many"
	urlHistoryInformation = "/internal/datastore/reader/history_information"
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
		url: url,
		client: &http.Client{
			Timeout: httpTimeout,
		},
		updater: updater,
	}
}

// Get fetches the request keys from the datastore-reader.
func (s *SourceDatastore) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	return s.GetPosition(ctx, 0, keys...)
}

// GetPosition gets keys from the datastore at a specifi position.
//
// Position 0 means the current position.
func (s *SourceDatastore) GetPosition(ctx context.Context, position int, keys ...string) (map[string][]byte, error) {
	requestData, err := keysToGetManyRequest(keys, position)
	if err != nil {
		return nil, fmt.Errorf("creating GetManyRequest: %w", err)
	}

	req, err := http.NewRequest("POST", s.url+urlGetMany, bytes.NewReader(requestData))
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

// HistoryInformation requests the history information for an fqid from the datastore.
func (s *SourceDatastore) HistoryInformation(ctx context.Context, fqid string, w io.Writer) error {
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		s.url+urlHistoryInformation,
		strings.NewReader(fmt.Sprintf(`{"fqid":[%q]}`, fqid)),
	)
	if err != nil {
		return fmt.Errorf("creating request for datastore: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("sending request to datastore: %w", err)
	}
	defer resp.Body.Close()
	defer io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("datastore returned %s", resp.Status)
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		return fmt.Errorf("copping datastore response to client: %w", err)
	}

	return nil
}
