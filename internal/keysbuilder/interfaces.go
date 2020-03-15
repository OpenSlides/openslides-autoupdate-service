package keysbuilder

import "context"

// IDer Returns relations from a key.
type IDer interface {
	ID(ctx context.Context, key string) (int, error)
	IDList(ctx context.Context, key string) ([]int, error)
	Template(ctx context.Context, key string) ([]string, error)
}

type fieldDescription interface {
	validate() error
	build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error)
}
