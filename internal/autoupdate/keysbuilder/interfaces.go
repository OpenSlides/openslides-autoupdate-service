package keysbuilder

import "context"

// IDer Returns ids for a key with suffix _id or _ids
type IDer interface {
	IDs(ctx context.Context, key string) ([]int, error)
}
