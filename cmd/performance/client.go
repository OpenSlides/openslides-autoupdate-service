package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type client struct{}

func (c *client) connect(ctx context.Context, keys chan<- string) error {
	hc := http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, requestBody())
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
		return fmt.Errorf("Server response: %s", body)
	}
	go func() {
		defer resp.Body.Close()
		buf := make([]byte, 1024)
		for {
			out := []byte{}
			for {
				n, err := resp.Body.Read(buf)
				if err != nil {
					log.Fatalf("Read error: %v", err)
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

func requestBody() *strings.Reader {
	return strings.NewReader(`[{"ids": [5], "collection": "user", "fields": {"name": null}}]`)
}
