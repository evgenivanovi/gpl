package xsql

import "context"

type Transactor interface {
	Within(ctx context.Context, fn func(ctx context.Context) error) error
}
