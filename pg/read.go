package pg

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReadRequester struct {
	pool *pgxpool.Pool
}

func ProvideReadRequester(pool *pgxpool.Pool) *ReadRequester {
	return &ReadRequester{
		pool: pool,
	}
}

// ExecOne
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) ExecOne(
	ctx context.Context,
	dst any, query string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecOneTx(ctx, tx, dst, query, args...)
	} else {
		return r.doExecOne(ctx, dst, query, args...)
	}

}

// doExecOneTx
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecOneTx(
	ctx context.Context,
	tx *pgxpool.Tx, dst any, query string, args ...any,
) error {

	rows, err := tx.Query(ctx, query, args...)

	if err != nil {
		return TranslateReadError(err)
	}

	err = pgxscan.ScanOne(dst, rows)
	if err != nil {
		return TranslateReadError(err)
	}

	return nil

}

// ExecOne
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecOne(
	ctx context.Context,
	dst any, query string, args ...any,
) error {

	rows, err := r.pool.Query(ctx, query, args...)

	if err != nil {
		return TranslateReadError(err)
	}

	err = pgxscan.ScanOne(dst, rows)
	if err != nil {
		return TranslateReadError(err)
	}

	return nil

}

// ExecOneWithScan
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) ExecOneWithScan(
	ctx context.Context,
	scan func(row pgx.Row) error, query string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecOneTxWithScan(ctx, tx, scan, query, args...)
	} else {
		return r.doExecOneWithScan(ctx, scan, query, args...)
	}

}

// doExecOneTx
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecOneTxWithScan(
	ctx context.Context,
	tx *pgxpool.Tx, scan func(row pgx.Row) error, query string, args ...any,
) error {

	row := tx.QueryRow(ctx, query, args...)

	err := scan(row)
	if err != nil {
		return TranslateReadError(err)
	}

	return nil

}

// ExecOne
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecOneWithScan(
	ctx context.Context,
	scan func(row pgx.Row) error, query string, args ...any,
) error {

	row := r.pool.QueryRow(ctx, query, args...)

	err := scan(row)
	if err != nil {
		return TranslateReadError(err)
	}

	return nil

}

// ExecMany
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) ExecMany(
	ctx context.Context,
	dst any, query string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecManyTx(ctx, tx, dst, query, args...)
	} else {
		return r.doExecMany(ctx, dst, query, args...)
	}

}

// doExecManyTx
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecManyTx(
	ctx context.Context,
	tx *pgxpool.Tx, dst any, query string, args ...any,
) error {

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return TranslateReadError(err)
	}

	err = pgxscan.ScanAll(dst, rows)
	if err != nil {
		return TranslateReadError(err)
	}

	return nil

}

// doExecManyTx
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecMany(
	ctx context.Context,
	dst any, query string, args ...any,
) error {

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return TranslateReadError(err)
	}

	err = pgxscan.ScanAll(dst, rows)
	if err != nil {
		return TranslateReadError(err)
	}

	return nil

}

// ExecManyWithScan
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) ExecManyWithScan(
	ctx context.Context,
	scan func(rows pgx.Rows) error, query string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecManyTxWithScan(ctx, tx, scan, query, args...)
	} else {
		return r.doExecManyWithScan(ctx, scan, query, args...)
	}

}

// doExecManyTxWithScan
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecManyTxWithScan(
	ctx context.Context,
	tx *pgxpool.Tx, scan func(rows pgx.Rows) error, query string, args ...any,
) error {

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return TranslateReadError(err)
	}

	innerScan := func(rows pgx.Rows) error {
		defer rows.Close()
		return scan(rows)
	}

	err = innerScan(rows)
	if err != nil {
		return TranslateReadError(err)
	}

	return nil

}

// doExecManyWithScan
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecManyWithScan(
	ctx context.Context,
	scan func(rows pgx.Rows) error, query string, args ...any,
) error {

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return TranslateReadError(err)
	}

	innerScan := func(rows pgx.Rows) error {
		defer rows.Close()
		return scan(rows)
	}

	err = innerScan(rows)
	if err != nil {
		return TranslateReadError(err)
	}

	return nil

}
