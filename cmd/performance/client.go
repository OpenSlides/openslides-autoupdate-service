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
// resonce of the server to the keys-channel.
// The function blocks until the connection is established.
func (c *client) connect(ctx context.Context, keys chan<- string) error {
	hc := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("can not create request: %w", err)
	}

	resp, err := hc.Do(req)
	if err != nil {
		return fmt.Errorf("can not send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			body = []byte(resp.Status)
		}
		return fmt.Errorf("server response: %s", body)
	}

	go func() {
		defer resp.Body.Close()
		buf := make([]byte, 1024)
		for {
			out := []byte{}
			for {
				n, err := resp.Body.Read(buf)
				if err != nil {
					log.Fatalf("Can not read from response body: %v", err)
				}
				out = append(out, buf[:n]...)

				if n < len(buf) {
					break
				}
			}
			keys <- string(out)
		}
	}()
	return nil
}
