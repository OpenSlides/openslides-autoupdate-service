package datastore

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"golang.org/x/sync/errgroup"
)

var (
	envDatastoreHost     = environment.NewVariable("DATASTORE_READER_HOST", "localhost", "Host of the datastore reader.")
	envDatastorePort     = environment.NewVariable("DATASTORE_READER_PORT", "9010", "Port of the datastore reader.")
	envDatastoreProtocol = environment.NewVariable("DATASTORE_READER_PROTOCOL", "openslides", "Protocol of the datastore reader.")

	envDatastoreTimeout         = environment.NewVariable("DATASTORE_TIMEOUT", "3s", "Time until a request to the datastore times out.")
	envDatastoreMaxParallelKeys = environment.NewVariable("DATASTORE_MAX_PARALLEL_KEYS", "1000", "Max keys that are send in one request to the datastore.")
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
	Update(context.Context) (map[Key][]byte, error)
}

// SourceDatastore receives the data from the datastore-reader via http and
// updates via the redis message bus.
type SourceDatastore struct {
	url     string
	client  *http.Client
	updater Updater

	metricDSHitCount  uint64
	maxKeysPerRequest int
}

// NewSourceDatastore initializes a SourceDatastore.
func NewSourceDatastore(lookup environment.Getenver, updater Updater) (*SourceDatastore, []environment.Variable, error) {
	url := fmt.Sprintf(
		"%s://%s:%s",
		envDatastoreProtocol.Value(lookup),
		envDatastoreHost.Value(lookup),
		envDatastorePort.Value(lookup),
	)

	timeout, err := environment.ParseDuration(envDatastoreTimeout.Value(lookup))
	if err != nil {
		return nil, nil, fmt.Errorf("parsing timeout: %w", err)
	}

	maxParallel, err := strconv.Atoi(envDatastoreMaxParallelKeys.Value(lookup))
	if err != nil {
		return nil, nil, fmt.Errorf(
			"environment variable MAX_PARALLEL_KEYS has to be a number, not %s",
			envDatastoreMaxParallelKeys.Value(lookup),
		)
	}

	source := SourceDatastore{
		url: url,
		client: &http.Client{
			Timeout: timeout,
		},
		updater:           updater,
		maxKeysPerRequest: maxParallel,
	}

	usedEnv := []environment.Variable{
		envDatastoreHost,
		envDatastorePort,
		envDatastoreProtocol,
	}

	return &source, usedEnv, nil
}

// Get fetches the request keys from the datastore-reader.
func (s *SourceDatastore) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	atomic.AddUint64(&s.metricDSHitCount, 1)
	return s.GetPosition(ctx, 0, keys...)
}

// GetPosition gets keys from the datastore at a specifi position.
//
// Position 0 means the current position.
func (s *SourceDatastore) GetPosition(ctx context.Context, position int, keys ...Key) (map[Key][]byte, error) {
	if len(keys) <= s.maxKeysPerRequest {
		return s.getPosition(ctx, position, keys...)
	}

	// Sort keys to help datastore
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Collection < keys[j].Collection {
			return true
		} else if keys[i].Collection < keys[j].Collection {
			return false
		}

		if keys[i].ID < keys[j].ID {
			return true
		} else if keys[i].ID < keys[j].ID {
			return false
		}

		return keys[i].Field < keys[j].Field
	})

	eg, ctx := errgroup.WithContext(ctx)

	requestCount := len(keys) / s.maxKeysPerRequest
	if len(keys)%s.maxKeysPerRequest != 0 {
		requestCount++
	}

	results := make([]map[Key][]byte, requestCount)
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

	combined := make(map[Key][]byte, len(keys))
	for _, r := range results {
		for k, v := range r {
			combined[k] = v
		}
	}

	return combined, nil
}

func (s *SourceDatastore) getPosition(ctx context.Context, position int, keys ...Key) (map[Key][]byte, error) {
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

// Update updates the data from the redis message bus.
func (s *SourceDatastore) Update(ctx context.Context) (map[Key][]byte, error) {
	return s.updater.Update(ctx)
}

// HistoryInformation requests the history information for an fqid from the datastore.
func (s *SourceDatastore) HistoryInformation(ctx context.Context, fqid string, w io.Writer) error {
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
