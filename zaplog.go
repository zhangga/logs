package logs

import (
	"context"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"time"
)

// ZapLogger Logger implements
type ZapLogger struct {
	Config *Config
	logger *zap.SugaredLogger
}

func NewZapLoggerWith(opts ...LogOption) Logger {
	logger := &ZapLogger{
		Config: newDefaultConfig(),
	}
	logger.ApplyConfig()
	return logger
}

func (l *ZapLogger) ApplyConfig() {
	conf := l.Config
	cores := []zapcore.Core{}

	var encoder zapcore.Encoder

	if conf.JsonFormat {
		encoder = zapcore.NewJSONEncoder(getEncoder())
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoder())
	}

	conf.AtomicLevel.SetLevel(getLevel(conf.DefaultLogLevel))

	if conf.ConsoleOut {
		writer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(encoder, writer, conf.AtomicLevel)
		cores = append(cores, core)
	}

	if conf.FileOut.Enable {
		fileWriter := getFileWriter(
			conf.FileOut.Path,
			conf.FileOut.Name,
			conf.FileOut.RotationTime,
			conf.FileOut.RotationCount,
		)
		writer := zapcore.AddSync(fileWriter)
		core := zapcore.NewCore(encoder, writer, conf.AtomicLevel)
		cores = append(cores, core)
	}

	combinedCore := zapcore.NewTee(cores...)

	logger := zap.New(combinedCore,
		zap.AddCallerSkip(conf.CallerSkip),
		zap.AddStacktrace(getLevel(conf.StacktraceLevel)),
		zap.AddCaller(),
	)

	if conf.ProjectName != "" {
		logger = logger.Named(conf.ProjectName)
	}

	defer logger.Sync()

	l.logger = logger.Sugar()
}

func getFileWriter(path, name string, rotationTime, rotationCount uint) io.Writer {
	writer, err := rotatelogs.New(
		filepath.Join(path, name+".%Y%m%d%H.log"),
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 日志切割时间间隔
		rotatelogs.WithRotationCount(rotationCount),                        // 文件最大保存份数
	)
	if err != nil {
		panic(err)
	}
	return writer
}

func getEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		LevelKey:       "L",
		TimeKey:        "T",
		MessageKey:     "M",
		NameKey:        "N",
		CallerKey:      "C",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func (l *ZapLogger) SetLogLevel(level Level) {
	l.Config.SetLevel(level)
}

func (l *ZapLogger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *ZapLogger) CtxDebugf(ctx context.Context, template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *ZapLogger) Info(args ...interface{}) {
	l.logger.Info(args)
}

func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *ZapLogger) CtxInfof(ctx context.Context, template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *ZapLogger) Warn(args ...interface{}) {
	l.logger.Warn(args)
}

func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *ZapLogger) CtxWarnf(ctx context.Context, template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *ZapLogger) Error(args ...interface{}) {
	l.logger.Error(args)
}

func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *ZapLogger) CtxErrorf(ctx context.Context, template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *ZapLogger) Panic(args ...interface{}) {
	l.logger.Panic(args)
}

func (l *ZapLogger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args)
}

func (l *ZapLogger) CtxPanicf(ctx context.Context, template string, args ...interface{}) {
	l.logger.Panicf(template, args)
}

func (l *ZapLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args)
}

func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

func (l *ZapLogger) CtxFatalf(ctx context.Context, template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}
