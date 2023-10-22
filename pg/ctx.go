package pg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type contextKey string

const trxCtxKey contextKey = "ctx.trx"

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
