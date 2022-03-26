package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	port, found := os.LookupEnv("AUTOUPDATE_PORT")
	if !found {
		port = "9012"
	}

	resp, err := http.Get("http://localhost:" + port + "/system/autoupdate/health")
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("health returned status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	expect := `{"healthy": true}`
	got := strings.TrimSpace(string(body))
	if got != expect {
		return fmt.Errorf("got `%s`, expected `%s`", body, expect)
	}

	return nil
}
