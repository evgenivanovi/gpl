package pg

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

/* __________________________________________________ */

type ReadRequester struct {
	pool *pgxpool.Pool
}

func ProvideReadRequester(pool *pgxpool.Pool) *ReadRequester {
	return &ReadRequester{
		pool: pool,
	}
}

/* __________________________________________________ */

// ExecuteOne
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) ExecuteOne(
	ctx context.Context,
	dst any, query string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecuteOneTx(ctx, tx, dst, query, args...)
	} else {
		return r.doExecuteOne(ctx, dst, query, args...)
	}

}

// doExecuteOneTx
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecuteOneTx(
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

// ExecuteOne
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecuteOne(
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

/* __________________________________________________ */

// ExecuteOneWithScan
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) ExecuteOneWithScan(
	ctx context.Context,
	scan func(row pgx.Row) error, query string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecuteOneTxWithScan(ctx, tx, scan, query, args...)
	} else {
		return r.doExecuteOneWithScan(ctx, scan, query, args...)
	}

}

// doExecuteOneTx
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecuteOneTxWithScan(
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

// ExecuteOne
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecuteOneWithScan(
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

/* __________________________________________________ */

// ExecuteMany
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) ExecuteMany(
	ctx context.Context,
	dst any, query string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecuteManyTx(ctx, tx, dst, query, args...)
	} else {
		return r.doExecuteMany(ctx, dst, query, args...)
	}

}

// doExecuteManyTx
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecuteManyTx(
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

// doExecuteManyTx
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecuteMany(
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

/* __________________________________________________ */

// ExecuteManyWithScan
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) ExecuteManyWithScan(
	ctx context.Context,
	scan func(rows pgx.Rows) error, query string, args ...any,
) error {

	tx := extractTx(ctx)

	if tx != nil {
		return r.doExecuteManyTxWithScan(ctx, tx, scan, query, args...)
	} else {
		return r.doExecuteManyWithScan(ctx, scan, query, args...)
	}

}

// doExecuteManyTxWithScan
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecuteManyTxWithScan(
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

// doExecuteManyWithScan
// Arguments should be referenced positionally from the SQL string as $1, $2, etc.
func (r *ReadRequester) doExecuteManyWithScan(
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

/* __________________________________________________ */
