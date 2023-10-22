package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WriteRequester struct {
	pool *pgxpool.Pool
}

func ProvideWriteRequester(pool *pgxpool.Pool) *WriteRequester {
	return &WriteRequester{
		pool: pool,
	}
}

// ExecWithDefaultOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) ExecWithDefaultOnError(
	ctx context.Context,
	command string, args ...any,
) error {
	return r.ExecWithOnError(
		ctx, TranslateWriteErrorFunc(), command, args...,
	)
}

// ExecWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) ExecWithOnError(
	ctx context.Context,
	onError func(error) error,
	command string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecTxWithOnError(ctx, tx, onError, command, args...)
	} else {
		return r.doExecWithOnError(ctx, onError, command, args...)
	}

}

// doExecTxWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) doExecTxWithOnError(
	ctx context.Context,
	tx *pgxpool.Tx,
	onError func(error) error,
	command string, args ...any,
) error {
	_, err := tx.Exec(ctx, command, args...)
	return onError(err)
}

// doExecWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) doExecWithOnError(
	ctx context.Context,
	onError func(error) error,
	command string, args ...any,
) error {
	_, err := r.pool.Exec(ctx, command, args...)
	return onError(err)
}

// ExecReturningWithDefaultOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) ExecReturningWithDefaultOnError(
	ctx context.Context,
	dst any, command string, args ...any,
) error {
	return r.ExecReturningWithOnError(
		ctx, TranslateWriteErrorFunc(), dst, command, args...,
	)
}

// ExecReturningWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) ExecReturningWithOnError(
	ctx context.Context,
	onError func(error) error,
	dst any, command string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecTxReturningWithOnError(ctx, tx, onError, dst, command, args...)
	} else {
		return r.doExecReturningWithOnError(ctx, onError, dst, command, args...)
	}

}

// ExecReturningWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) doExecTxReturningWithOnError(
	ctx context.Context,
	tx *pgxpool.Tx,
	onError func(error) error,
	dst any, command string, args ...any,
) error {
	row := tx.QueryRow(ctx, command, args...)
	err := row.Scan(dst)
	return onError(err)
}

// ExecReturningWithOnError
// https://github.com/jackc/pgx/issues/411#issuecomment-395987764
func (r *WriteRequester) doExecReturningWithOnError(
	ctx context.Context,
	onError func(error) error,
	dst any, command string, args ...any,
) error {
	row := r.pool.QueryRow(ctx, command, args...)
	err := row.Scan(dst)
	return onError(err)
}
