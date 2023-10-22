package zap

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

/* __________________________________________________ */

var level zap.AtomicLevel
var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

/* __________________________________________________ */

func init() {
	level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger = NewZap(level)
	sugarLogger = logger.Sugar()
}

func Log() *zap.Logger {
	return logger
}

func SugarLog() *zap.SugaredLogger {
	return sugarLogger
}

func Level() zapcore.Level {
	return level.Level()
}

func SetLog(new *zap.Logger) {
	logger = new
	sugarLogger = logger.Sugar()
}

func SetSugar(new *zap.SugaredLogger) {
	sugarLogger = new
	logger = sugarLogger.Desugar()
}

func SetLevel(new zapcore.Level) {
	level.SetLevel(new)
}

/* __________________________________________________ */

func NewZap(
	lvl zapcore.LevelEnabler,
	opts ...zap.Option,
) *zap.Logger {

	if lvl == nil {
		lvl = level
	}

	config := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(config)
	core := zapcore.NewCore(encoder, os.Stdout, level)

	return zap.New(core, opts...)

}

func LogAsStructured(log zap.Logger) *slog.Logger {
	return slog.New(
		zapslog.NewHandler(
			log.Core(),
			nil,
		),
	)
}

func SugarLogAsStructured(log zap.SugaredLogger) *slog.Logger {
	return slog.New(
		zapslog.NewHandler(
			log.Desugar().Core(),
			nil,
		),
	)
}

/* __________________________________________________ */
