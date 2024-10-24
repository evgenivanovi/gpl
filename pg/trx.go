package pg

import (
	"context"
	"fmt"

	"github.com/evgenivanovi/gpl/stdx/log/slog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type contextKey string

const trxCtxKey contextKey = "ctx.trx"

func IsWithinTx(ctx context.Context) bool {
	return ctx.Value(trxCtxKey) != nil
}

// injectTx injects transaction to context
func injectTx(ctx context.Context, tx *pgxpool.Tx) context.Context {
	return context.WithValue(ctx, trxCtxKey, tx)
}

// extractTx extracts transaction from context
func extractTx(ctx context.Context) *pgxpool.Tx {
	if tx, ok := ctx.Value(trxCtxKey).(*pgxpool.Tx); ok {
		return tx
	}
	return nil
}

type PgxTransactor struct {
	pool *pgxpool.Pool
}

func ProvidePgxTransactor(pool *pgxpool.Pool) *PgxTransactor {
	return &PgxTransactor{
		pool: pool,
	}
}

func (t PgxTransactor) MustStart(ctx context.Context) context.Context {
	ctx, err := t.Start(ctx)
	if err != nil {
		panic(err)
	}
	return ctx
}

func (t PgxTransactor) Start(ctx context.Context) (context.Context, error) {
	trx := extractTx(ctx)

	if trx == nil {
		tx, err := t.acquire(ctx)
		if err != nil {
			return ctx, err
		}
		return injectTx(ctx, tx), nil
	}

	return ctx, nil
}

func (t PgxTransactor) MustClose(ctx context.Context, err error) {
	if err := Close(ctx, err); err != nil {
		panic(err)
	}
}

func (t PgxTransactor) Close(ctx context.Context, err error) error {
	return Close(ctx, err)
}

func (t PgxTransactor) Within(
	ctx context.Context, in func(context.Context) error,
) error {

	trx, err := t.acquire(ctx)
	if err != nil {
		return err
	}

	err = in(injectTx(ctx, trx))
	if err != nil {
		// If rollback fails, there's nothing to do, the transaction will expire by itself
		_ = Rollback(ctx, trx)
		return err
	}

	return Commit(ctx, trx)

}

func (t PgxTransactor) acquire(ctx context.Context) (*pgxpool.Tx, error) {
	trx, err := t.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return trx.(*pgxpool.Tx), nil
}

func (t PgxTransactor) Commit(ctx context.Context) error {
	trx := extractTx(ctx)
	if trx != nil {
		return Commit(ctx, trx)
	}
	return nil
}

func (t PgxTransactor) Rollback(ctx context.Context) error {
	trx := extractTx(ctx)
	if trx != nil {
		return Rollback(ctx, trx)
	}
	return nil
}

func Close(ctx context.Context, err error) error {
	trx := extractTx(ctx)

	if trx != nil {
		if err != nil {
			return Rollback(ctx, trx)
		}
		return Commit(ctx, trx)
	}

	return nil
}

func MustClose(ctx context.Context, err error) {
	if err = Close(ctx, err); err != nil {
		panic(err)
	}
}

func Commit(ctx context.Context, trx pgx.Tx) error {
	err := trx.Commit(ctx)
	if err != nil {
		msg := fmt.Sprintf("trx.Commit failed: %s", err)
		slog.FromCtx(ctx).Debug(msg)
		return err
	}
	return nil
}

func MustCommit(ctx context.Context, trx pgx.Tx) {
	if err := Commit(ctx, trx); err != nil {
		panic(err)
	}
}

func Rollback(ctx context.Context, trx pgx.Tx) error {
	err := trx.Rollback(ctx)
	if err != nil {
		msg := fmt.Sprintf("trx.Rollback failed: %s", err)
		slog.FromCtx(ctx).Debug(msg)
		return err
	}
	return nil
}

func MustRollback(ctx context.Context, trx pgx.Tx) {
	if err := Rollback(ctx, trx); err != nil {
		panic(err)
	}
}
