package bootstrap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

const (
	levelInfo    = "info"
	levelWarning = "warning"
	levelError   = "error"
)

func NewLogger(level string) (l *zap.SugaredLogger, err error) {
	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true

	switch strings.ToLower(level) {
	default:
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	case levelInfo:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case levelWarning:
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case levelError:
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	basicLogger, err := cfg.Build()
	l = basicLogger.Sugar()

	return
}
