package slog

import (
	"log/slog"

	zapx "github.com/evgenivanovi/gpl/stdx/log/zap"
)

/* __________________________________________________ */

var logger *slog.Logger

/* __________________________________________________ */

func init() {
	logger = zapx.LogAsStructured(*zapx.Log())
}

func Log() *slog.Logger {
	return logger
}
