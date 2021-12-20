package logs

import "context"

type Logger interface {
	SetLogLevel(level Level)

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})

	CtxDebugf(ctx context.Context, template string, args ...interface{})
	CtxInfof(ctx context.Context, template string, args ...interface{})
	CtxWarnf(ctx context.Context, template string, args ...interface{})
	CtxErrorf(ctx context.Context, template string, args ...interface{})
	CtxPanicf(ctx context.Context, template string, args ...interface{})
	CtxFatalf(ctx context.Context, template string, args ...interface{})
}
