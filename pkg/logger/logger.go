package logger

import (
	"context"

	"github.com/sharemymusic/shared/pkg/env"
	"go.uber.org/zap"
)

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

func CreateContext(ctx context.Context, env env.Env) context.Context {
	ctx = context.WithValue(ctx, "logger", New(env))

	return ctx
}

func FromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value("logger").(*zap.Logger)

	if !ok {
		zap.L().Warn("logger not found in context")
		return zap.NewNop()
	}

	return logger
}
