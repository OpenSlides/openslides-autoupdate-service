package test

import (
	"context"
	"sync"
)

//MockPermission mocks the permission api.
type MockPermission struct {
	mu      sync.Mutex
	Data    map[string]bool
	Called  map[string]bool
	Default bool
}

// RestrictFQFields returns the fields where p.Data is true.
func (p *MockPermission) RestrictFQFields(ctx context.Context, uid int, fqids []string) (map[string]bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.Data == nil {
		p.Data = make(map[string]bool)
	}
	if p.Called == nil {
		p.Called = make(map[string]bool)
	}

	out := make(map[string]bool)
	var ok bool
	for _, k := range fqids {
		out[k], ok = p.Data[k]
		if !ok {
			out[k] = p.Default
		}
		p.Called[k] = true
	}

	return out, nil
}
