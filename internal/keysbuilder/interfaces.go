package keysbuilder

import "context"

// IDer returns relations from a key.
type IDer interface {
	ID(ctx context.Context, key string) (int, error)
	IDList(ctx context.Context, key string) ([]int, error)
	GenericID(ctx context.Context, key string) (string, error)
	GenericIDs(ctx context.Context, key string) ([]string, error)
	Template(ctx context.Context, key string) ([]string, error)
}

type fieldDescription interface {
	build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error)
}
