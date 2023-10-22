package pg

import (
	"context"
	"fmt"

	"github.com/evgenivanovi/gpl/stdx/log/slog"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

/* __________________________________________________ */

type TrxRequester struct {
	pool *pgxpool.Pool
}

func ProvideTrxRequester(pool *pgxpool.Pool) *TrxRequester {
	return &TrxRequester{
		pool: pool,
	}
}

func (t TrxRequester) StartEx(ctx context.Context) context.Context {
	ctx, err := t.Start(ctx)
	if err != nil {
		panic(err)
	}
	return ctx
}

func (t TrxRequester) Start(ctx context.Context) (context.Context, error) {
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

func (t TrxRequester) CloseEx(ctx context.Context, err error) {
	if err := Close(ctx, err); err != nil {
		panic(err)
	}
}

func (t TrxRequester) Close(ctx context.Context, err error) error {
	return Close(ctx, err)
}

func (t TrxRequester) Within(
	ctx context.Context, in func(context.Context) error,
) error {

	trx, err := t.acquire(ctx)
	if err != nil {
		return err
	}

	err = in(injectTx(ctx, trx))
	if err != nil {
		_ = Rollback(ctx, trx)
		return err
	}

	return Commit(ctx, trx)

}

func (t TrxRequester) acquire(ctx context.Context) (*pgxpool.Tx, error) {
	trx, err := t.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return trx.(*pgxpool.Tx), nil
}

func (t TrxRequester) Commit(ctx context.Context) error {
	trx := extractTx(ctx)
	if trx != nil {
		return Commit(ctx, trx)
	}
	return nil
}

func (t TrxRequester) Rollback(ctx context.Context) error {
	trx := extractTx(ctx)
	if trx != nil {
		return Rollback(ctx, trx)
	}
	return nil
}

/* __________________________________________________ */

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

func CloseEx(ctx context.Context, err error) {
	if err = Close(ctx, err); err != nil {
		panic(err)
	}
}

/* __________________________________________________ */

func Commit(ctx context.Context, trx pgx.Tx) error {
	err := trx.Commit(ctx)
	if err != nil {
		msg := fmt.Sprintf("trx.Commit failed: %s", err)
		slog.Log().Debug(msg)
		return err
	}
	return nil
}

func CommitEx(ctx context.Context, trx pgx.Tx) {
	if err := Commit(ctx, trx); err != nil {
		panic(err)
	}
}

/* __________________________________________________ */

func Rollback(ctx context.Context, trx pgx.Tx) error {
	err := trx.Rollback(ctx)
	if err != nil {
		msg := fmt.Sprintf("trx.Rollback failed: %s", err)
		slog.Log().Debug(msg)
		return err
	}
	return nil
}

func RollbackEx(ctx context.Context, trx pgx.Tx) {
	if err := Rollback(ctx, trx); err != nil {
		panic(err)
	}
}

/* __________________________________________________ */
