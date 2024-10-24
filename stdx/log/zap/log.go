package zap

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

var level zap.AtomicLevel
var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger = NewZap(level)
	sugar = logger.Sugar()
}

func Log() *zap.Logger {
	return logger
}

func SugarLog() *zap.SugaredLogger {
	return sugar
}

func Level() zapcore.Level {
	return level.Level()
}

func SetLog(new *zap.Logger) {
	logger = new
	sugar = logger.Sugar()
}

func SetSugar(new *zap.SugaredLogger) {
	sugar = new
	logger = sugar.Desugar()
}

func SetLevel(new zapcore.Level) {
	level.SetLevel(new)
}

func NewZap(
	lvl zapcore.LevelEnabler,
	ops ...zap.Option,
) *zap.Logger {
	if lvl == nil {
		lvl = level
	}

	config := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(config)
	core := zapcore.NewCore(encoder, os.Stdout, level)

	return zap.New(core, ops...)
}

func LogAsStructured(log zap.Logger) *slog.Logger {
	return slog.New(
		zapslog.NewHandler(
			log.Core(),
		),
	)
}

func SugarLogAsStructured(log zap.SugaredLogger) *slog.Logger {
	return slog.New(
		zapslog.NewHandler(
			log.Desugar().Core(),
		),
	)
}
