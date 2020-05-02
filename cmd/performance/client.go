package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// client holds one http connection to the autoupdate service.
type client struct{}

// connect creates a new connection to the autoupdate service. It returns the
// responce of the server to the given keys-channel. The function blocks until
// the connection is established. It is held open in the beckgrond.
func (c *client) connect(ctx context.Context, keys chan<- string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("server returned %s", resp.Status)
		}
		return fmt.Errorf("server returned %s: %s", resp.Status, body)
	}

	go func() {
		defer resp.Body.Close()

		buf := make([]byte, 1024)
		for {
			var data []byte
			for {
				n, err := resp.Body.Read(buf)
				if err != nil {
					log.Fatalf("Can not read from response body: %v", err)
				}
				data = append(data, buf[:n]...)

				if n < len(buf) {
					break
				}
			}
			keys <- string(data)
		}
	}()
	return nil
}
