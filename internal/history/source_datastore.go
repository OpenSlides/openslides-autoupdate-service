package history

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"golang.org/x/sync/errgroup"
)

var (
	envDatastoreHost     = environment.NewVariable("DATASTORE_READER_HOST", "localhost", "Host of the datastore reader.")
	envDatastorePort     = environment.NewVariable("DATASTORE_READER_PORT", "9010", "Port of the datastore reader.")
	envDatastoreProtocol = environment.NewVariable("DATASTORE_READER_PROTOCOL", "http", "Protocol of the datastore reader.")

	envDatastoreTimeout         = environment.NewVariable("DATASTORE_TIMEOUT", "3s", "Time until a request to the datastore times out.")
	envDatastoreMaxParallelKeys = environment.NewVariable("DATASTORE_MAX_PARALLEL_KEYS", "1000", "Max keys that are send in one request to the datastore.")
)

const (
	urlGetMany            = "/internal/datastore/reader/get_many"
	urlHistoryInformation = "/internal/datastore/reader/history_information"
)

// sourceDatastore receives the data from the datastore-reader via http and
// updates via the redis message bus.
type sourceDatastore struct {
	url    string
	client *http.Client

	metricDSHitCount  uint64
	maxKeysPerRequest int
}

// newSourceDatastore initializes a SourceDatastore.
func newSourceDatastore(lookup environment.Environmenter) (*sourceDatastore, error) {
	url := fmt.Sprintf(
		"%s://%s:%s",
		envDatastoreProtocol.Value(lookup),
		envDatastoreHost.Value(lookup),
		envDatastorePort.Value(lookup),
	)

	timeout, err := environment.ParseDuration(envDatastoreTimeout.Value(lookup))
	if err != nil {
		return nil, fmt.Errorf("parsing timeout: %w", err)
	}

	maxParallel, err := strconv.Atoi(envDatastoreMaxParallelKeys.Value(lookup))
	if err != nil {
		return nil, fmt.Errorf(
			"environment variable MAX_PARALLEL_KEYS has to be a number, not %s",
			envDatastoreMaxParallelKeys.Value(lookup),
		)
	}

	source := sourceDatastore{
		url: url,
		client: &http.Client{
			Timeout: timeout,
		},
		maxKeysPerRequest: maxParallel,
	}

	return &source, nil
}

func (s *sourceDatastore) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	return s.getPosition(ctx, 0, keys...)
}

// GetPosition gets keys from the datastore at a specifi position.
//
// Position 0 means the current position.
func (s *sourceDatastore) GetPosition(ctx context.Context, position int, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	atomic.AddUint64(&s.metricDSHitCount, 1)
	if len(keys) <= s.maxKeysPerRequest {
		return s.getPosition(ctx, position, keys...)
	}

	// Sort keys to help datastore
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Collection() < keys[j].Collection() {
			return true
		} else if keys[i].Collection() > keys[j].Collection() {
			return false
		}

		if keys[i].ID() < keys[j].ID() {
			return true
		} else if keys[i].ID() > keys[j].ID() {
			return false
		}

		return keys[i].Field() < keys[j].Field()
	})

	eg, ctx := errgroup.WithContext(ctx)

	requestCount := len(keys) / s.maxKeysPerRequest
	if len(keys)%s.maxKeysPerRequest != 0 {
		requestCount++
	}

	results := make([]map[dskey.Key][]byte, requestCount)
	for i := 0; i < len(results); i++ {
		i := i

		eg.Go(func() error {
			from := i * s.maxKeysPerRequest
			to := (i + 1) * s.maxKeysPerRequest
			if to > len(keys) {
				to = len(keys)
			}

			data, err := s.getPosition(ctx, position, keys[from:to]...)
			if err != nil {
				return fmt.Errorf("getting keys %d to %d: %w", from, to-1, err)
			}
			results[i] = data

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	combined := make(map[dskey.Key][]byte, len(keys))
	for _, r := range results {
		for k, v := range r {
			combined[k] = v
		}
	}

	return combined, nil
}

func (s *sourceDatastore) getPosition(ctx context.Context, position int, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	requestData, err := keysToGetManyRequest(keys, position)
	if err != nil {
		return nil, fmt.Errorf("creating GetManyRequest: %w", err)
	}

	req, err := http.NewRequest("POST", s.url+urlGetMany, bytes.NewReader(requestData))
	if err != nil {
		// TODO External Error
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		if oserror.Timeout(err) {
			return nil, oserror.ForAdmin(
				"A request to the datastore got a timeout. The current timeout value is %s. Make sure the datastore-reader is scalled, set a higher value to DATASTORE_TIMEOUT_SECONDS or set a lower value to the environment variable MAX_PARALLEL_KEYS.",
				s.client.Timeout,
			)
		}
		// TODO External Error
		return nil, fmt.Errorf("sending request to datastore: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("datastore returned status %s", resp.Status)
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			errMsg = fmt.Sprintf("%s :%s", errMsg, body)
		}
		// TODO External Error
		return nil, errors.New(errMsg)
	}

	responseData, err := parseGetManyResponse(resp.Body)
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

// HistoryInformation requests the history information for an fqid from the datastore.
func (s *sourceDatastore) HistoryInformation(ctx context.Context, fqid string, w io.Writer) error {
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		s.url+urlHistoryInformation,
		strings.NewReader(fmt.Sprintf(`{"fqids":[%q]}`, fqid)),
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
		// TODO LAST ERROR
		return fmt.Errorf("datastore returned %s", resp.Status)
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		// TODO External Error
		return fmt.Errorf("copping datastore response to client: %w", err)
	}

	return nil
}

// keysToGetManyRequest a json envoding of the get_many request.
func keysToGetManyRequest(keys []dskey.Key, position int) ([]byte, error) {
	request := struct {
		Requests []dskey.Key `json:"requests"`
		Position int         `json:"position,omitempty"`
	}{keys, position}
	return json.Marshal(request)
}

// parseGetManyResponse reads the response from the getMany request and
// returns the content as key-values.
func parseGetManyResponse(r io.Reader) (map[dskey.Key][]byte, error) {
	var data map[string]map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	keyValue := make(map[dskey.Key][]byte)
	for collection, idField := range data {
		for idstr, fieldValue := range idField {
			id, err := strconv.Atoi(idstr)
			if err != nil {
				// TODO LAST ERROR
				return nil, fmt.Errorf("invalid key. Id is no number: %s", idstr)
			}
			for field, value := range fieldValue {
				key, err := dskey.FromParts(collection, id, field)
				if err != nil {
					return nil, fmt.Errorf("invalid key: %w", err)
				}

				keyValue[key] = value
			}
		}
	}
	return keyValue, nil
}
