package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Level = zapcore.Level

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)

type Config struct {
	DefaultLogLevel string          `yaml:"LogLevel"`        //默认日志记录级别
	StacktraceLevel string          `yaml:"StacktraceLevel"` //记录堆栈的级别
	AtomicLevel     zap.AtomicLevel //用于动态更改日志记录级别
	ProjectName     string          `yaml:"ProjectName"` //项目名称
	CallerSkip      int             `yaml:"CallerSkip"`  //CallerSkip次数
	JsonFormat      bool            `yaml:"JsonFormat"`  //输出json格式
	ConsoleOut      bool            `yaml:"ConsoleOut"`  //是否输出到console
	FileOut         *fileOut        `yaml:"FileOut"`
}

type fileOut struct {
	Enable        bool   `yaml:"Enable"`        //是否将日志输出到文件
	Path          string `yaml:"Path"`          //日志保存路径
	Name          string `yaml:"Name"`          //日志保存的名称，不写随机生成
	RotationTime  uint   `yaml:"RotationTime"`  //日志切割时间间隔(小时)
	RotationCount uint   `yaml:"RotationCount"` //文件最大保存份数
}

func newDefaultConfig() *Config {
	return &Config{
		DefaultLogLevel: "info",
		StacktraceLevel: "panic",
		AtomicLevel:     zap.NewAtomicLevel(),
		ProjectName:     "",
		CallerSkip:      1,
		JsonFormat:      false,
		ConsoleOut:      true,
		FileOut: &fileOut{
			Enable:        true,
			Path:          "./logs/",
			Name:          "log",
			RotationTime:  24,
			RotationCount: 7,
		},
	}
}

func getLevel(level string) Level {
	switch strings.ToLower(level) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	default:
		return InfoLevel
	}
}

/*SetLevel 设置日志记录级别*/
func (c *Config) SetLevel(level Level) {
	c.AtomicLevel.SetLevel(level)
}

/*SetStacktraceLevel 设置堆栈跟踪的日志级别*/
func (c *Config) SetStacktraceLevel(level string) {
	c.StacktraceLevel = level
}

/*SetProjectName 设置ProjectName*/
func (c *Config) SetProjectName(projectName string) {
	c.ProjectName = projectName
}

/*SetCallerSkip 设置callerSkip次数*/
func (c *Config) SetCallerSkip(callerSkip int) {
	c.CallerSkip = callerSkip
}

/*EnableJSONFormat 开启JSON格式化输出*/
func (c *Config) EnableJSONFormat() {
	c.JsonFormat = true
}

/*DisableJSONFormat 关闭JSON格式化输出*/
func (c *Config) DisableJSONFormat() {
	c.JsonFormat = false
}

/*EnableConsoleOut 开启Console输出*/
func (c *Config) EnableConsoleOut() {
	c.ConsoleOut = true
}

/*DisableConsoleOut 关闭Console输出*/
func (c *Config) DisableConsoleOut() {
	c.ConsoleOut = false
}

/*SetFileOut 设置日志输出文件*/
func (c *Config) SetFileOut(path, name string, rotationTime, rotationCount uint) {
	c.FileOut.Enable = true
	c.FileOut.Path = path
	c.FileOut.Name = name
	c.FileOut.RotationTime = rotationTime
	c.FileOut.RotationCount = rotationCount
}
