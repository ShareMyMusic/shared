package logger

import (
	"context"

	"github.com/sharemymusic/shared/pkg/env"
	"go.uber.org/zap"
)

type loggerKey struct{}

// New returns a configured *zap.Logger for the given env.Env.
// For env.Development it creates a development logger; for env.Production it
// currently also creates a development logger. The created logger is set as
// the global logger via zap.ReplaceGlobals and then returned. New will panic
// if an unsupported env.Env value is provided.
func New(envType env.Env) *zap.Logger {
	var l *zap.Logger

	switch envType {
	case env.Development:
		l = zap.Must(zap.NewDevelopment())
	case env.Production:
		l = zap.Must(zap.NewDevelopment())
	default:
		panic("invalid env")
	}

	zap.ReplaceGlobals(l)
	return l
}

// CreateContext returns a derived context that carries a logger created from the
// provided env. The logger is stored under the package's unexported loggerKey;
// the original ctx is not modified and a new context containing the logger is
// returned.
func CreateContext(ctx context.Context, env env.Env) context.Context {
	ctx = context.WithValue(ctx, loggerKey{}, New(env))

	return ctx
}

// FromContext retrieves a *zap.Logger from ctx using the unexported loggerKey{}.
// If the value is present and of the correct type it is returned. If not, the
// global logger is warned and a no-op logger (zap.NewNop()) is returned so
// callers always receive a valid, non-nil *zap.Logger.
func FromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(loggerKey{}).(*zap.Logger)

	if !ok {
		zap.L().Warn("logger not found in context")
		return zap.NewNop()
	}

	return logger
}
