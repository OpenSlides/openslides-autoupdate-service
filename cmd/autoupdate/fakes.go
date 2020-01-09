package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
)

// fakeReceiver implements the Receiver interface. It reads on a Reader, for example stdin and
// takes each word on each line as changed key.
type faker struct {
	buf  *bufio.Reader
	data map[string][]byte
}

func (r faker) KeysChanged() (autoupdate.KeyChanges, error) {
	kc := autoupdate.KeyChanges{}
	msg, err := r.buf.ReadString('\n')
	if err == io.EOF {
		// Don't return anything (block forever) if the reader is empty.
		select {}
	}
	if err != nil {
		return kc, fmt.Errorf("can not read from buffer: %v", err) //TODO: in %w ändern
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
	kc.Updated = keys

	return kc, nil
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

func (r faker) IDsFromKey(ctx context.Context, uid int, mid int, key string) ([]int, error) {
	o := r.data[key]
	if len(o) != 0 {
		var id int
		if err := json.Unmarshal(o, &id); err != nil {
			var ids []int
			if err := json.Unmarshal(o, &ids); err != nil {
				return nil, fmt.Errorf("Invalid value %s for field %s", o, key)
			}
			return ids, nil
		}
		return []int{id}, nil
	}
	if strings.HasPrefix(key, "not_exist") {
		return nil, nil
	}
	if strings.HasSuffix(key, "_id") {
		return []int{1}, nil
	}
	if !strings.HasSuffix(key, "_ids") {
		return nil, fmt.Errorf("Key %s can not be a reference; expected suffex _id or _ids", key)
	}
	if mid == 1 {
		return []int{1}, nil
	}
	return []int{1, 2}, nil
}

func (r faker) IDsFromCollection(ctx context.Context, uid int, mid int, collection string) ([]int, error) {
	if mid == 1 {
		return []int{1}, nil
	}
	return []int{1, 2}, nil
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
