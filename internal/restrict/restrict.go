// Package restrict implements the autoupdate.Restricter interface by calleng in
// restricter in the openslides-backend service
package restrict

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	method  = "POST"
	urlPath = "/system/api/restrictions"
)

// Service holds the state of the restricter
type Service struct {
	client http.Client
	Addr   string
}

// Restrict returns the values for some keys for an user id
func (s *Service) Restrict(ctx context.Context, uid int, keys []string) (io.Reader, error) {
	var buf bytes.Buffer
	reqData := []struct {
		UID  int      `json:"user_id"`
		Keys []string `json:"fqfields"`
	}{
		{UID: uid, Keys: keys},
	}

	if err := json.NewEncoder(&buf).Encode(reqData); err != nil {
		return nil, fmt.Errorf("can not build json: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, s.Addr+urlPath, &buf)
	if err != nil {
		return nil, fmt.Errorf("can not build request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can not connect to backend: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		content, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("backend responded with status %s: %s", resp.Status, content)
	}

	var respData []json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("can not decode response body: %w", err)
	}

	if len(respData) == 0 {
		return nil, fmt.Errorf("backend did not return any data")
	}
	return bytes.NewReader(respData[0]), nil
}
