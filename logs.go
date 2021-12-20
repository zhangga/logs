package logs

import (
	"context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LogOption func() error

var (
	defLogger = newDefaultLogger()
)

func newDefaultLogger() Logger {
	return NewZapLoggerWith()
}

// WithLoggerConf 通过yaml配置文件初始化logger
func WithLoggerConf(filepath string) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		Errorf("ERROR! cannot found logger config file: %v", filepath)
		return
	}
	config := newDefaultConfig()
	err = yaml.Unmarshal(file, config)
	if err != nil {
		Errorf("ERROR! cannot yaml logger config file: %v, err: %v", filepath, err)
		return
	}
	logger := &ZapLogger{
		Config: config,
	}
	logger.ApplyConfig()
	defLogger = logger
}

func SetLogLevel(level Level) {
	defLogger.SetLogLevel(level)
}

func Debug(args ...interface{}) {
	defLogger.Debug(args)
}
func Info(args ...interface{}) {
	defLogger.Info(args)
}
func Warn(args ...interface{}) {
	defLogger.Warn(args)
}
func Error(args ...interface{}) {
	defLogger.Error(args)
}
func Panic(args ...interface{}) {
	defLogger.Panic(args)
}
func Fatal(args ...interface{}) {
	defLogger.Fatal(args)
}

func Debugf(template string, args ...interface{}) {
	defLogger.Debugf(template, args)
}
func Infof(template string, args ...interface{}) {
	defLogger.Infof(template, args)
}
func Warnf(template string, args ...interface{}) {
	defLogger.Warnf(template, args)
}
func Errorf(template string, args ...interface{}) {
	defLogger.Errorf(template, args)
}
func Panicf(template string, args ...interface{}) {
	defLogger.Panicf(template, args)
}
func Fatalf(template string, args ...interface{}) {
	defLogger.Fatalf(template, args)
}

func CtxDebugf(ctx context.Context, template string, args ...interface{}) {
	defLogger.CtxDebugf(ctx, template, args)
}
func CtxInfof(ctx context.Context, template string, args ...interface{}) {
	defLogger.CtxInfof(ctx, template, args)
}
func CtxWarnf(ctx context.Context, template string, args ...interface{}) {
	defLogger.CtxWarnf(ctx, template, args)
}
func CtxErrorf(ctx context.Context, template string, args ...interface{}) {
	defLogger.CtxErrorf(ctx, template, args)
}
func CtxPanicf(ctx context.Context, template string, args ...interface{}) {
	defLogger.CtxPanicf(ctx, template, args)
}
func CtxFatalf(ctx context.Context, template string, args ...interface{}) {
	defLogger.CtxFatalf(ctx, template, args)
}
