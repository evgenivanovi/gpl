package slog

import (
	"log/slog"

	zapx "github.com/evgenivanovi/gpl/stdx/log/zap"
)

var logger *slog.Logger

func init() {
	logger = zapx.LogAsStructured(*zapx.Log())
}

func Log() *slog.Logger {
	return logger
}
