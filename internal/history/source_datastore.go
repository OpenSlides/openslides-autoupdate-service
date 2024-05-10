package history

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

var (
	envDatastoreHost     = environment.NewVariable("DATASTORE_READER_HOST", "localhost", "Host of the datastore reader.")
	envDatastorePort     = environment.NewVariable("DATASTORE_READER_PORT", "9010", "Port of the datastore reader.")
	envDatastoreProtocol = environment.NewVariable("DATASTORE_READER_PROTOCOL", "http", "Protocol of the datastore reader.")

	envDatastoreTimeout = environment.NewVariable("DATASTORE_TIMEOUT", "3s", "Time until a request to the datastore times out.")
)

const (
	urlHistoryInformation = "/internal/datastore/reader/history_information"
)

// sourceDatastore receives the data from the datastore-reader via http and
// updates via the redis message bus.
type sourceDatastore struct {
	url    string
	client *http.Client

	metricDSHitCount uint64
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

	source := sourceDatastore{
		url: url,
		client: &http.Client{
			Timeout: timeout,
		},
	}

	return &source, nil
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
