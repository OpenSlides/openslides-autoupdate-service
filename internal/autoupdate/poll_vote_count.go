package autoupdate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func datastorePollVoteCount(ctx context.Context, fqfield string, changed map[string][]byte) ([]byte, error) {
	if changed != nil {
		return changed[fqfield], nil
	}

	values, err := requestKeys("http://localhost:9010/internal/vote/vote_count", []string{fqfield})
	if err != nil {
		return nil, fmt.Errorf("loading key %q from vote-service: %w", fqfield, err)
	}
	return values[fqfield], nil
}

// requestKeys request a list of keys by the datastore.
//
// If an error happens, no key is returned.
//
// The returned map contains exacply the given keys. If a key does not exist in
// the datastore, then the value of this key is <nil>.
func requestKeys(url string, keys []string) (map[string][]byte, error) {
	requestData, err := keysToGetManyRequest(keys)
	if err != nil {
		return nil, fmt.Errorf("creating GetManyRequest: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(requestData))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Do(req)
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

	responseData, err := getManyResponceToKeyValue(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse responce: %w", err)
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

// keysToGetManyRequest a json envoding of the get_many request.
func keysToGetManyRequest(keys []string) ([]byte, error) {
	request := struct {
		Requests []string `json:"requests"`
	}{keys}
	return json.Marshal(request)
}

// getManyResponceToKeyValue reads the responce from the getMany request and
// returns the content as key-values.
func getManyResponceToKeyValue(r io.Reader) (map[string][]byte, error) {
	var data map[string]map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding responce: %w", err)
	}

	keyValue := make(map[string][]byte)
	for collection, idField := range data {
		for id, fieldValue := range idField {
			for field, value := range fieldValue {
				keyValue[collection+"/"+id+"/"+field] = value
			}
		}
	}
	return keyValue, nil
}
