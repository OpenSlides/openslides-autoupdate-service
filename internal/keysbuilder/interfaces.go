package keysbuilder

import "context"

// Valuer decodes a restricted value for an key.
type Valuer interface {
	Value(ctx context.Context, uid int, key string, value interface{}) error
}

type fieldDescription interface {
	build(ctx context.Context, valuer Valuer, uid int, key string, keys chan<- string, errs chan<- error)
}

type keyDoesNotExister interface {
	KeyDoesNotExist() bool
}

type preloader interface {
	LoadKeys(context.Context, ...string) error
}
