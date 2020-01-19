package keysbuilder

import "context"

// Restricter restricts keys. See autoupdate.Restricter
type Restricter interface {
	Restrict(ctx context.Context, uid int, keys []string) (map[string][]byte, error)
	IDsFromKey(ctx context.Context, uid int, key string) ([]int, error)
}
