package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// fakeReceiver implements the Receiver interface. It reads on a Reader, for example stdin and
// takes each word on each line as changed key.
type faker struct {
	buf  *bufio.Reader
	data map[string][]byte
}

func (r faker) KeysChanged() ([]string, error) {
	msg, err := r.buf.ReadString('\n')
	if err == io.EOF {
		// Don't return anything (block forever) if the reader is empty.
		select {}
	}
	if err != nil {
		return nil, fmt.Errorf("can not read from buffer: %v", err) //TODO: in %w ändern
	}

	data := strings.Split(strings.TrimSpace(msg), " ")
	keys := []string{}
	for _, d := range data {
		keyValue := strings.SplitN(d, "=", 2)
		if len(keyValue) == 1 {
			keyValue = append(keyValue, "some new value")
		}
		keys = append(keys, keyValue[0])
		r.data[keyValue[0]] = []byte(keyValue[1])
	}
	return keys, nil
}

func (r faker) Restrict(ctx context.Context, uid int, keys []string) (map[string][]byte, error) {
	out := make(map[string][]byte, len(keys))
	for _, key := range keys {
		o := r.data[key]
		if len(o) != 0 {
			out[key] = o
			continue
		}
		switch {
		case strings.HasSuffix(key, "_id"):
			out[key] = []byte("1")
		case strings.HasSuffix(key, "_ids"):
			out[key] = []byte("[1,2]")
		default:
			out[key] = []byte("some data")
		}
	}
	return out, nil
}

// fake Auth implements the Authenticater interface. It always returns 1.
type fakeAuth struct {
	noConnection bool
}

func (a fakeAuth) Authenticate(ctx context.Context, _ *http.Request) (int, error) {
	if a.noConnection {
		<-ctx.Done()
		return 0, fmt.Errorf("can not connect to auth service")
	}
	return 1, nil
}
