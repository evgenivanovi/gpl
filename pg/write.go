package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

/* __________________________________________________ */

type WriteRequester struct {
	pool *pgxpool.Pool
}

func ProvideWriteRequester(pool *pgxpool.Pool) *WriteRequester {
	return &WriteRequester{
		pool: pool,
	}
}

/* __________________________________________________ */

// ExecuteWithDefaultOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) ExecuteWithDefaultOnError(
	ctx context.Context,
	command string, args ...any,
) error {
	return r.ExecuteWithOnError(
		ctx, TranslateWriteErrorFunc(), command, args...,
	)
}

// ExecuteWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) ExecuteWithOnError(
	ctx context.Context,
	onError func(error) error,
	command string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecuteTxWithOnError(ctx, tx, onError, command, args...)
	} else {
		return r.doExecuteWithOnError(ctx, onError, command, args...)
	}

}

// doExecuteTxWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) doExecuteTxWithOnError(
	ctx context.Context,
	tx *pgxpool.Tx,
	onError func(error) error,
	command string, args ...any,
) error {
	_, err := tx.Exec(ctx, command, args...)
	return onError(err)
}

// doExecuteWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) doExecuteWithOnError(
	ctx context.Context,
	onError func(error) error,
	command string, args ...any,
) error {
	_, err := r.pool.Exec(ctx, command, args...)
	return onError(err)
}

/* __________________________________________________ */

// ExecuteReturningWithDefaultOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) ExecuteReturningWithDefaultOnError(
	ctx context.Context,
	dst any, command string, args ...any,
) error {
	return r.ExecuteReturningWithOnError(
		ctx, TranslateWriteErrorFunc(), dst, command, args...,
	)
}

// ExecuteReturningWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) ExecuteReturningWithOnError(
	ctx context.Context,
	onError func(error) error,
	dst any, command string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecuteTxReturningWithOnError(ctx, tx, onError, dst, command, args...)
	} else {
		return r.doExecuteReturningWithOnError(ctx, onError, dst, command, args...)
	}

}

// ExecuteReturningWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) doExecuteTxReturningWithOnError(
	ctx context.Context,
	tx *pgxpool.Tx,
	onError func(error) error,
	dst any, command string, args ...any,
) error {
	row := tx.QueryRow(ctx, command, args...)
	err := row.Scan(dst)
	return onError(err)
}

// ExecuteReturningWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) doExecuteReturningWithOnError(
	ctx context.Context,
	onError func(error) error,
	dst any, command string, args ...any,
) error {
	row := r.pool.QueryRow(ctx, command, args...)
	err := row.Scan(dst)
	return onError(err)
}

/* __________________________________________________ */
