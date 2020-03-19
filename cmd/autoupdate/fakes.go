package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// fakeReceiver implements the Receiver interface. It reads on a Reader, for example stdin and
// takes each word on each line as changed key.
type faker struct {
	buf  *bufio.Reader
	data map[string]string
}

func (r faker) KeysChanged() ([]string, error) {
	msg, err := r.buf.ReadString('\n')
	if err == io.EOF {
		// Don't return anything (block forever) if the reader is empty.
		select {}
	}
	if err != nil {
		return nil, fmt.Errorf("can not read from buffer: %w", err)
	}

	data := strings.Split(strings.TrimSpace(msg), " ")
	keys := []string{}
	for _, d := range data {
		keyValue := strings.SplitN(d, "=", 2)
		if len(keyValue) == 1 {
			keyValue = append(keyValue, fmt.Sprintf(`"The time is: %s"`, time.Now()))
		}
		keys = append(keys, keyValue[0])
		r.data[keyValue[0]] = keyValue[1]
	}
	return keys, nil
}

func (r faker) Restrict(ctx context.Context, uid int, keys []string) (map[string]string, error) {
	out := make(map[string]string, len(keys))
	for _, key := range keys {
		o := r.data[key]
		if len(o) != 0 {
			out[key] = o
			continue
		}
		switch {
		case strings.HasSuffix(key, "_id"):
			out[key] = "1"
		case strings.HasSuffix(key, "_ids"):
			out[key] = "[1,2]"
		default:
			out[key] = fmt.Sprintf(`"The time is: %s"`, time.Now())
		}
	}
	return out, nil
}

// fake Auth implements the Authenticater interface. It always returns the given number.
type fakeAuth int

func (a fakeAuth) Authenticate(ctx context.Context, _ *http.Request) (int, error) {
	return int(a), nil
}
