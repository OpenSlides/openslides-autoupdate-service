package keysbuilder

import "context"

// IDer restricts keys. See autoupdate.IDer
type IDer interface {
	IDs(ctx context.Context, key string) ([]int, error)
}
